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

func MakeCalendar(ptr unsafe.Pointer) NSCalendar {
	return NSCalendar{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCalendar() NSCalendar {
	result_ := C.C_NSCalendar_AllocCalendar()
	return MakeCalendar(result_)
}

func (n NSCalendar) Autorelease() NSCalendar {
	result_ := C.C_NSCalendar_Autorelease(n.Ptr())
	return MakeCalendar(result_)
}

func (n NSCalendar) Retain() NSCalendar {
	result_ := C.C_NSCalendar_Retain(n.Ptr())
	return MakeCalendar(result_)
}

func CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	result_ := C.C_NSCalendar_CalendarWithIdentifier(NewString(string(calendarIdentifierConstant)).Ptr())
	return MakeCalendar(result_)
}

func (n NSCalendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	result_ := C.C_NSCalendar_InitWithCalendarIdentifier(n.Ptr(), NewString(string(ident)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSCalendar) Date_MatchesComponents(date Date, components DateComponents) bool {
	result_ := C.C_NSCalendar_Date_MatchesComponents(n.Ptr(), objc.ExtractPtr(date), objc.ExtractPtr(components))
	return bool(result_)
}

func (n NSCalendar) Component_FromDate(unit CalendarUnit, date Date) int {
	result_ := C.C_NSCalendar_Component_FromDate(n.Ptr(), C.uint(uint(unit)), objc.ExtractPtr(date))
	return int(result_)
}

func (n NSCalendar) Components_FromDate(unitFlags CalendarUnit, date Date) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDate(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(date))
	return MakeDateComponents(result_)
}

func (n NSCalendar) Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate Date, resultDate Date, opts CalendarOptions) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDate_ToDate_Options(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(startingDate), objc.ExtractPtr(resultDate), C.uint(uint(opts)))
	return MakeDateComponents(result_)
}

func (n NSCalendar) Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp DateComponents, resultDateComp DateComponents, options CalendarOptions) DateComponents {
	result_ := C.C_NSCalendar_Components_FromDateComponents_ToDateComponents_Options(n.Ptr(), C.uint(uint(unitFlags)), objc.ExtractPtr(startingDateComp), objc.ExtractPtr(resultDateComp), C.uint(uint(options)))
	return MakeDateComponents(result_)
}

func (n NSCalendar) ComponentsInTimeZone_FromDate(timezone TimeZone, date Date) DateComponents {
	result_ := C.C_NSCalendar_ComponentsInTimeZone_FromDate(n.Ptr(), objc.ExtractPtr(timezone), objc.ExtractPtr(date))
	return MakeDateComponents(result_)
}

func (n NSCalendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	result_ := C.C_NSCalendar_MaximumRangeOfUnit(n.Ptr(), C.uint(uint(unit)))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSCalendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	result_ := C.C_NSCalendar_MinimumRangeOfUnit(n.Ptr(), C.uint(uint(unit)))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSCalendar) OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) uint {
	result_ := C.C_NSCalendar_OrdinalityOfUnit_InUnit_ForDate(n.Ptr(), C.uint(uint(smaller)), C.uint(uint(larger)), objc.ExtractPtr(date))
	return uint(result_)
}

func (n NSCalendar) RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date Date) Range {
	result_ := C.C_NSCalendar_RangeOfUnit_InUnit_ForDate(n.Ptr(), C.uint(uint(smaller)), C.uint(uint(larger)), objc.ExtractPtr(date))
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSCalendar) StartOfDayForDate(date Date) Date {
	result_ := C.C_NSCalendar_StartOfDayForDate(n.Ptr(), objc.ExtractPtr(date))
	return MakeDate(result_)
}

