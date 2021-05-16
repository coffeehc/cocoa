#import <Appkit/Appkit.h>
#import "grid_column.h"

void* C_GridColumn_Alloc() {
    return [NSGridColumn alloc];
}

void* C_NSGridColumn_Init(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridColumn* result = [nSGridColumn init];
    return result;
}

void* C_NSGridColumn_CellAtIndex(void* ptr, int index) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCell* result = [nSGridColumn cellAtIndex:index];
    return result;
}

void C_NSGridColumn_MergeCellsInRange(void* ptr, NSRange _range) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn mergeCellsInRange:_range];
}

void* C_NSGridColumn_GridView(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridView* result = [nSGridColumn gridView];
    return result;
}

bool C_NSGridColumn_IsHidden(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    BOOL result = [nSGridColumn isHidden];
    return result;
}

void C_NSGridColumn_SetHidden(void* ptr, bool value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setHidden:value];
}

double C_NSGridColumn_LeadingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result = [nSGridColumn leadingPadding];
    return result;
}

void C_NSGridColumn_SetLeadingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setLeadingPadding:value];
}

int C_NSGridColumn_NumberOfCells(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSInteger result = [nSGridColumn numberOfCells];
    return result;
}

double C_NSGridColumn_TrailingPadding(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result = [nSGridColumn trailingPadding];
    return result;
}

void C_NSGridColumn_SetTrailingPadding(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setTrailingPadding:value];
}

double C_NSGridColumn_Width(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    CGFloat result = [nSGridColumn width];
    return result;
}

void C_NSGridColumn_SetWidth(void* ptr, double value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setWidth:value];
}

int C_NSGridColumn_XPlacement(void* ptr) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    NSGridCellPlacement result = [nSGridColumn xPlacement];
    return result;
}

void C_NSGridColumn_SetXPlacement(void* ptr, int value) {
    NSGridColumn* nSGridColumn = (NSGridColumn*)ptr;
    [nSGridColumn setXPlacement:value];
}
