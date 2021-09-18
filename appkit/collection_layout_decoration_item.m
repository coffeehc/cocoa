#import "collection_layout_decoration_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSCollectionViewCompositionalLayout.h>

void* C_CollectionLayoutDecorationItem_Alloc() {
    return [NSCollectionLayoutDecorationItem alloc];
}

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(void* elementKind) {
    NSCollectionLayoutDecorationItem* result_ = [NSCollectionLayoutDecorationItem backgroundDecorationItemWithElementKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize(void* layoutSize) {
    NSCollectionLayoutDecorationItem* result_ = [NSCollectionLayoutDecorationItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems) {
    NSMutableArray* objcSupplementaryItems = [NSMutableArray arrayWithCapacity:supplementaryItems.len];
    if (supplementaryItems.len > 0) {
    	void** supplementaryItemsData = (void**)supplementaryItems.data;
    	for (int i = 0; i < supplementaryItems.len; i++) {
    		void* p = supplementaryItemsData[i];
    		[objcSupplementaryItems addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    NSCollectionLayoutDecorationItem* result_ = [NSCollectionLayoutDecorationItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize supplementaryItems:objcSupplementaryItems];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_AllocCollectionLayoutDecorationItem() {
    NSCollectionLayoutDecorationItem* result_ = [NSCollectionLayoutDecorationItem alloc];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_Autorelease(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSCollectionLayoutDecorationItem* result_ = [nSCollectionLayoutDecorationItem autorelease];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_Retain(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSCollectionLayoutDecorationItem* result_ = [nSCollectionLayoutDecorationItem retain];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_ElementKind(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSString* result_ = [nSCollectionLayoutDecorationItem elementKind];
    return result_;
}

int C_NSCollectionLayoutDecorationItem_ZIndex(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSInteger result_ = [nSCollectionLayoutDecorationItem zIndex];
    return result_;
}

void C_NSCollectionLayoutDecorationItem_SetZIndex(void* ptr, int value) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    [nSCollectionLayoutDecorationItem setZIndex:value];
}
