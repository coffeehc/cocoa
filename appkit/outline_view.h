#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_OutlineView_Alloc();

void* C_NSOutlineView_InitWithCoder(void* ptr, void* coder);
void* C_NSOutlineView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSOutlineView_Init(void* ptr);
void* C_NSOutlineView_AllocOutlineView();
void* C_NSOutlineView_NewOutlineView();
void* C_NSOutlineView_Autorelease(void* ptr);
void* C_NSOutlineView_Retain(void* ptr);
bool C_NSOutlineView_IsExpandable(void* ptr, void* item);
bool C_NSOutlineView_IsItemExpanded(void* ptr, void* item);
void C_NSOutlineView_ExpandItem(void* ptr, void* item);
void C_NSOutlineView_ExpandItem_ExpandChildren(void* ptr, void* item, bool expandChildren);
void C_NSOutlineView_CollapseItem(void* ptr, void* item);
void C_NSOutlineView_CollapseItem_CollapseChildren(void* ptr, void* item, bool collapseChildren);
void C_NSOutlineView_ReloadItem(void* ptr, void* item);
void C_NSOutlineView_ReloadItem_ReloadChildren(void* ptr, void* item, bool reloadChildren);
void* C_NSOutlineView_ItemAtRow(void* ptr, int row);
int C_NSOutlineView_RowForItem(void* ptr, void* item);
int C_NSOutlineView_LevelForItem(void* ptr, void* item);
int C_NSOutlineView_LevelForRow(void* ptr, int row);
void C_NSOutlineView_SetDropItem_DropChildIndex(void* ptr, void* item, int index);
bool C_NSOutlineView_ShouldCollapseAutoExpandedItemsForDeposited(void* ptr, bool deposited);
void* C_NSOutlineView_ParentForItem(void* ptr, void* item);
int C_NSOutlineView_ChildIndexForItem(void* ptr, void* item);
void* C_NSOutlineView_Child_OfItem(void* ptr, int index, void* item);
int C_NSOutlineView_NumberOfChildrenOfItem(void* ptr, void* item);
CGRect C_NSOutlineView_FrameOfOutlineCellAtRow(void* ptr, int row);
void C_NSOutlineView_InsertItemsAtIndexes_InParent_WithAnimation(void* ptr, void* indexes, void* parent, unsigned int animationOptions);
void C_NSOutlineView_MoveItemAtIndex_InParent_ToIndex_InParent(void* ptr, int fromIndex, void* oldParent, int toIndex, void* newParent);
void C_NSOutlineView_RemoveItemsAtIndexes_InParent_WithAnimation(void* ptr, void* indexes, void* parent, unsigned int animationOptions);
bool C_NSOutlineView_StronglyReferencesItems(void* ptr);
void C_NSOutlineView_SetStronglyReferencesItems(void* ptr, bool value);
void* C_NSOutlineView_OutlineTableColumn(void* ptr);
void C_NSOutlineView_SetOutlineTableColumn(void* ptr, void* value);
bool C_NSOutlineView_AutoresizesOutlineColumn(void* ptr);
void C_NSOutlineView_SetAutoresizesOutlineColumn(void* ptr, bool value);
double C_NSOutlineView_IndentationPerLevel(void* ptr);
void C_NSOutlineView_SetIndentationPerLevel(void* ptr, double value);
bool C_NSOutlineView_IndentationMarkerFollowsCell(void* ptr);
void C_NSOutlineView_SetIndentationMarkerFollowsCell(void* ptr, bool value);
bool C_NSOutlineView_AutosaveExpandedItems(void* ptr);
void C_NSOutlineView_SetAutosaveExpandedItems(void* ptr, bool value);
