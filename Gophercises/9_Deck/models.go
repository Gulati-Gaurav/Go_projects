//go:generate stringer -type=Pill

package deck

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Suit uint8

const (
	Hearts Suit = iota
	Clubs
	Spades
	Diamonds
	Joker
)

type Card struct {
	Rank Rank
	Suit Suit
}
