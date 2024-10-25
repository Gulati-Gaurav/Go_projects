package main

import "fmt"

func main() {

	cards := newDeck()
	err := cards.saveToLocal("cards")
	if err != nil {
		fmt.Println("Saving Failed")
	} else { // In go we have to mention the else in the same line
		fmt.Println("Saving Succeeded")
	}

	newCards := getFromLocal("cards")
	newCards.printCards()
}
