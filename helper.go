package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var BLACK_JACK = 21

func readUserInput(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to read string", err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func displayCardsAndResultToUser(result int, userCards []card) {
	message := "Your cards are:"
	for i, card := range userCards {
		if len(userCards) == i+1 {
			message += fmt.Sprintf(" %s! \n", card.name)
		} else {
			message += fmt.Sprintf(" %s &", card.name)
		}
	}
	fmt.Printf(message)
	fmt.Printf("Result: %s \n", strconv.Itoa(result))
}

func setup() ([]card, int, set) {
	set := generateSet("./set.csv")
	userCards, newSet := drawUserCards(set)
	result := calculateResult(userCards)

	return userCards, result, newSet
}
