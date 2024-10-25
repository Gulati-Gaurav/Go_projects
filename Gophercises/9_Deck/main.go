package deck

import (
	"math/rand"
	"sort"
)

// when you wanna mention array but don't declare size. the length of the array is inferred from the number of elements provided.
var suits = [...]Suit{Spades, Diamonds, Clubs, Hearts}

const (
	minRank = Ace
	maxRank = King
)

func New(options ...func([]Card) []Card) []Card {
	cards := []Card{}
	for i := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: suits[i],
				Rank: rank,
			})
		}
	}
	for i := range options {
		cards = options[i](cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank < cards[j].Rank
	})
	return cards
}

// It will take in comparator function and return a function of Options type
func CustomSort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	index := rand.Intn(len(cards))
	for i := range cards {
		cards[index], cards[i] = cards[i], cards[index]
	}
	return cards
}

func AddJokers(numOfJokers int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < numOfJokers; i++ {
			cards = append(cards, Card{
				Suit: Joker,
			})
		}
		return cards
	}
}

func Filter(fn func(card Card) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		newCards := []Card{}
		for i := 0; i < len(cards); i++ {
			if !fn(cards[i]) {
				newCards = append(newCards, cards[i])
			}
		}
		return newCards
	}
}

func MultipleDecks(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		NewDeck := []Card{}
		for i := 0; i < n; i++ {
			NewDeck = append(NewDeck, cards...)
		}
		return NewDeck
	}
}

// Stringer is used to print constants in a nice way

func (card Card) String() string {
	if card.Suit == Joker {
		return card.Suit.String()
	}
	return card.Rank.String() + " of " + card.Suit.String()
}
