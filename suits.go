package durak

type Suit string

const (
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Clubs    Suit = "clubs"
	Spades   Suit = "spades"
)

func suits() []Suit {
	return []Suit{Hearts, Diamonds, Clubs, Spades}
}

type Rank string

const (
	Six   Rank = "six"
	Seven Rank = "seven"
	Eight Rank = "eight"
	Nine  Rank = "nine"
	Ten   Rank = "ten"
	Jack  Rank = "jack"
	Queen Rank = "queen"
	King  Rank = "king"
	Ace   Rank = "ace"
)

func ranks() []Rank {
	return []Rank{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
}
