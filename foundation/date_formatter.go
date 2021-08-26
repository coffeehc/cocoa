package foundation

// #include "date_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DateFormatter interface {
	Formatter
	DateFromString(_string string) Date
	StringFromDate(date Date) string
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	DateFormat() string
	SetDateFormat(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Calendar() Calendar
	SetCalendar(value Calendar)
	DefaultDate() Date
	SetDefaultDate(value Date)
	Locale() Locale
	SetLocale(value Locale)
	TimeZone() TimeZone
	SetTimeZone(value TimeZone)
	TwoDigitStartDate() Date
	SetTwoDigitStartDate(value Date)
	GregorianStartDate() Date
	SetGregorianStartDate(value Date)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	IsLenient() bool
	SetLenient(value bool)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	AMSymbol() string
	SetAMSymbol(value string)
	PMSymbol() string
	SetPMSymbol(value string)
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)
	EraSymbols() []string
	SetEraSymbols(value []string)
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
}

type NSDateFormatter struct {
	NSFormatter
}

func MakeDateFormatter(ptr unsafe.Pointer) NSDateFormatter {
	return NSDateFormatter{
		NSFormatter: MakeFormatter(ptr),
	}
}

func AllocDateFormatter() NSDateFormatter {
	return MakeDateFormatter(C.C_DateFormatter_Alloc())
}

func (n NSDateFormatter) Init() DateFormatter {
	result_ := C.C_NSDateFormatter_Init(n.Ptr())
	return MakeDateFormatter(result_)
}

func (n NSDateFormatter) DateFromString(_string string) Date {
	result_ := C.C_NSDateFormatter_DateFromString(n.Ptr(), NewString(_string).Ptr())
	return MakeDate(result_)
}

func (n NSDateFormatter) StringFromDate(date Date) string {
	result_ := C.C_NSDateFormatter_StringFromDate(n.Ptr(), objc.ExtractPtr(date))
	return MakeString(result_).String()
}

func DateFormatter_LocalizedStringFromDate_DateStyle_TimeStyle(date Date, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	result_ := C.C_NSDateFormatter_DateFormatter_LocalizedStringFromDate_DateStyle_TimeStyle(objc.ExtractPtr(date), C.uint(uint(dstyle)), C.uint(uint(tstyle)))
	return MakeString(result_).String()
}

func (n NSDateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	C.C_NSDateFormatter_SetLocalizedDateFormatFromTemplate(n.Ptr(), NewString(dateFormatTemplate).Ptr())
}

func DateFormatter_DateFormatFromTemplate_Options_Locale(tmplate string, opts uint, locale Locale) string {
	result_ := C.C_NSDateFormatter_DateFormatter_DateFormatFromTemplate_Options_Locale(NewString(tmplate).Ptr(), C.uint(opts), objc.ExtractPtr(locale))
	return MakeString(result_).String()
}

func (n NSDateFormatter) DateStyle() DateFormatterStyle {
	result_ := C.C_NSDateFormatter_DateStyle(n.Ptr())
	return DateFormatterStyle(uint(result_))
}

func (n NSDateFormatter) SetDateStyle(value DateFormatterStyle) {
	C.C_NSDateFormatter_SetDateStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateFormatter) TimeStyle() DateFormatterStyle {
	result_ := C.C_NSDateFormatter_TimeStyle(n.Ptr())
	return DateFormatterStyle(uint(result_))
}

func (n NSDateFormatter) SetTimeStyle(value DateFormatterStyle) {
	C.C_NSDateFormatter_SetTimeStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSDateFormatter) DateFormat() string {
	result_ := C.C_NSDateFormatter_DateFormat(n.Ptr())
	return MakeString(result_).String()
}

func (n NSDateFormatter) SetDateFormat(value string) {
	C.C_NSDateFormatter_SetDateFormat(n.Ptr(), NewString(value).Ptr())
}

func (n NSDateFormatter) FormattingContext() FormattingContext {
	result_ := C.C_NSDateFormatter_FormattingContext(n.Ptr())
	return FormattingContext(int(result_))
}

func (n NSDateFormatter) SetFormattingContext(value FormattingContext) {
	C.C_NSDateFormatter_SetFormattingContext(n.Ptr(), C.int(int(value)))
}

func (n NSDateFormatter) Calendar() Calendar {
	result_ := C.C_NSDateFormatter_Calendar(n.Ptr())
	return MakeCalendar(result_)
}

