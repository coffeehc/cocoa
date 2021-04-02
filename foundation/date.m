#import <Foundation/Foundation.h>
#import "date.h"

void* Date_DistantFuture() {
	return [NSDate distantFuture];
}

void* Date_DistantPast() {
	return [NSDate distantPast];
}

double Date_TimeIntervalSinceNow(void* ptr) {
	NSDate* date = (NSDate*)ptr;
	return [date timeIntervalSinceNow];
}

double Date_TimeIntervalSinceReferenceDate(void* ptr) {
	NSDate* date = (NSDate*)ptr;
	return [date timeIntervalSinceReferenceDate];
}

double Date_TimeIntervalSince1970(void* ptr) {
	NSDate* date = (NSDate*)ptr;
	return [date timeIntervalSince1970];
}

const char* Date_Description(void* ptr) {
	NSDate* date = (NSDate*)ptr;
	return [[date description] UTF8String];
}

void* Date_NewDate() {
	NSDate* date = [NSDate alloc];
	return [[date init] autorelease];
}

void* Date_CurrentDate() {
	return [NSDate date];
}

void* Date_DateWithTimeIntervalSinceNow(double secs) {
	return [NSDate dateWithTimeIntervalSinceNow:secs];
}

void* Date_dateWithTimeIntervalSinceDate(double secsToBeAdded, void* date) {
	return [NSDate dateWithTimeInterval:secsToBeAdded sinceDate:(NSDate*)date];
}

void* Date_DateWithTimeIntervalSinceReferenceDate(double secs) {
	return [NSDate dateWithTimeIntervalSinceReferenceDate:secs];
}

void* Date_DateWithTimeIntervalSince1970(double secs) {
	return [NSDate dateWithTimeIntervalSince1970:secs];
}

const char* Date_DescriptionWithLocale(void* ptr, void* locale) {
	NSDate* date = (NSDate*)ptr;
	return [[date descriptionWithLocale:(NSObject*)locale] UTF8String];
}

bool Date_IsEqualToDate(void* ptr, void* otherDate) {
	NSDate* date = (NSDate*)ptr;
	return [date isEqualToDate:(NSDate*)otherDate];
}

void* Date_EarlierDate(void* ptr, void* otherDate) {
	NSDate* date = (NSDate*)ptr;
	return [date earlierDate:(NSDate*)otherDate];
}

void* Date_LaterDate(void* ptr, void* otherDate) {
	NSDate* date = (NSDate*)ptr;
	return [date laterDate:(NSDate*)otherDate];
}

unsigned long Date_Compare(void* ptr, void* other) {
	NSDate* date = (NSDate*)ptr;
	return [date compare:(NSDate*)other];
}

double Date_TimeIntervalSinceDate(void* ptr, void* other) {
	NSDate* date = (NSDate*)ptr;
	return [date timeIntervalSinceDate:(NSDate*)other];
}
