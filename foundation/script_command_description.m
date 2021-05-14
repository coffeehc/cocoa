#import <Foundation/Foundation.h>
#import "script_command_description.h"

void* C_ScriptCommandDescription_Alloc() {
	return [NSScriptCommandDescription alloc];
}

void* C_NSScriptCommandDescription_InitWithCoder(void* ptr, void* inCoder) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSScriptCommandDescription* result = [nSScriptCommandDescription initWithCoder:(NSCoder*)inCoder];
	return result;
}

bool C_NSScriptCommandDescription_IsOptionalArgumentWithName(void* ptr, void* argumentName) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	bool result = [nSScriptCommandDescription isOptionalArgumentWithName:(NSString*)argumentName];
	return result;
}

void* C_NSScriptCommandDescription_TypeForArgumentWithName(void* ptr, void* argumentName) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSString* result = [nSScriptCommandDescription typeForArgumentWithName:(NSString*)argumentName];
	return result;
}

void* C_NSScriptCommandDescription_CreateCommandInstance(void* ptr) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSScriptCommand* result = [nSScriptCommandDescription createCommandInstance];
	return result;
}

void* C_NSScriptCommandDescription_CommandClassName(void* ptr) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSString* result = [nSScriptCommandDescription commandClassName];
	return result;
}

void* C_NSScriptCommandDescription_CommandName(void* ptr) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSString* result = [nSScriptCommandDescription commandName];
	return result;
}

void* C_NSScriptCommandDescription_SuiteName(void* ptr) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSString* result = [nSScriptCommandDescription suiteName];
	return result;
}

void* C_NSScriptCommandDescription_ReturnType(void* ptr) {
	NSScriptCommandDescription* nSScriptCommandDescription = (NSScriptCommandDescription*)ptr;
	NSString* result = [nSScriptCommandDescription returnType];
	return result;
}
