package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "progress_indicator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ProgressIndicator interface {
	View
	StartAnimation(sender objc.Object)
	StopAnimation(sender objc.Object)
	IncrementBy(delta float64)
	SizeToFit()
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	IsBezeled() bool
	SetBezeled(value bool)
	IsIndeterminate() bool
	SetIndeterminate(value bool)
	Style() ProgressIndicatorStyle
	SetStyle(value ProgressIndicatorStyle)
	IsDisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
}

type NSProgressIndicator struct {
	NSView
}

func MakeProgressIndicator(ptr unsafe.Pointer) *NSProgressIndicator {
	if ptr == nil {
		return nil
	}
	return &NSProgressIndicator{
		NSView: *MakeView(ptr),
	}
}

func AllocProgressIndicator() *NSProgressIndicator {
	return MakeProgressIndicator(C.C_ProgressIndicator_Alloc())
}

func (n *NSProgressIndicator) InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	result := C.C_NSProgressIndicator_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeProgressIndicator(result)
}

func (n *NSProgressIndicator) InitWithCoder(coder foundation.Coder) ProgressIndicator {
	result := C.C_NSProgressIndicator_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeProgressIndicator(result)
}

func (n *NSProgressIndicator) Init() ProgressIndicator {
	result := C.C_NSProgressIndicator_Init(n.Ptr())
	return MakeProgressIndicator(result)
}

func (n *NSProgressIndicator) StartAnimation(sender objc.Object) {
	C.C_NSProgressIndicator_StartAnimation(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSProgressIndicator) StopAnimation(sender objc.Object) {
	C.C_NSProgressIndicator_StopAnimation(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSProgressIndicator) IncrementBy(delta float64) {
	C.C_NSProgressIndicator_IncrementBy(n.Ptr(), C.double(delta))
}

func (n *NSProgressIndicator) SizeToFit() {
	C.C_NSProgressIndicator_SizeToFit(n.Ptr())
}

func (n *NSProgressIndicator) UsesThreadedAnimation() bool {
	result := C.C_NSProgressIndicator_UsesThreadedAnimation(n.Ptr())
	return bool(result)
}

func (n *NSProgressIndicator) SetUsesThreadedAnimation(value bool) {
	C.C_NSProgressIndicator_SetUsesThreadedAnimation(n.Ptr(), C.bool(value))
}

func (n *NSProgressIndicator) DoubleValue() float64 {
	result := C.C_NSProgressIndicator_DoubleValue(n.Ptr())
	return float64(result)
}

func (n *NSProgressIndicator) SetDoubleValue(value float64) {
	C.C_NSProgressIndicator_SetDoubleValue(n.Ptr(), C.double(value))
}

func (n *NSProgressIndicator) MinValue() float64 {
	result := C.C_NSProgressIndicator_MinValue(n.Ptr())
	return float64(result)
}

func (n *NSProgressIndicator) SetMinValue(value float64) {
	C.C_NSProgressIndicator_SetMinValue(n.Ptr(), C.double(value))
}

func (n *NSProgressIndicator) MaxValue() float64 {
	result := C.C_NSProgressIndicator_MaxValue(n.Ptr())
	return float64(result)
}

func (n *NSProgressIndicator) SetMaxValue(value float64) {
	C.C_NSProgressIndicator_SetMaxValue(n.Ptr(), C.double(value))
}

func (n *NSProgressIndicator) ControlSize() ControlSize {
	result := C.C_NSProgressIndicator_ControlSize(n.Ptr())
	return ControlSize(uint(result))
}

func (n *NSProgressIndicator) SetControlSize(value ControlSize) {
	C.C_NSProgressIndicator_SetControlSize(n.Ptr(), C.uint(uint(value)))
}

func (n *NSProgressIndicator) ControlTint() ControlTint {
	result := C.C_NSProgressIndicator_ControlTint(n.Ptr())
	return ControlTint(uint(result))
}

func (n *NSProgressIndicator) SetControlTint(value ControlTint) {
	C.C_NSProgressIndicator_SetControlTint(n.Ptr(), C.uint(uint(value)))
}

func (n *NSProgressIndicator) IsBezeled() bool {
	result := C.C_NSProgressIndicator_IsBezeled(n.Ptr())
	return bool(result)
}

func (n *NSProgressIndicator) SetBezeled(value bool) {
	C.C_NSProgressIndicator_SetBezeled(n.Ptr(), C.bool(value))
}

func (n *NSProgressIndicator) IsIndeterminate() bool {
	result := C.C_NSProgressIndicator_IsIndeterminate(n.Ptr())
	return bool(result)
}

func (n *NSProgressIndicator) SetIndeterminate(value bool) {
	C.C_NSProgressIndicator_SetIndeterminate(n.Ptr(), C.bool(value))
}

func (n *NSProgressIndicator) Style() ProgressIndicatorStyle {
	result := C.C_NSProgressIndicator_Style(n.Ptr())
	return ProgressIndicatorStyle(uint(result))
}

func (n *NSProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	C.C_NSProgressIndicator_SetStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSProgressIndicator) IsDisplayedWhenStopped() bool {
	result := C.C_NSProgressIndicator_IsDisplayedWhenStopped(n.Ptr())
	return bool(result)
}

func (n *NSProgressIndicator) SetDisplayedWhenStopped(value bool) {
	C.C_NSProgressIndicator_SetDisplayedWhenStopped(n.Ptr(), C.bool(value))
}
