#import "date_components.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSCalendar.h>

void* C_DateComponents_Alloc() {
    return [NSDateComponents alloc];
}

void* C_NSDateComponents_AllocDateComponents() {
    NSDateComponents* result_ = [NSDateComponents alloc];
    return result_;
}

void* C_NSDateComponents_Init(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSDateComponents* result_ = [nSDateComponents init];
    return result_;
}

void* C_NSDateComponents_NewDateComponents() {
    NSDateComponents* result_ = [NSDateComponents new];
    return result_;
}

void* C_NSDateComponents_Autorelease(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSDateComponents* result_ = [nSDateComponents autorelease];
    return result_;
}

void* C_NSDateComponents_Retain(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSDateComponents* result_ = [nSDateComponents retain];
    return result_;
}

bool C_NSDateComponents_IsValidDateInCalendar(void* ptr, void* calendar) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    BOOL result_ = [nSDateComponents isValidDateInCalendar:(NSCalendar*)calendar];
    return result_;
}

int C_NSDateComponents_ValueForComponent(void* ptr, unsigned int unit) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents valueForComponent:unit];
    return result_;
}

void C_NSDateComponents_SetValue_ForComponent(void* ptr, int value, unsigned int unit) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setValue:value forComponent:unit];
}

void* C_NSDateComponents_Calendar(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSCalendar* result_ = [nSDateComponents calendar];
    return result_;
}

void C_NSDateComponents_SetCalendar(void* ptr, void* value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setCalendar:(NSCalendar*)value];
}

void* C_NSDateComponents_TimeZone(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSTimeZone* result_ = [nSDateComponents timeZone];
    return result_;
}

void C_NSDateComponents_SetTimeZone(void* ptr, void* value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setTimeZone:(NSTimeZone*)value];
}

bool C_NSDateComponents_IsValidDate(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    BOOL result_ = [nSDateComponents isValidDate];
    return result_;
}

void* C_NSDateComponents_Date(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSDate* result_ = [nSDateComponents date];
    return result_;
}

int C_NSDateComponents_Era(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents era];
    return result_;
}

void C_NSDateComponents_SetEra(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setEra:value];
}

int C_NSDateComponents_Year(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents year];
    return result_;
}

void C_NSDateComponents_SetYear(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setYear:value];
}

int C_NSDateComponents_YearForWeekOfYear(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents yearForWeekOfYear];
    return result_;
}

void C_NSDateComponents_SetYearForWeekOfYear(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setYearForWeekOfYear:value];
}

int C_NSDateComponents_Quarter(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents quarter];
    return result_;
}

void C_NSDateComponents_SetQuarter(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setQuarter:value];
}

int C_NSDateComponents_Month(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents month];
    return result_;
}

void C_NSDateComponents_SetMonth(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setMonth:value];
}

bool C_NSDateComponents_IsLeapMonth(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    BOOL result_ = [nSDateComponents isLeapMonth];
    return result_;
}

void C_NSDateComponents_SetLeapMonth(void* ptr, bool value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setLeapMonth:value];
}

int C_NSDateComponents_Weekday(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents weekday];
    return result_;
}

void C_NSDateComponents_SetWeekday(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setWeekday:value];
}

int C_NSDateComponents_WeekdayOrdinal(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents weekdayOrdinal];
    return result_;
}

void C_NSDateComponents_SetWeekdayOrdinal(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setWeekdayOrdinal:value];
}

int C_NSDateComponents_WeekOfMonth(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents weekOfMonth];
    return result_;
}

void C_NSDateComponents_SetWeekOfMonth(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setWeekOfMonth:value];
}

int C_NSDateComponents_WeekOfYear(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents weekOfYear];
    return result_;
}

void C_NSDateComponents_SetWeekOfYear(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setWeekOfYear:value];
}

int C_NSDateComponents_Day(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents day];
    return result_;
}

void C_NSDateComponents_SetDay(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setDay:value];
}

int C_NSDateComponents_Hour(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents hour];
    return result_;
}

void C_NSDateComponents_SetHour(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setHour:value];
}

int C_NSDateComponents_Minute(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents minute];
    return result_;
}

void C_NSDateComponents_SetMinute(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setMinute:value];
}

int C_NSDateComponents_Second(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents second];
    return result_;
}

void C_NSDateComponents_SetSecond(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setSecond:value];
}

int C_NSDateComponents_Nanosecond(void* ptr) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    NSInteger result_ = [nSDateComponents nanosecond];
    return result_;
}

void C_NSDateComponents_SetNanosecond(void* ptr, int value) {
    NSDateComponents* nSDateComponents = (NSDateComponents*)ptr;
    [nSDateComponents setNanosecond:value];
}
