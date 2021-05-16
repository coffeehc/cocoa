#import <Appkit/Appkit.h>
#import "grid_view.h"

void* C_GridView_Alloc() {
    return [NSGridView alloc];
}

void* C_NSGridView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result = [nSGridView initWithFrame:frameRect];
    return result;
}

void* C_NSGridView_InitWithCoder(void* ptr, void* coder) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result = [nSGridView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSGridView_Init(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result = [nSGridView init];
    return result;
}

void* C_NSGridView_GridViewWithNumberOfColumns_Rows(int columnCount, int rowCount) {
    NSGridView* result = [NSGridView gridViewWithNumberOfColumns:columnCount rows:rowCount];
    return result;
}

int C_NSGridView_IndexOfColumn(void* ptr, void* column) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result = [nSGridView indexOfColumn:(NSGridColumn*)column];
    return result;
}

void* C_NSGridView_RowAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridRow* result = [nSGridView rowAtIndex:index];
    return result;
}

void* C_NSGridView_ColumnAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridColumn* result = [nSGridView columnAtIndex:index];
    return result;
}

int C_NSGridView_IndexOfRow(void* ptr, void* row) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result = [nSGridView indexOfRow:(NSGridRow*)row];
    return result;
}

void* C_NSGridView_AddRowWithViews(void* ptr, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSGridRow* result = [nSGridView addRowWithViews:objcViews];
    return result;
}

void* C_NSGridView_InsertRowAtIndex_WithViews(void* ptr, int index, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSGridRow* result = [nSGridView insertRowAtIndex:index withViews:objcViews];
    return result;
}

void C_NSGridView_RemoveRowAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView removeRowAtIndex:index];
}

void C_NSGridView_MoveRowAtIndex_ToIndex(void* ptr, int fromIndex, int toIndex) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView moveRowAtIndex:fromIndex toIndex:toIndex];
}

void* C_NSGridView_AddColumnWithViews(void* ptr, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSGridColumn* result = [nSGridView addColumnWithViews:objcViews];
    return result;
}

void* C_NSGridView_InsertColumnAtIndex_WithViews(void* ptr, int index, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSGridColumn* result = [nSGridView insertColumnAtIndex:index withViews:objcViews];
    return result;
}

void C_NSGridView_RemoveColumnAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView removeColumnAtIndex:index];
}

void C_NSGridView_MoveColumnAtIndex_ToIndex(void* ptr, int fromIndex, int toIndex) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView moveColumnAtIndex:fromIndex toIndex:toIndex];
}

void* C_NSGridView_CellAtColumnIndex_RowIndex(void* ptr, int columnIndex, int rowIndex) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCell* result = [nSGridView cellAtColumnIndex:columnIndex rowIndex:rowIndex];
    return result;
}

void* C_NSGridView_CellForView(void* ptr, void* view) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCell* result = [nSGridView cellForView:(NSView*)view];
    return result;
}

void C_NSGridView_MergeCellsInHorizontalRange_VerticalRange(void* ptr, NSRange hRange, NSRange vRange) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView mergeCellsInHorizontalRange:hRange verticalRange:vRange];
}

int C_NSGridView_NumberOfRows(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result = [nSGridView numberOfRows];
    return result;
}

int C_NSGridView_NumberOfColumns(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result = [nSGridView numberOfColumns];
    return result;
}

double C_NSGridView_ColumnSpacing(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    CGFloat result = [nSGridView columnSpacing];
    return result;
}

void C_NSGridView_SetColumnSpacing(void* ptr, double value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setColumnSpacing:value];
}

double C_NSGridView_RowSpacing(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    CGFloat result = [nSGridView rowSpacing];
    return result;
}

void C_NSGridView_SetRowSpacing(void* ptr, double value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setRowSpacing:value];
}

int C_NSGridView_RowAlignment(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridRowAlignment result = [nSGridView rowAlignment];
    return result;
}

void C_NSGridView_SetRowAlignment(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setRowAlignment:value];
}

int C_NSGridView_XPlacement(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCellPlacement result = [nSGridView xPlacement];
    return result;
}

void C_NSGridView_SetXPlacement(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setXPlacement:value];
}

int C_NSGridView_YPlacement(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCellPlacement result = [nSGridView yPlacement];
    return result;
}

void C_NSGridView_SetYPlacement(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setYPlacement:value];
}
