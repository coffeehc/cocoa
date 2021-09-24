#import "collection_view_compositional_layout.h"
#import <AppKit/NSCollectionViewCompositionalLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionViewCompositionalLayout_Alloc() {
    return [NSCollectionViewCompositionalLayout alloc];
}

void* C_NSCollectionViewCompositionalLayout_InitWithSection(void* ptr, void* section) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout initWithSection:(NSCollectionLayoutSection*)section];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(void* ptr, void* section, void* configuration) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout initWithSection:(NSCollectionLayoutSection*)section configuration:(NSCollectionViewCompositionalLayoutConfiguration*)configuration];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_AllocCollectionViewCompositionalLayout() {
    NSCollectionViewCompositionalLayout* result_ = [NSCollectionViewCompositionalLayout alloc];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Autorelease(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout autorelease];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Retain(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout retain];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Configuration(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayout configuration];
    return result_;
}

void C_NSCollectionViewCompositionalLayout_SetConfiguration(void* ptr, void* value) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    [nSCollectionViewCompositionalLayout setConfiguration:(NSCollectionViewCompositionalLayoutConfiguration*)value];
}

void* C_CollectionViewCompositionalLayoutConfiguration_Alloc() {
    return [NSCollectionViewCompositionalLayoutConfiguration alloc];
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_AllocCollectionViewCompositionalLayoutConfiguration() {
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [NSCollectionViewCompositionalLayoutConfiguration alloc];
    return result_;
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_Init(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayoutConfiguration init];
    return result_;
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_NewCollectionViewCompositionalLayoutConfiguration() {
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [NSCollectionViewCompositionalLayoutConfiguration new];
    return result_;
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_Autorelease(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayoutConfiguration autorelease];
    return result_;
}

void* C_NSCollectionViewCompositionalLayoutConfiguration_Retain(void* ptr) {
    NSCollectionViewCompositionalLayoutConfiguration* nSCollectionViewCompositionalLayoutConfiguration = (NSCollectionViewCompositionalLayoutConfiguration*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayoutConfiguration retain];
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
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSCollectionLayoutBoundarySupplementaryItem*)p];
    	}
    }
    [nSCollectionViewCompositionalLayoutConfiguration setBoundarySupplementaryItems:objcValue];
}

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

void* C_CollectionLayoutSpacing_Alloc() {
    return [NSCollectionLayoutSpacing alloc];
}

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(double fixedSpacing) {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing fixedSpacing:fixedSpacing];
    return result_;
}

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(double flexibleSpacing) {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing flexibleSpacing:flexibleSpacing];
    return result_;
}

void* C_NSCollectionLayoutSpacing_AllocCollectionLayoutSpacing() {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing alloc];
    return result_;
}

void* C_NSCollectionLayoutSpacing_Autorelease(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutSpacing autorelease];
    return result_;
}

void* C_NSCollectionLayoutSpacing_Retain(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutSpacing retain];
    return result_;
}

double C_NSCollectionLayoutSpacing_Spacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    CGFloat result_ = [nSCollectionLayoutSpacing spacing];
    return result_;
}

bool C_NSCollectionLayoutSpacing_IsFixedSpacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    BOOL result_ = [nSCollectionLayoutSpacing isFixedSpacing];
    return result_;
}

bool C_NSCollectionLayoutSpacing_IsFlexibleSpacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    BOOL result_ = [nSCollectionLayoutSpacing isFlexibleSpacing];
    return result_;
}

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

void* C_CollectionLayoutGroupCustomItem_Alloc() {
    return [NSCollectionLayoutGroupCustomItem alloc];
}

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(CGRect frame) {
    NSCollectionLayoutGroupCustomItem* result_ = [NSCollectionLayoutGroupCustomItem customItemWithFrame:frame];
    return result_;
}

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(CGRect frame, int zIndex) {
    NSCollectionLayoutGroupCustomItem* result_ = [NSCollectionLayoutGroupCustomItem customItemWithFrame:frame zIndex:zIndex];
    return result_;
}

void* C_NSCollectionLayoutGroupCustomItem_AllocCollectionLayoutGroupCustomItem() {
    NSCollectionLayoutGroupCustomItem* result_ = [NSCollectionLayoutGroupCustomItem alloc];
    return result_;
}

void* C_NSCollectionLayoutGroupCustomItem_Autorelease(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSCollectionLayoutGroupCustomItem* result_ = [nSCollectionLayoutGroupCustomItem autorelease];
    return result_;
}

