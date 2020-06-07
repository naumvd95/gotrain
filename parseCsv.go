package main

/*
Input:
csv string:
`
Username; Identifier;First name;Last name
booker12;9012;Rachel;Booker
grey07;2070;Laura;Grey
johnson81;4081;Craig;Johnson
jenkins46;9346;Mary;Jenkins
smith79;5079;Jamie;Smith
`
Output:
Prettified data sets like:
{"Username": "booker12", "Identifier": "9012", "First name": Rachel, "Last name": "Booker"}
*/
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type dataStruct struct {
	Username   string `json:"username"`
	Identifier int    `json:"identifier"`
	FirstName  string `json:"first name"`
	LastName   string `json:"last name"`
}

//csvAsJSON receives string interpolation of CSV data and returns parsed struct in string interpolated json
func csvAsJSON(in string) (string, error) {
	// 1. Use external csv lib
	r := csv.NewReader(strings.NewReader(in))
	// 2. set custom separator
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}
	fmt.Printf("here is parsed csv [][]string: %v \n", records)

	fmt.Println("Marshalling into predefined struct")
	// 3. create struct w/ predefined fields
	var persons []dataStruct
	// 4. skipping slice w/ keys
	for _, r := range records[1:] {
		parsedIdentifier, err := strconv.Atoi(r[1])
		if err != nil {
			return "", err
		}

		persons = append(persons, dataStruct{
			Username:   r[0],
			Identifier: parsedIdentifier,
			FirstName:  r[2],
			LastName:   r[3],
		})
	}
	// 4. Marshall in json to prettify output
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		return "", err
	}

	// 5. Use strings to convert json obj for pretty printing
	return string(personsJSON), nil
}

//csvAsMap receives string interpolation of CSV data and returns map w/ parsed keys/elements
func csvAsMap(in string) ([]map[string]string, error) {
	//TODO here is repeatable steps, used in another func as well,
	//duplicating it only for study perspectives, ideally need 1 more common func
	// 1. Use external csv lib
	r := csv.NewReader(strings.NewReader(in))
	// 2. set custom separator
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	fmt.Printf("here is parsed csv [][]string: %v \n", records)

	usernameKey := records[0][0]
	identKey := records[0][1]
	firstNameKey := records[0][2]
	lastNameKey := records[0][3]
	// 6. create map
	var personsMap []map[string]string
	// 7. skipping slice w/ keys
	for _, r := range records[1:] {
		personsMap = append(personsMap, map[string]string{
			usernameKey:  r[0],
			identKey:     r[1], // cannot use mixed types string+int easily, so present it as string
			firstNameKey: r[2],
			lastNameKey:  r[3],
		})
	}

	return personsMap, nil
}

//PrettyPrintStr just prettifies printing and used for checking go test ;)
func PrettyPrintStr(dataType, dataSet string) (string, error) {
	switch dataType {
	case "structedJson":
		return fmt.Sprintf("here is json struct from csv: %v \n\n", dataSet), nil
	case "map":
		return fmt.Sprintf("here is map from csv: %v \n\n", dataSet), nil
	default:
		return "", PrettyErr
	}

}

// PrettyErrorSet represents custom error interface
type PrettyErrorSet string

func (e PrettyErrorSet) Error() string {
	return string(e)
}

const (
	csvInput string = `Username;Identifier;First name;Last name
 booker12;9012;Rachel;Booker
 grey07;2070;Laura;Grey
 johnson81;4081;Craig;Johnson
 jenkins46;9346;Mary;Jenkins
 smith79;5079;Jamie;Smith`
	//PrettyErr custom error
	PrettyErr PrettyErrorSet = "Failed to prettify output"
)

func main() {
	var ppStr string
	var ppErr error

	fmt.Println("1. Marshalling into custom json struct")
	personsJSON, err := csvAsJSON(csvInput)
	if err != nil {
		log.Fatal(err)
	}
	ppStr, ppErr = PrettyPrintStr("structedJson", personsJSON)
	if err != nil {
		log.Fatal(ppErr)
	}
	fmt.Println(ppStr)

	fmt.Println("2. Alternative Marshalling into custom map")
	personsMap, err := csvAsMap(csvInput)
	if err != nil {
		log.Fatal(err)
	}
	// convert map to string
	var slicePersons string
	for _, el := range personsMap {
		for key, value := range el {
			slicePersons += fmt.Sprintf("%s=%s,\n", key, value)
		}
	}

	ppStr, ppErr = PrettyPrintStr("map", slicePersons)
	if err != nil {
		log.Fatal(ppErr)
	}
	fmt.Println(ppStr)

	fmt.Println("3. Read/Write file operations w/ data sets")
	// 8. Writing/Reading json
	f, err := os.Create("csv-to-json-output-example.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(string(personsJSON))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("Here is csv-to-json-output-example.json, try reading")
	// 9. Open json file and read bytes from it
	jsonFile, err := os.Open("csv-to-json-output-example.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// 10. Init struct object to unmarshall in
	var personsFromFile []dataStruct
	err = json.Unmarshal(jsonBytes, &personsFromFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("here is struct from json file: %v \n", personsFromFile)

}
