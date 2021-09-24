#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DateFormatter_Alloc();

void* C_NSDateFormatter_AllocDateFormatter();
void* C_NSDateFormatter_Init(void* ptr);
void* C_NSDateFormatter_NewDateFormatter();
void* C_NSDateFormatter_Autorelease(void* ptr);
void* C_NSDateFormatter_Retain(void* ptr);
void* C_NSDateFormatter_DateFromString(void* ptr, void* _string);
void* C_NSDateFormatter_StringFromDate(void* ptr, void* date);
void* C_NSDateFormatter_DateFormatter_LocalizedStringFromDate_DateStyle_TimeStyle(void* date, unsigned int dstyle, unsigned int tstyle);
void C_NSDateFormatter_SetLocalizedDateFormatFromTemplate(void* ptr, void* dateFormatTemplate);
void* C_NSDateFormatter_DateFormatter_DateFormatFromTemplate_Options_Locale(void* tmplate, unsigned int opts, void* locale);
unsigned int C_NSDateFormatter_DateStyle(void* ptr);
void C_NSDateFormatter_SetDateStyle(void* ptr, unsigned int value);
unsigned int C_NSDateFormatter_TimeStyle(void* ptr);
void C_NSDateFormatter_SetTimeStyle(void* ptr, unsigned int value);
void* C_NSDateFormatter_DateFormat(void* ptr);
void C_NSDateFormatter_SetDateFormat(void* ptr, void* value);
int C_NSDateFormatter_FormattingContext(void* ptr);
void C_NSDateFormatter_SetFormattingContext(void* ptr, int value);
void* C_NSDateFormatter_Calendar(void* ptr);
void C_NSDateFormatter_SetCalendar(void* ptr, void* value);
void* C_NSDateFormatter_DefaultDate(void* ptr);
void C_NSDateFormatter_SetDefaultDate(void* ptr, void* value);
void* C_NSDateFormatter_Locale(void* ptr);
void C_NSDateFormatter_SetLocale(void* ptr, void* value);
void* C_NSDateFormatter_TimeZone(void* ptr);
void C_NSDateFormatter_SetTimeZone(void* ptr, void* value);
void* C_NSDateFormatter_TwoDigitStartDate(void* ptr);
void C_NSDateFormatter_SetTwoDigitStartDate(void* ptr, void* value);
void* C_NSDateFormatter_GregorianStartDate(void* ptr);
void C_NSDateFormatter_SetGregorianStartDate(void* ptr, void* value);
unsigned int C_NSDateFormatter_FormatterBehavior(void* ptr);
void C_NSDateFormatter_SetFormatterBehavior(void* ptr, unsigned int value);
unsigned int C_NSDateFormatter_DateFormatter_DefaultFormatterBehavior();
void C_NSDateFormatter_DateFormatter_SetDefaultFormatterBehavior(unsigned int value);
bool C_NSDateFormatter_IsLenient(void* ptr);
void C_NSDateFormatter_SetLenient(void* ptr, bool value);
bool C_NSDateFormatter_DoesRelativeDateFormatting(void* ptr);
void C_NSDateFormatter_SetDoesRelativeDateFormatting(void* ptr, bool value);
void* C_NSDateFormatter_AMSymbol(void* ptr);
void C_NSDateFormatter_SetAMSymbol(void* ptr, void* value);
void* C_NSDateFormatter_PMSymbol(void* ptr);
void C_NSDateFormatter_SetPMSymbol(void* ptr, void* value);
Array C_NSDateFormatter_WeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortWeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetShortWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_VeryShortWeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetVeryShortWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_StandaloneWeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetStandaloneWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortStandaloneWeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetShortStandaloneWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_VeryShortStandaloneWeekdaySymbols(void* ptr);
void C_NSDateFormatter_SetVeryShortStandaloneWeekdaySymbols(void* ptr, Array value);
Array C_NSDateFormatter_MonthSymbols(void* ptr);
void C_NSDateFormatter_SetMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortMonthSymbols(void* ptr);
void C_NSDateFormatter_SetShortMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_VeryShortMonthSymbols(void* ptr);
void C_NSDateFormatter_SetVeryShortMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_StandaloneMonthSymbols(void* ptr);
void C_NSDateFormatter_SetStandaloneMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortStandaloneMonthSymbols(void* ptr);
void C_NSDateFormatter_SetShortStandaloneMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_VeryShortStandaloneMonthSymbols(void* ptr);
void C_NSDateFormatter_SetVeryShortStandaloneMonthSymbols(void* ptr, Array value);
Array C_NSDateFormatter_QuarterSymbols(void* ptr);
void C_NSDateFormatter_SetQuarterSymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortQuarterSymbols(void* ptr);
void C_NSDateFormatter_SetShortQuarterSymbols(void* ptr, Array value);
Array C_NSDateFormatter_StandaloneQuarterSymbols(void* ptr);
void C_NSDateFormatter_SetStandaloneQuarterSymbols(void* ptr, Array value);
Array C_NSDateFormatter_ShortStandaloneQuarterSymbols(void* ptr);
void C_NSDateFormatter_SetShortStandaloneQuarterSymbols(void* ptr, Array value);
Array C_NSDateFormatter_EraSymbols(void* ptr);
void C_NSDateFormatter_SetEraSymbols(void* ptr, Array value);
Array C_NSDateFormatter_LongEraSymbols(void* ptr);
void C_NSDateFormatter_SetLongEraSymbols(void* ptr, Array value);
bool C_NSDateFormatter_GeneratesCalendarDates(void* ptr);
void C_NSDateFormatter_SetGeneratesCalendarDates(void* ptr, bool value);
