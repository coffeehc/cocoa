#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Matrix_Alloc();

void* C_NSMatrix_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSMatrix_InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(void* ptr, CGRect frameRect, unsigned int mode, void* cell, int rowsHigh, int colsWide);
void* C_NSMatrix_InitWithCoder(void* ptr, void* coder);
void* C_NSMatrix_Init(void* ptr);
void* C_NSMatrix_AllocMatrix();
void* C_NSMatrix_NewMatrix();
void* C_NSMatrix_Autorelease(void* ptr);
void* C_NSMatrix_Retain(void* ptr);
void C_NSMatrix_AddColumn(void* ptr);
void C_NSMatrix_AddColumnWithCells(void* ptr, Array newCells);
void C_NSMatrix_AddRow(void* ptr);
void C_NSMatrix_AddRowWithCells(void* ptr, Array newCells);
CGRect C_NSMatrix_CellFrameAtRow_Column(void* ptr, int row, int col);
void C_NSMatrix_InsertColumn(void* ptr, int column);
void C_NSMatrix_InsertColumn_WithCells(void* ptr, int column, Array newCells);
void C_NSMatrix_InsertRow(void* ptr, int row);
void C_NSMatrix_InsertRow_WithCells(void* ptr, int row, Array newCells);
void* C_NSMatrix_MakeCellAtRow_Column(void* ptr, int row, int col);
void C_NSMatrix_PutCell_AtRow_Column(void* ptr, void* newCell, int row, int col);
void C_NSMatrix_RemoveColumn(void* ptr, int col);
void C_NSMatrix_RemoveRow(void* ptr, int row);
void C_NSMatrix_RenewRows_Columns(void* ptr, int newRows, int newCols);
void C_NSMatrix_SortUsingSelector(void* ptr, void* comparator);
void C_NSMatrix_SetState_AtRow_Column(void* ptr, int value, int row, int col);
void C_NSMatrix_SetToolTip_ForCell(void* ptr, void* toolTipString, void* cell);
void* C_NSMatrix_ToolTipForCell(void* ptr, void* cell);
void C_NSMatrix_SelectCellAtRow_Column(void* ptr, int row, int col);
bool C_NSMatrix_SelectCellWithTag(void* ptr, int tag);
void C_NSMatrix_SelectAll(void* ptr, void* sender);
void C_NSMatrix_SetSelectionFrom_To_Anchor_Highlight(void* ptr, int startPos, int endPos, int anchorPos, bool lit);
void C_NSMatrix_DeselectAllCells(void* ptr);
void C_NSMatrix_DeselectSelectedCell(void* ptr);
void* C_NSMatrix_CellAtRow_Column(void* ptr, int row, int col);
void* C_NSMatrix_CellWithTag(void* ptr, int tag);
void C_NSMatrix_SelectText(void* ptr, void* sender);
void* C_NSMatrix_SelectTextAtRow_Column(void* ptr, int row, int col);
bool C_NSMatrix_TextShouldBeginEditing(void* ptr, void* textObject);
void C_NSMatrix_TextDidBeginEditing(void* ptr, void* notification);
void C_NSMatrix_TextDidChange(void* ptr, void* notification);
bool C_NSMatrix_TextShouldEndEditing(void* ptr, void* textObject);
void C_NSMatrix_TextDidEndEditing(void* ptr, void* notification);
void C_NSMatrix_SetValidateSize(void* ptr, bool flag);
void C_NSMatrix_SizeToCells(void* ptr);
void C_NSMatrix_SetScrollable(void* ptr, bool flag);
void C_NSMatrix_ScrollCellToVisibleAtRow_Column(void* ptr, int row, int col);
void C_NSMatrix_DrawCellAtRow_Column(void* ptr, int row, int col);
void C_NSMatrix_HighlightCell_AtRow_Column(void* ptr, bool flag, int row, int col);
bool C_NSMatrix_SendAction(void* ptr);
void C_NSMatrix_SendAction_To_ForAllCells(void* ptr, void* selector, void* object, bool flag);
void C_NSMatrix_SendDoubleAction(void* ptr);
unsigned int C_NSMatrix_Mode(void* ptr);
void C_NSMatrix_SetMode(void* ptr, unsigned int value);
bool C_NSMatrix_AllowsEmptySelection(void* ptr);
void C_NSMatrix_SetAllowsEmptySelection(void* ptr, bool value);
bool C_NSMatrix_IsSelectionByRect(void* ptr);
void C_NSMatrix_SetSelectionByRect(void* ptr, bool value);
void* C_NSMatrix_Prototype(void* ptr);
void C_NSMatrix_SetPrototype(void* ptr, void* value);
CGSize C_NSMatrix_CellSize(void* ptr);
void C_NSMatrix_SetCellSize(void* ptr, CGSize value);
CGSize C_NSMatrix_IntercellSpacing(void* ptr);
void C_NSMatrix_SetIntercellSpacing(void* ptr, CGSize value);
int C_NSMatrix_NumberOfColumns(void* ptr);
int C_NSMatrix_NumberOfRows(void* ptr);
bool C_NSMatrix_AutorecalculatesCellSize(void* ptr);
void C_NSMatrix_SetAutorecalculatesCellSize(void* ptr, bool value);
void* C_NSMatrix_KeyCell(void* ptr);
void C_NSMatrix_SetKeyCell(void* ptr, void* value);
void* C_NSMatrix_SelectedCell(void* ptr);
Array C_NSMatrix_SelectedCells(void* ptr);
int C_NSMatrix_SelectedColumn(void* ptr);
int C_NSMatrix_SelectedRow(void* ptr);
Array C_NSMatrix_Cells(void* ptr);
void* C_NSMatrix_BackgroundColor(void* ptr);
void C_NSMatrix_SetBackgroundColor(void* ptr, void* value);
void* C_NSMatrix_CellBackgroundColor(void* ptr);
void C_NSMatrix_SetCellBackgroundColor(void* ptr, void* value);
bool C_NSMatrix_DrawsBackground(void* ptr);
void C_NSMatrix_SetDrawsBackground(void* ptr, bool value);
bool C_NSMatrix_DrawsCellBackground(void* ptr);
void C_NSMatrix_SetDrawsCellBackground(void* ptr, bool value);
bool C_NSMatrix_TabKeyTraversesCells(void* ptr);
void C_NSMatrix_SetTabKeyTraversesCells(void* ptr, bool value);
void* C_NSMatrix_Delegate(void* ptr);
void C_NSMatrix_SetDelegate(void* ptr, void* value);
bool C_NSMatrix_AutosizesCells(void* ptr);
void C_NSMatrix_SetAutosizesCells(void* ptr, bool value);
bool C_NSMatrix_IsAutoscroll(void* ptr);
void C_NSMatrix_SetAutoscroll(void* ptr, bool value);
void* C_NSMatrix_DoubleAction(void* ptr);
void C_NSMatrix_SetDoubleAction(void* ptr, void* value);
int C_NSMatrix_MouseDownFlags(void* ptr);
