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

bool C_NSValue_IsEqualToValue(void* ptr, void* value) {
	NSValue* nSValue = (NSValue*)ptr;
	bool result = [nSValue isEqualToValue:(NSValue*)value];
	return result;
}

void* C_NSValue_NonretainedObjectValue(void* ptr) {
	NSValue* nSValue = (NSValue*)ptr;
	id result = [nSValue nonretainedObjectValue];
	return result;
}
