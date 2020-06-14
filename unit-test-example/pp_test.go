package main

import (
	"fmt"
	"testing"
)

func TestPrettyPrintStr(t *testing.T) {

	jsonExpected := fmt.Sprintf("here is json struct from csv: %v \n\n", "{}")
	jsonRes, err := PrettyPrintStr("structedJson", "{}")
	if err != nil {
		t.Errorf(
			"PrettyPrintStr(\"structedJson\", \"{}\") failed, got err: %v", err)
	}
	if jsonRes != jsonExpected {
		t.Errorf(
			"PrettyPrintStr(\"structedJson\", \"{}\") failed, expected: %s, got: %v\n", jsonExpected, jsonRes)
	}

	mapExpected := fmt.Sprintf("here is map from csv: %v \n\n", "{}")
	mapRes, err := PrettyPrintStr("map", "{}")
	if err != nil {
		t.Errorf(
			"PrettyPrintStr(\"structedJson\", \"{}\") failed, got err: %v", err)
	}
	if mapRes != mapExpected {
		t.Errorf(
			"PrettyPrintStr(\"structedJson\", \"{}\") failed, expected: %s, got: %v\n", mapExpected, mapRes)
	}

	_, err = PrettyPrintStr("incorrectDataType", "{}")
	if err != PrettyErr {
		t.Errorf(
			"PrettyPrintStr(\"incorrectDataType\", \"{}\") failed, expected err: %v, got: %v\n", PrettyErr, err)
	}
}
