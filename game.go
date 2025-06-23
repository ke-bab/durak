package durak

import (
	"errors"
	"fmt"
	"sync"
)

const maxPlayers = 4 // in one game
const maxGames = 100
const minPlayersForStart = 2

type Game struct {
	Players      map[int]*Player
	State        GameState
	CardsOnTable CardsOnTable
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
		return nil, fmt.Errorf("player %d already joined", playerId)
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
		return fmt.Errorf("player %d not found", id)
	}

	player.IsReady = true

	if g.isEnoughPlayersForStart() && g.isEveryoneReady() {
		g.start()
	}

	return nil
}

func (g *Game) PlayerPlaysCard(id int, card *Card) error {
	if g.State != Play {
		return errors.New("game is in wrong state for playing cards")
	}

	player, ok := g.Players[id]
	if !ok {
		return fmt.Errorf("player %d not found", id)
	}

	if g.moveOrder.Current != player {
		return fmt.Errorf("it is not player's %d turn", id)
	}

	if !player.hasCard(card) {
		return fmt.Errorf("player %d has no card %s %s", player.ID, card.Rank, card.Suit)
	}

	// player's turn
	player.playCard(card, g.CardsOnTable)
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
