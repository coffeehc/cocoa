#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_WebpagePreferences_Alloc();

void* C_WKWebpagePreferences_Init(void* ptr);
bool C_WKWebpagePreferences_AllowsContentJavaScript(void* ptr);
void C_WKWebpagePreferences_SetAllowsContentJavaScript(void* ptr, bool value);
int C_WKWebpagePreferences_PreferredContentMode(void* ptr);
void C_WKWebpagePreferences_SetPreferredContentMode(void* ptr, int value);
