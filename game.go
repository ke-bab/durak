package durak

import (
	"errors"
	"fmt"
	"sync"
)

const maxPlayers = 4 // in one game
const maxGames = 100

type Game struct {
	Players      map[int]*Player
	State        GameState
	lock         sync.Mutex
	playerIdPool *IdPool
}

func NewGame(pool *IdPool) *Game {
	return &Game{
		Players:      make(map[int]*Player, maxPlayers),
		State:        Open,
		playerIdPool: pool,
	}
}

func (g *Game) JoinPlayer() (*Player, error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.State != Open {
		return nil, errors.New("game is not open for players")
	}

	if g.isFull() {
		return nil, errors.New("game already full")
	}

	playerId, ok := g.playerIdPool.getId()
	if !ok {
		return nil, errors.New("no free id in player id pool")
	}

	if g.hasPlayer(playerId) {
		return nil, errors.New(fmt.Sprintf("player %d already joined", playerId))
	}

	newPlayer := NewPlayer(playerId)
	g.Players[playerId] = newPlayer

	if len(g.Players) == maxPlayers {
		g.start()
	}

	return newPlayer, nil
}

func (g *Game) PlayerIsReady(id int) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.State != Open {
		return errors.New("game is not open for players")
	}

	player, ok := g.Players[id]
	if !ok {
		return errors.New(fmt.Sprintf("player %d not found", id))
	}

	player.IsReady = true

	if g.everyoneIsReady() {
		g.start()
	}

	return nil
}

func (g *Game) everyoneIsReady() bool {
	for _, p := range g.Players {
		if !p.IsReady {
			return false
		}
	}

	return true
}

func (g *Game) start() {
	g.State = Play
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
