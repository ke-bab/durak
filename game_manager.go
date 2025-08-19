package durak

import (
	"errors"
	"fmt"
	"sync"
)

const maxGames = 100

type GameManager struct {
	Games     map[int]*Game
	mu        sync.RWMutex
	gameIds   *IdPool
	playerIds *IdPool
}

func NewGameManager() (*GameManager, error) {
	gamePool, err := NewIdPool(maxGames)
	if err != nil {
		return nil, err
	}
	playerPool, err := NewIdPool(maxGames * playersInGame)
	if err != nil {
		return nil, err
	}

	return &GameManager{
		Games:     make(map[int]*Game, maxGames),
		gameIds:   gamePool,
		playerIds: playerPool,
	}, nil
}

func (gm *GameManager) CreateGame() (*Game, error) {
	game := NewGame()
	gm.mu.Lock()
	defer gm.mu.Unlock()

	gameId, err := gm.gameIds.Acquire()
	if err != nil {
		return nil, err
	}
	// return id if something goes wrong
	defer recoverGameId(gameId, gm)

	if _, ok := gm.Games[gameId]; ok {
		return nil, handleErrGameAlreadyExists(gameId, gm)
	}

	gm.Games[gameId] = game

	return game, nil
}

func (gm *GameManager) Find(id int) (*Game, error) {
	gm.mu.RLock()
	defer gm.mu.Unlock()
	game, ok := gm.Games[id]
	if !ok {
		return nil, errors.New("game not found")
	}

	return game, nil
}

/////////////
// utility //
/////////////

func recoverGameId(gameId int, gm *GameManager) {
	if r := recover(); r != nil {
		_ = gm.gameIds.Release(gameId) // ignore error because it will never happen (of course).
	}
}

func handleErrGameAlreadyExists(gameId int, gm *GameManager) error {
	errs := make([]error, 0)
	err := gm.gameIds.Release(gameId)
	if err != nil {
		errs = append(errs, err)
	}
	errs = append(errs, errors.New(fmt.Sprintf("game with id %d already exists", gameId)))

	return errors.Join(errs...)
}
