package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's play a simple version of black jack!")
	fmt.Println("To start the game, press 1")

	reader := bufio.NewReader(os.Stdin)
	userInput := readUserInput(reader)

	if userInput == "1" {
		fmt.Println("------------ START ------------")
		userResult, bankResult := play(reader)
		fmt.Println("------------ END ------------")
		fmt.Printf("The bank had: %d \n", bankResult)
		message := createEndGameMessage(userResult, bankResult)
		fmt.Println(message)

	} else {
		fmt.Println("See you next time!")
	}
}

func play(reader *bufio.Reader) (int, int) {
	userCards, set, bankCards := setup()
	userResult := calculateResult(userCards)
	bankResult := calculateResult(bankCards)
	displayCardsAndResultToUser(userResult, userCards, bankCards[0])

	round := 1

	for userResult < BlackJack && bankResult < BlackJack {
		fmt.Printf("------------ ROUND %d------------\n", round)
		fmt.Println("Do you want to draw another card (1) or end game (2)?")
		userInput := readUserInput(reader)
		if userInput == "1" {
			newUserCards, newUserResult, setV1 := playRound(set, userCards)
			newBankCards, newBankResult, newSet := playRound(setV1, bankCards)
			displayCardsAndResultToUser(newUserResult, newUserCards, bankCards[0])
			userCards = newUserCards
			userResult = newUserResult
			bankCards = newBankCards
			bankResult = newBankResult
			set = newSet
		} else {
			return userResult, bankResult
		}
		round += 1
	}
	return userResult, bankResult
}
