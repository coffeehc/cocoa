#import <dispatch/dispatch.h>
#import "base.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


void Dispatch_MainQueueAsync(uintptr_t ptr) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(ptr);
    });
}

void Run_WithAutoreleasePool(uintptr_t ptr) {
    @autoreleasepool {
        runTask(ptr);
    }
}
