#import "outline_view_delegate.h"
#import "_cgo_export.h"
#import <AppKit/NSOutlineView.h>

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
