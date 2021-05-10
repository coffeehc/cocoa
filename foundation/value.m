#import <Foundation/Foundation.h>
#import "url.h"

void* C_Value_Alloc() {
	return [NSValue alloc];
}

void* C_NSValue_InitWithCoder(void* ptr, void* coder) {
	NSValue* nSValue = (NSValue*)ptr;
	NSValue* result = [nSValue initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSValue_Init(void* ptr) {
	NSValue* nSValue = (NSValue*)ptr;
	NSValue* result = [nSValue init];
	return result;
}

void* C_NSValue_ValueWithNonretainedObject(void* anObject) {
	NSValue* result = [NSValue valueWithNonretainedObject:(id)anObject];
	return result;
}

void* C_NSValue_ValueWithRange(NSRange _range) {
	NSValue* result = [NSValue valueWithRange:_range];
	return result;
}

bool C_NSValue_IsEqualToValue(void* ptr, void* value) {
	NSValue* nSValue = (NSValue*)ptr;
	bool result = [nSValue isEqualToValue:(NSValue*)value];
	return result;
}

void* C_NSValue_ValueWithEdgeInsets(NSEdgeInsets insets) {
	NSValue* result = [NSValue valueWithEdgeInsets:insets];
	return result;
}

void* C_NSValue_NonretainedObjectValue(void* ptr) {
	NSValue* nSValue = (NSValue*)ptr;
	id result = [nSValue nonretainedObjectValue];
	return result;
}

NSRange C_NSValue_RangeValue(void* ptr) {
	NSValue* nSValue = (NSValue*)ptr;
	NSRange result = [nSValue rangeValue];
	return result;
}

NSEdgeInsets C_NSValue_EdgeInsetsValue(void* ptr) {
	NSValue* nSValue = (NSValue*)ptr;
	NSEdgeInsets result = [nSValue edgeInsetsValue];
	return result;
}
