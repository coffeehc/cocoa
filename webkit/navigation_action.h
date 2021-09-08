#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_NavigationAction_Alloc();

void* C_WKNavigationAction_AllocNavigationAction();
void* C_WKNavigationAction_Init(void* ptr);
void* C_WKNavigationAction_NewNavigationAction();
void* C_WKNavigationAction_Autorelease(void* ptr);
void* C_WKNavigationAction_Retain(void* ptr);
int C_WKNavigationAction_NavigationType(void* ptr);
void* C_WKNavigationAction_Request(void* ptr);
void* C_WKNavigationAction_SourceFrame(void* ptr);
void* C_WKNavigationAction_TargetFrame(void* ptr);
int C_WKNavigationAction_ButtonNumber(void* ptr);
unsigned int C_WKNavigationAction_ModifierFlags(void* ptr);
bool C_WKNavigationAction_ShouldPerformDownload(void* ptr);
