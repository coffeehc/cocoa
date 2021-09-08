#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutBoundarySupplementaryItem_Alloc();

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(void* layoutSize, void* elementKind, int alignment);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(void* layoutSize, void* elementKind, int alignment, CGPoint absoluteOffset);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutBoundarySupplementaryItem_AllocCollectionLayoutBoundarySupplementaryItem();
void* C_NSCollectionLayoutBoundarySupplementaryItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutBoundarySupplementaryItem_Retain(void* ptr);
bool C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(void* ptr);
void C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(void* ptr, bool value);
CGPoint C_NSCollectionLayoutBoundarySupplementaryItem_Offset(void* ptr);
int C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(void* ptr);
bool C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(void* ptr);
void C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(void* ptr, bool value);
