package durak

type Card struct {
	Suit Suit
	Rank Rank
}

func NewCard(s Suit, r Rank) *Card {
	return &Card{
		Suit: s,
		Rank: r,
	}
}

type CardsOnTable []*Card

func (cards *CardsOnTable) add(c *Card) {
	*cards = append(*cards, c)
}
