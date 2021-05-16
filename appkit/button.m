#import <Appkit/Appkit.h>
#import "button.h"

void* C_Button_Alloc() {
    return [NSButton alloc];
}

void* C_NSButton_InitWithFrame(void* ptr, CGRect frameRect) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result = [nSButton initWithFrame:frameRect];
    return result;
}

void* C_NSButton_InitWithCoder(void* ptr, void* coder) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result = [nSButton initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSButton_Init(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result = [nSButton init];
    return result;
}

void* C_NSButton_Button_CheckboxWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result = [NSButton checkboxWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result;
}

void* C_NSButton_ButtonWithImage_Target_Action(void* image, void* target, void* action) {
    NSButton* result = [NSButton buttonWithImage:(NSImage*)image target:(id)target action:(SEL)action];
    return result;
}

void* C_NSButton_Button_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result = [NSButton radioButtonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result;
}

void* C_NSButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action) {
    NSButton* result = [NSButton buttonWithTitle:(NSString*)title image:(NSImage*)image target:(id)target action:(SEL)action];
    return result;
}

void* C_NSButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result = [NSButton buttonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result;
}

void C_NSButton_SetButtonType(void* ptr, unsigned int _type) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setButtonType:_type];
}

void C_NSButton_SetPeriodicDelay_Interval(void* ptr, float delay, float interval) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setPeriodicDelay:delay interval:interval];
}

void C_NSButton_SetNextState(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setNextState];
}

void C_NSButton_Highlight(void* ptr, bool flag) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton highlight:flag];
}

void* C_NSButton_ContentTintColor(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSColor* result = [nSButton contentTintColor];
    return result;
}

void C_NSButton_SetContentTintColor(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setContentTintColor:(NSColor*)value];
}

bool C_NSButton_HasDestructiveAction(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton hasDestructiveAction];
    return result;
}

void C_NSButton_SetHasDestructiveAction(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setHasDestructiveAction:value];
}

void* C_NSButton_AlternateTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result = [nSButton alternateTitle];
    return result;
}

void C_NSButton_SetAlternateTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAlternateTitle:(NSString*)value];
}

void* C_NSButton_AttributedTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSAttributedString* result = [nSButton attributedTitle];
    return result;
}

void C_NSButton_SetAttributedTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSButton_AttributedAlternateTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSAttributedString* result = [nSButton attributedAlternateTitle];
    return result;
}

void C_NSButton_SetAttributedAlternateTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAttributedAlternateTitle:(NSAttributedString*)value];
}

void* C_NSButton_Title(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result = [nSButton title];
    return result;
}

void C_NSButton_SetTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setTitle:(NSString*)value];
}

void* C_NSButton_Sound(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSSound* result = [nSButton sound];
    return result;
}

void C_NSButton_SetSound(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSound:(NSSound*)value];
}

bool C_NSButton_IsSpringLoaded(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton isSpringLoaded];
    return result;
}

void C_NSButton_SetSpringLoaded(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSpringLoaded:value];
}

int C_NSButton_MaxAcceleratorLevel(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSInteger result = [nSButton maxAcceleratorLevel];
    return result;
}

void C_NSButton_SetMaxAcceleratorLevel(void* ptr, int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setMaxAcceleratorLevel:value];
}

void* C_NSButton_Image(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImage* result = [nSButton image];
    return result;
}

void C_NSButton_SetImage(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImage:(NSImage*)value];
}

void* C_NSButton_AlternateImage(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImage* result = [nSButton alternateImage];
    return result;
}

void C_NSButton_SetAlternateImage(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAlternateImage:(NSImage*)value];
}

unsigned int C_NSButton_ImagePosition(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSCellImagePosition result = [nSButton imagePosition];
    return result;
}

void C_NSButton_SetImagePosition(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImagePosition:value];
}

bool C_NSButton_IsBordered(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton isBordered];
    return result;
}

void C_NSButton_SetBordered(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBordered:value];
}

bool C_NSButton_IsTransparent(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton isTransparent];
    return result;
}

void C_NSButton_SetTransparent(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setTransparent:value];
}

unsigned int C_NSButton_BezelStyle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSBezelStyle result = [nSButton bezelStyle];
    return result;
}

void C_NSButton_SetBezelStyle(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBezelStyle:value];
}

void* C_NSButton_BezelColor(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSColor* result = [nSButton bezelColor];
    return result;
}

void C_NSButton_SetBezelColor(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBezelColor:(NSColor*)value];
}

bool C_NSButton_ShowsBorderOnlyWhileMouseInside(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton showsBorderOnlyWhileMouseInside];
    return result;
}

void C_NSButton_SetShowsBorderOnlyWhileMouseInside(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setShowsBorderOnlyWhileMouseInside:value];
}

bool C_NSButton_ImageHugsTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton imageHugsTitle];
    return result;
}

void C_NSButton_SetImageHugsTitle(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImageHugsTitle:value];
}

unsigned int C_NSButton_ImageScaling(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImageScaling result = [nSButton imageScaling];
    return result;
}

void C_NSButton_SetImageScaling(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImageScaling:value];
}

void* C_NSButton_ActiveCompressionOptions(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSButton activeCompressionOptions];
    return result;
}

bool C_NSButton_AllowsMixedState(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result = [nSButton allowsMixedState];
    return result;
}

void C_NSButton_SetAllowsMixedState(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAllowsMixedState:value];
}

int C_NSButton_State(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSControlStateValue result = [nSButton state];
    return result;
}

void C_NSButton_SetState(void* ptr, int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setState:value];
}

void* C_NSButton_KeyEquivalent(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result = [nSButton keyEquivalent];
    return result;
}

void C_NSButton_SetKeyEquivalent(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setKeyEquivalent:(NSString*)value];
}

unsigned int C_NSButton_KeyEquivalentModifierMask(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSEventModifierFlags result = [nSButton keyEquivalentModifierMask];
    return result;
}

void C_NSButton_SetKeyEquivalentModifierMask(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setKeyEquivalentModifierMask:value];
}

void* C_NSButton_SymbolConfiguration(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImageSymbolConfiguration* result = [nSButton symbolConfiguration];
    return result;
}

void C_NSButton_SetSymbolConfiguration(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSymbolConfiguration:(NSImageSymbolConfiguration*)value];
}
