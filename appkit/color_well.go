package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "color_well.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorWell interface {
	Control
	Color() Color
	SetColor(color Color)
	IsActive() bool
	IsBordered() bool
	SetBordered(bordered bool)
	TakeColorFrom(sender objc.Object)
	Activate(exclusive bool)
	Deactivate()
	DrawWellInside(insideRect foundation.Rect)
}

var _ ColorWell = (*NSColorWell)(nil)

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

func (c *NSColorWell) Color() Color {
	return MakeColor(C.ColorWell_Color(c.Ptr()))
}

func (c *NSColorWell) SetColor(color Color) {
	C.ColorWell_SetColor(c.Ptr(), toPointer(color))
}

func (c *NSColorWell) IsActive() bool {
	return bool(C.ColorWell_IsActive(c.Ptr()))
}

func (c *NSColorWell) IsBordered() bool {
	return bool(C.ColorWell_IsBordered(c.Ptr()))
}

func (c *NSColorWell) SetBordered(bordered bool) {
	C.ColorWell_SetBordered(c.Ptr(), C.bool(bordered))
}

func NewColorWell(frame foundation.Rect) ColorWell {
	return MakeColorWell(C.ColorWell_NewColorWell(toNSRect(frame)))
}

func (c *NSColorWell) TakeColorFrom(sender objc.Object) {
	C.ColorWell_TakeColorFrom(c.Ptr(), toPointer(sender))
}

func (c *NSColorWell) Activate(exclusive bool) {
	C.ColorWell_Activate(c.Ptr(), C.bool(exclusive))
}

func (c *NSColorWell) Deactivate() {
	C.ColorWell_Deactivate(c.Ptr())
}

func (c *NSColorWell) DrawWellInside(insideRect foundation.Rect) {
	C.ColorWell_DrawWellInside(c.Ptr(), toNSRect(insideRect))
}
