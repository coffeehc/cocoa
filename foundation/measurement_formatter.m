#import "measurement_formatter.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSMeasurementFormatter.h>

void* C_MeasurementFormatter_Alloc() {
    return [NSMeasurementFormatter alloc];
}

void* C_NSMeasurementFormatter_AllocMeasurementFormatter() {
    NSMeasurementFormatter* result_ = [NSMeasurementFormatter alloc];
    return result_;
}

void* C_NSMeasurementFormatter_Init(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSMeasurementFormatter* result_ = [nSMeasurementFormatter init];
    return result_;
}

void* C_NSMeasurementFormatter_NewMeasurementFormatter() {
    NSMeasurementFormatter* result_ = [NSMeasurementFormatter new];
    return result_;
}

void* C_NSMeasurementFormatter_Autorelease(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSMeasurementFormatter* result_ = [nSMeasurementFormatter autorelease];
    return result_;
}

void* C_NSMeasurementFormatter_Retain(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSMeasurementFormatter* result_ = [nSMeasurementFormatter retain];
    return result_;
}

void* C_NSMeasurementFormatter_StringFromMeasurement(void* ptr, void* measurement) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSString* result_ = [nSMeasurementFormatter stringFromMeasurement:(NSMeasurement*)measurement];
    return result_;
}

void* C_NSMeasurementFormatter_StringFromUnit(void* ptr, void* unit) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSString* result_ = [nSMeasurementFormatter stringFromUnit:(NSUnit*)unit];
    return result_;
}

unsigned int C_NSMeasurementFormatter_UnitOptions(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSMeasurementFormatterUnitOptions result_ = [nSMeasurementFormatter unitOptions];
    return result_;
}

void C_NSMeasurementFormatter_SetUnitOptions(void* ptr, unsigned int value) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    [nSMeasurementFormatter setUnitOptions:value];
}

int C_NSMeasurementFormatter_UnitStyle(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSFormattingUnitStyle result_ = [nSMeasurementFormatter unitStyle];
    return result_;
}

void C_NSMeasurementFormatter_SetUnitStyle(void* ptr, int value) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    [nSMeasurementFormatter setUnitStyle:value];
}

void* C_NSMeasurementFormatter_Locale(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSLocale* result_ = [nSMeasurementFormatter locale];
    return result_;
}

void C_NSMeasurementFormatter_SetLocale(void* ptr, void* value) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    [nSMeasurementFormatter setLocale:(NSLocale*)value];
}

void* C_NSMeasurementFormatter_NumberFormatter(void* ptr) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    NSNumberFormatter* result_ = [nSMeasurementFormatter numberFormatter];
    return result_;
}

void C_NSMeasurementFormatter_SetNumberFormatter(void* ptr, void* value) {
    NSMeasurementFormatter* nSMeasurementFormatter = (NSMeasurementFormatter*)ptr;
    [nSMeasurementFormatter setNumberFormatter:(NSNumberFormatter*)value];
}
