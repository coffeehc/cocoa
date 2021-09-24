#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionView_Alloc();

void* C_NSCollectionView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSCollectionView_InitWithCoder(void* ptr, void* coder);
void* C_NSCollectionView_Init(void* ptr);
void* C_NSCollectionView_AllocCollectionView();
void* C_NSCollectionView_NewCollectionView();
void* C_NSCollectionView_Autorelease(void* ptr);
void* C_NSCollectionView_Retain(void* ptr);
void* C_NSCollectionView_MakeItemWithIdentifier_ForIndexPath(void* ptr, void* identifier, void* indexPath);
void C_NSCollectionView_RegisterNib_ForItemWithIdentifier(void* ptr, void* nib, void* identifier);
void* C_NSCollectionView_MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(void* ptr, void* elementKind, void* identifier, void* indexPath);
void C_NSCollectionView_RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(void* ptr, void* nib, void* kind, void* identifier);
void C_NSCollectionView_ReloadData(void* ptr);
void C_NSCollectionView_ReloadSections(void* ptr, void* sections);
void C_NSCollectionView_ReloadItemsAtIndexPaths(void* ptr, void* indexPaths);
int C_NSCollectionView_NumberOfItemsInSection(void* ptr, int section);
void C_NSCollectionView_InsertItemsAtIndexPaths(void* ptr, void* indexPaths);
void C_NSCollectionView_MoveItemAtIndexPath_ToIndexPath(void* ptr, void* indexPath, void* newIndexPath);
void C_NSCollectionView_DeleteItemsAtIndexPaths(void* ptr, void* indexPaths);
void C_NSCollectionView_InsertSections(void* ptr, void* sections);
void C_NSCollectionView_MoveSection_ToSection(void* ptr, int section, int newSection);
void C_NSCollectionView_DeleteSections(void* ptr, void* sections);
void C_NSCollectionView_ToggleSectionCollapse(void* ptr, void* sender);
void C_NSCollectionView_SelectAll(void* ptr, void* sender);
void C_NSCollectionView_DeselectAll(void* ptr, void* sender);
void C_NSCollectionView_SelectItemsAtIndexPaths_ScrollPosition(void* ptr, void* indexPaths, unsigned int scrollPosition);
void C_NSCollectionView_DeselectItemsAtIndexPaths(void* ptr, void* indexPaths);
Array C_NSCollectionView_VisibleItems(void* ptr);
void* C_NSCollectionView_IndexPathsForVisibleItems(void* ptr);
void* C_NSCollectionView_IndexPathsForVisibleSupplementaryElementsOfKind(void* ptr, void* elementKind);
void* C_NSCollectionView_IndexPathForItem(void* ptr, void* item);
void* C_NSCollectionView_IndexPathForItemAtPoint(void* ptr, CGPoint point);
void* C_NSCollectionView_ItemAtIndexPath(void* ptr, void* indexPath);
void* C_NSCollectionView_SupplementaryViewForElementKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath);
void C_NSCollectionView_ScrollToItemsAtIndexPaths_ScrollPosition(void* ptr, void* indexPaths, unsigned int scrollPosition);
void* C_NSCollectionView_LayoutAttributesForItemAtIndexPath(void* ptr, void* indexPath);
void* C_NSCollectionView_LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(void* ptr, void* kind, void* indexPath);
void* C_NSCollectionView_ItemAtIndex(void* ptr, unsigned int index);
CGRect C_NSCollectionView_FrameForItemAtIndex(void* ptr, unsigned int index);
CGRect C_NSCollectionView_FrameForItemAtIndex_WithNumberOfItems(void* ptr, unsigned int index, unsigned int numberOfItems);
void C_NSCollectionView_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int dragOperationMask, bool localDestination);
void* C_NSCollectionView_DataSource(void* ptr);
void C_NSCollectionView_SetDataSource(void* ptr, void* value);
void* C_NSCollectionView_Delegate(void* ptr);
void C_NSCollectionView_SetDelegate(void* ptr, void* value);
Array C_NSCollectionView_Content(void* ptr);
void C_NSCollectionView_SetContent(void* ptr, Array value);
void* C_NSCollectionView_BackgroundView(void* ptr);
void C_NSCollectionView_SetBackgroundView(void* ptr, void* value);
Array C_NSCollectionView_BackgroundColors(void* ptr);
void C_NSCollectionView_SetBackgroundColors(void* ptr, Array value);
bool C_NSCollectionView_BackgroundViewScrollsWithContent(void* ptr);
void C_NSCollectionView_SetBackgroundViewScrollsWithContent(void* ptr, bool value);
void* C_NSCollectionView_CollectionViewLayout(void* ptr);
void C_NSCollectionView_SetCollectionViewLayout(void* ptr, void* value);
void* C_NSCollectionView_PrefetchDataSource(void* ptr);
void C_NSCollectionView_SetPrefetchDataSource(void* ptr, void* value);
int C_NSCollectionView_NumberOfSections(void* ptr);
bool C_NSCollectionView_IsSelectable(void* ptr);
void C_NSCollectionView_SetSelectable(void* ptr, bool value);
bool C_NSCollectionView_AllowsMultipleSelection(void* ptr);
void C_NSCollectionView_SetAllowsMultipleSelection(void* ptr, bool value);
bool C_NSCollectionView_AllowsEmptySelection(void* ptr);
void C_NSCollectionView_SetAllowsEmptySelection(void* ptr, bool value);
void* C_NSCollectionView_SelectionIndexPaths(void* ptr);
void C_NSCollectionView_SetSelectionIndexPaths(void* ptr, void* value);
bool C_NSCollectionView_IsFirstResponder(void* ptr);
void* C_NSCollectionView_SelectionIndexes(void* ptr);
void C_NSCollectionView_SetSelectionIndexes(void* ptr, void* value);

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

void* WrapCollectionViewDataSource(uintptr_t goID);

void* WrapCollectionViewDelegate(uintptr_t goID);
