// Code generated by "stringer -type=Suit"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Hearts-0]
	_ = x[Clubs-1]
	_ = x[Spades-2]
	_ = x[Diamonds-3]
	_ = x[Joker-4]
}

const _Suit_name = "HeartsClubsSpadesDiamondsJoker"

var _Suit_index = [...]uint8{0, 6, 11, 17, 25, 30}

func (i Suit) String() string {
	if i >= Suit(len(_Suit_index)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}
