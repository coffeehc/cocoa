#import <Appkit/Appkit.h>
#import "collection_layout_section.h"

void* C_CollectionLayoutSection_Alloc() {
    return [NSCollectionLayoutSection alloc];
}

void* C_NSCollectionLayoutSection_CollectionLayoutSection_SectionWithGroup(void* group) {
    NSCollectionLayoutSection* result_ = [NSCollectionLayoutSection sectionWithGroup:(NSCollectionLayoutGroup*)group];
    return result_;
}

void* C_NSCollectionLayoutSection_AllocCollectionLayoutSection() {
    NSCollectionLayoutSection* result_ = [NSCollectionLayoutSection alloc];
    return result_;
}

void* C_NSCollectionLayoutSection_Autorelease(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSCollectionLayoutSection* result_ = [nSCollectionLayoutSection autorelease];
    return result_;
}

void* C_NSCollectionLayoutSection_Retain(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSCollectionLayoutSection* result_ = [nSCollectionLayoutSection retain];
    return result_;
}

int C_NSCollectionLayoutSection_OrthogonalScrollingBehavior(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSCollectionLayoutSectionOrthogonalScrollingBehavior result_ = [nSCollectionLayoutSection orthogonalScrollingBehavior];
    return result_;
}

void C_NSCollectionLayoutSection_SetOrthogonalScrollingBehavior(void* ptr, int value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    [nSCollectionLayoutSection setOrthogonalScrollingBehavior:value];
}

double C_NSCollectionLayoutSection_InterGroupSpacing(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    CGFloat result_ = [nSCollectionLayoutSection interGroupSpacing];
    return result_;
}

void C_NSCollectionLayoutSection_SetInterGroupSpacing(void* ptr, double value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    [nSCollectionLayoutSection setInterGroupSpacing:value];
}

NSDirectionalEdgeInsets C_NSCollectionLayoutSection_ContentInsets(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSDirectionalEdgeInsets result_ = [nSCollectionLayoutSection contentInsets];
    return result_;
}

void C_NSCollectionLayoutSection_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    [nSCollectionLayoutSection setContentInsets:value];
}

bool C_NSCollectionLayoutSection_SupplementariesFollowContentInsets(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    BOOL result_ = [nSCollectionLayoutSection supplementariesFollowContentInsets];
    return result_;
}

void C_NSCollectionLayoutSection_SetSupplementariesFollowContentInsets(void* ptr, bool value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    [nSCollectionLayoutSection setSupplementariesFollowContentInsets:value];
}

Array C_NSCollectionLayoutSection_BoundarySupplementaryItems(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSArray* result_ = [nSCollectionLayoutSection boundarySupplementaryItems];
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

void C_NSCollectionLayoutSection_SetBoundarySupplementaryItems(void* ptr, Array value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSCollectionLayoutBoundarySupplementaryItem*)p];
    	}
    }
    [nSCollectionLayoutSection setBoundarySupplementaryItems:objcValue];
}

Array C_NSCollectionLayoutSection_DecorationItems(void* ptr) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSArray* result_ = [nSCollectionLayoutSection decorationItems];
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

void C_NSCollectionLayoutSection_SetDecorationItems(void* ptr, Array value) {
    NSCollectionLayoutSection* nSCollectionLayoutSection = (NSCollectionLayoutSection*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSCollectionLayoutDecorationItem*)p];
    	}
    }
    [nSCollectionLayoutSection setDecorationItems:objcValue];
}
