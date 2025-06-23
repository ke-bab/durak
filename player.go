package durak

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
	// todo
}

func (p *Player) removeCard(c *Card) {
	// todo
}

func (p *Player) playCard(c *Card, cards CardsOnTable) {
	cards.add(c)
	p.removeCard(c)
}
