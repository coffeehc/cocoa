package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Appkit
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
)

var resources = foundation.NewResourceRegistry()
