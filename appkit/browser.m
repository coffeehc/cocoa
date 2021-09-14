#import <Appkit/Appkit.h>
#import "browser.h"

void* C_Browser_Alloc() {
    return [NSBrowser alloc];
}

void* C_NSBrowser_InitWithFrame(void* ptr, CGRect frameRect) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowser* result_ = [nSBrowser initWithFrame:frameRect];
    return result_;
}

void* C_NSBrowser_InitWithCoder(void* ptr, void* coder) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowser* result_ = [nSBrowser initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSBrowser_Init(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowser* result_ = [nSBrowser init];
    return result_;
}

void* C_NSBrowser_AllocBrowser() {
    NSBrowser* result_ = [NSBrowser alloc];
    return result_;
}

void* C_NSBrowser_NewBrowser() {
    NSBrowser* result_ = [NSBrowser new];
    return result_;
}

void* C_NSBrowser_Autorelease(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowser* result_ = [nSBrowser autorelease];
    return result_;
}

void* C_NSBrowser_Retain(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowser* result_ = [nSBrowser retain];
    return result_;
}

void C_NSBrowser_Tile(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser tile];
}

void* C_NSBrowser_SelectedRowIndexesInColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSIndexSet* result_ = [nSBrowser selectedRowIndexesInColumn:column];
    return result_;
}

void C_NSBrowser_SelectRowIndexes_InColumn(void* ptr, void* indexes, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser selectRowIndexes:(NSIndexSet*)indexes inColumn:column];
}

void* C_NSBrowser_SelectedCellInColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser selectedCellInColumn:column];
    return result_;
}

void C_NSBrowser_SelectAll(void* ptr, void* sender) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser selectAll:(id)sender];
}

int C_NSBrowser_SelectedRowInColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser selectedRowInColumn:column];
    return result_;
}

void C_NSBrowser_SelectRow_InColumn(void* ptr, int row, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser selectRow:row inColumn:column];
}

void* C_NSBrowser_LoadedCellAtRow_Column(void* ptr, int row, int col) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser loadedCellAtRow:row column:col];
    return result_;
}

void C_NSBrowser_EditItemAtIndexPath_WithEvent_Select(void* ptr, void* indexPath, void* event, bool _select) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser editItemAtIndexPath:(NSIndexPath*)indexPath withEvent:(NSEvent*)event select:_select];
}

void* C_NSBrowser_ItemAtIndexPath(void* ptr, void* indexPath) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser itemAtIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSBrowser_ItemAtRow_InColumn(void* ptr, int row, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser itemAtRow:row inColumn:column];
    return result_;
}

void* C_NSBrowser_IndexPathForColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSIndexPath* result_ = [nSBrowser indexPathForColumn:column];
    return result_;
}

bool C_NSBrowser_IsLeafItem(void* ptr, void* item) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser isLeafItem:(id)item];
    return result_;
}

void* C_NSBrowser_ParentForItemsInColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser parentForItemsInColumn:column];
    return result_;
}

void* C_NSBrowser_Path(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSString* result_ = [nSBrowser path];
    return result_;
}

bool C_NSBrowser_SetPath(void* ptr, void* path) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser setPath:(NSString*)path];
    return result_;
}

void* C_NSBrowser_PathToColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSString* result_ = [nSBrowser pathToColumn:column];
    return result_;
}

void C_NSBrowser_AddColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser addColumn];
}

void C_NSBrowser_ValidateVisibleColumns(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser validateVisibleColumns];
}

void C_NSBrowser_LoadColumnZero(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser loadColumnZero];
}

void C_NSBrowser_ReloadColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser reloadColumn:column];
}

void* C_NSBrowser_TitleOfColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSString* result_ = [nSBrowser titleOfColumn:column];
    return result_;
}

void C_NSBrowser_SetTitle_OfColumn(void* ptr, void* _string, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setTitle:(NSString*)_string ofColumn:column];
}

void C_NSBrowser_DrawTitleOfColumn_InRect(void* ptr, int column, CGRect rect) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser drawTitleOfColumn:column inRect:rect];
}

CGRect C_NSBrowser_TitleFrameOfColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSRect result_ = [nSBrowser titleFrameOfColumn:column];
    return result_;
}

void C_NSBrowser_NoteHeightOfRowsWithIndexesChanged_InColumn(void* ptr, void* indexSet, int columnIndex) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser noteHeightOfRowsWithIndexesChanged:(NSIndexSet*)indexSet inColumn:columnIndex];
}

void C_NSBrowser_ReloadDataForRowIndexes_InColumn(void* ptr, void* rowIndexes, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser reloadDataForRowIndexes:(NSIndexSet*)rowIndexes inColumn:column];
}

void C_NSBrowser_ScrollColumnToVisible(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser scrollColumnToVisible:column];
}

