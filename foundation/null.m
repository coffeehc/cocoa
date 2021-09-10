#import <Foundation/Foundation.h>
#import "null.h"

void* C_Null_Alloc() {
    return [NSNull alloc];
}

void* C_NSNull_AllocNull() {
    NSNull* result_ = [NSNull alloc];
    return result_;
}

void* C_NSNull_Init(void* ptr) {
    NSNull* nSNull = (NSNull*)ptr;
    NSNull* result_ = [nSNull init];
    return result_;
}

void* C_NSNull_NewNull() {
    NSNull* result_ = [NSNull new];
    return result_;
}

void* C_NSNull_Autorelease(void* ptr) {
    NSNull* nSNull = (NSNull*)ptr;
    NSNull* result_ = [nSNull autorelease];
    return result_;
}

void* C_NSNull_Retain(void* ptr) {
    NSNull* nSNull = (NSNull*)ptr;
    NSNull* result_ = [nSNull retain];
    return result_;
}

void* C_NSNull_Null_() {
    NSNull* result_ = [NSNull null];
    return result_;
}
