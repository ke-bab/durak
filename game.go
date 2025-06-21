package durak

import (
	"errors"
	"fmt"
	"sync"
)

const maxPlayers = 6

type Game struct {
	Players map[int]Player
	lock    sync.Mutex
}

func NewGame() *Game {
	return &Game{
		Players: make(map[int]Player, maxPlayers),
	}
}

func (g *Game) JoinPlayer(p Player) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.isFull() {
		return errors.New("game already full")
	}

	if g.hasPlayer(p.ID) {
		return errors.New(fmt.Sprintf("player %d already joined", p.ID))
	}

	g.Players[p.ID] = p

	return nil
}

func (g *Game) hasPlayer(id int) bool {
	if _, ok := g.Players[id]; ok {
		return true
	}

	return false
}

func (g *Game) isFull() bool {
	if len(g.Players) >= maxPlayers {
		return true
	}

	return false
}
