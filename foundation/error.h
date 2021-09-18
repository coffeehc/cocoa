#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Error_Alloc();

void* C_NSError_ErrorWithDomain_Code_UserInfo(void* domain, int code, Dictionary dict);
void* C_NSError_InitWithDomain_Code_UserInfo(void* ptr, void* domain, int code, Dictionary dict);
void* C_NSError_AllocError();
void* C_NSError_Init(void* ptr);
void* C_NSError_NewError();
void* C_NSError_Autorelease(void* ptr);
void* C_NSError_Retain(void* ptr);
bool C_NSError_AttemptRecoveryFromError_OptionIndex(void* ptr, void* error, unsigned int recoveryOptionIndex);
int C_NSError_Code(void* ptr);
void* C_NSError_Domain(void* ptr);
Dictionary C_NSError_UserInfo(void* ptr);
void* C_NSError_LocalizedDescription(void* ptr);
Array C_NSError_LocalizedRecoveryOptions(void* ptr);
void* C_NSError_LocalizedRecoverySuggestion(void* ptr);
void* C_NSError_LocalizedFailureReason(void* ptr);
void* C_NSError_RecoveryAttempter(void* ptr);
void* C_NSError_HelpAnchor(void* ptr);
Array C_NSError_UnderlyingErrors(void* ptr);
