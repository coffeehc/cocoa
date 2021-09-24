#import "outline_view.h"
#import <AppKit/NSOutlineView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

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

@interface NSOutlineViewDataSourceAdaptor : NSObject <NSOutlineViewDataSource>
@property (assign) uintptr_t goID;
@end

@implementation NSOutlineViewDataSourceAdaptor


- (BOOL)outlineView:(NSOutlineView*)outlineView acceptDrop:(id)info item:(id)item childIndex:(NSInteger)index {
    bool result_ = outlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex([self goID], outlineView, info, item, index);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView child:(NSInteger)index ofItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_Child_OfItem([self goID], outlineView, index, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    outlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation([self goID], outlineView, session, screenPoint, operation);
}

- (void)outlineView:(NSOutlineView*)outlineView draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint forItems:(NSArray*)draggedItems {
    Array draggedItemsArray;
    int draggedItemscount = [draggedItems count];
    if (draggedItemscount > 0) {
    	void** draggedItemsData = malloc(draggedItemscount * sizeof(void*));
    	for (int i = 0; i < draggedItemscount; i++) {
    		 void* p = [draggedItems objectAtIndex:i];
    		 draggedItemsData[i] = p;
    	}
    	draggedItemsArray.data = draggedItemsData;
    	draggedItemsArray.len = draggedItemscount;
    }
    outlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems([self goID], outlineView, session, screenPoint, draggedItemsArray);
}

- (BOOL)outlineView:(NSOutlineView*)outlineView isItemExpandable:(id)item {
    bool result_ = outlineViewDataSource_OutlineView_IsItemExpandable([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView itemForPersistentObject:(id)object {
    void* result_ = outlineViewDataSource_OutlineView_ItemForPersistentObject([self goID], outlineView, object);
    return (id)result_;
}

- (NSInteger)outlineView:(NSOutlineView*)outlineView numberOfChildrenOfItem:(id)item {
    int result_ = outlineViewDataSource_OutlineView_NumberOfChildrenOfItem([self goID], outlineView, item);
    return result_;
}

- (id)outlineView:(NSOutlineView*)outlineView objectValueForTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem([self goID], outlineView, tableColumn, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView pasteboardWriterForItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_PasteboardWriterForItem([self goID], outlineView, item);
    return (id)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView persistentObjectForItem:(id)item {
    void* result_ = outlineViewDataSource_OutlineView_PersistentObjectForItem([self goID], outlineView, item);
    return (id)result_;
}

- (void)outlineView:(NSOutlineView*)outlineView setObjectValue:(id)object forTableColumn:(NSTableColumn*)tableColumn byItem:(id)item {
    outlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem([self goID], outlineView, object, tableColumn, item);
}

- (void)outlineView:(NSOutlineView*)outlineView sortDescriptorsDidChange:(NSArray*)oldDescriptors {
    Array oldDescriptorsArray;
    int oldDescriptorscount = [oldDescriptors count];
    if (oldDescriptorscount > 0) {
    	void** oldDescriptorsData = malloc(oldDescriptorscount * sizeof(void*));
    	for (int i = 0; i < oldDescriptorscount; i++) {
    		 void* p = [oldDescriptors objectAtIndex:i];
    		 oldDescriptorsData[i] = p;
    	}
    	oldDescriptorsArray.data = oldDescriptorsData;
    	oldDescriptorsArray.len = oldDescriptorscount;
    }
    outlineViewDataSource_OutlineView_SortDescriptorsDidChange([self goID], outlineView, oldDescriptorsArray);
}

- (void)outlineView:(NSOutlineView*)outlineView updateDraggingItemsForDrag:(id)draggingInfo {
    outlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag([self goID], outlineView, draggingInfo);
}

- (NSDragOperation)outlineView:(NSOutlineView*)outlineView validateDrop:(id)info proposedItem:(id)item proposedChildIndex:(NSInteger)index {
    unsigned int result_ = outlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex([self goID], outlineView, info, item, index);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return OutlineViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapOutlineViewDataSource(uintptr_t goID) {
    NSOutlineViewDataSourceAdaptor* adaptor = [[NSOutlineViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

@interface NSOutlineViewDelegateAdaptor : NSObject <NSOutlineViewDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSOutlineViewDelegateAdaptor


- (BOOL)outlineView:(NSOutlineView*)outlineView shouldExpandItem:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldExpandItem([self goID], outlineView, item);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldCollapseItem:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldCollapseItem([self goID], outlineView, item);
    return result_;
}

- (NSString*)outlineView:(NSOutlineView*)outlineView typeSelectStringForTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    void* result_ = outlineViewDelegate_OutlineView_TypeSelectStringForTableColumn_Item([self goID], outlineView, tableColumn, item);
    return (NSString*)result_;
}

- (id)outlineView:(NSOutlineView*)outlineView nextTypeSelectMatchFromItem:(id)startItem toItem:(id)endItem forString:(NSString*)searchString {
    void* result_ = outlineViewDelegate_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString([self goID], outlineView, startItem, endItem, searchString);
    return (id)result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldTypeSelectForEvent:(NSEvent*)event withCurrentSearchString:(NSString*)searchString {
    bool result_ = outlineViewDelegate_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString([self goID], outlineView, event, searchString);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldSelectTableColumn:(NSTableColumn*)tableColumn {
    bool result_ = outlineViewDelegate_OutlineView_ShouldSelectTableColumn([self goID], outlineView, tableColumn);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldSelectItem:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldSelectItem([self goID], outlineView, item);
    return result_;
}

- (NSIndexSet*)outlineView:(NSOutlineView*)outlineView selectionIndexesForProposedSelection:(NSIndexSet*)proposedSelectionIndexes {
    void* result_ = outlineViewDelegate_OutlineView_SelectionIndexesForProposedSelection([self goID], outlineView, proposedSelectionIndexes);
    return (NSIndexSet*)result_;
}

- (BOOL)selectionShouldChangeInOutlineView:(NSOutlineView*)outlineView {
    bool result_ = outlineViewDelegate_SelectionShouldChangeInOutlineView([self goID], outlineView);
    return result_;
}

- (void)outlineViewSelectionIsChanging:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewSelectionIsChanging([self goID], notification);
}

- (void)outlineViewSelectionDidChange:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewSelectionDidChange([self goID], notification);
}

- (void)outlineView:(NSOutlineView*)outlineView willDisplayCell:(id)cell forTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    outlineViewDelegate_OutlineView_WillDisplayCell_ForTableColumn_Item([self goID], outlineView, cell, tableColumn, item);
}

- (void)outlineView:(NSOutlineView*)outlineView willDisplayOutlineCell:(id)cell forTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    outlineViewDelegate_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item([self goID], outlineView, cell, tableColumn, item);
}

- (NSCell*)outlineView:(NSOutlineView*)outlineView dataCellForTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    void* result_ = outlineViewDelegate_OutlineView_DataCellForTableColumn_Item([self goID], outlineView, tableColumn, item);
    return (NSCell*)result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldShowOutlineCellForItem:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldShowOutlineCellForItem([self goID], outlineView, item);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldShowCellExpansionForTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldShowCellExpansionForTableColumn_Item([self goID], outlineView, tableColumn, item);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldReorderColumn:(NSInteger)columnIndex toColumn:(NSInteger)newColumnIndex {
    bool result_ = outlineViewDelegate_OutlineView_ShouldReorderColumn_ToColumn([self goID], outlineView, columnIndex, newColumnIndex);
    return result_;
}

- (void)outlineViewColumnDidMove:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewColumnDidMove([self goID], notification);
}

- (void)outlineViewColumnDidResize:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewColumnDidResize([self goID], notification);
}

- (void)outlineViewItemWillExpand:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewItemWillExpand([self goID], notification);
}

- (void)outlineViewItemDidExpand:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewItemDidExpand([self goID], notification);
}

- (void)outlineViewItemWillCollapse:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewItemWillCollapse([self goID], notification);
}

