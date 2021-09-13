#import <Foundation/Foundation.h>
#import "extension_context.h"

void* C_ExtensionContext_Alloc() {
    return [NSExtensionContext alloc];
}

void* C_NSExtensionContext_AllocExtensionContext() {
    NSExtensionContext* result_ = [NSExtensionContext alloc];
    return result_;
}

void* C_NSExtensionContext_Init(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSExtensionContext* result_ = [nSExtensionContext init];
    return result_;
}

void* C_NSExtensionContext_NewExtensionContext() {
    NSExtensionContext* result_ = [NSExtensionContext new];
    return result_;
}

void* C_NSExtensionContext_Autorelease(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSExtensionContext* result_ = [nSExtensionContext autorelease];
    return result_;
}

void* C_NSExtensionContext_Retain(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSExtensionContext* result_ = [nSExtensionContext retain];
    return result_;
}

void C_NSExtensionContext_CancelRequestWithError(void* ptr, void* error) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    [nSExtensionContext cancelRequestWithError:(NSError*)error];
}

Array C_NSExtensionContext_InputItems(void* ptr) {
    NSExtensionContext* nSExtensionContext = (NSExtensionContext*)ptr;
    NSArray* result_ = [nSExtensionContext inputItems];
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
