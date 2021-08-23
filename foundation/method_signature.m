#import <Foundation/Foundation.h>
#import "method_signature.h"

void* C_MethodSignature_Alloc() {
    return [NSMethodSignature alloc];
}

void* C_NSMethodSignature_Init(void* ptr) {
    NSMethodSignature* nSMethodSignature = (NSMethodSignature*)ptr;
    NSMethodSignature* result_ = [nSMethodSignature init];
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
