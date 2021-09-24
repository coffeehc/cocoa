#import "navigation_response.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKNavigationResponse.h>

void* C_NavigationResponse_Alloc() {
    return [WKNavigationResponse alloc];
}

void* C_WKNavigationResponse_AllocNavigationResponse() {
    WKNavigationResponse* result_ = [WKNavigationResponse alloc];
    return result_;
}

void* C_WKNavigationResponse_Init(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    WKNavigationResponse* result_ = [wKNavigationResponse init];
    return result_;
}

void* C_WKNavigationResponse_NewNavigationResponse() {
    WKNavigationResponse* result_ = [WKNavigationResponse new];
    return result_;
}

void* C_WKNavigationResponse_Autorelease(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    WKNavigationResponse* result_ = [wKNavigationResponse autorelease];
    return result_;
}

void* C_WKNavigationResponse_Retain(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    WKNavigationResponse* result_ = [wKNavigationResponse retain];
    return result_;
}

void* C_WKNavigationResponse_Response(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    NSURLResponse* result_ = [wKNavigationResponse response];
    return result_;
}

bool C_WKNavigationResponse_CanShowMIMEType(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    BOOL result_ = [wKNavigationResponse canShowMIMEType];
    return result_;
}

bool C_WKNavigationResponse_IsForMainFrame(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    BOOL result_ = [wKNavigationResponse isForMainFrame];
    return result_;
}
