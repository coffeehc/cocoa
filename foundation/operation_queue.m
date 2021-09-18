#import "operation_queue.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSOperation.h>

void* C_OperationQueue_Alloc() {
    return [NSOperationQueue alloc];
}

void* C_NSOperationQueue_AllocOperationQueue() {
    NSOperationQueue* result_ = [NSOperationQueue alloc];
    return result_;
}

void* C_NSOperationQueue_Init(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSOperationQueue* result_ = [nSOperationQueue init];
    return result_;
}

void* C_NSOperationQueue_NewOperationQueue() {
    NSOperationQueue* result_ = [NSOperationQueue new];
    return result_;
}

void* C_NSOperationQueue_Autorelease(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSOperationQueue* result_ = [nSOperationQueue autorelease];
    return result_;
}

void* C_NSOperationQueue_Retain(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSOperationQueue* result_ = [nSOperationQueue retain];
    return result_;
}

void C_NSOperationQueue_AddOperation(void* ptr, void* op) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue addOperation:(NSOperation*)op];
}

void C_NSOperationQueue_AddOperations_WaitUntilFinished(void* ptr, Array ops, bool wait) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSMutableArray* objcOps = [NSMutableArray arrayWithCapacity:ops.len];
    if (ops.len > 0) {
    	void** opsData = (void**)ops.data;
    	for (int i = 0; i < ops.len; i++) {
    		void* p = opsData[i];
    		[objcOps addObject:(NSOperation*)p];
    	}
    }
    [nSOperationQueue addOperations:objcOps waitUntilFinished:wait];
}

void C_NSOperationQueue_CancelAllOperations(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue cancelAllOperations];
}

void C_NSOperationQueue_WaitUntilAllOperationsAreFinished(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue waitUntilAllOperationsAreFinished];
}

void* C_NSOperationQueue_OperationQueue_MainQueue() {
    NSOperationQueue* result_ = [NSOperationQueue mainQueue];
    return result_;
}

void* C_NSOperationQueue_OperationQueue_CurrentQueue() {
    NSOperationQueue* result_ = [NSOperationQueue currentQueue];
    return result_;
}

int C_NSOperationQueue_QualityOfService(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSQualityOfService result_ = [nSOperationQueue qualityOfService];
    return result_;
}

void C_NSOperationQueue_SetQualityOfService(void* ptr, int value) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue setQualityOfService:value];
}

int C_NSOperationQueue_MaxConcurrentOperationCount(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSInteger result_ = [nSOperationQueue maxConcurrentOperationCount];
    return result_;
}

void C_NSOperationQueue_SetMaxConcurrentOperationCount(void* ptr, int value) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue setMaxConcurrentOperationCount:value];
}

bool C_NSOperationQueue_IsSuspended(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    BOOL result_ = [nSOperationQueue isSuspended];
    return result_;
}

void C_NSOperationQueue_SetSuspended(void* ptr, bool value) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue setSuspended:value];
}

void* C_NSOperationQueue_Name(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSString* result_ = [nSOperationQueue name];
    return result_;
}

void C_NSOperationQueue_SetName(void* ptr, void* value) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    [nSOperationQueue setName:(NSString*)value];
}

void* C_NSOperationQueue_Progress(void* ptr) {
    NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
    NSProgress* result_ = [nSOperationQueue progress];
    return result_;
}
