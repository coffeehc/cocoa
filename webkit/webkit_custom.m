#import <WebKit/WebKit.h>
#include "_cgo_export.h"
#import "webkit_custom.h"

void WebView_TakeSnapshotWithConfiguration(void* ptr, void* configuration, uintptr_t handler) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView takeSnapshotWithConfiguration:(WKSnapshotConfiguration*)configuration completionHandler:^(NSImage * _Nullable image, NSError * _Nullable error) {
        callTakeSnapshotHandler(handler, image, error);
    }];
}

void WebView_EvaluateJavaScript(void* ptr, void* javascript, uintptr_t handler) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView evaluateJavaScript:(NSString*)javascript completionHandler:^(id value, NSError * _Nullable error) {
        callEvaluateJavaScriptHandler(handler, value, error);
    }];
}