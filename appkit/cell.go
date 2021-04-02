package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Cell interface {
	objc.Object
	Wraps() bool
	SetWraps(wraps bool)
	IsScrollable() bool
	SetScrollable(scrollable bool)
	IsEditable() bool
	SetEditable(editable bool)
	IsSelectable() bool
	SetSelectable(selectable bool)
}

var _ Cell = (*NSCell)(nil)

type NSCell struct {
	objc.NSObject
}

func MakeCell(ptr unsafe.Pointer) *NSCell {
	if ptr == nil {
		return nil
	}
	return &NSCell{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (c *NSCell) Wraps() bool {
	return bool(C.Cell_Wraps(c.Ptr()))
}

func (c *NSCell) SetWraps(wraps bool) {
	C.Cell_SetWraps(c.Ptr(), C.bool(wraps))
}

func (c *NSCell) IsScrollable() bool {
	return bool(C.Cell_IsScrollable(c.Ptr()))
}

func (c *NSCell) SetScrollable(scrollable bool) {
	C.Cell_SetScrollable(c.Ptr(), C.bool(scrollable))
}

func (c *NSCell) IsEditable() bool {
	return bool(C.Cell_IsEditable(c.Ptr()))
}

func (c *NSCell) SetEditable(editable bool) {
	C.Cell_SetEditable(c.Ptr(), C.bool(editable))
}

func (c *NSCell) IsSelectable() bool {
	return bool(C.Cell_IsSelectable(c.Ptr()))
}

func (c *NSCell) SetSelectable(selectable bool) {
	C.Cell_SetSelectable(c.Ptr(), C.bool(selectable))
}
