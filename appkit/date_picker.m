#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "date_picker.h"

bool DatePicker_IsBezeled(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker isBezeled];
}

void DatePicker_SetBezeled(void* ptr, bool bezeled) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setBezeled:bezeled];
}

bool DatePicker_IsBordered(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker isBordered];
}

void DatePicker_SetBordered(void* ptr, bool bordered) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setBordered:bordered];
}

void* DatePicker_BackgroundColor(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker backgroundColor];
}

void DatePicker_SetBackgroundColor(void* ptr, void* backgroundColor) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setBackgroundColor:(NSColor*)backgroundColor];
}

bool DatePicker_DrawsBackground(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker drawsBackground];
}

void DatePicker_SetDrawsBackground(void* ptr, bool drawsBackground) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setDrawsBackground:drawsBackground];
}

void* DatePicker_TextColor(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker textColor];
}

void DatePicker_SetTextColor(void* ptr, void* textColor) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setTextColor:(NSColor*)textColor];
}

unsigned long DatePicker_DatePickerStyle(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker datePickerStyle];
}

void DatePicker_SetDatePickerStyle(void* ptr, unsigned long datePickerStyle) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setDatePickerStyle:datePickerStyle];
}

bool DatePicker_PresentsCalendarOverlay(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker presentsCalendarOverlay];
}

void DatePicker_SetPresentsCalendarOverlay(void* ptr, bool presentsCalendarOverlay) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setPresentsCalendarOverlay:presentsCalendarOverlay];
}

unsigned long DatePicker_DatePickerElements(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker datePickerElements];
}

void DatePicker_SetDatePickerElements(void* ptr, unsigned long datePickerElements) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setDatePickerElements:datePickerElements];
}

unsigned long DatePicker_DatePickerMode(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker datePickerMode];
}

void DatePicker_SetDatePickerMode(void* ptr, unsigned long datePickerMode) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setDatePickerMode:datePickerMode];
}

void* DatePicker_DateValue(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker dateValue];
}

void DatePicker_SetDateValue(void* ptr, void* dateValue) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setDateValue:(NSDate*)dateValue];
}

void* DatePicker_MaxDate(void* ptr) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	return [datePicker maxDate];
}

void DatePicker_SetMaxDate(void* ptr, void* maxDate) {
	NSDatePicker* datePicker = (NSDatePicker*)ptr;
	[datePicker setMaxDate:(NSDate*)maxDate];
}

void* DatePicker_NewDatePicker(NSRect frame) {
	NSDatePicker* datePicker = [NSDatePicker alloc];
	return [[datePicker initWithFrame:frame] autorelease];
}
