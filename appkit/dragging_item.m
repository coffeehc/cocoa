#import <Appkit/Appkit.h>
#import "dragging_item.h"

void* C_DraggingItem_Alloc() {
    return [NSDraggingItem alloc];
}

void* C_NSDraggingItem_InitWithPasteboardWriter(void* ptr, void* pasteboardWriter) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    NSDraggingItem* result_ = [nSDraggingItem initWithPasteboardWriter:(id)pasteboardWriter];
    return result_;
}

void* C_NSDraggingItem_AllocDraggingItem() {
    NSDraggingItem* result_ = [NSDraggingItem alloc];
    return result_;
}

void* C_NSDraggingItem_Autorelease(void* ptr) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    NSDraggingItem* result_ = [nSDraggingItem autorelease];
    return result_;
}

void* C_NSDraggingItem_Retain(void* ptr) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    NSDraggingItem* result_ = [nSDraggingItem retain];
    return result_;
}

void C_NSDraggingItem_SetDraggingFrame_Contents(void* ptr, CGRect frame, void* contents) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    [nSDraggingItem setDraggingFrame:frame contents:(id)contents];
}

CGRect C_NSDraggingItem_DraggingFrame(void* ptr) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    NSRect result_ = [nSDraggingItem draggingFrame];
    return result_;
}

void C_NSDraggingItem_SetDraggingFrame(void* ptr, CGRect value) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    [nSDraggingItem setDraggingFrame:value];
}

Array C_NSDraggingItem_ImageComponents(void* ptr) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    NSArray* result_ = [nSDraggingItem imageComponents];
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

void* C_NSDraggingItem_Item(void* ptr) {
    NSDraggingItem* nSDraggingItem = (NSDraggingItem*)ptr;
    id result_ = [nSDraggingItem item];
    return result_;
}
