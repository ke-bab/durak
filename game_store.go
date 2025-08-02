package durak

import (
	"errors"
	"fmt"
	"sync"
)

// GameStore is struct which holds list of all existing games on this server and
// allows to add or remove games to that list.
type GameStore struct {
	Games        map[int]*Game
	lock         sync.Mutex
	gameIdPool   *IdPool
	playerIdPool *IdPool
}

func NewGameStore() (*GameStore, error) {
	playerPool, err := NewIdPool(maxGames * maxPlayers)
	if err != nil {
		return nil, err
	}
	gamePool, err := NewIdPool(maxGames)
	if err != nil {
		return nil, err
	}

	return &GameStore{
		Games:        make(map[int]*Game),
		playerIdPool: playerPool,
		gameIdPool:   gamePool,
	}, nil
}

func (gm *GameStore) CreateGame() (*Game, error) {
	game := NewGame()
	gameId, ok := gm.gameIdPool.Reserve()
	if !ok {
		return nil, errors.New("no free id for new game")
	}
	gm.lock.Lock()
	defer gm.lock.Unlock()
	if _, ok := gm.Games[gameId]; ok {
		return nil, errors.New(fmt.Sprintf("game with id %d already exists", gameId))
	}
	gm.Games[gameId] = game

	return game, nil
}
