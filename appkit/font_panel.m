#import <Appkit/Appkit.h>
#import "font_panel.h"

void* C_FontPanel_Alloc() {
    return [NSFontPanel alloc];
}

void* C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFontPanel* result_ = [nSFontPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFontPanel* result_ = [nSFontPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSFontPanel_Init(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFontPanel* result_ = [nSFontPanel init];
    return result_;
}

void C_NSFontPanel_ReloadDefaultFontFamilies(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    [nSFontPanel reloadDefaultFontFamilies];
}

void C_NSFontPanel_SetPanelFont_IsMultiple(void* ptr, void* fontObj, bool flag) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    [nSFontPanel setPanelFont:(NSFont*)fontObj isMultiple:flag];
}

void* C_NSFontPanel_PanelConvertFont(void* ptr, void* fontObj) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFont* result_ = [nSFontPanel panelConvertFont:(NSFont*)fontObj];
    return result_;
}

void* C_NSFontPanel_SharedFontPanel() {
    NSFontPanel* result_ = [NSFontPanel sharedFontPanel];
    return result_;
}

bool C_NSFontPanel_SharedFontPanelExists() {
    BOOL result_ = [NSFontPanel sharedFontPanelExists];
    return result_;
}

bool C_NSFontPanel_IsEnabled(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    BOOL result_ = [nSFontPanel isEnabled];
    return result_;
}

void C_NSFontPanel_SetEnabled(void* ptr, bool value) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    [nSFontPanel setEnabled:value];
}

void* C_NSFontPanel_AccessoryView(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSView* result_ = [nSFontPanel accessoryView];
    return result_;
}

void C_NSFontPanel_SetAccessoryView(void* ptr, void* value) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    [nSFontPanel setAccessoryView:(NSView*)value];
}
