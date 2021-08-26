#import <WebKit/WebKit.h>
#include "_cgo_export.h"
#import "webkit_custom.h"

void TakeSnapshotWithConfiguration(void* ptr, void* configuration, uintptr_t handler) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView takeSnapshotWithConfiguration:(WKSnapshotConfiguration*)configuration completionHandler:^(NSImage * _Nullable image, NSError * _Nullable error) {
        callTakeSnapshotWithConfiguration(handler, image, error);
        deleteHandle(handler);
    }];
}

void EvaluateJavaScript(void* ptr, void* javascript, uintptr_t handler) {
        WKWebView* wKWebView = (WKWebView*)ptr;
        [wKWebView evaluateJavaScript:(NSString*)javascript completionHandler:^(id value, NSError * _Nullable error) {
            callEvaluateJavaScript(handler, value, error);
            deleteHandle(handler);
        }];
}