package webkit

// #include "content_world.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ContentWorld interface {
	objc.Object
	Name() string
}

type WKContentWorld struct {
	objc.NSObject
}

func MakeContentWorld(ptr unsafe.Pointer) WKContentWorld {
	return WKContentWorld{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocContentWorld() WKContentWorld {
	return MakeContentWorld(C.C_ContentWorld_Alloc())
}

func ContentWorld_WorldWithName(name string) ContentWorld {
	result_ := C.C_WKContentWorld_ContentWorld_WorldWithName(foundation.NewString(name).Ptr())
	return MakeContentWorld(result_)
}

func ContentWorld_DefaultClientWorld() ContentWorld {
	result_ := C.C_WKContentWorld_ContentWorld_DefaultClientWorld()
	return MakeContentWorld(result_)
}

func ContentWorld_PageWorld() ContentWorld {
	result_ := C.C_WKContentWorld_ContentWorld_PageWorld()
	return MakeContentWorld(result_)
}

func (w WKContentWorld) Name() string {
	result_ := C.C_WKContentWorld_Name(w.Ptr())
	return foundation.MakeString(result_).String()
}
