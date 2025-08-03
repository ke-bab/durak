package durak

type Deck []*Card

func (d Deck) isFull() bool {
	if len(d) == len(suits)*len(ranks) {
		return true
	}

	return false
}
