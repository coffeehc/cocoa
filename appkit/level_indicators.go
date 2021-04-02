package appkit

type TickMarkPosition uint

const (
	TickMarkPositionBelow    TickMarkPosition = 0
	TickMarkPositionAbove    TickMarkPosition = 1
	TickMarkPositionLeading  TickMarkPosition = TickMarkPositionAbove
	TickMarkPositionTrailing TickMarkPosition = TickMarkPositionBelow
)

type LevelIndicatorStyle uint

const (
	LevelIndicatorStyleRelevancy LevelIndicatorStyle = iota
	LevelIndicatorStyleContinuousCapacity
	LevelIndicatorStyleDiscreteCapacity
	LevelIndicatorStyleRating
)
