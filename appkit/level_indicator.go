package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "level_indicator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type LevelIndicator interface {
	Control
	MinValue() float64
	SetMinValue(minValue float64)
	MaxValue() float64
	SetMaxValue(maxValue float64)
	WarningValue() float64
	SetWarningValue(warningValue float64)
	CriticalValue() float64
	SetCriticalValue(criticalValue float64)
	IsEditable() bool
	SetEditable(editable bool)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(tickMarkPosition TickMarkPosition)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(numberOfTickMarks int)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(numberOfMajorTickMarks int)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(levelIndicatorStyle LevelIndicatorStyle)
	TickMarkValueAtIndex(index int) float64
	RectOfTickMarkAtIndex(index int) foundation.Rect
}

var _ LevelIndicator = (*NSLevelIndicator)(nil)

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

func (l *NSLevelIndicator) MinValue() float64 {
	return float64(C.LevelIndicator_MinValue(l.Ptr()))
}

func (l *NSLevelIndicator) SetMinValue(minValue float64) {
	C.LevelIndicator_SetMinValue(l.Ptr(), C.double(minValue))
}

func (l *NSLevelIndicator) MaxValue() float64 {
	return float64(C.LevelIndicator_MaxValue(l.Ptr()))
}

func (l *NSLevelIndicator) SetMaxValue(maxValue float64) {
	C.LevelIndicator_SetMaxValue(l.Ptr(), C.double(maxValue))
}

func (l *NSLevelIndicator) WarningValue() float64 {
	return float64(C.LevelIndicator_WarningValue(l.Ptr()))
}

func (l *NSLevelIndicator) SetWarningValue(warningValue float64) {
	C.LevelIndicator_SetWarningValue(l.Ptr(), C.double(warningValue))
}

func (l *NSLevelIndicator) CriticalValue() float64 {
	return float64(C.LevelIndicator_CriticalValue(l.Ptr()))
}

func (l *NSLevelIndicator) SetCriticalValue(criticalValue float64) {
	C.LevelIndicator_SetCriticalValue(l.Ptr(), C.double(criticalValue))
}

func (l *NSLevelIndicator) IsEditable() bool {
	return bool(C.LevelIndicator_IsEditable(l.Ptr()))
}

func (l *NSLevelIndicator) SetEditable(editable bool) {
	C.LevelIndicator_SetEditable(l.Ptr(), C.bool(editable))
}

func (l *NSLevelIndicator) TickMarkPosition() TickMarkPosition {
	return TickMarkPosition(C.LevelIndicator_TickMarkPosition(l.Ptr()))
}

func (l *NSLevelIndicator) SetTickMarkPosition(tickMarkPosition TickMarkPosition) {
	C.LevelIndicator_SetTickMarkPosition(l.Ptr(), C.ulong(tickMarkPosition))
}

func (l *NSLevelIndicator) NumberOfTickMarks() int {
	return int(C.LevelIndicator_NumberOfTickMarks(l.Ptr()))
}

func (l *NSLevelIndicator) SetNumberOfTickMarks(numberOfTickMarks int) {
	C.LevelIndicator_SetNumberOfTickMarks(l.Ptr(), C.long(numberOfTickMarks))
}

func (l *NSLevelIndicator) NumberOfMajorTickMarks() int {
	return int(C.LevelIndicator_NumberOfMajorTickMarks(l.Ptr()))
}

func (l *NSLevelIndicator) SetNumberOfMajorTickMarks(numberOfMajorTickMarks int) {
	C.LevelIndicator_SetNumberOfMajorTickMarks(l.Ptr(), C.long(numberOfMajorTickMarks))
}

func (l *NSLevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	return LevelIndicatorStyle(C.LevelIndicator_LevelIndicatorStyle(l.Ptr()))
}

func (l *NSLevelIndicator) SetLevelIndicatorStyle(levelIndicatorStyle LevelIndicatorStyle) {
	C.LevelIndicator_SetLevelIndicatorStyle(l.Ptr(), C.ulong(levelIndicatorStyle))
}

func NewLevelIndicator(frame foundation.Rect) LevelIndicator {
	return MakeLevelIndicator(C.LevelIndicator_NewLevelIndicator(toNSRect(frame)))
}

func (l *NSLevelIndicator) TickMarkValueAtIndex(index int) float64 {
	return float64(C.LevelIndicator_TickMarkValueAtIndex(l.Ptr(), C.long(index)))
}

func (l *NSLevelIndicator) RectOfTickMarkAtIndex(index int) foundation.Rect {
	return toRect(C.LevelIndicator_RectOfTickMarkAtIndex(l.Ptr(), C.long(index)))
}
