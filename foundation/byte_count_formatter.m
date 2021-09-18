#import "byte_count_formatter.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSByteCountFormatter.h>

void* C_ByteCountFormatter_Alloc() {
    return [NSByteCountFormatter alloc];
}

void* C_NSByteCountFormatter_AllocByteCountFormatter() {
    NSByteCountFormatter* result_ = [NSByteCountFormatter alloc];
    return result_;
}

void* C_NSByteCountFormatter_Init(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSByteCountFormatter* result_ = [nSByteCountFormatter init];
    return result_;
}

void* C_NSByteCountFormatter_NewByteCountFormatter() {
    NSByteCountFormatter* result_ = [NSByteCountFormatter new];
    return result_;
}

void* C_NSByteCountFormatter_Autorelease(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSByteCountFormatter* result_ = [nSByteCountFormatter autorelease];
    return result_;
}

void* C_NSByteCountFormatter_Retain(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSByteCountFormatter* result_ = [nSByteCountFormatter retain];
    return result_;
}

void* C_NSByteCountFormatter_StringFromMeasurement(void* ptr, void* measurement) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSString* result_ = [nSByteCountFormatter stringFromMeasurement:(NSMeasurement*)measurement];
    return result_;
}

void* C_NSByteCountFormatter_ByteCountFormatter_StringFromMeasurement_CountStyle(void* measurement, int countStyle) {
    NSString* result_ = [NSByteCountFormatter stringFromMeasurement:(NSMeasurement*)measurement countStyle:countStyle];
    return result_;
}

int C_NSByteCountFormatter_FormattingContext(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSFormattingContext result_ = [nSByteCountFormatter formattingContext];
    return result_;
}

void C_NSByteCountFormatter_SetFormattingContext(void* ptr, int value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setFormattingContext:value];
}

int C_NSByteCountFormatter_CountStyle(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSByteCountFormatterCountStyle result_ = [nSByteCountFormatter countStyle];
    return result_;
}

void C_NSByteCountFormatter_SetCountStyle(void* ptr, int value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setCountStyle:value];
}

bool C_NSByteCountFormatter_AllowsNonnumericFormatting(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter allowsNonnumericFormatting];
    return result_;
}

void C_NSByteCountFormatter_SetAllowsNonnumericFormatting(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setAllowsNonnumericFormatting:value];
}

bool C_NSByteCountFormatter_IncludesActualByteCount(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter includesActualByteCount];
    return result_;
}

void C_NSByteCountFormatter_SetIncludesActualByteCount(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setIncludesActualByteCount:value];
}

bool C_NSByteCountFormatter_IsAdaptive(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter isAdaptive];
    return result_;
}

void C_NSByteCountFormatter_SetAdaptive(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setAdaptive:value];
}

unsigned int C_NSByteCountFormatter_AllowedUnits(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    NSByteCountFormatterUnits result_ = [nSByteCountFormatter allowedUnits];
    return result_;
}

void C_NSByteCountFormatter_SetAllowedUnits(void* ptr, unsigned int value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setAllowedUnits:value];
}

bool C_NSByteCountFormatter_IncludesCount(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter includesCount];
    return result_;
}

void C_NSByteCountFormatter_SetIncludesCount(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setIncludesCount:value];
}

bool C_NSByteCountFormatter_IncludesUnit(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter includesUnit];
    return result_;
}

void C_NSByteCountFormatter_SetIncludesUnit(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setIncludesUnit:value];
}

bool C_NSByteCountFormatter_ZeroPadsFractionDigits(void* ptr) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    BOOL result_ = [nSByteCountFormatter zeroPadsFractionDigits];
    return result_;
}

void C_NSByteCountFormatter_SetZeroPadsFractionDigits(void* ptr, bool value) {
    NSByteCountFormatter* nSByteCountFormatter = (NSByteCountFormatter*)ptr;
    [nSByteCountFormatter setZeroPadsFractionDigits:value];
}
