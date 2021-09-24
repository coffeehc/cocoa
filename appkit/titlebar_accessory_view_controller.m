#import "titlebar_accessory_view_controller.h"
#import <AppKit/NSTitlebarAccessoryViewController.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TitlebarAccessoryViewController_Alloc() {
    return [NSTitlebarAccessoryViewController alloc];
}

void* C_NSTitlebarAccessoryViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSTitlebarAccessoryViewController* result_ = [nSTitlebarAccessoryViewController initWithNibName:(NSString*)nibNameOrNil bundle:(NSBundle*)nibBundleOrNil];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_InitWithCoder(void* ptr, void* coder) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSTitlebarAccessoryViewController* result_ = [nSTitlebarAccessoryViewController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_Init(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSTitlebarAccessoryViewController* result_ = [nSTitlebarAccessoryViewController init];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_AllocTitlebarAccessoryViewController() {
    NSTitlebarAccessoryViewController* result_ = [NSTitlebarAccessoryViewController alloc];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_NewTitlebarAccessoryViewController() {
    NSTitlebarAccessoryViewController* result_ = [NSTitlebarAccessoryViewController new];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_Autorelease(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSTitlebarAccessoryViewController* result_ = [nSTitlebarAccessoryViewController autorelease];
    return result_;
}

void* C_NSTitlebarAccessoryViewController_Retain(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSTitlebarAccessoryViewController* result_ = [nSTitlebarAccessoryViewController retain];
    return result_;
}

double C_NSTitlebarAccessoryViewController_FullScreenMinHeight(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    CGFloat result_ = [nSTitlebarAccessoryViewController fullScreenMinHeight];
    return result_;
}

void C_NSTitlebarAccessoryViewController_SetFullScreenMinHeight(void* ptr, double value) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    [nSTitlebarAccessoryViewController setFullScreenMinHeight:value];
}

int C_NSTitlebarAccessoryViewController_LayoutAttribute(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    NSLayoutAttribute result_ = [nSTitlebarAccessoryViewController layoutAttribute];
    return result_;
}

void C_NSTitlebarAccessoryViewController_SetLayoutAttribute(void* ptr, int value) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    [nSTitlebarAccessoryViewController setLayoutAttribute:value];
}

bool C_NSTitlebarAccessoryViewController_AutomaticallyAdjustsSize(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    BOOL result_ = [nSTitlebarAccessoryViewController automaticallyAdjustsSize];
    return result_;
}

void C_NSTitlebarAccessoryViewController_SetAutomaticallyAdjustsSize(void* ptr, bool value) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    [nSTitlebarAccessoryViewController setAutomaticallyAdjustsSize:value];
}

bool C_NSTitlebarAccessoryViewController_IsHidden(void* ptr) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    BOOL result_ = [nSTitlebarAccessoryViewController isHidden];
    return result_;
}

void C_NSTitlebarAccessoryViewController_SetHidden(void* ptr, bool value) {
    NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
    [nSTitlebarAccessoryViewController setHidden:value];
}
