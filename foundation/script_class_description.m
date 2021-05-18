#import <Foundation/Foundation.h>
#import "script_class_description.h"

void* C_ScriptClassDescription_Alloc() {
    return [NSScriptClassDescription alloc];
}

void* C_NSScriptClassDescription_Init(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSScriptClassDescription* result_ = [nSScriptClassDescription init];
    return result_;
}

void* C_NSScriptClassDescription_ClassDescriptionForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSScriptClassDescription* result_ = [nSScriptClassDescription classDescriptionForKey:(NSString*)key];
    return result_;
}

bool C_NSScriptClassDescription_IsLocationRequiredToCreateForKey(void* ptr, void* toManyRelationshipKey) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription isLocationRequiredToCreateForKey:(NSString*)toManyRelationshipKey];
    return result_;
}

bool C_NSScriptClassDescription_HasOrderedToManyRelationshipForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription hasOrderedToManyRelationshipForKey:(NSString*)key];
    return result_;
}

bool C_NSScriptClassDescription_HasPropertyForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription hasPropertyForKey:(NSString*)key];
    return result_;
}

bool C_NSScriptClassDescription_HasReadablePropertyForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription hasReadablePropertyForKey:(NSString*)key];
    return result_;
}

bool C_NSScriptClassDescription_HasWritablePropertyForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription hasWritablePropertyForKey:(NSString*)key];
    return result_;
}

void* C_NSScriptClassDescription_TypeForKey(void* ptr, void* key) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSString* result_ = [nSScriptClassDescription typeForKey:(NSString*)key];
    return result_;
}

void* C_NSScriptClassDescription_SelectorForCommand(void* ptr, void* commandDescription) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    SEL result_ = [nSScriptClassDescription selectorForCommand:(NSScriptCommandDescription*)commandDescription];
    return result_;
}

bool C_NSScriptClassDescription_SupportsCommand(void* ptr, void* commandDescription) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    BOOL result_ = [nSScriptClassDescription supportsCommand:(NSScriptCommandDescription*)commandDescription];
    return result_;
}

void* C_NSScriptClassDescription_SuperclassDescription(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSScriptClassDescription* result_ = [nSScriptClassDescription superclassDescription];
    return result_;
}

void* C_NSScriptClassDescription_ClassName(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSString* result_ = [nSScriptClassDescription className];
    return result_;
}

void* C_NSScriptClassDescription_DefaultSubcontainerAttributeKey(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSString* result_ = [nSScriptClassDescription defaultSubcontainerAttributeKey];
    return result_;
}

void* C_NSScriptClassDescription_ImplementationClassName(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSString* result_ = [nSScriptClassDescription implementationClassName];
    return result_;
}

void* C_NSScriptClassDescription_SuiteName(void* ptr) {
    NSScriptClassDescription* nSScriptClassDescription = (NSScriptClassDescription*)ptr;
    NSString* result_ = [nSScriptClassDescription suiteName];
    return result_;
}
