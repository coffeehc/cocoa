#import <Appkit/Appkit.h>
#import "collection_layout_group.h"

void* C_CollectionLayoutGroup_Alloc() {
    return [NSCollectionLayoutGroup alloc];
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems) {
    NSMutableArray* objcSubitems = [[NSMutableArray alloc] init];
    void** subitemsData = (void**)subitems.data;
    for (int i = 0; i < subitems.len; i++) {
    	void* p = subitemsData[i];
    	[objcSubitems addObject:(NSCollectionLayoutItem*)(NSCollectionLayoutItem*)p];
    }
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup horizontalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitems:objcSubitems];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count) {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup horizontalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitem:(NSCollectionLayoutItem*)subitem count:count];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems) {
    NSMutableArray* objcSubitems = [[NSMutableArray alloc] init];
    void** subitemsData = (void**)subitems.data;
    for (int i = 0; i < subitems.len; i++) {
    	void* p = subitemsData[i];
    	[objcSubitems addObject:(NSCollectionLayoutItem*)(NSCollectionLayoutItem*)p];
    }
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup verticalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitems:objcSubitems];
    return result_;
}

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count) {
    NSCollectionLayoutGroup* result_ = [NSCollectionLayoutGroup verticalGroupWithLayoutSize:(NSCollectionLayoutSize*)layoutSize subitem:(NSCollectionLayoutItem*)subitem count:count];
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
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
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
