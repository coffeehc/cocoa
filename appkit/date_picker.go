package appkit

// #include "date_picker.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DatePicker interface {
	Control
	IsBezeled() bool
	SetBezeled(value bool)
	IsBordered() bool
	SetBordered(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value Color)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.Calendar)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
	DateValue() foundation.Date
	SetDateValue(value foundation.Date)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	MinDate() foundation.Date
	SetMinDate(value foundation.Date)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.Date)
}

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

func AllocDatePicker() *NSDatePicker {
	return MakeDatePicker(C.C_DatePicker_Alloc())
}

func (n *NSDatePicker) InitWithFrame(frameRect foundation.Rect) DatePicker {
	result_ := C.C_NSDatePicker_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeDatePicker(result_)
}

func (n *NSDatePicker) InitWithCoder(coder foundation.Coder) DatePicker {
	result_ := C.C_NSDatePicker_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeDatePicker(result_)
}

func (n *NSDatePicker) Init() DatePicker {
	result_ := C.C_NSDatePicker_Init(n.Ptr())
	return MakeDatePicker(result_)
}

func (n *NSDatePicker) IsBezeled() bool {
	result_ := C.C_NSDatePicker_IsBezeled(n.Ptr())
	return bool(result_)
}

func (n *NSDatePicker) SetBezeled(value bool) {
	C.C_NSDatePicker_SetBezeled(n.Ptr(), C.bool(value))
}

func (n *NSDatePicker) IsBordered() bool {
	result_ := C.C_NSDatePicker_IsBordered(n.Ptr())
	return bool(result_)
}

func (n *NSDatePicker) SetBordered(value bool) {
	C.C_NSDatePicker_SetBordered(n.Ptr(), C.bool(value))
}

func (n *NSDatePicker) BackgroundColor() Color {
	result_ := C.C_NSDatePicker_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSDatePicker) SetBackgroundColor(value Color) {
	C.C_NSDatePicker_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) DrawsBackground() bool {
	result_ := C.C_NSDatePicker_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n *NSDatePicker) SetDrawsBackground(value bool) {
	C.C_NSDatePicker_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n *NSDatePicker) TextColor() Color {
	result_ := C.C_NSDatePicker_TextColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSDatePicker) SetTextColor(value Color) {
	C.C_NSDatePicker_SetTextColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) DatePickerStyle() DatePickerStyle {
	result_ := C.C_NSDatePicker_DatePickerStyle(n.Ptr())
	return DatePickerStyle(uint(result_))
}

func (n *NSDatePicker) SetDatePickerStyle(value DatePickerStyle) {
	C.C_NSDatePicker_SetDatePickerStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSDatePicker) PresentsCalendarOverlay() bool {
	result_ := C.C_NSDatePicker_PresentsCalendarOverlay(n.Ptr())
	return bool(result_)
}

func (n *NSDatePicker) SetPresentsCalendarOverlay(value bool) {
	C.C_NSDatePicker_SetPresentsCalendarOverlay(n.Ptr(), C.bool(value))
}

func (n *NSDatePicker) Delegate() objc.Object {
	result_ := C.C_NSDatePicker_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSDatePicker) SetDelegate(value objc.Object) {
	C.C_NSDatePicker_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) DatePickerElements() DatePickerElementFlags {
	result_ := C.C_NSDatePicker_DatePickerElements(n.Ptr())
	return DatePickerElementFlags(uint(result_))
}

func (n *NSDatePicker) SetDatePickerElements(value DatePickerElementFlags) {
	C.C_NSDatePicker_SetDatePickerElements(n.Ptr(), C.uint(uint(value)))
}

func (n *NSDatePicker) Calendar() foundation.Calendar {
	result_ := C.C_NSDatePicker_Calendar(n.Ptr())
	return foundation.MakeCalendar(result_)
}

func (n *NSDatePicker) SetCalendar(value foundation.Calendar) {
	C.C_NSDatePicker_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) Locale() foundation.Locale {
	result_ := C.C_NSDatePicker_Locale(n.Ptr())
	return foundation.MakeLocale(result_)
}

func (n *NSDatePicker) SetLocale(value foundation.Locale) {
	C.C_NSDatePicker_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) DatePickerMode() DatePickerMode {
	result_ := C.C_NSDatePicker_DatePickerMode(n.Ptr())
	return DatePickerMode(uint(result_))
}

func (n *NSDatePicker) SetDatePickerMode(value DatePickerMode) {
	C.C_NSDatePicker_SetDatePickerMode(n.Ptr(), C.uint(uint(value)))
}

func (n *NSDatePicker) TimeZone() foundation.TimeZone {
	result_ := C.C_NSDatePicker_TimeZone(n.Ptr())
	return foundation.MakeTimeZone(result_)
}

func (n *NSDatePicker) SetTimeZone(value foundation.TimeZone) {
	C.C_NSDatePicker_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) DateValue() foundation.Date {
	result_ := C.C_NSDatePicker_DateValue(n.Ptr())
	return foundation.MakeDate(result_)
}

func (n *NSDatePicker) SetDateValue(value foundation.Date) {
	C.C_NSDatePicker_SetDateValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) TimeInterval() foundation.TimeInterval {
	result_ := C.C_NSDatePicker_TimeInterval(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n *NSDatePicker) SetTimeInterval(value foundation.TimeInterval) {
	C.C_NSDatePicker_SetTimeInterval(n.Ptr(), C.double(float64(value)))
}

func (n *NSDatePicker) MinDate() foundation.Date {
	result_ := C.C_NSDatePicker_MinDate(n.Ptr())
	return foundation.MakeDate(result_)
}

func (n *NSDatePicker) SetMinDate(value foundation.Date) {
	C.C_NSDatePicker_SetMinDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDatePicker) MaxDate() foundation.Date {
	result_ := C.C_NSDatePicker_MaxDate(n.Ptr())
	return foundation.MakeDate(result_)
}

func (n *NSDatePicker) SetMaxDate(value foundation.Date) {
	C.C_NSDatePicker_SetMaxDate(n.Ptr(), objc.ExtractPtr(value))
}
