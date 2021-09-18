#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_NavigationResponse_Alloc();

void* C_WKNavigationResponse_AllocNavigationResponse();
void* C_WKNavigationResponse_Init(void* ptr);
void* C_WKNavigationResponse_NewNavigationResponse();
void* C_WKNavigationResponse_Autorelease(void* ptr);
void* C_WKNavigationResponse_Retain(void* ptr);
void* C_WKNavigationResponse_Response(void* ptr);
bool C_WKNavigationResponse_CanShowMIMEType(void* ptr);
bool C_WKNavigationResponse_IsForMainFrame(void* ptr);
