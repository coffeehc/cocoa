#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "menu_item.h"
#import "menu_item_delegate.h"

@implementation NSMenuItemHandler

- (void)onAction:(id)sender {
	return MenuItem_Target_Action([self goID], sender);
}

@end
void MenuItem_SetDelegate(void *ptr, long goID) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	NSMenuItemHandler* handler = [[NSMenuItemHandler alloc] init];
	[handler setGoID:goID];
	[menuItem setTarget:handler];
	[menuItem setAction:@selector(onAction:)];
}

bool MenuItem_IsEnabled(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isEnabled];
}

void MenuItem_SetEnabled(void* ptr, bool enabled) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setEnabled:enabled];
}

bool MenuItem_IsHidden(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isHidden];
}

void MenuItem_SetHidden(void* ptr, bool hidden) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setHidden:hidden];
}

const char* MenuItem_Title(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem title] UTF8String];
}

void MenuItem_SetTitle(void* ptr, const char* title) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setTitle:[NSString stringWithUTF8String:title]];
}

void* MenuItem_Submenu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem submenu];
}

void MenuItem_SetSubmenu(void* ptr, void* submenu) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setSubmenu:(NSMenu*)submenu];
}

bool MenuItem_HasSubmenu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem hasSubmenu];
}

bool MenuItem_IsSeparatorItem(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isSeparatorItem];
}

void* MenuItem_Menu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem menu];
}

void MenuItem_SetMenu(void* ptr, void* menu) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setMenu:(NSMenu*)menu];
}

const char* MenuItem_ToolTip(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem toolTip] UTF8String];
}

void MenuItem_SetToolTip(void* ptr, const char* toolTip) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setToolTip:[NSString stringWithUTF8String:toolTip]];
}

bool MenuItem_IsHighlighted(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isHighlighted];
}

const char* MenuItem_KeyEquivalent(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem keyEquivalent] UTF8String];
}

void MenuItem_SetKeyEquivalent(void* ptr, const char* keyEquivalent) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setKeyEquivalent:[NSString stringWithUTF8String:keyEquivalent]];
}

unsigned long MenuItem_KeyEquivalentModifierMask(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem keyEquivalentModifierMask];
}

void MenuItem_SetKeyEquivalentModifierMask(void* ptr, unsigned long keyEquivalentModifierMask) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setKeyEquivalentModifierMask:keyEquivalentModifierMask];
}

const char* MenuItem_UserKeyEquivalent(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem userKeyEquivalent] UTF8String];
}

bool MenuItem_IsAlternate(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isAlternate];
}

void MenuItem_SetAlternate(void* ptr, bool alternate) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setAlternate:alternate];
}

long MenuItem_IndentationLevel(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem indentationLevel];
}

void MenuItem_SetIndentationLevel(void* ptr, long indentationLevel) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setIndentationLevel:indentationLevel];
}

void* MenuItem_View(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem view];
}

void MenuItem_SetView(void* ptr, void* view) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setView:(NSView*)view];
}

bool MenuItem_AllowsKeyEquivalentWhenHidden(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem allowsKeyEquivalentWhenHidden];
}

void MenuItem_SetAllowsKeyEquivalentWhenHidden(void* ptr, bool allowsKeyEquivalentWhenHidden) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setAllowsKeyEquivalentWhenHidden:allowsKeyEquivalentWhenHidden];
}

bool MenuItem_UsesUserKeyEquivalents() {
	return [NSMenuItem usesUserKeyEquivalents];
}

void MenuItem_SetUsesUserKeyEquivalents(bool usesUserKeyEquivalents) {
	[NSMenuItem setUsesUserKeyEquivalents:usesUserKeyEquivalents];
}

long MenuItem_State(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem state];
}

void MenuItem_SetState(void* ptr, long state) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setState:state];
}

long MenuItem_Tag(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem tag];
}

void MenuItem_SetTag(void* ptr, long tag) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setTag:tag];
}

void* MenuItem_NewMenuItem(const char* title, void* selector, const char* charCode) {
	NSMenuItem* menuItem = [NSMenuItem alloc];
	return [[menuItem initWithTitle:[NSString stringWithUTF8String:title] action:(SEL)selector keyEquivalent:[NSString stringWithUTF8String:charCode]] autorelease];
}

void* MenuItem_NewSeparatorItem() {
	return [NSMenuItem separatorItem];
}
