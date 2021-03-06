#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionViewFlowLayout_Alloc();

void* C_NSCollectionViewFlowLayout_AllocCollectionViewFlowLayout();
void* C_NSCollectionViewFlowLayout_Init(void* ptr);
void* C_NSCollectionViewFlowLayout_NewCollectionViewFlowLayout();
void* C_NSCollectionViewFlowLayout_Autorelease(void* ptr);
void* C_NSCollectionViewFlowLayout_Retain(void* ptr);
void C_NSCollectionViewFlowLayout_CollapseSectionAtIndex(void* ptr, unsigned int sectionIndex);
void C_NSCollectionViewFlowLayout_ExpandSectionAtIndex(void* ptr, unsigned int sectionIndex);
bool C_NSCollectionViewFlowLayout_SectionAtIndexIsCollapsed(void* ptr, unsigned int sectionIndex);
int C_NSCollectionViewFlowLayout_ScrollDirection(void* ptr);
void C_NSCollectionViewFlowLayout_SetScrollDirection(void* ptr, int value);
double C_NSCollectionViewFlowLayout_MinimumLineSpacing(void* ptr);
void C_NSCollectionViewFlowLayout_SetMinimumLineSpacing(void* ptr, double value);
double C_NSCollectionViewFlowLayout_MinimumInteritemSpacing(void* ptr);
void C_NSCollectionViewFlowLayout_SetMinimumInteritemSpacing(void* ptr, double value);
CGSize C_NSCollectionViewFlowLayout_EstimatedItemSize(void* ptr);
void C_NSCollectionViewFlowLayout_SetEstimatedItemSize(void* ptr, CGSize value);
CGSize C_NSCollectionViewFlowLayout_ItemSize(void* ptr);
void C_NSCollectionViewFlowLayout_SetItemSize(void* ptr, CGSize value);
NSEdgeInsets C_NSCollectionViewFlowLayout_SectionInset(void* ptr);
void C_NSCollectionViewFlowLayout_SetSectionInset(void* ptr, NSEdgeInsets value);
CGSize C_NSCollectionViewFlowLayout_HeaderReferenceSize(void* ptr);
void C_NSCollectionViewFlowLayout_SetHeaderReferenceSize(void* ptr, CGSize value);
CGSize C_NSCollectionViewFlowLayout_FooterReferenceSize(void* ptr);
void C_NSCollectionViewFlowLayout_SetFooterReferenceSize(void* ptr, CGSize value);
bool C_NSCollectionViewFlowLayout_SectionFootersPinToVisibleBounds(void* ptr);
void C_NSCollectionViewFlowLayout_SetSectionFootersPinToVisibleBounds(void* ptr, bool value);
bool C_NSCollectionViewFlowLayout_SectionHeadersPinToVisibleBounds(void* ptr);
void C_NSCollectionViewFlowLayout_SetSectionHeadersPinToVisibleBounds(void* ptr, bool value);

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

void* WrapCollectionViewDelegateFlowLayout(uintptr_t goID);
