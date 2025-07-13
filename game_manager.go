package durak

import (
	"errors"
	"fmt"
	"sync"
)

type GameManager struct {
	Games        map[int]*Game
	lock         sync.Mutex
	gameIdPool   *IdPool
	playerIdPool *IdPool
}

func NewGameManager() (*GameManager, error) {
	playerPool, err := NewIdPool(maxGames * maxPlayers)
	if err != nil {
		return nil, err
	}
	gamePool, err := NewIdPool(maxGames)
	if err != nil {
		return nil, err
	}

	return &GameManager{
		Games:        make(map[int]*Game),
		playerIdPool: playerPool,
		gameIdPool:   gamePool,
	}, nil
}

func (gm *GameManager) CreateGame() (*Game, error) {
	game := NewGame(gm.playerIdPool)
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
