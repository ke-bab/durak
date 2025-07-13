package durak

import (
	"errors"
	"fmt"
)

type BeatCardAction struct {
	Player *Player
	Card   *Card
}

func (a *BeatCardAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Play {
		return false, errors.New("game is in wrong state for playing cards")
	}

	if !g.hasPlayer(a.Player) {
		return false, fmt.Errorf("player %d not found", a.Player.ID)
	}

	// todo

	return true, nil
}

func (a *BeatCardAction) Apply(g *Game) {
	// todo
}

func (a *BeatCardAction) Name() ActionName {
	return ActionBeatCard
}
