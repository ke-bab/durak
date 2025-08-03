package durak

const maxPlayers = 2

// Game is a pure state model of game.
// It has only fields which represents different game states.
// Validity of these states is done by "state" type structs.
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

type GameState string

const (
	Open   GameState = "open"
	Play   GameState = "play"
	Closed GameState = "closed"
)
