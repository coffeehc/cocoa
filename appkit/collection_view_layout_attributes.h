#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

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
