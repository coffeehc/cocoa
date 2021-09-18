#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

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
