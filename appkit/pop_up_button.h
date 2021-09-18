#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_PopUpButton_Alloc();

void* C_NSPopUpButton_InitWithFrame_PullsDown(void* ptr, CGRect buttonFrame, bool flag);
void* C_NSPopUpButton_PopUpButton_CheckboxWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSPopUpButton_PopUpButton_ButtonWithImage_Target_Action(void* image, void* target, void* action);
void* C_NSPopUpButton_PopUpButton_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSPopUpButton_PopUpButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action);
void* C_NSPopUpButton_PopUpButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSPopUpButton_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSPopUpButton_InitWithCoder(void* ptr, void* coder);
void* C_NSPopUpButton_Init(void* ptr);
void* C_NSPopUpButton_AllocPopUpButton();
void* C_NSPopUpButton_NewPopUpButton();
void* C_NSPopUpButton_Autorelease(void* ptr);
void* C_NSPopUpButton_Retain(void* ptr);
void C_NSPopUpButton_AddItemWithTitle(void* ptr, void* title);
void C_NSPopUpButton_AddItemsWithTitles(void* ptr, Array itemTitles);
void C_NSPopUpButton_InsertItemWithTitle_AtIndex(void* ptr, void* title, int index);
void C_NSPopUpButton_RemoveAllItems(void* ptr);
void C_NSPopUpButton_RemoveItemWithTitle(void* ptr, void* title);
void C_NSPopUpButton_RemoveItemAtIndex(void* ptr, int index);
void C_NSPopUpButton_SelectItem(void* ptr, void* item);
void C_NSPopUpButton_SelectItemAtIndex(void* ptr, int index);
bool C_NSPopUpButton_SelectItemWithTag(void* ptr, int tag);
void C_NSPopUpButton_SelectItemWithTitle(void* ptr, void* title);
void* C_NSPopUpButton_ItemAtIndex(void* ptr, int index);
void* C_NSPopUpButton_ItemTitleAtIndex(void* ptr, int index);
void* C_NSPopUpButton_ItemWithTitle(void* ptr, void* title);
int C_NSPopUpButton_IndexOfItem(void* ptr, void* item);
int C_NSPopUpButton_IndexOfItemWithTag(void* ptr, int tag);
int C_NSPopUpButton_IndexOfItemWithTitle(void* ptr, void* title);
int C_NSPopUpButton_IndexOfItemWithRepresentedObject(void* ptr, void* obj);
int C_NSPopUpButton_IndexOfItemWithTarget_AndAction(void* ptr, void* target, void* actionSelector);
void C_NSPopUpButton_SynchronizeTitleAndSelectedItem(void* ptr);
bool C_NSPopUpButton_PullsDown(void* ptr);
void C_NSPopUpButton_SetPullsDown(void* ptr, bool value);
bool C_NSPopUpButton_AutoenablesItems(void* ptr);
void C_NSPopUpButton_SetAutoenablesItems(void* ptr, bool value);
void* C_NSPopUpButton_SelectedItem(void* ptr);
void* C_NSPopUpButton_TitleOfSelectedItem(void* ptr);
int C_NSPopUpButton_IndexOfSelectedItem(void* ptr);
int C_NSPopUpButton_SelectedTag(void* ptr);
int C_NSPopUpButton_NumberOfItems(void* ptr);
Array C_NSPopUpButton_ItemArray(void* ptr);
Array C_NSPopUpButton_ItemTitles(void* ptr);
void* C_NSPopUpButton_LastItem(void* ptr);
unsigned int C_NSPopUpButton_PreferredEdge(void* ptr);
void C_NSPopUpButton_SetPreferredEdge(void* ptr, unsigned int value);
