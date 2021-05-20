#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutSupplementaryItem_Alloc();

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor);
void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor);
void* C_NSCollectionLayoutSupplementaryItem_ItemAnchor(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_ContainerAnchor(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_ElementKind(void* ptr);
int C_NSCollectionLayoutSupplementaryItem_ZIndex(void* ptr);
void C_NSCollectionLayoutSupplementaryItem_SetZIndex(void* ptr, int value);
