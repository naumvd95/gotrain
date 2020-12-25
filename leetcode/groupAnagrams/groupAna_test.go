package main

import (
	"reflect"
	"sort"
	"testing"
)

type data struct {
	Input         []string
	DesiredResult [][]string
}

var dataSet = []data{
	data{
		Input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		DesiredResult: [][]string{
			[]string{"bat"},
			[]string{"nat", "tan"},
			[]string{"ate", "eat", "tea"},
		},
	},
	data{
		Input: []string{""},
		DesiredResult: [][]string{
			[]string{""},
		},
	},
	data{
		Input: []string{"a"},
		DesiredResult: [][]string{
			[]string{"a"},
		},
	},
}

func TestGroupAnagramsByAlphabetHash(t *testing.T) {
	for _, v := range dataSet {
		res := groupAnagramsByAlphabetHash(v.Input)
		if len(res) != len(v.DesiredResult) {
			t.Errorf("[incorrect size] for %v\n expected %v\n got %v\n\n", v.Input, v.DesiredResult, res)
		}

		for _, expectedGroup := range v.DesiredResult {
			matches := false

			for _, resultGroup := range res {
				sort.Strings(resultGroup)

				if reflect.DeepEqual(resultGroup, expectedGroup) {
					matches = true
				}
			}
			if !matches {
				t.Errorf("[no case match] for %v\n expected %v\n got %v\n\n", v.Input, v.DesiredResult, res)
			}
		}
	}
}

func TestGroupAnagramsByStringSort(t *testing.T) {
	for _, v := range dataSet {
		res := groupAnagramsByStringSort(v.Input)
		if len(res) != len(v.DesiredResult) {
			t.Errorf("[incorrect size] for %v\n expected %v\n got %v\n\n", v.Input, v.DesiredResult, res)
		}

		for _, expectedGroup := range v.DesiredResult {
			matches := false

			for _, resultGroup := range res {
				sort.Strings(resultGroup)

				if reflect.DeepEqual(resultGroup, expectedGroup) {
					matches = true
				}
			}
			if !matches {
				t.Errorf("[no case match] for %v\n expected %v\n got %v\n\n", v.Input, v.DesiredResult, res)
			}
		}
	}
}
