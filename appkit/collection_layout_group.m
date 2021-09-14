#import <Appkit/Appkit.h>
#import "collection_layout_group.h"

void* C_CollectionLayoutGroup_Alloc() {
    return [NSCollectionLayoutGroup alloc];
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems) {
    NSMutableArray* objcSubitems = [NSMutableArray arrayWithCapacity:subitems.len];
    if (subitems.len > 0) {
    	void** subitemsData = (void**)subitems.data;
    	for (int i = 0; i < subitems.len; i++) {
    		void* p = subitemsData[i];
    		[objcSubitems addObject:(NSCollectionLayoutItem*)p];
    	}
    }
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup horizontalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitems:objcSubitems];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count) {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup horizontalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitem:(NSCollectionLayoutItem*)subitem count:count];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems) {
    NSMutableArray* objcSubitems = [NSMutableArray arrayWithCapacity:subitems.len];
    if (subitems.len > 0) {
    	void** subitemsData = (void**)subitems.data;
    	for (int i = 0; i < subitems.len; i++) {
    		void* p = subitemsData[i];
    		[objcSubitems addObject:(NSCollectionLayoutItem*)p];
    	}
    }
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup verticalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitems:objcSubitems];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count) {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup verticalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitem:(NSCollectionLayoutItem*)subitem count:count];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize(void* layoutSize) {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems) {
    NSMutableArray* objcSupplementaryItems = [NSMutableArray arrayWithCapacity:supplementaryItems.len];
    if (supplementaryItems.len > 0) {
    	void** supplementaryItemsData = (void**)supplementaryItems.data;
    	for (int i = 0; i < supplementaryItems.len; i++) {
    		void* p = supplementaryItemsData[i];
    		[objcSupplementaryItems addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup itemWithLayoutSize:(NSCollectionLayoutSize*)layoutSize supplementaryItems:objcSupplementaryItems];
    return result_;
}

void* C_NSCollectionLayoutGroup_AllocCollectionLayoutGroup() {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup alloc];
    return result_;
}

void* C_NSCollectionLayoutGroup_Autorelease(void* ptr) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSCollectionLayoutGroup* result_ = [nSCollectionLayoutGroup autorelease];
    return result_;
}

void* C_NSCollectionLayoutGroup_Retain(void* ptr) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSCollectionLayoutGroup* result_ = [nSCollectionLayoutGroup retain];
    return result_;
}

void* C_NSCollectionLayoutGroup_VisualDescription(void* ptr) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSString* result_ = [nSCollectionLayoutGroup visualDescription];
    return result_;
}

Array C_NSCollectionLayoutGroup_Subitems(void* ptr) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSArray* result_ = [nSCollectionLayoutGroup subitems];
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

void C_NSCollectionLayoutGroup_SetSupplementaryItems(void* ptr, Array value) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSCollectionLayoutSupplementaryItem*)p];
    	}
    }
    [nSCollectionLayoutGroup setSupplementaryItems:objcValue];
}

void* C_NSCollectionLayoutGroup_InterItemSpacing(void* ptr) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutGroup interItemSpacing];
    return result_;
}

void C_NSCollectionLayoutGroup_SetInterItemSpacing(void* ptr, void* value) {
    NSCollectionLayoutGroup* nSCollectionLayoutGroup = (NSCollectionLayoutGroup*)ptr;
    [nSCollectionLayoutGroup setInterItemSpacing:(NSCollectionLayoutSpacing*)value];
}
