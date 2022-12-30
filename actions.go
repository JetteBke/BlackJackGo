package main

import (
	"github.com/samber/lo"
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

func drawCards(oldSet set) ([]card, set) {
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
	return lo.Reduce(cards, func(agg int, item card, index int) int {
		return agg + item.value
	}, 0)
}

func playUserRound(set set, cards []card) ([]card, int, set) {
	newCard, newSet := drawCard(set)
	newBankCards := append(cards, newCard)
	result := calculateResult(newBankCards)
	return newBankCards, result, newSet
}

func playBankRound(set set, cards []card) ([]card, int, set) {
	bankResult := calculateResult(cards)
	switch {
	case bankResult >= 17:
		return cards, bankResult, set
	default:
		newCard, newSet := drawCard(set)
		newBankCards := append(cards, newCard)
		result := calculateResult(newBankCards)
		return newBankCards, result, newSet
	}
}
