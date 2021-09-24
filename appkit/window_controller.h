#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_WindowController_Alloc();

void* C_NSWindowController_InitWithWindow(void* ptr, void* window);
void* C_NSWindowController_InitWithWindowNibName(void* ptr, void* windowNibName);
void* C_NSWindowController_InitWithWindowNibName_Owner(void* ptr, void* windowNibName, void* owner);
void* C_NSWindowController_InitWithWindowNibPath_Owner(void* ptr, void* windowNibPath, void* owner);
void* C_NSWindowController_InitWithCoder(void* ptr, void* coder);
void* C_NSWindowController_Init(void* ptr);
void* C_NSWindowController_AllocWindowController();
void* C_NSWindowController_NewWindowController();
void* C_NSWindowController_Autorelease(void* ptr);
void* C_NSWindowController_Retain(void* ptr);
void C_NSWindowController_LoadWindow(void* ptr);
void C_NSWindowController_ShowWindow(void* ptr, void* sender);
void C_NSWindowController_WindowDidLoad(void* ptr);
void C_NSWindowController_WindowWillLoad(void* ptr);
void C_NSWindowController_SetDocumentEdited(void* ptr, bool dirtyFlag);
void C_NSWindowController_Close(void* ptr);
void C_NSWindowController_SynchronizeWindowTitleWithDocumentName(void* ptr);
void* C_NSWindowController_WindowTitleForDocumentDisplayName(void* ptr, void* displayName);
void C_NSWindowController_DismissController(void* ptr, void* sender);
bool C_NSWindowController_IsWindowLoaded(void* ptr);
void* C_NSWindowController_Window(void* ptr);
void C_NSWindowController_SetWindow(void* ptr, void* value);
void* C_NSWindowController_Document(void* ptr);
void C_NSWindowController_SetDocument(void* ptr, void* value);
bool C_NSWindowController_ShouldCloseDocument(void* ptr);
void C_NSWindowController_SetShouldCloseDocument(void* ptr, bool value);
void* C_NSWindowController_Owner(void* ptr);
void* C_NSWindowController_Storyboard(void* ptr);
void* C_NSWindowController_WindowNibName(void* ptr);
void* C_NSWindowController_WindowNibPath(void* ptr);
bool C_NSWindowController_ShouldCascadeWindows(void* ptr);
void C_NSWindowController_SetShouldCascadeWindows(void* ptr, bool value);
void* C_NSWindowController_WindowFrameAutosaveName(void* ptr);
void C_NSWindowController_SetWindowFrameAutosaveName(void* ptr, void* value);
void* C_NSWindowController_ContentViewController(void* ptr);
void C_NSWindowController_SetContentViewController(void* ptr, void* value);
