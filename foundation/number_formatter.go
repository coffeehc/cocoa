package foundation

// #include "number_formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type NumberFormatter interface {
	Formatter
	NumberFromString(_string string) Number
	StringFromNumber(number Number) string
	FormatterBehavior() NumberFormatterBehavior
	SetFormatterBehavior(value NumberFormatterBehavior)
	NumberStyle() NumberFormatterStyle
	SetNumberStyle(value NumberFormatterStyle)
	GeneratesDecimalNumbers() bool
	SetGeneratesDecimalNumbers(value bool)
	LocalizesFormat() bool
	SetLocalizesFormat(value bool)
	Locale() Locale
	SetLocale(value Locale)
	RoundingBehavior() DecimalNumberHandler
	SetRoundingBehavior(value DecimalNumberHandler)
	RoundingIncrement() Number
	SetRoundingIncrement(value Number)
	RoundingMode() NumberFormatterRoundingMode
	SetRoundingMode(value NumberFormatterRoundingMode)
	MinimumIntegerDigits() uint
	SetMinimumIntegerDigits(value uint)
	MaximumIntegerDigits() uint
	SetMaximumIntegerDigits(value uint)
	MinimumFractionDigits() uint
	SetMinimumFractionDigits(value uint)
	MaximumFractionDigits() uint
	SetMaximumFractionDigits(value uint)
	UsesSignificantDigits() bool
	SetUsesSignificantDigits(value bool)
	MinimumSignificantDigits() uint
	SetMinimumSignificantDigits(value uint)
	MaximumSignificantDigits() uint
	SetMaximumSignificantDigits(value uint)
	Format() string
	SetFormat(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	FormatWidth() uint
	SetFormatWidth(value uint)
	NegativeFormat() string
	SetNegativeFormat(value string)
	PositiveFormat() string
	SetPositiveFormat(value string)
	Multiplier() Number
	SetMultiplier(value Number)
	PercentSymbol() string
	SetPercentSymbol(value string)
	PerMillSymbol() string
	SetPerMillSymbol(value string)
	MinusSign() string
	SetMinusSign(value string)
	PlusSign() string
	SetPlusSign(value string)
	ExponentSymbol() string
	SetExponentSymbol(value string)
	ZeroSymbol() string
	SetZeroSymbol(value string)
	NilSymbol() string
	SetNilSymbol(value string)
	NotANumberSymbol() string
	SetNotANumberSymbol(value string)
	NegativeInfinitySymbol() string
	SetNegativeInfinitySymbol(value string)
	PositiveInfinitySymbol() string
	SetPositiveInfinitySymbol(value string)
	CurrencySymbol() string
	SetCurrencySymbol(value string)
	CurrencyCode() string
	SetCurrencyCode(value string)
	InternationalCurrencySymbol() string
	SetInternationalCurrencySymbol(value string)
	CurrencyGroupingSeparator() string
	SetCurrencyGroupingSeparator(value string)
	PositivePrefix() string
	SetPositivePrefix(value string)
	PositiveSuffix() string
	SetPositiveSuffix(value string)
	NegativePrefix() string
	SetNegativePrefix(value string)
	NegativeSuffix() string
	SetNegativeSuffix(value string)
	AttributedStringForZero() AttributedString
	SetAttributedStringForZero(value AttributedString)
	AttributedStringForNil() AttributedString
	SetAttributedStringForNil(value AttributedString)
	AttributedStringForNotANumber() AttributedString
	SetAttributedStringForNotANumber(value AttributedString)
	GroupingSeparator() string
	SetGroupingSeparator(value string)
	UsesGroupingSeparator() bool
	SetUsesGroupingSeparator(value bool)
	ThousandSeparator() string
	SetThousandSeparator(value string)
	HasThousandSeparators() bool
	SetHasThousandSeparators(value bool)
	DecimalSeparator() string
	SetDecimalSeparator(value string)
	AlwaysShowsDecimalSeparator() bool
	SetAlwaysShowsDecimalSeparator(value bool)
	CurrencyDecimalSeparator() string
	SetCurrencyDecimalSeparator(value string)
	GroupingSize() uint
	SetGroupingSize(value uint)
	SecondaryGroupingSize() uint
	SetSecondaryGroupingSize(value uint)
	PaddingCharacter() string
	SetPaddingCharacter(value string)
	PaddingPosition() NumberFormatterPadPosition
	SetPaddingPosition(value NumberFormatterPadPosition)
	AllowsFloats() bool
	SetAllowsFloats(value bool)
	Minimum() Number
	SetMinimum(value Number)
	Maximum() Number
	SetMaximum(value Number)
	IsLenient() bool
	SetLenient(value bool)
	IsPartialStringValidationEnabled() bool
	SetPartialStringValidationEnabled(value bool)
}

