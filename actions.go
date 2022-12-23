package main

import (
	"math/rand"
	"time"
)

type card struct {
	name  string
	value int
}

type set struct {
	set []card
}

func generateSet(filePath string) set {
	fileContent := readCsvFile(filePath)
	return convertToSet(fileContent)
}

func drawUserCards(oldSet set) ([]card, set) {
	cardOne, set := drawCard(oldSet)
	cardTwo, newSet := drawCard(set)
	return []card{cardOne, cardTwo}, newSet
}

func drawCard(oldSet set) (card, set) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(oldSet.set))
	newCard := oldSet.set[randomIndex]
	newSet := append(oldSet.set[:randomIndex], oldSet.set[randomIndex+1:]...)
	return newCard, set{set: newSet}
}

func calculateResult(cards []card) int {
	sum := 0
	for _, card := range cards {
		sum += card.value
	}
	return sum
}

func playRound(set set, userCards []card) ([]card, int, set) {
	newCard, newSet := drawCard(set)
	newUserCards := append(userCards, newCard)
	result := calculateResult(newUserCards)
	return newUserCards, result, newSet
}
