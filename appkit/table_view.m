#import <Appkit/Appkit.h>
#import "table_view.h"

void* C_TableView_Alloc() {
    return [NSTableView alloc];
}

void* C_NSTableView_InitWithCoder(void* ptr, void* coder) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableView* result_ = [nSTableView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableView* result_ = [nSTableView initWithFrame:frameRect];
    return result_;
}

void* C_NSTableView_Init(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableView* result_ = [nSTableView init];
    return result_;
}

void* C_NSTableView_AllocTableView() {
    NSTableView* result_ = [NSTableView alloc];
    return result_;
}

void* C_NSTableView_NewTableView() {
    NSTableView* result_ = [NSTableView new];
    return result_;
}

void* C_NSTableView_Autorelease(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableView* result_ = [nSTableView autorelease];
    return result_;
}

void* C_NSTableView_Retain(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableView* result_ = [nSTableView retain];
    return result_;
}

void C_NSTableView_ReloadData(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView reloadData];
}

void C_NSTableView_ReloadDataForRowIndexes_ColumnIndexes(void* ptr, void* rowIndexes, void* columnIndexes) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView reloadDataForRowIndexes:(NSIndexSet*)rowIndexes columnIndexes:(NSIndexSet*)columnIndexes];
}

void* C_NSTableView_MakeViewWithIdentifier_Owner(void* ptr, void* identifier, void* owner) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSView* result_ = [nSTableView makeViewWithIdentifier:(NSString*)identifier owner:(id)owner];
    return result_;
}

void* C_NSTableView_RowViewAtRow_MakeIfNecessary(void* ptr, int row, bool makeIfNecessary) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableRowView* result_ = [nSTableView rowViewAtRow:row makeIfNecessary:makeIfNecessary];
    return result_;
}

void* C_NSTableView_ViewAtColumn_Row_MakeIfNecessary(void* ptr, int column, int row, bool makeIfNecessary) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSView* result_ = [nSTableView viewAtColumn:column row:row makeIfNecessary:makeIfNecessary];
    return result_;
}

void C_NSTableView_BeginUpdates(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView beginUpdates];
}

void C_NSTableView_EndUpdates(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView endUpdates];
}

void C_NSTableView_MoveRowAtIndex_ToIndex(void* ptr, int oldIndex, int newIndex) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView moveRowAtIndex:oldIndex toIndex:newIndex];
}

void C_NSTableView_InsertRowsAtIndexes_WithAnimation(void* ptr, void* indexes, unsigned int animationOptions) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView insertRowsAtIndexes:(NSIndexSet*)indexes withAnimation:animationOptions];
}

void C_NSTableView_RemoveRowsAtIndexes_WithAnimation(void* ptr, void* indexes, unsigned int animationOptions) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView removeRowsAtIndexes:(NSIndexSet*)indexes withAnimation:animationOptions];
}

int C_NSTableView_RowForView(void* ptr, void* view) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView rowForView:(NSView*)view];
    return result_;
}

int C_NSTableView_ColumnForView(void* ptr, void* view) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView columnForView:(NSView*)view];
    return result_;
}

void C_NSTableView_RegisterNib_ForIdentifier(void* ptr, void* nib, void* identifier) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView registerNib:(NSNib*)nib forIdentifier:(NSString*)identifier];
}

void* C_NSTableView_IndicatorImageInTableColumn(void* ptr, void* tableColumn) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSImage* result_ = [nSTableView indicatorImageInTableColumn:(NSTableColumn*)tableColumn];
    return result_;
}

void C_NSTableView_SetIndicatorImage_InTableColumn(void* ptr, void* image, void* tableColumn) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setIndicatorImage:(NSImage*)image inTableColumn:(NSTableColumn*)tableColumn];
}

void C_NSTableView_AddTableColumn(void* ptr, void* tableColumn) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView addTableColumn:(NSTableColumn*)tableColumn];
}

void C_NSTableView_RemoveTableColumn(void* ptr, void* tableColumn) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView removeTableColumn:(NSTableColumn*)tableColumn];
}

void C_NSTableView_MoveColumn_ToColumn(void* ptr, int oldIndex, int newIndex) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView moveColumn:oldIndex toColumn:newIndex];
}

int C_NSTableView_ColumnWithIdentifier(void* ptr, void* identifier) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView columnWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSTableView_TableColumnWithIdentifier(void* ptr, void* identifier) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableColumn* result_ = [nSTableView tableColumnWithIdentifier:(NSString*)identifier];
    return result_;
}

void C_NSTableView_SelectColumnIndexes_ByExtendingSelection(void* ptr, void* indexes, bool extend) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView selectColumnIndexes:(NSIndexSet*)indexes byExtendingSelection:extend];
}

