#import <Foundation/Foundation.h>
#import "set.h"

void* C_Set_Alloc() {
    return [NSSet alloc];
}

void* C_NSSet_InitWithSet(void* ptr, void* set) {
    NSSet* nSSet = (NSSet*)ptr;
    NSSet* result_ = [nSSet initWithSet:(NSSet*)set];
    return result_;
}

void* C_NSSet_InitWithSet_CopyItems(void* ptr, void* set, bool flag) {
    NSSet* nSSet = (NSSet*)ptr;
    NSSet* result_ = [nSSet initWithSet:(NSSet*)set copyItems:flag];
    return result_;
}

void* C_NSSet_Init(void* ptr) {
    NSSet* nSSet = (NSSet*)ptr;
    NSSet* result_ = [nSSet init];
    return result_;
}

void* C_NSSet_InitWithCoder(void* ptr, void* coder) {
    NSSet* nSSet = (NSSet*)ptr;
    NSSet* result_ = [nSSet initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSet_Set_() {
    NSSet* result_ = [NSSet set];
    return result_;
}

void* C_NSSet_SetWithSet(void* set) {
    NSSet* result_ = [NSSet setWithSet:(NSSet*)set];
    return result_;
}

void* C_NSSet_SetByAddingObjectsFromSet(void* ptr, void* other) {
    NSSet* nSSet = (NSSet*)ptr;
    NSSet* result_ = [nSSet setByAddingObjectsFromSet:(NSSet*)other];
    return result_;
}

void C_NSSet_MakeObjectsPerformSelector(void* ptr, void* aSelector) {
    NSSet* nSSet = (NSSet*)ptr;
    [nSSet makeObjectsPerformSelector:(SEL)aSelector];
}

void C_NSSet_MakeObjectsPerformSelector_WithObject(void* ptr, void* aSelector, void* argument) {
    NSSet* nSSet = (NSSet*)ptr;
    [nSSet makeObjectsPerformSelector:(SEL)aSelector withObject:(id)argument];
}

bool C_NSSet_IsSubsetOfSet(void* ptr, void* otherSet) {
    NSSet* nSSet = (NSSet*)ptr;
    BOOL result_ = [nSSet isSubsetOfSet:(NSSet*)otherSet];
    return result_;
}

bool C_NSSet_IntersectsSet(void* ptr, void* otherSet) {
    NSSet* nSSet = (NSSet*)ptr;
    BOOL result_ = [nSSet intersectsSet:(NSSet*)otherSet];
    return result_;
}

bool C_NSSet_IsEqualToSet(void* ptr, void* otherSet) {
    NSSet* nSSet = (NSSet*)ptr;
    BOOL result_ = [nSSet isEqualToSet:(NSSet*)otherSet];
    return result_;
}

void* C_NSSet_ValueForKey(void* ptr, void* key) {
    NSSet* nSSet = (NSSet*)ptr;
    id result_ = [nSSet valueForKey:(NSString*)key];
    return result_;
}

void C_NSSet_SetValue_ForKey(void* ptr, void* value, void* key) {
    NSSet* nSSet = (NSSet*)ptr;
    [nSSet setValue:(id)value forKey:(NSString*)key];
}

void C_NSSet_RemoveObserver_ForKeyPath(void* ptr, void* observer, void* keyPath) {
    NSSet* nSSet = (NSSet*)ptr;
    [nSSet removeObserver:(NSObject*)observer forKeyPath:(NSString*)keyPath];
}

void* C_NSSet_DescriptionWithLocale(void* ptr, void* locale) {
    NSSet* nSSet = (NSSet*)ptr;
    NSString* result_ = [nSSet descriptionWithLocale:(id)locale];
    return result_;
}

unsigned int C_NSSet_Count(void* ptr) {
    NSSet* nSSet = (NSSet*)ptr;
    NSUInteger result_ = [nSSet count];
    return result_;
}

void* C_NSSet_Description(void* ptr) {
    NSSet* nSSet = (NSSet*)ptr;
    NSString* result_ = [nSSet description];
    return result_;
}
