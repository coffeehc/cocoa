#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_MenuItem_Alloc();

void* C_NSMenuItem_InitWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode);
void* C_NSMenuItem_InitWithCoder(void* ptr, void* coder);
void* C_NSMenuItem_AllocMenuItem();
void* C_NSMenuItem_Init(void* ptr);
void* C_NSMenuItem_NewMenuItem();
void* C_NSMenuItem_Autorelease(void* ptr);
void* C_NSMenuItem_Retain(void* ptr);
void* C_NSMenuItem_MenuItem_SeparatorItem();
bool C_NSMenuItem_IsEnabled(void* ptr);
void C_NSMenuItem_SetEnabled(void* ptr, bool value);
bool C_NSMenuItem_IsHidden(void* ptr);
void C_NSMenuItem_SetHidden(void* ptr, bool value);
bool C_NSMenuItem_IsHiddenOrHasHiddenAncestor(void* ptr);
void* C_NSMenuItem_Target(void* ptr);
void C_NSMenuItem_SetTarget(void* ptr, void* value);
void* C_NSMenuItem_Action(void* ptr);
void C_NSMenuItem_SetAction(void* ptr, void* value);
void* C_NSMenuItem_Title(void* ptr);
void C_NSMenuItem_SetTitle(void* ptr, void* value);
void* C_NSMenuItem_AttributedTitle(void* ptr);
void C_NSMenuItem_SetAttributedTitle(void* ptr, void* value);
int C_NSMenuItem_Tag(void* ptr);
void C_NSMenuItem_SetTag(void* ptr, int value);
int C_NSMenuItem_State(void* ptr);
void C_NSMenuItem_SetState(void* ptr, int value);
void* C_NSMenuItem_Image(void* ptr);
void C_NSMenuItem_SetImage(void* ptr, void* value);
void* C_NSMenuItem_OnStateImage(void* ptr);
void C_NSMenuItem_SetOnStateImage(void* ptr, void* value);
void* C_NSMenuItem_OffStateImage(void* ptr);
void C_NSMenuItem_SetOffStateImage(void* ptr, void* value);
void* C_NSMenuItem_MixedStateImage(void* ptr);
void C_NSMenuItem_SetMixedStateImage(void* ptr, void* value);
void* C_NSMenuItem_Submenu(void* ptr);
void C_NSMenuItem_SetSubmenu(void* ptr, void* value);
bool C_NSMenuItem_HasSubmenu(void* ptr);
void* C_NSMenuItem_ParentItem(void* ptr);
bool C_NSMenuItem_IsSeparatorItem(void* ptr);
void* C_NSMenuItem_Menu(void* ptr);
void C_NSMenuItem_SetMenu(void* ptr, void* value);
void* C_NSMenuItem_KeyEquivalent(void* ptr);
void C_NSMenuItem_SetKeyEquivalent(void* ptr, void* value);
unsigned int C_NSMenuItem_KeyEquivalentModifierMask(void* ptr);
void C_NSMenuItem_SetKeyEquivalentModifierMask(void* ptr, unsigned int value);
bool C_NSMenuItem_MenuItem_UsesUserKeyEquivalents();
void C_NSMenuItem_MenuItem_SetUsesUserKeyEquivalents(bool value);
void* C_NSMenuItem_UserKeyEquivalent(void* ptr);
bool C_NSMenuItem_IsAlternate(void* ptr);
void C_NSMenuItem_SetAlternate(void* ptr, bool value);
int C_NSMenuItem_IndentationLevel(void* ptr);
void C_NSMenuItem_SetIndentationLevel(void* ptr, int value);
void* C_NSMenuItem_ToolTip(void* ptr);
void C_NSMenuItem_SetToolTip(void* ptr, void* value);
void* C_NSMenuItem_RepresentedObject(void* ptr);
void C_NSMenuItem_SetRepresentedObject(void* ptr, void* value);
void* C_NSMenuItem_View(void* ptr);
void C_NSMenuItem_SetView(void* ptr, void* value);
bool C_NSMenuItem_IsHighlighted(void* ptr);
bool C_NSMenuItem_AllowsKeyEquivalentWhenHidden(void* ptr);
void C_NSMenuItem_SetAllowsKeyEquivalentWhenHidden(void* ptr, bool value);
