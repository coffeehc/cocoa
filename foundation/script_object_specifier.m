#import <Foundation/Foundation.h>
#import "script_object_specifier.h"

void* C_ScriptObjectSpecifier_Alloc() {
	return [NSScriptObjectSpecifier alloc];
}

void* C_NSScriptObjectSpecifier_InitWithContainerClassDescription_ContainerSpecifier_Key(void* ptr, void* classDesc, void* container, void* property) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier initWithContainerClassDescription:(NSScriptClassDescription*)classDesc containerSpecifier:(NSScriptObjectSpecifier*)container key:(NSString*)property];
	return result;
}

void* C_NSScriptObjectSpecifier_InitWithContainerSpecifier_Key(void* ptr, void* container, void* property) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier initWithContainerSpecifier:(NSScriptObjectSpecifier*)container key:(NSString*)property];
	return result;
}

void* C_NSScriptObjectSpecifier_InitWithCoder(void* ptr, void* inCoder) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier initWithCoder:(NSCoder*)inCoder];
	return result;
}

void* C_NSScriptObjectSpecifier_Init(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier init];
	return result;
}

void* C_NSScriptObjectSpecifier_ScriptObjectSpecifierObjectSpecifierWithDescriptor(void* descriptor) {
	NSScriptObjectSpecifier* result = [NSScriptObjectSpecifier objectSpecifierWithDescriptor:(NSAppleEventDescriptor*)descriptor];
	return result;
}

void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingWithContainers(void* ptr, void* containers) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	id result = [nSScriptObjectSpecifier objectsByEvaluatingWithContainers:(id)containers];
	return result;
}

void* C_NSScriptObjectSpecifier_ObjectsByEvaluatingSpecifier(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	id result = [nSScriptObjectSpecifier objectsByEvaluatingSpecifier];
	return result;
}

void* C_NSScriptObjectSpecifier_ContainerClassDescription(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptClassDescription* result = [nSScriptObjectSpecifier containerClassDescription];
	return result;
}

void C_NSScriptObjectSpecifier_SetContainerClassDescription(void* ptr, void* value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setContainerClassDescription:(NSScriptClassDescription*)value];
}

bool C_NSScriptObjectSpecifier_ContainerIsObjectBeingTested(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	bool result = [nSScriptObjectSpecifier containerIsObjectBeingTested];
	return result;
}

void C_NSScriptObjectSpecifier_SetContainerIsObjectBeingTested(void* ptr, bool value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setContainerIsObjectBeingTested:value];
}

bool C_NSScriptObjectSpecifier_ContainerIsRangeContainerObject(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	bool result = [nSScriptObjectSpecifier containerIsRangeContainerObject];
	return result;
}

void C_NSScriptObjectSpecifier_SetContainerIsRangeContainerObject(void* ptr, bool value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setContainerIsRangeContainerObject:value];
}

void* C_NSScriptObjectSpecifier_ContainerSpecifier(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier containerSpecifier];
	return result;
}

void C_NSScriptObjectSpecifier_SetContainerSpecifier(void* ptr, void* value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setContainerSpecifier:(NSScriptObjectSpecifier*)value];
}

void* C_NSScriptObjectSpecifier_ChildSpecifier(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier childSpecifier];
	return result;
}

void C_NSScriptObjectSpecifier_SetChildSpecifier(void* ptr, void* value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setChildSpecifier:(NSScriptObjectSpecifier*)value];
}

void* C_NSScriptObjectSpecifier_Key(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSString* result = [nSScriptObjectSpecifier key];
	return result;
}

void C_NSScriptObjectSpecifier_SetKey(void* ptr, void* value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setKey:(NSString*)value];
}

void* C_NSScriptObjectSpecifier_KeyClassDescription(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptClassDescription* result = [nSScriptObjectSpecifier keyClassDescription];
	return result;
}

void* C_NSScriptObjectSpecifier_EvaluationErrorSpecifier(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSScriptObjectSpecifier* result = [nSScriptObjectSpecifier evaluationErrorSpecifier];
	return result;
}

int C_NSScriptObjectSpecifier_EvaluationErrorNumber(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	int result = [nSScriptObjectSpecifier evaluationErrorNumber];
	return result;
}

void C_NSScriptObjectSpecifier_SetEvaluationErrorNumber(void* ptr, int value) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	[nSScriptObjectSpecifier setEvaluationErrorNumber:value];
}

void* C_NSScriptObjectSpecifier_Descriptor(void* ptr) {
	NSScriptObjectSpecifier* nSScriptObjectSpecifier = (NSScriptObjectSpecifier*)ptr;
	NSAppleEventDescriptor* result = [nSScriptObjectSpecifier descriptor];
	return result;
}
