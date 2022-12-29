package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

var BlackJack = 21

func readUserInput(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to read string", err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func createDisplayMessageForUser(cards []card) string {
	message := createDisplayMessageFor(cards)
	return fmt.Sprintf("Your cards are:%s \n", message)
}

func createDisplayMessageForBank(card card) string {
	return fmt.Sprintf("The bank's card visible to you is: %s (%d) \n", card.name, card.value)
}

func createDisplayMessageFor(cards []card) string {
	message := ""
	for i, card := range cards {
		if len(cards) == i+1 {
			message += fmt.Sprintf(" %s!", card.name)
		} else {
			message += fmt.Sprintf(" %s &", card.name)
		}
	}
	return message
}

func setup() ([]card, set, []card) {
	set := generateSet("./set.csv")

	userCards, setV1 := drawCards(set)
	bankCards, newSet := drawCards(setV1)

	return userCards, newSet, bankCards
}

func createEndGameMessage(userResult int, bankResult int) string {
	switch {
	case userResult > BlackJack:
		return "You lose!"
	case userResult == BlackJack && bankResult == BlackJack:
		return "It is a tie, nobody wins!"
	case userResult == BlackJack && bankResult != BlackJack:
		return "YOU WIN WITH BLACK JACK, YEEEESSSS!!"
	case bankResult == BlackJack:
		return "The bank wins with Black Jack!"
	case bankResult > BlackJack && userResult < BlackJack:
		return "You win yeeey!"
	case userResult > bankResult:
		return "You win yeeey!"
	default:
		return "You lose!"
	}
}
