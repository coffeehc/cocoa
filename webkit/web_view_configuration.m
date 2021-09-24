#import "web_view_configuration.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKWebViewConfiguration.h>

void* C_WebViewConfiguration_Alloc() {
    return [WKWebViewConfiguration alloc];
}

void* C_WKWebViewConfiguration_AllocWebViewConfiguration() {
    WKWebViewConfiguration* result_ = [WKWebViewConfiguration alloc];
    return result_;
}

void* C_WKWebViewConfiguration_Init(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKWebViewConfiguration* result_ = [wKWebViewConfiguration init];
    return result_;
}

void* C_WKWebViewConfiguration_NewWebViewConfiguration() {
    WKWebViewConfiguration* result_ = [WKWebViewConfiguration new];
    return result_;
}

void* C_WKWebViewConfiguration_Autorelease(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKWebViewConfiguration* result_ = [wKWebViewConfiguration autorelease];
    return result_;
}

void* C_WKWebViewConfiguration_Retain(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKWebViewConfiguration* result_ = [wKWebViewConfiguration retain];
    return result_;
}

void C_WKWebViewConfiguration_SetURLSchemeHandler_ForURLScheme(void* ptr, void* urlSchemeHandler, void* urlScheme) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setURLSchemeHandler:(id)urlSchemeHandler forURLScheme:(NSString*)urlScheme];
}

void* C_WKWebViewConfiguration_UrlSchemeHandlerForURLScheme(void* ptr, void* urlScheme) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    id result_ = [wKWebViewConfiguration urlSchemeHandlerForURLScheme:(NSString*)urlScheme];
    return result_;
}

void* C_WKWebViewConfiguration_WebsiteDataStore(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKWebsiteDataStore* result_ = [wKWebViewConfiguration websiteDataStore];
    return result_;
}

void C_WKWebViewConfiguration_SetWebsiteDataStore(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setWebsiteDataStore:(WKWebsiteDataStore*)value];
}

void* C_WKWebViewConfiguration_UserContentController(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKUserContentController* result_ = [wKWebViewConfiguration userContentController];
    return result_;
}

void C_WKWebViewConfiguration_SetUserContentController(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setUserContentController:(WKUserContentController*)value];
}

void* C_WKWebViewConfiguration_ProcessPool(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKProcessPool* result_ = [wKWebViewConfiguration processPool];
    return result_;
}

void C_WKWebViewConfiguration_SetProcessPool(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setProcessPool:(WKProcessPool*)value];
}

void* C_WKWebViewConfiguration_ApplicationNameForUserAgent(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    NSString* result_ = [wKWebViewConfiguration applicationNameForUserAgent];
    return result_;
}

void C_WKWebViewConfiguration_SetApplicationNameForUserAgent(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setApplicationNameForUserAgent:(NSString*)value];
}

bool C_WKWebViewConfiguration_LimitsNavigationsToAppBoundDomains(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    BOOL result_ = [wKWebViewConfiguration limitsNavigationsToAppBoundDomains];
    return result_;
}

void C_WKWebViewConfiguration_SetLimitsNavigationsToAppBoundDomains(void* ptr, bool value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setLimitsNavigationsToAppBoundDomains:value];
}

void* C_WKWebViewConfiguration_Preferences(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKPreferences* result_ = [wKWebViewConfiguration preferences];
    return result_;
}

void C_WKWebViewConfiguration_SetPreferences(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setPreferences:(WKPreferences*)value];
}

void* C_WKWebViewConfiguration_DefaultWebpagePreferences(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKWebpagePreferences* result_ = [wKWebViewConfiguration defaultWebpagePreferences];
    return result_;
}

void C_WKWebViewConfiguration_SetDefaultWebpagePreferences(void* ptr, void* value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setDefaultWebpagePreferences:(WKWebpagePreferences*)value];
}

bool C_WKWebViewConfiguration_SuppressesIncrementalRendering(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    BOOL result_ = [wKWebViewConfiguration suppressesIncrementalRendering];
    return result_;
}

void C_WKWebViewConfiguration_SetSuppressesIncrementalRendering(void* ptr, bool value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setSuppressesIncrementalRendering:value];
}

bool C_WKWebViewConfiguration_AllowsAirPlayForMediaPlayback(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    BOOL result_ = [wKWebViewConfiguration allowsAirPlayForMediaPlayback];
    return result_;
}

void C_WKWebViewConfiguration_SetAllowsAirPlayForMediaPlayback(void* ptr, bool value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setAllowsAirPlayForMediaPlayback:value];
}

unsigned int C_WKWebViewConfiguration_MediaTypesRequiringUserActionForPlayback(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKAudiovisualMediaTypes result_ = [wKWebViewConfiguration mediaTypesRequiringUserActionForPlayback];
    return result_;
}

void C_WKWebViewConfiguration_SetMediaTypesRequiringUserActionForPlayback(void* ptr, unsigned int value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setMediaTypesRequiringUserActionForPlayback:value];
}

int C_WKWebViewConfiguration_UserInterfaceDirectionPolicy(void* ptr) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    WKUserInterfaceDirectionPolicy result_ = [wKWebViewConfiguration userInterfaceDirectionPolicy];
    return result_;
}

void C_WKWebViewConfiguration_SetUserInterfaceDirectionPolicy(void* ptr, int value) {
    WKWebViewConfiguration* wKWebViewConfiguration = (WKWebViewConfiguration*)ptr;
    [wKWebViewConfiguration setUserInterfaceDirectionPolicy:value];
}
