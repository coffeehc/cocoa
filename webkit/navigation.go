package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
// #include "navigation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Navigation interface {
	objc.Object
	EffectiveContentMode() ContentMode
}

type WKNavigation struct {
	objc.NSObject
}

func MakeNavigation(ptr unsafe.Pointer) *WKNavigation {
	if ptr == nil {
		return nil
	}
	return &WKNavigation{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocNavigation() *WKNavigation {
	return MakeNavigation(C.C_Navigation_Alloc())
}

func (w *WKNavigation) Init() Navigation {
	result_ := C.C_WKNavigation_Init(w.Ptr())
	return MakeNavigation(result_)
}

func (w *WKNavigation) EffectiveContentMode() ContentMode {
	result_ := C.C_WKNavigation_EffectiveContentMode(w.Ptr())
	return ContentMode(int(result_))
}
