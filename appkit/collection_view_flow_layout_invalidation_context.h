#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewFlowLayoutInvalidationContext_Alloc();

void* C_NSCollectionViewFlowLayoutInvalidationContext_AllocCollectionViewFlowLayoutInvalidationContext();
void* C_NSCollectionViewFlowLayoutInvalidationContext_Init(void* ptr);
void* C_NSCollectionViewFlowLayoutInvalidationContext_NewCollectionViewFlowLayoutInvalidationContext();
void* C_NSCollectionViewFlowLayoutInvalidationContext_Autorelease(void* ptr);
void* C_NSCollectionViewFlowLayoutInvalidationContext_Retain(void* ptr);
bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(void* ptr);
void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(void* ptr, bool value);
bool C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(void* ptr);
void C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(void* ptr, bool value);
