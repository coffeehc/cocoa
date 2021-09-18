#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_UserScript_Alloc();

void* C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly(void* ptr, void* source, int injectionTime, bool forMainFrameOnly);
void* C_WKUserScript_InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(void* ptr, void* source, int injectionTime, bool forMainFrameOnly, void* contentWorld);
void* C_WKUserScript_AllocUserScript();
void* C_WKUserScript_Init(void* ptr);
void* C_WKUserScript_NewUserScript();
void* C_WKUserScript_Autorelease(void* ptr);
void* C_WKUserScript_Retain(void* ptr);
void* C_WKUserScript_Source(void* ptr);
int C_WKUserScript_InjectionTime(void* ptr);
bool C_WKUserScript_IsForMainFrameOnly(void* ptr);
