#import "image_view.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSImageView.h>

void* C_ImageView_Alloc() {
    return [NSImageView alloc];
}

void* C_NSImageView_ImageViewWithImage(void* image) {
    NSImageView* result_ = [NSImageView imageViewWithImage:(NSImage*)image];
    return result_;
}

void* C_NSImageView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result_ = [nSImageView initWithFrame:frameRect];
    return result_;
}

void* C_NSImageView_InitWithCoder(void* ptr, void* coder) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result_ = [nSImageView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSImageView_Init(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result_ = [nSImageView init];
    return result_;
}

void* C_NSImageView_AllocImageView() {
    NSImageView* result_ = [NSImageView alloc];
    return result_;
}

void* C_NSImageView_NewImageView() {
    NSImageView* result_ = [NSImageView new];
    return result_;
}

void* C_NSImageView_Autorelease(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result_ = [nSImageView autorelease];
    return result_;
}

void* C_NSImageView_Retain(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result_ = [nSImageView retain];
    return result_;
}

void* C_NSImageView_Image(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImage* result_ = [nSImageView image];
    return result_;
}

void C_NSImageView_SetImage(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImage:(NSImage*)value];
}

unsigned int C_NSImageView_ImageFrameStyle(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageFrameStyle result_ = [nSImageView imageFrameStyle];
    return result_;
}

void C_NSImageView_SetImageFrameStyle(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageFrameStyle:value];
}

unsigned int C_NSImageView_ImageAlignment(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageAlignment result_ = [nSImageView imageAlignment];
    return result_;
}

void C_NSImageView_SetImageAlignment(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageAlignment:value];
}

unsigned int C_NSImageView_ImageScaling(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageScaling result_ = [nSImageView imageScaling];
    return result_;
}

void C_NSImageView_SetImageScaling(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageScaling:value];
}

bool C_NSImageView_Animates(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result_ = [nSImageView animates];
    return result_;
}

void C_NSImageView_SetAnimates(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setAnimates:value];
}

bool C_NSImageView_IsEditable(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result_ = [nSImageView isEditable];
    return result_;
}

void C_NSImageView_SetEditable(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setEditable:value];
}

bool C_NSImageView_AllowsCutCopyPaste(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result_ = [nSImageView allowsCutCopyPaste];
    return result_;
}

void C_NSImageView_SetAllowsCutCopyPaste(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setAllowsCutCopyPaste:value];
}

void* C_NSImageView_ContentTintColor(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSColor* result_ = [nSImageView contentTintColor];
    return result_;
}

void C_NSImageView_SetContentTintColor(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setContentTintColor:(NSColor*)value];
}

void* C_NSImageView_SymbolConfiguration(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageView symbolConfiguration];
    return result_;
}

void C_NSImageView_SetSymbolConfiguration(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setSymbolConfiguration:(NSImageSymbolConfiguration*)value];
}
