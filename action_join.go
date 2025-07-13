package durak

import (
	"errors"
	"fmt"
)

type JoinAction struct {
	Player *Player
}

func (a *JoinAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Open {
		return false, errors.New("game is not open for players")
	}

	if g.isFull() {
		return false, errors.New("game already full")
	}

	if g.hasPlayer(a.Player) {
		return false, fmt.Errorf("player %d already joined", a.Player.ID)
	}

	return true, nil
}

func (a *JoinAction) Apply(g *Game) {
	g.Players[a.Player.ID] = a.Player

	if len(g.Players) == maxPlayers {
		g.start()
	}
}

func (a *JoinAction) Name() ActionName {
	return ActionJoin
}
