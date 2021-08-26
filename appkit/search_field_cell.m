#import <Appkit/Appkit.h>
#import "search_field_cell.h"

void* C_SearchFieldCell_Alloc() {
    return [NSSearchFieldCell alloc];
}

void* C_NSSearchFieldCell_InitWithCoder(void* ptr, void* coder) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSSearchFieldCell* result_ = [nSSearchFieldCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSearchFieldCell_InitTextCell(void* ptr, void* _string) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSSearchFieldCell* result_ = [nSSearchFieldCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSSearchFieldCell_Init(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSSearchFieldCell* result_ = [nSSearchFieldCell init];
    return result_;
}

void C_NSSearchFieldCell_ResetSearchButtonCell(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell resetSearchButtonCell];
}

void C_NSSearchFieldCell_ResetCancelButtonCell(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell resetCancelButtonCell];
}

CGRect C_NSSearchFieldCell_SearchTextRectForBounds(void* ptr, CGRect rect) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSRect result_ = [nSSearchFieldCell searchTextRectForBounds:rect];
    return result_;
}

CGRect C_NSSearchFieldCell_SearchButtonRectForBounds(void* ptr, CGRect rect) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSRect result_ = [nSSearchFieldCell searchButtonRectForBounds:rect];
    return result_;
}

CGRect C_NSSearchFieldCell_CancelButtonRectForBounds(void* ptr, CGRect rect) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSRect result_ = [nSSearchFieldCell cancelButtonRectForBounds:rect];
    return result_;
}

void* C_NSSearchFieldCell_SearchButtonCell(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSButtonCell* result_ = [nSSearchFieldCell searchButtonCell];
    return result_;
}

void C_NSSearchFieldCell_SetSearchButtonCell(void* ptr, void* value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setSearchButtonCell:(NSButtonCell*)value];
}

void* C_NSSearchFieldCell_CancelButtonCell(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSButtonCell* result_ = [nSSearchFieldCell cancelButtonCell];
    return result_;
}

void C_NSSearchFieldCell_SetCancelButtonCell(void* ptr, void* value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setCancelButtonCell:(NSButtonCell*)value];
}

void* C_NSSearchFieldCell_SearchMenuTemplate(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSMenu* result_ = [nSSearchFieldCell searchMenuTemplate];
    return result_;
}

void C_NSSearchFieldCell_SetSearchMenuTemplate(void* ptr, void* value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setSearchMenuTemplate:(NSMenu*)value];
}

bool C_NSSearchFieldCell_SendsWholeSearchString(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    BOOL result_ = [nSSearchFieldCell sendsWholeSearchString];
    return result_;
}

void C_NSSearchFieldCell_SetSendsWholeSearchString(void* ptr, bool value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setSendsWholeSearchString:value];
}

bool C_NSSearchFieldCell_SendsSearchStringImmediately(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    BOOL result_ = [nSSearchFieldCell sendsSearchStringImmediately];
    return result_;
}

void C_NSSearchFieldCell_SetSendsSearchStringImmediately(void* ptr, bool value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setSendsSearchStringImmediately:value];
}

int C_NSSearchFieldCell_MaximumRecents(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSInteger result_ = [nSSearchFieldCell maximumRecents];
    return result_;
}

void C_NSSearchFieldCell_SetMaximumRecents(void* ptr, int value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setMaximumRecents:value];
}

Array C_NSSearchFieldCell_RecentSearches(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSArray* result_ = [nSSearchFieldCell recentSearches];
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

void C_NSSearchFieldCell_SetRecentSearches(void* ptr, Array value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)(NSString*)p];
    	}
    }
    [nSSearchFieldCell setRecentSearches:objcValue];
}

void* C_NSSearchFieldCell_RecentsAutosaveName(void* ptr) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    NSSearchFieldRecentsAutosaveName result_ = [nSSearchFieldCell recentsAutosaveName];
    return result_;
}

void C_NSSearchFieldCell_SetRecentsAutosaveName(void* ptr, void* value) {
    NSSearchFieldCell* nSSearchFieldCell = (NSSearchFieldCell*)ptr;
    [nSSearchFieldCell setRecentsAutosaveName:(NSString*)value];
}
