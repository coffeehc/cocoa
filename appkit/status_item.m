#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "status_item.h"

void* StatusItem_StatusBar(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem statusBar];
}

unsigned long StatusItem_Behavior(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem behavior];
}

void StatusItem_SetBehavior(void* ptr, unsigned long behavior) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	[statusItem setBehavior:behavior];
}

void* StatusItem_Button(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem button];
}

void* StatusItem_Menu(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem menu];
}

void StatusItem_SetMenu(void* ptr, void* menu) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	[statusItem setMenu:(NSMenu*)menu];
}

bool StatusItem_IsVisible(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem isVisible];
}

void StatusItem_SetVisible(void* ptr, bool visible) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	[statusItem setVisible:visible];
}

double StatusItem_Length(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [statusItem length];
}

void StatusItem_SetLength(void* ptr, double length) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	[statusItem setLength:length];
}

const char* StatusItem_AutosaveName(void* ptr) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	return [[statusItem autosaveName] UTF8String];
}

void StatusItem_SetAutosaveName(void* ptr, const char* autosaveName) {
	NSStatusItem* statusItem = (NSStatusItem*)ptr;
	[statusItem setAutosaveName:[NSString stringWithUTF8String:autosaveName]];
}
