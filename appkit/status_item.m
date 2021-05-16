#import <Appkit/Appkit.h>
#import "status_item.h"

void* C_StatusItem_Alloc() {
    return [NSStatusItem alloc];
}

void* C_NSStatusItem_Init(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItem* result = [nSStatusItem init];
    return result;
}

void* C_NSStatusItem_StatusBar(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusBar* result = [nSStatusItem statusBar];
    return result;
}

unsigned int C_NSStatusItem_Behavior(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItemBehavior result = [nSStatusItem behavior];
    return result;
}

void C_NSStatusItem_SetBehavior(void* ptr, unsigned int value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setBehavior:value];
}

void* C_NSStatusItem_Button(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusBarButton* result = [nSStatusItem button];
    return result;
}

void* C_NSStatusItem_Menu(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSMenu* result = [nSStatusItem menu];
    return result;
}

void C_NSStatusItem_SetMenu(void* ptr, void* value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setMenu:(NSMenu*)value];
}

bool C_NSStatusItem_IsVisible(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    BOOL result = [nSStatusItem isVisible];
    return result;
}

void C_NSStatusItem_SetVisible(void* ptr, bool value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setVisible:value];
}

double C_NSStatusItem_Length(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    CGFloat result = [nSStatusItem length];
    return result;
}

void C_NSStatusItem_SetLength(void* ptr, double value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setLength:value];
}

void* C_NSStatusItem_AutosaveName(void* ptr) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    NSStatusItemAutosaveName result = [nSStatusItem autosaveName];
    return result;
}

void C_NSStatusItem_SetAutosaveName(void* ptr, void* value) {
    NSStatusItem* nSStatusItem = (NSStatusItem*)ptr;
    [nSStatusItem setAutosaveName:(NSString*)value];
}
