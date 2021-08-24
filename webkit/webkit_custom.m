#import <WebKit/WebKit.h>
#include "_cgo_export.h"
#import "webkit_custom.h"

// custom begin
void TakeSnapshotWithConfiguration(void* ptr, void* configuration, uintptr_t handler) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView takeSnapshotWithConfiguration:(WKSnapshotConfiguration*)configuration completionHandler:^(NSImage * _Nullable image, NSError * _Nullable error) {
        callTakeSnapshotWithConfiguration(handler, image, error);
        deleteTakeSnapshotWithConfiguration(handler);
    }];
}
// custom end