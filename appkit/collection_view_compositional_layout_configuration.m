#import <Appkit/Appkit.h>
#import "collection_view_compositional_layout_configuration.h"

void* C_CollectionViewCompositionalLayoutConfiguration_Alloc() {
    return [NSCollectionViewCompositionalLayoutConfiguration alloc];
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_Init(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayoutConfiguration init];
    return result_;
}

int C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSCollectionViewScrollDirection result_ = [nSCollectionViewCompositionalLayoutConfiguration scrollDirection];
    return result_;
}

void C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(void* ptr, int value) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    [nSCollectionViewCompositionalLayoutConfiguration setScrollDirection:value];
}

double C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    CGFloat result_ = [nSCollectionViewCompositionalLayoutConfiguration interSectionSpacing];
    return result_;
}

void C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(void* ptr, double value) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    [nSCollectionViewCompositionalLayoutConfiguration setInterSectionSpacing:value];
}

Array C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSArray* result_ = [nSCollectionViewCompositionalLayoutConfiguration boundarySupplementaryItems];
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

void C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(void* ptr, Array value) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSCollectionLayoutBoundarySupplementaryItem*)(NSCollectionLayoutBoundarySupplementaryItem*)p];
    	}
    }
    [nSCollectionViewCompositionalLayoutConfiguration setBoundarySupplementaryItems:objcValue];
}
