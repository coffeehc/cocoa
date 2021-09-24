#import "collection_view_layout.h"
#import <AppKit/NSCollectionViewLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionViewLayout_Alloc() {
    return [NSCollectionViewLayout alloc];
}

void* C_NSCollectionViewLayout_AllocCollectionViewLayout() {
    NSCollectionViewLayout* result_ = [NSCollectionViewLayout alloc];
    return result_;
}

void* C_NSCollectionViewLayout_Init(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionViewLayout init];
    return result_;
}

void* C_NSCollectionViewLayout_NewCollectionViewLayout() {
    NSCollectionViewLayout* result_ = [NSCollectionViewLayout new];
    return result_;
}

void* C_NSCollectionViewLayout_Autorelease(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionViewLayout autorelease];
    return result_;
}

void* C_NSCollectionViewLayout_Retain(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionViewLayout retain];
    return result_;
}

void C_NSCollectionViewLayout_PrepareLayout(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout prepareLayout];
}

Array C_NSCollectionViewLayout_LayoutAttributesForElementsInRect(void* ptr, CGRect rect) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSArray* result_ = [nSCollectionViewLayout layoutAttributesForElementsInRect:rect];
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

void* C_NSCollectionViewLayout_LayoutAttributesForItemAtIndexPath(void* ptr, void* indexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout layoutAttributesForItemAtIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayout_LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout layoutAttributesForSupplementaryViewOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayout_LayoutAttributesForDecorationViewOfKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout layoutAttributesForDecorationViewOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayout_LayoutAttributesForDropTargetAtPoint(void* ptr, CGPoint pointInCollectionView) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout layoutAttributesForDropTargetAtPoint:pointInCollectionView];
    return result_;
}

void* C_NSCollectionViewLayout_LayoutAttributesForInterItemGapBeforeIndexPath(void* ptr, void* indexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout layoutAttributesForInterItemGapBeforeIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

CGPoint C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset(void* ptr, CGPoint proposedContentOffset) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSPoint result_ = [nSCollectionViewLayout targetContentOffsetForProposedContentOffset:proposedContentOffset];
    return result_;
}

CGPoint C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(void* ptr, CGPoint proposedContentOffset, CGPoint velocity) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSPoint result_ = [nSCollectionViewLayout targetContentOffsetForProposedContentOffset:proposedContentOffset withScrollingVelocity:velocity];
    return result_;
}

void C_NSCollectionViewLayout_PrepareForCollectionViewUpdates(void* ptr, Array updateItems) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSMutableArray* objcUpdateItems = [NSMutableArray arrayWithCapacity:updateItems.len];
    if (updateItems.len > 0) {
    	void** updateItemsData = (void**)updateItems.data;
    	for (int i = 0; i < updateItems.len; i++) {
    		void* p = updateItemsData[i];
    		[objcUpdateItems addObject:(NSCollectionViewUpdateItem*)p];
    	}
    }
    [nSCollectionViewLayout prepareForCollectionViewUpdates:objcUpdateItems];
}

void C_NSCollectionViewLayout_FinalizeCollectionViewUpdates(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout finalizeCollectionViewUpdates];
}

void* C_NSCollectionViewLayout_IndexPathsToInsertForSupplementaryViewOfKind(void* ptr, void* elementKind) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSSet* result_ = [nSCollectionViewLayout indexPathsToInsertForSupplementaryViewOfKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionViewLayout_IndexPathsToInsertForDecorationViewOfKind(void* ptr, void* elementKind) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSSet* result_ = [nSCollectionViewLayout indexPathsToInsertForDecorationViewOfKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingItemAtIndexPath(void* ptr, void* itemIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout initialLayoutAttributesForAppearingItemAtIndexPath:(NSIndexPath*)itemIndexPath];
    return result_;
}

void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* elementIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout initialLayoutAttributesForAppearingSupplementaryElementOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)elementIndexPath];
    return result_;
}

void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* decorationIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout initialLayoutAttributesForAppearingDecorationElementOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)decorationIndexPath];
    return result_;
}

void* C_NSCollectionViewLayout_IndexPathsToDeleteForSupplementaryViewOfKind(void* ptr, void* elementKind) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSSet* result_ = [nSCollectionViewLayout indexPathsToDeleteForSupplementaryViewOfKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionViewLayout_IndexPathsToDeleteForDecorationViewOfKind(void* ptr, void* elementKind) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSSet* result_ = [nSCollectionViewLayout indexPathsToDeleteForDecorationViewOfKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingItemAtIndexPath(void* ptr, void* itemIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout finalLayoutAttributesForDisappearingItemAtIndexPath:(NSIndexPath*)itemIndexPath];
    return result_;
}

void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* elementIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout finalLayoutAttributesForDisappearingSupplementaryElementOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)elementIndexPath];
    return result_;
}

