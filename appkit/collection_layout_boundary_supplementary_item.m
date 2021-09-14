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

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor) {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem supplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind containerAnchor:(NSCollectionLayoutAnchor*)containerAnchor];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor) {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem supplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind containerAnchor:(NSCollectionLayoutAnchor*)containerAnchor itemAnchor:(NSCollectionLayoutAnchor*)itemAnchor];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(void* layoutSize) {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems) {
    NSMutableArray* objcSupplementaryItems = [NSMutableArray arrayWithCapacity:supplementaryItems.len];
    if (supplementaryItems.len > 0) {
    	void** supplementaryItemsData = (void**)supplementaryItems.data;
    	for (int i = 0; i < supplementaryItems.len; i++) {
    		void* p = supplementaryItemsData[i];
    		[objcSupplementaryItems addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize supplementaryItems:objcSupplementaryItems];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_AllocCollectionLayoutBoundarySupplementaryItem() {
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [NSCollectionLayoutBoundarySupplementaryItem alloc];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_Autorelease(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [nSCollectionLayoutBoundarySupplementaryItem autorelease];
    return result_;
}

void* C_NSCollectionLayoutBoundarySupplementaryItem_Retain(void* ptr) {
    NSCollectionLayoutBoundarySupplementaryItem* nSCollectionLayoutBoundarySupplementaryItem = (NSCollectionLayoutBoundarySupplementaryItem*)ptr;
    NSCollectionLayoutBoundarySupplementaryItem* result_ = [nSCollectionLayoutBoundarySupplementaryItem retain];
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
