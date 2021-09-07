#import <stdbool.h>
#import <stdlib.h>
#import <stdint.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void WebView_TakeSnapshotWithConfiguration(void* ptr, void* configuration, uintptr_t handler);
void WebView_EvaluateJavaScript(void* ptr, void* javascript, uintptr_t handler);