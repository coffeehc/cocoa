#import "grid_view.h"
#import <AppKit/NSGridView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_GridView_Alloc() {
    return [NSGridView alloc];
}

void* C_NSGridView_GridViewWithNumberOfColumns_Rows(int columnCount, int rowCount) {
    NSGridView* result_ = [NSGridView gridViewWithNumberOfColumns:columnCount rows:rowCount];
    return result_;
}

void* C_NSGridView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result_ = [nSGridView initWithFrame:frameRect];
    return result_;
}

void* C_NSGridView_InitWithCoder(void* ptr, void* coder) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result_ = [nSGridView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSGridView_Init(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result_ = [nSGridView init];
    return result_;
}

void* C_NSGridView_AllocGridView() {
    NSGridView* result_ = [NSGridView alloc];
    return result_;
}

void* C_NSGridView_NewGridView() {
    NSGridView* result_ = [NSGridView new];
    return result_;
}

void* C_NSGridView_Autorelease(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result_ = [nSGridView autorelease];
    return result_;
}

void* C_NSGridView_Retain(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridView* result_ = [nSGridView retain];
    return result_;
}

int C_NSGridView_IndexOfColumn(void* ptr, void* column) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result_ = [nSGridView indexOfColumn:(NSGridColumn*)column];
    return result_;
}

void* C_NSGridView_RowAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridRow* result_ = [nSGridView rowAtIndex:index];
    return result_;
}

void* C_NSGridView_ColumnAtIndex(void* ptr, int index) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridColumn* result_ = [nSGridView columnAtIndex:index];
    return result_;
}

int C_NSGridView_IndexOfRow(void* ptr, void* row) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result_ = [nSGridView indexOfRow:(NSGridRow*)row];
    return result_;
}

void* C_NSGridView_AddRowWithViews(void* ptr, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [NSMutableArray arrayWithCapacity:views.len];
    if (views.len > 0) {
    	void** viewsData = (void**)views.data;
    	for (int i = 0; i < views.len; i++) {
    		void* p = viewsData[i];
    		[objcViews addObject:(NSView*)p];
    	}
    }
    NSGridRow* result_ = [nSGridView addRowWithViews:objcViews];
    return result_;
}

void* C_NSGridView_InsertRowAtIndex_WithViews(void* ptr, int index, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [NSMutableArray arrayWithCapacity:views.len];
    if (views.len > 0) {
    	void** viewsData = (void**)views.data;
    	for (int i = 0; i < views.len; i++) {
    		void* p = viewsData[i];
    		[objcViews addObject:(NSView*)p];
    	}
    }
    NSGridRow* result_ = [nSGridView insertRowAtIndex:index withViews:objcViews];
    return result_;
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
    NSMutableArray* objcViews = [NSMutableArray arrayWithCapacity:views.len];
    if (views.len > 0) {
    	void** viewsData = (void**)views.data;
    	for (int i = 0; i < views.len; i++) {
    		void* p = viewsData[i];
    		[objcViews addObject:(NSView*)p];
    	}
    }
    NSGridColumn* result_ = [nSGridView addColumnWithViews:objcViews];
    return result_;
}

void* C_NSGridView_InsertColumnAtIndex_WithViews(void* ptr, int index, Array views) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [NSMutableArray arrayWithCapacity:views.len];
    if (views.len > 0) {
    	void** viewsData = (void**)views.data;
    	for (int i = 0; i < views.len; i++) {
    		void* p = viewsData[i];
    		[objcViews addObject:(NSView*)p];
    	}
    }
    NSGridColumn* result_ = [nSGridView insertColumnAtIndex:index withViews:objcViews];
    return result_;
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
    NSGridCell* result_ = [nSGridView cellAtColumnIndex:columnIndex rowIndex:rowIndex];
    return result_;
}

void* C_NSGridView_CellForView(void* ptr, void* view) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCell* result_ = [nSGridView cellForView:(NSView*)view];
    return result_;
}

