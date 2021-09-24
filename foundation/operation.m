#import "operation.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSOperation.h>

void* C_Operation_Alloc() {
    return [NSOperation alloc];
}

void* C_NSOperation_AllocOperation() {
    NSOperation* result_ = [NSOperation alloc];
    return result_;
}

void* C_NSOperation_Init(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSOperation* result_ = [nSOperation init];
    return result_;
}

void* C_NSOperation_NewOperation() {
    NSOperation* result_ = [NSOperation new];
    return result_;
}

void* C_NSOperation_Autorelease(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSOperation* result_ = [nSOperation autorelease];
    return result_;
}

void* C_NSOperation_Retain(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSOperation* result_ = [nSOperation retain];
    return result_;
}

void C_NSOperation_Start(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation start];
}

void C_NSOperation_Main(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation main];
}

void C_NSOperation_Cancel(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation cancel];
}

void C_NSOperation_AddDependency(void* ptr, void* op) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation addDependency:(NSOperation*)op];
}

void C_NSOperation_RemoveDependency(void* ptr, void* op) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation removeDependency:(NSOperation*)op];
}

void C_NSOperation_WaitUntilFinished(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation waitUntilFinished];
}

bool C_NSOperation_IsCancelled(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isCancelled];
    return result_;
}

bool C_NSOperation_IsExecuting(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isExecuting];
    return result_;
}

bool C_NSOperation_IsFinished(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isFinished];
    return result_;
}

bool C_NSOperation_IsConcurrent(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isConcurrent];
    return result_;
}

bool C_NSOperation_IsAsynchronous(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isAsynchronous];
    return result_;
}

bool C_NSOperation_IsReady(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    BOOL result_ = [nSOperation isReady];
    return result_;
}

void* C_NSOperation_Name(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSString* result_ = [nSOperation name];
    return result_;
}

void C_NSOperation_SetName(void* ptr, void* value) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation setName:(NSString*)value];
}

Array C_NSOperation_Dependencies(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSArray* result_ = [nSOperation dependencies];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

int C_NSOperation_QualityOfService(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSQualityOfService result_ = [nSOperation qualityOfService];
    return result_;
}

void C_NSOperation_SetQualityOfService(void* ptr, int value) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation setQualityOfService:value];
}

int C_NSOperation_QueuePriority(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSOperationQueuePriority result_ = [nSOperation queuePriority];
    return result_;
}

void C_NSOperation_SetQueuePriority(void* ptr, int value) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    [nSOperation setQueuePriority:value];
}

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
