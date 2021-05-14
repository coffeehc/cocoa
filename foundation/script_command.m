#import <Foundation/Foundation.h>
#import "script_command.h"

void* C_ScriptCommand_Alloc() {
	return [NSScriptCommand alloc];
}

void* C_NSScriptCommand_InitWithCommandDescription(void* ptr, void* commandDef) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSScriptCommand* result = [nSScriptCommand initWithCommandDescription:(NSScriptCommandDescription*)commandDef];
	return result;
}

void* C_NSScriptCommand_InitWithCoder(void* ptr, void* inCoder) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSScriptCommand* result = [nSScriptCommand initWithCoder:(NSCoder*)inCoder];
	return result;
}

void* C_NSScriptCommand_ScriptCommandCurrentCommand() {
	NSScriptCommand* result = [NSScriptCommand currentCommand];
	return result;
}

void* C_NSScriptCommand_ExecuteCommand(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	id result = [nSScriptCommand executeCommand];
	return result;
}

void* C_NSScriptCommand_PerformDefaultImplementation(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	id result = [nSScriptCommand performDefaultImplementation];
	return result;
}

void C_NSScriptCommand_SuspendExecution(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand suspendExecution];
}

void C_NSScriptCommand_ResumeExecutionWithResult(void* ptr, void* result) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand resumeExecutionWithResult:(id)result];
}

void* C_NSScriptCommand_AppleEvent(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSAppleEventDescriptor* result = [nSScriptCommand appleEvent];
	return result;
}

void* C_NSScriptCommand_EvaluatedReceivers(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	id result = [nSScriptCommand evaluatedReceivers];
	return result;
}

void* C_NSScriptCommand_ReceiversSpecifier(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptCommand receiversSpecifier];
	return result;
}

void C_NSScriptCommand_SetReceiversSpecifier(void* ptr, void* value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setReceiversSpecifier:(NSScriptObjectSpecifier*)value];
}

void* C_NSScriptCommand_DirectParameter(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	id result = [nSScriptCommand directParameter];
	return result;
}

void C_NSScriptCommand_SetDirectParameter(void* ptr, void* value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setDirectParameter:(id)value];
}

void* C_NSScriptCommand_CommandDescription(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSScriptCommandDescription* result = [nSScriptCommand commandDescription];
	return result;
}

void* C_NSScriptCommand_ScriptErrorExpectedTypeDescriptor(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSAppleEventDescriptor* result = [nSScriptCommand scriptErrorExpectedTypeDescriptor];
	return result;
}

void C_NSScriptCommand_SetScriptErrorExpectedTypeDescriptor(void* ptr, void* value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setScriptErrorExpectedTypeDescriptor:(NSAppleEventDescriptor*)value];
}

int C_NSScriptCommand_ScriptErrorNumber(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	int result = [nSScriptCommand scriptErrorNumber];
	return result;
}

void C_NSScriptCommand_SetScriptErrorNumber(void* ptr, int value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setScriptErrorNumber:value];
}

void* C_NSScriptCommand_ScriptErrorOffendingObjectDescriptor(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSAppleEventDescriptor* result = [nSScriptCommand scriptErrorOffendingObjectDescriptor];
	return result;
}

void C_NSScriptCommand_SetScriptErrorOffendingObjectDescriptor(void* ptr, void* value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setScriptErrorOffendingObjectDescriptor:(NSAppleEventDescriptor*)value];
}

void* C_NSScriptCommand_ScriptErrorString(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	NSString* result = [nSScriptCommand scriptErrorString];
	return result;
}

void C_NSScriptCommand_SetScriptErrorString(void* ptr, void* value) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	[nSScriptCommand setScriptErrorString:(NSString*)value];
}

bool C_NSScriptCommand_IsWellFormed(void* ptr) {
	NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
	bool result = [nSScriptCommand isWellFormed];
	return result;
}
