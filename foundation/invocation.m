#import <Foundation/Foundation.h>
#import "invocation.h"

void* C_Invocation_Alloc() {
    return [NSInvocation alloc];
}

void* C_NSInvocation_AllocInvocation() {
    NSInvocation* result_ = [NSInvocation alloc];
    return result_;
}

void* C_NSInvocation_Init(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    NSInvocation* result_ = [nSInvocation init];
    return result_;
}

void* C_NSInvocation_NewInvocation() {
    NSInvocation* result_ = [NSInvocation new];
    return result_;
}

void* C_NSInvocation_Autorelease(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    NSInvocation* result_ = [nSInvocation autorelease];
    return result_;
}

void* C_NSInvocation_Retain(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    NSInvocation* result_ = [nSInvocation retain];
    return result_;
}

void* C_NSInvocation_InvocationWithMethodSignature(void* sig) {
    NSInvocation* result_ = [NSInvocation invocationWithMethodSignature:(NSMethodSignature*)sig];
    return result_;
}

void C_NSInvocation_RetainArguments(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    [nSInvocation retainArguments];
}

void C_NSInvocation_Invoke(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    [nSInvocation invoke];
}

void C_NSInvocation_InvokeWithTarget(void* ptr, void* target) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    [nSInvocation invokeWithTarget:(id)target];
}

void* C_NSInvocation_Selector(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    SEL result_ = [nSInvocation selector];
    return result_;
}

void C_NSInvocation_SetSelector(void* ptr, void* value) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    [nSInvocation setSelector:(SEL)value];
}

void* C_NSInvocation_Target(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    id result_ = [nSInvocation target];
    return result_;
}

void C_NSInvocation_SetTarget(void* ptr, void* value) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    [nSInvocation setTarget:(id)value];
}

bool C_NSInvocation_ArgumentsRetained(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    BOOL result_ = [nSInvocation argumentsRetained];
    return result_;
}

void* C_NSInvocation_MethodSignature(void* ptr) {
    NSInvocation* nSInvocation = (NSInvocation*)ptr;
    NSMethodSignature* result_ = [nSInvocation methodSignature];
    return result_;
}
