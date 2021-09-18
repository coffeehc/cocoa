#import "index_set.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSIndexSet.h>

void* C_IndexSet_Alloc() {
    return [NSIndexSet alloc];
}

void* C_NSIndexSet_IndexSet_() {
    NSIndexSet* result_ = [NSIndexSet indexSet];
    return result_;
}

void* C_NSIndexSet_IndexSetWithIndex(unsigned int value) {
    NSIndexSet* result_ = [NSIndexSet indexSetWithIndex:value];
    return result_;
}

void* C_NSIndexSet_IndexSetWithIndexesInRange(NSRange _range) {
    NSIndexSet* result_ = [NSIndexSet indexSetWithIndexesInRange:_range];
    return result_;
}

void* C_NSIndexSet_InitWithIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet initWithIndex:value];
    return result_;
}

void* C_NSIndexSet_InitWithIndexesInRange(void* ptr, NSRange _range) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet initWithIndexesInRange:_range];
    return result_;
}

void* C_NSIndexSet_InitWithIndexSet(void* ptr, void* indexSet) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet initWithIndexSet:(NSIndexSet*)indexSet];
    return result_;
}

void* C_NSIndexSet_AllocIndexSet() {
    NSIndexSet* result_ = [NSIndexSet alloc];
    return result_;
}

void* C_NSIndexSet_Init(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet init];
    return result_;
}

void* C_NSIndexSet_NewIndexSet() {
    NSIndexSet* result_ = [NSIndexSet new];
    return result_;
}

void* C_NSIndexSet_Autorelease(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet autorelease];
    return result_;
}

void* C_NSIndexSet_Retain(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSIndexSet* result_ = [nSIndexSet retain];
    return result_;
}

bool C_NSIndexSet_ContainsIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    BOOL result_ = [nSIndexSet containsIndex:value];
    return result_;
}

bool C_NSIndexSet_ContainsIndexes(void* ptr, void* indexSet) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    BOOL result_ = [nSIndexSet containsIndexes:(NSIndexSet*)indexSet];
    return result_;
}

bool C_NSIndexSet_ContainsIndexesInRange(void* ptr, NSRange _range) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    BOOL result_ = [nSIndexSet containsIndexesInRange:_range];
    return result_;
}

bool C_NSIndexSet_IntersectsIndexesInRange(void* ptr, NSRange _range) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    BOOL result_ = [nSIndexSet intersectsIndexesInRange:_range];
    return result_;
}

unsigned int C_NSIndexSet_CountOfIndexesInRange(void* ptr, NSRange _range) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet countOfIndexesInRange:_range];
    return result_;
}

bool C_NSIndexSet_IsEqualToIndexSet(void* ptr, void* indexSet) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    BOOL result_ = [nSIndexSet isEqualToIndexSet:(NSIndexSet*)indexSet];
    return result_;
}

unsigned int C_NSIndexSet_IndexLessThanIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet indexLessThanIndex:value];
    return result_;
}

unsigned int C_NSIndexSet_IndexLessThanOrEqualToIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet indexLessThanOrEqualToIndex:value];
    return result_;
}

unsigned int C_NSIndexSet_IndexGreaterThanOrEqualToIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet indexGreaterThanOrEqualToIndex:value];
    return result_;
}

unsigned int C_NSIndexSet_IndexGreaterThanIndex(void* ptr, unsigned int value) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet indexGreaterThanIndex:value];
    return result_;
}

unsigned int C_NSIndexSet_Count(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet count];
    return result_;
}

unsigned int C_NSIndexSet_FirstIndex(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet firstIndex];
    return result_;
}

unsigned int C_NSIndexSet_LastIndex(void* ptr) {
    NSIndexSet* nSIndexSet = (NSIndexSet*)ptr;
    NSUInteger result_ = [nSIndexSet lastIndex];
    return result_;
}
