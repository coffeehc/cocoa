#import "collection_layout_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>

void* C_CollectionLayoutItem_Alloc() {
    return [NSCollectionLayoutItem alloc];
}

void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize(void* layoutSize) {
    NSCollectionLayoutItem* result_ = [NSCollectionLayoutItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize];
    return result_;
}

void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems) {
    NSMutableArray* objcSupplementaryItems = [NSMutableArray arrayWithCapacity:supplementaryItems.len];
    if (supplementaryItems.len > 0) {
    	void** supplementaryItemsData = (void**)supplementaryItems.data;
    	for (int i = 0; i < supplementaryItems.len; i++) {
    		void* p = supplementaryItemsData[i];
    		[objcSupplementaryItems addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    NSCollectionLayoutItem* result_ = [NSCollectionLayoutItem itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize supplementaryItems:objcSupplementaryItems];
    return result_;
}

void* C_NSCollectionLayoutItem_AllocCollectionLayoutItem() {
    NSCollectionLayoutItem* result_ = [NSCollectionLayoutItem alloc];
    return result_;
}

void* C_NSCollectionLayoutItem_Autorelease(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSCollectionLayoutItem* result_ = [nSCollectionLayoutItem autorelease];
    return result_;
}

void* C_NSCollectionLayoutItem_Retain(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSCollectionLayoutItem* result_ = [nSCollectionLayoutItem retain];
    return result_;
}

void* C_NSCollectionLayoutItem_LayoutSize(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSCollectionLayoutSize* result_ = [nSCollectionLayoutItem layoutSize];
    return result_;
}

Array C_NSCollectionLayoutItem_SupplementaryItems(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSArray* result_ = [nSCollectionLayoutItem supplementaryItems];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSCollectionLayoutItem_EdgeSpacing(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSCollectionLayoutEdgeSpacing* result_ = [nSCollectionLayoutItem edgeSpacing];
    return result_;
}

void C_NSCollectionLayoutItem_SetEdgeSpacing(void* ptr, void* value) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    [nSCollectionLayoutItem setEdgeSpacing:(NSCollectionLayoutEdgeSpacing*)value];
}

NSDirectionalEdgeInsets C_NSCollectionLayoutItem_ContentInsets(void* ptr) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    NSDirectionalEdgeInsets result_ = [nSCollectionLayoutItem contentInsets];
    return result_;
}

void C_NSCollectionLayoutItem_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value) {
    NSCollectionLayoutItem* nSCollectionLayoutItem = (NSCollectionLayoutItem*)ptr;
    [nSCollectionLayoutItem setContentInsets:value];
}
