package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "level_indicator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LevelIndicator interface {
	Control
	TickMarkValueAtIndex(index int) float64
	RectOfTickMarkAtIndex(index int) foundation.Rect
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	WarningValue() float64
	SetWarningValue(value float64)
	CriticalValue() float64
	SetCriticalValue(value float64)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	RatingImage() Image
	SetRatingImage(value Image)
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	FillColor() Color
	SetFillColor(value Color)
	WarningFillColor() Color
	SetWarningFillColor(value Color)
	CriticalFillColor() Color
	SetCriticalFillColor(value Color)
	RatingPlaceholderImage() Image
	SetRatingPlaceholderImage(value Image)
	PlaceholderVisibility() LevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility)
	IsEditable() bool
	SetEditable(value bool)
}

type NSLevelIndicator struct {
	NSControl
}

func MakeLevelIndicator(ptr unsafe.Pointer) *NSLevelIndicator {
	if ptr == nil {
		return nil
	}
	return &NSLevelIndicator{
		NSControl: *MakeControl(ptr),
	}
}

func AllocLevelIndicator() *NSLevelIndicator {
	return MakeLevelIndicator(C.C_LevelIndicator_Alloc())
}

func (n *NSLevelIndicator) InitWithFrame(frameRect foundation.Rect) LevelIndicator {
	result := C.C_NSLevelIndicator_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeLevelIndicator(result)
}

func (n *NSLevelIndicator) InitWithCoder(coder foundation.Coder) LevelIndicator {
	result := C.C_NSLevelIndicator_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeLevelIndicator(result)
}

func (n *NSLevelIndicator) Init() LevelIndicator {
	result := C.C_NSLevelIndicator_Init(n.Ptr())
	return MakeLevelIndicator(result)
}

func (n *NSLevelIndicator) TickMarkValueAtIndex(index int) float64 {
	result := C.C_NSLevelIndicator_TickMarkValueAtIndex(n.Ptr(), C.int(index))
	return float64(result)
}

func (n *NSLevelIndicator) RectOfTickMarkAtIndex(index int) foundation.Rect {
	result := C.C_NSLevelIndicator_RectOfTickMarkAtIndex(n.Ptr(), C.int(index))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSLevelIndicator) MinValue() float64 {
	result := C.C_NSLevelIndicator_MinValue(n.Ptr())
	return float64(result)
}

func (n *NSLevelIndicator) SetMinValue(value float64) {
	C.C_NSLevelIndicator_SetMinValue(n.Ptr(), C.double(value))
}

func (n *NSLevelIndicator) MaxValue() float64 {
	result := C.C_NSLevelIndicator_MaxValue(n.Ptr())
	return float64(result)
}

func (n *NSLevelIndicator) SetMaxValue(value float64) {
	C.C_NSLevelIndicator_SetMaxValue(n.Ptr(), C.double(value))
}

func (n *NSLevelIndicator) WarningValue() float64 {
	result := C.C_NSLevelIndicator_WarningValue(n.Ptr())
	return float64(result)
}

func (n *NSLevelIndicator) SetWarningValue(value float64) {
	C.C_NSLevelIndicator_SetWarningValue(n.Ptr(), C.double(value))
}

func (n *NSLevelIndicator) CriticalValue() float64 {
	result := C.C_NSLevelIndicator_CriticalValue(n.Ptr())
	return float64(result)
}

func (n *NSLevelIndicator) SetCriticalValue(value float64) {
	C.C_NSLevelIndicator_SetCriticalValue(n.Ptr(), C.double(value))
}

func (n *NSLevelIndicator) TickMarkPosition() TickMarkPosition {
	result := C.C_NSLevelIndicator_TickMarkPosition(n.Ptr())
	return TickMarkPosition(uint(result))
}

func (n *NSLevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	C.C_NSLevelIndicator_SetTickMarkPosition(n.Ptr(), C.uint(uint(value)))
}

func (n *NSLevelIndicator) NumberOfTickMarks() int {
	result := C.C_NSLevelIndicator_NumberOfTickMarks(n.Ptr())
	return int(result)
}

func (n *NSLevelIndicator) SetNumberOfTickMarks(value int) {
	C.C_NSLevelIndicator_SetNumberOfTickMarks(n.Ptr(), C.int(value))
}

func (n *NSLevelIndicator) NumberOfMajorTickMarks() int {
	result := C.C_NSLevelIndicator_NumberOfMajorTickMarks(n.Ptr())
	return int(result)
}

func (n *NSLevelIndicator) SetNumberOfMajorTickMarks(value int) {
	C.C_NSLevelIndicator_SetNumberOfMajorTickMarks(n.Ptr(), C.int(value))
}

func (n *NSLevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	result := C.C_NSLevelIndicator_LevelIndicatorStyle(n.Ptr())
	return LevelIndicatorStyle(uint(result))
}

func (n *NSLevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	C.C_NSLevelIndicator_SetLevelIndicatorStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSLevelIndicator) RatingImage() Image {
	result := C.C_NSLevelIndicator_RatingImage(n.Ptr())
	return MakeImage(result)
}

func (n *NSLevelIndicator) SetRatingImage(value Image) {
	C.C_NSLevelIndicator_SetRatingImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLevelIndicator) DrawsTieredCapacityLevels() bool {
	result := C.C_NSLevelIndicator_DrawsTieredCapacityLevels(n.Ptr())
	return bool(result)
}

func (n *NSLevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	C.C_NSLevelIndicator_SetDrawsTieredCapacityLevels(n.Ptr(), C.bool(value))
}

func (n *NSLevelIndicator) FillColor() Color {
	result := C.C_NSLevelIndicator_FillColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSLevelIndicator) SetFillColor(value Color) {
	C.C_NSLevelIndicator_SetFillColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLevelIndicator) WarningFillColor() Color {
	result := C.C_NSLevelIndicator_WarningFillColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSLevelIndicator) SetWarningFillColor(value Color) {
	C.C_NSLevelIndicator_SetWarningFillColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLevelIndicator) CriticalFillColor() Color {
	result := C.C_NSLevelIndicator_CriticalFillColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSLevelIndicator) SetCriticalFillColor(value Color) {
	C.C_NSLevelIndicator_SetCriticalFillColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLevelIndicator) RatingPlaceholderImage() Image {
	result := C.C_NSLevelIndicator_RatingPlaceholderImage(n.Ptr())
	return MakeImage(result)
}

func (n *NSLevelIndicator) SetRatingPlaceholderImage(value Image) {
	C.C_NSLevelIndicator_SetRatingPlaceholderImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	result := C.C_NSLevelIndicator_PlaceholderVisibility(n.Ptr())
	return LevelIndicatorPlaceholderVisibility(int(result))
}

func (n *NSLevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	C.C_NSLevelIndicator_SetPlaceholderVisibility(n.Ptr(), C.int(int(value)))
}

func (n *NSLevelIndicator) IsEditable() bool {
	result := C.C_NSLevelIndicator_IsEditable(n.Ptr())
	return bool(result)
}

func (n *NSLevelIndicator) SetEditable(value bool) {
	C.C_NSLevelIndicator_SetEditable(n.Ptr(), C.bool(value))
}
