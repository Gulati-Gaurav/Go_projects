package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"clubs", "diamonds", "hearts", "spades"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for i := range values {
		for j := range suits {
			cards = append(cards, values[i]+" "+suits[j])
		}
	}

	return cards
}

func (d deck) printCards() {
	for i := range d {
		fmt.Println(i, d[i], ",")
	}
}

func (d deck) pseudoShuffle() {
	for i := range d {
		index := rand.Intn(len(d))
		// This is how you swap 2 indexes easily
		d[i], d[index] = d[index], d[i]
	}
}

func deal(d deck, hands int) (deck, deck) {
	return d[:hands], d[hands:]
}

func (d deck) saveToLocal(fileName string) error {
	data := d.toByteArray()
	return os.WriteFile(fileName, data, 0666)
}

func (d deck) toByteArray() []byte {
	return []byte(strings.Join([]string(d), ","))
}

func getFromLocal(fileName string) deck {
	byteDeck, err := os.ReadFile(fileName)
	if err == nil { // in go no need of () in the if statement. Infact on saving vs will remove.
		return toDeck(byteDeck)
	}
	fmt.Println(err)
	os.Exit(1)
	return nil // see in such case we can return nil
}

func toDeck(data []byte) deck {
	return deck(strings.Split(string(data), ","))
}
