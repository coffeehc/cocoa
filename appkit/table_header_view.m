#import "table_header_view.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSTableHeaderView.h>

void* C_TableHeaderView_Alloc() {
    return [NSTableHeaderView alloc];
}

void* C_NSTableHeaderView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableHeaderView* result_ = [nSTableHeaderView initWithFrame:frameRect];
    return result_;
}

void* C_NSTableHeaderView_InitWithCoder(void* ptr, void* coder) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableHeaderView* result_ = [nSTableHeaderView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableHeaderView_Init(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableHeaderView* result_ = [nSTableHeaderView init];
    return result_;
}

void* C_NSTableHeaderView_AllocTableHeaderView() {
    NSTableHeaderView* result_ = [NSTableHeaderView alloc];
    return result_;
}

void* C_NSTableHeaderView_NewTableHeaderView() {
    NSTableHeaderView* result_ = [NSTableHeaderView new];
    return result_;
}

void* C_NSTableHeaderView_Autorelease(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableHeaderView* result_ = [nSTableHeaderView autorelease];
    return result_;
}

void* C_NSTableHeaderView_Retain(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableHeaderView* result_ = [nSTableHeaderView retain];
    return result_;
}

int C_NSTableHeaderView_ColumnAtPoint(void* ptr, CGPoint point) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSInteger result_ = [nSTableHeaderView columnAtPoint:point];
    return result_;
}

CGRect C_NSTableHeaderView_HeaderRectOfColumn(void* ptr, int column) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSRect result_ = [nSTableHeaderView headerRectOfColumn:column];
    return result_;
}

void* C_NSTableHeaderView_TableView(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSTableView* result_ = [nSTableHeaderView tableView];
    return result_;
}

void C_NSTableHeaderView_SetTableView(void* ptr, void* value) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    [nSTableHeaderView setTableView:(NSTableView*)value];
}

int C_NSTableHeaderView_DraggedColumn(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSInteger result_ = [nSTableHeaderView draggedColumn];
    return result_;
}

double C_NSTableHeaderView_DraggedDistance(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    CGFloat result_ = [nSTableHeaderView draggedDistance];
    return result_;
}

int C_NSTableHeaderView_ResizedColumn(void* ptr) {
    NSTableHeaderView* nSTableHeaderView = (NSTableHeaderView*)ptr;
    NSInteger result_ = [nSTableHeaderView resizedColumn];
    return result_;
}
