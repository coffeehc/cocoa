package appkit

import "math"

// StackViewGravity is the gravity areas available in a stack view.
type StackViewGravity int

const (
	StackViewGravityTop      StackViewGravity = 1
	StackViewGravityLeading  StackViewGravity = 1
	StackViewGravityCenter   StackViewGravity = 2
	StackViewGravityBottom   StackViewGravity = 3
	StackViewGravityTrailing StackViewGravity = 3
)

const StackViewSpacingUseDefault = math.MaxFloat32

type StackViewDistribution int

const (
	NSStackViewDistributionGravityAreas StackViewDistribution = -1
	NSStackViewDistributionFill         StackViewDistribution = iota
	NSStackViewDistributionFillEqually
	NSStackViewDistributionFillProportionally
	NSStackViewDistributionEqualSpacing
	NSStackViewDistributionEqualCentering
)

type StackViewVisibilityPriority float32

const (
	NSStackViewVisibilityPriorityMustHold              StackViewVisibilityPriority = 1000
	NSStackViewVisibilityPriorityDetachOnlyIfNecessary StackViewVisibilityPriority = 900
	NSStackViewVisibilityPriorityNotVisible            StackViewVisibilityPriority = 0
)
