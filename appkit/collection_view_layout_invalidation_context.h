#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

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
