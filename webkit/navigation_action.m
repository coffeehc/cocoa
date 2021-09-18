#import "navigation_action.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKNavigationAction.h>

void* C_NavigationAction_Alloc() {
    return [WKNavigationAction alloc];
}

void* C_WKNavigationAction_AllocNavigationAction() {
    WKNavigationAction* result_ = [WKNavigationAction alloc];
    return result_;
}

void* C_WKNavigationAction_Init(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKNavigationAction* result_ = [wKNavigationAction init];
    return result_;
}

void* C_WKNavigationAction_NewNavigationAction() {
    WKNavigationAction* result_ = [WKNavigationAction new];
    return result_;
}

void* C_WKNavigationAction_Autorelease(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKNavigationAction* result_ = [wKNavigationAction autorelease];
    return result_;
}

void* C_WKNavigationAction_Retain(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKNavigationAction* result_ = [wKNavigationAction retain];
    return result_;
}

int C_WKNavigationAction_NavigationType(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKNavigationType result_ = [wKNavigationAction navigationType];
    return result_;
}

void* C_WKNavigationAction_Request(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    NSURLRequest* result_ = [wKNavigationAction request];
    return result_;
}

void* C_WKNavigationAction_SourceFrame(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKFrameInfo* result_ = [wKNavigationAction sourceFrame];
    return result_;
}

void* C_WKNavigationAction_TargetFrame(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    WKFrameInfo* result_ = [wKNavigationAction targetFrame];
    return result_;
}

int C_WKNavigationAction_ButtonNumber(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    NSInteger result_ = [wKNavigationAction buttonNumber];
    return result_;
}

unsigned int C_WKNavigationAction_ModifierFlags(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    NSEventModifierFlags result_ = [wKNavigationAction modifierFlags];
    return result_;
}

bool C_WKNavigationAction_ShouldPerformDownload(void* ptr) {
    WKNavigationAction* wKNavigationAction = (WKNavigationAction*)ptr;
    BOOL result_ = [wKNavigationAction shouldPerformDownload];
    return result_;
}
