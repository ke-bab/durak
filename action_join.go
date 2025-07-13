package durak

import (
	"errors"
	"fmt"
)

type JoinAction struct {
	PlayerId int
}

func (a *JoinAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Open {
		return false, errors.New("game is not open for players")
	}

	if g.isFull() {
		return false, errors.New("game already full")
	}

	playerId, ok := g.playerIdPool.reserveId()
	if !ok {
		return false, errors.New("no free id in player id pool")
	}
	a.PlayerId = playerId

	if g.hasPlayer(playerId) {
		err := g.playerIdPool.releaseId(playerId)
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("player %d already joined", playerId)
	}

	return true, nil
}

func (a *JoinAction) Apply(g *Game) {
	newPlayer := NewPlayer(a.PlayerId)
	g.Players[a.PlayerId] = newPlayer

	if len(g.Players) == maxPlayers {
		g.start()
	}
}

func (a *JoinAction) Name() ActionName {
	return ActionJoin
}
