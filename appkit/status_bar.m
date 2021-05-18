#import <Appkit/Appkit.h>
#import "status_bar.h"

void* C_StatusBar_Alloc() {
    return [NSStatusBar alloc];
}

void* C_NSStatusBar_Init(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusBar* result_ = [nSStatusBar init];
    return result_;
}

void* C_NSStatusBar_StatusItemWithLength(void* ptr, double length) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusItem* result_ = [nSStatusBar statusItemWithLength:length];
    return result_;
}

void C_NSStatusBar_RemoveStatusItem(void* ptr, void* item) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    [nSStatusBar removeStatusItem:(NSStatusItem*)item];
}

void* C_NSStatusBar_SystemStatusBar() {
    NSStatusBar* result_ = [NSStatusBar systemStatusBar];
    return result_;
}

bool C_NSStatusBar_IsVertical(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    BOOL result_ = [nSStatusBar isVertical];
    return result_;
}

double C_NSStatusBar_Thickness(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    CGFloat result_ = [nSStatusBar thickness];
    return result_;
}
