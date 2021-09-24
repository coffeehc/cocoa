#import "controller.h"
#import <AppKit/NSController.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Controller_Alloc() {
    return [NSController alloc];
}

void* C_NSController_Init(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    NSController* result_ = [nSController init];
    return result_;
}

void* C_NSController_InitWithCoder(void* ptr, void* coder) {
    NSController* nSController = (NSController*)ptr;
    NSController* result_ = [nSController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSController_AllocController() {
    NSController* result_ = [NSController alloc];
    return result_;
}

void* C_NSController_NewController() {
    NSController* result_ = [NSController new];
    return result_;
}

void* C_NSController_Autorelease(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    NSController* result_ = [nSController autorelease];
    return result_;
}

void* C_NSController_Retain(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    NSController* result_ = [nSController retain];
    return result_;
}

void C_NSController_ObjectDidBeginEditing(void* ptr, void* editor) {
    NSController* nSController = (NSController*)ptr;
    [nSController objectDidBeginEditing:(id)editor];
}

void C_NSController_ObjectDidEndEditing(void* ptr, void* editor) {
    NSController* nSController = (NSController*)ptr;
    [nSController objectDidEndEditing:(id)editor];
}

bool C_NSController_CommitEditing(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    BOOL result_ = [nSController commitEditing];
    return result_;
}

void C_NSController_DiscardEditing(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    [nSController discardEditing];
}

bool C_NSController_IsEditing(void* ptr) {
    NSController* nSController = (NSController*)ptr;
    BOOL result_ = [nSController isEditing];
    return result_;
}
