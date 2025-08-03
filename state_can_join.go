package durak

import (
	"errors"
	"fmt"
)

type StateCanJoin struct {
	Slot **Player
}

func NewStateCanJoin(g *Game) (*StateCanJoin, error) {
	if g.State != Open {
		if g.State != Open {
			return nil, errors.New("game is not open")
		}
	}

	if g.Player1 == nil {
		return &StateCanJoin{Slot: &g.Player1}, nil
	}

	if g.Player2 == nil {
		return &StateCanJoin{Slot: &g.Player2}, nil
	}

	return nil, fmt.Errorf("no free slot to join")
}

func (state *StateCanJoin) Join(p *Player) {
	*state.Slot = p
}
