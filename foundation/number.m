#import <Foundation/Foundation.h>
#import "number.h"

bool Number_BoolValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number boolValue];
}

float Number_FloatValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number floatValue];
}

long Number_IntValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number intValue];
}

long Number_LongLongValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number longLongValue];
}

unsigned long Number_UnsignedIntValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number unsignedIntValue];
}

unsigned long Number_UnsignedLongLongValue(void* ptr) {
	NSNumber* number = (NSNumber*)ptr;
	return [number unsignedLongLongValue];
}

void* Number_NumberWithBool(bool value) {
	return [NSNumber numberWithBool:value];
}

void* Number_NumberWithDouble(double value) {
	return [NSNumber numberWithDouble:value];
}

void* Number_NumberWithFloat(float value) {
	return [NSNumber numberWithFloat:value];
}

void* Number_NumberWithInt(long value) {
	return [NSNumber numberWithInt:value];
}

void* Number_NumberWithLongLong(long value) {
	return [NSNumber numberWithLongLong:value];
}

void* Number_NumberWithUnsignedInt(unsigned long value) {
	return [NSNumber numberWithUnsignedInt:value];
}

void* Number_NumberWithUnsignedLongLong(unsigned long value) {
	return [NSNumber numberWithUnsignedLongLong:value];
}
