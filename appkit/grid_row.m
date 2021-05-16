#import <Appkit/Appkit.h>
#import "grid_row.h"

void* C_GridRow_Alloc() {
    return [NSGridRow alloc];
}

void* C_NSGridRow_Init(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRow* result = [nSGridRow init];
    return result;
}

void* C_NSGridRow_CellAtIndex(void* ptr, int index) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCell* result = [nSGridRow cellAtIndex:index];
    return result;
}

void C_NSGridRow_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow mergeCellsInRange:_range];
}

int C_NSGridRow_NumberOfCells(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSInteger result = [nSGridRow numberOfCells];
    return result;
}

bool C_NSGridRow_IsHidden(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    BOOL result = [nSGridRow isHidden];
    return result;
}

void C_NSGridRow_SetHidden(void* ptr, bool value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHidden:value];
}

double C_NSGridRow_TopPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result = [nSGridRow topPadding];
    return result;
}

void C_NSGridRow_SetTopPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setTopPadding:value];
}

double C_NSGridRow_BottomPadding(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result = [nSGridRow bottomPadding];
    return result;
}

void C_NSGridRow_SetBottomPadding(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setBottomPadding:value];
}

double C_NSGridRow_Height(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    CGFloat result = [nSGridRow height];
    return result;
}

void C_NSGridRow_SetHeight(void* ptr, double value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setHeight:value];
}

int C_NSGridRow_RowAlignment(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridRowAlignment result = [nSGridRow rowAlignment];
    return result;
}

void C_NSGridRow_SetRowAlignment(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setRowAlignment:value];
}

int C_NSGridRow_YPlacement(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridCellPlacement result = [nSGridRow yPlacement];
    return result;
}

void C_NSGridRow_SetYPlacement(void* ptr, int value) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    [nSGridRow setYPlacement:value];
}

void* C_NSGridRow_GridView(void* ptr) {
    NSGridRow* nSGridRow = (NSGridRow*)ptr;
    NSGridView* result = [nSGridRow gridView];
    return result;
}
