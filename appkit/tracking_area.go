package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "tracking_area.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TrackingArea interface {
	objc.Object
	Options() TrackingAreaOptions
	Owner() objc.Object
	Rect() foundation.Rect
}

type NSTrackingArea struct {
	objc.NSObject
}

func MakeTrackingArea(ptr unsafe.Pointer) *NSTrackingArea {
	if ptr == nil {
		return nil
	}
	return &NSTrackingArea{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTrackingArea() *NSTrackingArea {
	return MakeTrackingArea(C.C_TrackingArea_Alloc())
}

func (n *NSTrackingArea) Init() TrackingArea {
	result_ := C.C_NSTrackingArea_Init(n.Ptr())
	return MakeTrackingArea(result_)
}

func (n *NSTrackingArea) Options() TrackingAreaOptions {
	result_ := C.C_NSTrackingArea_Options(n.Ptr())
	return TrackingAreaOptions(uint(result_))
}

func (n *NSTrackingArea) Owner() objc.Object {
	result_ := C.C_NSTrackingArea_Owner(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSTrackingArea) Rect() foundation.Rect {
	result_ := C.C_NSTrackingArea_Rect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}
