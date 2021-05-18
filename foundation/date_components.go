package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "date_components.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DateComponents interface {
	objc.Object
	IsValidDateInCalendar(calendar Calendar) bool
	ValueForComponent(unit CalendarUnit) int
	SetValue_ForComponent(value int, unit CalendarUnit)
	Calendar() Calendar
	SetCalendar(value Calendar)
	TimeZone() TimeZone
	SetTimeZone(value TimeZone)
	IsValidDate() bool
	Date() Date
	Era() int
	SetEra(value int)
	Year() int
	SetYear(value int)
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
	Quarter() int
	SetQuarter(value int)
	Month() int
	SetMonth(value int)
	IsLeapMonth() bool
	SetLeapMonth(value bool)
	Weekday() int
	SetWeekday(value int)
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	WeekOfYear() int
	SetWeekOfYear(value int)
	Day() int
	SetDay(value int)
	Hour() int
	SetHour(value int)
	Minute() int
	SetMinute(value int)
	Second() int
	SetSecond(value int)
	Nanosecond() int
	SetNanosecond(value int)
}

type NSDateComponents struct {
	objc.NSObject
}

func MakeDateComponents(ptr unsafe.Pointer) *NSDateComponents {
	if ptr == nil {
		return nil
	}
	return &NSDateComponents{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocDateComponents() *NSDateComponents {
	return MakeDateComponents(C.C_DateComponents_Alloc())
}

func (n *NSDateComponents) Init() DateComponents {
	result_ := C.C_NSDateComponents_Init(n.Ptr())
	return MakeDateComponents(result_)
}

func (n *NSDateComponents) IsValidDateInCalendar(calendar Calendar) bool {
	result_ := C.C_NSDateComponents_IsValidDateInCalendar(n.Ptr(), objc.ExtractPtr(calendar))
	return bool(result_)
}

func (n *NSDateComponents) ValueForComponent(unit CalendarUnit) int {
	result_ := C.C_NSDateComponents_ValueForComponent(n.Ptr(), C.uint(uint(unit)))
	return int(result_)
}

func (n *NSDateComponents) SetValue_ForComponent(value int, unit CalendarUnit) {
	C.C_NSDateComponents_SetValue_ForComponent(n.Ptr(), C.int(value), C.uint(uint(unit)))
}

func (n *NSDateComponents) Calendar() Calendar {
	result_ := C.C_NSDateComponents_Calendar(n.Ptr())
	return MakeCalendar(result_)
}

func (n *NSDateComponents) SetCalendar(value Calendar) {
	C.C_NSDateComponents_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDateComponents) TimeZone() TimeZone {
	result_ := C.C_NSDateComponents_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n *NSDateComponents) SetTimeZone(value TimeZone) {
	C.C_NSDateComponents_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDateComponents) IsValidDate() bool {
	result_ := C.C_NSDateComponents_IsValidDate(n.Ptr())
	return bool(result_)
}

func (n *NSDateComponents) Date() Date {
	result_ := C.C_NSDateComponents_Date(n.Ptr())
	return MakeDate(result_)
}

func (n *NSDateComponents) Era() int {
	result_ := C.C_NSDateComponents_Era(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetEra(value int) {
	C.C_NSDateComponents_SetEra(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Year() int {
	result_ := C.C_NSDateComponents_Year(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetYear(value int) {
	C.C_NSDateComponents_SetYear(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) YearForWeekOfYear() int {
	result_ := C.C_NSDateComponents_YearForWeekOfYear(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetYearForWeekOfYear(value int) {
	C.C_NSDateComponents_SetYearForWeekOfYear(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Quarter() int {
	result_ := C.C_NSDateComponents_Quarter(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetQuarter(value int) {
	C.C_NSDateComponents_SetQuarter(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Month() int {
	result_ := C.C_NSDateComponents_Month(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetMonth(value int) {
	C.C_NSDateComponents_SetMonth(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) IsLeapMonth() bool {
	result_ := C.C_NSDateComponents_IsLeapMonth(n.Ptr())
	return bool(result_)
}

func (n *NSDateComponents) SetLeapMonth(value bool) {
	C.C_NSDateComponents_SetLeapMonth(n.Ptr(), C.bool(value))
}

func (n *NSDateComponents) Weekday() int {
	result_ := C.C_NSDateComponents_Weekday(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetWeekday(value int) {
	C.C_NSDateComponents_SetWeekday(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) WeekdayOrdinal() int {
	result_ := C.C_NSDateComponents_WeekdayOrdinal(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetWeekdayOrdinal(value int) {
	C.C_NSDateComponents_SetWeekdayOrdinal(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) WeekOfMonth() int {
	result_ := C.C_NSDateComponents_WeekOfMonth(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetWeekOfMonth(value int) {
	C.C_NSDateComponents_SetWeekOfMonth(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) WeekOfYear() int {
	result_ := C.C_NSDateComponents_WeekOfYear(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetWeekOfYear(value int) {
	C.C_NSDateComponents_SetWeekOfYear(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Day() int {
	result_ := C.C_NSDateComponents_Day(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetDay(value int) {
	C.C_NSDateComponents_SetDay(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Hour() int {
	result_ := C.C_NSDateComponents_Hour(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetHour(value int) {
	C.C_NSDateComponents_SetHour(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Minute() int {
	result_ := C.C_NSDateComponents_Minute(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetMinute(value int) {
	C.C_NSDateComponents_SetMinute(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Second() int {
	result_ := C.C_NSDateComponents_Second(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetSecond(value int) {
	C.C_NSDateComponents_SetSecond(n.Ptr(), C.int(value))
}

func (n *NSDateComponents) Nanosecond() int {
	result_ := C.C_NSDateComponents_Nanosecond(n.Ptr())
	return int(result_)
}

func (n *NSDateComponents) SetNanosecond(value int) {
	C.C_NSDateComponents_SetNanosecond(n.Ptr(), C.int(value))
}
