#import <Foundation/Foundation.h>
#import "calendar.h"

void* C_Calendar_Alloc() {
    return [NSCalendar alloc];
}

void* C_NSCalendar_CalendarWithIdentifier(void* calendarIdentifierConstant) {
    NSCalendar* result_ = [NSCalendar calendarWithIdentifier:(NSString*)calendarIdentifierConstant];
    return result_;
}

void* C_NSCalendar_InitWithCalendarIdentifier(void* ptr, void* ident) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    id result_ = [nSCalendar initWithCalendarIdentifier:(NSString*)ident];
    return result_;
}

bool C_NSCalendar_Date_MatchesComponents(void* ptr, void* date, void* components) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar date:(NSDate*)date matchesComponents:(NSDateComponents*)components];
    return result_;
}

int C_NSCalendar_Component_FromDate(void* ptr, unsigned int unit, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSInteger result_ = [nSCalendar component:unit fromDate:(NSDate*)date];
    return result_;
}

void* C_NSCalendar_Components_FromDate(void* ptr, unsigned int unitFlags, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result_ = [nSCalendar components:unitFlags fromDate:(NSDate*)date];
    return result_;
}

void* C_NSCalendar_Components_FromDate_ToDate_Options(void* ptr, unsigned int unitFlags, void* startingDate, void* resultDate, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result_ = [nSCalendar components:unitFlags fromDate:(NSDate*)startingDate toDate:(NSDate*)resultDate options:opts];
    return result_;
}

void* C_NSCalendar_Components_FromDateComponents_ToDateComponents_Options(void* ptr, unsigned int unitFlags, void* startingDateComp, void* resultDateComp, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result_ = [nSCalendar components:unitFlags fromDateComponents:(NSDateComponents*)startingDateComp toDateComponents:(NSDateComponents*)resultDateComp options:options];
    return result_;
}

void* C_NSCalendar_ComponentsInTimeZone_FromDate(void* ptr, void* timezone, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDateComponents* result_ = [nSCalendar componentsInTimeZone:(NSTimeZone*)timezone fromDate:(NSDate*)date];
    return result_;
}

NSRange C_NSCalendar_MaximumRangeOfUnit(void* ptr, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result_ = [nSCalendar maximumRangeOfUnit:unit];
    return result_;
}

NSRange C_NSCalendar_MinimumRangeOfUnit(void* ptr, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result_ = [nSCalendar minimumRangeOfUnit:unit];
    return result_;
}

unsigned int C_NSCalendar_OrdinalityOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result_ = [nSCalendar ordinalityOfUnit:smaller inUnit:larger forDate:(NSDate*)date];
    return result_;
}

NSRange C_NSCalendar_RangeOfUnit_InUnit_ForDate(void* ptr, unsigned int smaller, unsigned int larger, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSRange result_ = [nSCalendar rangeOfUnit:smaller inUnit:larger forDate:(NSDate*)date];
    return result_;
}

void* C_NSCalendar_StartOfDayForDate(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar startOfDayForDate:(NSDate*)date];
    return result_;
}

void* C_NSCalendar_NextDateAfterDate_MatchingComponents_Options(void* ptr, void* date, void* comps, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar nextDateAfterDate:(NSDate*)date matchingComponents:(NSDateComponents*)comps options:options];
    return result_;
}

void* C_NSCalendar_NextDateAfterDate_MatchingHour_Minute_Second_Options(void* ptr, void* date, int hourValue, int minuteValue, int secondValue, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar nextDateAfterDate:(NSDate*)date matchingHour:hourValue minute:minuteValue second:secondValue options:options];
    return result_;
}

void* C_NSCalendar_NextDateAfterDate_MatchingUnit_Value_Options(void* ptr, void* date, unsigned int unit, int value, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar nextDateAfterDate:(NSDate*)date matchingUnit:unit value:value options:options];
    return result_;
}

void* C_NSCalendar_DateFromComponents(void* ptr, void* comps) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateFromComponents:(NSDateComponents*)comps];
    return result_;
}

void* C_NSCalendar_DateByAddingComponents_ToDate_Options(void* ptr, void* comps, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateByAddingComponents:(NSDateComponents*)comps toDate:(NSDate*)date options:opts];
    return result_;
}

void* C_NSCalendar_DateByAddingUnit_Value_ToDate_Options(void* ptr, unsigned int unit, int value, void* date, unsigned int options) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateByAddingUnit:unit value:value toDate:(NSDate*)date options:options];
    return result_;
}

void* C_NSCalendar_DateBySettingHour_Minute_Second_OfDate_Options(void* ptr, int h, int m, int s, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateBySettingHour:h minute:m second:s ofDate:(NSDate*)date options:opts];
    return result_;
}

void* C_NSCalendar_DateBySettingUnit_Value_OfDate_Options(void* ptr, unsigned int unit, int v, void* date, unsigned int opts) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateBySettingUnit:unit value:v ofDate:(NSDate*)date options:opts];
    return result_;
}

void* C_NSCalendar_DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int monthValue, int dayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateWithEra:eraValue year:yearValue month:monthValue day:dayValue hour:hourValue minute:minuteValue second:secondValue nanosecond:nanosecondValue];
    return result_;
}

