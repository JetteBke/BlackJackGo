package main

import (
	"fmt"
)

func displayCardsAndResultToUser(result int, userCards []card, bankCard card) {
	fmt.Printf(createDisplayMessageForUser(userCards))
	fmt.Printf("Result: %d \n", result)
	fmt.Printf(createDisplayMessageForBank(bankCard))
}
