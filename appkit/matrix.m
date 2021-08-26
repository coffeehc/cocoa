#import <Appkit/Appkit.h>
#import "matrix.h"

void* C_Matrix_Alloc() {
    return [NSMatrix alloc];
}

void* C_NSMatrix_InitWithFrame(void* ptr, CGRect frameRect) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMatrix* result_ = [nSMatrix initWithFrame:frameRect];
    return result_;
}

void* C_NSMatrix_InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(void* ptr, CGRect frameRect, unsigned int mode, void* cell, int rowsHigh, int colsWide) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMatrix* result_ = [nSMatrix initWithFrame:frameRect mode:mode prototype:(NSCell*)cell numberOfRows:rowsHigh numberOfColumns:colsWide];
    return result_;
}

void* C_NSMatrix_InitWithCoder(void* ptr, void* coder) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMatrix* result_ = [nSMatrix initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSMatrix_Init(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMatrix* result_ = [nSMatrix init];
    return result_;
}

void C_NSMatrix_AddColumn(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix addColumn];
}

void C_NSMatrix_AddColumnWithCells(void* ptr, Array newCells) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMutableArray* objcNewCells = [[NSMutableArray alloc] init];
    if (newCells.len > 0) {
    	void** newCellsData = (void**)newCells.data;
    	for (int i = 0; i < newCells.len; i++) {
    		void* p = newCellsData[i];
    		[objcNewCells addObject:(NSCell*)(NSCell*)p];
    	}
    }
    [nSMatrix addColumnWithCells:objcNewCells];
}

void C_NSMatrix_AddRow(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix addRow];
}

void C_NSMatrix_AddRowWithCells(void* ptr, Array newCells) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMutableArray* objcNewCells = [[NSMutableArray alloc] init];
    if (newCells.len > 0) {
    	void** newCellsData = (void**)newCells.data;
    	for (int i = 0; i < newCells.len; i++) {
    		void* p = newCellsData[i];
    		[objcNewCells addObject:(NSCell*)(NSCell*)p];
    	}
    }
    [nSMatrix addRowWithCells:objcNewCells];
}

CGRect C_NSMatrix_CellFrameAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSRect result_ = [nSMatrix cellFrameAtRow:row column:col];
    return result_;
}

void C_NSMatrix_InsertColumn(void* ptr, int column) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix insertColumn:column];
}

void C_NSMatrix_InsertColumn_WithCells(void* ptr, int column, Array newCells) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMutableArray* objcNewCells = [[NSMutableArray alloc] init];
    if (newCells.len > 0) {
    	void** newCellsData = (void**)newCells.data;
    	for (int i = 0; i < newCells.len; i++) {
    		void* p = newCellsData[i];
    		[objcNewCells addObject:(NSCell*)(NSCell*)p];
    	}
    }
    [nSMatrix insertColumn:column withCells:objcNewCells];
}

void C_NSMatrix_InsertRow(void* ptr, int row) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix insertRow:row];
}

void C_NSMatrix_InsertRow_WithCells(void* ptr, int row, Array newCells) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMutableArray* objcNewCells = [[NSMutableArray alloc] init];
    if (newCells.len > 0) {
    	void** newCellsData = (void**)newCells.data;
    	for (int i = 0; i < newCells.len; i++) {
    		void* p = newCellsData[i];
    		[objcNewCells addObject:(NSCell*)(NSCell*)p];
    	}
    }
    [nSMatrix insertRow:row withCells:objcNewCells];
}

void* C_NSMatrix_MakeCellAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix makeCellAtRow:row column:col];
    return result_;
}

void C_NSMatrix_PutCell_AtRow_Column(void* ptr, void* newCell, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix putCell:(NSCell*)newCell atRow:row column:col];
}

void C_NSMatrix_RemoveColumn(void* ptr, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix removeColumn:col];
}

void C_NSMatrix_RemoveRow(void* ptr, int row) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix removeRow:row];
}

void C_NSMatrix_RenewRows_Columns(void* ptr, int newRows, int newCols) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix renewRows:newRows columns:newCols];
}

