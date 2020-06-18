package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NginxLogCombined is typical log format, set by default:
/*
# nginx.conf
http {
  ...
  log_format combined '$remote_addr - $remote_user [$time_local] '
                      '"$request" $status $body_bytes_sent '
                      '"$http_referer" "$http_user_agent"';
  ...
}
*/
type NginxLogCombined struct {
	RemoteAddr    string
	RemoteUser    string
	TimeLocal     time.Time
	Request       string
	ResponseCode  int
	BodyBytesSent int
	Referrer      string
	UserAgent     string
}

// TimeLayout represents nginx time in epoch go format
const TimeLayout = "[02/Jan/2006:15:04:05 +0000]"

// cleanOutNginxPodLog removes all pod-infra logs and left only access.log
func cleanOutNginxPodLog(filepath string) (string, error) {
	var txtlines string

	// reading file
	file, err := os.Open(filepath)
	if err != nil {
		return txtlines, err
	}

	// init scanner to go line-by-line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// we need to keep lines that started w/ IP, e.i. related to access log
	ipDetectedRegex := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).*`
	re := regexp.MustCompile(ipDetectedRegex)

	// go line by line
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			txtlines += "\n" + scanner.Text()
		} else {
			fmt.Printf("Skipping non-access logs from pod logs: %v \n", scanner.Text())
		}
	}
	file.Close()

	return txtlines, nil
}

//parseNginxLog gets planin data as string and returns objects of Nginx Log
func parseNginxLog(plainData string) ([]NginxLogCombined, error) {
	result := []NginxLogCombined{}

	r := csv.NewReader(strings.NewReader(plainData))
	// use whitespace as delimeter
	r.Comma = ' '
	r.LazyQuotes = true
	records, err := r.ReadAll()
	if err != nil {
		return result, err
	}

	for _, r := range records {
		// dirty hack to because time log contains whitespace that reacts on csv delimeter
		//		combineTimeLog := strings.ReplaceAll(r[3], "[", "") + " " + strings.ReplaceAll(r[4], "]", "")
		combineTimeLog := r[3] + " " + r[4]

		// transforming str data in desired format
		timestamp, err := time.Parse(TimeLayout, combineTimeLog)
		if err != nil {
			return result, err
		}
		response, err := strconv.Atoi(r[6])
		if err != nil {
			return result, err
		}
		bBytes, err := strconv.Atoi(r[7])
		if err != nil {
			return result, err
		}

		// getting all keys
		result = append(result, NginxLogCombined{
			RemoteAddr:    r[0],
			RemoteUser:    r[2],
			TimeLocal:     timestamp,
			Request:       r[5],
			ResponseCode:  response,
			BodyBytesSent: bBytes,
			Referrer:      r[8],
			UserAgent:     r[9],
		})
	}
	return result, nil
}

func main() {
	// get rid of infra pod logs
	cleanLog, err := cleanOutNginxPodLog("nginx-k8s-pod.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Here is clear access log w/o any infra pod logs: \n %v \n\n", cleanLog)

	// generate structs
	parsedLog, err := parseNginxLog(cleanLog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Here is cvs parsed access log: %v \n\n", parsedLog)

	// show struxt example
	fmt.Printf(`Here is first example:
				Request was sent from ip: %s
				Sender is: %s
				It was made at: %v
				Request was: %s
				Nginx response code was: %v
				Bytes in body payload: %v
				Referrer is: %s
				User agent: %s
				`, parsedLog[0].RemoteAddr, parsedLog[0].RemoteUser, parsedLog[0].TimeLocal,
		parsedLog[0].Request, parsedLog[0].ResponseCode, parsedLog[0].BodyBytesSent,
		parsedLog[0].Referrer, parsedLog[0].UserAgent)
}
