#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewCompositionalLayout_Alloc();

void* C_NSCollectionViewCompositionalLayout_InitWithSection(void* ptr, void* section);
void* C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(void* ptr, void* section, void* configuration);
void* C_NSCollectionViewCompositionalLayout_Configuration(void* ptr);
void C_NSCollectionViewCompositionalLayout_SetConfiguration(void* ptr, void* value);
