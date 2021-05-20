#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutSpacing_Alloc();

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(double fixedSpacing);
void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(double flexibleSpacing);
double C_NSCollectionLayoutSpacing_Spacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFixedSpacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFlexibleSpacing(void* ptr);
