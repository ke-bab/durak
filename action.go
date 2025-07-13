package durak

type ActionName string

const (
	ActionJoin     ActionName = "join"
	ActionReady    ActionName = "ready"
	ActionPlayCard ActionName = "play_card"
	ActionBeatCard ActionName = "beat_card"
	ActionTakeCard ActionName = "take_card"
)

type Action interface {
	CanBeApplied(g *Game) (bool, error)
	Apply(g *Game)
	Name() ActionName
}
