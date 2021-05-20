#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TableRowView_Alloc();

void* C_NSTableRowView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTableRowView_InitWithCoder(void* ptr, void* coder);
void* C_NSTableRowView_Init(void* ptr);
void C_NSTableRowView_DrawBackgroundInRect(void* ptr, CGRect dirtyRect);
void C_NSTableRowView_DrawDraggingDestinationFeedbackInRect(void* ptr, CGRect dirtyRect);
void C_NSTableRowView_DrawSelectionInRect(void* ptr, CGRect dirtyRect);
void C_NSTableRowView_DrawSeparatorInRect(void* ptr, CGRect dirtyRect);
void* C_NSTableRowView_ViewAtColumn(void* ptr, int column);
bool C_NSTableRowView_IsEmphasized(void* ptr);
void C_NSTableRowView_SetEmphasized(void* ptr, bool value);
int C_NSTableRowView_InteriorBackgroundStyle(void* ptr);
bool C_NSTableRowView_IsFloating(void* ptr);
void C_NSTableRowView_SetFloating(void* ptr, bool value);
bool C_NSTableRowView_IsSelected(void* ptr);
void C_NSTableRowView_SetSelected(void* ptr, bool value);
int C_NSTableRowView_SelectionHighlightStyle(void* ptr);
void C_NSTableRowView_SetSelectionHighlightStyle(void* ptr, int value);
int C_NSTableRowView_DraggingDestinationFeedbackStyle(void* ptr);
void C_NSTableRowView_SetDraggingDestinationFeedbackStyle(void* ptr, int value);
double C_NSTableRowView_IndentationForDropOperation(void* ptr);
void C_NSTableRowView_SetIndentationForDropOperation(void* ptr, double value);
bool C_NSTableRowView_IsTargetForDropOperation(void* ptr);
void C_NSTableRowView_SetTargetForDropOperation(void* ptr, bool value);
bool C_NSTableRowView_IsGroupRowStyle(void* ptr);
void C_NSTableRowView_SetGroupRowStyle(void* ptr, bool value);
int C_NSTableRowView_NumberOfColumns(void* ptr);
void* C_NSTableRowView_BackgroundColor(void* ptr);
void C_NSTableRowView_SetBackgroundColor(void* ptr, void* value);
bool C_NSTableRowView_IsNextRowSelected(void* ptr);
void C_NSTableRowView_SetNextRowSelected(void* ptr, bool value);
bool C_NSTableRowView_IsPreviousRowSelected(void* ptr);
void C_NSTableRowView_SetPreviousRowSelected(void* ptr, bool value);
