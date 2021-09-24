#import "value.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSValue.h>

void* C_Value_Alloc() {
    return [NSValue alloc];
}

void* C_NSValue_InitWithCoder(void* ptr, void* coder) {
    NSValue* nSValue = (NSValue*)ptr;
    NSValue* result_ = [nSValue initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSValue_AllocValue() {
    NSValue* result_ = [NSValue alloc];
    return result_;
}

void* C_NSValue_Init(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSValue* result_ = [nSValue init];
    return result_;
}

void* C_NSValue_NewValue() {
    NSValue* result_ = [NSValue new];
    return result_;
}

void* C_NSValue_Autorelease(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSValue* result_ = [nSValue autorelease];
    return result_;
}

void* C_NSValue_Retain(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSValue* result_ = [nSValue retain];
    return result_;
}

void* C_NSValue_ValueWithNonretainedObject(void* anObject) {
    NSValue* result_ = [NSValue valueWithNonretainedObject:(id)anObject];
    return result_;
}

void* C_NSValue_ValueWithRange(NSRange _range) {
    NSValue* result_ = [NSValue valueWithRange:_range];
    return result_;
}

void* C_NSValue_ValueWithPoint(CGPoint point) {
    NSValue* result_ = [NSValue valueWithPoint:point];
    return result_;
}

void* C_NSValue_ValueWithSize(CGSize size) {
    NSValue* result_ = [NSValue valueWithSize:size];
    return result_;
}

void* C_NSValue_ValueWithRect(CGRect rect) {
    NSValue* result_ = [NSValue valueWithRect:rect];
    return result_;
}

bool C_NSValue_IsEqualToValue(void* ptr, void* value) {
    NSValue* nSValue = (NSValue*)ptr;
    BOOL result_ = [nSValue isEqualToValue:(NSValue*)value];
    return result_;
}

void* C_NSValue_ValueWithEdgeInsets(NSEdgeInsets insets) {
    NSValue* result_ = [NSValue valueWithEdgeInsets:insets];
    return result_;
}

void* C_NSValue_NonretainedObjectValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    id result_ = [nSValue nonretainedObjectValue];
    return result_;
}

NSRange C_NSValue_RangeValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSRange result_ = [nSValue rangeValue];
    return result_;
}

CGPoint C_NSValue_PointValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSPoint result_ = [nSValue pointValue];
    return result_;
}

CGSize C_NSValue_SizeValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSSize result_ = [nSValue sizeValue];
    return result_;
}

CGRect C_NSValue_RectValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSRect result_ = [nSValue rectValue];
    return result_;
}

NSEdgeInsets C_NSValue_EdgeInsetsValue(void* ptr) {
    NSValue* nSValue = (NSValue*)ptr;
    NSEdgeInsets result_ = [nSValue edgeInsetsValue];
    return result_;
}

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
