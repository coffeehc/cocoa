#import "collection_layout_supplementary_item.h"
#import <AppKit/NSCollectionViewCompositionalLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionLayoutSupplementaryItem_Alloc() {
    return [NSCollectionLayoutSupplementaryItem alloc];
}

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor) {
    NSCollectionLayoutSupplementaryItem* result_ = [NSCollectionLayoutSupplementaryItem supplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind containerAnchor:(NSCollectionLayoutAnchor*)containerAnchor];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor) {
    NSCollectionLayoutSupplementaryItem* result_ = [NSCollectionLayoutSupplementaryItem supplementaryItemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize elementKind:(NSString*)elementKind containerAnchor:(NSCollectionLayoutAnchor*)containerAnchor itemAnchor:(NSCollectionLayoutAnchor*)itemAnchor];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize(void* layoutSize) {
    NSCollectionLayoutSupplementaryItem* result_ = [NSCollectionLayoutSupplementaryItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems) {
    NSMutableArray* objcSupplementaryItems = [NSMutableArray arrayWithCapacity:supplementaryItems.len];
    if (supplementaryItems.len > 0) {
    	void** supplementaryItemsData = (void**)supplementaryItems.data;
    	for (int i = 0; i < supplementaryItems.len; i++) {
    		void* p = supplementaryItemsData[i];
    		[objcSupplementaryItems addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    NSCollectionLayoutSupplementaryItem* result_ = [NSCollectionLayoutSupplementaryItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize supplementaryItems:objcSupplementaryItems];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_AllocCollectionLayoutSupplementaryItem() {
    NSCollectionLayoutSupplementaryItem* result_ = [NSCollectionLayoutSupplementaryItem alloc];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_Autorelease(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSCollectionLayoutSupplementaryItem* result_ = [nSCollectionLayoutSupplementaryItem autorelease];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_Retain(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSCollectionLayoutSupplementaryItem* result_ = [nSCollectionLayoutSupplementaryItem retain];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_ItemAnchor(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSCollectionLayoutAnchor* result_ = [nSCollectionLayoutSupplementaryItem itemAnchor];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_ContainerAnchor(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSCollectionLayoutAnchor* result_ = [nSCollectionLayoutSupplementaryItem containerAnchor];
    return result_;
}

void* C_NSCollectionLayoutSupplementaryItem_ElementKind(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSString* result_ = [nSCollectionLayoutSupplementaryItem elementKind];
    return result_;
}

int C_NSCollectionLayoutSupplementaryItem_ZIndex(void* ptr) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    NSInteger result_ = [nSCollectionLayoutSupplementaryItem zIndex];
    return result_;
}

void C_NSCollectionLayoutSupplementaryItem_SetZIndex(void* ptr, int value) {
    NSCollectionLayoutSupplementaryItem* nSCollectionLayoutSupplementaryItem = (NSCollectionLayoutSupplementaryItem*)ptr;
    [nSCollectionLayoutSupplementaryItem setZIndex:value];
}
