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
		result := play(reader)
		if result == BLACK_JACK {
			fmt.Println("YOU WIN YEEEESSSS!!")
		} else {
			fmt.Println("You lost!")
		}
	} else {
		fmt.Println("See you next time!")
	}
}

func play(reader *bufio.Reader) int {
	userCards, result, set := setup()
	displayCardsAndResultToUser(result, userCards)
	for result < BLACK_JACK {
		fmt.Println("Do you want to draw another card (1) or end game (2)?")
		userInput := readUserInput(reader)
		if userInput == "1" {
			newUserCards, newResult, newSet := playRound(set, userCards)
			displayCardsAndResultToUser(newResult, newUserCards)
			userCards = newUserCards
			result = newResult
			set = newSet
		} else {
			return result
		}
	}
	return result
}
