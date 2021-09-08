#import <Foundation/Foundation.h>
#import "date_interval_formatter.h"

void* C_DateIntervalFormatter_Alloc() {
    return [NSDateIntervalFormatter alloc];
}

void* C_NSDateIntervalFormatter_AllocDateIntervalFormatter() {
    NSDateIntervalFormatter* result_ = [NSDateIntervalFormatter alloc];
    return result_;
}

void* C_NSDateIntervalFormatter_Init(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSDateIntervalFormatter* result_ = [nSDateIntervalFormatter init];
    return result_;
}

void* C_NSDateIntervalFormatter_NewDateIntervalFormatter() {
    NSDateIntervalFormatter* result_ = [NSDateIntervalFormatter new];
    return result_;
}

void* C_NSDateIntervalFormatter_Autorelease(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSDateIntervalFormatter* result_ = [nSDateIntervalFormatter autorelease];
    return result_;
}

void* C_NSDateIntervalFormatter_Retain(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSDateIntervalFormatter* result_ = [nSDateIntervalFormatter retain];
    return result_;
}

void* C_NSDateIntervalFormatter_StringFromDate_ToDate(void* ptr, void* fromDate, void* toDate) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSString* result_ = [nSDateIntervalFormatter stringFromDate:(NSDate*)fromDate toDate:(NSDate*)toDate];
    return result_;
}

void* C_NSDateIntervalFormatter_StringFromDateInterval(void* ptr, void* dateInterval) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSString* result_ = [nSDateIntervalFormatter stringFromDateInterval:(NSDateInterval*)dateInterval];
    return result_;
}

unsigned int C_NSDateIntervalFormatter_DateStyle(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSDateIntervalFormatterStyle result_ = [nSDateIntervalFormatter dateStyle];
    return result_;
}

void C_NSDateIntervalFormatter_SetDateStyle(void* ptr, unsigned int value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setDateStyle:value];
}

unsigned int C_NSDateIntervalFormatter_TimeStyle(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSDateIntervalFormatterStyle result_ = [nSDateIntervalFormatter timeStyle];
    return result_;
}

void C_NSDateIntervalFormatter_SetTimeStyle(void* ptr, unsigned int value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setTimeStyle:value];
}

void* C_NSDateIntervalFormatter_DateTemplate(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSString* result_ = [nSDateIntervalFormatter dateTemplate];
    return result_;
}

void C_NSDateIntervalFormatter_SetDateTemplate(void* ptr, void* value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setDateTemplate:(NSString*)value];
}

void* C_NSDateIntervalFormatter_Calendar(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSCalendar* result_ = [nSDateIntervalFormatter calendar];
    return result_;
}

void C_NSDateIntervalFormatter_SetCalendar(void* ptr, void* value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setCalendar:(NSCalendar*)value];
}

void* C_NSDateIntervalFormatter_Locale(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSLocale* result_ = [nSDateIntervalFormatter locale];
    return result_;
}

void C_NSDateIntervalFormatter_SetLocale(void* ptr, void* value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setLocale:(NSLocale*)value];
}

void* C_NSDateIntervalFormatter_TimeZone(void* ptr) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    NSTimeZone* result_ = [nSDateIntervalFormatter timeZone];
    return result_;
}

void C_NSDateIntervalFormatter_SetTimeZone(void* ptr, void* value) {
    NSDateIntervalFormatter* nSDateIntervalFormatter = (NSDateIntervalFormatter*)ptr;
    [nSDateIntervalFormatter setTimeZone:(NSTimeZone*)value];
}
