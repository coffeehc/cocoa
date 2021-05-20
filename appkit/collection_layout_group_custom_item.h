#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutGroupCustomItem_Alloc();

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(CGRect frame);
void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(CGRect frame, int zIndex);
CGRect C_NSCollectionLayoutGroupCustomItem_Frame(void* ptr);
int C_NSCollectionLayoutGroupCustomItem_ZIndex(void* ptr);
