#import <Foundation/Foundation.h>
#import "calendar.h"

void* C_Calendar_Alloc() {
    return [NSCalendar alloc];
}

void* C_NSCalendar_CalendarWithIdentifier(void* calendarIdentifierConstant) {
    NSCalendar* result = [NSCalendar calendarWithIdentifier:(NSString*)calendarIdentifierConstant];
    return result;
}

void* C_NSCalendar_InitWithCalendarIdentifier(void* ptr, void* ident) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    id result = [nSCalendar initWithCalendarIdentifier:(NSString*)ident];
    return result;
}

bool C_NSCalendar_Date_MatchesComponents(void* ptr, void* date, void* components) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar date:(NSDate*)date matchesComponents:(NSDateComponents*)components];
    return result;
}

int C_NSCalendar_Component_FromDate(void* ptr, unsigned int unit, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSInteger result = [nSCalendar component:unit fromDate:(NSDate*)date];
    return result;
}

void* C_NSCalendar_Components_FromDate(void* ptr, unsigned int unitFlags, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result = [nSCalendar components:unitFlags fromDate:(NSDate*)date];
    return result;
}

void* C_NSCalendar_Components_FromDate_ToDate_Options(void* ptr, unsigned int unitFlags, void* startingDate, void* resultDate, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result = [nSCalendar components:unitFlags fromDate:(NSDate*)startingDate toDate:(NSDate*)resultDate options:opts];
    return result;
}

void* C_NSCalendar_Components_FromDateComponents_ToDateComponents_Options(void* ptr, unsigned int unitFlags, void* startingDateComp, void* resultDateComp, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result = [nSCalendar components:unitFlags fromDateComponents:(NSDateComponents*)startingDateComp toDateComponents:(NSDateComponents*)resultDateComp options:options];
    return result;
}

void* C_NSCalendar_ComponentsInTimeZone_FromDate(void* ptr, void* timezone, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result = [nSCalendar componentsInTimeZone:(NSTimeZone*)timezone fromDate:(NSDate*)date];
    return result;
}

NSRange C_NSCalendar_MaximumRangeOfUnit(void* ptr, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result = [nSCalendar maximumRangeOfUnit:unit];
    return result;
}

NSRange C_NSCalendar_MinimumRangeOfUnit(void* ptr, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result = [nSCalendar minimumRangeOfUnit:unit];
    return result;
}

unsigned int C_NSCalendar_OrdinalityOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result = [nSCalendar ordinalityOfUnit:smaller inUnit:larger forDate:(NSDate*)date];
    return result;
}

NSRange C_NSCalendar_RangeOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result = [nSCalendar rangeOfUnit:smaller inUnit:larger forDate:(NSDate*)date];
    return result;
}

void* C_NSCalendar_StartOfDayForDate(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar startOfDayForDate:(NSDate*)date];
    return result;
}

void* C_NSCalendar_NextDateAfterDate_MatchingComponents_Options(void* ptr, void* date, void* comps, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar nextDateAfterDate:(NSDate*)date matchingComponents:(NSDateComponents*)comps options:options];
    return result;
}

void* C_NSCalendar_NextDateAfterDate_MatchingHour_Minute_Second_Options(void* ptr, void* date, int hourValue, int minuteValue, int secondValue, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar nextDateAfterDate:(NSDate*)date matchingHour:hourValue minute:minuteValue second:secondValue options:options];
    return result;
}

void* C_NSCalendar_NextDateAfterDate_MatchingUnit_Value_Options(void* ptr, void* date, unsigned int unit, int value, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar nextDateAfterDate:(NSDate*)date matchingUnit:unit value:value options:options];
    return result;
}

void* C_NSCalendar_DateFromComponents(void* ptr, void* comps) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateFromComponents:(NSDateComponents*)comps];
    return result;
}

void* C_NSCalendar_DateByAddingComponents_ToDate_Options(void* ptr, void* comps, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateByAddingComponents:(NSDateComponents*)comps toDate:(NSDate*)date options:opts];
    return result;
}

void* C_NSCalendar_DateByAddingUnit_Value_ToDate_Options(void* ptr, unsigned int unit, int value, void* date, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateByAddingUnit:unit value:value toDate:(NSDate*)date options:options];
    return result;
}

void* C_NSCalendar_DateBySettingHour_Minute_Second_OfDate_Options(void* ptr, int h, int m, int s, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateBySettingHour:h minute:m second:s ofDate:(NSDate*)date options:opts];
    return result;
}

void* C_NSCalendar_DateBySettingUnit_Value_OfDate_Options(void* ptr, unsigned int unit, int v, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateBySettingUnit:unit value:v ofDate:(NSDate*)date options:opts];
    return result;
}

