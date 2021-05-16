#import <Appkit/Appkit.h>
#import "grid_cell.h"

void* C_GridCell_Alloc() {
    return [NSGridCell alloc];
}

void* C_NSGridCell_Init(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCell* result = [nSGridCell init];
    return result;
}

void* C_NSGridCell_Column(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridColumn* result = [nSGridCell column];
    return result;
}

void* C_NSGridCell_Row(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRow* result = [nSGridCell row];
    return result;
}

void* C_NSGridCell_ContentView(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSView* result = [nSGridCell contentView];
    return result;
}

void C_NSGridCell_SetContentView(void* ptr, void* value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setContentView:(NSView*)value];
}

void* C_NSGridCell_GridCell_EmptyContentView() {
    NSView* result = [NSGridCell emptyContentView];
    return result;
}

Array C_NSGridCell_CustomPlacementConstraints(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSArray* result = [nSGridCell customPlacementConstraints];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSGridCell_SetCustomPlacementConstraints(void* ptr, Array value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSLayoutConstraint*)(NSLayoutConstraint*)p];
    }
    [nSGridCell setCustomPlacementConstraints:objcValue];
}

int C_NSGridCell_RowAlignment(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRowAlignment result = [nSGridCell rowAlignment];
    return result;
}

void C_NSGridCell_SetRowAlignment(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setRowAlignment:value];
}

int C_NSGridCell_XPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result = [nSGridCell xPlacement];
    return result;
}

void C_NSGridCell_SetXPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setXPlacement:value];
}

int C_NSGridCell_YPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result = [nSGridCell yPlacement];
    return result;
}

void C_NSGridCell_SetYPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setYPlacement:value];
}
