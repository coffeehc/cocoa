#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionLayoutDecorationItem_Alloc();

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(void* elementKind);
void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutDecorationItem_AllocCollectionLayoutDecorationItem();
void* C_NSCollectionLayoutDecorationItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutDecorationItem_Retain(void* ptr);
void* C_NSCollectionLayoutDecorationItem_ElementKind(void* ptr);
int C_NSCollectionLayoutDecorationItem_ZIndex(void* ptr);
void C_NSCollectionLayoutDecorationItem_SetZIndex(void* ptr, int value);
