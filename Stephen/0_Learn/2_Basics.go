package main

import "fmt"

func arrays_slices() {
	// array = array
	// slice = vector
	// you mention size it is array. You don't it is slice.
	// However you can omit writing elements in array and can write elements initially in a slice

	cards := []string{newCard(), "Ace of diamonds"} // slice
	// put {} in end of []string, map[string]string
	cards = append(cards, "Six of Spades") // add element.
	// cards = append(cards, "Six of Spades", "Ace of Clubs") // add multiple elements.
	// imp The append doesn't modify existing slice instead returns new one.

	// To make a slice of certain size
	// _ := make([]int, 5)

	var arrayName [5]int                  // array
	numbers := [5]int{10, 20, 30, 40, 50} // array
}

func for_discard_subset() {
	cards2 := [5]int{1, 3, 4, 5, 8}
	
	// iterate. i = index, card = element at ith index.
	// range is keyword for every element in slice. first is index, second is element.
	for i, card := range cards2 {
		fmt.Println(i, card, cards2[i])
	}

	// for every iteration i, card are initialised i.e. card is not actual cards[i] but rather a copy of cards[i] created using card:=cards[i]. Proof:-

	for i, card := range cards2 {
		card = 12
		fmt.Println(i, card, cards2[i])
	}
	// Response
	// 0 12 1
	// 1 12 3
	// 2 12 4
	// 3 12 5
	// 4 12 8

	// In case it is a struct then use below (array of pointers of struct)
	// for i, card := range cards2 {
	// 	*card = 12
	// 	fmt.Println(i, *card, cards2[i])
	// }

	// the underscore (_) is a special identifier used to discard values. It's often used when you need to call a function or method that returns a value, but you don't need that value.
	for _, card := range cards2 {
		fmt.Println(card)
	}

	// To get subset. [startIndexIncluding : upToNotIncluding]. [:n] means start from 0. [1:] means go to end
	// will not modify the original slice. returns a new slice
	subCard := cards2[0:2]

	fmt.Println(cards2, subCard, arrayName, numbers)
}
