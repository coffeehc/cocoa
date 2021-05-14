#import <Appkit/Appkit.h>
#import "window_tab.h"

void* C_WindowTab_Alloc() {
	return [NSWindowTab alloc];
}

void* C_NSWindowTab_Init(void* ptr) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	NSWindowTab* result = [nSWindowTab init];
	return result;
}

void* C_NSWindowTab_Title(void* ptr) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	NSString* result = [nSWindowTab title];
	return result;
}

void C_NSWindowTab_SetTitle(void* ptr, void* value) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	[nSWindowTab setTitle:(NSString*)value];
}

void* C_NSWindowTab_AttributedTitle(void* ptr) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	NSAttributedString* result = [nSWindowTab attributedTitle];
	return result;
}

void C_NSWindowTab_SetAttributedTitle(void* ptr, void* value) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	[nSWindowTab setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSWindowTab_ToolTip(void* ptr) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	NSString* result = [nSWindowTab toolTip];
	return result;
}

void C_NSWindowTab_SetToolTip(void* ptr, void* value) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	[nSWindowTab setToolTip:(NSString*)value];
}

void* C_NSWindowTab_AccessoryView(void* ptr) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	NSView* result = [nSWindowTab accessoryView];
	return result;
}

void C_NSWindowTab_SetAccessoryView(void* ptr, void* value) {
	NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
	[nSWindowTab setAccessoryView:(NSView*)value];
}
