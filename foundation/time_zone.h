#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_TimeZone_Alloc();

void* C_NSTimeZone_InitWithName(void* ptr, void* tzName);
void* C_NSTimeZone_InitWithName_Data(void* ptr, void* tzName, Array aData);
void C_NSTimeZone_ResetSystemTimeZone();
void* C_NSTimeZone_TimeZoneWithName(void* tzName);
void* C_NSTimeZone_TimeZoneWithName_Data(void* tzName, Array aData);
void* C_NSTimeZone_TimeZoneWithAbbreviation(void* abbreviation);
void* C_NSTimeZone_TimeZoneForSecondsFromGMT(int seconds);
void* C_NSTimeZone_AbbreviationForDate(void* ptr, void* aDate);
int C_NSTimeZone_SecondsFromGMTForDate(void* ptr, void* aDate);
bool C_NSTimeZone_IsDaylightSavingTimeForDate(void* ptr, void* aDate);
double C_NSTimeZone_DaylightSavingTimeOffsetForDate(void* ptr, void* aDate);
void* C_NSTimeZone_NextDaylightSavingTimeTransitionAfterDate(void* ptr, void* aDate);
bool C_NSTimeZone_IsEqualToTimeZone(void* ptr, void* aTimeZone);
void* C_NSTimeZone_LocalizedName_Locale(void* ptr, int style, void* locale);
void* C_NSTimeZone_LocalTimeZone();
void* C_NSTimeZone_SystemTimeZone();
void* C_NSTimeZone_DefaultTimeZone();
void C_NSTimeZone_SetDefaultTimeZone(void* value);
Array C_NSTimeZone_TimeZone_KnownTimeZoneNames();
void* C_NSTimeZone_Name(void* ptr);
void* C_NSTimeZone_Abbreviation(void* ptr);
int C_NSTimeZone_SecondsFromGMT(void* ptr);
Array C_NSTimeZone_Data(void* ptr);
void* C_NSTimeZone_TimeZoneDataVersion();
bool C_NSTimeZone_IsDaylightSavingTime(void* ptr);
double C_NSTimeZone_DaylightSavingTimeOffset(void* ptr);
void* C_NSTimeZone_NextDaylightSavingTimeTransition(void* ptr);
void* C_NSTimeZone_Description(void* ptr);
