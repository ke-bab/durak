package durak

import "fmt"

type PlayerFactory struct {
	pool *IdPool
}

func NewPlayerFactory(pool *IdPool) *PlayerFactory {
	return &PlayerFactory{pool: pool}
}

func (f *PlayerFactory) CreatePlayer() (*Player, error) {
	id, ok := f.pool.Reserve()
	if !ok {
		return nil, fmt.Errorf("no ids in pool")
	}

	return NewPlayer(id), nil
}
