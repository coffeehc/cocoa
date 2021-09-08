#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewItem_Alloc();

void* C_NSCollectionViewItem_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil);
void* C_NSCollectionViewItem_InitWithCoder(void* ptr, void* coder);
void* C_NSCollectionViewItem_Init(void* ptr);
void* C_NSCollectionViewItem_AllocCollectionViewItem();
void* C_NSCollectionViewItem_NewCollectionViewItem();
void* C_NSCollectionViewItem_Autorelease(void* ptr);
void* C_NSCollectionViewItem_Retain(void* ptr);
bool C_NSCollectionViewItem_IsSelected(void* ptr);
void C_NSCollectionViewItem_SetSelected(void* ptr, bool value);
int C_NSCollectionViewItem_HighlightState(void* ptr);
void C_NSCollectionViewItem_SetHighlightState(void* ptr, int value);
void* C_NSCollectionViewItem_CollectionView(void* ptr);
Array C_NSCollectionViewItem_DraggingImageComponents(void* ptr);
