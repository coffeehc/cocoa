#import <Appkit/Appkit.h>
#import "window_controller.h"

void* C_WindowController_Alloc() {
    return [NSWindowController alloc];
}

void* C_NSWindowController_InitWithWindow(void* ptr, void* window) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController initWithWindow:(NSWindow*)window];
    return result_;
}

void* C_NSWindowController_InitWithWindowNibName(void* ptr, void* windowNibName) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController initWithWindowNibName:(NSString*)windowNibName];
    return result_;
}

void* C_NSWindowController_InitWithWindowNibName_Owner(void* ptr, void* windowNibName, void* owner) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController initWithWindowNibName:(NSString*)windowNibName owner:(id)owner];
    return result_;
}

void* C_NSWindowController_InitWithWindowNibPath_Owner(void* ptr, void* windowNibPath, void* owner) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController initWithWindowNibPath:(NSString*)windowNibPath owner:(id)owner];
    return result_;
}

void* C_NSWindowController_InitWithCoder(void* ptr, void* coder) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSWindowController_Init(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController init];
    return result_;
}

void* C_NSWindowController_AllocWindowController() {
    NSWindowController* result_ = [NSWindowController alloc];
    return result_;
}

void* C_NSWindowController_NewWindowController() {
    NSWindowController* result_ = [NSWindowController new];
    return result_;
}

void* C_NSWindowController_Autorelease(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController autorelease];
    return result_;
}

void* C_NSWindowController_Retain(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowController* result_ = [nSWindowController retain];
    return result_;
}

void C_NSWindowController_LoadWindow(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController loadWindow];
}

void C_NSWindowController_ShowWindow(void* ptr, void* sender) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController showWindow:(id)sender];
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
    NSString* result_ = [nSWindowController windowTitleForDocumentDisplayName:(NSString*)displayName];
    return result_;
}

void C_NSWindowController_DismissController(void* ptr, void* sender) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController dismissController:(id)sender];
}

bool C_NSWindowController_IsWindowLoaded(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    BOOL result_ = [nSWindowController isWindowLoaded];
    return result_;
}

void* C_NSWindowController_Window(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindow* result_ = [nSWindowController window];
    return result_;
}

void C_NSWindowController_SetWindow(void* ptr, void* value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setWindow:(NSWindow*)value];
}

void* C_NSWindowController_Document(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    id result_ = [nSWindowController document];
    return result_;
}

void C_NSWindowController_SetDocument(void* ptr, void* value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setDocument:(id)value];
}

bool C_NSWindowController_ShouldCloseDocument(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    BOOL result_ = [nSWindowController shouldCloseDocument];
    return result_;
}

void C_NSWindowController_SetShouldCloseDocument(void* ptr, bool value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setShouldCloseDocument:value];
}

void* C_NSWindowController_Owner(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    id result_ = [nSWindowController owner];
    return result_;
}

void* C_NSWindowController_Storyboard(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSStoryboard* result_ = [nSWindowController storyboard];
    return result_;
}

void* C_NSWindowController_WindowNibName(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSNibName result_ = [nSWindowController windowNibName];
    return result_;
}

void* C_NSWindowController_WindowNibPath(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSString* result_ = [nSWindowController windowNibPath];
    return result_;
}

bool C_NSWindowController_ShouldCascadeWindows(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    BOOL result_ = [nSWindowController shouldCascadeWindows];
    return result_;
}

void C_NSWindowController_SetShouldCascadeWindows(void* ptr, bool value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setShouldCascadeWindows:value];
}

void* C_NSWindowController_WindowFrameAutosaveName(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSWindowFrameAutosaveName result_ = [nSWindowController windowFrameAutosaveName];
    return result_;
}

void C_NSWindowController_SetWindowFrameAutosaveName(void* ptr, void* value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setWindowFrameAutosaveName:(NSString*)value];
}

void* C_NSWindowController_ContentViewController(void* ptr) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    NSViewController* result_ = [nSWindowController contentViewController];
    return result_;
}

void C_NSWindowController_SetContentViewController(void* ptr, void* value) {
    NSWindowController* nSWindowController = (NSWindowController*)ptr;
    [nSWindowController setContentViewController:(NSViewController*)value];
}
