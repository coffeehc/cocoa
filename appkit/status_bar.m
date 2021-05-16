#import <Appkit/Appkit.h>
#import "status_bar.h"

void* C_StatusBar_Alloc() {
    return [NSStatusBar alloc];
}

void* C_NSStatusBar_Init(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusBar* result = [nSStatusBar init];
    return result;
}

void* C_NSStatusBar_StatusItemWithLength(void* ptr, double length) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusItem* result = [nSStatusBar statusItemWithLength:length];
    return result;
}

void C_NSStatusBar_RemoveStatusItem(void* ptr, void* item) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    [nSStatusBar removeStatusItem:(NSStatusItem*)item];
}

void* C_NSStatusBar_SystemStatusBar() {
    NSStatusBar* result = [NSStatusBar systemStatusBar];
    return result;
}

bool C_NSStatusBar_IsVertical(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    BOOL result = [nSStatusBar isVertical];
    return result;
}

double C_NSStatusBar_Thickness(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    CGFloat result = [nSStatusBar thickness];
    return result;
}
