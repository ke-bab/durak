package durak

type MoveOrder struct {
	Current *Player
	Next    *MoveOrder
}

// NewMoveOrder makes circular single linked list which used to define order for player moves.
// p1 -> p2 -> p3 -> p4 -> p1 and so on.
func NewMoveOrder(list map[int]*Player) *MoveOrder {
	var first *MoveOrder
	var prev *MoveOrder
	for _, p := range list {
		current := &MoveOrder{
			Current: p,
			Next:    nil,
		}
		if prev != nil {
			// we are not in first element
			prev.Next = current
		} else {
			// we are in first element
			first = current
		}
		prev = current
	}
	// last node
	prev.Next = first

	return first
}
