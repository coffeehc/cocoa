#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Predicate_Alloc();

void* C_NSPredicate_PredicateWithSubstitutionVariables(void* ptr, Dictionary variables);
void* C_NSPredicate_AllocPredicate();
void* C_NSPredicate_Init(void* ptr);
void* C_NSPredicate_NewPredicate();
void* C_NSPredicate_Autorelease(void* ptr);
void* C_NSPredicate_Retain(void* ptr);
void* C_NSPredicate_PredicateWithFormat_ArgumentArray(void* predicateFormat, Array arguments);
void* C_NSPredicate_PredicateWithValue(bool value);
void* C_NSPredicate_PredicateFromMetadataQueryString(void* queryString);
bool C_NSPredicate_EvaluateWithObject(void* ptr, void* object);
bool C_NSPredicate_EvaluateWithObject_SubstitutionVariables(void* ptr, void* object, Dictionary bindings);
void C_NSPredicate_AllowEvaluation(void* ptr);
void* C_NSPredicate_PredicateFormat(void* ptr);
