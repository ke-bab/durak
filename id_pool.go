package durak

import (
	"errors"
	"fmt"
)

const minId = 1

// for this project we don't care about database and stuff so IDs are just numbers in particular range.
type IdPool struct {
	pool map[int]bool
	max  int
}

func NewIdPool(maxId int) (*IdPool, error) {
	if maxId < minId {
		return nil, errors.New("maxId argument is less than minId")
	}
	pool := make(map[int]bool, maxId)
	for i := 1; i <= maxId; i++ {
		pool[i] = true
	}

	return &IdPool{pool: pool, max: maxId}, nil
}

func (p *IdPool) Reserve() (int, bool) {
	for k, v := range p.pool {
		if v == true {
			p.pool[k] = false
			return k, true
		}
	}

	return 0, false
}

func (p *IdPool) Release(id int) error {
	if id < minId || id > p.max {
		return errors.New(fmt.Sprintf("id is out of bounds: min %d, max %d", minId, p.max))
	}
	exist, _ := p.pool[id]
	if exist {
		return errors.New("id already returned")
	}

	p.pool[id] = true

	return nil
}