void C_NSTableView_DeselectColumn(void* ptr, int column) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView deselectColumn:column];
}

bool C_NSTableView_IsColumnSelected(void* ptr, int column) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView isColumnSelected:column];
    return result_;
}

void C_NSTableView_SelectRowIndexes_ByExtendingSelection(void* ptr, void* indexes, bool extend) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView selectRowIndexes:(NSIndexSet*)indexes byExtendingSelection:extend];
}

void C_NSTableView_DeselectRow(void* ptr, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView deselectRow:row];
}

bool C_NSTableView_IsRowSelected(void* ptr, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView isRowSelected:row];
    return result_;
}

void C_NSTableView_SelectAll(void* ptr, void* sender) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView selectAll:(id)sender];
}

void C_NSTableView_DeselectAll(void* ptr, void* sender) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView deselectAll:(id)sender];
}

void C_NSTableView_EditColumn_Row_WithEvent_Select(void* ptr, int column, int row, void* event, bool _select) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView editColumn:column row:row withEvent:(NSEvent*)event select:_select];
}

void C_NSTableView_DidAddRowView_ForRow(void* ptr, void* rowView, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView didAddRowView:(NSTableRowView*)rowView forRow:row];
}

void C_NSTableView_DidRemoveRowView_ForRow(void* ptr, void* rowView, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView didRemoveRowView:(NSTableRowView*)rowView forRow:row];
}

CGRect C_NSTableView_RectOfColumn(void* ptr, int column) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSRect result_ = [nSTableView rectOfColumn:column];
    return result_;
}

CGRect C_NSTableView_RectOfRow(void* ptr, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSRect result_ = [nSTableView rectOfRow:row];
    return result_;
}

NSRange C_NSTableView_RowsInRect(void* ptr, CGRect rect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSRange result_ = [nSTableView rowsInRect:rect];
    return result_;
}

void* C_NSTableView_ColumnIndexesInRect(void* ptr, CGRect rect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSIndexSet* result_ = [nSTableView columnIndexesInRect:rect];
    return result_;
}

int C_NSTableView_ColumnAtPoint(void* ptr, CGPoint point) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView columnAtPoint:point];
    return result_;
}

int C_NSTableView_RowAtPoint(void* ptr, CGPoint point) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView rowAtPoint:point];
    return result_;
}

CGRect C_NSTableView_FrameOfCellAtColumn_Row(void* ptr, int column, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSRect result_ = [nSTableView frameOfCellAtColumn:column row:row];
    return result_;
}

void C_NSTableView_SizeLastColumnToFit(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView sizeLastColumnToFit];
}

void C_NSTableView_NoteNumberOfRowsChanged(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView noteNumberOfRowsChanged];
}

void C_NSTableView_Tile(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView tile];
}

void C_NSTableView_NoteHeightOfRowsWithIndexesChanged(void* ptr, void* indexSet) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView noteHeightOfRowsWithIndexesChanged:(NSIndexSet*)indexSet];
}

void C_NSTableView_DrawRow_ClipRect(void* ptr, int row, CGRect clipRect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView drawRow:row clipRect:clipRect];
}

void C_NSTableView_DrawGridInClipRect(void* ptr, CGRect clipRect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView drawGridInClipRect:clipRect];
}

void C_NSTableView_HighlightSelectionInClipRect(void* ptr, CGRect clipRect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView highlightSelectionInClipRect:clipRect];
}

void C_NSTableView_DrawBackgroundInClipRect(void* ptr, CGRect clipRect) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView drawBackgroundInClipRect:clipRect];
}

void C_NSTableView_ScrollRowToVisible(void* ptr, int row) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView scrollRowToVisible:row];
}

void C_NSTableView_ScrollColumnToVisible(void* ptr, int column) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView scrollColumnToVisible:column];
}

bool C_NSTableView_CanDragRowsWithIndexes_AtPoint(void* ptr, void* rowIndexes, CGPoint mouseDownPoint) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView canDragRowsWithIndexes:(NSIndexSet*)rowIndexes atPoint:mouseDownPoint];
    return result_;
}

void C_NSTableView_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int mask, bool isLocal) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDraggingSourceOperationMask:mask forLocal:isLocal];
}

void C_NSTableView_SetDropRow_DropOperation(void* ptr, int row, unsigned int dropOperation) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDropRow:row dropOperation:dropOperation];
}

void C_NSTableView_HideRowsAtIndexes_WithAnimation(void* ptr, void* indexes, unsigned int rowAnimation) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView hideRowsAtIndexes:(NSIndexSet*)indexes withAnimation:rowAnimation];
}

