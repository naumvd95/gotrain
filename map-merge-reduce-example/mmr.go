package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	pretty "github.com/inancgumus/prettyslice"
	"github.com/naumvd95/gotrain/common"
	sort "github.com/naumvd95/gotrain/sorting-example/merge-sort"
)

const (
	//ComparingDate desired date to show collect DeathByDate metric
	ComparingDate = "6/15/20"
	// CriticalDeathsPerDay represents low border of deaths amount to reduce data
	CriticalDeathsPerDay = 1000
	// DataSetURL csv data set w/ covid cases
	DataSetURL = "https://data.humdata.org/hxlproxy/api/data-preview.csv?url=https%3A%2F%2Fraw.githubusercontent.com%2FCSSEGISandData%2FCOVID-19%2Fmaster%2Fcsse_covid_19_data%2Fcsse_covid_19_time_series%2Ftime_series_covid19_confirmed_global.csv&filename=time_series_covid19_confirmed_global.csv"
)

//csvMetaPointer struct needed to manage parsed csv data objects and store date index position from the csv key set
type csvMetaPointer struct {
	data                [][]string
	desiredDatePosition int // store index in array which represents ComparingDate in csv data
}

//getCSVDataSet download csv dataset, parsed though csv lib, identifies index of desired date csv key.
func getCSVDataSet(desiredDate string) (csvMetaPointer, error) {
	var ds csvMetaPointer
	var dateIndex int

	resp, err := http.Get(DataSetURL)
	if err != nil {
		return ds, err
	}
	defer resp.Body.Close()

	// parsing csv
	body := csv.NewReader(resp.Body)
	records, err := body.ReadAll()
	if err != nil {
		return ds, err
	}

	//detecting desired date index, iterating over csv keys
	dateNotFound := true
	for index, date := range records[0] {
		if date == desiredDate {
			dateIndex = index
			dateNotFound = false
		}
	}
	if dateNotFound {
		return ds, errors.New("Failed to identify desired date in data set")
	}

	// cut off useless keys
	ds.data = records[1:]
	ds.desiredDatePosition = dateIndex

	fmt.Printf("Dataset received, number of lines: %v \n", len(ds.data))
	return ds, nil
}

//Map for each line in csv dataset we use that func to generate CovidUnit struct object w/ all needed key/values.
func Map(covidIncident []string, byDateIndex int) []common.CovidUnit {
	// create slice to perform multithread concatenation later
	mappedResList := []common.CovidUnit{}
	deathAmount, _ := strconv.Atoi(covidIncident[byDateIndex])

	mappedResList = append(mappedResList, common.CovidUnit{
		Province:    covidIncident[0],
		Country:     covidIncident[1],
		Latitude:    covidIncident[2],
		Longitude:   covidIncident[3],
		DeathByDate: deathAmount, // search only info by specified date
	})

	return mappedResList
}

//Reducer listens to results from Map channel, filter datasets and sends updated data into final channel
func Reducer(criticalDeath int, mapList chan []common.CovidUnit, sendFinalValue chan []common.CovidUnit) {
	finalRes := []common.CovidUnit{}

	// mapList channel has not got index, only data w/ slices
	for list := range mapList {
		// going deeper and parse slice in channel data
		for _, accident := range list {
			if accident.DeathByDate >= criticalDeath {
				finalRes = append(finalRes, accident)
			}
		}
	}

	// send final value
	sendFinalValue <- finalRes
}

func main() {
	covidCases, err := getCSVDataSet(ComparingDate)
	if err != nil {
		log.Fatal(err)
	}

	// Init channels for managing objects
	mappedLists := make(chan []common.CovidUnit)
	finalValue := make(chan []common.CovidUnit)
	var wg sync.WaitGroup
	// specify number of goroutines for wait group
	wg.Add(len(covidCases.data))

	startTime := time.Now()
	fmt.Printf("Dataset parsed, starting map-reduce ops at: %v \n", startTime)
	// for each slice in csv slices in dataset [][]string
	for _, line := range covidCases.data {
		// goroutine forwards accident slice metadata to Map
		go func(accident []string) {
			// forward all data in channel
			mappedLists <- Map(accident, covidCases.desiredDatePosition)
			// anyway close goroutine
			defer wg.Done()
		}(line)
	}

	// run Reduce filters by CriticalDeathsPerDay all data in channel mappedLists and sends results in finalValue
	go Reducer(CriticalDeathsPerDay, mappedLists, finalValue)
	wg.Wait()
	// make sure we closed channel
	close(mappedLists)
	/*
			notice:
		    value, more := <-mappedLists
			in above form, `more` is boolean value that will be false if we close mappedLists channel
	*/

	reducedData := <-finalValue
	close(finalValue)

	fmt.Printf("Here is list of critical amount of the covid accidents (more than %v) in all countries, happened at %s: \n\n %v \n\n",
		CriticalDeathsPerDay, ComparingDate, reducedData)

	sortedData := sort.MergeSortMultiThread(reducedData)
	fmt.Printf(`All operations w/ that dataset:
	Map - into object and by desired date
	Reduce - by critical death amount
	Merge sort (multithread)`)
	pretty.MaxPerLine = 1
	pretty.Show("Sorted Data:", sortedData)

	finishTime := time.Now()
	fmt.Printf("All map/reduce & sorting operations finised at: %v , overall it took %v \n", finishTime, finishTime.Sub(startTime))
}
