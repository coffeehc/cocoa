package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "control.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type Control interface {
	View
	IsEnabled() bool
	SetEnabled(enabled bool)
	DoubleValue() float64
	SetDoubleValue(doubleValue float64)
	FloatValue() float32
	SetFloatValue(floatValue float32)
	IntValue() int
	SetIntValue(intValue int)
	IntegerValue() int
	SetIntegerValue(integerValue int)
	StringValue() string
	SetStringValue(stringValue string)
	Cell() Cell
	SetCell(cell Cell)
	SizeToFit()
}

var _ Control = (*NSControl)(nil)

type NSControl struct {
	NSView
}

func MakeControl(ptr unsafe.Pointer) *NSControl {
	if ptr == nil {
		return nil
	}
	return &NSControl{
		NSView: *MakeView(ptr),
	}
}

func (c *NSControl) IsEnabled() bool {
	return bool(C.Control_IsEnabled(c.Ptr()))
}

func (c *NSControl) SetEnabled(enabled bool) {
	C.Control_SetEnabled(c.Ptr(), C.bool(enabled))
}

func (c *NSControl) DoubleValue() float64 {
	return float64(C.Control_DoubleValue(c.Ptr()))
}

func (c *NSControl) SetDoubleValue(doubleValue float64) {
	C.Control_SetDoubleValue(c.Ptr(), C.double(doubleValue))
}

func (c *NSControl) FloatValue() float32 {
	return float32(C.Control_FloatValue(c.Ptr()))
}

func (c *NSControl) SetFloatValue(floatValue float32) {
	C.Control_SetFloatValue(c.Ptr(), C.float(floatValue))
}

func (c *NSControl) IntValue() int {
	return int(C.Control_IntValue(c.Ptr()))
}

func (c *NSControl) SetIntValue(intValue int) {
	C.Control_SetIntValue(c.Ptr(), C.long(intValue))
}

func (c *NSControl) IntegerValue() int {
	return int(C.Control_IntegerValue(c.Ptr()))
}

func (c *NSControl) SetIntegerValue(integerValue int) {
	C.Control_SetIntegerValue(c.Ptr(), C.long(integerValue))
}

func (c *NSControl) StringValue() string {
	return C.GoString(C.Control_StringValue(c.Ptr()))
}

func (c *NSControl) SetStringValue(stringValue string) {
	cStringValue := C.CString(stringValue)
	defer C.free(unsafe.Pointer(cStringValue))
	C.Control_SetStringValue(c.Ptr(), cStringValue)
}

func (c *NSControl) Cell() Cell {
	return MakeCell(C.Control_Cell(c.Ptr()))
}

func (c *NSControl) SetCell(cell Cell) {
	C.Control_SetCell(c.Ptr(), toPointer(cell))
}

func NewControl(frame foundation.Rect) Control {
	return MakeControl(C.Control_NewControl(toNSRect(frame)))
}

func (c *NSControl) SizeToFit() {
	C.Control_SizeToFit(c.Ptr())
}
