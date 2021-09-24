#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DatePicker_Alloc();

void* C_NSDatePicker_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSDatePicker_InitWithCoder(void* ptr, void* coder);
void* C_NSDatePicker_Init(void* ptr);
void* C_NSDatePicker_AllocDatePicker();
void* C_NSDatePicker_NewDatePicker();
void* C_NSDatePicker_Autorelease(void* ptr);
void* C_NSDatePicker_Retain(void* ptr);
bool C_NSDatePicker_IsBezeled(void* ptr);
void C_NSDatePicker_SetBezeled(void* ptr, bool value);
bool C_NSDatePicker_IsBordered(void* ptr);
void C_NSDatePicker_SetBordered(void* ptr, bool value);
void* C_NSDatePicker_BackgroundColor(void* ptr);
void C_NSDatePicker_SetBackgroundColor(void* ptr, void* value);
bool C_NSDatePicker_DrawsBackground(void* ptr);
void C_NSDatePicker_SetDrawsBackground(void* ptr, bool value);
void* C_NSDatePicker_TextColor(void* ptr);
void C_NSDatePicker_SetTextColor(void* ptr, void* value);
unsigned int C_NSDatePicker_DatePickerStyle(void* ptr);
void C_NSDatePicker_SetDatePickerStyle(void* ptr, unsigned int value);
bool C_NSDatePicker_PresentsCalendarOverlay(void* ptr);
void C_NSDatePicker_SetPresentsCalendarOverlay(void* ptr, bool value);
void* C_NSDatePicker_Delegate(void* ptr);
void C_NSDatePicker_SetDelegate(void* ptr, void* value);
unsigned int C_NSDatePicker_DatePickerElements(void* ptr);
void C_NSDatePicker_SetDatePickerElements(void* ptr, unsigned int value);
void* C_NSDatePicker_Calendar(void* ptr);
void C_NSDatePicker_SetCalendar(void* ptr, void* value);
void* C_NSDatePicker_Locale(void* ptr);
void C_NSDatePicker_SetLocale(void* ptr, void* value);
unsigned int C_NSDatePicker_DatePickerMode(void* ptr);
void C_NSDatePicker_SetDatePickerMode(void* ptr, unsigned int value);
void* C_NSDatePicker_TimeZone(void* ptr);
void C_NSDatePicker_SetTimeZone(void* ptr, void* value);
void* C_NSDatePicker_DateValue(void* ptr);
void C_NSDatePicker_SetDateValue(void* ptr, void* value);
double C_NSDatePicker_TimeInterval(void* ptr);
void C_NSDatePicker_SetTimeInterval(void* ptr, double value);
void* C_NSDatePicker_MinDate(void* ptr);
void C_NSDatePicker_SetMinDate(void* ptr, void* value);
void* C_NSDatePicker_MaxDate(void* ptr);
void C_NSDatePicker_SetMaxDate(void* ptr, void* value);