type NSNumberFormatter struct {
	NSFormatter
}

func MakeNumberFormatter(ptr unsafe.Pointer) *NSNumberFormatter {
	if ptr == nil {
		return nil
	}
	return &NSNumberFormatter{
		NSFormatter: *MakeFormatter(ptr),
	}
}

func AllocNumberFormatter() *NSNumberFormatter {
	return MakeNumberFormatter(C.C_NumberFormatter_Alloc())
}

func (n *NSNumberFormatter) Init() NumberFormatter {
	result_ := C.C_NSNumberFormatter_Init(n.Ptr())
	return MakeNumberFormatter(result_)
}

func NumberFormatter_SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	C.C_NSNumberFormatter_NumberFormatter_SetDefaultFormatterBehavior(C.uint(uint(behavior)))
}

func NumberFormatter_DefaultFormatterBehavior() NumberFormatterBehavior {
	result_ := C.C_NSNumberFormatter_NumberFormatter_DefaultFormatterBehavior()
	return NumberFormatterBehavior(uint(result_))
}

func (n *NSNumberFormatter) NumberFromString(_string string) Number {
	result_ := C.C_NSNumberFormatter_NumberFromString(n.Ptr(), NewString(_string).Ptr())
	return MakeNumber(result_)
}

func (n *NSNumberFormatter) StringFromNumber(number Number) string {
	result_ := C.C_NSNumberFormatter_StringFromNumber(n.Ptr(), objc.ExtractPtr(number))
	return MakeString(result_).String()
}

func NumberFormatter_LocalizedStringFromNumber_NumberStyle(num Number, nstyle NumberFormatterStyle) string {
	result_ := C.C_NSNumberFormatter_NumberFormatter_LocalizedStringFromNumber_NumberStyle(objc.ExtractPtr(num), C.uint(uint(nstyle)))
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	result_ := C.C_NSNumberFormatter_FormatterBehavior(n.Ptr())
	return NumberFormatterBehavior(uint(result_))
}

func (n *NSNumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	C.C_NSNumberFormatter_SetFormatterBehavior(n.Ptr(), C.uint(uint(value)))
}

func (n *NSNumberFormatter) NumberStyle() NumberFormatterStyle {
	result_ := C.C_NSNumberFormatter_NumberStyle(n.Ptr())
	return NumberFormatterStyle(uint(result_))
}

