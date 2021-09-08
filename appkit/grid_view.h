#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

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
