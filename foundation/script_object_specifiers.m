#import "script_object_specifiers.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSScriptObjectSpecifiers.h>

void* C_ScriptObjectSpecifier_Alloc() {
    return [NSScriptObjectSpecifier alloc];
}

void* C_NSScriptObjectSpecifier_InitWithContainerClassDescription_ContainerSpecifier_Key(void* ptr, void* classDesc, void* container, void* property) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier initWithContainerClassDescription:(NSScriptClassDescription*)classDesc containerSpecifier:(NSScriptObjectSpecifier*)container key:(NSString*)property];
    return result_;
}

void* C_NSScriptObjectSpecifier_InitWithContainerSpecifier_Key(void* ptr, void* container, void* property) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier initWithContainerSpecifier:(NSScriptObjectSpecifier*)container key:(NSString*)property];
    return result_;
}

void* C_NSScriptObjectSpecifier_InitWithCoder(void* ptr, void* inCoder) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier initWithCoder:(NSCoder*)inCoder];
    return result_;
}

void* C_NSScriptObjectSpecifier_AllocScriptObjectSpecifier() {
    NSScriptObjectSpecifier* result_ = [NSScriptObjectSpecifier alloc];
    return result_;
}

void* C_NSScriptObjectSpecifier_Init(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier init];
    return result_;
}

void* C_NSScriptObjectSpecifier_NewScriptObjectSpecifier() {
    NSScriptObjectSpecifier* result_ = [NSScriptObjectSpecifier new];
    return result_;
}

void* C_NSScriptObjectSpecifier_Autorelease(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier autorelease];
    return result_;
}

void* C_NSScriptObjectSpecifier_Retain(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier retain];
    return result_;
}

void* C_NSScriptObjectSpecifier_ScriptObjectSpecifier_ObjectSpecifierWithDescriptor(void* descriptor) {
    NSScriptObjectSpecifier* result_ = [NSScriptObjectSpecifier objectSpecifierWithDescriptor:(NSAppleEventDescriptor*)descriptor];
    return result_;
}

void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingWithContainers(void* ptr, void* containers) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    id result_ = [nSScriptObjectSpecifier objectsByEvaluatingWithContainers:(id)containers];
    return result_;
}

void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingSpecifier(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    id result_ = [nSScriptObjectSpecifier objectsByEvaluatingSpecifier];
    return result_;
}

void* C_NSScriptObjectSpecifier_ContainerClassDescription(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptClassDescription* result_ = [nSScriptObjectSpecifier containerClassDescription];
    return result_;
}

void C_NSScriptObjectSpecifier_SetContainerClassDescription(void* ptr, void* value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setContainerClassDescription:(NSScriptClassDescription*)value];
}

bool C_NSScriptObjectSpecifier_ContainerIsObjectBeingTested(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    BOOL result_ = [nSScriptObjectSpecifier containerIsObjectBeingTested];
    return result_;
}

void C_NSScriptObjectSpecifier_SetContainerIsObjectBeingTested(void* ptr, bool value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setContainerIsObjectBeingTested:value];
}

bool C_NSScriptObjectSpecifier_ContainerIsRangeContainerObject(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    BOOL result_ = [nSScriptObjectSpecifier containerIsRangeContainerObject];
    return result_;
}

void C_NSScriptObjectSpecifier_SetContainerIsRangeContainerObject(void* ptr, bool value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setContainerIsRangeContainerObject:value];
}

void* C_NSScriptObjectSpecifier_ContainerSpecifier(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier containerSpecifier];
    return result_;
}

void C_NSScriptObjectSpecifier_SetContainerSpecifier(void* ptr, void* value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setContainerSpecifier:(NSScriptObjectSpecifier*)value];
}

void* C_NSScriptObjectSpecifier_ChildSpecifier(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier childSpecifier];
    return result_;
}

void C_NSScriptObjectSpecifier_SetChildSpecifier(void* ptr, void* value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setChildSpecifier:(NSScriptObjectSpecifier*)value];
}

void* C_NSScriptObjectSpecifier_Key(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSString* result_ = [nSScriptObjectSpecifier key];
    return result_;
}

void C_NSScriptObjectSpecifier_SetKey(void* ptr, void* value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setKey:(NSString*)value];
}

void* C_NSScriptObjectSpecifier_KeyClassDescription(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptClassDescription* result_ = [nSScriptObjectSpecifier keyClassDescription];
    return result_;
}

void* C_NSScriptObjectSpecifier_EvaluationErrorSpecifier(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSScriptObjectSpecifier* result_ = [nSScriptObjectSpecifier evaluationErrorSpecifier];
    return result_;
}

int C_NSScriptObjectSpecifier_EvaluationErrorNumber(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSInteger result_ = [nSScriptObjectSpecifier evaluationErrorNumber];
    return result_;
}

void C_NSScriptObjectSpecifier_SetEvaluationErrorNumber(void* ptr, int value) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    [nSScriptObjectSpecifier setEvaluationErrorNumber:value];
}

void* C_NSScriptObjectSpecifier_Descriptor(void* ptr) {
    NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
    NSAppleEventDescriptor* result_ = [nSScriptObjectSpecifier descriptor];
    return result_;
}
