#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Date_Alloc();

void* C_NSDate_Date_();
void* C_NSDate_DateWithTimeIntervalSinceNow(double secs);
void* C_NSDate_DateWithTimeInterval_SinceDate(double secsToBeAdded, void* date);
void* C_NSDate_DateWithTimeIntervalSinceReferenceDate(double ti);
void* C_NSDate_DateWithTimeIntervalSince1970(double secs);
void* C_NSDate_Init(void* ptr);
void* C_NSDate_InitWithTimeIntervalSinceNow(void* ptr, double secs);
void* C_NSDate_InitWithTimeInterval_SinceDate(void* ptr, double secsToBeAdded, void* date);
void* C_NSDate_InitWithTimeIntervalSinceReferenceDate(void* ptr, double ti);
void* C_NSDate_InitWithTimeIntervalSince1970(void* ptr, double secs);
void* C_NSDate_InitWithCoder(void* ptr, void* coder);
void* C_NSDate_DateByAddingTimeInterval(void* ptr, double ti);
void* C_NSDate_AllocDate();
void* C_NSDate_NewDate();
void* C_NSDate_Autorelease(void* ptr);
void* C_NSDate_Retain(void* ptr);
bool C_NSDate_IsEqualToDate(void* ptr, void* otherDate);
void* C_NSDate_EarlierDate(void* ptr, void* anotherDate);
void* C_NSDate_LaterDate(void* ptr, void* anotherDate);
int C_NSDate_Compare(void* ptr, void* other);
double C_NSDate_TimeIntervalSinceDate(void* ptr, void* anotherDate);
void* C_NSDate_DescriptionWithLocale(void* ptr, void* locale);
void* C_NSDate_Date_DistantFuture();
void* C_NSDate_Date_DistantPast();
void* C_NSDate_Date_Now();
double C_NSDate_TimeIntervalSinceNow(void* ptr);
double C_NSDate_TimeIntervalSinceReferenceDate(void* ptr);
double C_NSDate_TimeIntervalSince1970(void* ptr);
void* C_NSDate_Description(void* ptr);
