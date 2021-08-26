#import <Appkit/Appkit.h>
#import "grid_cell.h"

void* C_GridCell_Alloc() {
    return [NSGridCell alloc];
}

void* C_NSGridCell_Init(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCell* result_ = [nSGridCell init];
    return result_;
}

void* C_NSGridCell_Column(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridColumn* result_ = [nSGridCell column];
    return result_;
}

void* C_NSGridCell_Row(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRow* result_ = [nSGridCell row];
    return result_;
}

void* C_NSGridCell_ContentView(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSView* result_ = [nSGridCell contentView];
    return result_;
}

void C_NSGridCell_SetContentView(void* ptr, void* value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setContentView:(NSView*)value];
}

void* C_NSGridCell_GridCell_EmptyContentView() {
    NSView* result_ = [NSGridCell emptyContentView];
    return result_;
}

Array C_NSGridCell_CustomPlacementConstraints(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSArray* result_ = [nSGridCell customPlacementConstraints];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSGridCell_SetCustomPlacementConstraints(void* ptr, Array value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSLayoutConstraint*)(NSLayoutConstraint*)p];
    	}
    }
    [nSGridCell setCustomPlacementConstraints:objcValue];
}

int C_NSGridCell_RowAlignment(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridRowAlignment result_ = [nSGridCell rowAlignment];
    return result_;
}

void C_NSGridCell_SetRowAlignment(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setRowAlignment:value];
}

int C_NSGridCell_XPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result_ = [nSGridCell xPlacement];
    return result_;
}

void C_NSGridCell_SetXPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setXPlacement:value];
}

int C_NSGridCell_YPlacement(void* ptr) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    NSGridCellPlacement result_ = [nSGridCell yPlacement];
    return result_;
}

void C_NSGridCell_SetYPlacement(void* ptr, int value) {
    NSGridCell* nSGridCell = (NSGridCell*)ptr;
    [nSGridCell setYPlacement:value];
}
