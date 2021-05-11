#import <Foundation/Foundation.h>
#import "date.h"

void* C_Date_Alloc() {
	return [NSDate alloc];
}

void* C_NSDate_Init(void* ptr) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate init];
	return result;
}

void* C_NSDate_InitWithTimeIntervalSinceNow(void* ptr, double secs) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate initWithTimeIntervalSinceNow:secs];
	return result;
}

void* C_NSDate_InitWithTimeInterval_SinceDate(void* ptr, double secsToBeAdded, void* date) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate initWithTimeInterval:secsToBeAdded sinceDate:(NSDate*)date];
	return result;
}

void* C_NSDate_InitWithTimeIntervalSinceReferenceDate(void* ptr, double ti) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate initWithTimeIntervalSinceReferenceDate:ti];
	return result;
}

void* C_NSDate_InitWithTimeIntervalSince1970(void* ptr, double secs) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate initWithTimeIntervalSince1970:secs];
	return result;
}

void* C_NSDate_InitWithCoder(void* ptr, void* coder) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSDate_DateByAddingTimeInterval(void* ptr, double ti) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate dateByAddingTimeInterval:ti];
	return result;
}

void* C_NSDate_CurrentDate() {
	NSDate* result = [NSDate date];
	return result;
}

void* C_NSDate_DateWithTimeIntervalSinceNow(double secs) {
	NSDate* result = [NSDate dateWithTimeIntervalSinceNow:secs];
	return result;
}

void* C_NSDate_DateWithTimeInterval_SinceDate(double secsToBeAdded, void* date) {
	NSDate* result = [NSDate dateWithTimeInterval:secsToBeAdded sinceDate:(NSDate*)date];
	return result;
}

void* C_NSDate_DateWithTimeIntervalSinceReferenceDate(double ti) {
	NSDate* result = [NSDate dateWithTimeIntervalSinceReferenceDate:ti];
	return result;
}

void* C_NSDate_DateWithTimeIntervalSince1970(double secs) {
	NSDate* result = [NSDate dateWithTimeIntervalSince1970:secs];
	return result;
}

bool C_NSDate_IsEqualToDate(void* ptr, void* otherDate) {
	NSDate* nSDate = (NSDate*)ptr;
	bool result = [nSDate isEqualToDate:(NSDate*)otherDate];
	return result;
}

void* C_NSDate_EarlierDate(void* ptr, void* anotherDate) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate earlierDate:(NSDate*)anotherDate];
	return result;
}

void* C_NSDate_LaterDate(void* ptr, void* anotherDate) {
	NSDate* nSDate = (NSDate*)ptr;
	NSDate* result = [nSDate laterDate:(NSDate*)anotherDate];
	return result;
}

double C_NSDate_TimeIntervalSinceDate(void* ptr, void* anotherDate) {
	NSDate* nSDate = (NSDate*)ptr;
	double result = [nSDate timeIntervalSinceDate:(NSDate*)anotherDate];
	return result;
}

void* C_NSDate_DescriptionWithLocale(void* ptr, void* locale) {
	NSDate* nSDate = (NSDate*)ptr;
	NSString* result = [nSDate descriptionWithLocale:(id)locale];
	return result;
}

double C_NSDate_TimeIntervalSinceNow(void* ptr) {
	NSDate* nSDate = (NSDate*)ptr;
	double result = [nSDate timeIntervalSinceNow];
	return result;
}

double C_NSDate_TimeIntervalSinceReferenceDate(void* ptr) {
	NSDate* nSDate = (NSDate*)ptr;
	double result = [nSDate timeIntervalSinceReferenceDate];
	return result;
}

double C_NSDate_TimeIntervalSince1970(void* ptr) {
	NSDate* nSDate = (NSDate*)ptr;
	double result = [nSDate timeIntervalSince1970];
	return result;
}

void* C_NSDate_Description(void* ptr) {
	NSDate* nSDate = (NSDate*)ptr;
	NSString* result = [nSDate description];
	return result;
}