void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* decorationIndexPath) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayout finalLayoutAttributesForDisappearingDecorationElementOfKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)decorationIndexPath];
    return result_;
}

void C_NSCollectionViewLayout_InvalidateLayout(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout invalidateLayout];
}

void C_NSCollectionViewLayout_InvalidateLayoutWithContext(void* ptr, void* context) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout invalidateLayoutWithContext:(NSCollectionViewLayoutInvalidationContext*)context];
}

bool C_NSCollectionViewLayout_ShouldInvalidateLayoutForBoundsChange(void* ptr, CGRect newBounds) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    BOOL result_ = [nSCollectionViewLayout shouldInvalidateLayoutForBoundsChange:newBounds];
    return result_;
}

bool C_NSCollectionViewLayout_ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(void* ptr, void* preferredAttributes, void* originalAttributes) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    BOOL result_ = [nSCollectionViewLayout shouldInvalidateLayoutForPreferredLayoutAttributes:(NSCollectionViewLayoutAttributes*)preferredAttributes withOriginalAttributes:(NSCollectionViewLayoutAttributes*)originalAttributes];
    return result_;
}

void* C_NSCollectionViewLayout_InvalidationContextForBoundsChange(void* ptr, CGRect newBounds) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayout invalidationContextForBoundsChange:newBounds];
    return result_;
}

void* C_NSCollectionViewLayout_InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(void* ptr, void* preferredAttributes, void* originalAttributes) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayout invalidationContextForPreferredLayoutAttributes:(NSCollectionViewLayoutAttributes*)preferredAttributes withOriginalAttributes:(NSCollectionViewLayoutAttributes*)originalAttributes];
    return result_;
}

void C_NSCollectionViewLayout_PrepareForAnimatedBoundsChange(void* ptr, CGRect oldBounds) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout prepareForAnimatedBoundsChange:oldBounds];
}

void C_NSCollectionViewLayout_FinalizeAnimatedBoundsChange(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout finalizeAnimatedBoundsChange];
}

void C_NSCollectionViewLayout_RegisterNib_ForDecorationViewOfKind(void* ptr, void* nib, void* elementKind) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout registerNib:(NSNib*)nib forDecorationViewOfKind:(NSString*)elementKind];
}

void C_NSCollectionViewLayout_PrepareForTransitionFromLayout(void* ptr, void* oldLayout) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout prepareForTransitionFromLayout:(NSCollectionViewLayout*)oldLayout];
}

void C_NSCollectionViewLayout_PrepareForTransitionToLayout(void* ptr, void* newLayout) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout prepareForTransitionToLayout:(NSCollectionViewLayout*)newLayout];
}

void C_NSCollectionViewLayout_FinalizeLayoutTransition(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    [nSCollectionViewLayout finalizeLayoutTransition];
}

void* C_NSCollectionViewLayout_CollectionView(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSCollectionView* result_ = [nSCollectionViewLayout collectionView];
    return result_;
}

CGSize C_NSCollectionViewLayout_CollectionViewContentSize(void* ptr) {
    NSCollectionViewLayout* nSCollectionViewLayout = (NSCollectionViewLayout*)ptr;
    NSSize result_ = [nSCollectionViewLayout collectionViewContentSize];
    return result_;
}

