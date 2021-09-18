#import "status_bar.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSStatusBar.h>

void* C_StatusBar_Alloc() {
    return [NSStatusBar alloc];
}

void* C_NSStatusBar_AllocStatusBar() {
    NSStatusBar* result_ = [NSStatusBar alloc];
    return result_;
}

void* C_NSStatusBar_Init(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusBar* result_ = [nSStatusBar init];
    return result_;
}

void* C_NSStatusBar_NewStatusBar() {
    NSStatusBar* result_ = [NSStatusBar new];
    return result_;
}

void* C_NSStatusBar_Autorelease(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusBar* result_ = [nSStatusBar autorelease];
    return result_;
}

void* C_NSStatusBar_Retain(void* ptr) {
    NSStatusBar* nSStatusBar = (NSStatusBar*)ptr;
    NSStatusBar* result_ = [nSStatusBar retain];
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
