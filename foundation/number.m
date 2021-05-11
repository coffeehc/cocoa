#import <Foundation/Foundation.h>
#import "number.h"

void* C_Number_Alloc() {
	return [NSNumber alloc];
}

void* C_NSNumber_InitWithCoder(void* ptr, void* coder) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSNumber_Init(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber init];
	return result;
}

void* C_NSNumber_NumberWithBool(bool value) {
	NSNumber* result = [NSNumber numberWithBool:value];
	return result;
}

void* C_NSNumber_NumberWithDouble(double value) {
	NSNumber* result = [NSNumber numberWithDouble:value];
	return result;
}

void* C_NSNumber_NumberWithFloat(float value) {
	NSNumber* result = [NSNumber numberWithFloat:value];
	return result;
}

void* C_NSNumber_NumberWithInteger(int value) {
	NSNumber* result = [NSNumber numberWithInteger:value];
	return result;
}

void* C_NSNumber_NumberWithUnsignedInteger(unsigned int value) {
	NSNumber* result = [NSNumber numberWithUnsignedInteger:value];
	return result;
}

void* C_NSNumber_InitWithBool(void* ptr, bool value) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithBool:value];
	return result;
}

void* C_NSNumber_InitWithDouble(void* ptr, double value) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithDouble:value];
	return result;
}

void* C_NSNumber_InitWithFloat(void* ptr, float value) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithFloat:value];
	return result;
}

void* C_NSNumber_InitWithInteger(void* ptr, int value) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithInteger:value];
	return result;
}

void* C_NSNumber_InitWithUnsignedInteger(void* ptr, unsigned int value) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSNumber* result = [nSNumber initWithUnsignedInteger:value];
	return result;
}

void* C_NSNumber_DescriptionWithLocale(void* ptr, void* locale) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSString* result = [nSNumber descriptionWithLocale:(id)locale];
	return result;
}

bool C_NSNumber_IsEqualToNumber(void* ptr, void* number) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	bool result = [nSNumber isEqualToNumber:(NSNumber*)number];
	return result;
}

bool C_NSNumber_BoolValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	bool result = [nSNumber boolValue];
	return result;
}

NSDecimal C_NSNumber_DecimalValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSDecimal result = [nSNumber decimalValue];
	return result;
}

double C_NSNumber_DoubleValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	double result = [nSNumber doubleValue];
	return result;
}

float C_NSNumber_FloatValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	float result = [nSNumber floatValue];
	return result;
}

int C_NSNumber_IntegerValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	int result = [nSNumber integerValue];
	return result;
}

unsigned int C_NSNumber_UnsignedIntegerValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	unsigned int result = [nSNumber unsignedIntegerValue];
	return result;
}

void* C_NSNumber_StringValue(void* ptr) {
	NSNumber* nSNumber = (NSNumber*)ptr;
	NSString* result = [nSNumber stringValue];
	return result;
}
