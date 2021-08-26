#import <stdbool.h>
#import <stdlib.h>
#import <stdint.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void TakeSnapshotWithConfiguration(void* ptr, void* configuration, uintptr_t handler);
void EvaluateJavaScript(void* ptr, void* javascript, uintptr_t handler);