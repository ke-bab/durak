package durak

type Player struct {
	ID      int
	IsReady bool
}

func NewPlayer(id int) *Player {
	return &Player{ID: id}
}
