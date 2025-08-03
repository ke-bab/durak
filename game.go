package durak

import (
	"errors"
	"fmt"
)

const playersInGame = 2

type Game struct {
	Player1  *Player
	Player2  *Player
	State    GameState `json:"state"`
	Deck     Deck      `json:"deck"`
	Attacker *Player
}

func NewGame() *Game {
	return &Game{
		State: Open,
		Deck:  initDeck(),
	}
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

func (g *Game) CanJoinState() (*StateCanJoin, error) {
	if g.State != Open {
		if g.State != Open {
			return nil, errors.New("game is not open")
		}
	}

	if g.Player1 == nil {
		return &StateCanJoin{Slot: &g.Player1}, nil
	}

	if g.Player2 == nil {
		return &StateCanJoin{Slot: &g.Player2}, nil
	}

	return nil, fmt.Errorf("no free slot to join")
}
