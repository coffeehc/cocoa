#import "date_interval.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDateInterval.h>
#import <Foundation/NSDictionary.h>

void* C_DateInterval_Alloc() {
    return [NSDateInterval alloc];
}

void* C_NSDateInterval_Init(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval init];
    return result_;
}

void* C_NSDateInterval_InitWithStartDate_Duration(void* ptr, void* startDate, double duration) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval initWithStartDate:(NSDate*)startDate duration:duration];
    return result_;
}

void* C_NSDateInterval_InitWithStartDate_EndDate(void* ptr, void* startDate, void* endDate) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval initWithStartDate:(NSDate*)startDate endDate:(NSDate*)endDate];
    return result_;
}

void* C_NSDateInterval_InitWithCoder(void* ptr, void* coder) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSDateInterval_AllocDateInterval() {
    NSDateInterval* result_ = [NSDateInterval alloc];
    return result_;
}

void* C_NSDateInterval_NewDateInterval() {
    NSDateInterval* result_ = [NSDateInterval new];
    return result_;
}

void* C_NSDateInterval_Autorelease(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval autorelease];
    return result_;
}

void* C_NSDateInterval_Retain(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval retain];
    return result_;
}

int C_NSDateInterval_Compare(void* ptr, void* dateInterval) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSComparisonResult result_ = [nSDateInterval compare:(NSDateInterval*)dateInterval];
    return result_;
}

bool C_NSDateInterval_IsEqualToDateInterval(void* ptr, void* dateInterval) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    BOOL result_ = [nSDateInterval isEqualToDateInterval:(NSDateInterval*)dateInterval];
    return result_;
}

bool C_NSDateInterval_IntersectsDateInterval(void* ptr, void* dateInterval) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    BOOL result_ = [nSDateInterval intersectsDateInterval:(NSDateInterval*)dateInterval];
    return result_;
}

void* C_NSDateInterval_IntersectionWithDateInterval(void* ptr, void* dateInterval) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDateInterval* result_ = [nSDateInterval intersectionWithDateInterval:(NSDateInterval*)dateInterval];
    return result_;
}

bool C_NSDateInterval_ContainsDate(void* ptr, void* date) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    BOOL result_ = [nSDateInterval containsDate:(NSDate*)date];
    return result_;
}

void* C_NSDateInterval_StartDate(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDate* result_ = [nSDateInterval startDate];
    return result_;
}

void* C_NSDateInterval_EndDate(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSDate* result_ = [nSDateInterval endDate];
    return result_;
}

double C_NSDateInterval_Duration(void* ptr) {
    NSDateInterval* nSDateInterval = (NSDateInterval*)ptr;
    NSTimeInterval result_ = [nSDateInterval duration];
    return result_;
}
