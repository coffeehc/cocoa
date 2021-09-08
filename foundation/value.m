#import <Foundation/Foundation.h>
#import "value.h"

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
