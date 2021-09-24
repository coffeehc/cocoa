#import "touch.h"
#import <AppKit/NSTouch.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Touch_Alloc() {
    return [NSTouch alloc];
}

void* C_NSTouch_AllocTouch() {
    NSTouch* result_ = [NSTouch alloc];
    return result_;
}

void* C_NSTouch_Init(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouch* result_ = [nSTouch init];
    return result_;
}

void* C_NSTouch_NewTouch() {
    NSTouch* result_ = [NSTouch new];
    return result_;
}

void* C_NSTouch_Autorelease(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouch* result_ = [nSTouch autorelease];
    return result_;
}

void* C_NSTouch_Retain(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouch* result_ = [nSTouch retain];
    return result_;
}

CGPoint C_NSTouch_LocationInView(void* ptr, void* view) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result_ = [nSTouch locationInView:(NSView*)view];
    return result_;
}

CGPoint C_NSTouch_PreviousLocationInView(void* ptr, void* view) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result_ = [nSTouch previousLocationInView:(NSView*)view];
    return result_;
}

int C_NSTouch_Type(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouchType result_ = [nSTouch type];
    return result_;
}

void* C_NSTouch_Identity(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    id result_ = [nSTouch identity];
    return result_;
}

unsigned int C_NSTouch_Phase(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouchPhase result_ = [nSTouch phase];
    return result_;
}

CGPoint C_NSTouch_NormalizedPosition(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result_ = [nSTouch normalizedPosition];
    return result_;
}

bool C_NSTouch_IsResting(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    BOOL result_ = [nSTouch isResting];
    return result_;
}

void* C_NSTouch_Device(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    id result_ = [nSTouch device];
    return result_;
}

CGSize C_NSTouch_DeviceSize(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSSize result_ = [nSTouch deviceSize];
    return result_;
}
