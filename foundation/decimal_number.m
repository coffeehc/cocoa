#import "decimal_number.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDecimalNumber.h>
#import <Foundation/NSDictionary.h>

void* C_DecimalNumberHandler_Alloc() {
    return [NSDecimalNumberHandler alloc];
}

void* C_NSDecimalNumberHandler_AllocDecimalNumberHandler() {
    NSDecimalNumberHandler* result_ = [NSDecimalNumberHandler alloc];
    return result_;
}

void* C_NSDecimalNumberHandler_Init(void* ptr) {
    NSDecimalNumberHandler* nSDecimalNumberHandler = (NSDecimalNumberHandler*)ptr;
    NSDecimalNumberHandler* result_ = [nSDecimalNumberHandler init];
    return result_;
}

void* C_NSDecimalNumberHandler_NewDecimalNumberHandler() {
    NSDecimalNumberHandler* result_ = [NSDecimalNumberHandler new];
    return result_;
}

void* C_NSDecimalNumberHandler_Autorelease(void* ptr) {
    NSDecimalNumberHandler* nSDecimalNumberHandler = (NSDecimalNumberHandler*)ptr;
    NSDecimalNumberHandler* result_ = [nSDecimalNumberHandler autorelease];
    return result_;
}

void* C_NSDecimalNumberHandler_Retain(void* ptr) {
    NSDecimalNumberHandler* nSDecimalNumberHandler = (NSDecimalNumberHandler*)ptr;
    NSDecimalNumberHandler* result_ = [nSDecimalNumberHandler retain];
    return result_;
}

void* C_NSDecimalNumberHandler_DefaultDecimalNumberHandler() {
    NSDecimalNumberHandler* result_ = [NSDecimalNumberHandler defaultDecimalNumberHandler];
    return result_;
}
