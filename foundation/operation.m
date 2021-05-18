#import <Foundation/Foundation.h>
#import "operation.h"

void* C_Operation_Alloc() {
    return [NSOperation alloc];
}

void* C_NSOperation_Init(void* ptr) {
    NSOperation* nSOperation = (NSOperation*)ptr;
    NSOperation* result_ = [nSOperation init];
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
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
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
