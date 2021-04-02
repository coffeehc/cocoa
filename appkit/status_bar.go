package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "status_bar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusBar interface {
	objc.Object
	IsVertical() bool
	Thickness() float64
	StatusItemWithLength(length float64) StatusItem
	RemoveStatusItem(item StatusItem)
}

var _ StatusBar = (*NSStatusBar)(nil)

type NSStatusBar struct {
	objc.NSObject
}

func MakeStatusBar(ptr unsafe.Pointer) *NSStatusBar {
	if ptr == nil {
		return nil
	}
	return &NSStatusBar{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (s *NSStatusBar) IsVertical() bool {
	return bool(C.StatusBar_IsVertical(s.Ptr()))
}

func (s *NSStatusBar) Thickness() float64 {
	return float64(C.StatusBar_Thickness(s.Ptr()))
}

func SystemStatusBar() StatusBar {
	return MakeStatusBar(C.StatusBar_SystemStatusBar())
}

func (s *NSStatusBar) StatusItemWithLength(length float64) StatusItem {
	return MakeStatusItem(C.StatusBar_StatusItemWithLength(s.Ptr(), C.double(length)))
}

func (s *NSStatusBar) RemoveStatusItem(item StatusItem) {
	C.StatusBar_RemoveStatusItem(s.Ptr(), toPointer(item))
}
