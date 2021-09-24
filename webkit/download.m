#import "download.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKDownload.h>

void* C_Download_Alloc() {
    return [WKDownload alloc];
}

void* C_WKDownload_AllocDownload() {
    WKDownload* result_ = [WKDownload alloc];
    return result_;
}

void* C_WKDownload_Init(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    WKDownload* result_ = [wKDownload init];
    return result_;
}

void* C_WKDownload_NewDownload() {
    WKDownload* result_ = [WKDownload new];
    return result_;
}

void* C_WKDownload_Autorelease(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    WKDownload* result_ = [wKDownload autorelease];
    return result_;
}

void* C_WKDownload_Retain(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    WKDownload* result_ = [wKDownload retain];
    return result_;
}

void* C_WKDownload_Delegate(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    id result_ = [wKDownload delegate];
    return result_;
}

void C_WKDownload_SetDelegate(void* ptr, void* value) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    [wKDownload setDelegate:(id)value];
}

void* C_WKDownload_OriginalRequest(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    NSURLRequest* result_ = [wKDownload originalRequest];
    return result_;
}

void* C_WKDownload_WebView(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    WKWebView* result_ = [wKDownload webView];
    return result_;
}
