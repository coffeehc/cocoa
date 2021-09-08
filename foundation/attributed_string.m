#import <Foundation/Foundation.h>
#import "attributed_string.h"

void* C_AttributedString_Alloc() {
    return [NSAttributedString alloc];
}

void* C_NSAttributedString_AllocAttributedString() {
    NSAttributedString* result_ = [NSAttributedString alloc];
    return result_;
}

void* C_NSAttributedString_Init(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString init];
    return result_;
}

void* C_NSAttributedString_NewAttributedString() {
    NSAttributedString* result_ = [NSAttributedString new];
    return result_;
}

void* C_NSAttributedString_Autorelease(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString autorelease];
    return result_;
}

void* C_NSAttributedString_Retain(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString retain];
    return result_;
}

void* C_NSAttributedString_String(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSString* result_ = [nSAttributedString string];
    return result_;
}

unsigned int C_NSAttributedString_Length(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSUInteger result_ = [nSAttributedString length];
    return result_;
}

Array C_NSAttributedString_AttributedString_TextTypes() {
    NSArray* result_ = [NSAttributedString textTypes];
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

Array C_NSAttributedString_AttributedString_TextUnfilteredTypes() {
    NSArray* result_ = [NSAttributedString textUnfilteredTypes];
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
