#import <Appkit/Appkit.h>
#import "collection_layout_boundary_supplementary_item.h"

void* C_CollectionLayoutBoundarySupplementaryItem_Alloc() {
    return [NSCollectionLayoutBoundarySupplementaryItem alloc];
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(void* layoutSize, void* elementKind, int alignment) {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem boundarySupplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind alignment:alignment];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(void* layoutSize, void* elementKind, int alignment, CGPoint absoluteOffset) {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem boundarySupplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind alignment:alignment absoluteOffset:absoluteOffset];
    return result_;
}

bool C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    BOOL result_ = [nSCollectionLayoutBoundarySupplementaryItem pinToVisibleBounds];
    return result_;
}

void C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(void* ptr, bool value) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    [nSCollectionLayoutBoundarySupplementaryItem setPinToVisibleBounds:value];
}

CGPoint C_NSCollectionLayoutBoundarySupplementaryItem_Offset(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    NSPoint result_ = [nSCollectionLayoutBoundarySupplementaryItem offset];
    return result_;
}

int C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    NSRectAlignment result_ = [nSCollectionLayoutBoundarySupplementaryItem alignment];
    return result_;
}

bool C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    BOOL result_ = [nSCollectionLayoutBoundarySupplementaryItem extendsBoundary];
    return result_;
}

void C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(void* ptr, bool value) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    [nSCollectionLayoutBoundarySupplementaryItem setExtendsBoundary:value];
}
