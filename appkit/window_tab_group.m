#import "window_tab_group.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSWindowTabGroup.h>

void* C_WindowTabGroup_Alloc() {
    return [NSWindowTabGroup alloc];
}

void* C_NSWindowTabGroup_AllocWindowTabGroup() {
    NSWindowTabGroup* result_ = [NSWindowTabGroup alloc];
    return result_;
}

void* C_NSWindowTabGroup_Init(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSWindowTabGroup* result_ = [nSWindowTabGroup init];
    return result_;
}

void* C_NSWindowTabGroup_NewWindowTabGroup() {
    NSWindowTabGroup* result_ = [NSWindowTabGroup new];
    return result_;
}

void* C_NSWindowTabGroup_Autorelease(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSWindowTabGroup* result_ = [nSWindowTabGroup autorelease];
    return result_;
}

void* C_NSWindowTabGroup_Retain(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSWindowTabGroup* result_ = [nSWindowTabGroup retain];
    return result_;
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
    NSWindowTabbingIdentifier result_ = [nSWindowTabGroup identifier];
    return result_;
}

bool C_NSWindowTabGroup_IsOverviewVisible(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    BOOL result_ = [nSWindowTabGroup isOverviewVisible];
    return result_;
}

void C_NSWindowTabGroup_SetOverviewVisible(void* ptr, bool value) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    [nSWindowTabGroup setOverviewVisible:value];
}

bool C_NSWindowTabGroup_IsTabBarVisible(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    BOOL result_ = [nSWindowTabGroup isTabBarVisible];
    return result_;
}

Array C_NSWindowTabGroup_Windows(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSArray* result_ = [nSWindowTabGroup windows];
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

void* C_NSWindowTabGroup_SelectedWindow(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSWindow* result_ = [nSWindowTabGroup selectedWindow];
    return result_;
}

void C_NSWindowTabGroup_SetSelectedWindow(void* ptr, void* value) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    [nSWindowTabGroup setSelectedWindow:(NSWindow*)value];
}
