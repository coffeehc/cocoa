#import <AppKit/AppKit.h>
#import "combo_box.h"

bool ComboBox_HasVerticalScroller(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox hasVerticalScroller];
}

void ComboBox_SetHasVerticalScroller(void* ptr, bool hasVerticalScroller) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setHasVerticalScroller:hasVerticalScroller];
}

NSSize ComboBox_IntercellSpacing(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox intercellSpacing];
}

void ComboBox_SetIntercellSpacing(void* ptr, NSSize intercellSpacing) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setIntercellSpacing:intercellSpacing];
}

bool ComboBox_IsButtonBordered(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox isButtonBordered];
}

void ComboBox_SetButtonBordered(void* ptr, bool buttonBordered) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setButtonBordered:buttonBordered];
}

double ComboBox_ItemHeight(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox itemHeight];
}

void ComboBox_SetItemHeight(void* ptr, double itemHeight) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setItemHeight:itemHeight];
}

long ComboBox_NumberOfVisibleItems(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox numberOfVisibleItems];
}

void ComboBox_SetNumberOfVisibleItems(void* ptr, long numberOfVisibleItems) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setNumberOfVisibleItems:numberOfVisibleItems];
}

bool ComboBox_UsesDataSource(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox usesDataSource];
}

void ComboBox_SetUsesDataSource(void* ptr, bool usesDataSource) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setUsesDataSource:usesDataSource];
}

bool ComboBox_Completes(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox completes];
}

void ComboBox_SetCompletes(void* ptr, bool completes) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox setCompletes:completes];
}

Array ComboBox_ObjectValues(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	NSArray* ns_array = [comboBox objectValues];
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

long ComboBox_IndexOfSelectedItem(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox indexOfSelectedItem];
}

void* ComboBox_NewComboBox(NSRect frame) {
	NSComboBox* comboBox = [NSComboBox alloc];
	return [[comboBox initWithFrame:frame] autorelease];
}

void ComboBox_AddItemsWithObjectValues(void* ptr, Array objects) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
    NSMutableArray* objc_objects = [[NSMutableArray alloc] init];
    void** objectsData = (void**)objects.data;
    for (int i = 0; i < objects.len; i++) {
    	[objc_objects addObject:(NSObject*)objectsData[i]];
    }
	[comboBox addItemsWithObjectValues:objc_objects];
}

void ComboBox_AddItemWithObjectValue(void* ptr, void* object) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox addItemWithObjectValue:(NSObject*)object];
}

void ComboBox_InsertItemWithObjectValue(void* ptr, void* object, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox insertItemWithObjectValue:(NSObject*)object atIndex:index];
}

void ComboBox_RemoveAllItems(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox removeAllItems];
}

void ComboBox_RemoveItemAtIndex(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox removeItemAtIndex:index];
}

void ComboBox_RemoveItemWithObjectValue(void* ptr, void* object) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox removeItemWithObjectValue:(NSObject*)object];
}

long ComboBox_IndexOfItemWithObjectValue(void* ptr, void* object) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox indexOfItemWithObjectValue:(NSObject*)object];
}

void* ComboBox_ItemObjectValueAtIndex(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	return [comboBox itemObjectValueAtIndex:index];
}

void ComboBox_NoteNumberOfItemsChanged(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox noteNumberOfItemsChanged];
}

void ComboBox_ReloadData(void* ptr) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox reloadData];
}

void ComboBox_ScrollItemAtIndexToTop(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox scrollItemAtIndexToTop:index];
}

void ComboBox_ScrollItemAtIndexToVisible(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox scrollItemAtIndexToVisible:index];
}

void ComboBox_DeselectItemAtIndex(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox deselectItemAtIndex:index];
}

void ComboBox_SelectItemAtIndex(void* ptr, long index) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox selectItemAtIndex:index];
}

void ComboBox_SelectItemWithObjectValue(void* ptr, void* object) {
	NSComboBox* comboBox = (NSComboBox*)ptr;
	[comboBox selectItemWithObjectValue:(NSObject*)object];
}
