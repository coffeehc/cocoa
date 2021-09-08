#import <Foundation/Foundation.h>
#import "class_description.h"

void* C_ClassDescription_Alloc() {
    return [NSClassDescription alloc];
}

void* C_NSClassDescription_AllocClassDescription() {
    NSClassDescription* result_ = [NSClassDescription alloc];
    return result_;
}

void* C_NSClassDescription_Init(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSClassDescription* result_ = [nSClassDescription init];
    return result_;
}

void* C_NSClassDescription_NewClassDescription() {
    NSClassDescription* result_ = [NSClassDescription new];
    return result_;
}

void* C_NSClassDescription_Autorelease(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSClassDescription* result_ = [nSClassDescription autorelease];
    return result_;
}

void* C_NSClassDescription_Retain(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSClassDescription* result_ = [nSClassDescription retain];
    return result_;
}

void C_NSClassDescription_InvalidateClassDescriptionCache() {
    [NSClassDescription invalidateClassDescriptionCache];
}

void* C_NSClassDescription_InverseForRelationshipKey(void* ptr, void* relationshipKey) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSString* result_ = [nSClassDescription inverseForRelationshipKey:(NSString*)relationshipKey];
    return result_;
}

Array C_NSClassDescription_AttributeKeys(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSArray* result_ = [nSClassDescription attributeKeys];
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

Array C_NSClassDescription_ToManyRelationshipKeys(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSArray* result_ = [nSClassDescription toManyRelationshipKeys];
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

Array C_NSClassDescription_ToOneRelationshipKeys(void* ptr) {
    NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
    NSArray* result_ = [nSClassDescription toOneRelationshipKeys];
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
