package appkit

// #include "stepper.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Stepper interface {
	Control
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	Autorepeat() bool
	SetAutorepeat(value bool)
	ValueWraps() bool
	SetValueWraps(value bool)
}

type NSStepper struct {
	NSControl
}

func MakeStepper(ptr unsafe.Pointer) *NSStepper {
	if ptr == nil {
		return nil
	}
	return &NSStepper{
		NSControl: *MakeControl(ptr),
	}
}

func AllocStepper() *NSStepper {
	return MakeStepper(C.C_Stepper_Alloc())
}

func (n *NSStepper) InitWithFrame(frameRect foundation.Rect) Stepper {
	result_ := C.C_NSStepper_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeStepper(result_)
}

func (n *NSStepper) InitWithCoder(coder foundation.Coder) Stepper {
	result_ := C.C_NSStepper_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeStepper(result_)
}

func (n *NSStepper) Init() Stepper {
	result_ := C.C_NSStepper_Init(n.Ptr())
	return MakeStepper(result_)
}

func (n *NSStepper) MaxValue() float64 {
	result_ := C.C_NSStepper_MaxValue(n.Ptr())
	return float64(result_)
}

func (n *NSStepper) SetMaxValue(value float64) {
	C.C_NSStepper_SetMaxValue(n.Ptr(), C.double(value))
}

func (n *NSStepper) MinValue() float64 {
	result_ := C.C_NSStepper_MinValue(n.Ptr())
	return float64(result_)
}

func (n *NSStepper) SetMinValue(value float64) {
	C.C_NSStepper_SetMinValue(n.Ptr(), C.double(value))
}

func (n *NSStepper) Increment() float64 {
	result_ := C.C_NSStepper_Increment(n.Ptr())
	return float64(result_)
}

func (n *NSStepper) SetIncrement(value float64) {
	C.C_NSStepper_SetIncrement(n.Ptr(), C.double(value))
}

func (n *NSStepper) Autorepeat() bool {
	result_ := C.C_NSStepper_Autorepeat(n.Ptr())
	return bool(result_)
}

func (n *NSStepper) SetAutorepeat(value bool) {
	C.C_NSStepper_SetAutorepeat(n.Ptr(), C.bool(value))
}

func (n *NSStepper) ValueWraps() bool {
	result_ := C.C_NSStepper_ValueWraps(n.Ptr())
	return bool(result_)
}

func (n *NSStepper) SetValueWraps(value bool) {
	C.C_NSStepper_SetValueWraps(n.Ptr(), C.bool(value))
}
