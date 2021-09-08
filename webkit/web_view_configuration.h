#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_WebViewConfiguration_Alloc();

void* C_WKWebViewConfiguration_AllocWebViewConfiguration();
void* C_WKWebViewConfiguration_Init(void* ptr);
void* C_WKWebViewConfiguration_NewWebViewConfiguration();
void* C_WKWebViewConfiguration_Autorelease(void* ptr);
void* C_WKWebViewConfiguration_Retain(void* ptr);
void C_WKWebViewConfiguration_SetURLSchemeHandler_ForURLScheme(void* ptr, void* urlSchemeHandler, void* urlScheme);
void* C_WKWebViewConfiguration_UrlSchemeHandlerForURLScheme(void* ptr, void* urlScheme);
void* C_WKWebViewConfiguration_WebsiteDataStore(void* ptr);
void C_WKWebViewConfiguration_SetWebsiteDataStore(void* ptr, void* value);
void* C_WKWebViewConfiguration_UserContentController(void* ptr);
void C_WKWebViewConfiguration_SetUserContentController(void* ptr, void* value);
void* C_WKWebViewConfiguration_ProcessPool(void* ptr);
void C_WKWebViewConfiguration_SetProcessPool(void* ptr, void* value);
void* C_WKWebViewConfiguration_ApplicationNameForUserAgent(void* ptr);
void C_WKWebViewConfiguration_SetApplicationNameForUserAgent(void* ptr, void* value);
bool C_WKWebViewConfiguration_LimitsNavigationsToAppBoundDomains(void* ptr);
void C_WKWebViewConfiguration_SetLimitsNavigationsToAppBoundDomains(void* ptr, bool value);
void* C_WKWebViewConfiguration_Preferences(void* ptr);
void C_WKWebViewConfiguration_SetPreferences(void* ptr, void* value);
void* C_WKWebViewConfiguration_DefaultWebpagePreferences(void* ptr);
void C_WKWebViewConfiguration_SetDefaultWebpagePreferences(void* ptr, void* value);
bool C_WKWebViewConfiguration_SuppressesIncrementalRendering(void* ptr);
void C_WKWebViewConfiguration_SetSuppressesIncrementalRendering(void* ptr, bool value);
bool C_WKWebViewConfiguration_AllowsAirPlayForMediaPlayback(void* ptr);
void C_WKWebViewConfiguration_SetAllowsAirPlayForMediaPlayback(void* ptr, bool value);
unsigned int C_WKWebViewConfiguration_MediaTypesRequiringUserActionForPlayback(void* ptr);
void C_WKWebViewConfiguration_SetMediaTypesRequiringUserActionForPlayback(void* ptr, unsigned int value);
int C_WKWebViewConfiguration_UserInterfaceDirectionPolicy(void* ptr);
void C_WKWebViewConfiguration_SetUserInterfaceDirectionPolicy(void* ptr, int value);
