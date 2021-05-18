#import <Foundation/Foundation.h>
#import "error.h"

void* C_Error_Alloc() {
    return [NSError alloc];
}

void* C_NSError_Init(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSError* result_ = [nSError init];
    return result_;
}

bool C_NSError_AttemptRecoveryFromError_OptionIndex(void* ptr, void* error, unsigned int recoveryOptionIndex) {
    NSError* nSError = (NSError*)ptr;
    BOOL result_ = [nSError attemptRecoveryFromError:(NSError*)error optionIndex:recoveryOptionIndex];
    return result_;
}

int C_NSError_Code(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSInteger result_ = [nSError code];
    return result_;
}

void* C_NSError_Domain(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSErrorDomain result_ = [nSError domain];
    return result_;
}

void* C_NSError_LocalizedDescription(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSString* result_ = [nSError localizedDescription];
    return result_;
}

Array C_NSError_LocalizedRecoveryOptions(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSArray* result_ = [nSError localizedRecoveryOptions];
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

void* C_NSError_LocalizedRecoverySuggestion(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSString* result_ = [nSError localizedRecoverySuggestion];
    return result_;
}

void* C_NSError_LocalizedFailureReason(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSString* result_ = [nSError localizedFailureReason];
    return result_;
}

void* C_NSError_RecoveryAttempter(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    id result_ = [nSError recoveryAttempter];
    return result_;
}

void* C_NSError_HelpAnchor(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSString* result_ = [nSError helpAnchor];
    return result_;
}

Array C_NSError_UnderlyingErrors(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSArray* result_ = [nSError underlyingErrors];
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
