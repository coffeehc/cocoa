#import "grid_row.h"
#import <AppKit/NSGridView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_GridRow_Alloc() {
    return [NSGridRow alloc];
}

void* C_NSGridRow_AllocGridRow() {
    NSGridRow* result_ = [NSGridRow alloc];
    return result_;
}

void* C_NSGridRow_Init(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow init];
    return result_;
}

void* C_NSGridRow_NewGridRow() {
    NSGridRow* result_ = [NSGridRow new];
    return result_;
}

void* C_NSGridRow_Autorelease(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow autorelease];
    return result_;
}

void* C_NSGridRow_Retain(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result_ = [nSGridRow retain];
    return result_;
}

void* C_NSGridRow_CellAtIndex(void* ptr, int index) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCell* result_ = [nSGridRow cellAtIndex:index];
    return result_;
}

void C_NSGridRow_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow mergeCellsInRange:_range];
}

int C_NSGridRow_NumberOfCells(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSInteger result_ = [nSGridRow numberOfCells];
    return result_;
}

bool C_NSGridRow_IsHidden(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    BOOL result_ = [nSGridRow isHidden];
    return result_;
}

void C_NSGridRow_SetHidden(void* ptr, bool value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHidden:value];
}

double C_NSGridRow_TopPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow topPadding];
    return result_;
}

void C_NSGridRow_SetTopPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setTopPadding:value];
}

double C_NSGridRow_BottomPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow bottomPadding];
    return result_;
}

void C_NSGridRow_SetBottomPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setBottomPadding:value];
}

double C_NSGridRow_Height(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result_ = [nSGridRow height];
    return result_;
}

void C_NSGridRow_SetHeight(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHeight:value];
}

int C_NSGridRow_RowAlignment(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRowAlignment result_ = [nSGridRow rowAlignment];
    return result_;
}

void C_NSGridRow_SetRowAlignment(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setRowAlignment:value];
}

int C_NSGridRow_YPlacement(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCellPlacement result_ = [nSGridRow yPlacement];
    return result_;
}

void C_NSGridRow_SetYPlacement(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setYPlacement:value];
}

void* C_NSGridRow_GridView(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridView* result_ = [nSGridRow gridView];
    return result_;
}
