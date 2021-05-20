#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutDimension_Alloc();

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(double absoluteDimension);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(double estimatedDimension);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(double fractionalHeight);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(double fractionalWidth);
double C_NSCollectionLayoutDimension_Dimension(void* ptr);
bool C_NSCollectionLayoutDimension_IsAbsolute(void* ptr);
bool C_NSCollectionLayoutDimension_IsEstimated(void* ptr);
bool C_NSCollectionLayoutDimension_IsFractionalHeight(void* ptr);
bool C_NSCollectionLayoutDimension_IsFractionalWidth(void* ptr);
