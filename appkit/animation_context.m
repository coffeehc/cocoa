#import <Appkit/Appkit.h>
#import "animation_context.h"

void* C_AnimationContext_Alloc() {
    return [NSAnimationContext alloc];
}

void* C_NSAnimationContext_AllocAnimationContext() {
    NSAnimationContext* result_ = [NSAnimationContext alloc];
    return result_;
}

void* C_NSAnimationContext_Init(void* ptr) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    NSAnimationContext* result_ = [nSAnimationContext init];
    return result_;
}

void* C_NSAnimationContext_NewAnimationContext() {
    NSAnimationContext* result_ = [NSAnimationContext new];
    return result_;
}

void* C_NSAnimationContext_Autorelease(void* ptr) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    NSAnimationContext* result_ = [nSAnimationContext autorelease];
    return result_;
}

void* C_NSAnimationContext_Retain(void* ptr) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    NSAnimationContext* result_ = [nSAnimationContext retain];
    return result_;
}

void C_NSAnimationContext_AnimationContext_BeginGrouping() {
    [NSAnimationContext beginGrouping];
}

void C_NSAnimationContext_AnimationContext_EndGrouping() {
    [NSAnimationContext endGrouping];
}

void* C_NSAnimationContext_AnimationContext_CurrentContext() {
    NSAnimationContext* result_ = [NSAnimationContext currentContext];
    return result_;
}

double C_NSAnimationContext_Duration(void* ptr) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    NSTimeInterval result_ = [nSAnimationContext duration];
    return result_;
}

void C_NSAnimationContext_SetDuration(void* ptr, double value) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    [nSAnimationContext setDuration:value];
}

bool C_NSAnimationContext_AllowsImplicitAnimation(void* ptr) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    BOOL result_ = [nSAnimationContext allowsImplicitAnimation];
    return result_;
}

void C_NSAnimationContext_SetAllowsImplicitAnimation(void* ptr, bool value) {
    NSAnimationContext* nSAnimationContext = (NSAnimationContext*)ptr;
    [nSAnimationContext setAllowsImplicitAnimation:value];
}
