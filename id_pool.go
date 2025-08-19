package durak

import (
	"errors"
	"fmt"
	"sync"
)

const minId = 1

// IdPool for this project we don't care about database and stuff so IDs are just numbers in particular range.
type IdPool struct {
	pool  map[int]struct{}
	mu    sync.Mutex
	maxId int
}

func NewIdPool(maxId int) (*IdPool, error) {
	if maxId < minId {
		return nil, errors.New("maxId argument is less than minId")
	}
	pool := make(map[int]struct{}, maxId)
	for i := minId; i <= maxId; i++ {
		pool[i] = struct{}{}
	}

	return &IdPool{pool: pool, maxId: maxId}, nil
}

func (p *IdPool) Acquire() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for k := range p.pool {
		delete(p.pool, k)
		return k, nil
	}

	return 0, errors.New("no ids in pool")
}

func (p *IdPool) Release(id int) error {
	if id < minId || id > p.maxId {
		return fmt.Errorf("id is out of bounds: min %d, max %d", minId, p.maxId)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.pool[id]
	if ok {
		return errors.New("id already returned")
	}

	p.pool[id] = struct{}{}

	return nil
}
