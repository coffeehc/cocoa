#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Preferences_Alloc();

void* C_WKPreferences_AllocPreferences();
void* C_WKPreferences_Init(void* ptr);
void* C_WKPreferences_NewPreferences();
void* C_WKPreferences_Autorelease(void* ptr);
void* C_WKPreferences_Retain(void* ptr);
double C_WKPreferences_MinimumFontSize(void* ptr);
void C_WKPreferences_SetMinimumFontSize(void* ptr, double value);
bool C_WKPreferences_TabFocusesLinks(void* ptr);
void C_WKPreferences_SetTabFocusesLinks(void* ptr, bool value);
bool C_WKPreferences_JavaScriptCanOpenWindowsAutomatically(void* ptr);
void C_WKPreferences_SetJavaScriptCanOpenWindowsAutomatically(void* ptr, bool value);
bool C_WKPreferences_IsFraudulentWebsiteWarningEnabled(void* ptr);
void C_WKPreferences_SetFraudulentWebsiteWarningEnabled(void* ptr, bool value);
bool C_WKPreferences_TextInteractionEnabled(void* ptr);
void C_WKPreferences_SetTextInteractionEnabled(void* ptr, bool value);
