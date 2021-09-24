#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_GridView_Alloc();

void* C_NSGridView_GridViewWithNumberOfColumns_Rows(int columnCount, int rowCount);
void* C_NSGridView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSGridView_InitWithCoder(void* ptr, void* coder);
void* C_NSGridView_Init(void* ptr);
void* C_NSGridView_AllocGridView();
void* C_NSGridView_NewGridView();
void* C_NSGridView_Autorelease(void* ptr);
void* C_NSGridView_Retain(void* ptr);
int C_NSGridView_IndexOfColumn(void* ptr, void* column);
void* C_NSGridView_RowAtIndex(void* ptr, int index);
void* C_NSGridView_ColumnAtIndex(void* ptr, int index);
int C_NSGridView_IndexOfRow(void* ptr, void* row);
void* C_NSGridView_AddRowWithViews(void* ptr, Array views);
void* C_NSGridView_InsertRowAtIndex_WithViews(void* ptr, int index, Array views);
void C_NSGridView_RemoveRowAtIndex(void* ptr, int index);
void C_NSGridView_MoveRowAtIndex_ToIndex(void* ptr, int fromIndex, int toIndex);
void* C_NSGridView_AddColumnWithViews(void* ptr, Array views);
void* C_NSGridView_InsertColumnAtIndex_WithViews(void* ptr, int index, Array views);
void C_NSGridView_RemoveColumnAtIndex(void* ptr, int index);
void C_NSGridView_MoveColumnAtIndex_ToIndex(void* ptr, int fromIndex, int toIndex);
void* C_NSGridView_CellAtColumnIndex_RowIndex(void* ptr, int columnIndex, int rowIndex);
void* C_NSGridView_CellForView(void* ptr, void* view);
void C_NSGridView_MergeCellsInHorizontalRange_VerticalRange(void* ptr, NSRange hRange, NSRange vRange);
int C_NSGridView_NumberOfRows(void* ptr);
int C_NSGridView_NumberOfColumns(void* ptr);
double C_NSGridView_ColumnSpacing(void* ptr);
void C_NSGridView_SetColumnSpacing(void* ptr, double value);
double C_NSGridView_RowSpacing(void* ptr);
void C_NSGridView_SetRowSpacing(void* ptr, double value);
int C_NSGridView_RowAlignment(void* ptr);
void C_NSGridView_SetRowAlignment(void* ptr, int value);
int C_NSGridView_XPlacement(void* ptr);
void C_NSGridView_SetXPlacement(void* ptr, int value);
int C_NSGridView_YPlacement(void* ptr);
void C_NSGridView_SetYPlacement(void* ptr, int value);

void* C_GridCell_Alloc();

void* C_NSGridCell_AllocGridCell();
void* C_NSGridCell_Init(void* ptr);
void* C_NSGridCell_NewGridCell();
void* C_NSGridCell_Autorelease(void* ptr);
void* C_NSGridCell_Retain(void* ptr);
void* C_NSGridCell_Column(void* ptr);
void* C_NSGridCell_Row(void* ptr);
void* C_NSGridCell_ContentView(void* ptr);
void C_NSGridCell_SetContentView(void* ptr, void* value);
void* C_NSGridCell_GridCell_EmptyContentView();
Array C_NSGridCell_CustomPlacementConstraints(void* ptr);
void C_NSGridCell_SetCustomPlacementConstraints(void* ptr, Array value);
int C_NSGridCell_RowAlignment(void* ptr);
void C_NSGridCell_SetRowAlignment(void* ptr, int value);
int C_NSGridCell_XPlacement(void* ptr);
void C_NSGridCell_SetXPlacement(void* ptr, int value);
int C_NSGridCell_YPlacement(void* ptr);
void C_NSGridCell_SetYPlacement(void* ptr, int value);

void* C_GridColumn_Alloc();

void* C_NSGridColumn_AllocGridColumn();
void* C_NSGridColumn_Init(void* ptr);
void* C_NSGridColumn_NewGridColumn();
void* C_NSGridColumn_Autorelease(void* ptr);
void* C_NSGridColumn_Retain(void* ptr);
void* C_NSGridColumn_CellAtIndex(void* ptr, int index);
void C_NSGridColumn_MergeCellsInRange(void* ptr, NSRange _range);
void* C_NSGridColumn_GridView(void* ptr);
bool C_NSGridColumn_IsHidden(void* ptr);
void C_NSGridColumn_SetHidden(void* ptr, bool value);
double C_NSGridColumn_LeadingPadding(void* ptr);
void C_NSGridColumn_SetLeadingPadding(void* ptr, double value);
int C_NSGridColumn_NumberOfCells(void* ptr);
double C_NSGridColumn_TrailingPadding(void* ptr);
void C_NSGridColumn_SetTrailingPadding(void* ptr, double value);
double C_NSGridColumn_Width(void* ptr);
void C_NSGridColumn_SetWidth(void* ptr, double value);
int C_NSGridColumn_XPlacement(void* ptr);
void C_NSGridColumn_SetXPlacement(void* ptr, int value);

void* C_GridRow_Alloc();

void* C_NSGridRow_AllocGridRow();
void* C_NSGridRow_Init(void* ptr);
void* C_NSGridRow_NewGridRow();
void* C_NSGridRow_Autorelease(void* ptr);
void* C_NSGridRow_Retain(void* ptr);
void* C_NSGridRow_CellAtIndex(void* ptr, int index);
void C_NSGridRow_MergeCellsInRange(void* ptr, NSRange _range);
int C_NSGridRow_NumberOfCells(void* ptr);
bool C_NSGridRow_IsHidden(void* ptr);
void C_NSGridRow_SetHidden(void* ptr, bool value);
double C_NSGridRow_TopPadding(void* ptr);
void C_NSGridRow_SetTopPadding(void* ptr, double value);
double C_NSGridRow_BottomPadding(void* ptr);
void C_NSGridRow_SetBottomPadding(void* ptr, double value);
double C_NSGridRow_Height(void* ptr);
void C_NSGridRow_SetHeight(void* ptr, double value);
int C_NSGridRow_RowAlignment(void* ptr);
void C_NSGridRow_SetRowAlignment(void* ptr, int value);
int C_NSGridRow_YPlacement(void* ptr);
void C_NSGridRow_SetYPlacement(void* ptr, int value);
void* C_NSGridRow_GridView(void* ptr);
