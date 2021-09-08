#import <WebKit/WebKit.h>
#import "preferences.h"

void* C_Preferences_Alloc() {
    return [WKPreferences alloc];
}

void* C_WKPreferences_AllocPreferences() {
    WKPreferences* result_ = [WKPreferences alloc];
    return result_;
}

void* C_WKPreferences_Init(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    WKPreferences* result_ = [wKPreferences init];
    return result_;
}

void* C_WKPreferences_NewPreferences() {
    WKPreferences* result_ = [WKPreferences new];
    return result_;
}

void* C_WKPreferences_Autorelease(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    WKPreferences* result_ = [wKPreferences autorelease];
    return result_;
}

void* C_WKPreferences_Retain(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    WKPreferences* result_ = [wKPreferences retain];
    return result_;
}

double C_WKPreferences_MinimumFontSize(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    CGFloat result_ = [wKPreferences minimumFontSize];
    return result_;
}

void C_WKPreferences_SetMinimumFontSize(void* ptr, double value) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    [wKPreferences setMinimumFontSize:value];
}

bool C_WKPreferences_TabFocusesLinks(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    BOOL result_ = [wKPreferences tabFocusesLinks];
    return result_;
}

void C_WKPreferences_SetTabFocusesLinks(void* ptr, bool value) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    [wKPreferences setTabFocusesLinks:value];
}

bool C_WKPreferences_JavaScriptCanOpenWindowsAutomatically(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    BOOL result_ = [wKPreferences javaScriptCanOpenWindowsAutomatically];
    return result_;
}

void C_WKPreferences_SetJavaScriptCanOpenWindowsAutomatically(void* ptr, bool value) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    [wKPreferences setJavaScriptCanOpenWindowsAutomatically:value];
}

bool C_WKPreferences_IsFraudulentWebsiteWarningEnabled(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    BOOL result_ = [wKPreferences isFraudulentWebsiteWarningEnabled];
    return result_;
}

void C_WKPreferences_SetFraudulentWebsiteWarningEnabled(void* ptr, bool value) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    [wKPreferences setFraudulentWebsiteWarningEnabled:value];
}

bool C_WKPreferences_TextInteractionEnabled(void* ptr) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    BOOL result_ = [wKPreferences textInteractionEnabled];
    return result_;
}

void C_WKPreferences_SetTextInteractionEnabled(void* ptr, bool value) {
    WKPreferences* wKPreferences = (WKPreferences*)ptr;
    [wKPreferences setTextInteractionEnabled:value];
}
