package durak

import (
	"errors"
	"fmt"
	"sync"
)

const minId = 1

// IdPool for this project we don't care about database and stuff so IDs are just numbers in particular range.
type IdPool struct {
	pool  map[int]bool
	mu    sync.Mutex
	maxId int
}

func NewIdPool(maxId int) (*IdPool, error) {
	if maxId < minId {
		return nil, errors.New("maxId argument is less than minId")
	}
	pool := make(map[int]bool, maxId)
	for i := 1; i <= maxId; i++ {
		pool[i] = true
	}

	return &IdPool{pool: pool, maxId: maxId}, nil
}

func (p *IdPool) Acquire() (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for k, v := range p.pool {
		if v == true {
			p.pool[k] = false
			return k, nil
		}
	}

	return 0, errors.New("no free game id")
}

func (p *IdPool) Release(id int) error {
	if id < minId || id > p.maxId {
		return errors.New(fmt.Sprintf("id is out of bounds: min %d, max %d", minId, p.maxId))
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	exist, _ := p.pool[id]
	if exist {
		return errors.New("id already returned")
	}

	p.pool[id] = true

	return nil
}
