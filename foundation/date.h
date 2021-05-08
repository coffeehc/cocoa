#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* C_Date_Alloc();

void* C_NSDate_Init(void* ptr);
void* C_NSDate_InitWithTimeIntervalSinceNow(void* ptr, double secs);
void* C_NSDate_InitWithTimeInterval_SinceDate(void* ptr, double secsToBeAdded, void* date);
void* C_NSDate_InitWithTimeIntervalSinceReferenceDate(void* ptr, double ti);
void* C_NSDate_InitWithTimeIntervalSince1970(void* ptr, double secs);
void* C_NSDate_InitWithCoder(void* ptr, void* coder);
void* C_NSDate_DateByAddingTimeInterval(void* ptr, double ti);
void* C_NSDate_CurrentDate();
void* C_NSDate_DateWithTimeIntervalSinceNow(double secs);
void* C_NSDate_DateWithTimeInterval_SinceDate(double secsToBeAdded, void* date);
void* C_NSDate_DateWithTimeIntervalSinceReferenceDate(double ti);
void* C_NSDate_DateWithTimeIntervalSince1970(double secs);
bool C_NSDate_IsEqualToDate(void* ptr, void* otherDate);
void* C_NSDate_EarlierDate(void* ptr, void* anotherDate);
void* C_NSDate_LaterDate(void* ptr, void* anotherDate);
double C_NSDate_TimeIntervalSinceDate(void* ptr, void* anotherDate);
void* C_NSDate_DescriptionWithLocale(void* ptr, void* locale);
double C_NSDate_TimeIntervalSinceNow(void* ptr);
double C_NSDate_TimeIntervalSinceReferenceDate(void* ptr);
double C_NSDate_TimeIntervalSince1970(void* ptr);
void* C_NSDate_Description(void* ptr);
