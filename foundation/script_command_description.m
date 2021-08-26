#import <Foundation/Foundation.h>
#import "script_command_description.h"

void* C_ScriptCommandDescription_Alloc() {
    return [NSScriptCommandDescription alloc];
}

void* C_NSScriptCommandDescription_InitWithCoder(void* ptr, void* inCoder) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSScriptCommandDescription* result_ = [nSScriptCommandDescription initWithCoder:(NSCoder*)inCoder];
    return result_;
}

bool C_NSScriptCommandDescription_IsOptionalArgumentWithName(void* ptr, void* argumentName) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    BOOL result_ = [nSScriptCommandDescription isOptionalArgumentWithName:(NSString*)argumentName];
    return result_;
}

void* C_NSScriptCommandDescription_TypeForArgumentWithName(void* ptr, void* argumentName) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSString* result_ = [nSScriptCommandDescription typeForArgumentWithName:(NSString*)argumentName];
    return result_;
}

void* C_NSScriptCommandDescription_CreateCommandInstance(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSScriptCommand* result_ = [nSScriptCommandDescription createCommandInstance];
    return result_;
}

void* C_NSScriptCommandDescription_CommandClassName(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSString* result_ = [nSScriptCommandDescription commandClassName];
    return result_;
}

void* C_NSScriptCommandDescription_CommandName(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSString* result_ = [nSScriptCommandDescription commandName];
    return result_;
}

void* C_NSScriptCommandDescription_SuiteName(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSString* result_ = [nSScriptCommandDescription suiteName];
    return result_;
}

Array C_NSScriptCommandDescription_ArgumentNames(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSArray* result_ = [nSScriptCommandDescription argumentNames];
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

void* C_NSScriptCommandDescription_ReturnType(void* ptr) {
    NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
    NSString* result_ = [nSScriptCommandDescription returnType];
    return result_;
}
