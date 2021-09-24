#import "font_panel.h"
#import <AppKit/NSFontPanel.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_FontPanel_Alloc() {
    return [NSFontPanel alloc];
}

void* C_NSFontPanel_FontPanel_WindowWithContentViewController(void* contentViewController) {
    NSFontPanel* result_ = [NSFontPanel windowWithContentViewController:(NSViewController*)contentViewController];
    return result_;
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

void* C_NSFontPanel_AllocFontPanel() {
    NSFontPanel* result_ = [NSFontPanel alloc];
    return result_;
}

void* C_NSFontPanel_NewFontPanel() {
    NSFontPanel* result_ = [NSFontPanel new];
    return result_;
}

void* C_NSFontPanel_Autorelease(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFontPanel* result_ = [nSFontPanel autorelease];
    return result_;
}

void* C_NSFontPanel_Retain(void* ptr) {
    NSFontPanel* nSFontPanel = (NSFontPanel*)ptr;
    NSFontPanel* result_ = [nSFontPanel retain];
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

@interface NSFontChangingAdaptor : NSObject <NSFontChanging>
@property (assign) uintptr_t goID;
@end

@implementation NSFontChangingAdaptor


- (void)changeFont:(NSFontManager*)sender {
    fontChanging_ChangeFont([self goID], sender);
}

- (NSFontPanelModeMask)validModesForFontPanel:(NSFontPanel*)fontPanel {
    unsigned int result_ = fontChanging_ValidModesForFontPanel([self goID], fontPanel);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return FontChanging_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapFontChanging(uintptr_t goID) {
    NSFontChangingAdaptor* adaptor = [[NSFontChangingAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
