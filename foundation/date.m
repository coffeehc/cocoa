#import <Foundation/Foundation.h>
#import "date.h"

void* C_Date_Alloc() {
    return [NSDate alloc];
}

void* C_NSDate_Init(void* ptr) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate init];
    return result_;
}

void* C_NSDate_InitWithTimeIntervalSinceNow(void* ptr, double secs) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate initWithTimeIntervalSinceNow:secs];
    return result_;
}

void* C_NSDate_InitWithTimeInterval_SinceDate(void* ptr, double secsToBeAdded, void* date) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate initWithTimeInterval:secsToBeAdded sinceDate:(NSDate*)date];
    return result_;
}

void* C_NSDate_InitWithTimeIntervalSinceReferenceDate(void* ptr, double ti) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate initWithTimeIntervalSinceReferenceDate:ti];
    return result_;
}

void* C_NSDate_InitWithTimeIntervalSince1970(void* ptr, double secs) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate initWithTimeIntervalSince1970:secs];
    return result_;
}

void* C_NSDate_InitWithCoder(void* ptr, void* coder) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSDate_DateByAddingTimeInterval(void* ptr, double ti) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate dateByAddingTimeInterval:ti];
    return result_;
}

void* C_NSDate_CurrentDate() {
    NSDate* result_ = [NSDate date];
    return result_;
}

void* C_NSDate_DateWithTimeIntervalSinceNow(double secs) {
    NSDate* result_ = [NSDate dateWithTimeIntervalSinceNow:secs];
    return result_;
}

void* C_NSDate_DateWithTimeInterval_SinceDate(double secsToBeAdded, void* date) {
    NSDate* result_ = [NSDate dateWithTimeInterval:secsToBeAdded sinceDate:(NSDate*)date];
    return result_;
}

void* C_NSDate_DateWithTimeIntervalSinceReferenceDate(double ti) {
    NSDate* result_ = [NSDate dateWithTimeIntervalSinceReferenceDate:ti];
    return result_;
}

void* C_NSDate_DateWithTimeIntervalSince1970(double secs) {
    NSDate* result_ = [NSDate dateWithTimeIntervalSince1970:secs];
    return result_;
}

bool C_NSDate_IsEqualToDate(void* ptr, void* otherDate) {
    NSDate* nSDate = (NSDate*)ptr;
    BOOL result_ = [nSDate isEqualToDate:(NSDate*)otherDate];
    return result_;
}

void* C_NSDate_EarlierDate(void* ptr, void* anotherDate) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate earlierDate:(NSDate*)anotherDate];
    return result_;
}

void* C_NSDate_LaterDate(void* ptr, void* anotherDate) {
    NSDate* nSDate = (NSDate*)ptr;
    NSDate* result_ = [nSDate laterDate:(NSDate*)anotherDate];
    return result_;
}

int C_NSDate_Compare(void* ptr, void* other) {
    NSDate* nSDate = (NSDate*)ptr;
    NSComparisonResult result_ = [nSDate compare:(NSDate*)other];
    return result_;
}

double C_NSDate_TimeIntervalSinceDate(void* ptr, void* anotherDate) {
    NSDate* nSDate = (NSDate*)ptr;
    NSTimeInterval result_ = [nSDate timeIntervalSinceDate:(NSDate*)anotherDate];
    return result_;
}

void* C_NSDate_DescriptionWithLocale(void* ptr, void* locale) {
    NSDate* nSDate = (NSDate*)ptr;
    NSString* result_ = [nSDate descriptionWithLocale:(id)locale];
    return result_;
}

void* C_NSDate_Date_DistantFuture() {
    NSDate* result_ = [NSDate distantFuture];
    return result_;
}

void* C_NSDate_Date_DistantPast() {
    NSDate* result_ = [NSDate distantPast];
    return result_;
}

void* C_NSDate_Date_Now() {
    NSDate* result_ = [NSDate now];
    return result_;
}

double C_NSDate_TimeIntervalSinceNow(void* ptr) {
    NSDate* nSDate = (NSDate*)ptr;
    NSTimeInterval result_ = [nSDate timeIntervalSinceNow];
    return result_;
}

double C_NSDate_TimeIntervalSinceReferenceDate(void* ptr) {
    NSDate* nSDate = (NSDate*)ptr;
    NSTimeInterval result_ = [nSDate timeIntervalSinceReferenceDate];
    return result_;
}

double C_NSDate_TimeIntervalSince1970(void* ptr) {
    NSDate* nSDate = (NSDate*)ptr;
    NSTimeInterval result_ = [nSDate timeIntervalSince1970];
    return result_;
}

void* C_NSDate_Description(void* ptr) {
    NSDate* nSDate = (NSDate*)ptr;
    NSString* result_ = [nSDate description];
    return result_;
}