- (void)outlineViewItemDidCollapse:(NSNotification*)notification {
    outlineViewDelegate_OutlineViewItemDidCollapse([self goID], notification);
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldEditTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldEditTableColumn_Item([self goID], outlineView, tableColumn, item);
    return result_;
}

- (void)outlineView:(NSOutlineView*)outlineView mouseDownInHeaderOfTableColumn:(NSTableColumn*)tableColumn {
    outlineViewDelegate_OutlineView_MouseDownInHeaderOfTableColumn([self goID], outlineView, tableColumn);
}

- (void)outlineView:(NSOutlineView*)outlineView didClickTableColumn:(NSTableColumn*)tableColumn {
    outlineViewDelegate_OutlineView_DidClickTableColumn([self goID], outlineView, tableColumn);
}

- (void)outlineView:(NSOutlineView*)outlineView didDragTableColumn:(NSTableColumn*)tableColumn {
    outlineViewDelegate_OutlineView_DidDragTableColumn([self goID], outlineView, tableColumn);
}

- (CGFloat)outlineView:(NSOutlineView*)outlineView heightOfRowByItem:(id)item {
    double result_ = outlineViewDelegate_OutlineView_HeightOfRowByItem([self goID], outlineView, item);
    return result_;
}

- (CGFloat)outlineView:(NSOutlineView*)outlineView sizeToFitWidthOfColumn:(NSInteger)column {
    double result_ = outlineViewDelegate_OutlineView_SizeToFitWidthOfColumn([self goID], outlineView, column);
    return result_;
}

