#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_OperationQueue_Alloc();

void* C_NSOperationQueue_Init(void* ptr);
void C_NSOperationQueue_AddOperation(void* ptr, void* op);
void C_NSOperationQueue_CancelAllOperations(void* ptr);
void C_NSOperationQueue_WaitUntilAllOperationsAreFinished(void* ptr);
int C_NSOperationQueue_MaxConcurrentOperationCount(void* ptr);
void C_NSOperationQueue_SetMaxConcurrentOperationCount(void* ptr, int value);
bool C_NSOperationQueue_IsSuspended(void* ptr);
void C_NSOperationQueue_SetSuspended(void* ptr, bool value);
void* C_NSOperationQueue_Name(void* ptr);
void C_NSOperationQueue_SetName(void* ptr, void* value);
void* C_NSOperationQueue_Progress(void* ptr);
