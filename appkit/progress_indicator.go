package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "progress_indicator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ProgressIndicator interface {
	View
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(usesThreadedAnimation bool)
	DoubleValue() float64
	SetDoubleValue(doubleValue float64)
	MinValue() float64
	SetMinValue(minValue float64)
	MaxValue() float64
	SetMaxValue(maxValue float64)
	IsIndeterminate() bool
	SetIndeterminate(indeterminate bool)
	IsBezeled() bool
	SetBezeled(bezeled bool)
	IsDisplayedWhenStopped() bool
	SetDisplayedWhenStopped(displayedWhenStopped bool)
	StartAnimation(sender objc.Object)
	StopAnimation(sender objc.Object)
	IncrementBy(delta float64)
}

var _ ProgressIndicator = (*NSProgressIndicator)(nil)

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

func (p *NSProgressIndicator) UsesThreadedAnimation() bool {
	return bool(C.ProgressIndicator_UsesThreadedAnimation(p.Ptr()))
}

func (p *NSProgressIndicator) SetUsesThreadedAnimation(usesThreadedAnimation bool) {
	C.ProgressIndicator_SetUsesThreadedAnimation(p.Ptr(), C.bool(usesThreadedAnimation))
}

func (p *NSProgressIndicator) DoubleValue() float64 {
	return float64(C.ProgressIndicator_DoubleValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetDoubleValue(doubleValue float64) {
	C.ProgressIndicator_SetDoubleValue(p.Ptr(), C.double(doubleValue))
}

func (p *NSProgressIndicator) MinValue() float64 {
	return float64(C.ProgressIndicator_MinValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetMinValue(minValue float64) {
	C.ProgressIndicator_SetMinValue(p.Ptr(), C.double(minValue))
}

func (p *NSProgressIndicator) MaxValue() float64 {
	return float64(C.ProgressIndicator_MaxValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetMaxValue(maxValue float64) {
	C.ProgressIndicator_SetMaxValue(p.Ptr(), C.double(maxValue))
}

func (p *NSProgressIndicator) IsIndeterminate() bool {
	return bool(C.ProgressIndicator_IsIndeterminate(p.Ptr()))
}

func (p *NSProgressIndicator) SetIndeterminate(indeterminate bool) {
	C.ProgressIndicator_SetIndeterminate(p.Ptr(), C.bool(indeterminate))
}

func (p *NSProgressIndicator) IsBezeled() bool {
	return bool(C.ProgressIndicator_IsBezeled(p.Ptr()))
}

func (p *NSProgressIndicator) SetBezeled(bezeled bool) {
	C.ProgressIndicator_SetBezeled(p.Ptr(), C.bool(bezeled))
}

func (p *NSProgressIndicator) IsDisplayedWhenStopped() bool {
	return bool(C.ProgressIndicator_IsDisplayedWhenStopped(p.Ptr()))
}

func (p *NSProgressIndicator) SetDisplayedWhenStopped(displayedWhenStopped bool) {
	C.ProgressIndicator_SetDisplayedWhenStopped(p.Ptr(), C.bool(displayedWhenStopped))
}

func NewProgressIndicator(frame foundation.Rect) ProgressIndicator {
	return MakeProgressIndicator(C.ProgressIndicator_NewProgressIndicator(toNSRect(frame)))
}

func (p *NSProgressIndicator) StartAnimation(sender objc.Object) {
	C.ProgressIndicator_StartAnimation(p.Ptr(), toPointer(sender))
}

func (p *NSProgressIndicator) StopAnimation(sender objc.Object) {
	C.ProgressIndicator_StopAnimation(p.Ptr(), toPointer(sender))
}

func (p *NSProgressIndicator) IncrementBy(delta float64) {
	C.ProgressIndicator_IncrementBy(p.Ptr(), C.double(delta))
}
