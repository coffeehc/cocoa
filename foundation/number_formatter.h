#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_NumberFormatter_Alloc();

void* C_NSNumberFormatter_AllocNumberFormatter();
void* C_NSNumberFormatter_Init(void* ptr);
void* C_NSNumberFormatter_NewNumberFormatter();
void* C_NSNumberFormatter_Autorelease(void* ptr);
void* C_NSNumberFormatter_Retain(void* ptr);
void C_NSNumberFormatter_NumberFormatter_SetDefaultFormatterBehavior(unsigned int behavior);
unsigned int C_NSNumberFormatter_NumberFormatter_DefaultFormatterBehavior();
void* C_NSNumberFormatter_NumberFromString(void* ptr, void* _string);
void* C_NSNumberFormatter_StringFromNumber(void* ptr, void* number);
void* C_NSNumberFormatter_NumberFormatter_LocalizedStringFromNumber_NumberStyle(void* num, unsigned int nstyle);
unsigned int C_NSNumberFormatter_FormatterBehavior(void* ptr);
void C_NSNumberFormatter_SetFormatterBehavior(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_NumberStyle(void* ptr);
void C_NSNumberFormatter_SetNumberStyle(void* ptr, unsigned int value);
bool C_NSNumberFormatter_GeneratesDecimalNumbers(void* ptr);
void C_NSNumberFormatter_SetGeneratesDecimalNumbers(void* ptr, bool value);
bool C_NSNumberFormatter_LocalizesFormat(void* ptr);
void C_NSNumberFormatter_SetLocalizesFormat(void* ptr, bool value);
void* C_NSNumberFormatter_Locale(void* ptr);
void C_NSNumberFormatter_SetLocale(void* ptr, void* value);
void* C_NSNumberFormatter_RoundingBehavior(void* ptr);
void C_NSNumberFormatter_SetRoundingBehavior(void* ptr, void* value);
void* C_NSNumberFormatter_RoundingIncrement(void* ptr);
void C_NSNumberFormatter_SetRoundingIncrement(void* ptr, void* value);
unsigned int C_NSNumberFormatter_RoundingMode(void* ptr);
void C_NSNumberFormatter_SetRoundingMode(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_MinimumIntegerDigits(void* ptr);
void C_NSNumberFormatter_SetMinimumIntegerDigits(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_MaximumIntegerDigits(void* ptr);
void C_NSNumberFormatter_SetMaximumIntegerDigits(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_MinimumFractionDigits(void* ptr);
void C_NSNumberFormatter_SetMinimumFractionDigits(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_MaximumFractionDigits(void* ptr);
void C_NSNumberFormatter_SetMaximumFractionDigits(void* ptr, unsigned int value);
bool C_NSNumberFormatter_UsesSignificantDigits(void* ptr);
void C_NSNumberFormatter_SetUsesSignificantDigits(void* ptr, bool value);
unsigned int C_NSNumberFormatter_MinimumSignificantDigits(void* ptr);
void C_NSNumberFormatter_SetMinimumSignificantDigits(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_MaximumSignificantDigits(void* ptr);
void C_NSNumberFormatter_SetMaximumSignificantDigits(void* ptr, unsigned int value);
void* C_NSNumberFormatter_Format(void* ptr);
void C_NSNumberFormatter_SetFormat(void* ptr, void* value);
int C_NSNumberFormatter_FormattingContext(void* ptr);
void C_NSNumberFormatter_SetFormattingContext(void* ptr, int value);
unsigned int C_NSNumberFormatter_FormatWidth(void* ptr);
void C_NSNumberFormatter_SetFormatWidth(void* ptr, unsigned int value);
void* C_NSNumberFormatter_NegativeFormat(void* ptr);
void C_NSNumberFormatter_SetNegativeFormat(void* ptr, void* value);
void* C_NSNumberFormatter_PositiveFormat(void* ptr);
void C_NSNumberFormatter_SetPositiveFormat(void* ptr, void* value);
void* C_NSNumberFormatter_Multiplier(void* ptr);
void C_NSNumberFormatter_SetMultiplier(void* ptr, void* value);
void* C_NSNumberFormatter_PercentSymbol(void* ptr);
void C_NSNumberFormatter_SetPercentSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_PerMillSymbol(void* ptr);
void C_NSNumberFormatter_SetPerMillSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_MinusSign(void* ptr);
void C_NSNumberFormatter_SetMinusSign(void* ptr, void* value);
void* C_NSNumberFormatter_PlusSign(void* ptr);
void C_NSNumberFormatter_SetPlusSign(void* ptr, void* value);
void* C_NSNumberFormatter_ExponentSymbol(void* ptr);
void C_NSNumberFormatter_SetExponentSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_ZeroSymbol(void* ptr);
void C_NSNumberFormatter_SetZeroSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_NilSymbol(void* ptr);
void C_NSNumberFormatter_SetNilSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_NotANumberSymbol(void* ptr);
void C_NSNumberFormatter_SetNotANumberSymbol(void* ptr, void* value);
void* C_NSNumberFormatter_NegativeInfinitySymbol(void* ptr);
void C_NSNumberFormatter_SetNegativeInfinitySymbol(void* ptr, void* value);
void* C_NSNumberFormatter_PositiveInfinitySymbol(void* ptr);
void C_NSNumberFormatter_SetPositiveInfinitySymbol(void* ptr, void* value);
void* C_NSNumberFormatter_CurrencySymbol(void* ptr);
void C_NSNumberFormatter_SetCurrencySymbol(void* ptr, void* value);
void* C_NSNumberFormatter_CurrencyCode(void* ptr);
void C_NSNumberFormatter_SetCurrencyCode(void* ptr, void* value);
void* C_NSNumberFormatter_InternationalCurrencySymbol(void* ptr);
void C_NSNumberFormatter_SetInternationalCurrencySymbol(void* ptr, void* value);
void* C_NSNumberFormatter_CurrencyGroupingSeparator(void* ptr);
void C_NSNumberFormatter_SetCurrencyGroupingSeparator(void* ptr, void* value);
void* C_NSNumberFormatter_PositivePrefix(void* ptr);
void C_NSNumberFormatter_SetPositivePrefix(void* ptr, void* value);
void* C_NSNumberFormatter_PositiveSuffix(void* ptr);
void C_NSNumberFormatter_SetPositiveSuffix(void* ptr, void* value);
void* C_NSNumberFormatter_NegativePrefix(void* ptr);
void C_NSNumberFormatter_SetNegativePrefix(void* ptr, void* value);
void* C_NSNumberFormatter_NegativeSuffix(void* ptr);
void C_NSNumberFormatter_SetNegativeSuffix(void* ptr, void* value);
Dictionary C_NSNumberFormatter_TextAttributesForNegativeValues(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForNegativeValues(void* ptr, Dictionary value);
Dictionary C_NSNumberFormatter_TextAttributesForPositiveValues(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForPositiveValues(void* ptr, Dictionary value);
void* C_NSNumberFormatter_AttributedStringForZero(void* ptr);
void C_NSNumberFormatter_SetAttributedStringForZero(void* ptr, void* value);
Dictionary C_NSNumberFormatter_TextAttributesForZero(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForZero(void* ptr, Dictionary value);
void* C_NSNumberFormatter_AttributedStringForNil(void* ptr);
void C_NSNumberFormatter_SetAttributedStringForNil(void* ptr, void* value);
Dictionary C_NSNumberFormatter_TextAttributesForNil(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForNil(void* ptr, Dictionary value);
void* C_NSNumberFormatter_AttributedStringForNotANumber(void* ptr);
void C_NSNumberFormatter_SetAttributedStringForNotANumber(void* ptr, void* value);
Dictionary C_NSNumberFormatter_TextAttributesForNotANumber(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForNotANumber(void* ptr, Dictionary value);
Dictionary C_NSNumberFormatter_TextAttributesForPositiveInfinity(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForPositiveInfinity(void* ptr, Dictionary value);
Dictionary C_NSNumberFormatter_TextAttributesForNegativeInfinity(void* ptr);
void C_NSNumberFormatter_SetTextAttributesForNegativeInfinity(void* ptr, Dictionary value);
void* C_NSNumberFormatter_GroupingSeparator(void* ptr);
void C_NSNumberFormatter_SetGroupingSeparator(void* ptr, void* value);
bool C_NSNumberFormatter_UsesGroupingSeparator(void* ptr);
void C_NSNumberFormatter_SetUsesGroupingSeparator(void* ptr, bool value);
void* C_NSNumberFormatter_ThousandSeparator(void* ptr);
void C_NSNumberFormatter_SetThousandSeparator(void* ptr, void* value);
bool C_NSNumberFormatter_HasThousandSeparators(void* ptr);
void C_NSNumberFormatter_SetHasThousandSeparators(void* ptr, bool value);
void* C_NSNumberFormatter_DecimalSeparator(void* ptr);
void C_NSNumberFormatter_SetDecimalSeparator(void* ptr, void* value);
bool C_NSNumberFormatter_AlwaysShowsDecimalSeparator(void* ptr);
void C_NSNumberFormatter_SetAlwaysShowsDecimalSeparator(void* ptr, bool value);
void* C_NSNumberFormatter_CurrencyDecimalSeparator(void* ptr);
void C_NSNumberFormatter_SetCurrencyDecimalSeparator(void* ptr, void* value);
unsigned int C_NSNumberFormatter_GroupingSize(void* ptr);
void C_NSNumberFormatter_SetGroupingSize(void* ptr, unsigned int value);
unsigned int C_NSNumberFormatter_SecondaryGroupingSize(void* ptr);
void C_NSNumberFormatter_SetSecondaryGroupingSize(void* ptr, unsigned int value);
void* C_NSNumberFormatter_PaddingCharacter(void* ptr);
void C_NSNumberFormatter_SetPaddingCharacter(void* ptr, void* value);
unsigned int C_NSNumberFormatter_PaddingPosition(void* ptr);
void C_NSNumberFormatter_SetPaddingPosition(void* ptr, unsigned int value);
bool C_NSNumberFormatter_AllowsFloats(void* ptr);
void C_NSNumberFormatter_SetAllowsFloats(void* ptr, bool value);
void* C_NSNumberFormatter_Minimum(void* ptr);
void C_NSNumberFormatter_SetMinimum(void* ptr, void* value);
void* C_NSNumberFormatter_Maximum(void* ptr);
void C_NSNumberFormatter_SetMaximum(void* ptr, void* value);
bool C_NSNumberFormatter_IsLenient(void* ptr);
void C_NSNumberFormatter_SetLenient(void* ptr, bool value);
bool C_NSNumberFormatter_IsPartialStringValidationEnabled(void* ptr);
void C_NSNumberFormatter_SetPartialStringValidationEnabled(void* ptr, bool value);
