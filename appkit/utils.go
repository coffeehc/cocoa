package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Appkit
import "C"
import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal/utils"
	"github.com/hsiafan/cocoa/objc"
)

var resources = foundation.NewResourceRegistry()

func toPointer(o objc.PointerHolder) unsafe.Pointer {
	if utils.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}
