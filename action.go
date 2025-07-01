package durak

type ActionName string

const (
	ActionJoin     = "join"
	ActionReady    = "ready"
	ActionPlayCard = "play_card"
)

type Action interface {
	CanBeApplied(g *Game) (bool, error)
	Apply(g *Game)
	Name() ActionName
}
