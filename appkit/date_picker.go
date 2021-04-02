package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "date_picker.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type DatePicker interface {
	Control
	IsBezeled() bool
	SetBezeled(bezeled bool)
	IsBordered() bool
	SetBordered(bordered bool)
	BackgroundColor() Color
	SetBackgroundColor(backgroundColor Color)
	DrawsBackground() bool
	SetDrawsBackground(drawsBackground bool)
	TextColor() Color
	SetTextColor(textColor Color)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(datePickerStyle DatePickerStyle)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(presentsCalendarOverlay bool)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(datePickerElements DatePickerElementFlags)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(datePickerMode DatePickerMode)
	DateValue() foundation.Date
	SetDateValue(dateValue foundation.Date)
	MaxDate() foundation.Date
	SetMaxDate(maxDate foundation.Date)
}

var _ DatePicker = (*NSDatePicker)(nil)

type NSDatePicker struct {
	NSControl
}

func MakeDatePicker(ptr unsafe.Pointer) *NSDatePicker {
	if ptr == nil {
		return nil
	}
	return &NSDatePicker{
		NSControl: *MakeControl(ptr),
	}
}

func (d *NSDatePicker) IsBezeled() bool {
	return bool(C.DatePicker_IsBezeled(d.Ptr()))
}

func (d *NSDatePicker) SetBezeled(bezeled bool) {
	C.DatePicker_SetBezeled(d.Ptr(), C.bool(bezeled))
}

func (d *NSDatePicker) IsBordered() bool {
	return bool(C.DatePicker_IsBordered(d.Ptr()))
}

func (d *NSDatePicker) SetBordered(bordered bool) {
	C.DatePicker_SetBordered(d.Ptr(), C.bool(bordered))
}

func (d *NSDatePicker) BackgroundColor() Color {
	return MakeColor(C.DatePicker_BackgroundColor(d.Ptr()))
}

func (d *NSDatePicker) SetBackgroundColor(backgroundColor Color) {
	C.DatePicker_SetBackgroundColor(d.Ptr(), toPointer(backgroundColor))
}

func (d *NSDatePicker) DrawsBackground() bool {
	return bool(C.DatePicker_DrawsBackground(d.Ptr()))
}

func (d *NSDatePicker) SetDrawsBackground(drawsBackground bool) {
	C.DatePicker_SetDrawsBackground(d.Ptr(), C.bool(drawsBackground))
}

func (d *NSDatePicker) TextColor() Color {
	return MakeColor(C.DatePicker_TextColor(d.Ptr()))
}

func (d *NSDatePicker) SetTextColor(textColor Color) {
	C.DatePicker_SetTextColor(d.Ptr(), toPointer(textColor))
}

func (d *NSDatePicker) DatePickerStyle() DatePickerStyle {
	return DatePickerStyle(C.DatePicker_DatePickerStyle(d.Ptr()))
}

func (d *NSDatePicker) SetDatePickerStyle(datePickerStyle DatePickerStyle) {
	C.DatePicker_SetDatePickerStyle(d.Ptr(), C.ulong(datePickerStyle))
}

func (d *NSDatePicker) PresentsCalendarOverlay() bool {
	return bool(C.DatePicker_PresentsCalendarOverlay(d.Ptr()))
}

func (d *NSDatePicker) SetPresentsCalendarOverlay(presentsCalendarOverlay bool) {
	C.DatePicker_SetPresentsCalendarOverlay(d.Ptr(), C.bool(presentsCalendarOverlay))
}

func (d *NSDatePicker) DatePickerElements() DatePickerElementFlags {
	return DatePickerElementFlags(C.DatePicker_DatePickerElements(d.Ptr()))
}

func (d *NSDatePicker) SetDatePickerElements(datePickerElements DatePickerElementFlags) {
	C.DatePicker_SetDatePickerElements(d.Ptr(), C.ulong(datePickerElements))
}

func (d *NSDatePicker) DatePickerMode() DatePickerMode {
	return DatePickerMode(C.DatePicker_DatePickerMode(d.Ptr()))
}

func (d *NSDatePicker) SetDatePickerMode(datePickerMode DatePickerMode) {
	C.DatePicker_SetDatePickerMode(d.Ptr(), C.ulong(datePickerMode))
}

func (d *NSDatePicker) DateValue() foundation.Date {
	return foundation.MakeDate(C.DatePicker_DateValue(d.Ptr()))
}

func (d *NSDatePicker) SetDateValue(dateValue foundation.Date) {
	C.DatePicker_SetDateValue(d.Ptr(), toPointer(dateValue))
}

func (d *NSDatePicker) MaxDate() foundation.Date {
	return foundation.MakeDate(C.DatePicker_MaxDate(d.Ptr()))
}

func (d *NSDatePicker) SetMaxDate(maxDate foundation.Date) {
	C.DatePicker_SetMaxDate(d.Ptr(), toPointer(maxDate))
}

func NewDatePicker(frame foundation.Rect) DatePicker {
	return MakeDatePicker(C.DatePicker_NewDatePicker(toNSRect(frame)))
}
