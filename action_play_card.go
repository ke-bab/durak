package durak

import (
	"errors"
	"fmt"
)

type PlayCardAction struct {
	PlayerId int
	Card     *Card
	//
	player    *Player
	cardIndex int
}

func (a *PlayCardAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Play {
		return false, errors.New("game is in wrong state for playing cards")
	}

	player, ok := g.Players[a.PlayerId]
	if !ok {
		return false, fmt.Errorf("player %d not found", a.PlayerId)
	}
	a.player = player

	if g.moveOrder.Current != player {
		return false, fmt.Errorf("it is not player's %d turn", a.PlayerId)
	}

	index, ok := player.hasCard(a.Card)
	if !ok {
		return false, fmt.Errorf("player %d has no Card %s %s", player.ID, a.Card.Rank, a.Card.Suit)
	}
	a.cardIndex = index

	return true, nil
}

func (a *PlayCardAction) Apply(g *Game) {
	cards := CardCollection(g.CardsOnTable)
	cards.add(a.Card)
	a.player.removeExistingCard(a.cardIndex)
}

func (a *PlayCardAction) Name() ActionName {
	return ActionPlayCard
}
