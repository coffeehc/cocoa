#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Browser_Alloc();

void* C_NSBrowser_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSBrowser_InitWithCoder(void* ptr, void* coder);
void* C_NSBrowser_Init(void* ptr);
void* C_NSBrowser_AllocBrowser();
void* C_NSBrowser_NewBrowser();
void* C_NSBrowser_Autorelease(void* ptr);
void* C_NSBrowser_Retain(void* ptr);
void C_NSBrowser_Tile(void* ptr);
void* C_NSBrowser_SelectedRowIndexesInColumn(void* ptr, int column);
void C_NSBrowser_SelectRowIndexes_InColumn(void* ptr, void* indexes, int column);
void* C_NSBrowser_SelectedCellInColumn(void* ptr, int column);
void C_NSBrowser_SelectAll(void* ptr, void* sender);
int C_NSBrowser_SelectedRowInColumn(void* ptr, int column);
void C_NSBrowser_SelectRow_InColumn(void* ptr, int row, int column);
void* C_NSBrowser_LoadedCellAtRow_Column(void* ptr, int row, int col);
void C_NSBrowser_EditItemAtIndexPath_WithEvent_Select(void* ptr, void* indexPath, void* event, bool _select);
void* C_NSBrowser_ItemAtIndexPath(void* ptr, void* indexPath);
void* C_NSBrowser_ItemAtRow_InColumn(void* ptr, int row, int column);
void* C_NSBrowser_IndexPathForColumn(void* ptr, int column);
bool C_NSBrowser_IsLeafItem(void* ptr, void* item);
void* C_NSBrowser_ParentForItemsInColumn(void* ptr, int column);
void* C_NSBrowser_Path(void* ptr);
bool C_NSBrowser_SetPath(void* ptr, void* path);
void* C_NSBrowser_PathToColumn(void* ptr, int column);
void C_NSBrowser_AddColumn(void* ptr);
void C_NSBrowser_ValidateVisibleColumns(void* ptr);
void C_NSBrowser_LoadColumnZero(void* ptr);
void C_NSBrowser_ReloadColumn(void* ptr, int column);
void* C_NSBrowser_TitleOfColumn(void* ptr, int column);
void C_NSBrowser_SetTitle_OfColumn(void* ptr, void* _string, int column);
void C_NSBrowser_DrawTitleOfColumn_InRect(void* ptr, int column, CGRect rect);
CGRect C_NSBrowser_TitleFrameOfColumn(void* ptr, int column);
void C_NSBrowser_NoteHeightOfRowsWithIndexesChanged_InColumn(void* ptr, void* indexSet, int columnIndex);
void C_NSBrowser_ReloadDataForRowIndexes_InColumn(void* ptr, void* rowIndexes, int column);
void C_NSBrowser_ScrollColumnToVisible(void* ptr, int column);
void C_NSBrowser_ScrollColumnsLeftBy(void* ptr, int shiftAmount);
void C_NSBrowser_ScrollColumnsRightBy(void* ptr, int shiftAmount);
void C_NSBrowser_ScrollRowToVisible_InColumn(void* ptr, int row, int column);
void C_NSBrowser_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int mask, bool isLocal);
bool C_NSBrowser_CanDragRowsWithIndexes_InColumn_WithEvent(void* ptr, void* rowIndexes, int column, void* event);
CGRect C_NSBrowser_FrameOfColumn(void* ptr, int column);
CGRect C_NSBrowser_FrameOfInsideOfColumn(void* ptr, int column);
CGRect C_NSBrowser_FrameOfRow_InColumn(void* ptr, int row, int column);
bool C_NSBrowser_SendAction(void* ptr);
void C_NSBrowser_DoClick(void* ptr, void* sender);
void C_NSBrowser_DoDoubleClick(void* ptr, void* sender);
void C_NSBrowser_Browser_RemoveSavedColumnsWithAutosaveName(void* name);
double C_NSBrowser_ColumnContentWidthForColumnWidth(void* ptr, double columnWidth);
double C_NSBrowser_ColumnWidthForColumnContentWidth(void* ptr, double columnContentWidth);
double C_NSBrowser_WidthOfColumn(void* ptr, int column);
void C_NSBrowser_SetWidth_OfColumn(void* ptr, double columnWidth, int columnIndex);
double C_NSBrowser_DefaultColumnWidth(void* ptr);
void C_NSBrowser_SetDefaultColumnWidth(void* ptr, double columnWidth);
bool C_NSBrowser_ReusesColumns(void* ptr);
void C_NSBrowser_SetReusesColumns(void* ptr, bool value);
int C_NSBrowser_MaxVisibleColumns(void* ptr);
void C_NSBrowser_SetMaxVisibleColumns(void* ptr, int value);
bool C_NSBrowser_AutohidesScroller(void* ptr);
void C_NSBrowser_SetAutohidesScroller(void* ptr, bool value);
void* C_NSBrowser_BackgroundColor(void* ptr);
void C_NSBrowser_SetBackgroundColor(void* ptr, void* value);
double C_NSBrowser_MinColumnWidth(void* ptr);
void C_NSBrowser_SetMinColumnWidth(void* ptr, double value);
bool C_NSBrowser_SeparatesColumns(void* ptr);
void C_NSBrowser_SetSeparatesColumns(void* ptr, bool value);
bool C_NSBrowser_TakesTitleFromPreviousColumn(void* ptr);
void C_NSBrowser_SetTakesTitleFromPreviousColumn(void* ptr, bool value);
void* C_NSBrowser_Delegate(void* ptr);
void C_NSBrowser_SetDelegate(void* ptr, void* value);
void* C_NSBrowser_CellPrototype(void* ptr);
void C_NSBrowser_SetCellPrototype(void* ptr, void* value);
bool C_NSBrowser_AllowsBranchSelection(void* ptr);
void C_NSBrowser_SetAllowsBranchSelection(void* ptr, bool value);
bool C_NSBrowser_AllowsEmptySelection(void* ptr);
void C_NSBrowser_SetAllowsEmptySelection(void* ptr, bool value);
bool C_NSBrowser_AllowsMultipleSelection(void* ptr);
void C_NSBrowser_SetAllowsMultipleSelection(void* ptr, bool value);
bool C_NSBrowser_AllowsTypeSelect(void* ptr);
void C_NSBrowser_SetAllowsTypeSelect(void* ptr, bool value);
void* C_NSBrowser_SelectedCell(void* ptr);
Array C_NSBrowser_SelectedCells(void* ptr);
void* C_NSBrowser_SelectionIndexPath(void* ptr);
void C_NSBrowser_SetSelectionIndexPath(void* ptr, void* value);
Array C_NSBrowser_SelectionIndexPaths(void* ptr);
void C_NSBrowser_SetSelectionIndexPaths(void* ptr, Array value);
void* C_NSBrowser_PathSeparator(void* ptr);
void C_NSBrowser_SetPathSeparator(void* ptr, void* value);
int C_NSBrowser_SelectedColumn(void* ptr);
int C_NSBrowser_LastColumn(void* ptr);
void C_NSBrowser_SetLastColumn(void* ptr, int value);
int C_NSBrowser_FirstVisibleColumn(void* ptr);
int C_NSBrowser_NumberOfVisibleColumns(void* ptr);
int C_NSBrowser_LastVisibleColumn(void* ptr);
bool C_NSBrowser_IsLoaded(void* ptr);
bool C_NSBrowser_IsTitled(void* ptr);
void C_NSBrowser_SetTitled(void* ptr, bool value);
double C_NSBrowser_TitleHeight(void* ptr);
bool C_NSBrowser_HasHorizontalScroller(void* ptr);
void C_NSBrowser_SetHasHorizontalScroller(void* ptr, bool value);
void* C_NSBrowser_DoubleAction(void* ptr);
void C_NSBrowser_SetDoubleAction(void* ptr, void* value);
bool C_NSBrowser_SendsActionOnArrowKeys(void* ptr);
void C_NSBrowser_SetSendsActionOnArrowKeys(void* ptr, bool value);
int C_NSBrowser_ClickedColumn(void* ptr);
int C_NSBrowser_ClickedRow(void* ptr);
void* C_NSBrowser_ColumnsAutosaveName(void* ptr);
void C_NSBrowser_SetColumnsAutosaveName(void* ptr, void* value);
unsigned int C_NSBrowser_ColumnResizingType(void* ptr);
void C_NSBrowser_SetColumnResizingType(void* ptr, unsigned int value);
bool C_NSBrowser_PrefersAllColumnUserResizing(void* ptr);
void C_NSBrowser_SetPrefersAllColumnUserResizing(void* ptr, bool value);
double C_NSBrowser_RowHeight(void* ptr);
void C_NSBrowser_SetRowHeight(void* ptr, double value);
