#import <Appkit/Appkit.h>
#import "search_field.h"

void* C_SearchField_Alloc() {
    return [NSSearchField alloc];
}

void* C_NSSearchField_SearchField_LabelWithAttributedString(void* attributedStringValue) {
    NSSearchField* result_ = [NSSearchField labelWithAttributedString:(NSAttributedString*)attributedStringValue];
    return result_;
}

void* C_NSSearchField_SearchField_LabelWithString(void* stringValue) {
    NSSearchField* result_ = [NSSearchField labelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSearchField_SearchField_TextFieldWithString(void* stringValue) {
    NSSearchField* result_ = [NSSearchField textFieldWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSearchField_SearchField_WrappingLabelWithString(void* stringValue) {
    NSSearchField* result_ = [NSSearchField wrappingLabelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSearchField_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchField* result_ = [nSSearchField initWithFrame:frameRect];
    return result_;
}

void* C_NSSearchField_InitWithCoder(void* ptr, void* coder) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchField* result_ = [nSSearchField initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSearchField_Init(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchField* result_ = [nSSearchField init];
    return result_;
}

void* C_NSSearchField_AllocSearchField() {
    NSSearchField* result_ = [NSSearchField alloc];
    return result_;
}

void* C_NSSearchField_NewSearchField() {
    NSSearchField* result_ = [NSSearchField new];
    return result_;
}

void* C_NSSearchField_Autorelease(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchField* result_ = [nSSearchField autorelease];
    return result_;
}

void* C_NSSearchField_Retain(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchField* result_ = [nSSearchField retain];
    return result_;
}

void* C_NSSearchField_SearchMenuTemplate(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSMenu* result_ = [nSSearchField searchMenuTemplate];
    return result_;
}

void C_NSSearchField_SetSearchMenuTemplate(void* ptr, void* value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    [nSSearchField setSearchMenuTemplate:(NSMenu*)value];
}

bool C_NSSearchField_SendsSearchStringImmediately(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    BOOL result_ = [nSSearchField sendsSearchStringImmediately];
    return result_;
}

void C_NSSearchField_SetSendsSearchStringImmediately(void* ptr, bool value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    [nSSearchField setSendsSearchStringImmediately:value];
}

bool C_NSSearchField_SendsWholeSearchString(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    BOOL result_ = [nSSearchField sendsWholeSearchString];
    return result_;
}

void C_NSSearchField_SetSendsWholeSearchString(void* ptr, bool value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    [nSSearchField setSendsWholeSearchString:value];
}

Array C_NSSearchField_RecentSearches(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSArray* result_ = [nSSearchField recentSearches];
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

void C_NSSearchField_SetRecentSearches(void* ptr, Array value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSSearchField setRecentSearches:objcValue];
}

int C_NSSearchField_MaximumRecents(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSInteger result_ = [nSSearchField maximumRecents];
    return result_;
}

void C_NSSearchField_SetMaximumRecents(void* ptr, int value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    [nSSearchField setMaximumRecents:value];
}

void* C_NSSearchField_RecentsAutosaveName(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSSearchFieldRecentsAutosaveName result_ = [nSSearchField recentsAutosaveName];
    return result_;
}

void C_NSSearchField_SetRecentsAutosaveName(void* ptr, void* value) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    [nSSearchField setRecentsAutosaveName:(NSString*)value];
}

CGRect C_NSSearchField_CancelButtonBounds(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSRect result_ = [nSSearchField cancelButtonBounds];
    return result_;
}

CGRect C_NSSearchField_SearchButtonBounds(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSRect result_ = [nSSearchField searchButtonBounds];
    return result_;
}

CGRect C_NSSearchField_SearchTextBounds(void* ptr) {
    NSSearchField* nSSearchField = (NSSearchField*)ptr;
    NSRect result_ = [nSSearchField searchTextBounds];
    return result_;
}