void C_NSBrowser_ScrollColumnsLeftBy(void* ptr, int shiftAmount) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser scrollColumnsLeftBy:shiftAmount];
}

void C_NSBrowser_ScrollColumnsRightBy(void* ptr, int shiftAmount) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser scrollColumnsRightBy:shiftAmount];
}

void C_NSBrowser_ScrollRowToVisible_InColumn(void* ptr, int row, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser scrollRowToVisible:row inColumn:column];
}

void C_NSBrowser_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int mask, bool isLocal) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setDraggingSourceOperationMask:mask forLocal:isLocal];
}

bool C_NSBrowser_CanDragRowsWithIndexes_InColumn_WithEvent(void* ptr, void* rowIndexes, int column, void* event) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser canDragRowsWithIndexes:(NSIndexSet*)rowIndexes inColumn:column withEvent:(NSEvent*)event];
    return result_;
}

CGRect C_NSBrowser_FrameOfColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSRect result_ = [nSBrowser frameOfColumn:column];
    return result_;
}

CGRect C_NSBrowser_FrameOfInsideOfColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSRect result_ = [nSBrowser frameOfInsideOfColumn:column];
    return result_;
}

CGRect C_NSBrowser_FrameOfRow_InColumn(void* ptr, int row, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSRect result_ = [nSBrowser frameOfRow:row inColumn:column];
    return result_;
}

bool C_NSBrowser_SendAction(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser sendAction];
    return result_;
}

void C_NSBrowser_DoClick(void* ptr, void* sender) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser doClick:(id)sender];
}

void C_NSBrowser_DoDoubleClick(void* ptr, void* sender) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser doDoubleClick:(id)sender];
}

void C_NSBrowser_Browser_RemoveSavedColumnsWithAutosaveName(void* name) {
    [NSBrowser removeSavedColumnsWithAutosaveName:(NSString*)name];
}

double C_NSBrowser_ColumnContentWidthForColumnWidth(void* ptr, double columnWidth) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser columnContentWidthForColumnWidth:columnWidth];
    return result_;
}

double C_NSBrowser_ColumnWidthForColumnContentWidth(void* ptr, double columnContentWidth) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser columnWidthForColumnContentWidth:columnContentWidth];
    return result_;
}

double C_NSBrowser_WidthOfColumn(void* ptr, int column) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser widthOfColumn:column];
    return result_;
}

void C_NSBrowser_SetWidth_OfColumn(void* ptr, double columnWidth, int columnIndex) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setWidth:columnWidth ofColumn:columnIndex];
}

double C_NSBrowser_DefaultColumnWidth(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser defaultColumnWidth];
    return result_;
}

void C_NSBrowser_SetDefaultColumnWidth(void* ptr, double columnWidth) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setDefaultColumnWidth:columnWidth];
}

bool C_NSBrowser_ReusesColumns(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser reusesColumns];
    return result_;
}

void C_NSBrowser_SetReusesColumns(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setReusesColumns:value];
}

int C_NSBrowser_MaxVisibleColumns(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser maxVisibleColumns];
    return result_;
}

void C_NSBrowser_SetMaxVisibleColumns(void* ptr, int value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setMaxVisibleColumns:value];
}

bool C_NSBrowser_AutohidesScroller(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser autohidesScroller];
    return result_;
}

void C_NSBrowser_SetAutohidesScroller(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setAutohidesScroller:value];
}

void* C_NSBrowser_BackgroundColor(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSColor* result_ = [nSBrowser backgroundColor];
    return result_;
}

void C_NSBrowser_SetBackgroundColor(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setBackgroundColor:(NSColor*)value];
}

double C_NSBrowser_MinColumnWidth(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser minColumnWidth];
    return result_;
}

void C_NSBrowser_SetMinColumnWidth(void* ptr, double value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setMinColumnWidth:value];
}

bool C_NSBrowser_SeparatesColumns(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser separatesColumns];
    return result_;
}

void C_NSBrowser_SetSeparatesColumns(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setSeparatesColumns:value];
}

bool C_NSBrowser_TakesTitleFromPreviousColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser takesTitleFromPreviousColumn];
    return result_;
}

void C_NSBrowser_SetTakesTitleFromPreviousColumn(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setTakesTitleFromPreviousColumn:value];
}

void* C_NSBrowser_Delegate(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser delegate];
    return result_;
}

void C_NSBrowser_SetDelegate(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setDelegate:(id)value];
}

void* C_NSBrowser_CellPrototype(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser cellPrototype];
    return result_;
}

void C_NSBrowser_SetCellPrototype(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setCellPrototype:(id)value];
}

bool C_NSBrowser_AllowsBranchSelection(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser allowsBranchSelection];
    return result_;
}

