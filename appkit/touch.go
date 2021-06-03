package appkit

// #include "touch.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Touch interface {
	objc.Object
	LocationInView(view View) foundation.Point
	PreviousLocationInView(view View) foundation.Point
	Type() TouchType
	Identity() objc.Object
	Phase() TouchPhase
	NormalizedPosition() foundation.Point
	IsResting() bool
	Device() objc.Object
	DeviceSize() foundation.Size
}

type NSTouch struct {
	objc.NSObject
}

func MakeTouch(ptr unsafe.Pointer) NSTouch {
	return NSTouch{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTouch() NSTouch {
	return MakeTouch(C.C_Touch_Alloc())
}

func (n NSTouch) Init() Touch {
	result_ := C.C_NSTouch_Init(n.Ptr())
	return MakeTouch(result_)
}

func (n NSTouch) LocationInView(view View) foundation.Point {
	result_ := C.C_NSTouch_LocationInView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSTouch) PreviousLocationInView(view View) foundation.Point {
	result_ := C.C_NSTouch_PreviousLocationInView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSTouch) Type() TouchType {
	result_ := C.C_NSTouch_Type(n.Ptr())
	return TouchType(int(result_))
}

func (n NSTouch) Identity() objc.Object {
	result_ := C.C_NSTouch_Identity(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTouch) Phase() TouchPhase {
	result_ := C.C_NSTouch_Phase(n.Ptr())
	return TouchPhase(uint(result_))
}

func (n NSTouch) NormalizedPosition() foundation.Point {
	result_ := C.C_NSTouch_NormalizedPosition(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSTouch) IsResting() bool {
	result_ := C.C_NSTouch_IsResting(n.Ptr())
	return bool(result_)
}

func (n NSTouch) Device() objc.Object {
	result_ := C.C_NSTouch_Device(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTouch) DeviceSize() foundation.Size {
	result_ := C.C_NSTouch_DeviceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}
