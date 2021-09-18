#import "browser_cell.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSBrowserCell.h>

void* C_BrowserCell_Alloc() {
    return [NSBrowserCell alloc];
}

void* C_NSBrowserCell_InitWithCoder(void* ptr, void* coder) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSBrowserCell_InitImageCell(void* ptr, void* image) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell initImageCell:(NSImage*)image];
    return result_;
}

void* C_NSBrowserCell_InitTextCell(void* ptr, void* _string) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSBrowserCell_Init(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell init];
    return result_;
}

void* C_NSBrowserCell_AllocBrowserCell() {
    NSBrowserCell* result_ = [NSBrowserCell alloc];
    return result_;
}

void* C_NSBrowserCell_NewBrowserCell() {
    NSBrowserCell* result_ = [NSBrowserCell new];
    return result_;
}

void* C_NSBrowserCell_Autorelease(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell autorelease];
    return result_;
}

void* C_NSBrowserCell_Retain(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSBrowserCell* result_ = [nSBrowserCell retain];
    return result_;
}

void C_NSBrowserCell_Reset(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    [nSBrowserCell reset];
}

void C_NSBrowserCell_Set(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    [nSBrowserCell set];
}

void* C_NSBrowserCell_HighlightColorInView(void* ptr, void* controlView) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSColor* result_ = [nSBrowserCell highlightColorInView:(NSView*)controlView];
    return result_;
}

void* C_NSBrowserCell_BrowserCell_BranchImage() {
    NSImage* result_ = [NSBrowserCell branchImage];
    return result_;
}

void* C_NSBrowserCell_BrowserCell_HighlightedBranchImage() {
    NSImage* result_ = [NSBrowserCell highlightedBranchImage];
    return result_;
}

void* C_NSBrowserCell_AlternateImage(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    NSImage* result_ = [nSBrowserCell alternateImage];
    return result_;
}

void C_NSBrowserCell_SetAlternateImage(void* ptr, void* value) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    [nSBrowserCell setAlternateImage:(NSImage*)value];
}

bool C_NSBrowserCell_IsLeaf(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    BOOL result_ = [nSBrowserCell isLeaf];
    return result_;
}

void C_NSBrowserCell_SetLeaf(void* ptr, bool value) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    [nSBrowserCell setLeaf:value];
}

bool C_NSBrowserCell_IsLoaded(void* ptr) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    BOOL result_ = [nSBrowserCell isLoaded];
    return result_;
}

void C_NSBrowserCell_SetLoaded(void* ptr, bool value) {
    NSBrowserCell* nSBrowserCell = (NSBrowserCell*)ptr;
    [nSBrowserCell setLoaded:value];
}
