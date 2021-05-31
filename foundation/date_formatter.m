#import <Foundation/Foundation.h>
#import "date_formatter.h"

void* C_DateFormatter_Alloc() {
    return [NSDateFormatter alloc];
}

void* C_NSDateFormatter_Init(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDateFormatter* result_ = [nSDateFormatter init];
    return result_;
}

void* C_NSDateFormatter_DateFromString(void* ptr, void* _string) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDate* result_ = [nSDateFormatter dateFromString:(NSString*)_string];
    return result_;
}

void* C_NSDateFormatter_StringFromDate(void* ptr, void* date) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSString* result_ = [nSDateFormatter stringFromDate:(NSDate*)date];
    return result_;
}

void* C_NSDateFormatter_DateFormatter_LocalizedStringFromDate_DateStyle_TimeStyle(void* date, unsigned int dstyle, unsigned int tstyle) {
    NSString* result_ = [NSDateFormatter localizedStringFromDate:(NSDate*)date dateStyle:dstyle timeStyle:tstyle];
    return result_;
}

void C_NSDateFormatter_SetLocalizedDateFormatFromTemplate(void* ptr, void* dateFormatTemplate) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setLocalizedDateFormatFromTemplate:(NSString*)dateFormatTemplate];
}

void* C_NSDateFormatter_DateFormatter_DateFormatFromTemplate_Options_Locale(void* tmplate, unsigned int opts, void* locale) {
    NSString* result_ = [NSDateFormatter dateFormatFromTemplate:(NSString*)tmplate options:opts locale:(NSLocale*)locale];
    return result_;
}

unsigned int C_NSDateFormatter_DateStyle(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDateFormatterStyle result_ = [nSDateFormatter dateStyle];
    return result_;
}

void C_NSDateFormatter_SetDateStyle(void* ptr, unsigned int value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setDateStyle:value];
}

unsigned int C_NSDateFormatter_TimeStyle(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDateFormatterStyle result_ = [nSDateFormatter timeStyle];
    return result_;
}

void C_NSDateFormatter_SetTimeStyle(void* ptr, unsigned int value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setTimeStyle:value];
}

void* C_NSDateFormatter_DateFormat(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSString* result_ = [nSDateFormatter dateFormat];
    return result_;
}

void C_NSDateFormatter_SetDateFormat(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setDateFormat:(NSString*)value];
}

int C_NSDateFormatter_FormattingContext(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSFormattingContext result_ = [nSDateFormatter formattingContext];
    return result_;
}

void C_NSDateFormatter_SetFormattingContext(void* ptr, int value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setFormattingContext:value];
}

void* C_NSDateFormatter_Calendar(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSCalendar* result_ = [nSDateFormatter calendar];
    return result_;
}

void C_NSDateFormatter_SetCalendar(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setCalendar:(NSCalendar*)value];
}

void* C_NSDateFormatter_DefaultDate(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDate* result_ = [nSDateFormatter defaultDate];
    return result_;
}

void C_NSDateFormatter_SetDefaultDate(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setDefaultDate:(NSDate*)value];
}

void* C_NSDateFormatter_Locale(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSLocale* result_ = [nSDateFormatter locale];
    return result_;
}

void C_NSDateFormatter_SetLocale(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setLocale:(NSLocale*)value];
}

void* C_NSDateFormatter_TimeZone(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSTimeZone* result_ = [nSDateFormatter timeZone];
    return result_;
}

void C_NSDateFormatter_SetTimeZone(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setTimeZone:(NSTimeZone*)value];
}

void* C_NSDateFormatter_TwoDigitStartDate(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDate* result_ = [nSDateFormatter twoDigitStartDate];
    return result_;
}

void C_NSDateFormatter_SetTwoDigitStartDate(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setTwoDigitStartDate:(NSDate*)value];
}

void* C_NSDateFormatter_GregorianStartDate(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDate* result_ = [nSDateFormatter gregorianStartDate];
    return result_;
}

void C_NSDateFormatter_SetGregorianStartDate(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setGregorianStartDate:(NSDate*)value];
}

unsigned int C_NSDateFormatter_FormatterBehavior(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSDateFormatterBehavior result_ = [nSDateFormatter formatterBehavior];
    return result_;
}

void C_NSDateFormatter_SetFormatterBehavior(void* ptr, unsigned int value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setFormatterBehavior:value];
}

unsigned int C_NSDateFormatter_DateFormatter_DefaultFormatterBehavior() {
    NSDateFormatterBehavior result_ = [NSDateFormatter defaultFormatterBehavior];
    return result_;
}

void C_NSDateFormatter_DateFormatter_SetDefaultFormatterBehavior(unsigned int value) {
    [NSDateFormatter setDefaultFormatterBehavior:value];
}

bool C_NSDateFormatter_IsLenient(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    BOOL result_ = [nSDateFormatter isLenient];
    return result_;
}

