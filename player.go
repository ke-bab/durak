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

// why we do not check that index exists?
// because we assume that we already checked it, and that's why function named as "existing".
// but in go impossible to define type which "has particular struct in slice", so we just hope this code
// will be used correctly.
//
// it is possible to encapsulate this behavior into another struct, but it will lead to more complex code,
// and will require tests to prove that it does what it should.
func (p *Player) removeExistingCard(i int) {
	p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
}
