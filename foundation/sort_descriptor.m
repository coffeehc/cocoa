#import "sort_descriptor.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSSortDescriptor.h>

void* C_SortDescriptor_Alloc() {
    return [NSSortDescriptor alloc];
}

void* C_NSSortDescriptor_SortDescriptorWithKey_Ascending(void* key, bool ascending) {
    NSSortDescriptor* result_ = [NSSortDescriptor sortDescriptorWithKey:(NSString*)key ascending:ascending];
    return result_;
}

void* C_NSSortDescriptor_InitWithKey_Ascending(void* ptr, void* key, bool ascending) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor initWithKey:(NSString*)key ascending:ascending];
    return result_;
}

void* C_NSSortDescriptor_SortDescriptorWithKey_Ascending_Selector(void* key, bool ascending, void* selector) {
    NSSortDescriptor* result_ = [NSSortDescriptor sortDescriptorWithKey:(NSString*)key ascending:ascending selector:(SEL)selector];
    return result_;
}

void* C_NSSortDescriptor_InitWithKey_Ascending_Selector(void* ptr, void* key, bool ascending, void* selector) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor initWithKey:(NSString*)key ascending:ascending selector:(SEL)selector];
    return result_;
}

void* C_NSSortDescriptor_InitWithCoder(void* ptr, void* coder) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSortDescriptor_AllocSortDescriptor() {
    NSSortDescriptor* result_ = [NSSortDescriptor alloc];
    return result_;
}

void* C_NSSortDescriptor_Init(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor init];
    return result_;
}

void* C_NSSortDescriptor_NewSortDescriptor() {
    NSSortDescriptor* result_ = [NSSortDescriptor new];
    return result_;
}

void* C_NSSortDescriptor_Autorelease(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor autorelease];
    return result_;
}

void* C_NSSortDescriptor_Retain(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSSortDescriptor* result_ = [nSSortDescriptor retain];
    return result_;
}

int C_NSSortDescriptor_CompareObject_ToObject(void* ptr, void* object1, void* object2) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSComparisonResult result_ = [nSSortDescriptor compareObject:(id)object1 toObject:(id)object2];
    return result_;
}

void C_NSSortDescriptor_AllowEvaluation(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    [nSSortDescriptor allowEvaluation];
}

bool C_NSSortDescriptor_Ascending(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    BOOL result_ = [nSSortDescriptor ascending];
    return result_;
}

void* C_NSSortDescriptor_Key(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    NSString* result_ = [nSSortDescriptor key];
    return result_;
}

void* C_NSSortDescriptor_Selector(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    SEL result_ = [nSSortDescriptor selector];
    return result_;
}

void* C_NSSortDescriptor_ReversedSortDescriptor(void* ptr) {
    NSSortDescriptor* nSSortDescriptor = (NSSortDescriptor*)ptr;
    id result_ = [nSSortDescriptor reversedSortDescriptor];
    return result_;
}
