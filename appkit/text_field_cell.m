#import <Appkit/Appkit.h>
#import "text_field_cell.h"

void* C_TextFieldCell_Alloc() {
    return [NSTextFieldCell alloc];
}

void* C_NSTextFieldCell_InitTextCell(void* ptr, void* _string) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldCell* result_ = [nSTextFieldCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSTextFieldCell_InitWithCoder(void* ptr, void* coder) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldCell* result_ = [nSTextFieldCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTextFieldCell_Init(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldCell* result_ = [nSTextFieldCell init];
    return result_;
}

void* C_NSTextFieldCell_AllocTextFieldCell() {
    NSTextFieldCell* result_ = [NSTextFieldCell alloc];
    return result_;
}

void* C_NSTextFieldCell_NewTextFieldCell() {
    NSTextFieldCell* result_ = [NSTextFieldCell new];
    return result_;
}

void* C_NSTextFieldCell_Autorelease(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldCell* result_ = [nSTextFieldCell autorelease];
    return result_;
}

void* C_NSTextFieldCell_Retain(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldCell* result_ = [nSTextFieldCell retain];
    return result_;
}

void C_NSTextFieldCell_SetWantsNotificationForMarkedText(void* ptr, bool flag) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setWantsNotificationForMarkedText:flag];
}

void* C_NSTextFieldCell_TextColor(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSColor* result_ = [nSTextFieldCell textColor];
    return result_;
}

void C_NSTextFieldCell_SetTextColor(void* ptr, void* value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setTextColor:(NSColor*)value];
}

unsigned int C_NSTextFieldCell_BezelStyle(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSTextFieldBezelStyle result_ = [nSTextFieldCell bezelStyle];
    return result_;
}

void C_NSTextFieldCell_SetBezelStyle(void* ptr, unsigned int value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setBezelStyle:value];
}

void* C_NSTextFieldCell_BackgroundColor(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSColor* result_ = [nSTextFieldCell backgroundColor];
    return result_;
}

void C_NSTextFieldCell_SetBackgroundColor(void* ptr, void* value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setBackgroundColor:(NSColor*)value];
}

bool C_NSTextFieldCell_DrawsBackground(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    BOOL result_ = [nSTextFieldCell drawsBackground];
    return result_;
}

void C_NSTextFieldCell_SetDrawsBackground(void* ptr, bool value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setDrawsBackground:value];
}

void* C_NSTextFieldCell_PlaceholderString(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSString* result_ = [nSTextFieldCell placeholderString];
    return result_;
}

void C_NSTextFieldCell_SetPlaceholderString(void* ptr, void* value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setPlaceholderString:(NSString*)value];
}

void* C_NSTextFieldCell_PlaceholderAttributedString(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSAttributedString* result_ = [nSTextFieldCell placeholderAttributedString];
    return result_;
}

void C_NSTextFieldCell_SetPlaceholderAttributedString(void* ptr, void* value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    [nSTextFieldCell setPlaceholderAttributedString:(NSAttributedString*)value];
}

Array C_NSTextFieldCell_AllowedInputSourceLocales(void* ptr) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSArray* result_ = [nSTextFieldCell allowedInputSourceLocales];
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

void C_NSTextFieldCell_SetAllowedInputSourceLocales(void* ptr, Array value) {
    NSTextFieldCell* nSTextFieldCell = (NSTextFieldCell*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTextFieldCell setAllowedInputSourceLocales:objcValue];
}
