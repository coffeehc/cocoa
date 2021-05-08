#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* C_Error_Alloc();

void* C_NSError_Init(void* ptr);
void C_NSError_SetUserInfoValueProviderForDomain_Provider(void* errorDomain, void* provider);
int C_NSError_Code(void* ptr);
void* C_NSError_Domain(void* ptr);
void* C_NSError_LocalizedDescription(void* ptr);
void* C_NSError_LocalizedRecoverySuggestion(void* ptr);
void* C_NSError_LocalizedFailureReason(void* ptr);
void* C_NSError_RecoveryAttempter(void* ptr);
void* C_NSError_HelpAnchor(void* ptr);
