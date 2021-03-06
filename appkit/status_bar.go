package appkit

// #include "status_bar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusBar interface {
	objc.Object
	StatusItemWithLength(length coregraphics.Float) StatusItem
	RemoveStatusItem(item StatusItem)
	IsVertical() bool
	Thickness() coregraphics.Float
}

type NSStatusBar struct {
	objc.NSObject
}

func MakeStatusBar(ptr unsafe.Pointer) NSStatusBar {
	return NSStatusBar{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocStatusBar() NSStatusBar {
	result_ := C.C_NSStatusBar_AllocStatusBar()
	return MakeStatusBar(result_)
}

func (n NSStatusBar) Init() NSStatusBar {
	result_ := C.C_NSStatusBar_Init(n.Ptr())
	return MakeStatusBar(result_)
}

func NewStatusBar() NSStatusBar {
	result_ := C.C_NSStatusBar_NewStatusBar()
	return MakeStatusBar(result_)
}

func (n NSStatusBar) Autorelease() NSStatusBar {
	result_ := C.C_NSStatusBar_Autorelease(n.Ptr())
	return MakeStatusBar(result_)
}

func (n NSStatusBar) Retain() NSStatusBar {
	result_ := C.C_NSStatusBar_Retain(n.Ptr())
	return MakeStatusBar(result_)
}

func (n NSStatusBar) StatusItemWithLength(length coregraphics.Float) StatusItem {
	result_ := C.C_NSStatusBar_StatusItemWithLength(n.Ptr(), C.double(float64(length)))
	return MakeStatusItem(result_)
}

func (n NSStatusBar) RemoveStatusItem(item StatusItem) {
	C.C_NSStatusBar_RemoveStatusItem(n.Ptr(), objc.ExtractPtr(item))
}

func SystemStatusBar() StatusBar {
	result_ := C.C_NSStatusBar_SystemStatusBar()
	return MakeStatusBar(result_)
}

func (n NSStatusBar) IsVertical() bool {
	result_ := C.C_NSStatusBar_IsVertical(n.Ptr())
	return bool(result_)
}

func (n NSStatusBar) Thickness() coregraphics.Float {
	result_ := C.C_NSStatusBar_Thickness(n.Ptr())
	return coregraphics.Float(float64(result_))
}
