#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_CollectionLayoutSpacing_Alloc();

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(double fixedSpacing);
void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(double flexibleSpacing);
void* C_NSCollectionLayoutSpacing_AllocCollectionLayoutSpacing();
void* C_NSCollectionLayoutSpacing_Autorelease(void* ptr);
void* C_NSCollectionLayoutSpacing_Retain(void* ptr);
double C_NSCollectionLayoutSpacing_Spacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFixedSpacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFlexibleSpacing(void* ptr);
