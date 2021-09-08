#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

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
