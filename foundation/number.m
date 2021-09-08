#import <Foundation/Foundation.h>
#import "number.h"

void* C_Number_Alloc() {
    return [NSNumber alloc];
}

void* C_NSNumber_InitWithCoder(void* ptr, void* coder) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSNumber_AllocNumber() {
    NSNumber* result_ = [NSNumber alloc];
    return result_;
}

void* C_NSNumber_Init(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber init];
    return result_;
}

void* C_NSNumber_NewNumber() {
    NSNumber* result_ = [NSNumber new];
    return result_;
}

void* C_NSNumber_Autorelease(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber autorelease];
    return result_;
}

void* C_NSNumber_Retain(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber retain];
    return result_;
}

void* C_NSNumber_NumberWithBool(bool value) {
    NSNumber* result_ = [NSNumber numberWithBool:value];
    return result_;
}

void* C_NSNumber_NumberWithDouble(double value) {
    NSNumber* result_ = [NSNumber numberWithDouble:value];
    return result_;
}

void* C_NSNumber_NumberWithFloat(float value) {
    NSNumber* result_ = [NSNumber numberWithFloat:value];
    return result_;
}

void* C_NSNumber_NumberWithInteger(int value) {
    NSNumber* result_ = [NSNumber numberWithInteger:value];
    return result_;
}

void* C_NSNumber_NumberWithUnsignedInteger(unsigned int value) {
    NSNumber* result_ = [NSNumber numberWithUnsignedInteger:value];
    return result_;
}

void* C_NSNumber_InitWithBool(void* ptr, bool value) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithBool:value];
    return result_;
}

void* C_NSNumber_InitWithDouble(void* ptr, double value) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithDouble:value];
    return result_;
}

void* C_NSNumber_InitWithFloat(void* ptr, float value) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithFloat:value];
    return result_;
}

void* C_NSNumber_InitWithInteger(void* ptr, int value) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithInteger:value];
    return result_;
}

void* C_NSNumber_InitWithUnsignedInteger(void* ptr, unsigned int value) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSNumber* result_ = [nSNumber initWithUnsignedInteger:value];
    return result_;
}

void* C_NSNumber_DescriptionWithLocale(void* ptr, void* locale) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSString* result_ = [nSNumber descriptionWithLocale:(id)locale];
    return result_;
}

int C_NSNumber_Compare(void* ptr, void* otherNumber) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSComparisonResult result_ = [nSNumber compare:(NSNumber*)otherNumber];
    return result_;
}

bool C_NSNumber_IsEqualToNumber(void* ptr, void* number) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    BOOL result_ = [nSNumber isEqualToNumber:(NSNumber*)number];
    return result_;
}

bool C_NSNumber_BoolValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    BOOL result_ = [nSNumber boolValue];
    return result_;
}

double C_NSNumber_DoubleValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    double result_ = [nSNumber doubleValue];
    return result_;
}

float C_NSNumber_FloatValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    float result_ = [nSNumber floatValue];
    return result_;
}

int C_NSNumber_IntegerValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSInteger result_ = [nSNumber integerValue];
    return result_;
}

unsigned int C_NSNumber_UnsignedIntegerValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSUInteger result_ = [nSNumber unsignedIntegerValue];
    return result_;
}

void* C_NSNumber_StringValue(void* ptr) {
    NSNumber* nSNumber = (NSNumber*)ptr;
    NSString* result_ = [nSNumber stringValue];
    return result_;
}