void C_NSTableView_UnhideRowsAtIndexes_WithAnimation(void* ptr, void* indexes, unsigned int rowAnimation) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView unhideRowsAtIndexes:(NSIndexSet*)indexes withAnimation:rowAnimation];
}

void* C_NSTableView_DataSource(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    id result_ = [nSTableView dataSource];
    return result_;
}

void C_NSTableView_SetDataSource(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDataSource:(id)value];
}

bool C_NSTableView_UsesStaticContents(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView usesStaticContents];
    return result_;
}

void C_NSTableView_SetUsesStaticContents(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setUsesStaticContents:value];
}

Dictionary C_NSTableView_RegisteredNibsByIdentifier(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSDictionary* result_ = [nSTableView registeredNibsByIdentifier];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSUserInterfaceItemIdentifier kp = [result_Keys objectAtIndex:i];
    		NSNib* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSTableView_DoubleAction(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    SEL result_ = [nSTableView doubleAction];
    return result_;
}

void C_NSTableView_SetDoubleAction(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDoubleAction:(SEL)value];
}

int C_NSTableView_ClickedColumn(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView clickedColumn];
    return result_;
}

int C_NSTableView_ClickedRow(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView clickedRow];
    return result_;
}

bool C_NSTableView_AllowsColumnReordering(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsColumnReordering];
    return result_;
}

void C_NSTableView_SetAllowsColumnReordering(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsColumnReordering:value];
}

bool C_NSTableView_AllowsColumnResizing(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsColumnResizing];
    return result_;
}

void C_NSTableView_SetAllowsColumnResizing(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsColumnResizing:value];
}

bool C_NSTableView_AllowsMultipleSelection(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsMultipleSelection];
    return result_;
}

void C_NSTableView_SetAllowsMultipleSelection(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsMultipleSelection:value];
}

bool C_NSTableView_AllowsEmptySelection(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsEmptySelection];
    return result_;
}

void C_NSTableView_SetAllowsEmptySelection(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsEmptySelection:value];
}

bool C_NSTableView_AllowsColumnSelection(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsColumnSelection];
    return result_;
}

void C_NSTableView_SetAllowsColumnSelection(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsColumnSelection:value];
}

bool C_NSTableView_UsesAutomaticRowHeights(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView usesAutomaticRowHeights];
    return result_;
}

void C_NSTableView_SetUsesAutomaticRowHeights(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setUsesAutomaticRowHeights:value];
}

CGSize C_NSTableView_IntercellSpacing(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSSize result_ = [nSTableView intercellSpacing];
    return result_;
}

void C_NSTableView_SetIntercellSpacing(void* ptr, CGSize value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setIntercellSpacing:value];
}

double C_NSTableView_RowHeight(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    CGFloat result_ = [nSTableView rowHeight];
    return result_;
}

void C_NSTableView_SetRowHeight(void* ptr, double value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setRowHeight:value];
}

void* C_NSTableView_BackgroundColor(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSColor* result_ = [nSTableView backgroundColor];
    return result_;
}

void C_NSTableView_SetBackgroundColor(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setBackgroundColor:(NSColor*)value];
}

bool C_NSTableView_UsesAlternatingRowBackgroundColors(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView usesAlternatingRowBackgroundColors];
    return result_;
}

void C_NSTableView_SetUsesAlternatingRowBackgroundColors(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setUsesAlternatingRowBackgroundColors:value];
}

int C_NSTableView_Style(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewStyle result_ = [nSTableView style];
    return result_;
}

void C_NSTableView_SetStyle(void* ptr, int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setStyle:value];
}

int C_NSTableView_EffectiveStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewStyle result_ = [nSTableView effectiveStyle];
    return result_;
}

int C_NSTableView_SelectionHighlightStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewSelectionHighlightStyle result_ = [nSTableView selectionHighlightStyle];
    return result_;
}

void C_NSTableView_SetSelectionHighlightStyle(void* ptr, int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setSelectionHighlightStyle:value];
}

void* C_NSTableView_GridColor(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSColor* result_ = [nSTableView gridColor];
    return result_;
}

void C_NSTableView_SetGridColor(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setGridColor:(NSColor*)value];
}

unsigned int C_NSTableView_GridStyleMask(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewGridLineStyle result_ = [nSTableView gridStyleMask];
    return result_;
}

void C_NSTableView_SetGridStyleMask(void* ptr, unsigned int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setGridStyleMask:value];
}

int C_NSTableView_EffectiveRowSizeStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewRowSizeStyle result_ = [nSTableView effectiveRowSizeStyle];
    return result_;
}

int C_NSTableView_RowSizeStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewRowSizeStyle result_ = [nSTableView rowSizeStyle];
    return result_;
}

