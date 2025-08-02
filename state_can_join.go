package durak

import "errors"

type StateCanJoin struct {
	Player *Player
}

func NewStateCanJoin(g *Game) (*StateCanJoin, error) {
	if g.State != Open {
		if g.State != Open {
			return nil, errors.New("game is not open")
		}
	}

	hasSlot, err := NewStateHasSlot(g)
	if err != nil {
		return nil, err
	}

	return &StateCanJoin{Player: hasSlot.Slot}, nil
}

func (state *StateCanJoin) Join(p *Player) {
	state.Player = p
}
