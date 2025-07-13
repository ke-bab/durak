package durak

import (
	"errors"
	"fmt"
)

type PlayCardAction struct {
	Player *Player
	Card   *Card
	//
	cardIndex int
}

func (a *PlayCardAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Play {
		return false, errors.New("game is in wrong state for playing cards")
	}

	if !g.hasPlayer(a.Player) {
		return false, fmt.Errorf("player %d not found", a.Player.ID)
	}

	if g.attacker != a.Player {
		return false, fmt.Errorf("it is not player's %d turn", a.Player.ID)
	}

	index, ok := a.Player.hasCard(a.Card)
	if !ok {
		return false, fmt.Errorf("player %d has no Card %s %s", a.Player.ID, a.Card.Rank, a.Card.Suit)
	}
	a.cardIndex = index

	return true, nil
}

func (a *PlayCardAction) Apply(g *Game) {
	cards := CardCollection(g.CardsOnTable)
	cards.add(a.Card)
	a.Player.removeExistingCard(a.cardIndex)
}

func (a *PlayCardAction) Name() ActionName {
	return ActionPlayCard
}
