#import "button.h"
#import <AppKit/NSButton.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Button_Alloc() {
    return [NSButton alloc];
}

void* C_NSButton_Button_CheckboxWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result_ = [NSButton checkboxWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSButton_ButtonWithImage_Target_Action(void* image, void* target, void* action) {
    NSButton* result_ = [NSButton buttonWithImage:(NSImage*)image target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSButton_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result_ = [NSButton radioButtonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action) {
    NSButton* result_ = [NSButton buttonWithTitle:(NSString*)title image:(NSImage*)image target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSButton* result_ = [NSButton buttonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSButton_InitWithFrame(void* ptr, CGRect frameRect) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result_ = [nSButton initWithFrame:frameRect];
    return result_;
}

void* C_NSButton_InitWithCoder(void* ptr, void* coder) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result_ = [nSButton initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSButton_Init(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result_ = [nSButton init];
    return result_;
}

void* C_NSButton_AllocButton() {
    NSButton* result_ = [NSButton alloc];
    return result_;
}

void* C_NSButton_NewButton() {
    NSButton* result_ = [NSButton new];
    return result_;
}

void* C_NSButton_Autorelease(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result_ = [nSButton autorelease];
    return result_;
}

void* C_NSButton_Retain(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSButton* result_ = [nSButton retain];
    return result_;
}

void C_NSButton_SetButtonType(void* ptr, unsigned int _type) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setButtonType:_type];
}

void C_NSButton_SetPeriodicDelay_Interval(void* ptr, float delay, float interval) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setPeriodicDelay:delay interval:interval];
}

void C_NSButton_CompressWithPrioritizedCompressionOptions(void* ptr, Array prioritizedOptions) {
    NSButton* nSButton = (NSButton*)ptr;
    NSMutableArray* objcPrioritizedOptions = [NSMutableArray arrayWithCapacity:prioritizedOptions.len];
    if (prioritizedOptions.len > 0) {
    	void** prioritizedOptionsData = (void**)prioritizedOptions.data;
    	for (int i = 0; i < prioritizedOptions.len; i++) {
    		void* p = prioritizedOptionsData[i];
    		[objcPrioritizedOptions addObject:(NSUserInterfaceCompressionOptions*)p];
    	}
    }
    [nSButton compressWithPrioritizedCompressionOptions:objcPrioritizedOptions];
}

CGSize C_NSButton_MinimumSizeWithPrioritizedCompressionOptions(void* ptr, Array prioritizedOptions) {
    NSButton* nSButton = (NSButton*)ptr;
    NSMutableArray* objcPrioritizedOptions = [NSMutableArray arrayWithCapacity:prioritizedOptions.len];
    if (prioritizedOptions.len > 0) {
    	void** prioritizedOptionsData = (void**)prioritizedOptions.data;
    	for (int i = 0; i < prioritizedOptions.len; i++) {
    		void* p = prioritizedOptionsData[i];
    		[objcPrioritizedOptions addObject:(NSUserInterfaceCompressionOptions*)p];
    	}
    }
    NSSize result_ = [nSButton minimumSizeWithPrioritizedCompressionOptions:objcPrioritizedOptions];
    return result_;
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
    NSColor* result_ = [nSButton contentTintColor];
    return result_;
}

void C_NSButton_SetContentTintColor(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setContentTintColor:(NSColor*)value];
}

bool C_NSButton_HasDestructiveAction(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton hasDestructiveAction];
    return result_;
}

void C_NSButton_SetHasDestructiveAction(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setHasDestructiveAction:value];
}

void* C_NSButton_AlternateTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result_ = [nSButton alternateTitle];
    return result_;
}

void C_NSButton_SetAlternateTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAlternateTitle:(NSString*)value];
}

void* C_NSButton_AttributedTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSAttributedString* result_ = [nSButton attributedTitle];
    return result_;
}

void C_NSButton_SetAttributedTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSButton_AttributedAlternateTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSAttributedString* result_ = [nSButton attributedAlternateTitle];
    return result_;
}

void C_NSButton_SetAttributedAlternateTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAttributedAlternateTitle:(NSAttributedString*)value];
}

void* C_NSButton_Title(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result_ = [nSButton title];
    return result_;
}

