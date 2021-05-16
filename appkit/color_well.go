package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "color_well.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorWell interface {
	Control
	TakeColorFrom(sender objc.Object)
	Activate(exclusive bool)
	Deactivate()
	DrawWellInside(insideRect foundation.Rect)
	Color() Color
	SetColor(value Color)
	IsActive() bool
	IsBordered() bool
	SetBordered(value bool)
}

type NSColorWell struct {
	NSControl
}

func MakeColorWell(ptr unsafe.Pointer) *NSColorWell {
	if ptr == nil {
		return nil
	}
	return &NSColorWell{
		NSControl: *MakeControl(ptr),
	}
}

func AllocColorWell() *NSColorWell {
	return MakeColorWell(C.C_ColorWell_Alloc())
}

func (n *NSColorWell) InitWithFrame(frameRect foundation.Rect) ColorWell {
	result := C.C_NSColorWell_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeColorWell(result)
}

func (n *NSColorWell) InitWithCoder(coder foundation.Coder) ColorWell {
	result := C.C_NSColorWell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeColorWell(result)
}

func (n *NSColorWell) Init() ColorWell {
	result := C.C_NSColorWell_Init(n.Ptr())
	return MakeColorWell(result)
}

func (n *NSColorWell) TakeColorFrom(sender objc.Object) {
	C.C_NSColorWell_TakeColorFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSColorWell) Activate(exclusive bool) {
	C.C_NSColorWell_Activate(n.Ptr(), C.bool(exclusive))
}

func (n *NSColorWell) Deactivate() {
	C.C_NSColorWell_Deactivate(n.Ptr())
}

func (n *NSColorWell) DrawWellInside(insideRect foundation.Rect) {
	C.C_NSColorWell_DrawWellInside(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(insideRect))))
}

func (n *NSColorWell) Color() Color {
	result := C.C_NSColorWell_Color(n.Ptr())
	return MakeColor(result)
}

func (n *NSColorWell) SetColor(value Color) {
	C.C_NSColorWell_SetColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSColorWell) IsActive() bool {
	result := C.C_NSColorWell_IsActive(n.Ptr())
	return bool(result)
}

func (n *NSColorWell) IsBordered() bool {
	result := C.C_NSColorWell_IsBordered(n.Ptr())
	return bool(result)
}

func (n *NSColorWell) SetBordered(value bool) {
	C.C_NSColorWell_SetBordered(n.Ptr(), C.bool(value))
}
