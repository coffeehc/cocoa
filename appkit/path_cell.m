#import "path_cell.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSPathCell.h>

void* C_PathCell_Alloc() {
    return [NSPathCell alloc];
}

void* C_NSPathCell_InitImageCell(void* ptr, void* image) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell initImageCell:(NSImage*)image];
    return result_;
}

void* C_NSPathCell_InitTextCell(void* ptr, void* _string) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSPathCell_Init(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell init];
    return result_;
}

void* C_NSPathCell_InitWithCoder(void* ptr, void* coder) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPathCell_AllocPathCell() {
    NSPathCell* result_ = [NSPathCell alloc];
    return result_;
}

void* C_NSPathCell_NewPathCell() {
    NSPathCell* result_ = [NSPathCell new];
    return result_;
}

void* C_NSPathCell_Autorelease(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell autorelease];
    return result_;
}

void* C_NSPathCell_Retain(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathCell* result_ = [nSPathCell retain];
    return result_;
}

void C_NSPathCell_MouseEntered_WithFrame_InView(void* ptr, void* event, CGRect frame, void* view) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell mouseEntered:(NSEvent*)event withFrame:frame inView:(NSView*)view];
}

void C_NSPathCell_MouseExited_WithFrame_InView(void* ptr, void* event, CGRect frame, void* view) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell mouseExited:(NSEvent*)event withFrame:frame inView:(NSView*)view];
}

CGRect C_NSPathCell_RectOfPathComponentCell_WithFrame_InView(void* ptr, void* cell, CGRect frame, void* view) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSRect result_ = [nSPathCell rectOfPathComponentCell:(NSPathComponentCell*)cell withFrame:frame inView:(NSView*)view];
    return result_;
}

void* C_NSPathCell_PathComponentCellAtPoint_WithFrame_InView(void* ptr, CGPoint point, CGRect frame, void* view) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathComponentCell* result_ = [nSPathCell pathComponentCellAtPoint:point withFrame:frame inView:(NSView*)view];
    return result_;
}

Array C_NSPathCell_AllowedTypes(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSArray* result_ = [nSPathCell allowedTypes];
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

void C_NSPathCell_SetAllowedTypes(void* ptr, Array value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSPathCell setAllowedTypes:objcValue];
}

int C_NSPathCell_PathStyle(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathStyle result_ = [nSPathCell pathStyle];
    return result_;
}

void C_NSPathCell_SetPathStyle(void* ptr, int value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setPathStyle:value];
}

void* C_NSPathCell_PlaceholderAttributedString(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSAttributedString* result_ = [nSPathCell placeholderAttributedString];
    return result_;
}

void C_NSPathCell_SetPlaceholderAttributedString(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setPlaceholderAttributedString:(NSAttributedString*)value];
}

void* C_NSPathCell_PlaceholderString(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSString* result_ = [nSPathCell placeholderString];
    return result_;
}

void C_NSPathCell_SetPlaceholderString(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setPlaceholderString:(NSString*)value];
}

void* C_NSPathCell_BackgroundColor(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSColor* result_ = [nSPathCell backgroundColor];
    return result_;
}

void C_NSPathCell_SetBackgroundColor(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setBackgroundColor:(NSColor*)value];
}

void* C_NSPathCell_ClickedPathComponentCell(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSPathComponentCell* result_ = [nSPathCell clickedPathComponentCell];
    return result_;
}

Array C_NSPathCell_PathComponentCells(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSArray* result_ = [nSPathCell pathComponentCells];
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

void C_NSPathCell_SetPathComponentCells(void* ptr, Array value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSPathComponentCell*)p];
    	}
    }
    [nSPathCell setPathComponentCells:objcValue];
}

void* C_NSPathCell_DoubleAction(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    SEL result_ = [nSPathCell doubleAction];
    return result_;
}

void C_NSPathCell_SetDoubleAction(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setDoubleAction:(SEL)value];
}

void* C_NSPathCell_URL(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    NSURL* result_ = [nSPathCell URL];
    return result_;
}

void C_NSPathCell_SetURL(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setURL:(NSURL*)value];
}

void* C_NSPathCell_Delegate(void* ptr) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    id result_ = [nSPathCell delegate];
    return result_;
}

void C_NSPathCell_SetDelegate(void* ptr, void* value) {
    NSPathCell* nSPathCell = (NSPathCell*)ptr;
    [nSPathCell setDelegate:(id)value];
}
