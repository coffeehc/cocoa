#import <WebKit/WebKit.h>
#import "webpage_preferences.h"

void* C_WebpagePreferences_Alloc() {
    return [WKWebpagePreferences alloc];
}

void* C_WKWebpagePreferences_Init(void* ptr) {
    WKWebpagePreferences* wKWebpagePreferences = (WKWebpagePreferences*)ptr;
    WKWebpagePreferences* result_ = [wKWebpagePreferences init];
    return result_;
}

bool C_WKWebpagePreferences_AllowsContentJavaScript(void* ptr) {
    WKWebpagePreferences* wKWebpagePreferences = (WKWebpagePreferences*)ptr;
    BOOL result_ = [wKWebpagePreferences allowsContentJavaScript];
    return result_;
}

void C_WKWebpagePreferences_SetAllowsContentJavaScript(void* ptr, bool value) {
    WKWebpagePreferences* wKWebpagePreferences = (WKWebpagePreferences*)ptr;
    [wKWebpagePreferences setAllowsContentJavaScript:value];
}

int C_WKWebpagePreferences_PreferredContentMode(void* ptr) {
    WKWebpagePreferences* wKWebpagePreferences = (WKWebpagePreferences*)ptr;
    WKContentMode result_ = [wKWebpagePreferences preferredContentMode];
    return result_;
}

void C_WKWebpagePreferences_SetPreferredContentMode(void* ptr, int value) {
    WKWebpagePreferences* wKWebpagePreferences = (WKWebpagePreferences*)ptr;
    [wKWebpagePreferences setPreferredContentMode:value];
}
