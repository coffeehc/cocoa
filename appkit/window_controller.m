#import <Appkit/Appkit.h>
#import "window_controller.h"

void* C_WindowController_Alloc() {
	return [NSWindowController alloc];
}

void* C_NSWindowController_InitWithWindow(void* ptr, void* window) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController initWithWindow:(NSWindow*)window];
	return result;
}

void* C_NSWindowController_InitWithWindowNibName(void* ptr, void* windowNibName) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController initWithWindowNibName:(NSString*)windowNibName];
	return result;
}

void* C_NSWindowController_InitWithWindowNibName_Owner(void* ptr, void* windowNibName, void* owner) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController initWithWindowNibName:(NSString*)windowNibName owner:(id)owner];
	return result;
}

void* C_NSWindowController_InitWithWindowNibPath_Owner(void* ptr, void* windowNibPath, void* owner) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController initWithWindowNibPath:(NSString*)windowNibPath owner:(id)owner];
	return result;
}

void* C_NSWindowController_InitWithCoder(void* ptr, void* coder) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSWindowController_Init(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindowController* result = [nSWindowController init];
	return result;
}

void C_NSWindowController_LoadWindow(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController loadWindow];
}

void C_NSWindowController_WindowDidLoad(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController windowDidLoad];
}

void C_NSWindowController_WindowWillLoad(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController windowWillLoad];
}

void C_NSWindowController_SetDocumentEdited(void* ptr, bool dirtyFlag) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setDocumentEdited:dirtyFlag];
}

void C_NSWindowController_Close(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController close];
}

void C_NSWindowController_SynchronizeWindowTitleWithDocumentName(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController synchronizeWindowTitleWithDocumentName];
}

void* C_NSWindowController_WindowTitleForDocumentDisplayName(void* ptr, void* displayName) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSString* result = [nSWindowController windowTitleForDocumentDisplayName:(NSString*)displayName];
	return result;
}

bool C_NSWindowController_IsWindowLoaded(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	bool result = [nSWindowController isWindowLoaded];
	return result;
}

void* C_NSWindowController_Window(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSWindow* result = [nSWindowController window];
	return result;
}

void C_NSWindowController_SetWindow(void* ptr, void* value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setWindow:(NSWindow*)value];
}

void* C_NSWindowController_Document(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	id result = [nSWindowController document];
	return result;
}

void C_NSWindowController_SetDocument(void* ptr, void* value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setDocument:(id)value];
}

bool C_NSWindowController_ShouldCloseDocument(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	bool result = [nSWindowController shouldCloseDocument];
	return result;
}

void C_NSWindowController_SetShouldCloseDocument(void* ptr, bool value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setShouldCloseDocument:value];
}

void* C_NSWindowController_Owner(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	id result = [nSWindowController owner];
	return result;
}

void* C_NSWindowController_Storyboard(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSStoryboard* result = [nSWindowController storyboard];
	return result;
}

void* C_NSWindowController_WindowNibName(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSString* result = [nSWindowController windowNibName];
	return result;
}

void* C_NSWindowController_WindowNibPath(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSString* result = [nSWindowController windowNibPath];
	return result;
}

bool C_NSWindowController_ShouldCascadeWindows(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	bool result = [nSWindowController shouldCascadeWindows];
	return result;
}

void C_NSWindowController_SetShouldCascadeWindows(void* ptr, bool value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setShouldCascadeWindows:value];
}

void* C_NSWindowController_WindowFrameAutosaveName(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSString* result = [nSWindowController windowFrameAutosaveName];
	return result;
}

void C_NSWindowController_SetWindowFrameAutosaveName(void* ptr, void* value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setWindowFrameAutosaveName:(NSString*)value];
}

void* C_NSWindowController_ContentViewController(void* ptr) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	NSViewController* result = [nSWindowController contentViewController];
	return result;
}

void C_NSWindowController_SetContentViewController(void* ptr, void* value) {
	NSWindowController* nSWindowController = (NSWindowController*)ptr;
	[nSWindowController setContentViewController:(NSViewController*)value];
}
