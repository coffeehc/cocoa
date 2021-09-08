#import <Appkit/Appkit.h>
#import "clip_view.h"

void* C_ClipView_Alloc() {
    return [NSClipView alloc];
}

void* C_NSClipView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result_ = [nSClipView initWithFrame:frameRect];
    return result_;
}

void* C_NSClipView_InitWithCoder(void* ptr, void* coder) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result_ = [nSClipView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSClipView_Init(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result_ = [nSClipView init];
    return result_;
}

void* C_NSClipView_AllocClipView() {
    NSClipView* result_ = [NSClipView alloc];
    return result_;
}

void* C_NSClipView_NewClipView() {
    NSClipView* result_ = [NSClipView new];
    return result_;
}

void* C_NSClipView_Autorelease(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result_ = [nSClipView autorelease];
    return result_;
}

void* C_NSClipView_Retain(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSClipView* result_ = [nSClipView retain];
    return result_;
}

void C_NSClipView_ScrollToPoint(void* ptr, CGPoint newOrigin) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView scrollToPoint:newOrigin];
}

CGRect C_NSClipView_ConstrainBoundsRect(void* ptr, CGRect proposedBounds) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result_ = [nSClipView constrainBoundsRect:proposedBounds];
    return result_;
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
    NSView* result_ = [nSClipView documentView];
    return result_;
}

void C_NSClipView_SetDocumentView(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDocumentView:(NSView*)value];
}

NSEdgeInsets C_NSClipView_ContentInsets(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSEdgeInsets result_ = [nSClipView contentInsets];
    return result_;
}

void C_NSClipView_SetContentInsets(void* ptr, NSEdgeInsets value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setContentInsets:value];
}

bool C_NSClipView_AutomaticallyAdjustsContentInsets(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    BOOL result_ = [nSClipView automaticallyAdjustsContentInsets];
    return result_;
}

void C_NSClipView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setAutomaticallyAdjustsContentInsets:value];
}

CGRect C_NSClipView_DocumentRect(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result_ = [nSClipView documentRect];
    return result_;
}

CGRect C_NSClipView_DocumentVisibleRect(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSRect result_ = [nSClipView documentVisibleRect];
    return result_;
}

void* C_NSClipView_DocumentCursor(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSCursor* result_ = [nSClipView documentCursor];
    return result_;
}

void C_NSClipView_SetDocumentCursor(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDocumentCursor:(NSCursor*)value];
}

bool C_NSClipView_DrawsBackground(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    BOOL result_ = [nSClipView drawsBackground];
    return result_;
}

void C_NSClipView_SetDrawsBackground(void* ptr, bool value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setDrawsBackground:value];
}

void* C_NSClipView_BackgroundColor(void* ptr) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    NSColor* result_ = [nSClipView backgroundColor];
    return result_;
}

void C_NSClipView_SetBackgroundColor(void* ptr, void* value) {
    NSClipView* nSClipView = (NSClipView*)ptr;
    [nSClipView setBackgroundColor:(NSColor*)value];
}