void C_NSDateFormatter_SetLenient(void* ptr, bool value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setLenient:value];
}

bool C_NSDateFormatter_DoesRelativeDateFormatting(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    BOOL result_ = [nSDateFormatter doesRelativeDateFormatting];
    return result_;
}

void C_NSDateFormatter_SetDoesRelativeDateFormatting(void* ptr, bool value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setDoesRelativeDateFormatting:value];
}

void* C_NSDateFormatter_AMSymbol(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSString* result_ = [nSDateFormatter AMSymbol];
    return result_;
}

void C_NSDateFormatter_SetAMSymbol(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setAMSymbol:(NSString*)value];
}

void* C_NSDateFormatter_PMSymbol(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSString* result_ = [nSDateFormatter PMSymbol];
    return result_;
}

void C_NSDateFormatter_SetPMSymbol(void* ptr, void* value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setPMSymbol:(NSString*)value];
}

Array C_NSDateFormatter_WeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter weekdaySymbols];
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

void C_NSDateFormatter_SetWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_ShortWeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortWeekdaySymbols];
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

void C_NSDateFormatter_SetShortWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_VeryShortWeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter veryShortWeekdaySymbols];
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

void C_NSDateFormatter_SetVeryShortWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setVeryShortWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_StandaloneWeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter standaloneWeekdaySymbols];
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

void C_NSDateFormatter_SetStandaloneWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setStandaloneWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_ShortStandaloneWeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortStandaloneWeekdaySymbols];
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

void C_NSDateFormatter_SetShortStandaloneWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortStandaloneWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_VeryShortStandaloneWeekdaySymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter veryShortStandaloneWeekdaySymbols];
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

void C_NSDateFormatter_SetVeryShortStandaloneWeekdaySymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setVeryShortStandaloneWeekdaySymbols:objcValue];
}

Array C_NSDateFormatter_MonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter monthSymbols];
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

void C_NSDateFormatter_SetMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setMonthSymbols:objcValue];
}

Array C_NSDateFormatter_ShortMonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortMonthSymbols];
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

void C_NSDateFormatter_SetShortMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortMonthSymbols:objcValue];
}

Array C_NSDateFormatter_VeryShortMonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter veryShortMonthSymbols];
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

void C_NSDateFormatter_SetVeryShortMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setVeryShortMonthSymbols:objcValue];
}

Array C_NSDateFormatter_StandaloneMonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter standaloneMonthSymbols];
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

void C_NSDateFormatter_SetStandaloneMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setStandaloneMonthSymbols:objcValue];
}

Array C_NSDateFormatter_ShortStandaloneMonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortStandaloneMonthSymbols];
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

void C_NSDateFormatter_SetShortStandaloneMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortStandaloneMonthSymbols:objcValue];
}

Array C_NSDateFormatter_VeryShortStandaloneMonthSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter veryShortStandaloneMonthSymbols];
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

void C_NSDateFormatter_SetVeryShortStandaloneMonthSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setVeryShortStandaloneMonthSymbols:objcValue];
}

Array C_NSDateFormatter_QuarterSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter quarterSymbols];
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

void C_NSDateFormatter_SetQuarterSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setQuarterSymbols:objcValue];
}

Array C_NSDateFormatter_ShortQuarterSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortQuarterSymbols];
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

void C_NSDateFormatter_SetShortQuarterSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortQuarterSymbols:objcValue];
}

Array C_NSDateFormatter_StandaloneQuarterSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter standaloneQuarterSymbols];
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

void C_NSDateFormatter_SetStandaloneQuarterSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setStandaloneQuarterSymbols:objcValue];
}

Array C_NSDateFormatter_ShortStandaloneQuarterSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter shortStandaloneQuarterSymbols];
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

void C_NSDateFormatter_SetShortStandaloneQuarterSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setShortStandaloneQuarterSymbols:objcValue];
}

Array C_NSDateFormatter_EraSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter eraSymbols];
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

void C_NSDateFormatter_SetEraSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setEraSymbols:objcValue];
}

Array C_NSDateFormatter_LongEraSymbols(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSArray* result_ = [nSDateFormatter longEraSymbols];
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

void C_NSDateFormatter_SetLongEraSymbols(void* ptr, Array value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSDateFormatter setLongEraSymbols:objcValue];
}

bool C_NSDateFormatter_GeneratesCalendarDates(void* ptr) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    BOOL result_ = [nSDateFormatter generatesCalendarDates];
    return result_;
}

void C_NSDateFormatter_SetGeneratesCalendarDates(void* ptr, bool value) {
    NSDateFormatter* nSDateFormatter = (NSDateFormatter*)ptr;
    [nSDateFormatter setGeneratesCalendarDates:value];
}
