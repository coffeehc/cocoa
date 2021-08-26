#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_NavigationResponse_Alloc();

void* C_WKNavigationResponse_Init(void* ptr);
void* C_WKNavigationResponse_Response(void* ptr);
bool C_WKNavigationResponse_CanShowMIMEType(void* ptr);
bool C_WKNavigationResponse_IsForMainFrame(void* ptr);
