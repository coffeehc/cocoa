#import <Foundation/Foundation.h>
#import "index_path.h"

void* C_IndexPath_Alloc() {
    return [NSIndexPath alloc];
}

void* C_NSIndexPath_InitWithIndex(void* ptr, unsigned int index) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSIndexPath* result_ = [nSIndexPath initWithIndex:index];
    return result_;
}

void* C_NSIndexPath_Init(void* ptr) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSIndexPath* result_ = [nSIndexPath init];
    return result_;
}

void* C_NSIndexPath_IndexPathWithIndex(unsigned int index) {
    NSIndexPath* result_ = [NSIndexPath indexPathWithIndex:index];
    return result_;
}

void* C_NSIndexPath_IndexPathForItem_InSection(int item, int section) {
    NSIndexPath* result_ = [NSIndexPath indexPathForItem:item inSection:section];
    return result_;
}

void* C_NSIndexPath_IndexPathByAddingIndex(void* ptr, unsigned int index) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSIndexPath* result_ = [nSIndexPath indexPathByAddingIndex:index];
    return result_;
}

void* C_NSIndexPath_IndexPathByRemovingLastIndex(void* ptr) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSIndexPath* result_ = [nSIndexPath indexPathByRemovingLastIndex];
    return result_;
}

int C_NSIndexPath_Compare(void* ptr, void* otherObject) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSComparisonResult result_ = [nSIndexPath compare:(NSIndexPath*)otherObject];
    return result_;
}

unsigned int C_NSIndexPath_IndexAtPosition(void* ptr, unsigned int position) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSUInteger result_ = [nSIndexPath indexAtPosition:position];
    return result_;
}

int C_NSIndexPath_Section(void* ptr) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSInteger result_ = [nSIndexPath section];
    return result_;
}

int C_NSIndexPath_Item(void* ptr) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSInteger result_ = [nSIndexPath item];
    return result_;
}

unsigned int C_NSIndexPath_Length(void* ptr) {
    NSIndexPath* nSIndexPath = (NSIndexPath*)ptr;
    NSUInteger result_ = [nSIndexPath length];
    return result_;
}
