package main

import (
	"reflect"
	"testing"
)

func TestReadCsvFile(t *testing.T) {
	strArr := [][]string{{"Two of Hearts", "2"}, {"King of Hearts", "10"}, {"Three of Hearts", "3"}, {"Four of Hearts", "4"}}
	if !reflect.DeepEqual(readCsvFile("./test_set.csv"), strArr) {
		t.Error("No deep equality!")
	}
}

func TestConvertToCard(t *testing.T) {
	strArr := [][]string{{"Two of Hearts", "2"}, {"King of Hearts", "10"}}
	set := set{set: []card{
		{
			name:  "Two of Hearts",
			value: 2,
		},
		{
			name:  "King of Hearts",
			value: 10,
		},
	}}
	if !reflect.DeepEqual(convertToSet(strArr), set) {
		t.Error("No deep equality!")
	}
}
