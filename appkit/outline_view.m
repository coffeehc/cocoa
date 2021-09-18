#import "outline_view.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSOutlineView.h>

void* C_OutlineView_Alloc() {
    return [NSOutlineView alloc];
}

void* C_NSOutlineView_InitWithCoder(void* ptr, void* coder) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSOutlineView* result_ = [nSOutlineView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSOutlineView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSOutlineView* result_ = [nSOutlineView initWithFrame:frameRect];
    return result_;
}

void* C_NSOutlineView_Init(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSOutlineView* result_ = [nSOutlineView init];
    return result_;
}

void* C_NSOutlineView_AllocOutlineView() {
    NSOutlineView* result_ = [NSOutlineView alloc];
    return result_;
}

void* C_NSOutlineView_NewOutlineView() {
    NSOutlineView* result_ = [NSOutlineView new];
    return result_;
}

void* C_NSOutlineView_Autorelease(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSOutlineView* result_ = [nSOutlineView autorelease];
    return result_;
}

void* C_NSOutlineView_Retain(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSOutlineView* result_ = [nSOutlineView retain];
    return result_;
}

bool C_NSOutlineView_IsExpandable(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView isExpandable:(id)item];
    return result_;
}

bool C_NSOutlineView_IsItemExpanded(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView isItemExpanded:(id)item];
    return result_;
}

void C_NSOutlineView_ExpandItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView expandItem:(id)item];
}

void C_NSOutlineView_ExpandItem_ExpandChildren(void* ptr, void* item, bool expandChildren) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView expandItem:(id)item expandChildren:expandChildren];
}

void C_NSOutlineView_CollapseItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView collapseItem:(id)item];
}

void C_NSOutlineView_CollapseItem_CollapseChildren(void* ptr, void* item, bool collapseChildren) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView collapseItem:(id)item collapseChildren:collapseChildren];
}

void C_NSOutlineView_ReloadItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView reloadItem:(id)item];
}

void C_NSOutlineView_ReloadItem_ReloadChildren(void* ptr, void* item, bool reloadChildren) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView reloadItem:(id)item reloadChildren:reloadChildren];
}

void* C_NSOutlineView_ItemAtRow(void* ptr, int row) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    id result_ = [nSOutlineView itemAtRow:row];
    return result_;
}

int C_NSOutlineView_RowForItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSInteger result_ = [nSOutlineView rowForItem:(id)item];
    return result_;
}

int C_NSOutlineView_LevelForItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSInteger result_ = [nSOutlineView levelForItem:(id)item];
    return result_;
}

int C_NSOutlineView_LevelForRow(void* ptr, int row) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSInteger result_ = [nSOutlineView levelForRow:row];
    return result_;
}

void C_NSOutlineView_SetDropItem_DropChildIndex(void* ptr, void* item, int index) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setDropItem:(id)item dropChildIndex:index];
}

bool C_NSOutlineView_ShouldCollapseAutoExpandedItemsForDeposited(void* ptr, bool deposited) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView shouldCollapseAutoExpandedItemsForDeposited:deposited];
    return result_;
}

void* C_NSOutlineView_ParentForItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    id result_ = [nSOutlineView parentForItem:(id)item];
    return result_;
}

int C_NSOutlineView_ChildIndexForItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSInteger result_ = [nSOutlineView childIndexForItem:(id)item];
    return result_;
}

void* C_NSOutlineView_Child_OfItem(void* ptr, int index, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    id result_ = [nSOutlineView child:index ofItem:(id)item];
    return result_;
}

int C_NSOutlineView_NumberOfChildrenOfItem(void* ptr, void* item) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSInteger result_ = [nSOutlineView numberOfChildrenOfItem:(id)item];
    return result_;
}

CGRect C_NSOutlineView_FrameOfOutlineCellAtRow(void* ptr, int row) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSRect result_ = [nSOutlineView frameOfOutlineCellAtRow:row];
    return result_;
}

void C_NSOutlineView_InsertItemsAtIndexes_InParent_WithAnimation(void* ptr, void* indexes, void* parent, unsigned int animationOptions) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView insertItemsAtIndexes:(NSIndexSet*)indexes inParent:(id)parent withAnimation:animationOptions];
}

void C_NSOutlineView_MoveItemAtIndex_InParent_ToIndex_InParent(void* ptr, int fromIndex, void* oldParent, int toIndex, void* newParent) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView moveItemAtIndex:fromIndex inParent:(id)oldParent toIndex:toIndex inParent:(id)newParent];
}

void C_NSOutlineView_RemoveItemsAtIndexes_InParent_WithAnimation(void* ptr, void* indexes, void* parent, unsigned int animationOptions) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView removeItemsAtIndexes:(NSIndexSet*)indexes inParent:(id)parent withAnimation:animationOptions];
}

bool C_NSOutlineView_StronglyReferencesItems(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView stronglyReferencesItems];
    return result_;
}

void C_NSOutlineView_SetStronglyReferencesItems(void* ptr, bool value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setStronglyReferencesItems:value];
}

void* C_NSOutlineView_OutlineTableColumn(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    NSTableColumn* result_ = [nSOutlineView outlineTableColumn];
    return result_;
}

void C_NSOutlineView_SetOutlineTableColumn(void* ptr, void* value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setOutlineTableColumn:(NSTableColumn*)value];
}

bool C_NSOutlineView_AutoresizesOutlineColumn(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView autoresizesOutlineColumn];
    return result_;
}

void C_NSOutlineView_SetAutoresizesOutlineColumn(void* ptr, bool value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setAutoresizesOutlineColumn:value];
}

double C_NSOutlineView_IndentationPerLevel(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    CGFloat result_ = [nSOutlineView indentationPerLevel];
    return result_;
}

void C_NSOutlineView_SetIndentationPerLevel(void* ptr, double value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setIndentationPerLevel:value];
}

bool C_NSOutlineView_IndentationMarkerFollowsCell(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView indentationMarkerFollowsCell];
    return result_;
}

void C_NSOutlineView_SetIndentationMarkerFollowsCell(void* ptr, bool value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setIndentationMarkerFollowsCell:value];
}

bool C_NSOutlineView_AutosaveExpandedItems(void* ptr) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    BOOL result_ = [nSOutlineView autosaveExpandedItems];
    return result_;
}

void C_NSOutlineView_SetAutosaveExpandedItems(void* ptr, bool value) {
    NSOutlineView* nSOutlineView = (NSOutlineView*)ptr;
    [nSOutlineView setAutosaveExpandedItems:value];
}
