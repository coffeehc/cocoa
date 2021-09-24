#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionViewLayout_Alloc();

void* C_NSCollectionViewLayout_AllocCollectionViewLayout();
void* C_NSCollectionViewLayout_Init(void* ptr);
void* C_NSCollectionViewLayout_NewCollectionViewLayout();
void* C_NSCollectionViewLayout_Autorelease(void* ptr);
void* C_NSCollectionViewLayout_Retain(void* ptr);
void C_NSCollectionViewLayout_PrepareLayout(void* ptr);
Array C_NSCollectionViewLayout_LayoutAttributesForElementsInRect(void* ptr, CGRect rect);
void* C_NSCollectionViewLayout_LayoutAttributesForItemAtIndexPath(void* ptr, void* indexPath);
void* C_NSCollectionViewLayout_LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath);
void* C_NSCollectionViewLayout_LayoutAttributesForDecorationViewOfKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath);
void* C_NSCollectionViewLayout_LayoutAttributesForDropTargetAtPoint(void* ptr, CGPoint pointInCollectionView);
void* C_NSCollectionViewLayout_LayoutAttributesForInterItemGapBeforeIndexPath(void* ptr, void* indexPath);
CGPoint C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset(void* ptr, CGPoint proposedContentOffset);
CGPoint C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(void* ptr, CGPoint proposedContentOffset, CGPoint velocity);
void C_NSCollectionViewLayout_PrepareForCollectionViewUpdates(void* ptr, Array updateItems);
void C_NSCollectionViewLayout_FinalizeCollectionViewUpdates(void* ptr);
void* C_NSCollectionViewLayout_IndexPathsToInsertForSupplementaryViewOfKind(void* ptr, void* elementKind);
void* C_NSCollectionViewLayout_IndexPathsToInsertForDecorationViewOfKind(void* ptr, void* elementKind);
void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingItemAtIndexPath(void* ptr, void* itemIndexPath);
void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* elementIndexPath);
void* C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* decorationIndexPath);
void* C_NSCollectionViewLayout_IndexPathsToDeleteForSupplementaryViewOfKind(void* ptr, void* elementKind);
void* C_NSCollectionViewLayout_IndexPathsToDeleteForDecorationViewOfKind(void* ptr, void* elementKind);
void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingItemAtIndexPath(void* ptr, void* itemIndexPath);
void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* elementIndexPath);
void* C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(void* ptr, void* elementKind, void* decorationIndexPath);
void C_NSCollectionViewLayout_InvalidateLayout(void* ptr);
void C_NSCollectionViewLayout_InvalidateLayoutWithContext(void* ptr, void* context);
bool C_NSCollectionViewLayout_ShouldInvalidateLayoutForBoundsChange(void* ptr, CGRect newBounds);
bool C_NSCollectionViewLayout_ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(void* ptr, void* preferredAttributes, void* originalAttributes);
void* C_NSCollectionViewLayout_InvalidationContextForBoundsChange(void* ptr, CGRect newBounds);
void* C_NSCollectionViewLayout_InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(void* ptr, void* preferredAttributes, void* originalAttributes);
void C_NSCollectionViewLayout_PrepareForAnimatedBoundsChange(void* ptr, CGRect oldBounds);
void C_NSCollectionViewLayout_FinalizeAnimatedBoundsChange(void* ptr);
void C_NSCollectionViewLayout_RegisterNib_ForDecorationViewOfKind(void* ptr, void* nib, void* elementKind);
void C_NSCollectionViewLayout_PrepareForTransitionFromLayout(void* ptr, void* oldLayout);
void C_NSCollectionViewLayout_PrepareForTransitionToLayout(void* ptr, void* newLayout);
void C_NSCollectionViewLayout_FinalizeLayoutTransition(void* ptr);
void* C_NSCollectionViewLayout_CollectionView(void* ptr);
CGSize C_NSCollectionViewLayout_CollectionViewContentSize(void* ptr);

