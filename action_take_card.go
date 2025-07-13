package durak

import (
	"errors"
	"fmt"
)

type TakeCardAction struct {
	Player *Player
	Card   *Card
	//
	cardIndex int
}

func (a *TakeCardAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Play {
		return false, errors.New("game is in wrong state for playing cards")
	}

	if !g.hasPlayer(a.Player) {
		return false, fmt.Errorf("player %d not found", a.Player.ID)
	}

	// todo

	return true, nil
}

func (a *TakeCardAction) Apply(g *Game) {
	// todo
}

func (a *TakeCardAction) Name() ActionName {
	return ActionTakeCard
}