void* C_NSCalendar_DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(void* ptr, int eraValue, int yearValue, int weekValue, int weekdayValue, int hourValue, int minuteValue, int secondValue, int nanosecondValue) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSDate* result_ = [nSCalendar dateWithEra:eraValue yearForWeekOfYear:yearValue weekOfYear:weekValue weekday:weekdayValue hour:hourValue minute:minuteValue second:secondValue nanosecond:nanosecondValue];
    return result_;
}

int C_NSCalendar_CompareDate_ToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSComparisonResult result_ = [nSCalendar compareDate:(NSDate*)date1 toDate:(NSDate*)date2 toUnitGranularity:unit];
    return result_;
}

bool C_NSCalendar_IsDate_EqualToDate_ToUnitGranularity(void* ptr, void* date1, void* date2, unsigned int unit) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDate:(NSDate*)date1 equalToDate:(NSDate*)date2 toUnitGranularity:unit];
    return result_;
}

bool C_NSCalendar_IsDate_InSameDayAsDate(void* ptr, void* date1, void* date2) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDate:(NSDate*)date1 inSameDayAsDate:(NSDate*)date2];
    return result_;
}

bool C_NSCalendar_IsDateInToday(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDateInToday:(NSDate*)date];
    return result_;
}

bool C_NSCalendar_IsDateInTomorrow(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDateInTomorrow:(NSDate*)date];
    return result_;
}

bool C_NSCalendar_IsDateInWeekend(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDateInWeekend:(NSDate*)date];
    return result_;
}

bool C_NSCalendar_IsDateInYesterday(void* ptr, void* date) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    BOOL result_ = [nSCalendar isDateInYesterday:(NSDate*)date];
    return result_;
}

void* C_NSCalendar_CurrentCalendar() {
    NSCalendar* result_ = [NSCalendar currentCalendar];
    return result_;
}

void* C_NSCalendar_AutoupdatingCurrentCalendar() {
    NSCalendar* result_ = [NSCalendar autoupdatingCurrentCalendar];
    return result_;
}

void* C_NSCalendar_CalendarIdentifier(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSCalendarIdentifier result_ = [nSCalendar calendarIdentifier];
    return result_;
}

unsigned int C_NSCalendar_FirstWeekday(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result_ = [nSCalendar firstWeekday];
    return result_;
}

void C_NSCalendar_SetFirstWeekday(void* ptr, unsigned int value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setFirstWeekday:value];
}

void* C_NSCalendar_Locale(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSLocale* result_ = [nSCalendar locale];
    return result_;
}

void C_NSCalendar_SetLocale(void* ptr, void* value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setLocale:(NSLocale*)value];
}

void* C_NSCalendar_TimeZone(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSTimeZone* result_ = [nSCalendar timeZone];
    return result_;
}

void C_NSCalendar_SetTimeZone(void* ptr, void* value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setTimeZone:(NSTimeZone*)value];
}

unsigned int C_NSCalendar_MinimumDaysInFirstWeek(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSUInteger result_ = [nSCalendar minimumDaysInFirstWeek];
    return result_;
}

void C_NSCalendar_SetMinimumDaysInFirstWeek(void* ptr, unsigned int value) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    [nSCalendar setMinimumDaysInFirstWeek:value];
}

void* C_NSCalendar_AMSymbol(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSString* result_ = [nSCalendar AMSymbol];
    return result_;
}

void* C_NSCalendar_PMSymbol(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSString* result_ = [nSCalendar PMSymbol];
    return result_;
}

Array C_NSCalendar_WeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar weekdaySymbols];
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

Array C_NSCalendar_ShortWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortWeekdaySymbols];
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

Array C_NSCalendar_VeryShortWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar veryShortWeekdaySymbols];
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

Array C_NSCalendar_StandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar standaloneWeekdaySymbols];
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

Array C_NSCalendar_ShortStandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortStandaloneWeekdaySymbols];
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

Array C_NSCalendar_VeryShortStandaloneWeekdaySymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar veryShortStandaloneWeekdaySymbols];
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

Array C_NSCalendar_MonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar monthSymbols];
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

Array C_NSCalendar_ShortMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortMonthSymbols];
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

Array C_NSCalendar_VeryShortMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar veryShortMonthSymbols];
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

Array C_NSCalendar_StandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar standaloneMonthSymbols];
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

Array C_NSCalendar_ShortStandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortStandaloneMonthSymbols];
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

Array C_NSCalendar_VeryShortStandaloneMonthSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar veryShortStandaloneMonthSymbols];
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

Array C_NSCalendar_QuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar quarterSymbols];
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

Array C_NSCalendar_ShortQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortQuarterSymbols];
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

Array C_NSCalendar_StandaloneQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar standaloneQuarterSymbols];
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

Array C_NSCalendar_ShortStandaloneQuarterSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar shortStandaloneQuarterSymbols];
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

Array C_NSCalendar_EraSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar eraSymbols];
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

Array C_NSCalendar_LongEraSymbols(void* ptr) {
    NSCalendar* nSCalendar = (NSCalendar*)ptr;
    NSArray* result_ = [nSCalendar longEraSymbols];
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