void C_NSButton_SetTitle(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setTitle:(NSString*)value];
}

void* C_NSButton_Sound(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSSound* result_ = [nSButton sound];
    return result_;
}

void C_NSButton_SetSound(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSound:(NSSound*)value];
}

bool C_NSButton_IsSpringLoaded(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton isSpringLoaded];
    return result_;
}

void C_NSButton_SetSpringLoaded(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSpringLoaded:value];
}

int C_NSButton_MaxAcceleratorLevel(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSInteger result_ = [nSButton maxAcceleratorLevel];
    return result_;
}

void C_NSButton_SetMaxAcceleratorLevel(void* ptr, int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setMaxAcceleratorLevel:value];
}

void* C_NSButton_Image(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImage* result_ = [nSButton image];
    return result_;
}

void C_NSButton_SetImage(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImage:(NSImage*)value];
}

void* C_NSButton_AlternateImage(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImage* result_ = [nSButton alternateImage];
    return result_;
}

void C_NSButton_SetAlternateImage(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAlternateImage:(NSImage*)value];
}

unsigned int C_NSButton_ImagePosition(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSCellImagePosition result_ = [nSButton imagePosition];
    return result_;
}

void C_NSButton_SetImagePosition(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImagePosition:value];
}

bool C_NSButton_IsBordered(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton isBordered];
    return result_;
}

void C_NSButton_SetBordered(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBordered:value];
}

bool C_NSButton_IsTransparent(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton isTransparent];
    return result_;
}

void C_NSButton_SetTransparent(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setTransparent:value];
}

unsigned int C_NSButton_BezelStyle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSBezelStyle result_ = [nSButton bezelStyle];
    return result_;
}

void C_NSButton_SetBezelStyle(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBezelStyle:value];
}

void* C_NSButton_BezelColor(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSColor* result_ = [nSButton bezelColor];
    return result_;
}

void C_NSButton_SetBezelColor(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setBezelColor:(NSColor*)value];
}

bool C_NSButton_ShowsBorderOnlyWhileMouseInside(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton showsBorderOnlyWhileMouseInside];
    return result_;
}

void C_NSButton_SetShowsBorderOnlyWhileMouseInside(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setShowsBorderOnlyWhileMouseInside:value];
}

bool C_NSButton_ImageHugsTitle(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton imageHugsTitle];
    return result_;
}

void C_NSButton_SetImageHugsTitle(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImageHugsTitle:value];
}

unsigned int C_NSButton_ImageScaling(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImageScaling result_ = [nSButton imageScaling];
    return result_;
}

void C_NSButton_SetImageScaling(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setImageScaling:value];
}

void* C_NSButton_ActiveCompressionOptions(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSButton activeCompressionOptions];
    return result_;
}

bool C_NSButton_AllowsMixedState(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    BOOL result_ = [nSButton allowsMixedState];
    return result_;
}

void C_NSButton_SetAllowsMixedState(void* ptr, bool value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setAllowsMixedState:value];
}

int C_NSButton_State(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSControlStateValue result_ = [nSButton state];
    return result_;
}

void C_NSButton_SetState(void* ptr, int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setState:value];
}

void* C_NSButton_KeyEquivalent(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSString* result_ = [nSButton keyEquivalent];
    return result_;
}

void C_NSButton_SetKeyEquivalent(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setKeyEquivalent:(NSString*)value];
}

unsigned int C_NSButton_KeyEquivalentModifierMask(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSEventModifierFlags result_ = [nSButton keyEquivalentModifierMask];
    return result_;
}

void C_NSButton_SetKeyEquivalentModifierMask(void* ptr, unsigned int value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setKeyEquivalentModifierMask:value];
}

void* C_NSButton_SymbolConfiguration(void* ptr) {
    NSButton* nSButton = (NSButton*)ptr;
    NSImageSymbolConfiguration* result_ = [nSButton symbolConfiguration];
    return result_;
}

void C_NSButton_SetSymbolConfiguration(void* ptr, void* value) {
    NSButton* nSButton = (NSButton*)ptr;
    [nSButton setSymbolConfiguration:(NSImageSymbolConfiguration*)value];
}
