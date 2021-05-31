#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_DateComponentsFormatter_Alloc();

void* C_NSDateComponentsFormatter_Init(void* ptr);
void* C_NSDateComponentsFormatter_StringFromDateComponents(void* ptr, void* components);
void* C_NSDateComponentsFormatter_StringFromDate_ToDate(void* ptr, void* startDate, void* endDate);
void* C_NSDateComponentsFormatter_StringFromTimeInterval(void* ptr, double ti);
void* C_NSDateComponentsFormatter_DateComponentsFormatter_LocalizedStringFromDateComponents_UnitsStyle(void* components, int unitsStyle);
unsigned int C_NSDateComponentsFormatter_AllowedUnits(void* ptr);
void C_NSDateComponentsFormatter_SetAllowedUnits(void* ptr, unsigned int value);
bool C_NSDateComponentsFormatter_AllowsFractionalUnits(void* ptr);
void C_NSDateComponentsFormatter_SetAllowsFractionalUnits(void* ptr, bool value);
void* C_NSDateComponentsFormatter_Calendar(void* ptr);
void C_NSDateComponentsFormatter_SetCalendar(void* ptr, void* value);
bool C_NSDateComponentsFormatter_CollapsesLargestUnit(void* ptr);
void C_NSDateComponentsFormatter_SetCollapsesLargestUnit(void* ptr, bool value);
bool C_NSDateComponentsFormatter_IncludesApproximationPhrase(void* ptr);
void C_NSDateComponentsFormatter_SetIncludesApproximationPhrase(void* ptr, bool value);
bool C_NSDateComponentsFormatter_IncludesTimeRemainingPhrase(void* ptr);
void C_NSDateComponentsFormatter_SetIncludesTimeRemainingPhrase(void* ptr, bool value);
int C_NSDateComponentsFormatter_MaximumUnitCount(void* ptr);
void C_NSDateComponentsFormatter_SetMaximumUnitCount(void* ptr, int value);
int C_NSDateComponentsFormatter_UnitsStyle(void* ptr);
void C_NSDateComponentsFormatter_SetUnitsStyle(void* ptr, int value);
unsigned int C_NSDateComponentsFormatter_ZeroFormattingBehavior(void* ptr);
void C_NSDateComponentsFormatter_SetZeroFormattingBehavior(void* ptr, unsigned int value);
int C_NSDateComponentsFormatter_FormattingContext(void* ptr);
void C_NSDateComponentsFormatter_SetFormattingContext(void* ptr, int value);
void* C_NSDateComponentsFormatter_ReferenceDate(void* ptr);
void C_NSDateComponentsFormatter_SetReferenceDate(void* ptr, void* value);
