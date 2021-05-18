#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Menu_Alloc();

void* C_NSMenu_InitWithTitle(void* ptr, void* title);
void* C_NSMenu_InitWithCoder(void* ptr, void* coder);
void* C_NSMenu_Init(void* ptr);
bool C_NSMenu_MenuBarVisible();
void C_NSMenu_Menu_SetMenuBarVisible(bool visible);
void C_NSMenu_InsertItem_AtIndex(void* ptr, void* newItem, int index);
void* C_NSMenu_InsertItemWithTitle_Action_KeyEquivalent_AtIndex(void* ptr, void* _string, void* selector, void* charCode, int index);
void C_NSMenu_AddItem(void* ptr, void* newItem);
void* C_NSMenu_AddItemWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode);
void C_NSMenu_RemoveItem(void* ptr, void* item);
void C_NSMenu_RemoveItemAtIndex(void* ptr, int index);
void C_NSMenu_ItemChanged(void* ptr, void* item);
void C_NSMenu_RemoveAllItems(void* ptr);
void* C_NSMenu_ItemWithTag(void* ptr, int tag);
void* C_NSMenu_ItemWithTitle(void* ptr, void* title);
void* C_NSMenu_ItemAtIndex(void* ptr, int index);
int C_NSMenu_IndexOfItem(void* ptr, void* item);
int C_NSMenu_IndexOfItemWithTitle(void* ptr, void* title);
int C_NSMenu_IndexOfItemWithTag(void* ptr, int tag);
int C_NSMenu_IndexOfItemWithTarget_AndAction(void* ptr, void* target, void* actionSelector);
int C_NSMenu_IndexOfItemWithRepresentedObject(void* ptr, void* object);
int C_NSMenu_IndexOfItemWithSubmenu(void* ptr, void* submenu);
void C_NSMenu_SetSubmenu_ForItem(void* ptr, void* menu, void* item);
void C_NSMenu_SubmenuAction(void* ptr, void* sender);
void C_NSMenu_Update(void* ptr);
bool C_NSMenu_PerformKeyEquivalent(void* ptr, void* event);
void C_NSMenu_PerformActionForItemAtIndex(void* ptr, int index);
void C_NSMenu_PopUpContextMenu_WithEvent_ForView(void* menu, void* event, void* view);
void C_NSMenu_PopUpContextMenu_WithEvent_ForView_WithFont(void* menu, void* event, void* view, void* font);
bool C_NSMenu_PopUpMenuPositioningItem_AtLocation_InView(void* ptr, void* item, CGPoint location, void* view);
void C_NSMenu_CancelTracking(void* ptr);
void C_NSMenu_CancelTrackingWithoutAnimation(void* ptr);
double C_NSMenu_MenuBarHeight(void* ptr);
int C_NSMenu_NumberOfItems(void* ptr);
Array C_NSMenu_ItemArray(void* ptr);
void C_NSMenu_SetItemArray(void* ptr, Array value);
void* C_NSMenu_Supermenu(void* ptr);
void C_NSMenu_SetSupermenu(void* ptr, void* value);
bool C_NSMenu_AutoenablesItems(void* ptr);
void C_NSMenu_SetAutoenablesItems(void* ptr, bool value);
void* C_NSMenu_Font(void* ptr);
void C_NSMenu_SetFont(void* ptr, void* value);
void* C_NSMenu_Title(void* ptr);
void C_NSMenu_SetTitle(void* ptr, void* value);
double C_NSMenu_MinimumWidth(void* ptr);
void C_NSMenu_SetMinimumWidth(void* ptr, double value);
CGSize C_NSMenu_Size(void* ptr);
unsigned int C_NSMenu_PropertiesToUpdate(void* ptr);
bool C_NSMenu_AllowsContextMenuPlugIns(void* ptr);
void C_NSMenu_SetAllowsContextMenuPlugIns(void* ptr, bool value);
bool C_NSMenu_ShowsStateColumn(void* ptr);
void C_NSMenu_SetShowsStateColumn(void* ptr, bool value);
void* C_NSMenu_HighlightedItem(void* ptr);
int C_NSMenu_UserInterfaceLayoutDirection(void* ptr);
void C_NSMenu_SetUserInterfaceLayoutDirection(void* ptr, int value);
void* C_NSMenu_Delegate(void* ptr);
void C_NSMenu_SetDelegate(void* ptr, void* value);