func (n NSDateFormatter) SetCalendar(value Calendar) {
	C.C_NSDateFormatter_SetCalendar(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) DefaultDate() Date {
	result_ := C.C_NSDateFormatter_DefaultDate(n.Ptr())
	return MakeDate(result_)
}

func (n NSDateFormatter) SetDefaultDate(value Date) {
	C.C_NSDateFormatter_SetDefaultDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) Locale() Locale {
	result_ := C.C_NSDateFormatter_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n NSDateFormatter) SetLocale(value Locale) {
	C.C_NSDateFormatter_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) TimeZone() TimeZone {
	result_ := C.C_NSDateFormatter_TimeZone(n.Ptr())
	return MakeTimeZone(result_)
}

func (n NSDateFormatter) SetTimeZone(value TimeZone) {
	C.C_NSDateFormatter_SetTimeZone(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) TwoDigitStartDate() Date {
	result_ := C.C_NSDateFormatter_TwoDigitStartDate(n.Ptr())
	return MakeDate(result_)
}

func (n NSDateFormatter) SetTwoDigitStartDate(value Date) {
	C.C_NSDateFormatter_SetTwoDigitStartDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) GregorianStartDate() Date {
	result_ := C.C_NSDateFormatter_GregorianStartDate(n.Ptr())
	return MakeDate(result_)
}

func (n NSDateFormatter) SetGregorianStartDate(value Date) {
	C.C_NSDateFormatter_SetGregorianStartDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDateFormatter) FormatterBehavior() DateFormatterBehavior {
	result_ := C.C_NSDateFormatter_FormatterBehavior(n.Ptr())
	return DateFormatterBehavior(uint(result_))
}

func (n NSDateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	C.C_NSDateFormatter_SetFormatterBehavior(n.Ptr(), C.uint(uint(value)))
}

func DateFormatter_DefaultFormatterBehavior() DateFormatterBehavior {
	result_ := C.C_NSDateFormatter_DateFormatter_DefaultFormatterBehavior()
	return DateFormatterBehavior(uint(result_))
}

func DateFormatter_SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	C.C_NSDateFormatter_DateFormatter_SetDefaultFormatterBehavior(C.uint(uint(value)))
}

func (n NSDateFormatter) IsLenient() bool {
	result_ := C.C_NSDateFormatter_IsLenient(n.Ptr())
	return bool(result_)
}

func (n NSDateFormatter) SetLenient(value bool) {
	C.C_NSDateFormatter_SetLenient(n.Ptr(), C.bool(value))
}

func (n NSDateFormatter) DoesRelativeDateFormatting() bool {
	result_ := C.C_NSDateFormatter_DoesRelativeDateFormatting(n.Ptr())
	return bool(result_)
}

func (n NSDateFormatter) SetDoesRelativeDateFormatting(value bool) {
	C.C_NSDateFormatter_SetDoesRelativeDateFormatting(n.Ptr(), C.bool(value))
}

func (n NSDateFormatter) AMSymbol() string {
	result_ := C.C_NSDateFormatter_AMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n NSDateFormatter) SetAMSymbol(value string) {
	C.C_NSDateFormatter_SetAMSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n NSDateFormatter) PMSymbol() string {
	result_ := C.C_NSDateFormatter_PMSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n NSDateFormatter) SetPMSymbol(value string) {
	C.C_NSDateFormatter_SetPMSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n NSDateFormatter) WeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_WeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortWeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_ShortWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) VeryShortWeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_VeryShortWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetVeryShortWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) StandaloneWeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_StandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetStandaloneWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_ShortStandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortStandaloneWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	result_ := C.C_NSDateFormatter_VeryShortStandaloneWeekdaySymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetVeryShortStandaloneWeekdaySymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) MonthSymbols() []string {
	result_ := C.C_NSDateFormatter_MonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortMonthSymbols() []string {
	result_ := C.C_NSDateFormatter_ShortMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) VeryShortMonthSymbols() []string {
	result_ := C.C_NSDateFormatter_VeryShortMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetVeryShortMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetVeryShortMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) StandaloneMonthSymbols() []string {
	result_ := C.C_NSDateFormatter_StandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetStandaloneMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetStandaloneMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSDateFormatter_ShortStandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortStandaloneMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) VeryShortStandaloneMonthSymbols() []string {
	result_ := C.C_NSDateFormatter_VeryShortStandaloneMonthSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetVeryShortStandaloneMonthSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) QuarterSymbols() []string {
	result_ := C.C_NSDateFormatter_QuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetQuarterSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetQuarterSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortQuarterSymbols() []string {
	result_ := C.C_NSDateFormatter_ShortQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortQuarterSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortQuarterSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) StandaloneQuarterSymbols() []string {
	result_ := C.C_NSDateFormatter_StandaloneQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetStandaloneQuarterSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetStandaloneQuarterSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) ShortStandaloneQuarterSymbols() []string {
	result_ := C.C_NSDateFormatter_ShortStandaloneQuarterSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetShortStandaloneQuarterSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) EraSymbols() []string {
	result_ := C.C_NSDateFormatter_EraSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetEraSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetEraSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) LongEraSymbols() []string {
	result_ := C.C_NSDateFormatter_LongEraSymbols(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSDateFormatter) SetLongEraSymbols(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSDateFormatter_SetLongEraSymbols(n.Ptr(), cValue)
}

func (n NSDateFormatter) GeneratesCalendarDates() bool {
	result_ := C.C_NSDateFormatter_GeneratesCalendarDates(n.Ptr())
	return bool(result_)
}

func (n NSDateFormatter) SetGeneratesCalendarDates(value bool) {
	C.C_NSDateFormatter_SetGeneratesCalendarDates(n.Ptr(), C.bool(value))
}
