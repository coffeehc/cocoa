#import <AppKit/NSCollectionViewCompositionalLayout.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionLayoutSection_Alloc();

void* C_NSCollectionLayoutSection_CollectionLayoutSection_SectionWithGroup(void* group);
void* C_NSCollectionLayoutSection_AllocCollectionLayoutSection();
void* C_NSCollectionLayoutSection_Autorelease(void* ptr);
void* C_NSCollectionLayoutSection_Retain(void* ptr);
int C_NSCollectionLayoutSection_OrthogonalScrollingBehavior(void* ptr);
void C_NSCollectionLayoutSection_SetOrthogonalScrollingBehavior(void* ptr, int value);
double C_NSCollectionLayoutSection_InterGroupSpacing(void* ptr);
void C_NSCollectionLayoutSection_SetInterGroupSpacing(void* ptr, double value);
NSDirectionalEdgeInsets C_NSCollectionLayoutSection_ContentInsets(void* ptr);
void C_NSCollectionLayoutSection_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value);
bool C_NSCollectionLayoutSection_SupplementariesFollowContentInsets(void* ptr);
void C_NSCollectionLayoutSection_SetSupplementariesFollowContentInsets(void* ptr, bool value);
Array C_NSCollectionLayoutSection_BoundarySupplementaryItems(void* ptr);
void C_NSCollectionLayoutSection_SetBoundarySupplementaryItems(void* ptr, Array value);
Array C_NSCollectionLayoutSection_DecorationItems(void* ptr);
void C_NSCollectionLayoutSection_SetDecorationItems(void* ptr, Array value);
