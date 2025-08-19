package durak

import (
	"errors"
	"sync"
)

type GameState string

const (
	Open   GameState = "open"
	Play   GameState = "play"
	Closed GameState = "closed"
)

const playersInGame = 2

type Game struct {
	ID       int `json:"ID"`
	Player1  *Player
	Player2  *Player
	State    GameState `json:"state"`
	Deck     Deck      `json:"deck"`
	Attacker *Player

	mu sync.RWMutex
}

func NewGame(id int) *Game {
	return &Game{
		ID:    id,
		State: Open,
		Deck:  initDeck(),
	}
}

func (g *Game) Join() error {

	// WRONG! rework with services and DI.
	slot, ok := g.findFreePlayerSlot()
	if !ok {
		return errors.New("game is full")
	}

	gameId, err := g.Acquire()
	if err != nil {
		return nil, err
	}
	// return id if something goes wrong
	defer recoverGameId(gameId, gm)

	*slot = NewPlayer()
}

func (g *Game) start() {
	g.State = Play
	g.Attacker = g.Player1
	g.dealCardsOnStart()
}

func (g *Game) dealCardsOnStart() {
	deck := CardCollection(g.Deck)
	g.Player1.Hand = deck.takeXCards(6)
	g.Player2.Hand = deck.takeXCards(6)
}

func initDeck() []*Card {
	d := make([]*Card, 0, len(suits)*len(ranks))
	for _, s := range suits {
		for _, r := range ranks {
			d = append(d, NewCard(s, r))
		}
	}

	return d
}

// utility

func (g *Game) findFreePlayerSlot() (**Player, bool) {
	if g.Player1 == nil {
		return &g.Player1, true
	}
	if g.Player2 == nil {
		return &g.Player2, true
	}

	return nil, false
}
