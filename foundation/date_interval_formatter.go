package foundation

// #include "date_interval_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DateIntervalFormatter interface {
	Formatter
	StringFromDate_ToDate(fromDate Date, toDate Date) string
	StringFromDateInterval(dateInterval DateInterval) string
	DateStyle() DateIntervalFormatterStyle
	SetDateStyle(value DateIntervalFormatterStyle)
	TimeStyle() DateIntervalFormatterStyle
	SetTimeStyle(value DateIntervalFormatterStyle)
	DateTemplate() string
	SetDateTemplate(value string)
	Calendar() Calendar
	SetCalendar(value Calendar)
	Locale() Locale
	SetLocale(value Locale)
	TimeZone() TimeZone
	SetTimeZone(value TimeZone)
}

type NSDateIntervalFormatter struct {
	NSFormatter
}

func MakeDateIntervalFormatter(ptr unsafe.Pointer) NSDateIntervalFormatter {
	return NSDateIntervalFormatter{
		NSFormatter: MakeFormatter(ptr),
	}
}

func AllocDateIntervalFormatter() NSDateIntervalFormatter {
	return MakeDateIntervalFormatter(C.C_DateIntervalFormatter_Alloc())
}

func (n NSDateIntervalFormatter) Init() DateIntervalFormatter {
	result_ := C.C_NSDateIntervalFormatter_Init(n.Ptr())
	return MakeDateIntervalFormatter(result_)
}

func (n NSDateIntervalFormatter) StringFromDate_ToDate(fromDate Date, toDate Date) string {
	result_ := C.C_NSDateIntervalFormatter_StringFromDate_ToDate(n.Ptr(), objc.ExtractPtr(fromDate), objc.ExtractPtr(toDate))
	return MakeString(result_).String()
}

func (n NSDateIntervalFormatter) StringFromDateInterval(dateInterval DateInterval) string {
	result_ := C.C_NSDateIntervalFormatter_StringFromDateInterval(n.Ptr(), objc.ExtractPtr(dateInterval))
	return MakeString(result_).String()
}

func (n NSDateIntervalFormatter) DateStyle() DateIntervalFormatterStyle {
	result_ := C.C_NSDateIntervalFormatter_DateStyle(n.Ptr())
	return DateIntervalFormatterStyle(uint(result_))
}

func (n NSDateIntervalFormatter) SetDateStyle(value DateIntervalFormatterStyle) {
	C.C_NSDateIntervalFormatter_SetDateStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateIntervalFormatter) TimeStyle() DateIntervalFormatterStyle {
	result_ := C.C_NSDateIntervalFormatter_TimeStyle(n.Ptr())
	return DateIntervalFormatterStyle(uint(result_))
}

func (n NSDateIntervalFormatter) SetTimeStyle(value DateIntervalFormatterStyle) {
	C.C_NSDateIntervalFormatter_SetTimeStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateIntervalFormatter) DateTemplate() string {
	result_ := C.C_NSDateIntervalFormatter_DateTemplate(n.Ptr())
	return MakeString(result_).String()
}

func (n NSDateIntervalFormatter) SetDateTemplate(value string) {
	C.C_NSDateIntervalFormatter_SetDateTemplate(n.Ptr(), NewString(value).Ptr())
}

func (n NSDateIntervalFormatter) Calendar() Calendar {
	result_ := C.C_NSDateIntervalFormatter_Calendar(n.Ptr())
	return MakeCalendar(result_)
}

func (n NSDateIntervalFormatter) SetCalendar(value Calendar) {
	C.C_NSDateIntervalFormatter_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateIntervalFormatter) Locale() Locale {
	result_ := C.C_NSDateIntervalFormatter_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n NSDateIntervalFormatter) SetLocale(value Locale) {
	C.C_NSDateIntervalFormatter_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateIntervalFormatter) TimeZone() TimeZone {
	result_ := C.C_NSDateIntervalFormatter_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n NSDateIntervalFormatter) SetTimeZone(value TimeZone) {
	C.C_NSDateIntervalFormatter_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}
