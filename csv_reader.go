package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	cards, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return cards
}

func convertToSet(cards [][]string) set {
	newSet := set{}
	for _, c := range cards {
		intVar, _ := strconv.Atoi(c[1])
		newSet.set = append(newSet.set, card{
			name:  c[0],
			value: intVar,
		})
	}
	return newSet
}
