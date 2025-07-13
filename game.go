package durak

import (
	"sync"
)

const maxPlayers = 4 // in one game
const maxGames = 100
const minPlayersForStart = 2

type Game struct {
	Players      map[int]*Player `json:"players"`
	State        GameState       `json:"state"`
	CardsOnTable []*Card         `json:"cardsOnTable"`
	Deck         []*Card         `json:"deck"`

	attacker     *Player
	defender     *Player
	lock         sync.Mutex
	playerIdPool *IdPool
}

func NewGame(pool *IdPool) *Game {
	return &Game{
		Players:      make(map[int]*Player, maxPlayers),
		State:        Open,
		playerIdPool: pool,
		CardsOnTable: make([]*Card, 0),
		Deck:         initDeck(),
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
	g.attacker = g.Players[0]
	g.defender = g.Players[1]
	g.dealCardsOnStart()
}

func (g *Game) dealCardsOnStart() {
	deck := CardCollection(g.Deck)
	for _, p := range g.Players {
		p.Hand = deck.takeXCards(6)
	}
}

func (g *Game) hasPlayer(p *Player) bool {
	for _, player := range g.Players {
		if player == p {
			return true
		}
	}

	return false
}

func (g *Game) isFull() bool {
	if len(g.Players) >= maxPlayers {
		return true
	}

	return false
}

func initDeck() []*Card {
	suits := suits()
	ranks := ranks()
	d := make([]*Card, 0, len(suits)*len(ranks))
	for _, s := range suits {
		for _, r := range ranks {
			d = append(d, NewCard(s, r))
		}
	}

	return d
}
