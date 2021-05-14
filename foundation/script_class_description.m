#import <Foundation/Foundation.h>
#import "script_class_description.h"

void* C_ScriptClassDescription_Alloc() {
	return [NSScriptClassDescription alloc];
}

void* C_NSScriptClassDescription_Init(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSScriptClassDescription* result = [nSScriptClassDescription init];
	return result;
}

void* C_NSScriptClassDescription_ClassDescriptionForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSScriptClassDescription* result = [nSScriptClassDescription classDescriptionForKey:(NSString*)key];
	return result;
}

bool C_NSScriptClassDescription_IsLocationRequiredToCreateForKey(void* ptr, void* toManyRelationshipKey) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription isLocationRequiredToCreateForKey:(NSString*)toManyRelationshipKey];
	return result;
}

bool C_NSScriptClassDescription_HasOrderedToManyRelationshipForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription hasOrderedToManyRelationshipForKey:(NSString*)key];
	return result;
}

bool C_NSScriptClassDescription_HasPropertyForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription hasPropertyForKey:(NSString*)key];
	return result;
}

bool C_NSScriptClassDescription_HasReadablePropertyForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription hasReadablePropertyForKey:(NSString*)key];
	return result;
}

bool C_NSScriptClassDescription_HasWritablePropertyForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription hasWritablePropertyForKey:(NSString*)key];
	return result;
}

void* C_NSScriptClassDescription_TypeForKey(void* ptr, void* key) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSString* result = [nSScriptClassDescription typeForKey:(NSString*)key];
	return result;
}

void* C_NSScriptClassDescription_SelectorForCommand(void* ptr, void* commandDescription) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	SEL result = [nSScriptClassDescription selectorForCommand:(NSScriptCommandDescription*)commandDescription];
	return result;
}

bool C_NSScriptClassDescription_SupportsCommand(void* ptr, void* commandDescription) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	bool result = [nSScriptClassDescription supportsCommand:(NSScriptCommandDescription*)commandDescription];
	return result;
}

void* C_NSScriptClassDescription_SuperclassDescription(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSScriptClassDescription* result = [nSScriptClassDescription superclassDescription];
	return result;
}

void* C_NSScriptClassDescription_ClassName(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSString* result = [nSScriptClassDescription className];
	return result;
}

void* C_NSScriptClassDescription_DefaultSubcontainerAttributeKey(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSString* result = [nSScriptClassDescription defaultSubcontainerAttributeKey];
	return result;
}

void* C_NSScriptClassDescription_ImplementationClassName(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSString* result = [nSScriptClassDescription implementationClassName];
	return result;
}

void* C_NSScriptClassDescription_SuiteName(void* ptr) {
	NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
	NSString* result = [nSScriptClassDescription suiteName];
	return result;
}
