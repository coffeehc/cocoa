#import "user_script.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKUserScript.h>

void* C_UserScript_Alloc() {
    return [WKUserScript alloc];
}

void* C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly(void* ptr, void* source, int injectionTime, bool forMainFrameOnly) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScript* result_ = [wKUserScript initWithSource:(NSString*)source injectionTime:injectionTime forMainFrameOnly:forMainFrameOnly];
    return result_;
}

void* C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(void* ptr, void* source, int injectionTime, bool forMainFrameOnly, void* contentWorld) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScript* result_ = [wKUserScript initWithSource:(NSString*)source injectionTime:injectionTime forMainFrameOnly:forMainFrameOnly inContentWorld:(WKContentWorld*)contentWorld];
    return result_;
}

void* C_WKUserScript_AllocUserScript() {
    WKUserScript* result_ = [WKUserScript alloc];
    return result_;
}

void* C_WKUserScript_Init(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScript* result_ = [wKUserScript init];
    return result_;
}

void* C_WKUserScript_NewUserScript() {
    WKUserScript* result_ = [WKUserScript new];
    return result_;
}

void* C_WKUserScript_Autorelease(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScript* result_ = [wKUserScript autorelease];
    return result_;
}

void* C_WKUserScript_Retain(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScript* result_ = [wKUserScript retain];
    return result_;
}

void* C_WKUserScript_Source(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    NSString* result_ = [wKUserScript source];
    return result_;
}

int C_WKUserScript_InjectionTime(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    WKUserScriptInjectionTime result_ = [wKUserScript injectionTime];
    return result_;
}

bool C_WKUserScript_IsForMainFrameOnly(void* ptr) {
    WKUserScript* wKUserScript = (WKUserScript*)ptr;
    BOOL result_ = [wKUserScript isForMainFrameOnly];
    return result_;
}
