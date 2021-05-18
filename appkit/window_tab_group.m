#import <Appkit/Appkit.h>
#import "window_tab_group.h"

void* C_WindowTabGroup_Alloc() {
    return [NSWindowTabGroup alloc];
}

void* C_NSWindowTabGroup_Init(void* ptr) {
    NSWindowTabGroup* nSWindowTabGroup = (NSWindowTabGroup*)ptr;
    NSWindowTabGroup* result_ = [nSWindowTabGroup init];
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
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
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
