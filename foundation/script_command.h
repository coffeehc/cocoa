#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ScriptCommand_Alloc();

void* C_NSScriptCommand_InitWithCommandDescription(void* ptr, void* commandDef);
void* C_NSScriptCommand_InitWithCoder(void* ptr, void* inCoder);
void* C_NSScriptCommand_AllocScriptCommand();
void* C_NSScriptCommand_Autorelease(void* ptr);
void* C_NSScriptCommand_Retain(void* ptr);
void* C_NSScriptCommand_ScriptCommand_CurrentCommand();
void* C_NSScriptCommand_ExecuteCommand(void* ptr);
void* C_NSScriptCommand_PerformDefaultImplementation(void* ptr);
void C_NSScriptCommand_SuspendExecution(void* ptr);
void C_NSScriptCommand_ResumeExecutionWithResult(void* ptr, void* result);
void* C_NSScriptCommand_AppleEvent(void* ptr);
void* C_NSScriptCommand_EvaluatedReceivers(void* ptr);
void* C_NSScriptCommand_ReceiversSpecifier(void* ptr);
void C_NSScriptCommand_SetReceiversSpecifier(void* ptr, void* value);
Dictionary C_NSScriptCommand_Arguments(void* ptr);
void C_NSScriptCommand_SetArguments(void* ptr, Dictionary value);
Dictionary C_NSScriptCommand_EvaluatedArguments(void* ptr);
void* C_NSScriptCommand_DirectParameter(void* ptr);
void C_NSScriptCommand_SetDirectParameter(void* ptr, void* value);
void* C_NSScriptCommand_CommandDescription(void* ptr);
void* C_NSScriptCommand_ScriptErrorExpectedTypeDescriptor(void* ptr);
void C_NSScriptCommand_SetScriptErrorExpectedTypeDescriptor(void* ptr, void* value);
int C_NSScriptCommand_ScriptErrorNumber(void* ptr);
void C_NSScriptCommand_SetScriptErrorNumber(void* ptr, int value);
void* C_NSScriptCommand_ScriptErrorOffendingObjectDescriptor(void* ptr);
void C_NSScriptCommand_SetScriptErrorOffendingObjectDescriptor(void* ptr, void* value);
void* C_NSScriptCommand_ScriptErrorString(void* ptr);
void C_NSScriptCommand_SetScriptErrorString(void* ptr, void* value);
bool C_NSScriptCommand_IsWellFormed(void* ptr);
