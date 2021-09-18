#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_SearchFieldCell_Alloc();

void* C_NSSearchFieldCell_InitWithCoder(void* ptr, void* coder);
void* C_NSSearchFieldCell_InitTextCell(void* ptr, void* _string);
void* C_NSSearchFieldCell_Init(void* ptr);
void* C_NSSearchFieldCell_AllocSearchFieldCell();
void* C_NSSearchFieldCell_NewSearchFieldCell();
void* C_NSSearchFieldCell_Autorelease(void* ptr);
void* C_NSSearchFieldCell_Retain(void* ptr);
void C_NSSearchFieldCell_ResetSearchButtonCell(void* ptr);
void C_NSSearchFieldCell_ResetCancelButtonCell(void* ptr);
CGRect C_NSSearchFieldCell_SearchTextRectForBounds(void* ptr, CGRect rect);
CGRect C_NSSearchFieldCell_SearchButtonRectForBounds(void* ptr, CGRect rect);
CGRect C_NSSearchFieldCell_CancelButtonRectForBounds(void* ptr, CGRect rect);
void* C_NSSearchFieldCell_SearchButtonCell(void* ptr);
void C_NSSearchFieldCell_SetSearchButtonCell(void* ptr, void* value);
void* C_NSSearchFieldCell_CancelButtonCell(void* ptr);
void C_NSSearchFieldCell_SetCancelButtonCell(void* ptr, void* value);
void* C_NSSearchFieldCell_SearchMenuTemplate(void* ptr);
void C_NSSearchFieldCell_SetSearchMenuTemplate(void* ptr, void* value);
bool C_NSSearchFieldCell_SendsWholeSearchString(void* ptr);
void C_NSSearchFieldCell_SetSendsWholeSearchString(void* ptr, bool value);
bool C_NSSearchFieldCell_SendsSearchStringImmediately(void* ptr);
void C_NSSearchFieldCell_SetSendsSearchStringImmediately(void* ptr, bool value);
int C_NSSearchFieldCell_MaximumRecents(void* ptr);
void C_NSSearchFieldCell_SetMaximumRecents(void* ptr, int value);
Array C_NSSearchFieldCell_RecentSearches(void* ptr);
void C_NSSearchFieldCell_SetRecentSearches(void* ptr, Array value);
void* C_NSSearchFieldCell_RecentsAutosaveName(void* ptr);
void C_NSSearchFieldCell_SetRecentsAutosaveName(void* ptr, void* value);
