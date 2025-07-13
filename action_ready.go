package durak

import (
	"errors"
	"fmt"
)

type ReadyAction struct {
	Player *Player
}

func (a *ReadyAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Open {
		return false, errors.New("game is not open for players")
	}

	if !g.hasPlayer(a.Player) {
		return false, fmt.Errorf("player %d not found", a.Player.ID)
	}

	return true, nil
}

func (a *ReadyAction) Apply(g *Game) {
	a.Player.IsReady = true

	if g.isEnoughPlayersForStart() && g.isEveryoneReady() {
		g.start()
	}
}

func (a *ReadyAction) Name() ActionName {
	return ActionReady
}