void* C_NSCollectionLayoutGroupCustomItem_Retain(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSCollectionLayoutGroupCustomItem* result_ = [nSCollectionLayoutGroupCustomItem retain];
    return result_;
}

CGRect C_NSCollectionLayoutGroupCustomItem_Frame(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSRect result_ = [nSCollectionLayoutGroupCustomItem frame];
    return result_;
}

int C_NSCollectionLayoutGroupCustomItem_ZIndex(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSInteger result_ = [nSCollectionLayoutGroupCustomItem zIndex];
    return result_;
}

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

void* C_CollectionLayoutSize_Alloc() {
    return [NSCollectionLayoutSize alloc];
}

void* C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(void* width, void* height) {
    NSCollectionLayoutSize* result_ = [NSCollectionLayoutSize sizeWithWidthDimension:(NSCollectionLayoutDimension*)width heightDimension:(NSCollectionLayoutDimension*)height];
    return result_;
}

void* C_NSCollectionLayoutSize_AllocCollectionLayoutSize() {
    NSCollectionLayoutSize* result_ = [NSCollectionLayoutSize alloc];
    return result_;
}

void* C_NSCollectionLayoutSize_Autorelease(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutSize* result_ = [nSCollectionLayoutSize autorelease];
    return result_;
}

void* C_NSCollectionLayoutSize_Retain(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutSize* result_ = [nSCollectionLayoutSize retain];
    return result_;
}

void* C_NSCollectionLayoutSize_WidthDimension(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutSize widthDimension];
    return result_;
}

void* C_NSCollectionLayoutSize_HeightDimension(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutSize heightDimension];
    return result_;
}

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

void* C_CollectionLayoutAnchor_Alloc() {
    return [NSCollectionLayoutAnchor alloc];
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(unsigned int edges) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges];
    return result_;
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(unsigned int edges, CGPoint absoluteOffset) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges absoluteOffset:absoluteOffset];
    return result_;
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(unsigned int edges, CGPoint fractionalOffset) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges fractionalOffset:fractionalOffset];
    return result_;
}

void* C_NSCollectionLayoutAnchor_AllocCollectionLayoutAnchor() {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor alloc];
    return result_;
}

void* C_NSCollectionLayoutAnchor_Autorelease(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSCollectionLayoutAnchor* result_ = [nSCollectionLayoutAnchor autorelease];
    return result_;
}

void* C_NSCollectionLayoutAnchor_Retain(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSCollectionLayoutAnchor* result_ = [nSCollectionLayoutAnchor retain];
    return result_;
}

unsigned int C_NSCollectionLayoutAnchor_Edges(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSDirectionalRectEdge result_ = [nSCollectionLayoutAnchor edges];
    return result_;
}

CGPoint C_NSCollectionLayoutAnchor_Offset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSPoint result_ = [nSCollectionLayoutAnchor offset];
    return result_;
}

bool C_NSCollectionLayoutAnchor_IsAbsoluteOffset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    BOOL result_ = [nSCollectionLayoutAnchor isAbsoluteOffset];
    return result_;
}

bool C_NSCollectionLayoutAnchor_IsFractionalOffset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    BOOL result_ = [nSCollectionLayoutAnchor isFractionalOffset];
    return result_;
}

void* C_CollectionLayoutDimension_Alloc() {
    return [NSCollectionLayoutDimension alloc];
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(double absoluteDimension) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension absoluteDimension:absoluteDimension];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(double estimatedDimension) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension estimatedDimension:estimatedDimension];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(double fractionalHeight) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension fractionalHeightDimension:fractionalHeight];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(double fractionalWidth) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension fractionalWidthDimension:fractionalWidth];
    return result_;
}

void* C_NSCollectionLayoutDimension_AllocCollectionLayoutDimension() {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension alloc];
    return result_;
}

void* C_NSCollectionLayoutDimension_Autorelease(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutDimension autorelease];
    return result_;
}

void* C_NSCollectionLayoutDimension_Retain(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutDimension retain];
    return result_;
}

double C_NSCollectionLayoutDimension_Dimension(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    CGFloat result_ = [nSCollectionLayoutDimension dimension];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsAbsolute(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isAbsolute];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsEstimated(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isEstimated];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsFractionalHeight(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isFractionalHeight];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsFractionalWidth(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isFractionalWidth];
    return result_;
}

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
