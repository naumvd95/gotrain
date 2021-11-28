package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/* simple algo:
1. loop over start file, collect data [][]string == [[id,timestamp]], filtered by IP == startTrans
2. loop over end file, collect data map[string]string == "id": "timestamp", filtered by IP == endTrans
3. loop over startTrans and for each value get diff using endTrans key
4. during loooping, calculate avg == allTimestamps / len(startTrans)

smart algo -- use map reduce scenarios
1. map for start file + goroutines
2. doing reduce job during iterations in end file
*/

func main() {
	desiredIP := "192.168.10.15"
	//1.
	startData, err := os.Open("startTransactionslog.csv")
	if err != nil {
		panic(err)
	}
	defer startData.Close()

	csvReader := csv.NewReader(startData)
	startTrans, err := csvReader.ReadAll() // [][]string == [[headers], [item1].....[itemN]]
	if err != nil {
		panic(err)
	}

	startTransTimings := [][]string{}
	for _, v := range startTrans[1:] {
		tID := v[0]
		tTime := v[1]
		ip := v[2]
		if ip == desiredIP {
			startTransTimings = append(startTransTimings, []string{tID, tTime})
		}
	}

	//2.
	endData, err := os.Open("endTransactionslog.csv")
	if err != nil {
		panic(err)
	}
	defer endData.Close()
	csvReader = csv.NewReader(endData)
	endTrans, err := csvReader.ReadAll() // [][]string == [[headers], [item1].....[itemN]]
	if err != nil {
		panic(err)
	}

	endTransTimings := make(map[string]string)
	for _, v := range endTrans[1:] {
		tID := v[0]
		tTime := v[1]
		ip := v[2]
		if ip == desiredIP {
			endTransTimings[tID] = tTime
		}
	}

	//3.
	totalDiffSum := 0
	totalTransAmount := len(startTransTimings)

	for _, v := range startTransTimings {
		start, err := strconv.Atoi(v[1])
		if err != nil {
			panic(err)
		}

		id := v[0]
		end, err := strconv.Atoi(endTransTimings[id])
		if err != nil {
			panic(err)
		}

		totalDiffSum += end - start
	}
	fmt.Printf("total diff sum %v \n", totalDiffSum)

	//4.
	avgTiming := totalTransAmount / totalDiffSum
	fmt.Printf("Average timing for transactions start %v and end %v for ip %v is: %v", startTransTimings, endTransTimings, desiredIP, avgTiming)
}
