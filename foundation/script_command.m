#import <Foundation/Foundation.h>
#import "script_command.h"

void* C_ScriptCommand_Alloc() {
    return [NSScriptCommand alloc];
}

void* C_NSScriptCommand_InitWithCommandDescription(void* ptr, void* commandDef) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptCommand* result_ = [nSScriptCommand initWithCommandDescription:(NSScriptCommandDescription*)commandDef];
    return result_;
}

void* C_NSScriptCommand_InitWithCoder(void* ptr, void* inCoder) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptCommand* result_ = [nSScriptCommand initWithCoder:(NSCoder*)inCoder];
    return result_;
}

void* C_NSScriptCommand_AllocScriptCommand() {
    NSScriptCommand* result_ = [NSScriptCommand alloc];
    return result_;
}

void* C_NSScriptCommand_Autorelease(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptCommand* result_ = [nSScriptCommand autorelease];
    return result_;
}

void* C_NSScriptCommand_Retain(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptCommand* result_ = [nSScriptCommand retain];
    return result_;
}

void* C_NSScriptCommand_ScriptCommand_CurrentCommand() {
    NSScriptCommand* result_ = [NSScriptCommand currentCommand];
    return result_;
}

void* C_NSScriptCommand_ExecuteCommand(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    id result_ = [nSScriptCommand executeCommand];
    return result_;
}

void* C_NSScriptCommand_PerformDefaultImplementation(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    id result_ = [nSScriptCommand performDefaultImplementation];
    return result_;
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
    NSAppleEventDescriptor* result_ = [nSScriptCommand appleEvent];
    return result_;
}

void* C_NSScriptCommand_EvaluatedReceivers(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    id result_ = [nSScriptCommand evaluatedReceivers];
    return result_;
}

void* C_NSScriptCommand_ReceiversSpecifier(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptCommand receiversSpecifier];
    return result_;
}

void C_NSScriptCommand_SetReceiversSpecifier(void* ptr, void* value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setReceiversSpecifier:(NSScriptObjectSpecifier*)value];
}

Dictionary C_NSScriptCommand_Arguments(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSDictionary* result_ = [nSScriptCommand arguments];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
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

void C_NSScriptCommand_SetArguments(void* ptr, Dictionary value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSScriptCommand setArguments:objcValue];
}

Dictionary C_NSScriptCommand_EvaluatedArguments(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSDictionary* result_ = [nSScriptCommand evaluatedArguments];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
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

void* C_NSScriptCommand_DirectParameter(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    id result_ = [nSScriptCommand directParameter];
    return result_;
}

void C_NSScriptCommand_SetDirectParameter(void* ptr, void* value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setDirectParameter:(id)value];
}

void* C_NSScriptCommand_CommandDescription(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSScriptCommandDescription* result_ = [nSScriptCommand commandDescription];
    return result_;
}

void* C_NSScriptCommand_ScriptErrorExpectedTypeDescriptor(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSAppleEventDescriptor* result_ = [nSScriptCommand scriptErrorExpectedTypeDescriptor];
    return result_;
}

void C_NSScriptCommand_SetScriptErrorExpectedTypeDescriptor(void* ptr, void* value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setScriptErrorExpectedTypeDescriptor:(NSAppleEventDescriptor*)value];
}

int C_NSScriptCommand_ScriptErrorNumber(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSInteger result_ = [nSScriptCommand scriptErrorNumber];
    return result_;
}

void C_NSScriptCommand_SetScriptErrorNumber(void* ptr, int value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setScriptErrorNumber:value];
}

void* C_NSScriptCommand_ScriptErrorOffendingObjectDescriptor(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSAppleEventDescriptor* result_ = [nSScriptCommand scriptErrorOffendingObjectDescriptor];
    return result_;
}

void C_NSScriptCommand_SetScriptErrorOffendingObjectDescriptor(void* ptr, void* value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setScriptErrorOffendingObjectDescriptor:(NSAppleEventDescriptor*)value];
}

void* C_NSScriptCommand_ScriptErrorString(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    NSString* result_ = [nSScriptCommand scriptErrorString];
    return result_;
}

void C_NSScriptCommand_SetScriptErrorString(void* ptr, void* value) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    [nSScriptCommand setScriptErrorString:(NSString*)value];
}

bool C_NSScriptCommand_IsWellFormed(void* ptr) {
    NSScriptCommand* nSScriptCommand = (NSScriptCommand*)ptr;
    BOOL result_ = [nSScriptCommand isWellFormed];
    return result_;
}
