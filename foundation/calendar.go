package foundation

// #include "calendar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Calendar interface {
	objc.Object
	InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object
	Date_MatchesComponents(date Date, components DateComponents) bool
	Component_FromDate(unit CalendarUnit, date Date) int
	Components_FromDate(unitFlags CalendarUnit, date Date) DateComponents
	Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate Date, resultDate Date, opts CalendarOptions) DateComponents
	Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp DateComponents, resultDateComp DateComponents, options CalendarOptions) DateComponents
	ComponentsInTimeZone_FromDate(timezone TimeZone, date Date) DateComponents
	MaximumRangeOfUnit(unit CalendarUnit) Range
	MinimumRangeOfUnit(unit CalendarUnit) Range
	OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) uint
	RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) Range
	StartOfDayForDate(date Date) Date
	NextDateAfterDate_MatchingComponents_Options(date Date, comps DateComponents, options CalendarOptions) Date
	NextDateAfterDate_MatchingHour_Minute_Second_Options(date Date, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date
	NextDateAfterDate_MatchingUnit_Value_Options(date Date, unit CalendarUnit, value int, options CalendarOptions) Date
	DateFromComponents(comps DateComponents) Date
	DateByAddingComponents_ToDate_Options(comps DateComponents, date Date, opts CalendarOptions) Date
	DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date Date, options CalendarOptions) Date
	DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date Date, opts CalendarOptions) Date
	DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date Date, opts CalendarOptions) Date
	DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	CompareDate_ToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) ComparisonResult
	IsDate_EqualToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) bool
	IsDate_InSameDayAsDate(date1 Date, date2 Date) bool
	IsDateInToday(date Date) bool
	IsDateInTomorrow(date Date) bool
	IsDateInWeekend(date Date) bool
	IsDateInYesterday(date Date) bool
	CalendarIdentifier() CalendarIdentifier
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() Locale
	SetLocale(value Locale)
	TimeZone() TimeZone
	SetTimeZone(value TimeZone)
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	AMSymbol() string
	PMSymbol() string
	WeekdaySymbols() []string
	ShortWeekdaySymbols() []string
	VeryShortWeekdaySymbols() []string
	StandaloneWeekdaySymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	VeryShortStandaloneWeekdaySymbols() []string
	MonthSymbols() []string
	ShortMonthSymbols() []string
	VeryShortMonthSymbols() []string
	StandaloneMonthSymbols() []string
	ShortStandaloneMonthSymbols() []string
	VeryShortStandaloneMonthSymbols() []string
	QuarterSymbols() []string
	ShortQuarterSymbols() []string
	StandaloneQuarterSymbols() []string
	ShortStandaloneQuarterSymbols() []string
	EraSymbols() []string
	LongEraSymbols() []string
}

type NSCalendar struct {
	objc.NSObject
}

func MakeCalendar(ptr unsafe.Pointer) *NSCalendar {
	if ptr == nil {
		return nil
	}
	return &NSCalendar{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCalendar() *NSCalendar {
	return MakeCalendar(C.C_Calendar_Alloc())
}

func CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	result_ := C.C_NSCalendar_CalendarWithIdentifier(NewString(string(calendarIdentifierConstant)).Ptr())
	return MakeCalendar(result_)
}

func (n *NSCalendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	result_ := C.C_NSCalendar_InitWithCalendarIdentifier(n.Ptr(), NewString(string(ident)).Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCalendar) Date_MatchesComponents(date Date, components DateComponents) bool {
	result_ := C.C_NSCalendar_Date_MatchesComponents(n.Ptr(), objc.ExtractPtr(date), objc.ExtractPtr(components))
	return bool(result_)
}

func (n *NSCalendar) Component_FromDate(unit CalendarUnit, date Date) int {
	result_ := C.C_NSCalendar_Component_FromDate(n.Ptr(), C.uint(uint(unit)), objc.ExtractPtr(date))
	return int(result_)
}

func (n *NSCalendar) Components_FromDate(unitFlags CalendarUnit, date Date) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDate(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(date))
	return MakeDateComponents(result_)
}

func (n *NSCalendar) Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate Date, resultDate Date, opts CalendarOptions) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDate_ToDate_Options(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(startingDate), objc.ExtractPtr(resultDate), C.uint(uint(opts)))
	return MakeDateComponents(result_)
}

