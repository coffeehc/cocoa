#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_WebpagePreferences_Alloc();

void* C_WKWebpagePreferences_AllocWebpagePreferences();
void* C_WKWebpagePreferences_Init(void* ptr);
void* C_WKWebpagePreferences_NewWebpagePreferences();
void* C_WKWebpagePreferences_Autorelease(void* ptr);
void* C_WKWebpagePreferences_Retain(void* ptr);
bool C_WKWebpagePreferences_AllowsContentJavaScript(void* ptr);
void C_WKWebpagePreferences_SetAllowsContentJavaScript(void* ptr, bool value);
int C_WKWebpagePreferences_PreferredContentMode(void* ptr);
void C_WKWebpagePreferences_SetPreferredContentMode(void* ptr, int value);
