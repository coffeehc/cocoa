package appkit

// #include "storyboard.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Storyboard interface {
	objc.Object
	InstantiateInitialController() objc.Object
}

type NSStoryboard struct {
	objc.NSObject
}

func MakeStoryboard(ptr unsafe.Pointer) *NSStoryboard {
	if ptr == nil {
		return nil
	}
	return &NSStoryboard{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocStoryboard() *NSStoryboard {
	return MakeStoryboard(C.C_Storyboard_Alloc())
}

func (n *NSStoryboard) Init() Storyboard {
	result_ := C.C_NSStoryboard_Init(n.Ptr())
	return MakeStoryboard(result_)
}

func StoryboardWithName_Bundle(name StoryboardName, storyboardBundleOrNil foundation.Bundle) Storyboard {
	result_ := C.C_NSStoryboard_StoryboardWithName_Bundle(foundation.NewString(string(name)).Ptr(), objc.ExtractPtr(storyboardBundleOrNil))
	return MakeStoryboard(result_)
}

func (n *NSStoryboard) InstantiateInitialController() objc.Object {
	result_ := C.C_NSStoryboard_InstantiateInitialController(n.Ptr())
	return objc.MakeObject(result_)
}

func MainStoryboard() Storyboard {
	result_ := C.C_NSStoryboard_MainStoryboard()
	return MakeStoryboard(result_)
}
