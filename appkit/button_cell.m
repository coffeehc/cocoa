#import "button_cell.h"
#import <AppKit/NSButtonCell.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ButtonCell_Alloc() {
    return [NSButtonCell alloc];
}

void* C_NSButtonCell_InitWithCoder(void* ptr, void* coder) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSButtonCell_InitImageCell(void* ptr, void* image) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell initImageCell:(NSImage*)image];
    return result_;
}

void* C_NSButtonCell_InitTextCell(void* ptr, void* _string) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSButtonCell_Init(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell init];
    return result_;
}

void* C_NSButtonCell_AllocButtonCell() {
    NSButtonCell* result_ = [NSButtonCell alloc];
    return result_;
}

void* C_NSButtonCell_NewButtonCell() {
    NSButtonCell* result_ = [NSButtonCell new];
    return result_;
}

void* C_NSButtonCell_Autorelease(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell autorelease];
    return result_;
}

void* C_NSButtonCell_Retain(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSButtonCell* result_ = [nSButtonCell retain];
    return result_;
}

void C_NSButtonCell_SetPeriodicDelay_Interval(void* ptr, float delay, float interval) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setPeriodicDelay:delay interval:interval];
}

void C_NSButtonCell_SetButtonType(void* ptr, unsigned int _type) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setButtonType:_type];
}

void C_NSButtonCell_MouseEntered(void* ptr, void* event) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell mouseEntered:(NSEvent*)event];
}

void C_NSButtonCell_MouseExited(void* ptr, void* event) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell mouseExited:(NSEvent*)event];
}

void C_NSButtonCell_DrawBezelWithFrame_InView(void* ptr, CGRect frame, void* controlView) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell drawBezelWithFrame:frame inView:(NSView*)controlView];
}

void C_NSButtonCell_DrawImage_WithFrame_InView(void* ptr, void* image, CGRect frame, void* controlView) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell drawImage:(NSImage*)image withFrame:frame inView:(NSView*)controlView];
}

CGRect C_NSButtonCell_DrawTitle_WithFrame_InView(void* ptr, void* title, CGRect frame, void* controlView) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSRect result_ = [nSButtonCell drawTitle:(NSAttributedString*)title withFrame:frame inView:(NSView*)controlView];
    return result_;
}

void* C_NSButtonCell_AlternateTitle(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSString* result_ = [nSButtonCell alternateTitle];
    return result_;
}

void C_NSButtonCell_SetAlternateTitle(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setAlternateTitle:(NSString*)value];
}

void* C_NSButtonCell_AttributedAlternateTitle(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSAttributedString* result_ = [nSButtonCell attributedAlternateTitle];
    return result_;
}

void C_NSButtonCell_SetAttributedAlternateTitle(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setAttributedAlternateTitle:(NSAttributedString*)value];
}

void* C_NSButtonCell_AttributedTitle(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSAttributedString* result_ = [nSButtonCell attributedTitle];
    return result_;
}

void C_NSButtonCell_SetAttributedTitle(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setAttributedTitle:(NSAttributedString*)value];
}

void* C_NSButtonCell_AlternateImage(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSImage* result_ = [nSButtonCell alternateImage];
    return result_;
}

void C_NSButtonCell_SetAlternateImage(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setAlternateImage:(NSImage*)value];
}

unsigned int C_NSButtonCell_ImagePosition(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSCellImagePosition result_ = [nSButtonCell imagePosition];
    return result_;
}

void C_NSButtonCell_SetImagePosition(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setImagePosition:value];
}

unsigned int C_NSButtonCell_ImageScaling(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSImageScaling result_ = [nSButtonCell imageScaling];
    return result_;
}

void C_NSButtonCell_SetImageScaling(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setImageScaling:value];
}

void C_NSButtonCell_SetKeyEquivalent(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setKeyEquivalent:(NSString*)value];
}

unsigned int C_NSButtonCell_KeyEquivalentModifierMask(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSEventModifierFlags result_ = [nSButtonCell keyEquivalentModifierMask];
    return result_;
}

void C_NSButtonCell_SetKeyEquivalentModifierMask(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setKeyEquivalentModifierMask:value];
}

void* C_NSButtonCell_BackgroundColor(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSColor* result_ = [nSButtonCell backgroundColor];
    return result_;
}

void C_NSButtonCell_SetBackgroundColor(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setBackgroundColor:(NSColor*)value];
}

unsigned int C_NSButtonCell_BezelStyle(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSBezelStyle result_ = [nSButtonCell bezelStyle];
    return result_;
}

void C_NSButtonCell_SetBezelStyle(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setBezelStyle:value];
}

bool C_NSButtonCell_ImageDimsWhenDisabled(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    BOOL result_ = [nSButtonCell imageDimsWhenDisabled];
    return result_;
}

void C_NSButtonCell_SetImageDimsWhenDisabled(void* ptr, bool value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setImageDimsWhenDisabled:value];
}

bool C_NSButtonCell_IsTransparent(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    BOOL result_ = [nSButtonCell isTransparent];
    return result_;
}

void C_NSButtonCell_SetTransparent(void* ptr, bool value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setTransparent:value];
}

bool C_NSButtonCell_ShowsBorderOnlyWhileMouseInside(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    BOOL result_ = [nSButtonCell showsBorderOnlyWhileMouseInside];
    return result_;
}

void C_NSButtonCell_SetShowsBorderOnlyWhileMouseInside(void* ptr, bool value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setShowsBorderOnlyWhileMouseInside:value];
}

unsigned int C_NSButtonCell_HighlightsBy(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSCellStyleMask result_ = [nSButtonCell highlightsBy];
    return result_;
}

void C_NSButtonCell_SetHighlightsBy(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setHighlightsBy:value];
}

unsigned int C_NSButtonCell_ShowsStateBy(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSCellStyleMask result_ = [nSButtonCell showsStateBy];
    return result_;
}

void C_NSButtonCell_SetShowsStateBy(void* ptr, unsigned int value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setShowsStateBy:value];
}

void* C_NSButtonCell_Sound(void* ptr) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    NSSound* result_ = [nSButtonCell sound];
    return result_;
}

void C_NSButtonCell_SetSound(void* ptr, void* value) {
    NSButtonCell* nSButtonCell = (NSButtonCell*)ptr;
    [nSButtonCell setSound:(NSSound*)value];
}