func (n NSCalendar) NextDateAfterDate_MatchingComponents_Options(date Date, comps DateComponents, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingComponents_Options(n.Ptr(), objc.ExtractPtr(date), objc.ExtractPtr(comps), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n NSCalendar) NextDateAfterDate_MatchingHour_Minute_Second_Options(date Date, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingHour_Minute_Second_Options(n.Ptr(), objc.ExtractPtr(date), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n NSCalendar) NextDateAfterDate_MatchingUnit_Value_Options(date Date, unit CalendarUnit, value int, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_NextDateAfterDate_MatchingUnit_Value_Options(n.Ptr(), objc.ExtractPtr(date), C.uint(uint(unit)), C.int(value), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n NSCalendar) DateFromComponents(comps DateComponents) Date {
	result_ := C.C_NSCalendar_DateFromComponents(n.Ptr(), objc.ExtractPtr(comps))
	return MakeDate(result_)
}

func (n NSCalendar) DateByAddingComponents_ToDate_Options(comps DateComponents, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateByAddingComponents_ToDate_Options(n.Ptr(), objc.ExtractPtr(comps), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n NSCalendar) DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date Date, options CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateByAddingUnit_Value_ToDate_Options(n.Ptr(), C.uint(uint(unit)), C.int(value), objc.ExtractPtr(date), C.uint(uint(options)))
	return MakeDate(result_)
}

func (n NSCalendar) DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateBySettingHour_Minute_Second_OfDate_Options(n.Ptr(), C.int(h), C.int(m), C.int(s), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n NSCalendar) DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date Date, opts CalendarOptions) Date {
	result_ := C.C_NSCalendar_DateBySettingUnit_Value_OfDate_Options(n.Ptr(), C.uint(uint(unit)), C.int(v), objc.ExtractPtr(date), C.uint(uint(opts)))
	return MakeDate(result_)
}

func (n NSCalendar) DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	result_ := C.C_NSCalendar_DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(n.Ptr(), C.int(eraValue), C.int(yearValue), C.int(monthValue), C.int(dayValue), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.int(nanosecondValue))
	return MakeDate(result_)
}

func (n NSCalendar) DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	result_ := C.C_NSCalendar_DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(n.Ptr(), C.int(eraValue), C.int(yearValue), C.int(weekValue), C.int(weekdayValue), C.int(hourValue), C.int(minuteValue), C.int(secondValue), C.int(nanosecondValue))
	return MakeDate(result_)
}

func (n NSCalendar) CompareDate_ToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) ComparisonResult {
	result_ := C.C_NSCalendar_CompareDate_ToDate_ToUnitGranularity(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2), C.uint(uint(unit)))
	return ComparisonResult(int(result_))
}

func (n NSCalendar) IsDate_EqualToDate_ToUnitGranularity(date1 Date, date2 Date, unit CalendarUnit) bool {
	result_ := C.C_NSCalendar_IsDate_EqualToDate_ToUnitGranularity(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2), C.uint(uint(unit)))
	return bool(result_)
}

func (n NSCalendar) IsDate_InSameDayAsDate(date1 Date, date2 Date) bool {
	result_ := C.C_NSCalendar_IsDate_InSameDayAsDate(n.Ptr(), objc.ExtractPtr(date1), objc.ExtractPtr(date2))
	return bool(result_)
}

