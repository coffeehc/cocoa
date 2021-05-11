#import <Foundation/Foundation.h>
#import "operation.h"

void* C_Operation_Alloc() {
	return [NSOperation alloc];
}

void* C_NSOperation_Init(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	NSOperation* result = [nSOperation init];
	return result;
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
	bool result = [nSOperation isCancelled];
	return result;
}

bool C_NSOperation_IsExecuting(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	bool result = [nSOperation isExecuting];
	return result;
}

bool C_NSOperation_IsFinished(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	bool result = [nSOperation isFinished];
	return result;
}

bool C_NSOperation_IsConcurrent(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	bool result = [nSOperation isConcurrent];
	return result;
}

bool C_NSOperation_IsAsynchronous(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	bool result = [nSOperation isAsynchronous];
	return result;
}

bool C_NSOperation_IsReady(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	bool result = [nSOperation isReady];
	return result;
}

void* C_NSOperation_Name(void* ptr) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	NSString* result = [nSOperation name];
	return result;
}

void C_NSOperation_SetName(void* ptr, void* value) {
	NSOperation* nSOperation = (NSOperation*)ptr;
	[nSOperation setName:(NSString*)value];
}
