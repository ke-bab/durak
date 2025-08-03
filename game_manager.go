package durak

import (
	"errors"
	"fmt"
	"sync"
)

const maxGames = 100

// GameManager is struct which holds list of all existing games on this server and
// allows to add or remove games to that list.
type GameManager struct {
	Games map[int]*Game
	lock  sync.Mutex
}

func (gm *GameManager) CreateGame() (*Game, error) {
	game := NewGame()
	gm.lock.Lock()
	defer gm.lock.Unlock()

	if err != nil {
		return nil, err
	}

	if _, ok := gm.Games[gameId]; ok {
		return nil, errors.New(fmt.Sprintf("game with id %d already exists", gameId))
	}
	gm.Games[gameId] = game

	return game, nil
}