func (n *NSCalendar) Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp DateComponents, resultDateComp DateComponents, options CalendarOptions) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDateComponents_ToDateComponents_Options(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(startingDateComp), objc.ExtractPtr(resultDateComp), C.uint(uint(options)))
	return MakeDateComponents(result_)
}

func (n *NSCalendar) ComponentsInTimeZone_FromDate(timezone TimeZone, date Date) DateComponents {
	result_ := C.C_NSCalendar_ComponentsInTimeZone_FromDate(n.Ptr(), objc.ExtractPtr(timezone), objc.ExtractPtr(date))
	return MakeDateComponents(result_)
}

func (n *NSCalendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	result_ := C.C_NSCalendar_MaximumRangeOfUnit(n.Ptr(), C.uint(uint(unit)))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSCalendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	result_ := C.C_NSCalendar_MinimumRangeOfUnit(n.Ptr(), C.uint(uint(unit)))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSCalendar) OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) uint {
	result_ := C.C_NSCalendar_OrdinalityOfUnit_InUnit_ForDate(n.Ptr(), C.uint(uint(smaller)), C.uint(uint(larger)), objc.ExtractPtr(date))
	return uint(result_)
}

func (n *NSCalendar) RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) Range {
	result_ := C.C_NSCalendar_RangeOfUnit_InUnit_ForDate(n.Ptr(), C.uint(uint(smaller)), C.uint(uint(larger)), objc.ExtractPtr(date))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSCalendar) StartOfDayForDate(date Date) Date {
	result_ := C.C_NSCalendar_StartOfDayForDate(n.Ptr(), objc.ExtractPtr(date))
	return MakeDate(result_)
}

func (n *NSCalendar) NextDateAfterDate_MatchingComponents_Options(date Date, comps DateComponents, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingComponents_Options(n.Ptr(), objc.ExtractPtr(date), objc.ExtractPtr(comps), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n *NSCalendar) NextDateAfterDate_MatchingHour_Minute_Second_Options(date Date, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingHour_Minute_Second_Options(n.Ptr(), objc.ExtractPtr(date), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n *NSCalendar) NextDateAfterDate_MatchingUnit_Value_Options(date Date, unit CalendarUnit, value int, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingUnit_Value_Options(n.Ptr(), objc.ExtractPtr(date), C.uint(uint(unit)), C.int(value), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n *NSCalendar) DateFromComponents(comps DateComponents) Date {
	result_ := C.C_NSCalendar_DateFromComponents(n.Ptr(), objc.ExtractPtr(comps))
	return MakeDate(result_)
}

func (n *NSCalendar) DateByAddingComponents_ToDate_Options(comps DateComponents, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateByAddingComponents_ToDate_Options(n.Ptr(), objc.ExtractPtr(comps), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n *NSCalendar) DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date Date, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateByAddingUnit_Value_ToDate_Options(n.Ptr(), C.uint(uint(unit)), C.int(value), objc.ExtractPtr(date), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n *NSCalendar) DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateBySettingHour_Minute_Second_OfDate_Options(n.Ptr(), C.int(h), C.int(m), C.int(s), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n *NSCalendar) DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateBySettingUnit_Value_OfDate_Options(n.Ptr(), C.uint(uint(unit)), C.int(v), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n *NSCalendar) DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	result_ := C.C_NSCalendar_DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(n.Ptr(), C.int(eraValue), C.int(yearValue), C.int(monthValue), C.int(dayValue), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.int(nanosecondValue))
	return MakeDate(result_)
}

func (n *NSCalendar) DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	result_ := C.C_NSCalendar_DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(n.Ptr(), C.int(eraValue), C.int(yearValue), C.int(weekValue), C.int(weekdayValue), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.int(nanosecondValue))
	return MakeDate(result_)
}

func (n *NSCalendar) CompareDate_ToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) ComparisonResult {
	result_ := C.C_NSCalendar_CompareDate_ToDate_ToUnitGranularity(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2), C.uint(uint(unit)))
	return ComparisonResult(int(result_))
}

