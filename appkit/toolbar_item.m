#import "toolbar_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSToolbarItem.h>

void* C_ToolbarItem_Alloc() {
    return [NSToolbarItem alloc];
}

void* C_NSToolbarItem_InitWithItemIdentifier(void* ptr, void* itemIdentifier) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbarItem* result_ = [nSToolbarItem initWithItemIdentifier:(NSString*)itemIdentifier];
    return result_;
}

void* C_NSToolbarItem_AllocToolbarItem() {
    NSToolbarItem* result_ = [NSToolbarItem alloc];
    return result_;
}

void* C_NSToolbarItem_Autorelease(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbarItem* result_ = [nSToolbarItem autorelease];
    return result_;
}

void* C_NSToolbarItem_Retain(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbarItem* result_ = [nSToolbarItem retain];
    return result_;
}

void C_NSToolbarItem_Validate(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem validate];
}

void* C_NSToolbarItem_ItemIdentifier(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbarItemIdentifier result_ = [nSToolbarItem itemIdentifier];
    return result_;
}

void* C_NSToolbarItem_Toolbar(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbar* result_ = [nSToolbarItem toolbar];
    return result_;
}

void* C_NSToolbarItem_Label(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSString* result_ = [nSToolbarItem label];
    return result_;
}

void C_NSToolbarItem_SetLabel(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setLabel:(NSString*)value];
}

void* C_NSToolbarItem_PaletteLabel(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSString* result_ = [nSToolbarItem paletteLabel];
    return result_;
}

void C_NSToolbarItem_SetPaletteLabel(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setPaletteLabel:(NSString*)value];
}

void* C_NSToolbarItem_ToolTip(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSString* result_ = [nSToolbarItem toolTip];
    return result_;
}

void C_NSToolbarItem_SetToolTip(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setToolTip:(NSString*)value];
}

void* C_NSToolbarItem_Title(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSString* result_ = [nSToolbarItem title];
    return result_;
}

void C_NSToolbarItem_SetTitle(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setTitle:(NSString*)value];
}

void* C_NSToolbarItem_MenuFormRepresentation(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSMenuItem* result_ = [nSToolbarItem menuFormRepresentation];
    return result_;
}

void C_NSToolbarItem_SetMenuFormRepresentation(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setMenuFormRepresentation:(NSMenuItem*)value];
}

int C_NSToolbarItem_Tag(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSInteger result_ = [nSToolbarItem tag];
    return result_;
}

void C_NSToolbarItem_SetTag(void* ptr, int value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setTag:value];
}

void* C_NSToolbarItem_Target(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    id result_ = [nSToolbarItem target];
    return result_;
}

void C_NSToolbarItem_SetTarget(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setTarget:(id)value];
}

void* C_NSToolbarItem_Action(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    SEL result_ = [nSToolbarItem action];
    return result_;
}

void C_NSToolbarItem_SetAction(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setAction:(SEL)value];
}

bool C_NSToolbarItem_IsBordered(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    BOOL result_ = [nSToolbarItem isBordered];
    return result_;
}

void C_NSToolbarItem_SetBordered(void* ptr, bool value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setBordered:value];
}

bool C_NSToolbarItem_IsNavigational(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    BOOL result_ = [nSToolbarItem isNavigational];
    return result_;
}

void C_NSToolbarItem_SetNavigational(void* ptr, bool value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setNavigational:value];
}

bool C_NSToolbarItem_IsEnabled(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    BOOL result_ = [nSToolbarItem isEnabled];
    return result_;
}

void C_NSToolbarItem_SetEnabled(void* ptr, bool value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setEnabled:value];
}

void* C_NSToolbarItem_Image(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSImage* result_ = [nSToolbarItem image];
    return result_;
}

void C_NSToolbarItem_SetImage(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setImage:(NSImage*)value];
}

void* C_NSToolbarItem_View(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSView* result_ = [nSToolbarItem view];
    return result_;
}

void C_NSToolbarItem_SetView(void* ptr, void* value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setView:(NSView*)value];
}

int C_NSToolbarItem_VisibilityPriority(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    NSToolbarItemVisibilityPriority result_ = [nSToolbarItem visibilityPriority];
    return result_;
}

void C_NSToolbarItem_SetVisibilityPriority(void* ptr, int value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setVisibilityPriority:value];
}

bool C_NSToolbarItem_Autovalidates(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    BOOL result_ = [nSToolbarItem autovalidates];
    return result_;
}

void C_NSToolbarItem_SetAutovalidates(void* ptr, bool value) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    [nSToolbarItem setAutovalidates:value];
}

bool C_NSToolbarItem_AllowsDuplicatesInToolbar(void* ptr) {
    NSToolbarItem* nSToolbarItem = (NSToolbarItem*)ptr;
    BOOL result_ = [nSToolbarItem allowsDuplicatesInToolbar];
    return result_;
}
