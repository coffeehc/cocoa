#import "method_signature.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSMethodSignature.h>

void* C_MethodSignature_Alloc() {
    return [NSMethodSignature alloc];
}

void* C_NSMethodSignature_AllocMethodSignature() {
    NSMethodSignature* result_ = [NSMethodSignature alloc];
    return result_;
}

void* C_NSMethodSignature_Init(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSMethodSignature* result_ = [nSMethodSignature init];
    return result_;
}

void* C_NSMethodSignature_NewMethodSignature() {
    NSMethodSignature* result_ = [NSMethodSignature new];
    return result_;
}

void* C_NSMethodSignature_Autorelease(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSMethodSignature* result_ = [nSMethodSignature autorelease];
    return result_;
}

void* C_NSMethodSignature_Retain(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSMethodSignature* result_ = [nSMethodSignature retain];
    return result_;
}

bool C_NSMethodSignature_IsOneway(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    BOOL result_ = [nSMethodSignature isOneway];
    return result_;
}

unsigned int C_NSMethodSignature_NumberOfArguments(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSUInteger result_ = [nSMethodSignature numberOfArguments];
    return result_;
}

unsigned int C_NSMethodSignature_FrameLength(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSUInteger result_ = [nSMethodSignature frameLength];
    return result_;
}

unsigned int C_NSMethodSignature_MethodReturnLength(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSUInteger result_ = [nSMethodSignature methodReturnLength];
    return result_;
}
