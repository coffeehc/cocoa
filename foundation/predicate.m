#import <Foundation/Foundation.h>
#import "predicate.h"

void* C_Predicate_Alloc() {
    return [NSPredicate alloc];
}

void* C_NSPredicate_PredicateWithSubstitutionVariables(void* ptr, Dictionary variables) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSMutableDictionary* objcVariables = [NSMutableDictionary dictionaryWithCapacity:variables.len];
    if (variables.len > 0) {
    	void** variablesKeyData = (void**)variables.key_data;
    	void** variablesValueData = (void**)variables.value_data;
    	for (int i = 0; i < variables.len; i++) {
    		void* kp = variablesKeyData[i];
    		void* vp = variablesValueData[i];
    		[objcVariables setObject:(NSString*)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSPredicate* result_ = [nSPredicate predicateWithSubstitutionVariables:objcVariables];
    return result_;
}

void* C_NSPredicate_AllocPredicate() {
    NSPredicate* result_ = [NSPredicate alloc];
    return result_;
}

void* C_NSPredicate_Init(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSPredicate* result_ = [nSPredicate init];
    return result_;
}

void* C_NSPredicate_NewPredicate() {
    NSPredicate* result_ = [NSPredicate new];
    return result_;
}

void* C_NSPredicate_Autorelease(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSPredicate* result_ = [nSPredicate autorelease];
    return result_;
}

void* C_NSPredicate_Retain(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSPredicate* result_ = [nSPredicate retain];
    return result_;
}

void* C_NSPredicate_PredicateWithFormat_ArgumentArray(void* predicateFormat, Array arguments) {
    NSMutableArray* objcArguments = [NSMutableArray arrayWithCapacity:arguments.len];
    if (arguments.len > 0) {
    	void** argumentsData = (void**)arguments.data;
    	for (int i = 0; i < arguments.len; i++) {
    		void* p = argumentsData[i];
    		[objcArguments addObject:(NSObject*)(NSObject*)p];
    	}
    }
    NSPredicate* result_ = [NSPredicate predicateWithFormat:(NSString*)predicateFormat argumentArray:objcArguments];
    return result_;
}

void* C_NSPredicate_PredicateWithValue(bool value) {
    NSPredicate* result_ = [NSPredicate predicateWithValue:value];
    return result_;
}

void* C_NSPredicate_PredicateFromMetadataQueryString(void* queryString) {
    NSPredicate* result_ = [NSPredicate predicateFromMetadataQueryString:(NSString*)queryString];
    return result_;
}

bool C_NSPredicate_EvaluateWithObject(void* ptr, void* object) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    BOOL result_ = [nSPredicate evaluateWithObject:(id)object];
    return result_;
}

bool C_NSPredicate_EvaluateWithObject_SubstitutionVariables(void* ptr, void* object, Dictionary bindings) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSMutableDictionary* objcBindings = [NSMutableDictionary dictionaryWithCapacity:bindings.len];
    if (bindings.len > 0) {
    	void** bindingsKeyData = (void**)bindings.key_data;
    	void** bindingsValueData = (void**)bindings.value_data;
    	for (int i = 0; i < bindings.len; i++) {
    		void* kp = bindingsKeyData[i];
    		void* vp = bindingsValueData[i];
    		[objcBindings setObject:(NSString*)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    BOOL result_ = [nSPredicate evaluateWithObject:(id)object substitutionVariables:objcBindings];
    return result_;
}

void C_NSPredicate_AllowEvaluation(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    [nSPredicate allowEvaluation];
}

void* C_NSPredicate_PredicateFormat(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSString* result_ = [nSPredicate predicateFormat];
    return result_;
}
