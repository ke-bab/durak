package durak

import "fmt"

type StateHasSlot struct {
	Slot *Player
}

func NewStateHasSlot(g *Game) (*StateHasSlot, error) {
	if g.Player1 == nil {
		return &StateHasSlot{Slot: g.Player1}, nil
	}

	if g.Player2 == nil {
		return &StateHasSlot{Slot: g.Player2}, nil
	}

	return nil, fmt.Errorf("game has no slots")
}
