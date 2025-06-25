package durak

type Action interface {
	CanBeApplied(g *Game) bool
	Apply(g *Game) error
	Name() string
}
