#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewUpdateItem_Alloc();

void* C_NSCollectionViewUpdateItem_Init(void* ptr);
void* C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(void* ptr);
void* C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(void* ptr);
int C_NSCollectionViewUpdateItem_UpdateAction(void* ptr);
