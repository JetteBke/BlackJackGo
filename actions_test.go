package main

import (
	"reflect"
	"testing"
)

func TestDrawCard(t *testing.T) {
	set := generateSet("./test_set.csv")
	newCard, newSet := drawCard(set)
	if !(len(newSet.set) == 3) && !reflect.DeepEqual(newCard, newSet.set) {
		t.Error("Did not draw a card correctly!")
	}
}

func TestDrawUserCards(t *testing.T) {
	set := generateSet("./test_set.csv")
	userCards, newSet := drawCards(set)
	if !(len(newSet.set) == 2) && !(len(userCards) == 2) {
		t.Error("Did not draw user cards correctly!")
	}
}

func TestCalculateResult(t *testing.T) {
	result := calculateResult([]card{{name: "hi", value: 3}, {name: "hi2", value: 2}})
	if !(result == 5) {
		t.Error("Did not calculate result of cards correctly!")
	}
}
