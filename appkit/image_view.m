#import <Appkit/Appkit.h>
#import "image_view.h"

void* C_ImageView_Alloc() {
    return [NSImageView alloc];
}

void* C_NSImageView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result = [nSImageView initWithFrame:frameRect];
    return result;
}

void* C_NSImageView_InitWithCoder(void* ptr, void* coder) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result = [nSImageView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSImageView_Init(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageView* result = [nSImageView init];
    return result;
}

void* C_NSImageView_ImageViewWithImage(void* image) {
    NSImageView* result = [NSImageView imageViewWithImage:(NSImage*)image];
    return result;
}

void* C_NSImageView_Image(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImage* result = [nSImageView image];
    return result;
}

void C_NSImageView_SetImage(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImage:(NSImage*)value];
}

unsigned int C_NSImageView_ImageFrameStyle(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageFrameStyle result = [nSImageView imageFrameStyle];
    return result;
}

void C_NSImageView_SetImageFrameStyle(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageFrameStyle:value];
}

unsigned int C_NSImageView_ImageAlignment(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageAlignment result = [nSImageView imageAlignment];
    return result;
}

void C_NSImageView_SetImageAlignment(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageAlignment:value];
}

unsigned int C_NSImageView_ImageScaling(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageScaling result = [nSImageView imageScaling];
    return result;
}

void C_NSImageView_SetImageScaling(void* ptr, unsigned int value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setImageScaling:value];
}

bool C_NSImageView_Animates(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result = [nSImageView animates];
    return result;
}

void C_NSImageView_SetAnimates(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setAnimates:value];
}

bool C_NSImageView_IsEditable(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result = [nSImageView isEditable];
    return result;
}

void C_NSImageView_SetEditable(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setEditable:value];
}

bool C_NSImageView_AllowsCutCopyPaste(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    BOOL result = [nSImageView allowsCutCopyPaste];
    return result;
}

void C_NSImageView_SetAllowsCutCopyPaste(void* ptr, bool value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setAllowsCutCopyPaste:value];
}

void* C_NSImageView_ContentTintColor(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSColor* result = [nSImageView contentTintColor];
    return result;
}

void C_NSImageView_SetContentTintColor(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setContentTintColor:(NSColor*)value];
}

void* C_NSImageView_SymbolConfiguration(void* ptr) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    NSImageSymbolConfiguration* result = [nSImageView symbolConfiguration];
    return result;
}

void C_NSImageView_SetSymbolConfiguration(void* ptr, void* value) {
    NSImageView* nSImageView = (NSImageView*)ptr;
    [nSImageView setSymbolConfiguration:(NSImageSymbolConfiguration*)value];
}
