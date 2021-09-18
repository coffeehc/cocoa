#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ScriptClassDescription_Alloc();

void* C_NSScriptClassDescription_AllocScriptClassDescription();
void* C_NSScriptClassDescription_Init(void* ptr);
void* C_NSScriptClassDescription_NewScriptClassDescription();
void* C_NSScriptClassDescription_Autorelease(void* ptr);
void* C_NSScriptClassDescription_Retain(void* ptr);
void* C_NSScriptClassDescription_ClassDescriptionForKey(void* ptr, void* key);
bool C_NSScriptClassDescription_IsLocationRequiredToCreateForKey(void* ptr, void* toManyRelationshipKey);
bool C_NSScriptClassDescription_HasOrderedToManyRelationshipForKey(void* ptr, void* key);
bool C_NSScriptClassDescription_HasPropertyForKey(void* ptr, void* key);
bool C_NSScriptClassDescription_HasReadablePropertyForKey(void* ptr, void* key);
bool C_NSScriptClassDescription_HasWritablePropertyForKey(void* ptr, void* key);
void* C_NSScriptClassDescription_TypeForKey(void* ptr, void* key);
void* C_NSScriptClassDescription_SelectorForCommand(void* ptr, void* commandDescription);
bool C_NSScriptClassDescription_SupportsCommand(void* ptr, void* commandDescription);
void* C_NSScriptClassDescription_SuperclassDescription(void* ptr);
void* C_NSScriptClassDescription_ClassName(void* ptr);
void* C_NSScriptClassDescription_DefaultSubcontainerAttributeKey(void* ptr);
void* C_NSScriptClassDescription_ImplementationClassName(void* ptr);
void* C_NSScriptClassDescription_SuiteName(void* ptr);
