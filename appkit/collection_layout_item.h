#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutItem_Alloc();

void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutItem_AllocCollectionLayoutItem();
void* C_NSCollectionLayoutItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutItem_Retain(void* ptr);
void* C_NSCollectionLayoutItem_LayoutSize(void* ptr);
Array C_NSCollectionLayoutItem_SupplementaryItems(void* ptr);
void* C_NSCollectionLayoutItem_EdgeSpacing(void* ptr);
void C_NSCollectionLayoutItem_SetEdgeSpacing(void* ptr, void* value);
NSDirectionalEdgeInsets C_NSCollectionLayoutItem_ContentInsets(void* ptr);
void C_NSCollectionLayoutItem_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value);
