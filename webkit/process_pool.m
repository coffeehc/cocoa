#import "process_pool.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKProcessPool.h>

void* C_ProcessPool_Alloc() {
    return [WKProcessPool alloc];
}

void* C_WKProcessPool_AllocProcessPool() {
    WKProcessPool* result_ = [WKProcessPool alloc];
    return result_;
}

void* C_WKProcessPool_Init(void* ptr) {
    WKProcessPool* wKProcessPool = (WKProcessPool*)ptr;
    WKProcessPool* result_ = [wKProcessPool init];
    return result_;
}

void* C_WKProcessPool_NewProcessPool() {
    WKProcessPool* result_ = [WKProcessPool new];
    return result_;
}

void* C_WKProcessPool_Autorelease(void* ptr) {
    WKProcessPool* wKProcessPool = (WKProcessPool*)ptr;
    WKProcessPool* result_ = [wKProcessPool autorelease];
    return result_;
}

void* C_WKProcessPool_Retain(void* ptr) {
    WKProcessPool* wKProcessPool = (WKProcessPool*)ptr;
    WKProcessPool* result_ = [wKProcessPool retain];
    return result_;
}
