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
	// 									 (useful in case of constant and global variables)
	card := "Ace of Spades" // := is called the walrus operator

	// for const variables, global variables, without initialisation CAN'T use walrus
	const name = "Gaurav"

	// global variables
	// var jtoken string = "dsfv"
	// or
	// var jtoken = "dsfv" 
	// but can't do
	// jtoken := "dsfv"

	// Now how to initialise but not give value :-
	var deckSize int
	deckSize = 52

	// to get value from a function
	newcard := newCard()
	fmt.Println(newcard, card, name, deckSize)
}