func (n NSCalendar) IsDateInToday(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInToday(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n NSCalendar) IsDateInTomorrow(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInTomorrow(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n NSCalendar) IsDateInWeekend(date Date) bool {
	result_ := C.C_NSCalendar_IsDateInWeekend(n.Ptr(), objc.ExtractPtr(date))
	return bool(result_)
}

func (n NSCalendar) IsDateInYesterday(date Date) bool {
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

func (n NSCalendar) CalendarIdentifier() CalendarIdentifier {
	result_ := C.C_NSCalendar_CalendarIdentifier(n.Ptr())
	return CalendarIdentifier(MakeString(result_).String())
}

func (n NSCalendar) FirstWeekday() uint {
	result_ := C.C_NSCalendar_FirstWeekday(n.Ptr())
	return uint(result_)
}

func (n NSCalendar) SetFirstWeekday(value uint) {
	C.C_NSCalendar_SetFirstWeekday(n.Ptr(), C.uint(value))
}

func (n NSCalendar) Locale() Locale {
	result_ := C.C_NSCalendar_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n NSCalendar) SetLocale(value Locale) {
	C.C_NSCalendar_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCalendar) TimeZone() TimeZone {
	result_ := C.C_NSCalendar_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n NSCalendar) SetTimeZone(value TimeZone) {
	C.C_NSCalendar_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCalendar) MinimumDaysInFirstWeek() uint {
	result_ := C.C_NSCalendar_MinimumDaysInFirstWeek(n.Ptr())
	return uint(result_)
}

func (n NSCalendar) SetMinimumDaysInFirstWeek(value uint) {
	C.C_NSCalendar_SetMinimumDaysInFirstWeek(n.Ptr(), C.uint(value))
}

func (n NSCalendar) AMSymbol() string {
	result_ := C.C_NSCalendar_AMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n NSCalendar) PMSymbol() string {
	result_ := C.C_NSCalendar_PMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n NSCalendar) WeekdaySymbols() []string {
	result_ := C.C_NSCalendar_WeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_ShortWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) VeryShortWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_VeryShortWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) StandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_StandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) VeryShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSCalendar_VeryShortStandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) MonthSymbols() []string {
	result_ := C.C_NSCalendar_MonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortMonthSymbols() []string {
	result_ := C.C_NSCalendar_ShortMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) VeryShortMonthSymbols() []string {
	result_ := C.C_NSCalendar_VeryShortMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) StandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_StandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) VeryShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSCalendar_VeryShortStandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) QuarterSymbols() []string {
	result_ := C.C_NSCalendar_QuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortQuarterSymbols() []string {
	result_ := C.C_NSCalendar_ShortQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) StandaloneQuarterSymbols() []string {
	result_ := C.C_NSCalendar_StandaloneQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) ShortStandaloneQuarterSymbols() []string {
	result_ := C.C_NSCalendar_ShortStandaloneQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) EraSymbols() []string {
	result_ := C.C_NSCalendar_EraSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSCalendar) LongEraSymbols() []string {
	result_ := C.C_NSCalendar_LongEraSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

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

func MakeDateComponents(ptr unsafe.Pointer) NSDateComponents {
	return NSDateComponents{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDateComponents() NSDateComponents {
	result_ := C.C_NSDateComponents_AllocDateComponents()
	return MakeDateComponents(result_)
}

func (n NSDateComponents) Init() NSDateComponents {
	result_ := C.C_NSDateComponents_Init(n.Ptr())
	return MakeDateComponents(result_)
}

func NewDateComponents() NSDateComponents {
	result_ := C.C_NSDateComponents_NewDateComponents()
	return MakeDateComponents(result_)
}

func (n NSDateComponents) Autorelease() NSDateComponents {
	result_ := C.C_NSDateComponents_Autorelease(n.Ptr())
	return MakeDateComponents(result_)
}

func (n NSDateComponents) Retain() NSDateComponents {
	result_ := C.C_NSDateComponents_Retain(n.Ptr())
	return MakeDateComponents(result_)
}

func (n NSDateComponents) IsValidDateInCalendar(calendar Calendar) bool {
	result_ := C.C_NSDateComponents_IsValidDateInCalendar(n.Ptr(), objc.ExtractPtr(calendar))
	return bool(result_)
}

func (n NSDateComponents) ValueForComponent(unit CalendarUnit) int {
	result_ := C.C_NSDateComponents_ValueForComponent(n.Ptr(), C.uint(uint(unit)))
	return int(result_)
}

func (n NSDateComponents) SetValue_ForComponent(value int, unit CalendarUnit) {
	C.C_NSDateComponents_SetValue_ForComponent(n.Ptr(), C.int(value), C.uint(uint(unit)))
}

func (n NSDateComponents) Calendar() Calendar {
	result_ := C.C_NSDateComponents_Calendar(n.Ptr())
	return MakeCalendar(result_)
}

func (n NSDateComponents) SetCalendar(value Calendar) {
	C.C_NSDateComponents_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateComponents) TimeZone() TimeZone {
	result_ := C.C_NSDateComponents_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n NSDateComponents) SetTimeZone(value TimeZone) {
	C.C_NSDateComponents_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateComponents) IsValidDate() bool {
	result_ := C.C_NSDateComponents_IsValidDate(n.Ptr())
	return bool(result_)
}

func (n NSDateComponents) Date() Date {
	result_ := C.C_NSDateComponents_Date(n.Ptr())
	return MakeDate(result_)
}

func (n NSDateComponents) Era() int {
	result_ := C.C_NSDateComponents_Era(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetEra(value int) {
	C.C_NSDateComponents_SetEra(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Year() int {
	result_ := C.C_NSDateComponents_Year(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetYear(value int) {
	C.C_NSDateComponents_SetYear(n.Ptr(), C.int(value))
}

func (n NSDateComponents) YearForWeekOfYear() int {
	result_ := C.C_NSDateComponents_YearForWeekOfYear(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetYearForWeekOfYear(value int) {
	C.C_NSDateComponents_SetYearForWeekOfYear(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Quarter() int {
	result_ := C.C_NSDateComponents_Quarter(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetQuarter(value int) {
	C.C_NSDateComponents_SetQuarter(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Month() int {
	result_ := C.C_NSDateComponents_Month(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetMonth(value int) {
	C.C_NSDateComponents_SetMonth(n.Ptr(), C.int(value))
}

func (n NSDateComponents) IsLeapMonth() bool {
	result_ := C.C_NSDateComponents_IsLeapMonth(n.Ptr())
	return bool(result_)
}

func (n NSDateComponents) SetLeapMonth(value bool) {
	C.C_NSDateComponents_SetLeapMonth(n.Ptr(), C.bool(value))
}

func (n NSDateComponents) Weekday() int {
	result_ := C.C_NSDateComponents_Weekday(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetWeekday(value int) {
	C.C_NSDateComponents_SetWeekday(n.Ptr(), C.int(value))
}

func (n NSDateComponents) WeekdayOrdinal() int {
	result_ := C.C_NSDateComponents_WeekdayOrdinal(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetWeekdayOrdinal(value int) {
	C.C_NSDateComponents_SetWeekdayOrdinal(n.Ptr(), C.int(value))
}

func (n NSDateComponents) WeekOfMonth() int {
	result_ := C.C_NSDateComponents_WeekOfMonth(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetWeekOfMonth(value int) {
	C.C_NSDateComponents_SetWeekOfMonth(n.Ptr(), C.int(value))
}

func (n NSDateComponents) WeekOfYear() int {
	result_ := C.C_NSDateComponents_WeekOfYear(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetWeekOfYear(value int) {
	C.C_NSDateComponents_SetWeekOfYear(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Day() int {
	result_ := C.C_NSDateComponents_Day(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetDay(value int) {
	C.C_NSDateComponents_SetDay(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Hour() int {
	result_ := C.C_NSDateComponents_Hour(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetHour(value int) {
	C.C_NSDateComponents_SetHour(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Minute() int {
	result_ := C.C_NSDateComponents_Minute(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetMinute(value int) {
	C.C_NSDateComponents_SetMinute(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Second() int {
	result_ := C.C_NSDateComponents_Second(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetSecond(value int) {
	C.C_NSDateComponents_SetSecond(n.Ptr(), C.int(value))
}

func (n NSDateComponents) Nanosecond() int {
	result_ := C.C_NSDateComponents_Nanosecond(n.Ptr())
	return int(result_)
}

func (n NSDateComponents) SetNanosecond(value int) {
	C.C_NSDateComponents_SetNanosecond(n.Ptr(), C.int(value))
}
