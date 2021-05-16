#import <Appkit/Appkit.h>
#import "clip_view.h"

void* C_ClipView_Alloc() {
    return [NSClipView alloc];
}

void* C_NSClipView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result = [nSClipView initWithFrame:frameRect];
    return result;
}

void* C_NSClipView_InitWithCoder(void* ptr, void* coder) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result = [nSClipView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSClipView_Init(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result = [nSClipView init];
    return result;
}

void C_NSClipView_ScrollToPoint(void* ptr, CGPoint newOrigin) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView scrollToPoint:newOrigin];
}

CGRect C_NSClipView_ConstrainBoundsRect(void* ptr, CGRect proposedBounds) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result = [nSClipView constrainBoundsRect:proposedBounds];
    return result;
}

void C_NSClipView_ViewBoundsChanged(void* ptr, void* notification) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView viewBoundsChanged:(NSNotification*)notification];
}

void C_NSClipView_ViewFrameChanged(void* ptr, void* notification) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView viewFrameChanged:(NSNotification*)notification];
}

void* C_NSClipView_DocumentView(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSView* result = [nSClipView documentView];
    return result;
}

void C_NSClipView_SetDocumentView(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDocumentView:(NSView*)value];
}

NSEdgeInsets C_NSClipView_ContentInsets(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSEdgeInsets result = [nSClipView contentInsets];
    return result;
}

void C_NSClipView_SetContentInsets(void* ptr, NSEdgeInsets value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setContentInsets:value];
}

bool C_NSClipView_AutomaticallyAdjustsContentInsets(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    BOOL result = [nSClipView automaticallyAdjustsContentInsets];
    return result;
}

void C_NSClipView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setAutomaticallyAdjustsContentInsets:value];
}

CGRect C_NSClipView_DocumentRect(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result = [nSClipView documentRect];
    return result;
}

CGRect C_NSClipView_DocumentVisibleRect(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result = [nSClipView documentVisibleRect];
    return result;
}

void* C_NSClipView_DocumentCursor(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSCursor* result = [nSClipView documentCursor];
    return result;
}

void C_NSClipView_SetDocumentCursor(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDocumentCursor:(NSCursor*)value];
}

bool C_NSClipView_DrawsBackground(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    BOOL result = [nSClipView drawsBackground];
    return result;
}

void C_NSClipView_SetDrawsBackground(void* ptr, bool value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDrawsBackground:value];
}

void* C_NSClipView_BackgroundColor(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSColor* result = [nSClipView backgroundColor];
    return result;
}

void C_NSClipView_SetBackgroundColor(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setBackgroundColor:(NSColor*)value];
}
