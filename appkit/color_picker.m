#import <Appkit/Appkit.h>
#import "color_picker.h"

void* C_ColorPicker_Alloc() {
    return [NSColorPicker alloc];
}

void* C_NSColorPicker_InitWithPickerMask_ColorPanel(void* ptr, unsigned int mask, void* owningColorPanel) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSColorPicker* result_ = [nSColorPicker initWithPickerMask:mask colorPanel:(NSColorPanel*)owningColorPanel];
    return result_;
}

void* C_NSColorPicker_AllocColorPicker() {
    NSColorPicker* result_ = [NSColorPicker alloc];
    return result_;
}

void* C_NSColorPicker_Init(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSColorPicker* result_ = [nSColorPicker init];
    return result_;
}

void* C_NSColorPicker_NewColorPicker() {
    NSColorPicker* result_ = [NSColorPicker new];
    return result_;
}

void* C_NSColorPicker_Autorelease(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSColorPicker* result_ = [nSColorPicker autorelease];
    return result_;
}

void* C_NSColorPicker_Retain(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSColorPicker* result_ = [nSColorPicker retain];
    return result_;
}

void C_NSColorPicker_InsertNewButtonImage_In(void* ptr, void* newButtonImage, void* buttonCell) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    [nSColorPicker insertNewButtonImage:(NSImage*)newButtonImage in:(NSButtonCell*)buttonCell];
}

void C_NSColorPicker_SetMode(void* ptr, int mode) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    [nSColorPicker setMode:mode];
}

void C_NSColorPicker_AttachColorList(void* ptr, void* colorList) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    [nSColorPicker attachColorList:(NSColorList*)colorList];
}

void C_NSColorPicker_DetachColorList(void* ptr, void* colorList) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    [nSColorPicker detachColorList:(NSColorList*)colorList];
}

void C_NSColorPicker_ViewSizeChanged(void* ptr, void* sender) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    [nSColorPicker viewSizeChanged:(id)sender];
}

void* C_NSColorPicker_ColorPanel(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSColorPanel* result_ = [nSColorPicker colorPanel];
    return result_;
}

void* C_NSColorPicker_ProvideNewButtonImage(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSImage* result_ = [nSColorPicker provideNewButtonImage];
    return result_;
}

void* C_NSColorPicker_ButtonToolTip(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSString* result_ = [nSColorPicker buttonToolTip];
    return result_;
}

CGSize C_NSColorPicker_MinContentSize(void* ptr) {
    NSColorPicker* nSColorPicker = (NSColorPicker*)ptr;
    NSSize result_ = [nSColorPicker minContentSize];
    return result_;
}
