package durak

type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
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

// we assume that we have enough cards to do that.
// todo: try safer approach using state struct.
func (cards *CardCollection) takeXCards(n int) []*Card {
	firstXCards := (*cards)[:n]
	// remove them from deck
	*cards = (*cards)[n:]

	return firstXCards
}