void* C_CollectionViewLayoutAttributes_Alloc();

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(void* indexPath);
void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(void* elementKind, void* indexPath);
void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(void* decorationViewKind, void* indexPath);
void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(void* indexPath);
void* C_NSCollectionViewLayoutAttributes_AllocCollectionViewLayoutAttributes();
void* C_NSCollectionViewLayoutAttributes_Init(void* ptr);
void* C_NSCollectionViewLayoutAttributes_NewCollectionViewLayoutAttributes();
void* C_NSCollectionViewLayoutAttributes_Autorelease(void* ptr);
void* C_NSCollectionViewLayoutAttributes_Retain(void* ptr);
int C_NSCollectionViewLayoutAttributes_RepresentedElementCategory(void* ptr);
void* C_NSCollectionViewLayoutAttributes_IndexPath(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetIndexPath(void* ptr, void* value);
void* C_NSCollectionViewLayoutAttributes_RepresentedElementKind(void* ptr);
CGRect C_NSCollectionViewLayoutAttributes_Frame(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetFrame(void* ptr, CGRect value);
CGSize C_NSCollectionViewLayoutAttributes_Size(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetSize(void* ptr, CGSize value);
double C_NSCollectionViewLayoutAttributes_Alpha(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetAlpha(void* ptr, double value);
bool C_NSCollectionViewLayoutAttributes_IsHidden(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetHidden(void* ptr, bool value);
int C_NSCollectionViewLayoutAttributes_ZIndex(void* ptr);
void C_NSCollectionViewLayoutAttributes_SetZIndex(void* ptr, int value);

void* C_CollectionViewUpdateItem_Alloc();

void* C_NSCollectionViewUpdateItem_AllocCollectionViewUpdateItem();
void* C_NSCollectionViewUpdateItem_Init(void* ptr);
void* C_NSCollectionViewUpdateItem_NewCollectionViewUpdateItem();
void* C_NSCollectionViewUpdateItem_Autorelease(void* ptr);
void* C_NSCollectionViewUpdateItem_Retain(void* ptr);
void* C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(void* ptr);
void* C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(void* ptr);
int C_NSCollectionViewUpdateItem_UpdateAction(void* ptr);

void* C_CollectionViewLayoutInvalidationContext_Alloc();

void* C_NSCollectionViewLayoutInvalidationContext_AllocCollectionViewLayoutInvalidationContext();
void* C_NSCollectionViewLayoutInvalidationContext_Init(void* ptr);
void* C_NSCollectionViewLayoutInvalidationContext_NewCollectionViewLayoutInvalidationContext();
void* C_NSCollectionViewLayoutInvalidationContext_Autorelease(void* ptr);
void* C_NSCollectionViewLayoutInvalidationContext_Retain(void* ptr);
void C_NSCollectionViewLayoutInvalidationContext_InvalidateItemsAtIndexPaths(void* ptr, void* indexPaths);
void C_NSCollectionViewLayoutInvalidationContext_InvalidateSupplementaryElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths);
void C_NSCollectionViewLayoutInvalidationContext_InvalidateDecorationElementsOfKind_AtIndexPaths(void* ptr, void* elementKind, void* indexPaths);
bool C_NSCollectionViewLayoutInvalidationContext_InvalidateEverything(void* ptr);
bool C_NSCollectionViewLayoutInvalidationContext_InvalidateDataSourceCounts(void* ptr);
CGPoint C_NSCollectionViewLayoutInvalidationContext_ContentOffsetAdjustment(void* ptr);
void C_NSCollectionViewLayoutInvalidationContext_SetContentOffsetAdjustment(void* ptr, CGPoint value);
CGSize C_NSCollectionViewLayoutInvalidationContext_ContentSizeAdjustment(void* ptr);
void C_NSCollectionViewLayoutInvalidationContext_SetContentSizeAdjustment(void* ptr, CGSize value);
void* C_NSCollectionViewLayoutInvalidationContext_InvalidatedItemIndexPaths(void* ptr);
