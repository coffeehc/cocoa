#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "grid_cell.h"

void* GridCell_Column(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell column];
}

void* GridCell_Row(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell row];
}

void* GridCell_ContentView(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell contentView];
}

Array GridCell_CustomPlacementConstraints(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	NSArray* ns_array = [gridCell customPlacementConstraints];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [ns_array objectAtIndex:i];
	}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

void GridCell_SetCustomPlacementConstraints(void* ptr, Array customPlacementConstraints) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
    NSMutableArray* objc_customPlacementConstraints = [[NSMutableArray alloc] init];
    void** customPlacementConstraintsData = (void**)customPlacementConstraints.data;
    for (int i = 0; i < customPlacementConstraints.len; i++) {
    	[objc_customPlacementConstraints addObject:(NSLayoutConstraint*)customPlacementConstraintsData[i]];
    }
	[gridCell setCustomPlacementConstraints:objc_customPlacementConstraints];
}

long GridCell_RowAlignment(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell rowAlignment];
}

void GridCell_SetRowAlignment(void* ptr, long rowAlignment) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	[gridCell setRowAlignment:rowAlignment];
}

long GridCell_XPlacement(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell xPlacement];
}

void GridCell_SetXPlacement(void* ptr, long xPlacement) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	[gridCell setXPlacement:xPlacement];
}

long GridCell_YPlacement(void* ptr) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	return [gridCell yPlacement];
}

void GridCell_SetYPlacement(void* ptr, long yPlacement) {
	NSGridCell* gridCell = (NSGridCell*)ptr;
	[gridCell setYPlacement:yPlacement];
}
