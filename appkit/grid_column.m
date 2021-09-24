#import "grid_column.h"
#import <AppKit/NSGridView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_GridColumn_Alloc() {
    return [NSGridColumn alloc];
}

void* C_NSGridColumn_AllocGridColumn() {
    NSGridColumn* result_ = [NSGridColumn alloc];
    return result_;
}

void* C_NSGridColumn_Init(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn init];
    return result_;
}

void* C_NSGridColumn_NewGridColumn() {
    NSGridColumn* result_ = [NSGridColumn new];
    return result_;
}

void* C_NSGridColumn_Autorelease(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn autorelease];
    return result_;
}

void* C_NSGridColumn_Retain(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result_ = [nSGridColumn retain];
    return result_;
}

void* C_NSGridColumn_CellAtIndex(void* ptr, int index) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCell* result_ = [nSGridColumn cellAtIndex:index];
    return result_;
}

void C_NSGridColumn_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn mergeCellsInRange:_range];
}

void* C_NSGridColumn_GridView(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridView* result_ = [nSGridColumn gridView];
    return result_;
}

bool C_NSGridColumn_IsHidden(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    BOOL result_ = [nSGridColumn isHidden];
    return result_;
}

void C_NSGridColumn_SetHidden(void* ptr, bool value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setHidden:value];
}

double C_NSGridColumn_LeadingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn leadingPadding];
    return result_;
}

void C_NSGridColumn_SetLeadingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setLeadingPadding:value];
}

int C_NSGridColumn_NumberOfCells(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSInteger result_ = [nSGridColumn numberOfCells];
    return result_;
}

double C_NSGridColumn_TrailingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn trailingPadding];
    return result_;
}

void C_NSGridColumn_SetTrailingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setTrailingPadding:value];
}

double C_NSGridColumn_Width(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result_ = [nSGridColumn width];
    return result_;
}

void C_NSGridColumn_SetWidth(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setWidth:value];
}

int C_NSGridColumn_XPlacement(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCellPlacement result_ = [nSGridColumn xPlacement];
    return result_;
}

void C_NSGridColumn_SetXPlacement(void* ptr, int value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setXPlacement:value];
}
