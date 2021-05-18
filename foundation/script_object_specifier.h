#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_ScriptObjectSpecifier_Alloc();

void* C_NSScriptObjectSpecifier_InitWithContainerClassDescription_ContainerSpecifier_Key(void* ptr, void* classDesc, void* container, void* property);
void* C_NSScriptObjectSpecifier_InitWithContainerSpecifier_Key(void* ptr, void* container, void* property);
void* C_NSScriptObjectSpecifier_InitWithCoder(void* ptr, void* inCoder);
void* C_NSScriptObjectSpecifier_Init(void* ptr);
void* C_NSScriptObjectSpecifier_ScriptObjectSpecifier_ObjectSpecifierWithDescriptor(void* descriptor);
void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingWithContainers(void* ptr, void* containers);
void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingSpecifier(void* ptr);
void* C_NSScriptObjectSpecifier_ContainerClassDescription(void* ptr);
void C_NSScriptObjectSpecifier_SetContainerClassDescription(void* ptr, void* value);
bool C_NSScriptObjectSpecifier_ContainerIsObjectBeingTested(void* ptr);
void C_NSScriptObjectSpecifier_SetContainerIsObjectBeingTested(void* ptr, bool value);
bool C_NSScriptObjectSpecifier_ContainerIsRangeContainerObject(void* ptr);
void C_NSScriptObjectSpecifier_SetContainerIsRangeContainerObject(void* ptr, bool value);
void* C_NSScriptObjectSpecifier_ContainerSpecifier(void* ptr);
void C_NSScriptObjectSpecifier_SetContainerSpecifier(void* ptr, void* value);
void* C_NSScriptObjectSpecifier_ChildSpecifier(void* ptr);
void C_NSScriptObjectSpecifier_SetChildSpecifier(void* ptr, void* value);
void* C_NSScriptObjectSpecifier_Key(void* ptr);
void C_NSScriptObjectSpecifier_SetKey(void* ptr, void* value);
void* C_NSScriptObjectSpecifier_KeyClassDescription(void* ptr);
void* C_NSScriptObjectSpecifier_EvaluationErrorSpecifier(void* ptr);
int C_NSScriptObjectSpecifier_EvaluationErrorNumber(void* ptr);
void C_NSScriptObjectSpecifier_SetEvaluationErrorNumber(void* ptr, int value);
void* C_NSScriptObjectSpecifier_Descriptor(void* ptr);
