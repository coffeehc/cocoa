#import "date_picker.h"
#import <AppKit/NSDatePicker.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_DatePicker_Alloc() {
    return [NSDatePicker alloc];
}

void* C_NSDatePicker_InitWithFrame(void* ptr, CGRect frameRect) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result_ = [nSDatePicker initWithFrame:frameRect];
    return result_;
}

void* C_NSDatePicker_InitWithCoder(void* ptr, void* coder) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result_ = [nSDatePicker initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSDatePicker_Init(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result_ = [nSDatePicker init];
    return result_;
}

void* C_NSDatePicker_AllocDatePicker() {
    NSDatePicker* result_ = [NSDatePicker alloc];
    return result_;
}

void* C_NSDatePicker_NewDatePicker() {
    NSDatePicker* result_ = [NSDatePicker new];
    return result_;
}

void* C_NSDatePicker_Autorelease(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result_ = [nSDatePicker autorelease];
    return result_;
}

void* C_NSDatePicker_Retain(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result_ = [nSDatePicker retain];
    return result_;
}

bool C_NSDatePicker_IsBezeled(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result_ = [nSDatePicker isBezeled];
    return result_;
}

void C_NSDatePicker_SetBezeled(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBezeled:value];
}

bool C_NSDatePicker_IsBordered(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result_ = [nSDatePicker isBordered];
    return result_;
}

void C_NSDatePicker_SetBordered(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBordered:value];
}

void* C_NSDatePicker_BackgroundColor(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSColor* result_ = [nSDatePicker backgroundColor];
    return result_;
}

void C_NSDatePicker_SetBackgroundColor(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBackgroundColor:(NSColor*)value];
}

bool C_NSDatePicker_DrawsBackground(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result_ = [nSDatePicker drawsBackground];
    return result_;
}

void C_NSDatePicker_SetDrawsBackground(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDrawsBackground:value];
}

void* C_NSDatePicker_TextColor(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSColor* result_ = [nSDatePicker textColor];
    return result_;
}

void C_NSDatePicker_SetTextColor(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTextColor:(NSColor*)value];
}

unsigned int C_NSDatePicker_DatePickerStyle(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerStyle result_ = [nSDatePicker datePickerStyle];
    return result_;
}

void C_NSDatePicker_SetDatePickerStyle(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerStyle:value];
}

bool C_NSDatePicker_PresentsCalendarOverlay(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result_ = [nSDatePicker presentsCalendarOverlay];
    return result_;
}

void C_NSDatePicker_SetPresentsCalendarOverlay(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setPresentsCalendarOverlay:value];
}

void* C_NSDatePicker_Delegate(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    id result_ = [nSDatePicker delegate];
    return result_;
}

void C_NSDatePicker_SetDelegate(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDelegate:(id)value];
}

unsigned int C_NSDatePicker_DatePickerElements(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerElementFlags result_ = [nSDatePicker datePickerElements];
    return result_;
}

void C_NSDatePicker_SetDatePickerElements(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerElements:value];
}

void* C_NSDatePicker_Calendar(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSCalendar* result_ = [nSDatePicker calendar];
    return result_;
}

void C_NSDatePicker_SetCalendar(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setCalendar:(NSCalendar*)value];
}

void* C_NSDatePicker_Locale(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSLocale* result_ = [nSDatePicker locale];
    return result_;
}

void C_NSDatePicker_SetLocale(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setLocale:(NSLocale*)value];
}

unsigned int C_NSDatePicker_DatePickerMode(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerMode result_ = [nSDatePicker datePickerMode];
    return result_;
}

void C_NSDatePicker_SetDatePickerMode(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerMode:value];
}

void* C_NSDatePicker_TimeZone(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSTimeZone* result_ = [nSDatePicker timeZone];
    return result_;
}

void C_NSDatePicker_SetTimeZone(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTimeZone:(NSTimeZone*)value];
}

void* C_NSDatePicker_DateValue(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result_ = [nSDatePicker dateValue];
    return result_;
}

void C_NSDatePicker_SetDateValue(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDateValue:(NSDate*)value];
}

double C_NSDatePicker_TimeInterval(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSTimeInterval result_ = [nSDatePicker timeInterval];
    return result_;
}

void C_NSDatePicker_SetTimeInterval(void* ptr, double value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTimeInterval:value];
}

void* C_NSDatePicker_MinDate(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result_ = [nSDatePicker minDate];
    return result_;
}

void C_NSDatePicker_SetMinDate(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setMinDate:(NSDate*)value];
}

void* C_NSDatePicker_MaxDate(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result_ = [nSDatePicker maxDate];
    return result_;
}

void C_NSDatePicker_SetMaxDate(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setMaxDate:(NSDate*)value];
}
