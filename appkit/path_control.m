#import <Appkit/Appkit.h>
#import "path_control.h"

void* C_PathControl_Alloc() {
    return [NSPathControl alloc];
}

void* C_NSPathControl_InitWithFrame(void* ptr, CGRect frameRect) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControl* result_ = [nSPathControl initWithFrame:frameRect];
    return result_;
}

void* C_NSPathControl_InitWithCoder(void* ptr, void* coder) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControl* result_ = [nSPathControl initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPathControl_Init(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControl* result_ = [nSPathControl init];
    return result_;
}

void* C_NSPathControl_AllocPathControl() {
    NSPathControl* result_ = [NSPathControl alloc];
    return result_;
}

void* C_NSPathControl_NewPathControl() {
    NSPathControl* result_ = [NSPathControl new];
    return result_;
}

void* C_NSPathControl_Autorelease(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControl* result_ = [nSPathControl autorelease];
    return result_;
}

void* C_NSPathControl_Retain(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControl* result_ = [nSPathControl retain];
    return result_;
}

void C_NSPathControl_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int mask, bool isLocal) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setDraggingSourceOperationMask:mask forLocal:isLocal];
}

int C_NSPathControl_PathStyle(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathStyle result_ = [nSPathControl pathStyle];
    return result_;
}

void C_NSPathControl_SetPathStyle(void* ptr, int value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setPathStyle:value];
}

void* C_NSPathControl_BackgroundColor(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSColor* result_ = [nSPathControl backgroundColor];
    return result_;
}

void C_NSPathControl_SetBackgroundColor(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setBackgroundColor:(NSColor*)value];
}

void* C_NSPathControl_DoubleAction(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    SEL result_ = [nSPathControl doubleAction];
    return result_;
}

void C_NSPathControl_SetDoubleAction(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setDoubleAction:(SEL)value];
}

void* C_NSPathControl_URL(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSURL* result_ = [nSPathControl URL];
    return result_;
}

void C_NSPathControl_SetURL(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setURL:(NSURL*)value];
}

void* C_NSPathControl_Delegate(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    id result_ = [nSPathControl delegate];
    return result_;
}

void C_NSPathControl_SetDelegate(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setDelegate:(id)value];
}

Array C_NSPathControl_AllowedTypes(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSArray* result_ = [nSPathControl allowedTypes];
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

void C_NSPathControl_SetAllowedTypes(void* ptr, Array value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)(NSString*)p];
    	}
    }
    [nSPathControl setAllowedTypes:objcValue];
}

void* C_NSPathControl_ClickedPathItem(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSPathControlItem* result_ = [nSPathControl clickedPathItem];
    return result_;
}

bool C_NSPathControl_IsEditable(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    BOOL result_ = [nSPathControl isEditable];
    return result_;
}

void C_NSPathControl_SetEditable(void* ptr, bool value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setEditable:value];
}

Array C_NSPathControl_PathItems(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSArray* result_ = [nSPathControl pathItems];
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

void C_NSPathControl_SetPathItems(void* ptr, Array value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSPathControlItem*)(NSPathControlItem*)p];
    	}
    }
    [nSPathControl setPathItems:objcValue];
}

void* C_NSPathControl_PlaceholderAttributedString(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSAttributedString* result_ = [nSPathControl placeholderAttributedString];
    return result_;
}

void C_NSPathControl_SetPlaceholderAttributedString(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setPlaceholderAttributedString:(NSAttributedString*)value];
}

void* C_NSPathControl_PlaceholderString(void* ptr) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    NSString* result_ = [nSPathControl placeholderString];
    return result_;
}

void C_NSPathControl_SetPlaceholderString(void* ptr, void* value) {
    NSPathControl* nSPathControl = (NSPathControl*)ptr;
    [nSPathControl setPlaceholderString:(NSString*)value];
}
