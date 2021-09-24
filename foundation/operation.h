#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Operation_Alloc();

void* C_NSOperation_AllocOperation();
void* C_NSOperation_Init(void* ptr);
void* C_NSOperation_NewOperation();
void* C_NSOperation_Autorelease(void* ptr);
void* C_NSOperation_Retain(void* ptr);
void C_NSOperation_Start(void* ptr);
void C_NSOperation_Main(void* ptr);
void C_NSOperation_Cancel(void* ptr);
void C_NSOperation_AddDependency(void* ptr, void* op);
void C_NSOperation_RemoveDependency(void* ptr, void* op);
void C_NSOperation_WaitUntilFinished(void* ptr);
bool C_NSOperation_IsCancelled(void* ptr);
bool C_NSOperation_IsExecuting(void* ptr);
bool C_NSOperation_IsFinished(void* ptr);
bool C_NSOperation_IsConcurrent(void* ptr);
bool C_NSOperation_IsAsynchronous(void* ptr);
bool C_NSOperation_IsReady(void* ptr);
void* C_NSOperation_Name(void* ptr);
void C_NSOperation_SetName(void* ptr, void* value);
Array C_NSOperation_Dependencies(void* ptr);
int C_NSOperation_QualityOfService(void* ptr);
void C_NSOperation_SetQualityOfService(void* ptr, int value);
int C_NSOperation_QueuePriority(void* ptr);
void C_NSOperation_SetQueuePriority(void* ptr, int value);

void* C_OperationQueue_Alloc();

void* C_NSOperationQueue_AllocOperationQueue();
void* C_NSOperationQueue_Init(void* ptr);
void* C_NSOperationQueue_NewOperationQueue();
void* C_NSOperationQueue_Autorelease(void* ptr);
void* C_NSOperationQueue_Retain(void* ptr);
void C_NSOperationQueue_AddOperation(void* ptr, void* op);
void C_NSOperationQueue_AddOperations_WaitUntilFinished(void* ptr, Array ops, bool wait);
void C_NSOperationQueue_CancelAllOperations(void* ptr);
void C_NSOperationQueue_WaitUntilAllOperationsAreFinished(void* ptr);
void* C_NSOperationQueue_OperationQueue_MainQueue();
void* C_NSOperationQueue_OperationQueue_CurrentQueue();
int C_NSOperationQueue_QualityOfService(void* ptr);
void C_NSOperationQueue_SetQualityOfService(void* ptr, int value);
int C_NSOperationQueue_MaxConcurrentOperationCount(void* ptr);
void C_NSOperationQueue_SetMaxConcurrentOperationCount(void* ptr, int value);
bool C_NSOperationQueue_IsSuspended(void* ptr);
void C_NSOperationQueue_SetSuspended(void* ptr, bool value);
void* C_NSOperationQueue_Name(void* ptr);
void C_NSOperationQueue_SetName(void* ptr, void* value);
void* C_NSOperationQueue_Progress(void* ptr);