void C_NSMatrix_SortUsingSelector(void* ptr, void* comparator) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix sortUsingSelector:(SEL)comparator];
}

void C_NSMatrix_SetState_AtRow_Column(void* ptr, int value, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setState:value atRow:row column:col];
}

void C_NSMatrix_SetToolTip_ForCell(void* ptr, void* toolTipString, void* cell) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setToolTip:(NSString*)toolTipString forCell:(NSCell*)cell];
}

void* C_NSMatrix_ToolTipForCell(void* ptr, void* cell) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSString* result_ = [nSMatrix toolTipForCell:(NSCell*)cell];
    return result_;
}

void C_NSMatrix_SelectCellAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix selectCellAtRow:row column:col];
}

bool C_NSMatrix_SelectCellWithTag(void* ptr, int tag) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix selectCellWithTag:tag];
    return result_;
}

void C_NSMatrix_SelectAll(void* ptr, void* sender) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix selectAll:(id)sender];
}

void C_NSMatrix_SetSelectionFrom_To_Anchor_Highlight(void* ptr, int startPos, int endPos, int anchorPos, bool lit) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setSelectionFrom:startPos to:endPos anchor:anchorPos highlight:lit];
}

void C_NSMatrix_DeselectAllCells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix deselectAllCells];
}

void C_NSMatrix_DeselectSelectedCell(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix deselectSelectedCell];
}

void* C_NSMatrix_CellAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix cellAtRow:row column:col];
    return result_;
}

void* C_NSMatrix_CellWithTag(void* ptr, int tag) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix cellWithTag:tag];
    return result_;
}

void C_NSMatrix_SelectText(void* ptr, void* sender) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix selectText:(id)sender];
}

void* C_NSMatrix_SelectTextAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix selectTextAtRow:row column:col];
    return result_;
}

bool C_NSMatrix_TextShouldBeginEditing(void* ptr, void* textObject) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix textShouldBeginEditing:(NSText*)textObject];
    return result_;
}

void C_NSMatrix_TextDidBeginEditing(void* ptr, void* notification) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix textDidBeginEditing:(NSNotification*)notification];
}

void C_NSMatrix_TextDidChange(void* ptr, void* notification) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix textDidChange:(NSNotification*)notification];
}

bool C_NSMatrix_TextShouldEndEditing(void* ptr, void* textObject) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix textShouldEndEditing:(NSText*)textObject];
    return result_;
}

void C_NSMatrix_TextDidEndEditing(void* ptr, void* notification) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix textDidEndEditing:(NSNotification*)notification];
}

void C_NSMatrix_SetValidateSize(void* ptr, bool flag) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setValidateSize:flag];
}

void C_NSMatrix_SizeToCells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix sizeToCells];
}

void C_NSMatrix_SetScrollable(void* ptr, bool flag) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setScrollable:flag];
}

void C_NSMatrix_ScrollCellToVisibleAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix scrollCellToVisibleAtRow:row column:col];
}

void C_NSMatrix_DrawCellAtRow_Column(void* ptr, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix drawCellAtRow:row column:col];
}

void C_NSMatrix_HighlightCell_AtRow_Column(void* ptr, bool flag, int row, int col) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix highlightCell:flag atRow:row column:col];
}

bool C_NSMatrix_SendAction(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix sendAction];
    return result_;
}

void C_NSMatrix_SendAction_To_ForAllCells(void* ptr, void* selector, void* object, bool flag) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix sendAction:(SEL)selector to:(id)object forAllCells:flag];
}

void C_NSMatrix_SendDoubleAction(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix sendDoubleAction];
}

unsigned int C_NSMatrix_Mode(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSMatrixMode result_ = [nSMatrix mode];
    return result_;
}

void C_NSMatrix_SetMode(void* ptr, unsigned int value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setMode:value];
}

bool C_NSMatrix_AllowsEmptySelection(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix allowsEmptySelection];
    return result_;
}

void C_NSMatrix_SetAllowsEmptySelection(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setAllowsEmptySelection:value];
}

bool C_NSMatrix_IsSelectionByRect(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix isSelectionByRect];
    return result_;
}

