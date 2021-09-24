#import "status_item.h"
#import <AppKit/NSStatusItem.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_StatusItem_Alloc() {
    return [NSStatusItem alloc];
}

void* C_NSStatusItem_AllocStatusItem() {
    NSStatusItem* result_ = [NSStatusItem alloc];
    return result_;
}

void* C_NSStatusItem_Init(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItem* result_ = [nSStatusItem init];
    return result_;
}

void* C_NSStatusItem_NewStatusItem() {
    NSStatusItem* result_ = [NSStatusItem new];
    return result_;
}

void* C_NSStatusItem_Autorelease(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItem* result_ = [nSStatusItem autorelease];
    return result_;
}

void* C_NSStatusItem_Retain(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItem* result_ = [nSStatusItem retain];
    return result_;
}

void* C_NSStatusItem_StatusBar(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusBar* result_ = [nSStatusItem statusBar];
    return result_;
}

unsigned int C_NSStatusItem_Behavior(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItemBehavior result_ = [nSStatusItem behavior];
    return result_;
}

void C_NSStatusItem_SetBehavior(void* ptr, unsigned int value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setBehavior:value];
}

void* C_NSStatusItem_Button(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusBarButton* result_ = [nSStatusItem button];
    return result_;
}

void* C_NSStatusItem_Menu(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSMenu* result_ = [nSStatusItem menu];
    return result_;
}

void C_NSStatusItem_SetMenu(void* ptr, void* value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setMenu:(NSMenu*)value];
}

bool C_NSStatusItem_IsVisible(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    BOOL result_ = [nSStatusItem isVisible];
    return result_;
}

void C_NSStatusItem_SetVisible(void* ptr, bool value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setVisible:value];
}

double C_NSStatusItem_Length(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    CGFloat result_ = [nSStatusItem length];
    return result_;
}

void C_NSStatusItem_SetLength(void* ptr, double value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setLength:value];
}

void* C_NSStatusItem_AutosaveName(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItemAutosaveName result_ = [nSStatusItem autosaveName];
    return result_;
}

void C_NSStatusItem_SetAutosaveName(void* ptr, void* value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setAutosaveName:(NSString*)value];
}
