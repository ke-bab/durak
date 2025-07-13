package durak

import (
	"errors"
	"fmt"
)

type ReadyAction struct {
	PlayerId int
	player   *Player
}

func (a *ReadyAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Open {
		return false, errors.New("game is not open for players")
	}

	player, ok := g.Players[a.PlayerId]
	if !ok {
		return false, fmt.Errorf("player %d not found", a.PlayerId)
	}
	a.player = player

	return true, nil
}

func (a *ReadyAction) Apply(g *Game) {
	a.player.IsReady = true

	if g.isEnoughPlayersForStart() && g.isEveryoneReady() {
		g.start()
	}
}

func (a *ReadyAction) Name() ActionName {
	return ActionReady
}
