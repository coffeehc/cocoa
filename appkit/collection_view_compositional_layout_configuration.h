#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionViewCompositionalLayoutConfiguration_Alloc();

void* C_NSCollectionViewCompositionalLayoutConfiguration_AllocCollectionViewCompositionalLayoutConfiguration();
void* C_NSCollectionViewCompositionalLayoutConfiguration_Init(void* ptr);
void* C_NSCollectionViewCompositionalLayoutConfiguration_NewCollectionViewCompositionalLayoutConfiguration();
void* C_NSCollectionViewCompositionalLayoutConfiguration_Autorelease(void* ptr);
void* C_NSCollectionViewCompositionalLayoutConfiguration_Retain(void* ptr);
int C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(void* ptr, int value);
double C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(void* ptr, double value);
Array C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(void* ptr, Array value);
