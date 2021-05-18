#import <Foundation/Foundation.h>
#import "time_zone.h"

void* C_TimeZone_Alloc() {
    return [NSTimeZone alloc];
}

void* C_NSTimeZone_InitWithName(void* ptr, void* tzName) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeZone* result_ = [nSTimeZone initWithName:(NSString*)tzName];
    return result_;
}

void* C_NSTimeZone_InitWithName_Data(void* ptr, void* tzName, Array aData) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeZone* result_ = [nSTimeZone initWithName:(NSString*)tzName data:[[NSData alloc] initWithBytes:(Byte *)aData.data length:aData.len]];
    return result_;
}

void C_NSTimeZone_ResetSystemTimeZone() {
    [NSTimeZone resetSystemTimeZone];
}

void* C_NSTimeZone_TimeZoneWithName(void* tzName) {
    NSTimeZone* result_ = [NSTimeZone timeZoneWithName:(NSString*)tzName];
    return result_;
}

void* C_NSTimeZone_TimeZoneWithName_Data(void* tzName, Array aData) {
    NSTimeZone* result_ = [NSTimeZone timeZoneWithName:(NSString*)tzName data:[[NSData alloc] initWithBytes:(Byte *)aData.data length:aData.len]];
    return result_;
}

void* C_NSTimeZone_TimeZoneWithAbbreviation(void* abbreviation) {
    NSTimeZone* result_ = [NSTimeZone timeZoneWithAbbreviation:(NSString*)abbreviation];
    return result_;
}

void* C_NSTimeZone_TimeZoneForSecondsFromGMT(int seconds) {
    NSTimeZone* result_ = [NSTimeZone timeZoneForSecondsFromGMT:seconds];
    return result_;
}

void* C_NSTimeZone_AbbreviationForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result_ = [nSTimeZone abbreviationForDate:(NSDate*)aDate];
    return result_;
}

int C_NSTimeZone_SecondsFromGMTForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSInteger result_ = [nSTimeZone secondsFromGMTForDate:(NSDate*)aDate];
    return result_;
}

bool C_NSTimeZone_IsDaylightSavingTimeForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result_ = [nSTimeZone isDaylightSavingTimeForDate:(NSDate*)aDate];
    return result_;
}

double C_NSTimeZone_DaylightSavingTimeOffsetForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeInterval result_ = [nSTimeZone daylightSavingTimeOffsetForDate:(NSDate*)aDate];
    return result_;
}

void* C_NSTimeZone_NextDaylightSavingTimeTransitionAfterDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSDate* result_ = [nSTimeZone nextDaylightSavingTimeTransitionAfterDate:(NSDate*)aDate];
    return result_;
}

bool C_NSTimeZone_IsEqualToTimeZone(void* ptr, void* aTimeZone) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result_ = [nSTimeZone isEqualToTimeZone:(NSTimeZone*)aTimeZone];
    return result_;
}

void* C_NSTimeZone_LocalizedName_Locale(void* ptr, int style, void* locale) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result_ = [nSTimeZone localizedName:style locale:(NSLocale*)locale];
    return result_;
}

void* C_NSTimeZone_LocalTimeZone() {
    NSTimeZone* result_ = [NSTimeZone localTimeZone];
    return result_;
}

void* C_NSTimeZone_SystemTimeZone() {
    NSTimeZone* result_ = [NSTimeZone systemTimeZone];
    return result_;
}

void* C_NSTimeZone_DefaultTimeZone() {
    NSTimeZone* result_ = [NSTimeZone defaultTimeZone];
    return result_;
}

void C_NSTimeZone_SetDefaultTimeZone(void* value) {
    [NSTimeZone setDefaultTimeZone:(NSTimeZone*)value];
}

Array C_NSTimeZone_TimeZone_KnownTimeZoneNames() {
    NSArray* result_ = [NSTimeZone knownTimeZoneNames];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

void* C_NSTimeZone_Name(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result_ = [nSTimeZone name];
    return result_;
}

void* C_NSTimeZone_Abbreviation(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result_ = [nSTimeZone abbreviation];
    return result_;
}

int C_NSTimeZone_SecondsFromGMT(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSInteger result_ = [nSTimeZone secondsFromGMT];
    return result_;
}

Array C_NSTimeZone_Data(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSData* result_ = [nSTimeZone data];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSTimeZone_TimeZoneDataVersion() {
    NSString* result_ = [NSTimeZone timeZoneDataVersion];
    return result_;
}

bool C_NSTimeZone_IsDaylightSavingTime(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result_ = [nSTimeZone isDaylightSavingTime];
    return result_;
}

double C_NSTimeZone_DaylightSavingTimeOffset(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeInterval result_ = [nSTimeZone daylightSavingTimeOffset];
    return result_;
}

void* C_NSTimeZone_NextDaylightSavingTimeTransition(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSDate* result_ = [nSTimeZone nextDaylightSavingTimeTransition];
    return result_;
}

void* C_NSTimeZone_Description(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result_ = [nSTimeZone description];
    return result_;
}
