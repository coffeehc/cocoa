#import <Appkit/Appkit.h>
#import "menu_item.h"

void* C_MenuItem_Alloc() {
    return [NSMenuItem alloc];
}

void* C_NSMenuItem_InitWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result = [nSMenuItem initWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode];
    return result;
}

void* C_NSMenuItem_InitWithCoder(void* ptr, void* coder) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result = [nSMenuItem initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSMenuItem_Init(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result = [nSMenuItem init];
    return result;
}

void* C_NSMenuItem_MenuItem_SeparatorItem() {
    NSMenuItem* result = [NSMenuItem separatorItem];
    return result;
}

bool C_NSMenuItem_IsEnabled(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isEnabled];
    return result;
}

void C_NSMenuItem_SetEnabled(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setEnabled:value];
}

bool C_NSMenuItem_IsHidden(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isHidden];
    return result;
}

void C_NSMenuItem_SetHidden(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setHidden:value];
}

bool C_NSMenuItem_IsHiddenOrHasHiddenAncestor(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isHiddenOrHasHiddenAncestor];
    return result;
}

void* C_NSMenuItem_Target(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    id result = [nSMenuItem target];
    return result;
}

void C_NSMenuItem_SetTarget(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTarget:(id)value];
}

void* C_NSMenuItem_Action(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    SEL result = [nSMenuItem action];
    return result;
}

void C_NSMenuItem_SetAction(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAction:(SEL)value];
}

void* C_NSMenuItem_Title(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result = [nSMenuItem title];
    return result;
}

void C_NSMenuItem_SetTitle(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTitle:(NSString*)value];
}

void* C_NSMenuItem_AttributedTitle(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSAttributedString* result = [nSMenuItem attributedTitle];
    return result;
}

void C_NSMenuItem_SetAttributedTitle(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAttributedTitle:(NSAttributedString*)value];
}

int C_NSMenuItem_Tag(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSInteger result = [nSMenuItem tag];
    return result;
}

void C_NSMenuItem_SetTag(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTag:value];
}

int C_NSMenuItem_State(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSControlStateValue result = [nSMenuItem state];
    return result;
}

void C_NSMenuItem_SetState(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setState:value];
}

void* C_NSMenuItem_Image(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result = [nSMenuItem image];
    return result;
}

void C_NSMenuItem_SetImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setImage:(NSImage*)value];
}

void* C_NSMenuItem_OnStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result = [nSMenuItem onStateImage];
    return result;
}

void C_NSMenuItem_SetOnStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setOnStateImage:(NSImage*)value];
}

void* C_NSMenuItem_OffStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result = [nSMenuItem offStateImage];
    return result;
}

void C_NSMenuItem_SetOffStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setOffStateImage:(NSImage*)value];
}

void* C_NSMenuItem_MixedStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result = [nSMenuItem mixedStateImage];
    return result;
}

void C_NSMenuItem_SetMixedStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setMixedStateImage:(NSImage*)value];
}

void* C_NSMenuItem_Submenu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenu* result = [nSMenuItem submenu];
    return result;
}

void C_NSMenuItem_SetSubmenu(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setSubmenu:(NSMenu*)value];
}

bool C_NSMenuItem_HasSubmenu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem hasSubmenu];
    return result;
}

void* C_NSMenuItem_ParentItem(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result = [nSMenuItem parentItem];
    return result;
}

bool C_NSMenuItem_IsSeparatorItem(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isSeparatorItem];
    return result;
}

void* C_NSMenuItem_Menu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenu* result = [nSMenuItem menu];
    return result;
}

void C_NSMenuItem_SetMenu(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setMenu:(NSMenu*)value];
}

void* C_NSMenuItem_KeyEquivalent(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result = [nSMenuItem keyEquivalent];
    return result;
}

void C_NSMenuItem_SetKeyEquivalent(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setKeyEquivalent:(NSString*)value];
}

unsigned int C_NSMenuItem_KeyEquivalentModifierMask(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSEventModifierFlags result = [nSMenuItem keyEquivalentModifierMask];
    return result;
}

void C_NSMenuItem_SetKeyEquivalentModifierMask(void* ptr, unsigned int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setKeyEquivalentModifierMask:value];
}

bool C_NSMenuItem_MenuItem_UsesUserKeyEquivalents() {
    BOOL result = [NSMenuItem usesUserKeyEquivalents];
    return result;
}

void C_NSMenuItem_MenuItem_SetUsesUserKeyEquivalents(bool value) {
    [NSMenuItem setUsesUserKeyEquivalents:value];
}

void* C_NSMenuItem_UserKeyEquivalent(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result = [nSMenuItem userKeyEquivalent];
    return result;
}

bool C_NSMenuItem_IsAlternate(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isAlternate];
    return result;
}

void C_NSMenuItem_SetAlternate(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAlternate:value];
}

int C_NSMenuItem_IndentationLevel(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSInteger result = [nSMenuItem indentationLevel];
    return result;
}

void C_NSMenuItem_SetIndentationLevel(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setIndentationLevel:value];
}

void* C_NSMenuItem_ToolTip(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result = [nSMenuItem toolTip];
    return result;
}

void C_NSMenuItem_SetToolTip(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setToolTip:(NSString*)value];
}

void* C_NSMenuItem_RepresentedObject(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    id result = [nSMenuItem representedObject];
    return result;
}

void C_NSMenuItem_SetRepresentedObject(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setRepresentedObject:(id)value];
}

void* C_NSMenuItem_View(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSView* result = [nSMenuItem view];
    return result;
}

void C_NSMenuItem_SetView(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setView:(NSView*)value];
}

bool C_NSMenuItem_IsHighlighted(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem isHighlighted];
    return result;
}

bool C_NSMenuItem_AllowsKeyEquivalentWhenHidden(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result = [nSMenuItem allowsKeyEquivalentWhenHidden];
    return result;
}

void C_NSMenuItem_SetAllowsKeyEquivalentWhenHidden(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAllowsKeyEquivalentWhenHidden:value];
}
