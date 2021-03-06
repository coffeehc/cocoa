#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DateInterval_Alloc();

void* C_NSDateInterval_Init(void* ptr);
void* C_NSDateInterval_InitWithStartDate_Duration(void* ptr, void* startDate, double duration);
void* C_NSDateInterval_InitWithStartDate_EndDate(void* ptr, void* startDate, void* endDate);
void* C_NSDateInterval_InitWithCoder(void* ptr, void* coder);
void* C_NSDateInterval_AllocDateInterval();
void* C_NSDateInterval_NewDateInterval();
void* C_NSDateInterval_Autorelease(void* ptr);
void* C_NSDateInterval_Retain(void* ptr);
int C_NSDateInterval_Compare(void* ptr, void* dateInterval);
bool C_NSDateInterval_IsEqualToDateInterval(void* ptr, void* dateInterval);
bool C_NSDateInterval_IntersectsDateInterval(void* ptr, void* dateInterval);
void* C_NSDateInterval_IntersectionWithDateInterval(void* ptr, void* dateInterval);
bool C_NSDateInterval_ContainsDate(void* ptr, void* date);
void* C_NSDateInterval_StartDate(void* ptr);
void* C_NSDateInterval_EndDate(void* ptr);
double C_NSDateInterval_Duration(void* ptr);
