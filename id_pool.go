package durak

import (
	"errors"
	"fmt"
)

const minId = 1

// for this project we don't care about database and stuff so IDs are just numbers in particular range.
type IdPool struct {
	pool chan int
	max  int
}

func NewIdPool(maxId int) (*IdPool, error) {
	if maxId < minId {
		return nil, errors.New("maxId argument is less than minId")
	}
	pool := make(chan int, maxId)
	for i := 1; i <= maxId; i++ {
		pool <- i
	}

	return &IdPool{pool: pool, max: maxId}, nil
}

func (p *IdPool) Reserve() (int, bool) {
	select {
	case id := <-p.pool:
		return id, true
	default:
		return 0, false
	}
}

func (p *IdPool) Release(id int) error {
	if id < minId || id > p.max {
		return errors.New(fmt.Sprintf("id is out of bounds: min %d, max %d", minId, p.max))
	}
	// maybe add unique check?
	select {
	case p.pool <- id:
		return nil
	default:
		return fmt.Errorf("cannot release id %d: channel is full", id)
	}
}
