#import <Appkit/Appkit.h>
#import "window_tab_group.h"

void* C_WindowTabGroup_Alloc() {
	return [NSWindowTabGroup alloc];
}

void* C_NSWindowTabGroup_Init(void* ptr) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	NSWindowTabGroup* result = [nSWindowTabGroup init];
	return result;
}

void C_NSWindowTabGroup_AddWindow(void* ptr, void* window) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	[nSWindowTabGroup addWindow:(NSWindow*)window];
}

void C_NSWindowTabGroup_InsertWindow_AtIndex(void* ptr, void* window, int index) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	[nSWindowTabGroup insertWindow:(NSWindow*)window atIndex:index];
}

void C_NSWindowTabGroup_RemoveWindow(void* ptr, void* window) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	[nSWindowTabGroup removeWindow:(NSWindow*)window];
}

void* C_NSWindowTabGroup_Identifier(void* ptr) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	NSString* result = [nSWindowTabGroup identifier];
	return result;
}

bool C_NSWindowTabGroup_IsOverviewVisible(void* ptr) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	bool result = [nSWindowTabGroup isOverviewVisible];
	return result;
}

void C_NSWindowTabGroup_SetOverviewVisible(void* ptr, bool value) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	[nSWindowTabGroup setOverviewVisible:value];
}

bool C_NSWindowTabGroup_IsTabBarVisible(void* ptr) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	bool result = [nSWindowTabGroup isTabBarVisible];
	return result;
}

void* C_NSWindowTabGroup_SelectedWindow(void* ptr) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	NSWindow* result = [nSWindowTabGroup selectedWindow];
	return result;
}

void C_NSWindowTabGroup_SetSelectedWindow(void* ptr, void* value) {
	NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
	[nSWindowTabGroup setSelectedWindow:(NSWindow*)value];
}