void C_NSGridView_MergeCellsInHorizontalRange_VerticalRange(void* ptr, NSRange hRange, NSRange vRange) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView mergeCellsInHorizontalRange:hRange verticalRange:vRange];
}

int C_NSGridView_NumberOfRows(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result_ = [nSGridView numberOfRows];
    return result_;
}

int C_NSGridView_NumberOfColumns(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSInteger result_ = [nSGridView numberOfColumns];
    return result_;
}

double C_NSGridView_ColumnSpacing(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    CGFloat result_ = [nSGridView columnSpacing];
    return result_;
}

void C_NSGridView_SetColumnSpacing(void* ptr, double value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setColumnSpacing:value];
}

double C_NSGridView_RowSpacing(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    CGFloat result_ = [nSGridView rowSpacing];
    return result_;
}

void C_NSGridView_SetRowSpacing(void* ptr, double value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setRowSpacing:value];
}

int C_NSGridView_RowAlignment(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridRowAlignment result_ = [nSGridView rowAlignment];
    return result_;
}

void C_NSGridView_SetRowAlignment(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setRowAlignment:value];
}

int C_NSGridView_XPlacement(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCellPlacement result_ = [nSGridView xPlacement];
    return result_;
}

void C_NSGridView_SetXPlacement(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setXPlacement:value];
}

int C_NSGridView_YPlacement(void* ptr) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    NSGridCellPlacement result_ = [nSGridView yPlacement];
    return result_;
}

void C_NSGridView_SetYPlacement(void* ptr, int value) {
    NSGridView* nSGridView = (NSGridView*)ptr;
    [nSGridView setYPlacement:value];
}

void* C_GridCell_Alloc() {
    return [NSGridCell alloc];
}

void* C_NSGridCell_AllocGridCell() {
    NSGridCell* result_ = [NSGridCell alloc];
    return result_;
}

void* C_NSGridCell_Init(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCell* result_ = [nSGridCell init];
    return result_;
}

void* C_NSGridCell_NewGridCell() {
    NSGridCell* result_ = [NSGridCell new];
    return result_;
}

void* C_NSGridCell_Autorelease(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCell* result_ = [nSGridCell autorelease];
    return result_;
}

void* C_NSGridCell_Retain(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCell* result_ = [nSGridCell retain];
    return result_;
}

void* C_NSGridCell_Column(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridColumn* result_ = [nSGridCell column];
    return result_;
}

void* C_NSGridCell_Row(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRow* result_ = [nSGridCell row];
    return result_;
}

void* C_NSGridCell_ContentView(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSView* result_ = [nSGridCell contentView];
    return result_;
}

void C_NSGridCell_SetContentView(void* ptr, void* value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setContentView:(NSView*)value];
}

void* C_NSGridCell_GridCell_EmptyContentView() {
    NSView* result_ = [NSGridCell emptyContentView];
    return result_;
}

Array C_NSGridCell_CustomPlacementConstraints(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSArray* result_ = [nSGridCell customPlacementConstraints];
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

void C_NSGridCell_SetCustomPlacementConstraints(void* ptr, Array value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSLayoutConstraint*)p];
    	}
    }
    [nSGridCell setCustomPlacementConstraints:objcValue];
}

int C_NSGridCell_RowAlignment(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRowAlignment result_ = [nSGridCell rowAlignment];
    return result_;
}

void C_NSGridCell_SetRowAlignment(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setRowAlignment:value];
}

int C_NSGridCell_XPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result_ = [nSGridCell xPlacement];
    return result_;
}

void C_NSGridCell_SetXPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setXPlacement:value];
}

int C_NSGridCell_YPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result_ = [nSGridCell yPlacement];
    return result_;
}

void C_NSGridCell_SetYPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setYPlacement:value];
}

void* C_GridColumn_Alloc() {
    return [NSGridColumn alloc];
}

