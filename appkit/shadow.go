package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "shadow.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Shadow interface {
	objc.Object
	Set()
	ShadowOffset() coregraphics.Size
	SetShadowOffset(value coregraphics.Size)
	ShadowBlurRadius() coregraphics.Float
	SetShadowBlurRadius(value coregraphics.Float)
	ShadowColor() objc.Object
	SetShadowColor(value objc.Object)
}

type NSShadow struct {
	objc.NSObject
}

func MakeShadow(ptr unsafe.Pointer) *NSShadow {
	if ptr == nil {
		return nil
	}
	return &NSShadow{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocShadow() *NSShadow {
	return MakeShadow(C.C_Shadow_Alloc())
}

func (n *NSShadow) Init() Shadow {
	result := C.C_NSShadow_Init(n.Ptr())
	return MakeShadow(result)
}

func (n *NSShadow) Set() {
	C.C_NSShadow_Set(n.Ptr())
}

func (n *NSShadow) ShadowOffset() coregraphics.Size {
	result := C.C_NSShadow_ShadowOffset(n.Ptr())
	return coregraphics.FromCGSizePointer(unsafe.Pointer(&result))
}

func (n *NSShadow) SetShadowOffset(value coregraphics.Size) {
	C.C_NSShadow_SetShadowOffset(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(value)))
}

func (n *NSShadow) ShadowBlurRadius() coregraphics.Float {
	result := C.C_NSShadow_ShadowBlurRadius(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSShadow) SetShadowBlurRadius(value coregraphics.Float) {
	C.C_NSShadow_SetShadowBlurRadius(n.Ptr(), C.double(float64(value)))
}

func (n *NSShadow) ShadowColor() objc.Object {
	result := C.C_NSShadow_ShadowColor(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSShadow) SetShadowColor(value objc.Object) {
	C.C_NSShadow_SetShadowColor(n.Ptr(), objc.ExtractPtr(value))
}
