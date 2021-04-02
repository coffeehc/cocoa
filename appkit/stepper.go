package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "stepper.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type Stepper interface {
	Control
	MaxValue() float64
	SetMaxValue(maxValue float64)
	MinValue() float64
	SetMinValue(minValue float64)
	Increment() float64
	SetIncrement(increment float64)
	Autorepeat() bool
	SetAutorepeat(autorepeat bool)
	ValueWraps() bool
	SetValueWraps(valueWraps bool)
}

var _ Stepper = (*NSStepper)(nil)

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

func (s *NSStepper) MaxValue() float64 {
	return float64(C.Stepper_MaxValue(s.Ptr()))
}

func (s *NSStepper) SetMaxValue(maxValue float64) {
	C.Stepper_SetMaxValue(s.Ptr(), C.double(maxValue))
}

func (s *NSStepper) MinValue() float64 {
	return float64(C.Stepper_MinValue(s.Ptr()))
}

func (s *NSStepper) SetMinValue(minValue float64) {
	C.Stepper_SetMinValue(s.Ptr(), C.double(minValue))
}

func (s *NSStepper) Increment() float64 {
	return float64(C.Stepper_Increment(s.Ptr()))
}

func (s *NSStepper) SetIncrement(increment float64) {
	C.Stepper_SetIncrement(s.Ptr(), C.double(increment))
}

func (s *NSStepper) Autorepeat() bool {
	return bool(C.Stepper_Autorepeat(s.Ptr()))
}

func (s *NSStepper) SetAutorepeat(autorepeat bool) {
	C.Stepper_SetAutorepeat(s.Ptr(), C.bool(autorepeat))
}

func (s *NSStepper) ValueWraps() bool {
	return bool(C.Stepper_ValueWraps(s.Ptr()))
}

func (s *NSStepper) SetValueWraps(valueWraps bool) {
	C.Stepper_SetValueWraps(s.Ptr(), C.bool(valueWraps))
}

func NewStepper(frame foundation.Rect) Stepper {
	return MakeStepper(C.Stepper_NewStepper(toNSRect(frame)))
}
