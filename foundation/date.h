#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* Date_DistantFuture();
void* Date_DistantPast();
double Date_TimeIntervalSinceNow(void* ptr);
double Date_TimeIntervalSinceReferenceDate(void* ptr);
double Date_TimeIntervalSince1970(void* ptr);
const char* Date_Description(void* ptr);

void* Date_NewDate();
void* Date_CurrentDate();
void* Date_DateWithTimeIntervalSinceNow(double secs);
void* Date_dateWithTimeIntervalSinceDate(double secsToBeAdded, void* date);
void* Date_DateWithTimeIntervalSinceReferenceDate(double secs);
void* Date_DateWithTimeIntervalSince1970(double secs);
const char* Date_DescriptionWithLocale(void* ptr, void* locale);
bool Date_IsEqualToDate(void* ptr, void* otherDate);
void* Date_EarlierDate(void* ptr, void* otherDate);
void* Date_LaterDate(void* ptr, void* otherDate);
unsigned long Date_Compare(void* ptr, void* other);
double Date_TimeIntervalSinceDate(void* ptr, void* other);
