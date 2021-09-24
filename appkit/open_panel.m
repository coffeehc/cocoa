#import "open_panel.h"
#import <AppKit/NSOpenPanel.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_OpenPanel_Alloc() {
    return [NSOpenPanel alloc];
}

void* C_NSOpenPanel_OpenPanel_WindowWithContentViewController(void* contentViewController) {
    NSOpenPanel* result_ = [NSOpenPanel windowWithContentViewController:(NSViewController*)contentViewController];
    return result_;
}

void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result_ = [nSOpenPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result_ = [nSOpenPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSOpenPanel_Init(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result_ = [nSOpenPanel init];
    return result_;
}

void* C_NSOpenPanel_AllocOpenPanel() {
    NSOpenPanel* result_ = [NSOpenPanel alloc];
    return result_;
}

void* C_NSOpenPanel_NewOpenPanel() {
    NSOpenPanel* result_ = [NSOpenPanel new];
    return result_;
}

void* C_NSOpenPanel_Autorelease(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result_ = [nSOpenPanel autorelease];
    return result_;
}

void* C_NSOpenPanel_Retain(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result_ = [nSOpenPanel retain];
    return result_;
}

void* C_NSOpenPanel_OpenPanel_() {
    NSOpenPanel* result_ = [NSOpenPanel openPanel];
    return result_;
}

bool C_NSOpenPanel_CanChooseFiles(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel canChooseFiles];
    return result_;
}

void C_NSOpenPanel_SetCanChooseFiles(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanChooseFiles:value];
}

bool C_NSOpenPanel_CanChooseDirectories(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel canChooseDirectories];
    return result_;
}

void C_NSOpenPanel_SetCanChooseDirectories(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanChooseDirectories:value];
}

bool C_NSOpenPanel_ResolvesAliases(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel resolvesAliases];
    return result_;
}

void C_NSOpenPanel_SetResolvesAliases(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setResolvesAliases:value];
}

bool C_NSOpenPanel_AllowsMultipleSelection(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel allowsMultipleSelection];
    return result_;
}

void C_NSOpenPanel_SetAllowsMultipleSelection(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setAllowsMultipleSelection:value];
}

bool C_NSOpenPanel_IsAccessoryViewDisclosed(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel isAccessoryViewDisclosed];
    return result_;
}

void C_NSOpenPanel_SetAccessoryViewDisclosed(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setAccessoryViewDisclosed:value];
}

Array C_NSOpenPanel_URLs(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSArray* result_ = [nSOpenPanel URLs];
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

bool C_NSOpenPanel_CanDownloadUbiquitousContents(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel canDownloadUbiquitousContents];
    return result_;
}

void C_NSOpenPanel_SetCanDownloadUbiquitousContents(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanDownloadUbiquitousContents:value];
}

bool C_NSOpenPanel_CanResolveUbiquitousConflicts(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result_ = [nSOpenPanel canResolveUbiquitousConflicts];
    return result_;
}

void C_NSOpenPanel_SetCanResolveUbiquitousConflicts(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanResolveUbiquitousConflicts:value];
}
