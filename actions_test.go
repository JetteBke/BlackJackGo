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

func Test_playBankRound(t *testing.T) {
	type args struct {
		set   set
		cards []card
	}

	testCardsOne := []card{{"Seven of Hearts", 7}, {"King of Hearts", 10}}
	testSet := generateSet("./test_set.csv")

	tests := []struct {
		name         string
		args         args
		wantedCards  []card
		wantedResult int
		wantedSet    set
	}{
		{
			name:         "test user has 17 or more",
			args:         args{set: testSet, cards: testCardsOne},
			wantedCards:  testCardsOne,
			wantedResult: 17,
			wantedSet:    testSet,
		},
		//	TODO, add other test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCards, gotResult, gotSet := playBankRound(tt.args.set, tt.args.cards)
			if !reflect.DeepEqual(gotCards, tt.wantedCards) {
				t.Errorf("playBankRound() got cards = %v, want %v", gotCards, tt.wantedCards)
			}
			if gotResult != tt.wantedResult {
				t.Errorf("playBankRound() got result = %v, want %v", gotResult, tt.wantedResult)
			}
			if !reflect.DeepEqual(gotSet, tt.wantedSet) {
				t.Errorf("playBankRound() got set = %v, want %v", gotSet, tt.wantedSet)
			}
		})
	}
}