func (n *NSCalendar) IsDate_EqualToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) bool {
	result_ := C.C_NSCalendar_IsDate_EqualToDate_ToUnitGranularity(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2), C.uint(uint(unit)))
	return bool(result_)
}

func (n *NSCalendar) IsDate_InSameDayAsDate(date1 Date, date2 Date) bool {
	result_ := C.C_NSCalendar_IsDate_InSameDayAsDate(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2))
	return bool(result_)
}

func (n *NSCalendar) IsDateInToday(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInToday(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n *NSCalendar) IsDateInTomorrow(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInTomorrow(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n *NSCalendar) IsDateInWeekend(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInWeekend(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n *NSCalendar) IsDateInYesterday(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInYesterday(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func CurrentCalendar() Calendar {
	result_ := C.C_NSCalendar_CurrentCalendar()
	return MakeCalendar(result_)
}

func AutoupdatingCurrentCalendar() Calendar {
	result_ := C.C_NSCalendar_AutoupdatingCurrentCalendar()
	return MakeCalendar(result_)
}

func (n *NSCalendar) CalendarIdentifier() CalendarIdentifier {
	result_ := C.C_NSCalendar_CalendarIdentifier(n.Ptr())
	return CalendarIdentifier(MakeString(result_).String())
}

func (n *NSCalendar) FirstWeekday() uint {
	result_ := C.C_NSCalendar_FirstWeekday(n.Ptr())
	return uint(result_)
}

func (n *NSCalendar) SetFirstWeekday(value uint) {
	C.C_NSCalendar_SetFirstWeekday(n.Ptr(), C.uint(value))
}

func (n *NSCalendar) Locale() Locale {
	result_ := C.C_NSCalendar_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n *NSCalendar) SetLocale(value Locale) {
	C.C_NSCalendar_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSCalendar) TimeZone() TimeZone {
	result_ := C.C_NSCalendar_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n *NSCalendar) SetTimeZone(value TimeZone) {
	C.C_NSCalendar_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSCalendar) MinimumDaysInFirstWeek() uint {
	result_ := C.C_NSCalendar_MinimumDaysInFirstWeek(n.Ptr())
	return uint(result_)
}

func (n *NSCalendar) SetMinimumDaysInFirstWeek(value uint) {
	C.C_NSCalendar_SetMinimumDaysInFirstWeek(n.Ptr(), C.uint(value))
}

func (n *NSCalendar) AMSymbol() string {
	result_ := C.C_NSCalendar_AMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSCalendar) PMSymbol() string {
	result_ := C.C_NSCalendar_PMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSCalendar) WeekdaySymbols() []string {
	result_ := C.C_NSCalendar_WeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_ShortWeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) VeryShortWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_VeryShortWeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) StandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_StandaloneWeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneWeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) VeryShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_VeryShortStandaloneWeekdaySymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) MonthSymbols() []string {
	result_ := C.C_NSCalendar_MonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortMonthSymbols() []string {
	result_ := C.C_NSCalendar_ShortMonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) VeryShortMonthSymbols() []string {
	result_ := C.C_NSCalendar_VeryShortMonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) StandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_StandaloneMonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneMonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) VeryShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_VeryShortStandaloneMonthSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) QuarterSymbols() []string {
	result_ := C.C_NSCalendar_QuarterSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortQuarterSymbols() []string {
	result_ := C.C_NSCalendar_ShortQuarterSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) StandaloneQuarterSymbols() []string {
	result_ := C.C_NSCalendar_StandaloneQuarterSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) ShortStandaloneQuarterSymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneQuarterSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) EraSymbols() []string {
	result_ := C.C_NSCalendar_EraSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSCalendar) LongEraSymbols() []string {
	result_ := C.C_NSCalendar_LongEraSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
