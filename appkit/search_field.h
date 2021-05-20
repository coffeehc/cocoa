#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_SearchField_Alloc();

void* C_NSSearchField_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSearchField_InitWithCoder(void* ptr, void* coder);
void* C_NSSearchField_Init(void* ptr);
void* C_NSSearchField_SearchMenuTemplate(void* ptr);
void C_NSSearchField_SetSearchMenuTemplate(void* ptr, void* value);
bool C_NSSearchField_SendsSearchStringImmediately(void* ptr);
void C_NSSearchField_SetSendsSearchStringImmediately(void* ptr, bool value);
bool C_NSSearchField_SendsWholeSearchString(void* ptr);
void C_NSSearchField_SetSendsWholeSearchString(void* ptr, bool value);
Array C_NSSearchField_RecentSearches(void* ptr);
void C_NSSearchField_SetRecentSearches(void* ptr, Array value);
int C_NSSearchField_MaximumRecents(void* ptr);
void C_NSSearchField_SetMaximumRecents(void* ptr, int value);
void* C_NSSearchField_RecentsAutosaveName(void* ptr);
void C_NSSearchField_SetRecentsAutosaveName(void* ptr, void* value);
CGRect C_NSSearchField_CancelButtonBounds(void* ptr);
CGRect C_NSSearchField_SearchButtonBounds(void* ptr);
CGRect C_NSSearchField_SearchTextBounds(void* ptr);