func (n *NSNumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	C.C_NSNumberFormatter_SetNumberStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSNumberFormatter) GeneratesDecimalNumbers() bool {
	result_ := C.C_NSNumberFormatter_GeneratesDecimalNumbers(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	C.C_NSNumberFormatter_SetGeneratesDecimalNumbers(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) LocalizesFormat() bool {
	result_ := C.C_NSNumberFormatter_LocalizesFormat(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetLocalizesFormat(value bool) {
	C.C_NSNumberFormatter_SetLocalizesFormat(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) Locale() Locale {
	result_ := C.C_NSNumberFormatter_Locale(n.Ptr())
	return MakeLocale(result_)
}

func (n *NSNumberFormatter) SetLocale(value Locale) {
	C.C_NSNumberFormatter_SetLocale(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) RoundingBehavior() DecimalNumberHandler {
	result_ := C.C_NSNumberFormatter_RoundingBehavior(n.Ptr())
	return MakeDecimalNumberHandler(result_)
}

func (n *NSNumberFormatter) SetRoundingBehavior(value DecimalNumberHandler) {
	C.C_NSNumberFormatter_SetRoundingBehavior(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) RoundingIncrement() Number {
	result_ := C.C_NSNumberFormatter_RoundingIncrement(n.Ptr())
	return MakeNumber(result_)
}

func (n *NSNumberFormatter) SetRoundingIncrement(value Number) {
	C.C_NSNumberFormatter_SetRoundingIncrement(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	result_ := C.C_NSNumberFormatter_RoundingMode(n.Ptr())
	return NumberFormatterRoundingMode(uint(result_))
}

func (n *NSNumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	C.C_NSNumberFormatter_SetRoundingMode(n.Ptr(), C.uint(uint(value)))
}

func (n *NSNumberFormatter) MinimumIntegerDigits() uint {
	result_ := C.C_NSNumberFormatter_MinimumIntegerDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMinimumIntegerDigits(value uint) {
	C.C_NSNumberFormatter_SetMinimumIntegerDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) MaximumIntegerDigits() uint {
	result_ := C.C_NSNumberFormatter_MaximumIntegerDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMaximumIntegerDigits(value uint) {
	C.C_NSNumberFormatter_SetMaximumIntegerDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) MinimumFractionDigits() uint {
	result_ := C.C_NSNumberFormatter_MinimumFractionDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMinimumFractionDigits(value uint) {
	C.C_NSNumberFormatter_SetMinimumFractionDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) MaximumFractionDigits() uint {
	result_ := C.C_NSNumberFormatter_MaximumFractionDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMaximumFractionDigits(value uint) {
	C.C_NSNumberFormatter_SetMaximumFractionDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) UsesSignificantDigits() bool {
	result_ := C.C_NSNumberFormatter_UsesSignificantDigits(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetUsesSignificantDigits(value bool) {
	C.C_NSNumberFormatter_SetUsesSignificantDigits(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) MinimumSignificantDigits() uint {
	result_ := C.C_NSNumberFormatter_MinimumSignificantDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMinimumSignificantDigits(value uint) {
	C.C_NSNumberFormatter_SetMinimumSignificantDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) MaximumSignificantDigits() uint {
	result_ := C.C_NSNumberFormatter_MaximumSignificantDigits(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetMaximumSignificantDigits(value uint) {
	C.C_NSNumberFormatter_SetMaximumSignificantDigits(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) Format() string {
	result_ := C.C_NSNumberFormatter_Format(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetFormat(value string) {
	C.C_NSNumberFormatter_SetFormat(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) FormattingContext() FormattingContext {
	result_ := C.C_NSNumberFormatter_FormattingContext(n.Ptr())
	return FormattingContext(int(result_))
}

func (n *NSNumberFormatter) SetFormattingContext(value FormattingContext) {
	C.C_NSNumberFormatter_SetFormattingContext(n.Ptr(), C.int(int(value)))
}

func (n *NSNumberFormatter) FormatWidth() uint {
	result_ := C.C_NSNumberFormatter_FormatWidth(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetFormatWidth(value uint) {
	C.C_NSNumberFormatter_SetFormatWidth(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) NegativeFormat() string {
	result_ := C.C_NSNumberFormatter_NegativeFormat(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNegativeFormat(value string) {
	C.C_NSNumberFormatter_SetNegativeFormat(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PositiveFormat() string {
	result_ := C.C_NSNumberFormatter_PositiveFormat(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPositiveFormat(value string) {
	C.C_NSNumberFormatter_SetPositiveFormat(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) Multiplier() Number {
	result_ := C.C_NSNumberFormatter_Multiplier(n.Ptr())
	return MakeNumber(result_)
}

func (n *NSNumberFormatter) SetMultiplier(value Number) {
	C.C_NSNumberFormatter_SetMultiplier(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) PercentSymbol() string {
	result_ := C.C_NSNumberFormatter_PercentSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPercentSymbol(value string) {
	C.C_NSNumberFormatter_SetPercentSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PerMillSymbol() string {
	result_ := C.C_NSNumberFormatter_PerMillSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPerMillSymbol(value string) {
	C.C_NSNumberFormatter_SetPerMillSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) MinusSign() string {
	result_ := C.C_NSNumberFormatter_MinusSign(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetMinusSign(value string) {
	C.C_NSNumberFormatter_SetMinusSign(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PlusSign() string {
	result_ := C.C_NSNumberFormatter_PlusSign(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPlusSign(value string) {
	C.C_NSNumberFormatter_SetPlusSign(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) ExponentSymbol() string {
	result_ := C.C_NSNumberFormatter_ExponentSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetExponentSymbol(value string) {
	C.C_NSNumberFormatter_SetExponentSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) ZeroSymbol() string {
	result_ := C.C_NSNumberFormatter_ZeroSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetZeroSymbol(value string) {
	C.C_NSNumberFormatter_SetZeroSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) NilSymbol() string {
	result_ := C.C_NSNumberFormatter_NilSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNilSymbol(value string) {
	C.C_NSNumberFormatter_SetNilSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) NotANumberSymbol() string {
	result_ := C.C_NSNumberFormatter_NotANumberSymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNotANumberSymbol(value string) {
	C.C_NSNumberFormatter_SetNotANumberSymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) NegativeInfinitySymbol() string {
	result_ := C.C_NSNumberFormatter_NegativeInfinitySymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNegativeInfinitySymbol(value string) {
	C.C_NSNumberFormatter_SetNegativeInfinitySymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PositiveInfinitySymbol() string {
	result_ := C.C_NSNumberFormatter_PositiveInfinitySymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPositiveInfinitySymbol(value string) {
	C.C_NSNumberFormatter_SetPositiveInfinitySymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) CurrencySymbol() string {
	result_ := C.C_NSNumberFormatter_CurrencySymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetCurrencySymbol(value string) {
	C.C_NSNumberFormatter_SetCurrencySymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) CurrencyCode() string {
	result_ := C.C_NSNumberFormatter_CurrencyCode(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetCurrencyCode(value string) {
	C.C_NSNumberFormatter_SetCurrencyCode(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) InternationalCurrencySymbol() string {
	result_ := C.C_NSNumberFormatter_InternationalCurrencySymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetInternationalCurrencySymbol(value string) {
	C.C_NSNumberFormatter_SetInternationalCurrencySymbol(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) CurrencyGroupingSeparator() string {
	result_ := C.C_NSNumberFormatter_CurrencyGroupingSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetCurrencyGroupingSeparator(value string) {
	C.C_NSNumberFormatter_SetCurrencyGroupingSeparator(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PositivePrefix() string {
	result_ := C.C_NSNumberFormatter_PositivePrefix(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPositivePrefix(value string) {
	C.C_NSNumberFormatter_SetPositivePrefix(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PositiveSuffix() string {
	result_ := C.C_NSNumberFormatter_PositiveSuffix(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPositiveSuffix(value string) {
	C.C_NSNumberFormatter_SetPositiveSuffix(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) NegativePrefix() string {
	result_ := C.C_NSNumberFormatter_NegativePrefix(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNegativePrefix(value string) {
	C.C_NSNumberFormatter_SetNegativePrefix(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) NegativeSuffix() string {
	result_ := C.C_NSNumberFormatter_NegativeSuffix(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetNegativeSuffix(value string) {
	C.C_NSNumberFormatter_SetNegativeSuffix(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) AttributedStringForZero() AttributedString {
	result_ := C.C_NSNumberFormatter_AttributedStringForZero(n.Ptr())
	return MakeAttributedString(result_)
}

func (n *NSNumberFormatter) SetAttributedStringForZero(value AttributedString) {
	C.C_NSNumberFormatter_SetAttributedStringForZero(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) AttributedStringForNil() AttributedString {
	result_ := C.C_NSNumberFormatter_AttributedStringForNil(n.Ptr())
	return MakeAttributedString(result_)
}

func (n *NSNumberFormatter) SetAttributedStringForNil(value AttributedString) {
	C.C_NSNumberFormatter_SetAttributedStringForNil(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) AttributedStringForNotANumber() AttributedString {
	result_ := C.C_NSNumberFormatter_AttributedStringForNotANumber(n.Ptr())
	return MakeAttributedString(result_)
}

func (n *NSNumberFormatter) SetAttributedStringForNotANumber(value AttributedString) {
	C.C_NSNumberFormatter_SetAttributedStringForNotANumber(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) GroupingSeparator() string {
	result_ := C.C_NSNumberFormatter_GroupingSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetGroupingSeparator(value string) {
	C.C_NSNumberFormatter_SetGroupingSeparator(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) UsesGroupingSeparator() bool {
	result_ := C.C_NSNumberFormatter_UsesGroupingSeparator(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetUsesGroupingSeparator(value bool) {
	C.C_NSNumberFormatter_SetUsesGroupingSeparator(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) ThousandSeparator() string {
	result_ := C.C_NSNumberFormatter_ThousandSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetThousandSeparator(value string) {
	C.C_NSNumberFormatter_SetThousandSeparator(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) HasThousandSeparators() bool {
	result_ := C.C_NSNumberFormatter_HasThousandSeparators(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetHasThousandSeparators(value bool) {
	C.C_NSNumberFormatter_SetHasThousandSeparators(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) DecimalSeparator() string {
	result_ := C.C_NSNumberFormatter_DecimalSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetDecimalSeparator(value string) {
	C.C_NSNumberFormatter_SetDecimalSeparator(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) AlwaysShowsDecimalSeparator() bool {
	result_ := C.C_NSNumberFormatter_AlwaysShowsDecimalSeparator(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	C.C_NSNumberFormatter_SetAlwaysShowsDecimalSeparator(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) CurrencyDecimalSeparator() string {
	result_ := C.C_NSNumberFormatter_CurrencyDecimalSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetCurrencyDecimalSeparator(value string) {
	C.C_NSNumberFormatter_SetCurrencyDecimalSeparator(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) GroupingSize() uint {
	result_ := C.C_NSNumberFormatter_GroupingSize(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetGroupingSize(value uint) {
	C.C_NSNumberFormatter_SetGroupingSize(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) SecondaryGroupingSize() uint {
	result_ := C.C_NSNumberFormatter_SecondaryGroupingSize(n.Ptr())
	return uint(result_)
}

func (n *NSNumberFormatter) SetSecondaryGroupingSize(value uint) {
	C.C_NSNumberFormatter_SetSecondaryGroupingSize(n.Ptr(), C.uint(value))
}

func (n *NSNumberFormatter) PaddingCharacter() string {
	result_ := C.C_NSNumberFormatter_PaddingCharacter(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSNumberFormatter) SetPaddingCharacter(value string) {
	C.C_NSNumberFormatter_SetPaddingCharacter(n.Ptr(), NewString(value).Ptr())
}

func (n *NSNumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	result_ := C.C_NSNumberFormatter_PaddingPosition(n.Ptr())
	return NumberFormatterPadPosition(uint(result_))
}

func (n *NSNumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	C.C_NSNumberFormatter_SetPaddingPosition(n.Ptr(), C.uint(uint(value)))
}

func (n *NSNumberFormatter) AllowsFloats() bool {
	result_ := C.C_NSNumberFormatter_AllowsFloats(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetAllowsFloats(value bool) {
	C.C_NSNumberFormatter_SetAllowsFloats(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) Minimum() Number {
	result_ := C.C_NSNumberFormatter_Minimum(n.Ptr())
	return MakeNumber(result_)
}

func (n *NSNumberFormatter) SetMinimum(value Number) {
	C.C_NSNumberFormatter_SetMinimum(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) Maximum() Number {
	result_ := C.C_NSNumberFormatter_Maximum(n.Ptr())
	return MakeNumber(result_)
}

func (n *NSNumberFormatter) SetMaximum(value Number) {
	C.C_NSNumberFormatter_SetMaximum(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSNumberFormatter) IsLenient() bool {
	result_ := C.C_NSNumberFormatter_IsLenient(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetLenient(value bool) {
	C.C_NSNumberFormatter_SetLenient(n.Ptr(), C.bool(value))
}

func (n *NSNumberFormatter) IsPartialStringValidationEnabled() bool {
	result_ := C.C_NSNumberFormatter_IsPartialStringValidationEnabled(n.Ptr())
	return bool(result_)
}

func (n *NSNumberFormatter) SetPartialStringValidationEnabled(value bool) {
	C.C_NSNumberFormatter_SetPartialStringValidationEnabled(n.Ptr(), C.bool(value))
}