void C_NSMatrix_SetSelectionByRect(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setSelectionByRect:value];
}

void* C_NSMatrix_Prototype(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix prototype];
    return result_;
}

void C_NSMatrix_SetPrototype(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setPrototype:(NSCell*)value];
}

CGSize C_NSMatrix_CellSize(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSSize result_ = [nSMatrix cellSize];
    return result_;
}

void C_NSMatrix_SetCellSize(void* ptr, CGSize value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setCellSize:value];
}

CGSize C_NSMatrix_IntercellSpacing(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSSize result_ = [nSMatrix intercellSpacing];
    return result_;
}

void C_NSMatrix_SetIntercellSpacing(void* ptr, CGSize value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setIntercellSpacing:value];
}

int C_NSMatrix_NumberOfColumns(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSInteger result_ = [nSMatrix numberOfColumns];
    return result_;
}

int C_NSMatrix_NumberOfRows(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSInteger result_ = [nSMatrix numberOfRows];
    return result_;
}

bool C_NSMatrix_AutorecalculatesCellSize(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix autorecalculatesCellSize];
    return result_;
}

void C_NSMatrix_SetAutorecalculatesCellSize(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setAutorecalculatesCellSize:value];
}

void* C_NSMatrix_KeyCell(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix keyCell];
    return result_;
}

void C_NSMatrix_SetKeyCell(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setKeyCell:(NSCell*)value];
}

void* C_NSMatrix_SelectedCell(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSCell* result_ = [nSMatrix selectedCell];
    return result_;
}

Array C_NSMatrix_SelectedCells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSArray* result_ = [nSMatrix selectedCells];
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

int C_NSMatrix_SelectedColumn(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSInteger result_ = [nSMatrix selectedColumn];
    return result_;
}

int C_NSMatrix_SelectedRow(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSInteger result_ = [nSMatrix selectedRow];
    return result_;
}

Array C_NSMatrix_Cells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSArray* result_ = [nSMatrix cells];
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

void* C_NSMatrix_BackgroundColor(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSColor* result_ = [nSMatrix backgroundColor];
    return result_;
}

void C_NSMatrix_SetBackgroundColor(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setBackgroundColor:(NSColor*)value];
}

void* C_NSMatrix_CellBackgroundColor(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSColor* result_ = [nSMatrix cellBackgroundColor];
    return result_;
}

void C_NSMatrix_SetCellBackgroundColor(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setCellBackgroundColor:(NSColor*)value];
}

bool C_NSMatrix_DrawsBackground(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix drawsBackground];
    return result_;
}

void C_NSMatrix_SetDrawsBackground(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setDrawsBackground:value];
}

bool C_NSMatrix_DrawsCellBackground(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix drawsCellBackground];
    return result_;
}

void C_NSMatrix_SetDrawsCellBackground(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setDrawsCellBackground:value];
}

bool C_NSMatrix_TabKeyTraversesCells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix tabKeyTraversesCells];
    return result_;
}

void C_NSMatrix_SetTabKeyTraversesCells(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setTabKeyTraversesCells:value];
}

void* C_NSMatrix_Delegate(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    id result_ = [nSMatrix delegate];
    return result_;
}

void C_NSMatrix_SetDelegate(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setDelegate:(id)value];
}

bool C_NSMatrix_AutosizesCells(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix autosizesCells];
    return result_;
}

void C_NSMatrix_SetAutosizesCells(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setAutosizesCells:value];
}

bool C_NSMatrix_IsAutoscroll(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    BOOL result_ = [nSMatrix isAutoscroll];
    return result_;
}

void C_NSMatrix_SetAutoscroll(void* ptr, bool value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setAutoscroll:value];
}

void* C_NSMatrix_DoubleAction(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    SEL result_ = [nSMatrix doubleAction];
    return result_;
}

void C_NSMatrix_SetDoubleAction(void* ptr, void* value) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    [nSMatrix setDoubleAction:(SEL)value];
}

int C_NSMatrix_MouseDownFlags(void* ptr) {
    NSMatrix* nSMatrix = (NSMatrix*)ptr;
    NSInteger result_ = [nSMatrix mouseDownFlags];
    return result_;
}
