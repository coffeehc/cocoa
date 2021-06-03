package foundation

// #include "date_components_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DateComponentsFormatter interface {
	Formatter
	StringFromDateComponents(components DateComponents) string
	StringFromDate_ToDate(startDate Date, endDate Date) string
	StringFromTimeInterval(ti TimeInterval) string
	AllowedUnits() CalendarUnit
	SetAllowedUnits(value CalendarUnit)
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	Calendar() Calendar
	SetCalendar(value Calendar)
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	UnitsStyle() DateComponentsFormatterUnitsStyle
	SetUnitsStyle(value DateComponentsFormatterUnitsStyle)
	ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	ReferenceDate() Date
	SetReferenceDate(value Date)
}

type NSDateComponentsFormatter struct {
	NSFormatter
}

func MakeDateComponentsFormatter(ptr unsafe.Pointer) NSDateComponentsFormatter {
	return NSDateComponentsFormatter{
		NSFormatter: MakeFormatter(ptr),
	}
}

func AllocDateComponentsFormatter() NSDateComponentsFormatter {
	return MakeDateComponentsFormatter(C.C_DateComponentsFormatter_Alloc())
}

func (n NSDateComponentsFormatter) Init() DateComponentsFormatter {
	result_ := C.C_NSDateComponentsFormatter_Init(n.Ptr())
	return MakeDateComponentsFormatter(result_)
}

func (n NSDateComponentsFormatter) StringFromDateComponents(components DateComponents) string {
	result_ := C.C_NSDateComponentsFormatter_StringFromDateComponents(n.Ptr(), objc.ExtractPtr(components))
	return MakeString(result_).String()
}

func (n NSDateComponentsFormatter) StringFromDate_ToDate(startDate Date, endDate Date) string {
	result_ := C.C_NSDateComponentsFormatter_StringFromDate_ToDate(n.Ptr(), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return MakeString(result_).String()
}

func (n NSDateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	result_ := C.C_NSDateComponentsFormatter_StringFromTimeInterval(n.Ptr(), C.double(float64(ti)))
	return MakeString(result_).String()
}

func DateComponentsFormatter_LocalizedStringFromDateComponents_UnitsStyle(components DateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	result_ := C.C_NSDateComponentsFormatter_DateComponentsFormatter_LocalizedStringFromDateComponents_UnitsStyle(objc.ExtractPtr(components), C.int(int(unitsStyle)))
	return MakeString(result_).String()
}

func (n NSDateComponentsFormatter) AllowedUnits() CalendarUnit {
	result_ := C.C_NSDateComponentsFormatter_AllowedUnits(n.Ptr())
	return CalendarUnit(uint(result_))
}

func (n NSDateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	C.C_NSDateComponentsFormatter_SetAllowedUnits(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateComponentsFormatter) AllowsFractionalUnits() bool {
	result_ := C.C_NSDateComponentsFormatter_AllowsFractionalUnits(n.Ptr())
	return bool(result_)
}

func (n NSDateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	C.C_NSDateComponentsFormatter_SetAllowsFractionalUnits(n.Ptr(), C.bool(value))
}

func (n NSDateComponentsFormatter) Calendar() Calendar {
	result_ := C.C_NSDateComponentsFormatter_Calendar(n.Ptr())
	return MakeCalendar(result_)
}

func (n NSDateComponentsFormatter) SetCalendar(value Calendar) {
	C.C_NSDateComponentsFormatter_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateComponentsFormatter) CollapsesLargestUnit() bool {
	result_ := C.C_NSDateComponentsFormatter_CollapsesLargestUnit(n.Ptr())
	return bool(result_)
}

func (n NSDateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	C.C_NSDateComponentsFormatter_SetCollapsesLargestUnit(n.Ptr(), C.bool(value))
}

func (n NSDateComponentsFormatter) IncludesApproximationPhrase() bool {
	result_ := C.C_NSDateComponentsFormatter_IncludesApproximationPhrase(n.Ptr())
	return bool(result_)
}

func (n NSDateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	C.C_NSDateComponentsFormatter_SetIncludesApproximationPhrase(n.Ptr(), C.bool(value))
}

func (n NSDateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	result_ := C.C_NSDateComponentsFormatter_IncludesTimeRemainingPhrase(n.Ptr())
	return bool(result_)
}

func (n NSDateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	C.C_NSDateComponentsFormatter_SetIncludesTimeRemainingPhrase(n.Ptr(), C.bool(value))
}

func (n NSDateComponentsFormatter) MaximumUnitCount() int {
	result_ := C.C_NSDateComponentsFormatter_MaximumUnitCount(n.Ptr())
	return int(result_)
}

func (n NSDateComponentsFormatter) SetMaximumUnitCount(value int) {
	C.C_NSDateComponentsFormatter_SetMaximumUnitCount(n.Ptr(), C.int(value))
}

func (n NSDateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	result_ := C.C_NSDateComponentsFormatter_UnitsStyle(n.Ptr())
	return DateComponentsFormatterUnitsStyle(int(result_))
}

func (n NSDateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	C.C_NSDateComponentsFormatter_SetUnitsStyle(n.Ptr(), C.int(int(value)))
}

func (n NSDateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	result_ := C.C_NSDateComponentsFormatter_ZeroFormattingBehavior(n.Ptr())
	return DateComponentsFormatterZeroFormattingBehavior(uint(result_))
}

func (n NSDateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	C.C_NSDateComponentsFormatter_SetZeroFormattingBehavior(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateComponentsFormatter) FormattingContext() FormattingContext {
	result_ := C.C_NSDateComponentsFormatter_FormattingContext(n.Ptr())
	return FormattingContext(int(result_))
}

func (n NSDateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	C.C_NSDateComponentsFormatter_SetFormattingContext(n.Ptr(), C.int(int(value)))
}

func (n NSDateComponentsFormatter) ReferenceDate() Date {
	result_ := C.C_NSDateComponentsFormatter_ReferenceDate(n.Ptr())
	return MakeDate(result_)
}

func (n NSDateComponentsFormatter) SetReferenceDate(value Date) {
	C.C_NSDateComponentsFormatter_SetReferenceDate(n.Ptr(), objc.ExtractPtr(value))
}
