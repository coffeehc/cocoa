#import "color_panel.h"
#import <AppKit/NSColorPanel.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ColorPanel_Alloc() {
    return [NSColorPanel alloc];
}

void* C_NSColorPanel_ColorPanel_WindowWithContentViewController(void* contentViewController) {
    NSColorPanel* result_ = [NSColorPanel windowWithContentViewController:(NSViewController*)contentViewController];
    return result_;
}

void* C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanel* result_ = [nSColorPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanel* result_ = [nSColorPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSColorPanel_Init(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanel* result_ = [nSColorPanel init];
    return result_;
}

void* C_NSColorPanel_AllocColorPanel() {
    NSColorPanel* result_ = [NSColorPanel alloc];
    return result_;
}

void* C_NSColorPanel_NewColorPanel() {
    NSColorPanel* result_ = [NSColorPanel new];
    return result_;
}

void* C_NSColorPanel_Autorelease(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanel* result_ = [nSColorPanel autorelease];
    return result_;
}

void* C_NSColorPanel_Retain(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanel* result_ = [nSColorPanel retain];
    return result_;
}

void C_NSColorPanel_ColorPanel_SetPickerMode(int mode) {
    [NSColorPanel setPickerMode:mode];
}

void C_NSColorPanel_ColorPanel_SetPickerMask(unsigned int mask) {
    [NSColorPanel setPickerMask:mask];
}

void C_NSColorPanel_SetAction(void* ptr, void* selector) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setAction:(SEL)selector];
}

void C_NSColorPanel_SetTarget(void* ptr, void* target) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setTarget:(id)target];
}

void C_NSColorPanel_AttachColorList(void* ptr, void* colorList) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel attachColorList:(NSColorList*)colorList];
}

void C_NSColorPanel_DetachColorList(void* ptr, void* colorList) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel detachColorList:(NSColorList*)colorList];
}

bool C_NSColorPanel_ColorPanel_DragColor_WithEvent_FromView(void* color, void* event, void* sourceView) {
    BOOL result_ = [NSColorPanel dragColor:(NSColor*)color withEvent:(NSEvent*)event fromView:(NSView*)sourceView];
    return result_;
}

void* C_NSColorPanel_SharedColorPanel() {
    NSColorPanel* result_ = [NSColorPanel sharedColorPanel];
    return result_;
}

bool C_NSColorPanel_SharedColorPanelExists() {
    BOOL result_ = [NSColorPanel sharedColorPanelExists];
    return result_;
}

int C_NSColorPanel_Mode(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColorPanelMode result_ = [nSColorPanel mode];
    return result_;
}

void C_NSColorPanel_SetMode(void* ptr, int value) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setMode:value];
}

void* C_NSColorPanel_AccessoryView(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSView* result_ = [nSColorPanel accessoryView];
    return result_;
}

void C_NSColorPanel_SetAccessoryView(void* ptr, void* value) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setAccessoryView:(NSView*)value];
}

bool C_NSColorPanel_IsContinuous(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    BOOL result_ = [nSColorPanel isContinuous];
    return result_;
}

void C_NSColorPanel_SetContinuous(void* ptr, bool value) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setContinuous:value];
}

bool C_NSColorPanel_ShowsAlpha(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    BOOL result_ = [nSColorPanel showsAlpha];
    return result_;
}

void C_NSColorPanel_SetShowsAlpha(void* ptr, bool value) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setShowsAlpha:value];
}

void* C_NSColorPanel_Color(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    NSColor* result_ = [nSColorPanel color];
    return result_;
}

void C_NSColorPanel_SetColor(void* ptr, void* value) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    [nSColorPanel setColor:(NSColor*)value];
}

double C_NSColorPanel_Alpha(void* ptr) {
    NSColorPanel* nSColorPanel = (NSColorPanel*)ptr;
    CGFloat result_ = [nSColorPanel alpha];
    return result_;
}
