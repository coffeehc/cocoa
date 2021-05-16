#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Calendar_Alloc();

void* C_NSCalendar_CalendarWithIdentifier(void* calendarIdentifierConstant);
void* C_NSCalendar_InitWithCalendarIdentifier(void* ptr, void* ident);
bool C_NSCalendar_Date_MatchesComponents(void* ptr, void* date, void* components);
int C_NSCalendar_Component_FromDate(void* ptr, unsigned int unit, void* date);
void* C_NSCalendar_Components_FromDate(void* ptr, unsigned int unitFlags, void* date);
void* C_NSCalendar_Components_FromDate_ToDate_Options(void* ptr, unsigned int unitFlags, void* startingDate, void* resultDate, unsigned int opts);
void* C_NSCalendar_Components_FromDateComponents_ToDateComponents_Options(void* ptr, unsigned int unitFlags, void* startingDateComp, void* resultDateComp, unsigned int options);
void* C_NSCalendar_ComponentsInTimeZone_FromDate(void* ptr, void* timezone, void* date);
NSRange C_NSCalendar_MaximumRangeOfUnit(void* ptr, unsigned int unit);
NSRange C_NSCalendar_MinimumRangeOfUnit(void* ptr, unsigned int unit);
unsigned int C_NSCalendar_OrdinalityOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date);
NSRange C_NSCalendar_RangeOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date);
void* C_NSCalendar_StartOfDayForDate(void* ptr, void* date);
void* C_NSCalendar_NextDateAfterDate_MatchingComponents_Options(void* ptr, void* date, void* comps, unsigned int options);
void* C_NSCalendar_NextDateAfterDate_MatchingHour_Minute_Second_Options(void* ptr, void* date, int hourValue, int minuteValue, int secondValue, unsigned int options);
void* C_NSCalendar_NextDateAfterDate_MatchingUnit_Value_Options(void* ptr, void* date, unsigned int unit, int value, unsigned int options);
void* C_NSCalendar_DateFromComponents(void* ptr, void* comps);
void* C_NSCalendar_DateByAddingComponents_ToDate_Options(void* ptr, void* comps, void* date, unsigned int opts);
void* C_NSCalendar_DateByAddingUnit_Value_ToDate_Options(void* ptr, unsigned int unit, int value, void* date, unsigned int options);
void* C_NSCalendar_DateBySettingHour_Minute_Second_OfDate_Options(void* ptr, int h, int m, int s, void* date, unsigned int opts);
void* C_NSCalendar_DateBySettingUnit_Value_OfDate_Options(void* ptr, unsigned int unit, int v, void* date, unsigned int opts);
void* C_NSCalendar_DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int monthValue, int dayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue);
void* C_NSCalendar_DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int weekValue, int weekdayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue);
int C_NSCalendar_CompareDate_ToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit);
bool C_NSCalendar_IsDate_EqualToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit);
bool C_NSCalendar_IsDate_InSameDayAsDate(void* ptr, void* date1, void* date2);
bool C_NSCalendar_IsDateInToday(void* ptr, void* date);
bool C_NSCalendar_IsDateInTomorrow(void* ptr, void* date);
bool C_NSCalendar_IsDateInWeekend(void* ptr, void* date);
bool C_NSCalendar_IsDateInYesterday(void* ptr, void* date);
void* C_NSCalendar_CurrentCalendar();
void* C_NSCalendar_AutoupdatingCurrentCalendar();
void* C_NSCalendar_CalendarIdentifier(void* ptr);
unsigned int C_NSCalendar_FirstWeekday(void* ptr);
void C_NSCalendar_SetFirstWeekday(void* ptr, unsigned int value);
void* C_NSCalendar_Locale(void* ptr);
void C_NSCalendar_SetLocale(void* ptr, void* value);
void* C_NSCalendar_TimeZone(void* ptr);
void C_NSCalendar_SetTimeZone(void* ptr, void* value);
unsigned int C_NSCalendar_MinimumDaysInFirstWeek(void* ptr);
void C_NSCalendar_SetMinimumDaysInFirstWeek(void* ptr, unsigned int value);
void* C_NSCalendar_AMSymbol(void* ptr);
void* C_NSCalendar_PMSymbol(void* ptr);
Array C_NSCalendar_WeekdaySymbols(void* ptr);
Array C_NSCalendar_ShortWeekdaySymbols(void* ptr);
Array C_NSCalendar_VeryShortWeekdaySymbols(void* ptr);
Array C_NSCalendar_StandaloneWeekdaySymbols(void* ptr);
Array C_NSCalendar_ShortStandaloneWeekdaySymbols(void* ptr);
Array C_NSCalendar_VeryShortStandaloneWeekdaySymbols(void* ptr);
Array C_NSCalendar_MonthSymbols(void* ptr);
Array C_NSCalendar_ShortMonthSymbols(void* ptr);
Array C_NSCalendar_VeryShortMonthSymbols(void* ptr);
Array C_NSCalendar_StandaloneMonthSymbols(void* ptr);
Array C_NSCalendar_ShortStandaloneMonthSymbols(void* ptr);
Array C_NSCalendar_VeryShortStandaloneMonthSymbols(void* ptr);
Array C_NSCalendar_QuarterSymbols(void* ptr);
Array C_NSCalendar_ShortQuarterSymbols(void* ptr);
Array C_NSCalendar_StandaloneQuarterSymbols(void* ptr);
Array C_NSCalendar_ShortStandaloneQuarterSymbols(void* ptr);
Array C_NSCalendar_EraSymbols(void* ptr);
Array C_NSCalendar_LongEraSymbols(void* ptr);