void* C_CollectionViewLayoutAttributes_Alloc() {
    return [NSCollectionViewLayoutAttributes alloc];
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForItemWithIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(void* elementKind, void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForSupplementaryViewOfKind:(NSString*)elementKind withIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(void* decorationViewKind, void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForDecorationViewOfKind:(NSString*)decorationViewKind withIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForInterItemGapBeforeIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_AllocCollectionViewLayoutAttributes() {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes alloc];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Init(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes init];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_NewCollectionViewLayoutAttributes() {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes new];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Autorelease(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes autorelease];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Retain(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes retain];
    return result_;
}

int C_NSCollectionViewLayoutAttributes_RepresentedElementCategory(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionElementCategory result_ = [nSCollectionViewLayoutAttributes representedElementCategory];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_IndexPath(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSIndexPath* result_ = [nSCollectionViewLayoutAttributes indexPath];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetIndexPath(void* ptr, void* value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setIndexPath:(NSIndexPath*)value];
}

void* C_NSCollectionViewLayoutAttributes_RepresentedElementKind(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSString* result_ = [nSCollectionViewLayoutAttributes representedElementKind];
    return result_;
}

CGRect C_NSCollectionViewLayoutAttributes_Frame(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSRect result_ = [nSCollectionViewLayoutAttributes frame];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetFrame(void* ptr, CGRect value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setFrame:value];
}

CGSize C_NSCollectionViewLayoutAttributes_Size(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSSize result_ = [nSCollectionViewLayoutAttributes size];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetSize(void* ptr, CGSize value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setSize:value];
}

double C_NSCollectionViewLayoutAttributes_Alpha(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    CGFloat result_ = [nSCollectionViewLayoutAttributes alpha];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetAlpha(void* ptr, double value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setAlpha:value];
}

bool C_NSCollectionViewLayoutAttributes_IsHidden(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    BOOL result_ = [nSCollectionViewLayoutAttributes isHidden];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetHidden(void* ptr, bool value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setHidden:value];
}

int C_NSCollectionViewLayoutAttributes_ZIndex(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSInteger result_ = [nSCollectionViewLayoutAttributes zIndex];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetZIndex(void* ptr, int value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setZIndex:value];
}

void* C_CollectionViewUpdateItem_Alloc() {
    return [NSCollectionViewUpdateItem alloc];
}

void* C_NSCollectionViewUpdateItem_AllocCollectionViewUpdateItem() {
    NSCollectionViewUpdateItem* result_ = [NSCollectionViewUpdateItem alloc];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Init(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem init];
    return result_;
}

void* C_NSCollectionViewUpdateItem_NewCollectionViewUpdateItem() {
    NSCollectionViewUpdateItem* result_ = [NSCollectionViewUpdateItem new];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Autorelease(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem autorelease];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Retain(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem retain];
    return result_;
}

void* C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSIndexPath* result_ = [nSCollectionViewUpdateItem indexPathBeforeUpdate];
    return result_;
}

void* C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSIndexPath* result_ = [nSCollectionViewUpdateItem indexPathAfterUpdate];
    return result_;
}

int C_NSCollectionViewUpdateItem_UpdateAction(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionUpdateAction result_ = [nSCollectionViewUpdateItem updateAction];
    return result_;
}

void* C_CollectionViewLayoutInvalidationContext_Alloc() {
    return [NSCollectionViewLayoutInvalidationContext alloc];
}

void* C_NSCollectionViewLayoutInvalidationContext_AllocCollectionViewLayoutInvalidationContext() {
    NSCollectionViewLayoutInvalidationContext* result_ = [NSCollectionViewLayoutInvalidationContext alloc];
    return result_;
}

void* C_NSCollectionViewLayoutInvalidationContext_Init(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayoutInvalidationContext init];
    return result_;
}

void* C_NSCollectionViewLayoutInvalidationContext_NewCollectionViewLayoutInvalidationContext() {
    NSCollectionViewLayoutInvalidationContext* result_ = [NSCollectionViewLayoutInvalidationContext new];
    return result_;
}

void* C_NSCollectionViewLayoutInvalidationContext_Autorelease(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayoutInvalidationContext autorelease];
    return result_;
}

void* C_NSCollectionViewLayoutInvalidationContext_Retain(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSCollectionViewLayoutInvalidationContext* result_ = [nSCollectionViewLayoutInvalidationContext retain];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateItemsAtIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateSupplementaryElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateSupplementaryElementsOfKind:(NSString*)elementKind atIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionViewLayoutInvalidationContext_InvalidateDecorationElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext invalidateDecorationElementsOfKind:(NSString*)elementKind atIndexPaths:(NSSet*)indexPaths];
}

bool C_NSCollectionViewLayoutInvalidationContext_InvalidateEverything(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewLayoutInvalidationContext invalidateEverything];
    return result_;
}

bool C_NSCollectionViewLayoutInvalidationContext_InvalidateDataSourceCounts(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    BOOL result_ = [nSCollectionViewLayoutInvalidationContext invalidateDataSourceCounts];
    return result_;
}

CGPoint C_NSCollectionViewLayoutInvalidationContext_ContentOffsetAdjustment(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSPoint result_ = [nSCollectionViewLayoutInvalidationContext contentOffsetAdjustment];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_SetContentOffsetAdjustment(void* ptr, CGPoint value) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext setContentOffsetAdjustment:value];
}

CGSize C_NSCollectionViewLayoutInvalidationContext_ContentSizeAdjustment(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSSize result_ = [nSCollectionViewLayoutInvalidationContext contentSizeAdjustment];
    return result_;
}

void C_NSCollectionViewLayoutInvalidationContext_SetContentSizeAdjustment(void* ptr, CGSize value) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    [nSCollectionViewLayoutInvalidationContext setContentSizeAdjustment:value];
}

void* C_NSCollectionViewLayoutInvalidationContext_InvalidatedItemIndexPaths(void* ptr) {
    NSCollectionViewLayoutInvalidationContext* nSCollectionViewLayoutInvalidationContext = (NSCollectionViewLayoutInvalidationContext*)ptr;
    NSSet* result_ = [nSCollectionViewLayoutInvalidationContext invalidatedItemIndexPaths];
    return result_;
}
