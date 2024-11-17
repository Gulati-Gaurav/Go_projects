package main

import "fmt"

// Mention the return type in function declaration like (not mentioning means void)
func newCard() string {
	return "five of diamonds"
}

func variables() {
	// 3 Types of declaring variables
	// var Card string = "Ace of spades"
	// var Card        = "Ace of spades" (no need of mentioning the type if declaring also)
	card := "Ace of Spades" // var is replaced by :
	// := is called the walrus operator
	fmt.Println(card)

	// for const variables, global variables, without initialisation CAN'T use walrus
	const name = "Gaurav"

	// global variables
	// var jtoken string = "dsfv"
	// or
	// var jtoken = "dsfv" // now able to understand its significance
	// but can't do
	// jtoken := "dsfv"

	// Now how to initialise but not give value :-
	var deckSize int
	deckSize = 52
	fmt.Println(deckSize)

	// to get value from a function
	newcard := newCard()
	fmt.Println(newcard)
}
