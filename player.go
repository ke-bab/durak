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

func (p *Player) hasCard(c *Card) (int, bool) {
	for i, handCard := range p.Hand {
		if c.isSame(handCard) {
			return i, true
		}
	}

	return 0, false
}

func (p *Player) removeExistingCard(i int) {
	p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
}
