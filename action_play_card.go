package durak

import (
	"errors"
	"fmt"
)

type PlayCardAction struct {
	playerId  int
	card      *Card
	player    *Player
	cardIndex int
}

func (a *PlayCardAction) CanBeApplied(g *Game) (bool, error) {
	if g.State != Play {
		return false, errors.New("game is in wrong state for playing cards")
	}

	player, ok := g.Players[a.playerId]
	if !ok {
		return false, fmt.Errorf("player %d not found", a.playerId)
	}
	a.player = player

	if g.moveOrder.Current != player {
		return false, fmt.Errorf("it is not player's %d turn", a.playerId)
	}

	index, ok := player.hasCard(a.card)
	if !ok {
		return false, fmt.Errorf("player %d has no card %s %s", player.ID, a.card.Rank, a.card.Suit)
	}
	a.cardIndex = index

	return true, nil
}

func (a *PlayCardAction) Apply(g *Game) {
	cards := CardCollection(g.CardsOnTable)
	cards.add(a.card)
	a.player.removeExistingCard(a.cardIndex)
}

func (a *PlayCardAction) Name() ActionName {
	return ActionPlayCard
}
