#import "table_row_view.h"
#import <AppKit/NSTableRowView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TableRowView_Alloc() {
    return [NSTableRowView alloc];
}

void* C_NSTableRowView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableRowView* result_ = [nSTableRowView initWithFrame:frameRect];
    return result_;
}

void* C_NSTableRowView_InitWithCoder(void* ptr, void* coder) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableRowView* result_ = [nSTableRowView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableRowView_Init(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableRowView* result_ = [nSTableRowView init];
    return result_;
}

void* C_NSTableRowView_AllocTableRowView() {
    NSTableRowView* result_ = [NSTableRowView alloc];
    return result_;
}

void* C_NSTableRowView_NewTableRowView() {
    NSTableRowView* result_ = [NSTableRowView new];
    return result_;
}

void* C_NSTableRowView_Autorelease(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableRowView* result_ = [nSTableRowView autorelease];
    return result_;
}

void* C_NSTableRowView_Retain(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableRowView* result_ = [nSTableRowView retain];
    return result_;
}

void C_NSTableRowView_DrawBackgroundInRect(void* ptr, CGRect dirtyRect) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView drawBackgroundInRect:dirtyRect];
}

void C_NSTableRowView_DrawDraggingDestinationFeedbackInRect(void* ptr, CGRect dirtyRect) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView drawDraggingDestinationFeedbackInRect:dirtyRect];
}

void C_NSTableRowView_DrawSelectionInRect(void* ptr, CGRect dirtyRect) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView drawSelectionInRect:dirtyRect];
}

void C_NSTableRowView_DrawSeparatorInRect(void* ptr, CGRect dirtyRect) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView drawSeparatorInRect:dirtyRect];
}

void* C_NSTableRowView_ViewAtColumn(void* ptr, int column) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    id result_ = [nSTableRowView viewAtColumn:column];
    return result_;
}

bool C_NSTableRowView_IsEmphasized(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isEmphasized];
    return result_;
}

void C_NSTableRowView_SetEmphasized(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setEmphasized:value];
}

int C_NSTableRowView_InteriorBackgroundStyle(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSBackgroundStyle result_ = [nSTableRowView interiorBackgroundStyle];
    return result_;
}

bool C_NSTableRowView_IsFloating(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isFloating];
    return result_;
}

void C_NSTableRowView_SetFloating(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setFloating:value];
}

bool C_NSTableRowView_IsSelected(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isSelected];
    return result_;
}

void C_NSTableRowView_SetSelected(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setSelected:value];
}

int C_NSTableRowView_SelectionHighlightStyle(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableViewSelectionHighlightStyle result_ = [nSTableRowView selectionHighlightStyle];
    return result_;
}

void C_NSTableRowView_SetSelectionHighlightStyle(void* ptr, int value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setSelectionHighlightStyle:value];
}

int C_NSTableRowView_DraggingDestinationFeedbackStyle(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSTableViewDraggingDestinationFeedbackStyle result_ = [nSTableRowView draggingDestinationFeedbackStyle];
    return result_;
}

void C_NSTableRowView_SetDraggingDestinationFeedbackStyle(void* ptr, int value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setDraggingDestinationFeedbackStyle:value];
}

double C_NSTableRowView_IndentationForDropOperation(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    CGFloat result_ = [nSTableRowView indentationForDropOperation];
    return result_;
}

void C_NSTableRowView_SetIndentationForDropOperation(void* ptr, double value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setIndentationForDropOperation:value];
}

bool C_NSTableRowView_IsTargetForDropOperation(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isTargetForDropOperation];
    return result_;
}

void C_NSTableRowView_SetTargetForDropOperation(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setTargetForDropOperation:value];
}

bool C_NSTableRowView_IsGroupRowStyle(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isGroupRowStyle];
    return result_;
}

void C_NSTableRowView_SetGroupRowStyle(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setGroupRowStyle:value];
}

int C_NSTableRowView_NumberOfColumns(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSInteger result_ = [nSTableRowView numberOfColumns];
    return result_;
}

void* C_NSTableRowView_BackgroundColor(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    NSColor* result_ = [nSTableRowView backgroundColor];
    return result_;
}

void C_NSTableRowView_SetBackgroundColor(void* ptr, void* value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setBackgroundColor:(NSColor*)value];
}

bool C_NSTableRowView_IsNextRowSelected(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isNextRowSelected];
    return result_;
}

void C_NSTableRowView_SetNextRowSelected(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setNextRowSelected:value];
}

bool C_NSTableRowView_IsPreviousRowSelected(void* ptr) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    BOOL result_ = [nSTableRowView isPreviousRowSelected];
    return result_;
}

void C_NSTableRowView_SetPreviousRowSelected(void* ptr, bool value) {
    NSTableRowView* nSTableRowView = (NSTableRowView*)ptr;
    [nSTableRowView setPreviousRowSelected:value];
}
