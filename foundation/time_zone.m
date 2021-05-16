#import <Foundation/Foundation.h>
#import "time_zone.h"

void* C_TimeZone_Alloc() {
    return [NSTimeZone alloc];
}

void* C_NSTimeZone_InitWithName(void* ptr, void* tzName) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeZone* result = [nSTimeZone initWithName:(NSString*)tzName];
    return result;
}

void* C_NSTimeZone_InitWithName_Data(void* ptr, void* tzName, Array aData) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeZone* result = [nSTimeZone initWithName:(NSString*)tzName data:[[NSData alloc] initWithBytes:(Byte *)aData.data length:aData.len]];
    return result;
}

void C_NSTimeZone_ResetSystemTimeZone() {
    [NSTimeZone resetSystemTimeZone];
}

void* C_NSTimeZone_TimeZoneWithName(void* tzName) {
    NSTimeZone* result = [NSTimeZone timeZoneWithName:(NSString*)tzName];
    return result;
}

void* C_NSTimeZone_TimeZoneWithName_Data(void* tzName, Array aData) {
    NSTimeZone* result = [NSTimeZone timeZoneWithName:(NSString*)tzName data:[[NSData alloc] initWithBytes:(Byte *)aData.data length:aData.len]];
    return result;
}

void* C_NSTimeZone_TimeZoneWithAbbreviation(void* abbreviation) {
    NSTimeZone* result = [NSTimeZone timeZoneWithAbbreviation:(NSString*)abbreviation];
    return result;
}

void* C_NSTimeZone_TimeZoneForSecondsFromGMT(int seconds) {
    NSTimeZone* result = [NSTimeZone timeZoneForSecondsFromGMT:seconds];
    return result;
}

void* C_NSTimeZone_AbbreviationForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result = [nSTimeZone abbreviationForDate:(NSDate*)aDate];
    return result;
}

int C_NSTimeZone_SecondsFromGMTForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSInteger result = [nSTimeZone secondsFromGMTForDate:(NSDate*)aDate];
    return result;
}

bool C_NSTimeZone_IsDaylightSavingTimeForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result = [nSTimeZone isDaylightSavingTimeForDate:(NSDate*)aDate];
    return result;
}

double C_NSTimeZone_DaylightSavingTimeOffsetForDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeInterval result = [nSTimeZone daylightSavingTimeOffsetForDate:(NSDate*)aDate];
    return result;
}

void* C_NSTimeZone_NextDaylightSavingTimeTransitionAfterDate(void* ptr, void* aDate) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSDate* result = [nSTimeZone nextDaylightSavingTimeTransitionAfterDate:(NSDate*)aDate];
    return result;
}

bool C_NSTimeZone_IsEqualToTimeZone(void* ptr, void* aTimeZone) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result = [nSTimeZone isEqualToTimeZone:(NSTimeZone*)aTimeZone];
    return result;
}

void* C_NSTimeZone_LocalizedName_Locale(void* ptr, int style, void* locale) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result = [nSTimeZone localizedName:style locale:(NSLocale*)locale];
    return result;
}

void* C_NSTimeZone_LocalTimeZone() {
    NSTimeZone* result = [NSTimeZone localTimeZone];
    return result;
}

void* C_NSTimeZone_SystemTimeZone() {
    NSTimeZone* result = [NSTimeZone systemTimeZone];
    return result;
}

void* C_NSTimeZone_DefaultTimeZone() {
    NSTimeZone* result = [NSTimeZone defaultTimeZone];
    return result;
}

void C_NSTimeZone_SetDefaultTimeZone(void* value) {
    [NSTimeZone setDefaultTimeZone:(NSTimeZone*)value];
}

Array C_NSTimeZone_TimeZone_KnownTimeZoneNames() {
    NSArray* result = [NSTimeZone knownTimeZoneNames];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void* C_NSTimeZone_Name(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result = [nSTimeZone name];
    return result;
}

void* C_NSTimeZone_Abbreviation(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result = [nSTimeZone abbreviation];
    return result;
}

int C_NSTimeZone_SecondsFromGMT(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSInteger result = [nSTimeZone secondsFromGMT];
    return result;
}

Array C_NSTimeZone_Data(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSData* result = [nSTimeZone data];
    Array resultarray;
    resultarray.data = [result bytes];
    resultarray.len = result.length;
    return resultarray;
}

void* C_NSTimeZone_TimeZoneDataVersion() {
    NSString* result = [NSTimeZone timeZoneDataVersion];
    return result;
}

bool C_NSTimeZone_IsDaylightSavingTime(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    BOOL result = [nSTimeZone isDaylightSavingTime];
    return result;
}

double C_NSTimeZone_DaylightSavingTimeOffset(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSTimeInterval result = [nSTimeZone daylightSavingTimeOffset];
    return result;
}

void* C_NSTimeZone_NextDaylightSavingTimeTransition(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSDate* result = [nSTimeZone nextDaylightSavingTimeTransition];
    return result;
}

void* C_NSTimeZone_Description(void* ptr) {
    NSTimeZone* nSTimeZone = (NSTimeZone*)ptr;
    NSString* result = [nSTimeZone description];
    return result;
}
