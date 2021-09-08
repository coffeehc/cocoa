package appkit

// #include "slider.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Slider interface {
	Control
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point foundation.Point) int
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	SliderType() SliderType
	SetSliderType(value SliderType)
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	KnobThickness() coregraphics.Float
	IsVertical() bool
	SetVertical(value bool)
	TrackFillColor() Color
	SetTrackFillColor(value Color)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
}

type NSSlider struct {
	NSControl
}

func MakeSlider(ptr unsafe.Pointer) NSSlider {
	return NSSlider{
		NSControl: MakeControl(ptr),
	}
}

func SliderWithTarget_Action(target objc.Object, action objc.Selector) NSSlider {
	result_ := C.C_NSSlider_SliderWithTarget_Action(objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeSlider(result_)
}

func SliderWithValue_MinValue_MaxValue_Target_Action(value float64, minValue float64, maxValue float64, target objc.Object, action objc.Selector) NSSlider {
	result_ := C.C_NSSlider_SliderWithValue_MinValue_MaxValue_Target_Action(C.double(value), C.double(minValue), C.double(maxValue), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeSlider(result_)
}

func (n NSSlider) InitWithFrame(frameRect foundation.Rect) NSSlider {
	result_ := C.C_NSSlider_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSlider(result_)
}

func (n NSSlider) InitWithCoder(coder foundation.Coder) NSSlider {
	result_ := C.C_NSSlider_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSlider(result_)
}

func (n NSSlider) Init() NSSlider {
	result_ := C.C_NSSlider_Init(n.Ptr())
	return MakeSlider(result_)
}

func AllocSlider() NSSlider {
	result_ := C.C_NSSlider_AllocSlider()
	return MakeSlider(result_)
}

func NewSlider() NSSlider {
	result_ := C.C_NSSlider_NewSlider()
	return MakeSlider(result_)
}

func (n NSSlider) Autorelease() NSSlider {
	result_ := C.C_NSSlider_Autorelease(n.Ptr())
	return MakeSlider(result_)
}

func (n NSSlider) Retain() NSSlider {
	result_ := C.C_NSSlider_Retain(n.Ptr())
	return MakeSlider(result_)
}

func (n NSSlider) ClosestTickMarkValueToValue(value float64) float64 {
	result_ := C.C_NSSlider_ClosestTickMarkValueToValue(n.Ptr(), C.double(value))
	return float64(result_)
}

func (n NSSlider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	result_ := C.C_NSSlider_IndexOfTickMarkAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return int(result_)
}

func (n NSSlider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	result_ := C.C_NSSlider_RectOfTickMarkAtIndex(n.Ptr(), C.int(index))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSlider) TickMarkValueAtIndex(index int) float64 {
	result_ := C.C_NSSlider_TickMarkValueAtIndex(n.Ptr(), C.int(index))
	return float64(result_)
}

func (n NSSlider) SliderType() SliderType {
	result_ := C.C_NSSlider_SliderType(n.Ptr())
	return SliderType(uint(result_))
}

func (n NSSlider) SetSliderType(value SliderType) {
	C.C_NSSlider_SetSliderType(n.Ptr(), C.uint(uint(value)))
}

func (n NSSlider) AltIncrementValue() float64 {
	result_ := C.C_NSSlider_AltIncrementValue(n.Ptr())
	return float64(result_)
}

func (n NSSlider) SetAltIncrementValue(value float64) {
	C.C_NSSlider_SetAltIncrementValue(n.Ptr(), C.double(value))
}

func (n NSSlider) KnobThickness() coregraphics.Float {
	result_ := C.C_NSSlider_KnobThickness(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSSlider) IsVertical() bool {
	result_ := C.C_NSSlider_IsVertical(n.Ptr())
	return bool(result_)
}

func (n NSSlider) SetVertical(value bool) {
	C.C_NSSlider_SetVertical(n.Ptr(), C.bool(value))
}

func (n NSSlider) TrackFillColor() Color {
	result_ := C.C_NSSlider_TrackFillColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSSlider) SetTrackFillColor(value Color) {
	C.C_NSSlider_SetTrackFillColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSlider) MaxValue() float64 {
	result_ := C.C_NSSlider_MaxValue(n.Ptr())
	return float64(result_)
}

func (n NSSlider) SetMaxValue(value float64) {
	C.C_NSSlider_SetMaxValue(n.Ptr(), C.double(value))
}

func (n NSSlider) MinValue() float64 {
	result_ := C.C_NSSlider_MinValue(n.Ptr())
	return float64(result_)
}

func (n NSSlider) SetMinValue(value float64) {
	C.C_NSSlider_SetMinValue(n.Ptr(), C.double(value))
}

func (n NSSlider) AllowsTickMarkValuesOnly() bool {
	result_ := C.C_NSSlider_AllowsTickMarkValuesOnly(n.Ptr())
	return bool(result_)
}

func (n NSSlider) SetAllowsTickMarkValuesOnly(value bool) {
	C.C_NSSlider_SetAllowsTickMarkValuesOnly(n.Ptr(), C.bool(value))
}

func (n NSSlider) NumberOfTickMarks() int {
	result_ := C.C_NSSlider_NumberOfTickMarks(n.Ptr())
	return int(result_)
}

func (n NSSlider) SetNumberOfTickMarks(value int) {
	C.C_NSSlider_SetNumberOfTickMarks(n.Ptr(), C.int(value))
}

func (n NSSlider) TickMarkPosition() TickMarkPosition {
	result_ := C.C_NSSlider_TickMarkPosition(n.Ptr())
	return TickMarkPosition(uint(result_))
}

func (n NSSlider) SetTickMarkPosition(value TickMarkPosition) {
	C.C_NSSlider_SetTickMarkPosition(n.Ptr(), C.uint(uint(value)))
}
