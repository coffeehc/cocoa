#import <Foundation/Foundation.h>
#import "date_components_formatter.h"

void* C_DateComponentsFormatter_Alloc() {
    return [NSDateComponentsFormatter alloc];
}

void* C_NSDateComponentsFormatter_Init(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSDateComponentsFormatter* result_ = [nSDateComponentsFormatter init];
    return result_;
}

void* C_NSDateComponentsFormatter_StringFromDateComponents(void* ptr, void* components) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSString* result_ = [nSDateComponentsFormatter stringFromDateComponents:(NSDateComponents*)components];
    return result_;
}

void* C_NSDateComponentsFormatter_StringFromDate_ToDate(void* ptr, void* startDate, void* endDate) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSString* result_ = [nSDateComponentsFormatter stringFromDate:(NSDate*)startDate toDate:(NSDate*)endDate];
    return result_;
}

void* C_NSDateComponentsFormatter_StringFromTimeInterval(void* ptr, double ti) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSString* result_ = [nSDateComponentsFormatter stringFromTimeInterval:ti];
    return result_;
}

void* C_NSDateComponentsFormatter_DateComponentsFormatter_LocalizedStringFromDateComponents_UnitsStyle(void* components, int unitsStyle) {
    NSString* result_ = [NSDateComponentsFormatter localizedStringFromDateComponents:(NSDateComponents*)components unitsStyle:unitsStyle];
    return result_;
}

unsigned int C_NSDateComponentsFormatter_AllowedUnits(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSCalendarUnit result_ = [nSDateComponentsFormatter allowedUnits];
    return result_;
}

void C_NSDateComponentsFormatter_SetAllowedUnits(void* ptr, unsigned int value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setAllowedUnits:value];
}

bool C_NSDateComponentsFormatter_AllowsFractionalUnits(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    BOOL result_ = [nSDateComponentsFormatter allowsFractionalUnits];
    return result_;
}

void C_NSDateComponentsFormatter_SetAllowsFractionalUnits(void* ptr, bool value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setAllowsFractionalUnits:value];
}

void* C_NSDateComponentsFormatter_Calendar(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSCalendar* result_ = [nSDateComponentsFormatter calendar];
    return result_;
}

void C_NSDateComponentsFormatter_SetCalendar(void* ptr, void* value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setCalendar:(NSCalendar*)value];
}

bool C_NSDateComponentsFormatter_CollapsesLargestUnit(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    BOOL result_ = [nSDateComponentsFormatter collapsesLargestUnit];
    return result_;
}

void C_NSDateComponentsFormatter_SetCollapsesLargestUnit(void* ptr, bool value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setCollapsesLargestUnit:value];
}

bool C_NSDateComponentsFormatter_IncludesApproximationPhrase(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    BOOL result_ = [nSDateComponentsFormatter includesApproximationPhrase];
    return result_;
}

void C_NSDateComponentsFormatter_SetIncludesApproximationPhrase(void* ptr, bool value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setIncludesApproximationPhrase:value];
}

bool C_NSDateComponentsFormatter_IncludesTimeRemainingPhrase(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    BOOL result_ = [nSDateComponentsFormatter includesTimeRemainingPhrase];
    return result_;
}

void C_NSDateComponentsFormatter_SetIncludesTimeRemainingPhrase(void* ptr, bool value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setIncludesTimeRemainingPhrase:value];
}

int C_NSDateComponentsFormatter_MaximumUnitCount(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSInteger result_ = [nSDateComponentsFormatter maximumUnitCount];
    return result_;
}

void C_NSDateComponentsFormatter_SetMaximumUnitCount(void* ptr, int value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setMaximumUnitCount:value];
}

int C_NSDateComponentsFormatter_UnitsStyle(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSDateComponentsFormatterUnitsStyle result_ = [nSDateComponentsFormatter unitsStyle];
    return result_;
}

void C_NSDateComponentsFormatter_SetUnitsStyle(void* ptr, int value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setUnitsStyle:value];
}

unsigned int C_NSDateComponentsFormatter_ZeroFormattingBehavior(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSDateComponentsFormatterZeroFormattingBehavior result_ = [nSDateComponentsFormatter zeroFormattingBehavior];
    return result_;
}

void C_NSDateComponentsFormatter_SetZeroFormattingBehavior(void* ptr, unsigned int value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setZeroFormattingBehavior:value];
}

int C_NSDateComponentsFormatter_FormattingContext(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSFormattingContext result_ = [nSDateComponentsFormatter formattingContext];
    return result_;
}

void C_NSDateComponentsFormatter_SetFormattingContext(void* ptr, int value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setFormattingContext:value];
}

void* C_NSDateComponentsFormatter_ReferenceDate(void* ptr) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    NSDate* result_ = [nSDateComponentsFormatter referenceDate];
    return result_;
}

void C_NSDateComponentsFormatter_SetReferenceDate(void* ptr, void* value) {
    NSDateComponentsFormatter* nSDateComponentsFormatter = (NSDateComponentsFormatter*)ptr;
    [nSDateComponentsFormatter setReferenceDate:(NSDate*)value];
}
