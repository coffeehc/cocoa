#import <WebKit/WebKit.h>
#import "download.h"

void* C_Download_Alloc() {
    return [WKDownload alloc];
}

void* C_WKDownload_Init(void* ptr) {
    WKDownload* wKDownload = (WKDownload*)ptr;
    WKDownload* result_ = [wKDownload init];
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
