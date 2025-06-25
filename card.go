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

func (c *Card) isSame(card *Card) bool {
	return c.Suit == card.Suit && c.Rank == card.Rank
}

type CardCollection []*Card

func (cards *CardCollection) add(c *Card) {
	*cards = append(*cards, c)
}
