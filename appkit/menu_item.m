#import <Appkit/Appkit.h>
#import "menu_item.h"

void* C_MenuItem_Alloc() {
    return [NSMenuItem alloc];
}

void* C_NSMenuItem_InitWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result_ = [nSMenuItem initWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode];
    return result_;
}

void* C_NSMenuItem_InitWithCoder(void* ptr, void* coder) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result_ = [nSMenuItem initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSMenuItem_Init(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result_ = [nSMenuItem init];
    return result_;
}

void* C_NSMenuItem_MenuItem_SeparatorItem() {
    NSMenuItem* result_ = [NSMenuItem separatorItem];
    return result_;
}

bool C_NSMenuItem_IsEnabled(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isEnabled];
    return result_;
}

void C_NSMenuItem_SetEnabled(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setEnabled:value];
}

bool C_NSMenuItem_IsHidden(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isHidden];
    return result_;
}

void C_NSMenuItem_SetHidden(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setHidden:value];
}

bool C_NSMenuItem_IsHiddenOrHasHiddenAncestor(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isHiddenOrHasHiddenAncestor];
    return result_;
}

void* C_NSMenuItem_Target(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    id result_ = [nSMenuItem target];
    return result_;
}

void C_NSMenuItem_SetTarget(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTarget:(id)value];
}

void* C_NSMenuItem_Action(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    SEL result_ = [nSMenuItem action];
    return result_;
}

void C_NSMenuItem_SetAction(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAction:(SEL)value];
}

void* C_NSMenuItem_Title(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result_ = [nSMenuItem title];
    return result_;
}

void C_NSMenuItem_SetTitle(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTitle:(NSString*)value];
}

void* C_NSMenuItem_AttributedTitle(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSAttributedString* result_ = [nSMenuItem attributedTitle];
    return result_;
}

void C_NSMenuItem_SetAttributedTitle(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAttributedTitle:(NSAttributedString*)value];
}

int C_NSMenuItem_Tag(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSInteger result_ = [nSMenuItem tag];
    return result_;
}

void C_NSMenuItem_SetTag(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setTag:value];
}

int C_NSMenuItem_State(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSControlStateValue result_ = [nSMenuItem state];
    return result_;
}

void C_NSMenuItem_SetState(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setState:value];
}

void* C_NSMenuItem_Image(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result_ = [nSMenuItem image];
    return result_;
}

void C_NSMenuItem_SetImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setImage:(NSImage*)value];
}

void* C_NSMenuItem_OnStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result_ = [nSMenuItem onStateImage];
    return result_;
}

void C_NSMenuItem_SetOnStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setOnStateImage:(NSImage*)value];
}

void* C_NSMenuItem_OffStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result_ = [nSMenuItem offStateImage];
    return result_;
}

void C_NSMenuItem_SetOffStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setOffStateImage:(NSImage*)value];
}

void* C_NSMenuItem_MixedStateImage(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSImage* result_ = [nSMenuItem mixedStateImage];
    return result_;
}

void C_NSMenuItem_SetMixedStateImage(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setMixedStateImage:(NSImage*)value];
}

void* C_NSMenuItem_Submenu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenu* result_ = [nSMenuItem submenu];
    return result_;
}

void C_NSMenuItem_SetSubmenu(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setSubmenu:(NSMenu*)value];
}

bool C_NSMenuItem_HasSubmenu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem hasSubmenu];
    return result_;
}

void* C_NSMenuItem_ParentItem(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenuItem* result_ = [nSMenuItem parentItem];
    return result_;
}

bool C_NSMenuItem_IsSeparatorItem(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isSeparatorItem];
    return result_;
}

void* C_NSMenuItem_Menu(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSMenu* result_ = [nSMenuItem menu];
    return result_;
}

void C_NSMenuItem_SetMenu(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setMenu:(NSMenu*)value];
}

void* C_NSMenuItem_KeyEquivalent(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result_ = [nSMenuItem keyEquivalent];
    return result_;
}

void C_NSMenuItem_SetKeyEquivalent(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setKeyEquivalent:(NSString*)value];
}

unsigned int C_NSMenuItem_KeyEquivalentModifierMask(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSEventModifierFlags result_ = [nSMenuItem keyEquivalentModifierMask];
    return result_;
}

void C_NSMenuItem_SetKeyEquivalentModifierMask(void* ptr, unsigned int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setKeyEquivalentModifierMask:value];
}

bool C_NSMenuItem_MenuItem_UsesUserKeyEquivalents() {
    BOOL result_ = [NSMenuItem usesUserKeyEquivalents];
    return result_;
}

void C_NSMenuItem_MenuItem_SetUsesUserKeyEquivalents(bool value) {
    [NSMenuItem setUsesUserKeyEquivalents:value];
}

void* C_NSMenuItem_UserKeyEquivalent(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result_ = [nSMenuItem userKeyEquivalent];
    return result_;
}

bool C_NSMenuItem_IsAlternate(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isAlternate];
    return result_;
}

void C_NSMenuItem_SetAlternate(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAlternate:value];
}

int C_NSMenuItem_IndentationLevel(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSInteger result_ = [nSMenuItem indentationLevel];
    return result_;
}

void C_NSMenuItem_SetIndentationLevel(void* ptr, int value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setIndentationLevel:value];
}

void* C_NSMenuItem_ToolTip(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSString* result_ = [nSMenuItem toolTip];
    return result_;
}

void C_NSMenuItem_SetToolTip(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setToolTip:(NSString*)value];
}

void* C_NSMenuItem_RepresentedObject(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    id result_ = [nSMenuItem representedObject];
    return result_;
}

void C_NSMenuItem_SetRepresentedObject(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setRepresentedObject:(id)value];
}

void* C_NSMenuItem_View(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    NSView* result_ = [nSMenuItem view];
    return result_;
}

void C_NSMenuItem_SetView(void* ptr, void* value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setView:(NSView*)value];
}

bool C_NSMenuItem_IsHighlighted(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem isHighlighted];
    return result_;
}

bool C_NSMenuItem_AllowsKeyEquivalentWhenHidden(void* ptr) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    BOOL result_ = [nSMenuItem allowsKeyEquivalentWhenHidden];
    return result_;
}

void C_NSMenuItem_SetAllowsKeyEquivalentWhenHidden(void* ptr, bool value) {
    NSMenuItem* nSMenuItem = (NSMenuItem*)ptr;
    [nSMenuItem setAllowsKeyEquivalentWhenHidden:value];
}
