package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "event.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Event interface {
	objc.Object
	LocationInWindow() foundation.Point
	Window() Window
	WindowNumber() int64
	Timestamp() float64
}

var _ Event = (*NSEvent)(nil)

type NSEvent struct {
	objc.NSObject
}

func MakeEvent(ptr unsafe.Pointer) *NSEvent {
	if ptr == nil {
		return nil
	}
	return &NSEvent{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (e *NSEvent) LocationInWindow() foundation.Point {
	return toPoint(C.Event_LocationInWindow(e.Ptr()))
}

func (e *NSEvent) Window() Window {
	return MakeWindow(C.Event_Window(e.Ptr()))
}

func (e *NSEvent) WindowNumber() int64 {
	return int64(C.Event_WindowNumber(e.Ptr()))
}

func (e *NSEvent) Timestamp() float64 {
	return float64(C.Event_Timestamp(e.Ptr()))
}
