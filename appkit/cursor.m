#import "cursor.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSCursor.h>

void* C_Cursor_Alloc() {
    return [NSCursor alloc];
}

void* C_NSCursor_InitWithImage_HotSpot(void* ptr, void* newImage, CGPoint point) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSCursor* result_ = [nSCursor initWithImage:(NSImage*)newImage hotSpot:point];
    return result_;
}

void* C_NSCursor_InitWithCoder(void* ptr, void* coder) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSCursor* result_ = [nSCursor initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSCursor_AllocCursor() {
    NSCursor* result_ = [NSCursor alloc];
    return result_;
}

void* C_NSCursor_Init(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSCursor* result_ = [nSCursor init];
    return result_;
}

void* C_NSCursor_NewCursor() {
    NSCursor* result_ = [NSCursor new];
    return result_;
}

void* C_NSCursor_Autorelease(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSCursor* result_ = [nSCursor autorelease];
    return result_;
}

void* C_NSCursor_Retain(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSCursor* result_ = [nSCursor retain];
    return result_;
}

void C_NSCursor_Cursor_Hide() {
    [NSCursor hide];
}

void C_NSCursor_Cursor_Unhide() {
    [NSCursor unhide];
}

void C_NSCursor_Cursor_SetHiddenUntilMouseMoves(bool flag) {
    [NSCursor setHiddenUntilMouseMoves:flag];
}

void C_NSCursor_Cursor_Pop() {
    [NSCursor pop];
}

void C_NSCursor_Push(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    [nSCursor push];
}

void C_NSCursor_Set(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    [nSCursor set];
}

void* C_NSCursor_Image(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSImage* result_ = [nSCursor image];
    return result_;
}

CGPoint C_NSCursor_HotSpot(void* ptr) {
    NSCursor* nSCursor = (NSCursor*)ptr;
    NSPoint result_ = [nSCursor hotSpot];
    return result_;
}

void* C_NSCursor_CurrentCursor() {
    NSCursor* result_ = [NSCursor currentCursor];
    return result_;
}

void* C_NSCursor_CurrentSystemCursor() {
    NSCursor* result_ = [NSCursor currentSystemCursor];
    return result_;
}

void* C_NSCursor_ArrowCursor() {
    NSCursor* result_ = [NSCursor arrowCursor];
    return result_;
}

void* C_NSCursor_ContextualMenuCursor() {
    NSCursor* result_ = [NSCursor contextualMenuCursor];
    return result_;
}

void* C_NSCursor_ClosedHandCursor() {
    NSCursor* result_ = [NSCursor closedHandCursor];
    return result_;
}

void* C_NSCursor_CrosshairCursor() {
    NSCursor* result_ = [NSCursor crosshairCursor];
    return result_;
}

void* C_NSCursor_DisappearingItemCursor() {
    NSCursor* result_ = [NSCursor disappearingItemCursor];
    return result_;
}

void* C_NSCursor_DragCopyCursor() {
    NSCursor* result_ = [NSCursor dragCopyCursor];
    return result_;
}

void* C_NSCursor_DragLinkCursor() {
    NSCursor* result_ = [NSCursor dragLinkCursor];
    return result_;
}

void* C_NSCursor_IBeamCursor() {
    NSCursor* result_ = [NSCursor IBeamCursor];
    return result_;
}

void* C_NSCursor_OpenHandCursor() {
    NSCursor* result_ = [NSCursor openHandCursor];
    return result_;
}

void* C_NSCursor_OperationNotAllowedCursor() {
    NSCursor* result_ = [NSCursor operationNotAllowedCursor];
    return result_;
}

void* C_NSCursor_PointingHandCursor() {
    NSCursor* result_ = [NSCursor pointingHandCursor];
    return result_;
}

void* C_NSCursor_ResizeDownCursor() {
    NSCursor* result_ = [NSCursor resizeDownCursor];
    return result_;
}

void* C_NSCursor_ResizeLeftCursor() {
    NSCursor* result_ = [NSCursor resizeLeftCursor];
    return result_;
}

void* C_NSCursor_ResizeLeftRightCursor() {
    NSCursor* result_ = [NSCursor resizeLeftRightCursor];
    return result_;
}

void* C_NSCursor_ResizeRightCursor() {
    NSCursor* result_ = [NSCursor resizeRightCursor];
    return result_;
}

void* C_NSCursor_ResizeUpCursor() {
    NSCursor* result_ = [NSCursor resizeUpCursor];
    return result_;
}

void* C_NSCursor_ResizeUpDownCursor() {
    NSCursor* result_ = [NSCursor resizeUpDownCursor];
    return result_;
}

void* C_NSCursor_IBeamCursorForVerticalLayout() {
    NSCursor* result_ = [NSCursor IBeamCursorForVerticalLayout];
    return result_;
}
