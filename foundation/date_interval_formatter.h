#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DateIntervalFormatter_Alloc();

void* C_NSDateIntervalFormatter_AllocDateIntervalFormatter();
void* C_NSDateIntervalFormatter_Init(void* ptr);
void* C_NSDateIntervalFormatter_NewDateIntervalFormatter();
void* C_NSDateIntervalFormatter_Autorelease(void* ptr);
void* C_NSDateIntervalFormatter_Retain(void* ptr);
void* C_NSDateIntervalFormatter_StringFromDate_ToDate(void* ptr, void* fromDate, void* toDate);
void* C_NSDateIntervalFormatter_StringFromDateInterval(void* ptr, void* dateInterval);
unsigned int C_NSDateIntervalFormatter_DateStyle(void* ptr);
void C_NSDateIntervalFormatter_SetDateStyle(void* ptr, unsigned int value);
unsigned int C_NSDateIntervalFormatter_TimeStyle(void* ptr);
void C_NSDateIntervalFormatter_SetTimeStyle(void* ptr, unsigned int value);
void* C_NSDateIntervalFormatter_DateTemplate(void* ptr);
void C_NSDateIntervalFormatter_SetDateTemplate(void* ptr, void* value);
void* C_NSDateIntervalFormatter_Calendar(void* ptr);
void C_NSDateIntervalFormatter_SetCalendar(void* ptr, void* value);
void* C_NSDateIntervalFormatter_Locale(void* ptr);
void C_NSDateIntervalFormatter_SetLocale(void* ptr, void* value);
void* C_NSDateIntervalFormatter_TimeZone(void* ptr);
void C_NSDateIntervalFormatter_SetTimeZone(void* ptr, void* value);
