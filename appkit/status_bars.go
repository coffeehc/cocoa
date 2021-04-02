package appkit

type StatusItemBehavior uint

const (
	StatusItemBehaviorRemovalAllowed       StatusItemBehavior = (1 << 1)
	StatusItemBehaviorTerminationOnRemoval StatusItemBehavior = (1 << 2)
)
const (
	SquareStatusItemLength   float64 = -2.0
	VariableStatusItemLength float64 = -1.0
)
