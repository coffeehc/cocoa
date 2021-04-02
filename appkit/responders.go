package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

// ActionHandler is a callback function for an Action.
type ActionHandler func(sender objc.Object)
