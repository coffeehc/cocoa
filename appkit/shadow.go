package appkit

// #include "shadow.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Shadow interface {
	objc.Object
	Set()
	ShadowOffset() foundation.Size
	SetShadowOffset(value foundation.Size)
	ShadowBlurRadius() coregraphics.Float
	SetShadowBlurRadius(value coregraphics.Float)
	ShadowColor() Color
	SetShadowColor(value Color)
}

type NSShadow struct {
	objc.NSObject
}

func MakeShadow(ptr unsafe.Pointer) NSShadow {
	return NSShadow{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocShadow() NSShadow {
	return MakeShadow(C.C_Shadow_Alloc())
}

func (n NSShadow) Init() Shadow {
	result_ := C.C_NSShadow_Init(n.Ptr())
	return MakeShadow(result_)
}

func (n NSShadow) Set() {
	C.C_NSShadow_Set(n.Ptr())
}

func (n NSShadow) ShadowOffset() foundation.Size {
	result_ := C.C_NSShadow_ShadowOffset(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSShadow) SetShadowOffset(value foundation.Size) {
	C.C_NSShadow_SetShadowOffset(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSShadow) ShadowBlurRadius() coregraphics.Float {
	result_ := C.C_NSShadow_ShadowBlurRadius(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSShadow) SetShadowBlurRadius(value coregraphics.Float) {
	C.C_NSShadow_SetShadowBlurRadius(n.Ptr(), C.double(float64(value)))
}

func (n NSShadow) ShadowColor() Color {
	result_ := C.C_NSShadow_ShadowColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSShadow) SetShadowColor(value Color) {
	C.C_NSShadow_SetShadowColor(n.Ptr(), objc.ExtractPtr(value))
}
