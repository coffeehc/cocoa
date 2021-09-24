#import "collection_layout_edge_spacing.h"
#import <AppKit/NSCollectionViewCompositionalLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionLayoutEdgeSpacing_Alloc() {
    return [NSCollectionLayoutEdgeSpacing alloc];
}

void* C_NSCollectionLayoutEdgeSpacing_CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(void* leading, void* top, void* trailing, void* bottom) {
    NSCollectionLayoutEdgeSpacing* result_ = [NSCollectionLayoutEdgeSpacing spacingForLeading:(NSCollectionLayoutSpacing*)leading top:(NSCollectionLayoutSpacing*)top trailing:(NSCollectionLayoutSpacing*)trailing bottom:(NSCollectionLayoutSpacing*)bottom];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_AllocCollectionLayoutEdgeSpacing() {
    NSCollectionLayoutEdgeSpacing* result_ = [NSCollectionLayoutEdgeSpacing alloc];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Autorelease(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutEdgeSpacing* result_ = [nSCollectionLayoutEdgeSpacing autorelease];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Retain(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutEdgeSpacing* result_ = [nSCollectionLayoutEdgeSpacing retain];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Leading(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutEdgeSpacing leading];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Top(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutEdgeSpacing top];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Trailing(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutEdgeSpacing trailing];
    return result_;
}

void* C_NSCollectionLayoutEdgeSpacing_Bottom(void* ptr) {
    NSCollectionLayoutEdgeSpacing* nSCollectionLayoutEdgeSpacing = (NSCollectionLayoutEdgeSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutEdgeSpacing bottom];
    return result_;
}