void C_NSBrowser_SetAllowsBranchSelection(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setAllowsBranchSelection:value];
}

bool C_NSBrowser_AllowsEmptySelection(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser allowsEmptySelection];
    return result_;
}

void C_NSBrowser_SetAllowsEmptySelection(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setAllowsEmptySelection:value];
}

bool C_NSBrowser_AllowsMultipleSelection(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser allowsMultipleSelection];
    return result_;
}

void C_NSBrowser_SetAllowsMultipleSelection(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setAllowsMultipleSelection:value];
}

bool C_NSBrowser_AllowsTypeSelect(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser allowsTypeSelect];
    return result_;
}

void C_NSBrowser_SetAllowsTypeSelect(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setAllowsTypeSelect:value];
}

void* C_NSBrowser_SelectedCell(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    id result_ = [nSBrowser selectedCell];
    return result_;
}

Array C_NSBrowser_SelectedCells(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSArray* result_ = [nSBrowser selectedCells];
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

void* C_NSBrowser_SelectionIndexPath(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSIndexPath* result_ = [nSBrowser selectionIndexPath];
    return result_;
}

void C_NSBrowser_SetSelectionIndexPath(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setSelectionIndexPath:(NSIndexPath*)value];
}

Array C_NSBrowser_SelectionIndexPaths(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSArray* result_ = [nSBrowser selectionIndexPaths];
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

void C_NSBrowser_SetSelectionIndexPaths(void* ptr, Array value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSIndexPath*)p];
    	}
    }
    [nSBrowser setSelectionIndexPaths:objcValue];
}

void* C_NSBrowser_PathSeparator(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSString* result_ = [nSBrowser pathSeparator];
    return result_;
}

void C_NSBrowser_SetPathSeparator(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setPathSeparator:(NSString*)value];
}

int C_NSBrowser_SelectedColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser selectedColumn];
    return result_;
}

int C_NSBrowser_LastColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser lastColumn];
    return result_;
}

void C_NSBrowser_SetLastColumn(void* ptr, int value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setLastColumn:value];
}

int C_NSBrowser_FirstVisibleColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser firstVisibleColumn];
    return result_;
}

int C_NSBrowser_NumberOfVisibleColumns(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser numberOfVisibleColumns];
    return result_;
}

int C_NSBrowser_LastVisibleColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser lastVisibleColumn];
    return result_;
}

bool C_NSBrowser_IsLoaded(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser isLoaded];
    return result_;
}

bool C_NSBrowser_IsTitled(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser isTitled];
    return result_;
}

void C_NSBrowser_SetTitled(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setTitled:value];
}

double C_NSBrowser_TitleHeight(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser titleHeight];
    return result_;
}

bool C_NSBrowser_HasHorizontalScroller(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser hasHorizontalScroller];
    return result_;
}

void C_NSBrowser_SetHasHorizontalScroller(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setHasHorizontalScroller:value];
}

void* C_NSBrowser_DoubleAction(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    SEL result_ = [nSBrowser doubleAction];
    return result_;
}

void C_NSBrowser_SetDoubleAction(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setDoubleAction:(SEL)value];
}

bool C_NSBrowser_SendsActionOnArrowKeys(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser sendsActionOnArrowKeys];
    return result_;
}

void C_NSBrowser_SetSendsActionOnArrowKeys(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setSendsActionOnArrowKeys:value];
}

int C_NSBrowser_ClickedColumn(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser clickedColumn];
    return result_;
}

int C_NSBrowser_ClickedRow(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSInteger result_ = [nSBrowser clickedRow];
    return result_;
}

void* C_NSBrowser_ColumnsAutosaveName(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowserColumnsAutosaveName result_ = [nSBrowser columnsAutosaveName];
    return result_;
}

void C_NSBrowser_SetColumnsAutosaveName(void* ptr, void* value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setColumnsAutosaveName:(NSString*)value];
}

unsigned int C_NSBrowser_ColumnResizingType(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    NSBrowserColumnResizingType result_ = [nSBrowser columnResizingType];
    return result_;
}

void C_NSBrowser_SetColumnResizingType(void* ptr, unsigned int value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setColumnResizingType:value];
}

bool C_NSBrowser_PrefersAllColumnUserResizing(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    BOOL result_ = [nSBrowser prefersAllColumnUserResizing];
    return result_;
}

void C_NSBrowser_SetPrefersAllColumnUserResizing(void* ptr, bool value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setPrefersAllColumnUserResizing:value];
}

double C_NSBrowser_RowHeight(void* ptr) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    CGFloat result_ = [nSBrowser rowHeight];
    return result_;
}

void C_NSBrowser_SetRowHeight(void* ptr, double value) {
    NSBrowser* nSBrowser = (NSBrowser*)ptr;
    [nSBrowser setRowHeight:value];
}