void C_NSTableView_SetRowSizeStyle(void* ptr, int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setRowSizeStyle:value];
}

Array C_NSTableView_TableColumns(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSArray* result_ = [nSTableView tableColumns];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

int C_NSTableView_SelectedColumn(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView selectedColumn];
    return result_;
}

void* C_NSTableView_SelectedColumnIndexes(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSIndexSet* result_ = [nSTableView selectedColumnIndexes];
    return result_;
}

int C_NSTableView_NumberOfSelectedColumns(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView numberOfSelectedColumns];
    return result_;
}

int C_NSTableView_SelectedRow(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView selectedRow];
    return result_;
}

void* C_NSTableView_SelectedRowIndexes(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSIndexSet* result_ = [nSTableView selectedRowIndexes];
    return result_;
}

int C_NSTableView_NumberOfSelectedRows(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView numberOfSelectedRows];
    return result_;
}

bool C_NSTableView_AllowsTypeSelect(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView allowsTypeSelect];
    return result_;
}

void C_NSTableView_SetAllowsTypeSelect(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAllowsTypeSelect:value];
}

int C_NSTableView_NumberOfColumns(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView numberOfColumns];
    return result_;
}

int C_NSTableView_NumberOfRows(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView numberOfRows];
    return result_;
}

bool C_NSTableView_FloatsGroupRows(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView floatsGroupRows];
    return result_;
}

void C_NSTableView_SetFloatsGroupRows(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setFloatsGroupRows:value];
}

int C_NSTableView_EditedColumn(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView editedColumn];
    return result_;
}

int C_NSTableView_EditedRow(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSInteger result_ = [nSTableView editedRow];
    return result_;
}

void* C_NSTableView_HeaderView(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableHeaderView* result_ = [nSTableView headerView];
    return result_;
}

void C_NSTableView_SetHeaderView(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setHeaderView:(NSTableHeaderView*)value];
}

void* C_NSTableView_CornerView(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSView* result_ = [nSTableView cornerView];
    return result_;
}

void C_NSTableView_SetCornerView(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setCornerView:(NSView*)value];
}

unsigned int C_NSTableView_ColumnAutoresizingStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewColumnAutoresizingStyle result_ = [nSTableView columnAutoresizingStyle];
    return result_;
}

void C_NSTableView_SetColumnAutoresizingStyle(void* ptr, unsigned int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setColumnAutoresizingStyle:value];
}

bool C_NSTableView_AutosaveTableColumns(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView autosaveTableColumns];
    return result_;
}

void C_NSTableView_SetAutosaveTableColumns(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAutosaveTableColumns:value];
}

void* C_NSTableView_AutosaveName(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewAutosaveName result_ = [nSTableView autosaveName];
    return result_;
}

void C_NSTableView_SetAutosaveName(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setAutosaveName:(NSString*)value];
}

void* C_NSTableView_Delegate(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    id result_ = [nSTableView delegate];
    return result_;
}

void C_NSTableView_SetDelegate(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDelegate:(id)value];
}

void* C_NSTableView_HighlightedTableColumn(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableColumn* result_ = [nSTableView highlightedTableColumn];
    return result_;
}

void C_NSTableView_SetHighlightedTableColumn(void* ptr, void* value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setHighlightedTableColumn:(NSTableColumn*)value];
}

bool C_NSTableView_VerticalMotionCanBeginDrag(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView verticalMotionCanBeginDrag];
    return result_;
}

void C_NSTableView_SetVerticalMotionCanBeginDrag(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setVerticalMotionCanBeginDrag:value];
}

int C_NSTableView_DraggingDestinationFeedbackStyle(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSTableViewDraggingDestinationFeedbackStyle result_ = [nSTableView draggingDestinationFeedbackStyle];
    return result_;
}

void C_NSTableView_SetDraggingDestinationFeedbackStyle(void* ptr, int value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setDraggingDestinationFeedbackStyle:value];
}

Array C_NSTableView_SortDescriptors(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSArray* result_ = [nSTableView sortDescriptors];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSTableView_SetSortDescriptors(void* ptr, Array value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSSortDescriptor*)p];
    	}
    }
    [nSTableView setSortDescriptors:objcValue];
}

bool C_NSTableView_RowActionsVisible(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    BOOL result_ = [nSTableView rowActionsVisible];
    return result_;
}

void C_NSTableView_SetRowActionsVisible(void* ptr, bool value) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    [nSTableView setRowActionsVisible:value];
}

void* C_NSTableView_HiddenRowIndexes(void* ptr) {
    NSTableView* nSTableView = (NSTableView*)ptr;
    NSIndexSet* result_ = [nSTableView hiddenRowIndexes];
    return result_;
}
