#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_ScriptCommandDescription_Alloc();

void* C_NSScriptCommandDescription_InitWithCoder(void* ptr, void* inCoder);
void* C_NSScriptCommandDescription_AllocScriptCommandDescription();
void* C_NSScriptCommandDescription_Autorelease(void* ptr);
void* C_NSScriptCommandDescription_Retain(void* ptr);
bool C_NSScriptCommandDescription_IsOptionalArgumentWithName(void* ptr, void* argumentName);
void* C_NSScriptCommandDescription_TypeForArgumentWithName(void* ptr, void* argumentName);
void* C_NSScriptCommandDescription_CreateCommandInstance(void* ptr);
void* C_NSScriptCommandDescription_CommandClassName(void* ptr);
void* C_NSScriptCommandDescription_CommandName(void* ptr);
void* C_NSScriptCommandDescription_SuiteName(void* ptr);
Array C_NSScriptCommandDescription_ArgumentNames(void* ptr);
void* C_NSScriptCommandDescription_ReturnType(void* ptr);
