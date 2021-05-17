#import <WebKit/WebKit.h>
#import "process_pool.h"

void* C_ProcessPool_Alloc() {
    return [WKProcessPool alloc];
}

void* C_WKProcessPool_Init(void* ptr) {
    WKProcessPool* wKProcessPool = (WKProcessPool*)ptr;
    WKProcessPool* result_ = [wKProcessPool init];
    return result_;
}
