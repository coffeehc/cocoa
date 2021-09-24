#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ToolbarItem_Alloc();

void* C_NSToolbarItem_InitWithItemIdentifier(void* ptr, void* itemIdentifier);
void* C_NSToolbarItem_AllocToolbarItem();
void* C_NSToolbarItem_Autorelease(void* ptr);
void* C_NSToolbarItem_Retain(void* ptr);
void C_NSToolbarItem_Validate(void* ptr);
void* C_NSToolbarItem_ItemIdentifier(void* ptr);
void* C_NSToolbarItem_Toolbar(void* ptr);
void* C_NSToolbarItem_Label(void* ptr);
void C_NSToolbarItem_SetLabel(void* ptr, void* value);
void* C_NSToolbarItem_PaletteLabel(void* ptr);
void C_NSToolbarItem_SetPaletteLabel(void* ptr, void* value);
void* C_NSToolbarItem_ToolTip(void* ptr);
void C_NSToolbarItem_SetToolTip(void* ptr, void* value);
void* C_NSToolbarItem_Title(void* ptr);
void C_NSToolbarItem_SetTitle(void* ptr, void* value);
void* C_NSToolbarItem_MenuFormRepresentation(void* ptr);
void C_NSToolbarItem_SetMenuFormRepresentation(void* ptr, void* value);
int C_NSToolbarItem_Tag(void* ptr);
void C_NSToolbarItem_SetTag(void* ptr, int value);
void* C_NSToolbarItem_Target(void* ptr);
void C_NSToolbarItem_SetTarget(void* ptr, void* value);
void* C_NSToolbarItem_Action(void* ptr);
void C_NSToolbarItem_SetAction(void* ptr, void* value);
bool C_NSToolbarItem_IsBordered(void* ptr);
void C_NSToolbarItem_SetBordered(void* ptr, bool value);
bool C_NSToolbarItem_IsNavigational(void* ptr);
void C_NSToolbarItem_SetNavigational(void* ptr, bool value);
bool C_NSToolbarItem_IsEnabled(void* ptr);
void C_NSToolbarItem_SetEnabled(void* ptr, bool value);
void* C_NSToolbarItem_Image(void* ptr);
void C_NSToolbarItem_SetImage(void* ptr, void* value);
void* C_NSToolbarItem_View(void* ptr);
void C_NSToolbarItem_SetView(void* ptr, void* value);
int C_NSToolbarItem_VisibilityPriority(void* ptr);
void C_NSToolbarItem_SetVisibilityPriority(void* ptr, int value);
bool C_NSToolbarItem_Autovalidates(void* ptr);
void C_NSToolbarItem_SetAutovalidates(void* ptr, bool value);
bool C_NSToolbarItem_AllowsDuplicatesInToolbar(void* ptr);
