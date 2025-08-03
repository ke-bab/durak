package durak

type StateCanJoin struct {
	factory *PlayerFactory
	Slot    **Player
}

func (s *StateCanJoin) Join() (*Player, error) {
	p, err := s.factory.CreatePlayer()
	if err != nil {
		return nil, err
	}
	*s.Slot = p

	return p, nil
}
