#import <Foundation/Foundation.h>
#import "operation_queue.h"

void* C_OperationQueue_Alloc() {
	return [NSOperationQueue alloc];
}

void* C_NSOperationQueue_Init(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	NSOperationQueue* result = [nSOperationQueue init];
	return result;
}

void C_NSOperationQueue_AddOperation(void* ptr, void* op) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue addOperation:(NSOperation*)op];
}

void C_NSOperationQueue_CancelAllOperations(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue cancelAllOperations];
}

void C_NSOperationQueue_WaitUntilAllOperationsAreFinished(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue waitUntilAllOperationsAreFinished];
}

int C_NSOperationQueue_MaxConcurrentOperationCount(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	int result = [nSOperationQueue maxConcurrentOperationCount];
	return result;
}

void C_NSOperationQueue_SetMaxConcurrentOperationCount(void* ptr, int value) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue setMaxConcurrentOperationCount:value];
}

bool C_NSOperationQueue_IsSuspended(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	bool result = [nSOperationQueue isSuspended];
	return result;
}

void C_NSOperationQueue_SetSuspended(void* ptr, bool value) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue setSuspended:value];
}

void* C_NSOperationQueue_Name(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	NSString* result = [nSOperationQueue name];
	return result;
}

void C_NSOperationQueue_SetName(void* ptr, void* value) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	[nSOperationQueue setName:(NSString*)value];
}

void* C_NSOperationQueue_Progress(void* ptr) {
	NSOperationQueue* nSOperationQueue = (NSOperationQueue*)ptr;
	NSProgress* result = [nSOperationQueue progress];
	return result;
}
