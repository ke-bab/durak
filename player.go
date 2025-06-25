package durak

import "fmt"

type Player struct {
	ID      int
	IsReady bool
	Hand    []*Card
}

func NewPlayer(id int) *Player {
	return &Player{
		ID:   id,
		Hand: make([]*Card, 0),
	}
}

func (p *Player) hasCard(c *Card) bool {
	if _, ok := p.findCardIndex(c); !ok {
		return false
	}

	return true
}

func (p *Player) removeCard(c *Card) error {
	index, ok := p.findCardIndex(c)
	if !ok {
		return fmt.Errorf("card %s %s cannot be removed", c.Rank, c.Suit)
	}

	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)

	return nil
}

func (p *Player) findCardIndex(c *Card) (int, bool) {
	for i, handCard := range p.Hand {
		if c.isSame(handCard) {
			return i, true
		}
	}

	return 0, false
}

func (p *Player) playCard(c *Card, cards CardCollection) error {
	err := p.removeCard(c)
	if err != nil {
		return err
	}
	cards.add(c)

	return nil
}
