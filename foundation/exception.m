#import "exception.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSException.h>

void* C_Exception_Alloc() {
    return [NSException alloc];
}

void* C_NSException_AllocException() {
    NSException* result_ = [NSException alloc];
    return result_;
}

void* C_NSException_Init(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSException* result_ = [nSException init];
    return result_;
}

void* C_NSException_NewException() {
    NSException* result_ = [NSException new];
    return result_;
}

void* C_NSException_Autorelease(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSException* result_ = [nSException autorelease];
    return result_;
}

void* C_NSException_Retain(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSException* result_ = [nSException retain];
    return result_;
}

void C_NSException_Raise(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    [nSException raise];
}

void* C_NSException_Name(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSExceptionName result_ = [nSException name];
    return result_;
}

void* C_NSException_Reason(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSString* result_ = [nSException reason];
    return result_;
}

Array C_NSException_CallStackReturnAddresses(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSArray* result_ = [nSException callStackReturnAddresses];
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

Array C_NSException_CallStackSymbols(void* ptr) {
    NSException* nSException = (NSException*)ptr;
    NSArray* result_ = [nSException callStackSymbols];
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
