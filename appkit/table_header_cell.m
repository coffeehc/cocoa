#import <Appkit/Appkit.h>
#import "table_header_cell.h"

void* C_TableHeaderCell_Alloc() {
    return [NSTableHeaderCell alloc];
}

void* C_NSTableHeaderCell_InitTextCell(void* ptr, void* _string) {
    NSTableHeaderCell* nSTableHeaderCell = (NSTableHeaderCell*)ptr;
    NSTableHeaderCell* result_ = [nSTableHeaderCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSTableHeaderCell_InitWithCoder(void* ptr, void* coder) {
    NSTableHeaderCell* nSTableHeaderCell = (NSTableHeaderCell*)ptr;
    NSTableHeaderCell* result_ = [nSTableHeaderCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTableHeaderCell_Init(void* ptr) {
    NSTableHeaderCell* nSTableHeaderCell = (NSTableHeaderCell*)ptr;
    NSTableHeaderCell* result_ = [nSTableHeaderCell init];
    return result_;
}

void C_NSTableHeaderCell_DrawSortIndicatorWithFrame_InView_Ascending_Priority(void* ptr, CGRect cellFrame, void* controlView, bool ascending, int priority) {
    NSTableHeaderCell* nSTableHeaderCell = (NSTableHeaderCell*)ptr;
    [nSTableHeaderCell drawSortIndicatorWithFrame:cellFrame inView:(NSView*)controlView ascending:ascending priority:priority];
}

CGRect C_NSTableHeaderCell_SortIndicatorRectForBounds(void* ptr, CGRect rect) {
    NSTableHeaderCell* nSTableHeaderCell = (NSTableHeaderCell*)ptr;
    NSRect result_ = [nSTableHeaderCell sortIndicatorRectForBounds:rect];
    return result_;
}
