#import <Appkit/Appkit.h>
#import "collection_view_layout.h"

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
    		[objcUpdateItems addObject:(NSCollectionViewUpdateItem*)(NSCollectionViewUpdateItem*)p];
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
