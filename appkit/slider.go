package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "slider.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Slider interface {
	Control
	SliderType() SliderType
	SetSliderType(sliderType SliderType)
	AltIncrementValue() float64
	SetAltIncrementValue(altIncrementValue float64)
	KnobThickness() float64
	IsVertical() bool
	SetVertical(vertical bool)
	TrackFillColor() Color
	SetTrackFillColor(trackFillColor Color)
	MaxValue() float64
	SetMaxValue(maxValue float64)
	MinValue() float64
	SetMinValue(minValue float64)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(allowsTickMarkValuesOnly bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(numberOfTickMarks int)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(tickMarkPosition TickMarkPosition)
	AcceptsFirstMouse(event Event) bool
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point foundation.Point) int
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	SetAction(handler ActionHandler)
}

var _ Slider = (*NSSlider)(nil)

type NSSlider struct {
	NSControl
	action ActionHandler
}

func MakeSlider(ptr unsafe.Pointer) *NSSlider {
	if ptr == nil {
		return nil
	}
	return &NSSlider{
		NSControl: *MakeControl(ptr),
	}
}
func (s *NSSlider) setDelegate() {
	id := resources.Register(s)
	C.Slider_SetDelegate(s.Ptr(), C.long(id))
}

func (s *NSSlider) SliderType() SliderType {
	return SliderType(C.Slider_SliderType(s.Ptr()))
}

func (s *NSSlider) SetSliderType(sliderType SliderType) {
	C.Slider_SetSliderType(s.Ptr(), C.ulong(sliderType))
}

func (s *NSSlider) AltIncrementValue() float64 {
	return float64(C.Slider_AltIncrementValue(s.Ptr()))
}

func (s *NSSlider) SetAltIncrementValue(altIncrementValue float64) {
	C.Slider_SetAltIncrementValue(s.Ptr(), C.double(altIncrementValue))
}

func (s *NSSlider) KnobThickness() float64 {
	return float64(C.Slider_KnobThickness(s.Ptr()))
}

func (s *NSSlider) IsVertical() bool {
	return bool(C.Slider_IsVertical(s.Ptr()))
}

func (s *NSSlider) SetVertical(vertical bool) {
	C.Slider_SetVertical(s.Ptr(), C.bool(vertical))
}

func (s *NSSlider) TrackFillColor() Color {
	return MakeColor(C.Slider_TrackFillColor(s.Ptr()))
}

func (s *NSSlider) SetTrackFillColor(trackFillColor Color) {
	C.Slider_SetTrackFillColor(s.Ptr(), toPointer(trackFillColor))
}

func (s *NSSlider) MaxValue() float64 {
	return float64(C.Slider_MaxValue(s.Ptr()))
}

func (s *NSSlider) SetMaxValue(maxValue float64) {
	C.Slider_SetMaxValue(s.Ptr(), C.double(maxValue))
}

func (s *NSSlider) MinValue() float64 {
	return float64(C.Slider_MinValue(s.Ptr()))
}

func (s *NSSlider) SetMinValue(minValue float64) {
	C.Slider_SetMinValue(s.Ptr(), C.double(minValue))
}

func (s *NSSlider) AllowsTickMarkValuesOnly() bool {
	return bool(C.Slider_AllowsTickMarkValuesOnly(s.Ptr()))
}

func (s *NSSlider) SetAllowsTickMarkValuesOnly(allowsTickMarkValuesOnly bool) {
	C.Slider_SetAllowsTickMarkValuesOnly(s.Ptr(), C.bool(allowsTickMarkValuesOnly))
}

func (s *NSSlider) NumberOfTickMarks() int {
	return int(C.Slider_NumberOfTickMarks(s.Ptr()))
}

func (s *NSSlider) SetNumberOfTickMarks(numberOfTickMarks int) {
	C.Slider_SetNumberOfTickMarks(s.Ptr(), C.long(numberOfTickMarks))
}

func (s *NSSlider) TickMarkPosition() TickMarkPosition {
	return TickMarkPosition(C.Slider_TickMarkPosition(s.Ptr()))
}

func (s *NSSlider) SetTickMarkPosition(tickMarkPosition TickMarkPosition) {
	C.Slider_SetTickMarkPosition(s.Ptr(), C.ulong(tickMarkPosition))
}

func NewSlider(frame foundation.Rect) Slider {
	instance := MakeSlider(C.Slider_NewSlider(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func SliderWithTarget(target objc.Object, action *objc.Selector) Control {
	return MakeControl(C.Slider_SliderWithTarget(toPointer(target), toPointer(action)))
}

func SliderWithValue(value float64, minValue float64, maxValue float64, target objc.Object, action *objc.Selector) Control {
	return MakeControl(C.Slider_SliderWithValue(C.double(value), C.double(minValue), C.double(maxValue), toPointer(target), toPointer(action)))
}

func (s *NSSlider) AcceptsFirstMouse(event Event) bool {
	return bool(C.Slider_AcceptsFirstMouse(s.Ptr(), toPointer(event)))
}

func (s *NSSlider) ClosestTickMarkValueToValue(value float64) float64 {
	return float64(C.Slider_ClosestTickMarkValueToValue(s.Ptr(), C.double(value)))
}

func (s *NSSlider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	return int(C.Slider_IndexOfTickMarkAtPoint(s.Ptr(), toNSPoint(point)))
}

func (s *NSSlider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	return toRect(C.Slider_RectOfTickMarkAtIndex(s.Ptr(), C.long(index)))
}

func (s *NSSlider) TickMarkValueAtIndex(index int) float64 {
	return float64(C.Slider_TickMarkValueAtIndex(s.Ptr(), C.long(index)))
}

func (s *NSSlider) SetAction(handler ActionHandler) {
	s.action = handler
}

//export Slider_Target_Action
func Slider_Target_Action(id int64, sender unsafe.Pointer) {
	slider := resources.Get(id).(*NSSlider)
	if slider.action != nil {
		slider.action(objc.MakeObject(sender))
	}
}
