#import <Appkit/Appkit.h>
#import "date_picker.h"

void* C_DatePicker_Alloc() {
    return [NSDatePicker alloc];
}

void* C_NSDatePicker_InitWithFrame(void* ptr, CGRect frameRect) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result = [nSDatePicker initWithFrame:frameRect];
    return result;
}

void* C_NSDatePicker_InitWithCoder(void* ptr, void* coder) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result = [nSDatePicker initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSDatePicker_Init(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePicker* result = [nSDatePicker init];
    return result;
}

bool C_NSDatePicker_IsBezeled(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result = [nSDatePicker isBezeled];
    return result;
}

void C_NSDatePicker_SetBezeled(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBezeled:value];
}

bool C_NSDatePicker_IsBordered(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result = [nSDatePicker isBordered];
    return result;
}

void C_NSDatePicker_SetBordered(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBordered:value];
}

void* C_NSDatePicker_BackgroundColor(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSColor* result = [nSDatePicker backgroundColor];
    return result;
}

void C_NSDatePicker_SetBackgroundColor(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setBackgroundColor:(NSColor*)value];
}

bool C_NSDatePicker_DrawsBackground(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result = [nSDatePicker drawsBackground];
    return result;
}

void C_NSDatePicker_SetDrawsBackground(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDrawsBackground:value];
}

void* C_NSDatePicker_TextColor(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSColor* result = [nSDatePicker textColor];
    return result;
}

void C_NSDatePicker_SetTextColor(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTextColor:(NSColor*)value];
}

unsigned int C_NSDatePicker_DatePickerStyle(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerStyle result = [nSDatePicker datePickerStyle];
    return result;
}

void C_NSDatePicker_SetDatePickerStyle(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerStyle:value];
}

bool C_NSDatePicker_PresentsCalendarOverlay(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    BOOL result = [nSDatePicker presentsCalendarOverlay];
    return result;
}

void C_NSDatePicker_SetPresentsCalendarOverlay(void* ptr, bool value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setPresentsCalendarOverlay:value];
}

unsigned int C_NSDatePicker_DatePickerElements(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerElementFlags result = [nSDatePicker datePickerElements];
    return result;
}

void C_NSDatePicker_SetDatePickerElements(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerElements:value];
}

void* C_NSDatePicker_Calendar(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSCalendar* result = [nSDatePicker calendar];
    return result;
}

void C_NSDatePicker_SetCalendar(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setCalendar:(NSCalendar*)value];
}

void* C_NSDatePicker_Locale(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSLocale* result = [nSDatePicker locale];
    return result;
}

void C_NSDatePicker_SetLocale(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setLocale:(NSLocale*)value];
}

unsigned int C_NSDatePicker_DatePickerMode(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDatePickerMode result = [nSDatePicker datePickerMode];
    return result;
}

void C_NSDatePicker_SetDatePickerMode(void* ptr, unsigned int value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDatePickerMode:value];
}

void* C_NSDatePicker_TimeZone(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSTimeZone* result = [nSDatePicker timeZone];
    return result;
}

void C_NSDatePicker_SetTimeZone(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTimeZone:(NSTimeZone*)value];
}

void* C_NSDatePicker_DateValue(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result = [nSDatePicker dateValue];
    return result;
}

void C_NSDatePicker_SetDateValue(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setDateValue:(NSDate*)value];
}

double C_NSDatePicker_TimeInterval(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSTimeInterval result = [nSDatePicker timeInterval];
    return result;
}

void C_NSDatePicker_SetTimeInterval(void* ptr, double value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setTimeInterval:value];
}

void* C_NSDatePicker_MinDate(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result = [nSDatePicker minDate];
    return result;
}

void C_NSDatePicker_SetMinDate(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setMinDate:(NSDate*)value];
}

void* C_NSDatePicker_MaxDate(void* ptr) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    NSDate* result = [nSDatePicker maxDate];
    return result;
}

void C_NSDatePicker_SetMaxDate(void* ptr, void* value) {
    NSDatePicker* nSDatePicker = (NSDatePicker*)ptr;
    [nSDatePicker setMaxDate:(NSDate*)value];
}