- (NSTintConfiguration*)outlineView:(NSOutlineView*)outlineView tintConfigurationForItem:(id)item {
    void* result_ = outlineViewDelegate_OutlineView_TintConfigurationForItem([self goID], outlineView, item);
    return (NSTintConfiguration*)result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView shouldTrackCell:(NSCell*)cell forTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_ShouldTrackCell_ForTableColumn_Item([self goID], outlineView, cell, tableColumn, item);
    return result_;
}

- (BOOL)outlineView:(NSOutlineView*)outlineView isGroupItem:(id)item {
    bool result_ = outlineViewDelegate_OutlineView_IsGroupItem([self goID], outlineView, item);
    return result_;
}

- (void)outlineView:(NSOutlineView*)outlineView didAddRowView:(NSTableRowView*)rowView forRow:(NSInteger)row {
    outlineViewDelegate_OutlineView_DidAddRowView_ForRow([self goID], outlineView, rowView, row);
}

- (void)outlineView:(NSOutlineView*)outlineView didRemoveRowView:(NSTableRowView*)rowView forRow:(NSInteger)row {
    outlineViewDelegate_OutlineView_DidRemoveRowView_ForRow([self goID], outlineView, rowView, row);
}

- (NSTableRowView*)outlineView:(NSOutlineView*)outlineView rowViewForItem:(id)item {
    void* result_ = outlineViewDelegate_OutlineView_RowViewForItem([self goID], outlineView, item);
    return (NSTableRowView*)result_;
}

- (NSView*)outlineView:(NSOutlineView*)outlineView viewForTableColumn:(NSTableColumn*)tableColumn item:(id)item {
    void* result_ = outlineViewDelegate_OutlineView_ViewForTableColumn_Item([self goID], outlineView, tableColumn, item);
    return (NSView*)result_;
}

- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = outlineViewDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    outlineViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = outlineViewDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = outlineViewDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = outlineViewDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = outlineViewDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    outlineViewDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    outlineViewDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    outlineViewDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return OutlineViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapOutlineViewDelegate(uintptr_t goID) {
    NSOutlineViewDelegateAdaptor* adaptor = [[NSOutlineViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
