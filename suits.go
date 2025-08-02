package durak

type Suit string

const (
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Clubs    Suit = "clubs"
	Spades   Suit = "spades"
)

var suits = []Suit{Hearts, Diamonds, Clubs, Spades}