void* C_NSCalendar_DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int monthValue, int dayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateWithEra:eraValue year:yearValue month:monthValue day:dayValue hour:hourValue minute:minuteValue second:secondValue nanosecond:nanosecondValue];
    return result;
}

void* C_NSCalendar_DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int weekValue, int weekdayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result = [nSCalendar dateWithEra:eraValue yearForWeekOfYear:yearValue weekOfYear:weekValue weekday:weekdayValue hour:hourValue minute:minuteValue second:secondValue nanosecond:nanosecondValue];
    return result;
}

int C_NSCalendar_CompareDate_ToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSComparisonResult result = [nSCalendar compareDate:(NSDate*)date1 toDate:(NSDate*)date2 toUnitGranularity:unit];
    return result;
}

bool C_NSCalendar_IsDate_EqualToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDate:(NSDate*)date1 equalToDate:(NSDate*)date2 toUnitGranularity:unit];
    return result;
}

bool C_NSCalendar_IsDate_InSameDayAsDate(void* ptr, void* date1, void* date2) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDate:(NSDate*)date1 inSameDayAsDate:(NSDate*)date2];
    return result;
}

bool C_NSCalendar_IsDateInToday(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDateInToday:(NSDate*)date];
    return result;
}

bool C_NSCalendar_IsDateInTomorrow(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDateInTomorrow:(NSDate*)date];
    return result;
}

bool C_NSCalendar_IsDateInWeekend(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDateInWeekend:(NSDate*)date];
    return result;
}

bool C_NSCalendar_IsDateInYesterday(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result = [nSCalendar isDateInYesterday:(NSDate*)date];
    return result;
}

void* C_NSCalendar_CurrentCalendar() {
    NSCalendar* result = [NSCalendar currentCalendar];
    return result;
}

void* C_NSCalendar_AutoupdatingCurrentCalendar() {
    NSCalendar* result = [NSCalendar autoupdatingCurrentCalendar];
    return result;
}

void* C_NSCalendar_CalendarIdentifier(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSCalendarIdentifier result = [nSCalendar calendarIdentifier];
    return result;
}

unsigned int C_NSCalendar_FirstWeekday(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result = [nSCalendar firstWeekday];
    return result;
}

void C_NSCalendar_SetFirstWeekday(void* ptr, unsigned int value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setFirstWeekday:value];
}

void* C_NSCalendar_Locale(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSLocale* result = [nSCalendar locale];
    return result;
}

void C_NSCalendar_SetLocale(void* ptr, void* value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setLocale:(NSLocale*)value];
}

void* C_NSCalendar_TimeZone(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSTimeZone* result = [nSCalendar timeZone];
    return result;
}

void C_NSCalendar_SetTimeZone(void* ptr, void* value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setTimeZone:(NSTimeZone*)value];
}

unsigned int C_NSCalendar_MinimumDaysInFirstWeek(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result = [nSCalendar minimumDaysInFirstWeek];
    return result;
}

void C_NSCalendar_SetMinimumDaysInFirstWeek(void* ptr, unsigned int value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setMinimumDaysInFirstWeek:value];
}

void* C_NSCalendar_AMSymbol(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSString* result = [nSCalendar AMSymbol];
    return result;
}

void* C_NSCalendar_PMSymbol(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSString* result = [nSCalendar PMSymbol];
    return result;
}

Array C_NSCalendar_WeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar weekdaySymbols];
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

Array C_NSCalendar_ShortWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortWeekdaySymbols];
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

Array C_NSCalendar_VeryShortWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar veryShortWeekdaySymbols];
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

Array C_NSCalendar_StandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar standaloneWeekdaySymbols];
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

Array C_NSCalendar_ShortStandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortStandaloneWeekdaySymbols];
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

Array C_NSCalendar_VeryShortStandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar veryShortStandaloneWeekdaySymbols];
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

Array C_NSCalendar_MonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar monthSymbols];
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

Array C_NSCalendar_ShortMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortMonthSymbols];
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

Array C_NSCalendar_VeryShortMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar veryShortMonthSymbols];
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

Array C_NSCalendar_StandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar standaloneMonthSymbols];
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

Array C_NSCalendar_ShortStandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortStandaloneMonthSymbols];
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

Array C_NSCalendar_VeryShortStandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar veryShortStandaloneMonthSymbols];
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

Array C_NSCalendar_QuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar quarterSymbols];
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

Array C_NSCalendar_ShortQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortQuarterSymbols];
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

Array C_NSCalendar_StandaloneQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar standaloneQuarterSymbols];
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

Array C_NSCalendar_ShortStandaloneQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar shortStandaloneQuarterSymbols];
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

Array C_NSCalendar_EraSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar eraSymbols];
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

Array C_NSCalendar_LongEraSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result = [nSCalendar longEraSymbols];
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
