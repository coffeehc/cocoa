#import <Foundation/Foundation.h>
#import "predicate.h"

void* C_Predicate_Alloc() {
    return [NSPredicate alloc];
}

void* C_NSPredicate_Init(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSPredicate* result_ = [nSPredicate init];
    return result_;
}

void* C_NSPredicate_PredicateWithFormat_ArgumentArray(void* predicateFormat, Array arguments) {
    NSMutableArray* objcArguments = [[NSMutableArray alloc] init];
    void** argumentsData = (void**)arguments.data;
    for (int i = 0; i < arguments.len; i++) {
    	void* p = argumentsData[i];
    	[objcArguments addObject:(NSObject*)(NSObject*)p];
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

void C_NSPredicate_AllowEvaluation(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    [nSPredicate allowEvaluation];
}

void* C_NSPredicate_PredicateFormat(void* ptr) {
    NSPredicate* nSPredicate = (NSPredicate*)ptr;
    NSString* result_ = [nSPredicate predicateFormat];
    return result_;
}
