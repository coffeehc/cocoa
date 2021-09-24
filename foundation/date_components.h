#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DateComponents_Alloc();

void* C_NSDateComponents_AllocDateComponents();
void* C_NSDateComponents_Init(void* ptr);
void* C_NSDateComponents_NewDateComponents();
void* C_NSDateComponents_Autorelease(void* ptr);
void* C_NSDateComponents_Retain(void* ptr);
bool C_NSDateComponents_IsValidDateInCalendar(void* ptr, void* calendar);
int C_NSDateComponents_ValueForComponent(void* ptr, unsigned int unit);
void C_NSDateComponents_SetValue_ForComponent(void* ptr, int value, unsigned int unit);
void* C_NSDateComponents_Calendar(void* ptr);
void C_NSDateComponents_SetCalendar(void* ptr, void* value);
void* C_NSDateComponents_TimeZone(void* ptr);
void C_NSDateComponents_SetTimeZone(void* ptr, void* value);
bool C_NSDateComponents_IsValidDate(void* ptr);
void* C_NSDateComponents_Date(void* ptr);
int C_NSDateComponents_Era(void* ptr);
void C_NSDateComponents_SetEra(void* ptr, int value);
int C_NSDateComponents_Year(void* ptr);
void C_NSDateComponents_SetYear(void* ptr, int value);
int C_NSDateComponents_YearForWeekOfYear(void* ptr);
void C_NSDateComponents_SetYearForWeekOfYear(void* ptr, int value);
int C_NSDateComponents_Quarter(void* ptr);
void C_NSDateComponents_SetQuarter(void* ptr, int value);
int C_NSDateComponents_Month(void* ptr);
void C_NSDateComponents_SetMonth(void* ptr, int value);
bool C_NSDateComponents_IsLeapMonth(void* ptr);
void C_NSDateComponents_SetLeapMonth(void* ptr, bool value);
int C_NSDateComponents_Weekday(void* ptr);
void C_NSDateComponents_SetWeekday(void* ptr, int value);
int C_NSDateComponents_WeekdayOrdinal(void* ptr);
void C_NSDateComponents_SetWeekdayOrdinal(void* ptr, int value);
int C_NSDateComponents_WeekOfMonth(void* ptr);
void C_NSDateComponents_SetWeekOfMonth(void* ptr, int value);
int C_NSDateComponents_WeekOfYear(void* ptr);
void C_NSDateComponents_SetWeekOfYear(void* ptr, int value);
int C_NSDateComponents_Day(void* ptr);
void C_NSDateComponents_SetDay(void* ptr, int value);
int C_NSDateComponents_Hour(void* ptr);
void C_NSDateComponents_SetHour(void* ptr, int value);
int C_NSDateComponents_Minute(void* ptr);
void C_NSDateComponents_SetMinute(void* ptr, int value);
int C_NSDateComponents_Second(void* ptr);
void C_NSDateComponents_SetSecond(void* ptr, int value);
int C_NSDateComponents_Nanosecond(void* ptr);
void C_NSDateComponents_SetNanosecond(void* ptr, int value);
