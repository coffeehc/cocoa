#import "alert.h"
#import <AppKit/NSAlert.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_Alert_Alloc() {
    return [NSAlert alloc];
}

void* C_NSAlert_AllocAlert() {
    NSAlert* result_ = [NSAlert alloc];
    return result_;
}

void* C_NSAlert_Init(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSAlert* result_ = [nSAlert init];
    return result_;
}

void* C_NSAlert_NewAlert() {
    NSAlert* result_ = [NSAlert new];
    return result_;
}

void* C_NSAlert_Autorelease(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSAlert* result_ = [nSAlert autorelease];
    return result_;
}

void* C_NSAlert_Retain(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSAlert* result_ = [nSAlert retain];
    return result_;
}

void* C_NSAlert_AlertWithError(void* error) {
    NSAlert* result_ = [NSAlert alertWithError:(NSError*)error];
    return result_;
}

void C_NSAlert_Layout(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert layout];
}

int C_NSAlert_RunModal(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSModalResponse result_ = [nSAlert runModal];
    return result_;
}

void* C_NSAlert_AddButtonWithTitle(void* ptr, void* title) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSButton* result_ = [nSAlert addButtonWithTitle:(NSString*)title];
    return result_;
}

unsigned int C_NSAlert_AlertStyle(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSAlertStyle result_ = [nSAlert alertStyle];
    return result_;
}

void C_NSAlert_SetAlertStyle(void* ptr, unsigned int value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setAlertStyle:value];
}

void* C_NSAlert_AccessoryView(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSView* result_ = [nSAlert accessoryView];
    return result_;
}

void C_NSAlert_SetAccessoryView(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setAccessoryView:(NSView*)value];
}

bool C_NSAlert_ShowsHelp(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    BOOL result_ = [nSAlert showsHelp];
    return result_;
}

void C_NSAlert_SetShowsHelp(void* ptr, bool value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setShowsHelp:value];
}

void* C_NSAlert_HelpAnchor(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSHelpAnchorName result_ = [nSAlert helpAnchor];
    return result_;
}

void C_NSAlert_SetHelpAnchor(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setHelpAnchor:(NSString*)value];
}

void* C_NSAlert_Delegate(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    id result_ = [nSAlert delegate];
    return result_;
}

void C_NSAlert_SetDelegate(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setDelegate:(id)value];
}

void* C_NSAlert_SuppressionButton(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSButton* result_ = [nSAlert suppressionButton];
    return result_;
}

bool C_NSAlert_ShowsSuppressionButton(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    BOOL result_ = [nSAlert showsSuppressionButton];
    return result_;
}

void C_NSAlert_SetShowsSuppressionButton(void* ptr, bool value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setShowsSuppressionButton:value];
}

void* C_NSAlert_InformativeText(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSString* result_ = [nSAlert informativeText];
    return result_;
}

void C_NSAlert_SetInformativeText(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setInformativeText:(NSString*)value];
}

void* C_NSAlert_MessageText(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSString* result_ = [nSAlert messageText];
    return result_;
}

void C_NSAlert_SetMessageText(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setMessageText:(NSString*)value];
}

void* C_NSAlert_Icon(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSImage* result_ = [nSAlert icon];
    return result_;
}

void C_NSAlert_SetIcon(void* ptr, void* value) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    [nSAlert setIcon:(NSImage*)value];
}

Array C_NSAlert_Buttons(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSArray* result_ = [nSAlert buttons];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSAlert_Window(void* ptr) {
    NSAlert* nSAlert = (NSAlert*)ptr;
    NSWindow* result_ = [nSAlert window];
    return result_;
}

@interface NSAlertDelegateAdaptor : NSObject <NSAlertDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSAlertDelegateAdaptor


- (BOOL)alertShowHelp:(NSAlert*)alert {
    bool result_ = alertDelegate_AlertShowHelp([self goID], alert);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return AlertDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapAlertDelegate(uintptr_t goID) {
    NSAlertDelegateAdaptor* adaptor = [[NSAlertDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
