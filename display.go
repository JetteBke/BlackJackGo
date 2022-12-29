package main

import (
	"fmt"
	"strconv"
)

func displayCardsAndResultToUser(result int, userCards []card, bankCard card) {
	fmt.Printf(createDisplayMessageForUser(userCards))
	fmt.Printf(createDisplayMessageForBank(bankCard))
	fmt.Printf("Result: %s \n", strconv.Itoa(result))
}
