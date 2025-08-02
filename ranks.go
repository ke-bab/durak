package durak

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

var ranks = []Rank{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
