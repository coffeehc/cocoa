#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewGridLayout_Alloc();

void* C_NSCollectionViewGridLayout_AllocCollectionViewGridLayout();
void* C_NSCollectionViewGridLayout_Init(void* ptr);
void* C_NSCollectionViewGridLayout_NewCollectionViewGridLayout();
void* C_NSCollectionViewGridLayout_Autorelease(void* ptr);
void* C_NSCollectionViewGridLayout_Retain(void* ptr);
unsigned int C_NSCollectionViewGridLayout_MaximumNumberOfRows(void* ptr);
void C_NSCollectionViewGridLayout_SetMaximumNumberOfRows(void* ptr, unsigned int value);
unsigned int C_NSCollectionViewGridLayout_MaximumNumberOfColumns(void* ptr);
void C_NSCollectionViewGridLayout_SetMaximumNumberOfColumns(void* ptr, unsigned int value);
CGSize C_NSCollectionViewGridLayout_MinimumItemSize(void* ptr);
void C_NSCollectionViewGridLayout_SetMinimumItemSize(void* ptr, CGSize value);
CGSize C_NSCollectionViewGridLayout_MaximumItemSize(void* ptr);
void C_NSCollectionViewGridLayout_SetMaximumItemSize(void* ptr, CGSize value);
double C_NSCollectionViewGridLayout_MinimumInteritemSpacing(void* ptr);
void C_NSCollectionViewGridLayout_SetMinimumInteritemSpacing(void* ptr, double value);
double C_NSCollectionViewGridLayout_MinimumLineSpacing(void* ptr);
void C_NSCollectionViewGridLayout_SetMinimumLineSpacing(void* ptr, double value);
NSEdgeInsets C_NSCollectionViewGridLayout_Margins(void* ptr);
void C_NSCollectionViewGridLayout_SetMargins(void* ptr, NSEdgeInsets value);
Array C_NSCollectionViewGridLayout_BackgroundColors(void* ptr);
void C_NSCollectionViewGridLayout_SetBackgroundColors(void* ptr, Array value);
