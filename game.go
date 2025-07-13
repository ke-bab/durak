package durak

import (
	"sync"
)

const maxPlayers = 4 // in one game
const maxGames = 100
const minPlayersForStart = 2

type Game struct {
	Players      map[int]*Player
	State        GameState
	CardsOnTable []*Card

	moveOrder    *MoveOrder
	lock         sync.Mutex
	playerIdPool *IdPool
}

func NewGame(pool *IdPool) *Game {
	return &Game{
		Players:      make(map[int]*Player, maxPlayers),
		State:        Open,
		playerIdPool: pool,
		CardsOnTable: make([]*Card, 0),
	}
}

func (g *Game) DoAction(a Action) error {
	ok, err := a.CanBeApplied(g)
	if ok {
		g.lock.Lock()
		defer g.lock.Unlock()

		a.Apply(g)
		return nil
	}

	return err
}

func (g *Game) isEnoughPlayersForStart() bool {
	if len(g.Players) >= minPlayersForStart {
		return true
	}

	return false
}

func (g *Game) isEveryoneReady() bool {
	for _, p := range g.Players {
		if !p.IsReady {
			return false
		}
	}

	return true
}

func (g *Game) start() {
	g.State = Play
	g.moveOrder = NewMoveOrder(g.Players)
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