void* C_NSGridColumn_AllocGridColumn() {
    NSGridColumn* result_ = [NSGridColumn alloc];
    return result_;
}

void* C_NSGridColumn_Init(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn init];
    return result_;
}

void* C_NSGridColumn_NewGridColumn() {
    NSGridColumn* result_ = [NSGridColumn new];
    return result_;
}

void* C_NSGridColumn_Autorelease(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn autorelease];
    return result_;
}

void* C_NSGridColumn_Retain(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn retain];
    return result_;
}

void* C_NSGridColumn_CellAtIndex(void* ptr, int index) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCell* result_ = [nSGridColumn cellAtIndex:index];
    return result_;
}

void C_NSGridColumn_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn mergeCellsInRange:_range];
}

void* C_NSGridColumn_GridView(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridView* result_ = [nSGridColumn gridView];
    return result_;
}

bool C_NSGridColumn_IsHidden(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    BOOL result_ = [nSGridColumn isHidden];
    return result_;
}

void C_NSGridColumn_SetHidden(void* ptr, bool value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setHidden:value];
}

double C_NSGridColumn_LeadingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn leadingPadding];
    return result_;
}

void C_NSGridColumn_SetLeadingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setLeadingPadding:value];
}

int C_NSGridColumn_NumberOfCells(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSInteger result_ = [nSGridColumn numberOfCells];
    return result_;
}

double C_NSGridColumn_TrailingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn trailingPadding];
    return result_;
}

void C_NSGridColumn_SetTrailingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setTrailingPadding:value];
}

double C_NSGridColumn_Width(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn width];
    return result_;
}

void C_NSGridColumn_SetWidth(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setWidth:value];
}

int C_NSGridColumn_XPlacement(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCellPlacement result_ = [nSGridColumn xPlacement];
    return result_;
}

void C_NSGridColumn_SetXPlacement(void* ptr, int value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setXPlacement:value];
}

void* C_GridRow_Alloc() {
    return [NSGridRow alloc];
}

void* C_NSGridRow_AllocGridRow() {
    NSGridRow* result_ = [NSGridRow alloc];
    return result_;
}

void* C_NSGridRow_Init(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow init];
    return result_;
}

void* C_NSGridRow_NewGridRow() {
    NSGridRow* result_ = [NSGridRow new];
    return result_;
}

void* C_NSGridRow_Autorelease(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow autorelease];
    return result_;
}

void* C_NSGridRow_Retain(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow retain];
    return result_;
}

void* C_NSGridRow_CellAtIndex(void* ptr, int index) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCell* result_ = [nSGridRow cellAtIndex:index];
    return result_;
}

void C_NSGridRow_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow mergeCellsInRange:_range];
}

int C_NSGridRow_NumberOfCells(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSInteger result_ = [nSGridRow numberOfCells];
    return result_;
}

bool C_NSGridRow_IsHidden(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    BOOL result_ = [nSGridRow isHidden];
    return result_;
}

void C_NSGridRow_SetHidden(void* ptr, bool value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHidden:value];
}

double C_NSGridRow_TopPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow topPadding];
    return result_;
}

void C_NSGridRow_SetTopPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setTopPadding:value];
}

double C_NSGridRow_BottomPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow bottomPadding];
    return result_;
}

void C_NSGridRow_SetBottomPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setBottomPadding:value];
}

double C_NSGridRow_Height(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow height];
    return result_;
}

void C_NSGridRow_SetHeight(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHeight:value];
}

int C_NSGridRow_RowAlignment(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRowAlignment result_ = [nSGridRow rowAlignment];
    return result_;
}

void C_NSGridRow_SetRowAlignment(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setRowAlignment:value];
}

int C_NSGridRow_YPlacement(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCellPlacement result_ = [nSGridRow yPlacement];
    return result_;
}

void C_NSGridRow_SetYPlacement(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setYPlacement:value];
}

void* C_NSGridRow_GridView(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridView* result_ = [nSGridRow gridView];
    return result_;
}
