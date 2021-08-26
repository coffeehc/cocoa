#import <Foundation/Foundation.h>
#import "error.h"

void* C_Error_Alloc() {
    return [NSError alloc];
}

void* C_NSError_InitWithDomain_Code_UserInfo(void* ptr, void* domain, int code, Dictionary dict) {
    NSError* nSError = (NSError*)ptr;
    NSMutableDictionary* objcDict = [[NSMutableDictionary alloc] initWithCapacity: dict.len];
    if (dict.len > 0) {
    	void** dictKeyData = (void**)dict.key_data;
    	void** dictValueData = (void**)dict.value_data;
    	for (int i = 0; i < dict.len; i++) {
    		void* kp = dictKeyData[i];
    		void* vp = dictValueData[i];
    		[objcDict setObject:(NSErrorUserInfoKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSError* result_ = [nSError initWithDomain:(NSString*)domain code:code userInfo:objcDict];
    return result_;
}

void* C_NSError_Init(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSError* result_ = [nSError init];
    return result_;
}

void* C_NSError_ErrorWithDomain_Code_UserInfo(void* domain, int code, Dictionary dict) {
    NSMutableDictionary* objcDict = [[NSMutableDictionary alloc] initWithCapacity: dict.len];
    if (dict.len > 0) {
    	void** dictKeyData = (void**)dict.key_data;
    	void** dictValueData = (void**)dict.value_data;
    	for (int i = 0; i < dict.len; i++) {
    		void* kp = dictKeyData[i];
    		void* vp = dictValueData[i];
    		[objcDict setObject:(NSErrorUserInfoKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSError* result_ = [NSError errorWithDomain:(NSString*)domain code:code userInfo:objcDict];
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

Dictionary C_NSError_UserInfo(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSDictionary* result_ = [nSError userInfo];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSErrorUserInfoKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSError_LocalizedDescription(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSString* result_ = [nSError localizedDescription];
    return result_;
}

Array C_NSError_LocalizedRecoveryOptions(void* ptr) {
    NSError* nSError = (NSError*)ptr;
    NSArray* result_ = [nSError localizedRecoveryOptions];
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
