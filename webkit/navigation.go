package webkit

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

func MakeNavigation(ptr unsafe.Pointer) WKNavigation {
	return WKNavigation{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocNavigation() WKNavigation {
	result_ := C.C_WKNavigation_AllocNavigation()
	return MakeNavigation(result_)
}

func (w WKNavigation) Init() WKNavigation {
	result_ := C.C_WKNavigation_Init(w.Ptr())
	return MakeNavigation(result_)
}

func NewNavigation() WKNavigation {
	result_ := C.C_WKNavigation_NewNavigation()
	return MakeNavigation(result_)
}

func (w WKNavigation) Autorelease() WKNavigation {
	result_ := C.C_WKNavigation_Autorelease(w.Ptr())
	return MakeNavigation(result_)
}

func (w WKNavigation) Retain() WKNavigation {
	result_ := C.C_WKNavigation_Retain(w.Ptr())
	return MakeNavigation(result_)
}

func (w WKNavigation) EffectiveContentMode() ContentMode {
	result_ := C.C_WKNavigation_EffectiveContentMode(w.Ptr())
	return ContentMode(int(result_))
}
