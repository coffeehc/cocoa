package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	Phase() TouchPhase
	NormalizedPosition() foundation.Point
	IsResting() bool
	Device() objc.Object
	DeviceSize() foundation.Size
}

type NSTouch struct {
	objc.NSObject
}

func MakeTouch(ptr unsafe.Pointer) *NSTouch {
	if ptr == nil {
		return nil
	}
	return &NSTouch{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTouch() *NSTouch {
	return MakeTouch(C.C_Touch_Alloc())
}

func (n *NSTouch) Init() Touch {
	result := C.C_NSTouch_Init(n.Ptr())
	return MakeTouch(result)
}

func (n *NSTouch) LocationInView(view View) foundation.Point {
	result := C.C_NSTouch_LocationInView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSTouch) PreviousLocationInView(view View) foundation.Point {
	result := C.C_NSTouch_PreviousLocationInView(n.Ptr(), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSTouch) Type() TouchType {
	result := C.C_NSTouch_Type(n.Ptr())
	return TouchType(int(result))
}

func (n *NSTouch) Phase() TouchPhase {
	result := C.C_NSTouch_Phase(n.Ptr())
	return TouchPhase(uint(result))
}

func (n *NSTouch) NormalizedPosition() foundation.Point {
	result := C.C_NSTouch_NormalizedPosition(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSTouch) IsResting() bool {
	result := C.C_NSTouch_IsResting(n.Ptr())
	return bool(result)
}

func (n *NSTouch) Device() objc.Object {
	result := C.C_NSTouch_Device(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSTouch) DeviceSize() foundation.Size {
	result := C.C_NSTouch_DeviceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}
