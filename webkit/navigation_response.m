#import <WebKit/WebKit.h>
#import "navigation_response.h"

void* C_NavigationResponse_Alloc() {
    return [WKNavigationResponse alloc];
}

void* C_WKNavigationResponse_Init(void* ptr) {
    WKNavigationResponse* wKNavigationResponse = (WKNavigationResponse*)ptr;
    WKNavigationResponse* result_ = [wKNavigationResponse init];
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
