#import "frame_info.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKFrameInfo.h>

void* C_FrameInfo_Alloc() {
    return [WKFrameInfo alloc];
}

void* C_WKFrameInfo_AllocFrameInfo() {
    WKFrameInfo* result_ = [WKFrameInfo alloc];
    return result_;
}

void* C_WKFrameInfo_Init(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    WKFrameInfo* result_ = [wKFrameInfo init];
    return result_;
}

void* C_WKFrameInfo_NewFrameInfo() {
    WKFrameInfo* result_ = [WKFrameInfo new];
    return result_;
}

void* C_WKFrameInfo_Autorelease(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    WKFrameInfo* result_ = [wKFrameInfo autorelease];
    return result_;
}

void* C_WKFrameInfo_Retain(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    WKFrameInfo* result_ = [wKFrameInfo retain];
    return result_;
}

bool C_WKFrameInfo_IsMainFrame(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    BOOL result_ = [wKFrameInfo isMainFrame];
    return result_;
}

void* C_WKFrameInfo_Request(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    NSURLRequest* result_ = [wKFrameInfo request];
    return result_;
}

void* C_WKFrameInfo_SecurityOrigin(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    WKSecurityOrigin* result_ = [wKFrameInfo securityOrigin];
    return result_;
}

void* C_WKFrameInfo_WebView(void* ptr) {
    WKFrameInfo* wKFrameInfo = (WKFrameInfo*)ptr;
    WKWebView* result_ = [wKFrameInfo webView];
    return result_;
}
