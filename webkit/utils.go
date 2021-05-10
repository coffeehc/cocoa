package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include <Foundation/NSGeometry.h>
// #include <Foundation/NSRange.h>
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal/utils"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

func toNSRect(rect foundation.Rect) C.NSRect {
	return C.NSRect{
		C.NSPoint{C.double(rect.Origin.X), C.double(rect.Origin.Y)},
		C.NSSize{C.double(rect.Size.Width), C.double(rect.Size.Height)},
	}
}

func toPointer(o objc.Object) unsafe.Pointer {
	if utils.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}
