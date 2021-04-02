#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool DatePicker_IsBezeled(void* ptr);
void DatePicker_SetBezeled(void* ptr, bool bezeled);
bool DatePicker_IsBordered(void* ptr);
void DatePicker_SetBordered(void* ptr, bool bordered);
void* DatePicker_BackgroundColor(void* ptr);
void DatePicker_SetBackgroundColor(void* ptr, void* backgroundColor);
bool DatePicker_DrawsBackground(void* ptr);
void DatePicker_SetDrawsBackground(void* ptr, bool drawsBackground);
void* DatePicker_TextColor(void* ptr);
void DatePicker_SetTextColor(void* ptr, void* textColor);
unsigned long DatePicker_DatePickerStyle(void* ptr);
void DatePicker_SetDatePickerStyle(void* ptr, unsigned long datePickerStyle);
bool DatePicker_PresentsCalendarOverlay(void* ptr);
void DatePicker_SetPresentsCalendarOverlay(void* ptr, bool presentsCalendarOverlay);
unsigned long DatePicker_DatePickerElements(void* ptr);
void DatePicker_SetDatePickerElements(void* ptr, unsigned long datePickerElements);
unsigned long DatePicker_DatePickerMode(void* ptr);
void DatePicker_SetDatePickerMode(void* ptr, unsigned long datePickerMode);
void* DatePicker_DateValue(void* ptr);
void DatePicker_SetDateValue(void* ptr, void* dateValue);
void* DatePicker_MaxDate(void* ptr);
void DatePicker_SetMaxDate(void* ptr, void* maxDate);

void* DatePicker_NewDatePicker(NSRect frame);
