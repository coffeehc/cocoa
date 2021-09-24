#import "dragging_item.h"
#import <AppKit/NSDraggingItem.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

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

void* C_DraggingImageComponent_Alloc() {
    return [NSDraggingImageComponent alloc];
}

void* C_NSDraggingImageComponent_InitWithKey(void* ptr, void* key) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponent* result_ = [nSDraggingImageComponent initWithKey:(NSString*)key];
    return result_;
}

void* C_NSDraggingImageComponent_AllocDraggingImageComponent() {
    NSDraggingImageComponent* result_ = [NSDraggingImageComponent alloc];
    return result_;
}

void* C_NSDraggingImageComponent_Autorelease(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponent* result_ = [nSDraggingImageComponent autorelease];
    return result_;
}

void* C_NSDraggingImageComponent_Retain(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponent* result_ = [nSDraggingImageComponent retain];
    return result_;
}

void* C_NSDraggingImageComponent_DraggingImageComponentWithKey(void* key) {
    NSDraggingImageComponent* result_ = [NSDraggingImageComponent draggingImageComponentWithKey:(NSString*)key];
    return result_;
}

void* C_NSDraggingImageComponent_Key(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponentKey result_ = [nSDraggingImageComponent key];
    return result_;
}

void C_NSDraggingImageComponent_SetKey(void* ptr, void* value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setKey:(NSString*)value];
}

void* C_NSDraggingImageComponent_Contents(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    id result_ = [nSDraggingImageComponent contents];
    return result_;
}

void C_NSDraggingImageComponent_SetContents(void* ptr, void* value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setContents:(id)value];
}

CGRect C_NSDraggingImageComponent_Frame(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSRect result_ = [nSDraggingImageComponent frame];
    return result_;
}

void C_NSDraggingImageComponent_SetFrame(void* ptr, CGRect value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setFrame:value];
}
