package deck

import (
	"fmt"
	"testing"
)

// This is how we can mention the expected output in the test case
func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Hearts})

	// Output:
	// Ace of Hearts
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if cards[0].Rank.String() != "Ace" || cards[0].Suit.String() != "Hearts" {
		t.Error("First card not correct", cards[0])
	}
}

func TestCustomSort(t *testing.T) {
	less := func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool {
			if cards[i].Suit != cards[j].Suit {
				return cards[i].Suit > cards[j].Suit
			}
			return cards[i].Rank < cards[j].Rank
		}
	}
	cards := New(CustomSort(less))

	if cards[0].Rank != Ace || cards[0].Suit != Diamonds {
		t.Error("Not sorted correctly")
	}
}

func TestShuffle(t *testing.T) {
	cards1 := New(Shuffle)
	cards2 := New(Shuffle)

	if cards1[0].Rank == cards2[0].Rank && cards1[0].Suit == cards2[0].Suit {
		t.Error("Not Shuffled")
	}

}

func TestAddJokers(t *testing.T) {
	cards := New(AddJokers(4))
	j := len(cards) - 1
	for i := 0; i < 4; i++ {
		if cards[j].Suit != Joker {
			t.Error("Adding Joker not succcess")
		}
		j--
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(card Card) bool {
		return card.Rank%2 == 0
	}))

	for i := range cards {
		if cards[i].Rank%2 == 0 {
			t.Error("Filter gone wrong. even found")
			return
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	cards := New(MultipleDecks(2))
	if len(cards) != 2*52 {
		t.Error("Not correct Multiple Decks size")
	}
	for i := 0; i < 52; i++ {
		if cards[i].Rank != cards[i+52].Rank || cards[i].Suit != cards[i+52].Suit {
			t.Error("Incorrect cards duplicacy")
		}
	}
}
