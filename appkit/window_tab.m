#import <Appkit/Appkit.h>
#import "window_tab.h"

void* C_WindowTab_Alloc() {
    return [NSWindowTab alloc];
}

void* C_NSWindowTab_Init(void* ptr) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    NSWindowTab* result_ = [nSWindowTab init];
    return result_;
}

void* C_NSWindowTab_Title(void* ptr) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    NSString* result_ = [nSWindowTab title];
    return result_;
}

void C_NSWindowTab_SetTitle(void* ptr, void* value) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    [nSWindowTab setTitle:(NSString*)value];
}

void* C_NSWindowTab_AttributedTitle(void* ptr) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    NSAttributedString* result_ = [nSWindowTab attributedTitle];
    return result_;
}

void C_NSWindowTab_SetAttributedTitle(void* ptr, void* value) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    [nSWindowTab setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSWindowTab_ToolTip(void* ptr) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    NSString* result_ = [nSWindowTab toolTip];
    return result_;
}

void C_NSWindowTab_SetToolTip(void* ptr, void* value) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    [nSWindowTab setToolTip:(NSString*)value];
}

void* C_NSWindowTab_AccessoryView(void* ptr) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    NSView* result_ = [nSWindowTab accessoryView];
    return result_;
}

void C_NSWindowTab_SetAccessoryView(void* ptr, void* value) {
    NSWindowTab* nSWindowTab = (NSWindowTab*)ptr;
    [nSWindowTab setAccessoryView:(NSView*)value];
}
