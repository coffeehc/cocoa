#import "storyboard.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSStoryboard.h>

void* C_Storyboard_Alloc() {
    return [NSStoryboard alloc];
}

void* C_NSStoryboard_StoryboardWithName_Bundle(void* name, void* storyboardBundleOrNil) {
    NSStoryboard* result_ = [NSStoryboard storyboardWithName:(NSString*)name bundle:(NSBundle*)storyboardBundleOrNil];
    return result_;
}

void* C_NSStoryboard_AllocStoryboard() {
    NSStoryboard* result_ = [NSStoryboard alloc];
    return result_;
}

void* C_NSStoryboard_Init(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    NSStoryboard* result_ = [nSStoryboard init];
    return result_;
}

void* C_NSStoryboard_NewStoryboard() {
    NSStoryboard* result_ = [NSStoryboard new];
    return result_;
}

void* C_NSStoryboard_Autorelease(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    NSStoryboard* result_ = [nSStoryboard autorelease];
    return result_;
}

void* C_NSStoryboard_Retain(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    NSStoryboard* result_ = [nSStoryboard retain];
    return result_;
}

void* C_NSStoryboard_InstantiateInitialController(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    id result_ = [nSStoryboard instantiateInitialController];
    return result_;
}

void* C_NSStoryboard_MainStoryboard() {
    NSStoryboard* result_ = [NSStoryboard mainStoryboard];
    return result_;
}
