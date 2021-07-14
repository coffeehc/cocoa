#import <Appkit/Appkit.h>
#import "table_view_delegate.h"
#import "_cgo_export.h"

@implementation NSTableViewDelegateAdaptor


- (NSView*)tableView:(NSTableView*)tableView viewForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    void* result_ = tableViewDelegate_TableView_ViewForTableColumn_Row([self goID], tableView, tableColumn, row);
    return (NSView*)result_;
}

- (NSTableRowView*)tableView:(NSTableView*)tableView rowViewForRow:(NSInteger)row {
    void* result_ = tableViewDelegate_TableView_RowViewForRow([self goID], tableView, row);
    return (NSTableRowView*)result_;
}

- (void)tableView:(NSTableView*)tableView didAddRowView:(NSTableRowView*)rowView forRow:(NSInteger)row {
    tableViewDelegate_TableView_DidAddRowView_ForRow([self goID], tableView, rowView, row);
}

- (void)tableView:(NSTableView*)tableView didRemoveRowView:(NSTableRowView*)rowView forRow:(NSInteger)row {
    tableViewDelegate_TableView_DidRemoveRowView_ForRow([self goID], tableView, rowView, row);
}

- (BOOL)tableView:(NSTableView*)tableView isGroupRow:(NSInteger)row {
    bool result_ = tableViewDelegate_TableView_IsGroupRow([self goID], tableView, row);
    return result_;
}

- (void)tableView:(NSTableView*)tableView willDisplayCell:(id)cell forTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    tableViewDelegate_TableView_WillDisplayCell_ForTableColumn_Row([self goID], tableView, cell, tableColumn, row);
}

- (NSCell*)tableView:(NSTableView*)tableView dataCellForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    void* result_ = tableViewDelegate_TableView_DataCellForTableColumn_Row([self goID], tableView, tableColumn, row);
    return (NSCell*)result_;
}

- (BOOL)tableView:(NSTableView*)tableView shouldShowCellExpansionForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    bool result_ = tableViewDelegate_TableView_ShouldShowCellExpansionForTableColumn_Row([self goID], tableView, tableColumn, row);
    return result_;
}

- (BOOL)tableView:(NSTableView*)tableView shouldEditTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    bool result_ = tableViewDelegate_TableView_ShouldEditTableColumn_Row([self goID], tableView, tableColumn, row);
    return result_;
}

- (CGFloat)tableView:(NSTableView*)tableView heightOfRow:(NSInteger)row {
    double result_ = tableViewDelegate_TableView_HeightOfRow([self goID], tableView, row);
    return result_;
}

- (CGFloat)tableView:(NSTableView*)tableView sizeToFitWidthOfColumn:(NSInteger)column {
    double result_ = tableViewDelegate_TableView_SizeToFitWidthOfColumn([self goID], tableView, column);
    return result_;
}

- (BOOL)selectionShouldChangeInTableView:(NSTableView*)tableView {
    bool result_ = tableViewDelegate_SelectionShouldChangeInTableView([self goID], tableView);
    return result_;
}

- (BOOL)tableView:(NSTableView*)tableView shouldSelectRow:(NSInteger)row {
    bool result_ = tableViewDelegate_TableView_ShouldSelectRow([self goID], tableView, row);
    return result_;
}

- (NSIndexSet*)tableView:(NSTableView*)tableView selectionIndexesForProposedSelection:(NSIndexSet*)proposedSelectionIndexes {
    void* result_ = tableViewDelegate_TableView_SelectionIndexesForProposedSelection([self goID], tableView, proposedSelectionIndexes);
    return (NSIndexSet*)result_;
}

- (BOOL)tableView:(NSTableView*)tableView shouldSelectTableColumn:(NSTableColumn*)tableColumn {
    bool result_ = tableViewDelegate_TableView_ShouldSelectTableColumn([self goID], tableView, tableColumn);
    return result_;
}

- (void)tableViewSelectionIsChanging:(NSNotification*)notification {
    tableViewDelegate_TableViewSelectionIsChanging([self goID], notification);
}

- (void)tableViewSelectionDidChange:(NSNotification*)notification {
    tableViewDelegate_TableViewSelectionDidChange([self goID], notification);
}

- (BOOL)tableView:(NSTableView*)tableView shouldTypeSelectForEvent:(NSEvent*)event withCurrentSearchString:(NSString*)searchString {
    bool result_ = tableViewDelegate_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString([self goID], tableView, event, searchString);
    return result_;
}

- (NSString*)tableView:(NSTableView*)tableView typeSelectStringForTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    void* result_ = tableViewDelegate_TableView_TypeSelectStringForTableColumn_Row([self goID], tableView, tableColumn, row);
    return (NSString*)result_;
}

- (NSInteger)tableView:(NSTableView*)tableView nextTypeSelectMatchFromRow:(NSInteger)startRow toRow:(NSInteger)endRow forString:(NSString*)searchString {
    int result_ = tableViewDelegate_TableView_NextTypeSelectMatchFromRow_ToRow_ForString([self goID], tableView, startRow, endRow, searchString);
    return result_;
}

- (BOOL)tableView:(NSTableView*)tableView shouldReorderColumn:(NSInteger)columnIndex toColumn:(NSInteger)newColumnIndex {
    bool result_ = tableViewDelegate_TableView_ShouldReorderColumn_ToColumn([self goID], tableView, columnIndex, newColumnIndex);
    return result_;
}

- (void)tableView:(NSTableView*)tableView didDragTableColumn:(NSTableColumn*)tableColumn {
    tableViewDelegate_TableView_DidDragTableColumn([self goID], tableView, tableColumn);
}

- (void)tableViewColumnDidMove:(NSNotification*)notification {
    tableViewDelegate_TableViewColumnDidMove([self goID], notification);
}

- (void)tableViewColumnDidResize:(NSNotification*)notification {
    tableViewDelegate_TableViewColumnDidResize([self goID], notification);
}

- (void)tableView:(NSTableView*)tableView didClickTableColumn:(NSTableColumn*)tableColumn {
    tableViewDelegate_TableView_DidClickTableColumn([self goID], tableView, tableColumn);
}

- (void)tableView:(NSTableView*)tableView mouseDownInHeaderOfTableColumn:(NSTableColumn*)tableColumn {
    tableViewDelegate_TableView_MouseDownInHeaderOfTableColumn([self goID], tableView, tableColumn);
}

- (BOOL)tableView:(NSTableView*)tableView shouldTrackCell:(NSCell*)cell forTableColumn:(NSTableColumn*)tableColumn row:(NSInteger)row {
    bool result_ = tableViewDelegate_TableView_ShouldTrackCell_ForTableColumn_Row([self goID], tableView, cell, tableColumn, row);
    return result_;
}

- (NSArray*)tableView:(NSTableView*)tableView rowActionsForRow:(NSInteger)row edge:(NSTableRowActionEdge)edge {
    Array result_ = tableViewDelegate_TableView_RowActionsForRow_Edge([self goID], tableView, row, edge);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSTableViewRowAction*)(NSTableViewRowAction*)p];
    }
    return objcResult_;
}

- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = tableViewDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    tableViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = tableViewDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = tableViewDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = tableViewDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = tableViewDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    tableViewDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    tableViewDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    tableViewDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TableViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTableViewDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTableViewDelegate(uintptr_t goID) {
    NSTableViewDelegateAdaptor* adaptor = [[NSTableViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
