#import <Appkit/Appkit.h>
#import "combo_box.h"

void* C_ComboBox_Alloc() {
    return [NSComboBox alloc];
}

void* C_NSComboBox_InitWithFrame(void* ptr, CGRect frameRect) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result_ = [nSComboBox initWithFrame:frameRect];
    return result_;
}

void* C_NSComboBox_InitWithCoder(void* ptr, void* coder) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result_ = [nSComboBox initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSComboBox_Init(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result_ = [nSComboBox init];
    return result_;
}

void C_NSComboBox_AddItemsWithObjectValues(void* ptr, Array objects) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSMutableArray* objcObjects = [[NSMutableArray alloc] init];
    void** objectsData = (void**)objects.data;
    for (int i = 0; i < objects.len; i++) {
    	void* p = objectsData[i];
    	[objcObjects addObject:(NSObject*)(NSObject*)p];
    }
    [nSComboBox addItemsWithObjectValues:objcObjects];
}

void C_NSComboBox_AddItemWithObjectValue(void* ptr, void* object) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox addItemWithObjectValue:(id)object];
}

void C_NSComboBox_InsertItemWithObjectValue_AtIndex(void* ptr, void* object, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox insertItemWithObjectValue:(id)object atIndex:index];
}

void C_NSComboBox_RemoveAllItems(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox removeAllItems];
}

void C_NSComboBox_RemoveItemAtIndex(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox removeItemAtIndex:index];
}

void C_NSComboBox_RemoveItemWithObjectValue(void* ptr, void* object) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox removeItemWithObjectValue:(id)object];
}

int C_NSComboBox_IndexOfItemWithObjectValue(void* ptr, void* object) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result_ = [nSComboBox indexOfItemWithObjectValue:(id)object];
    return result_;
}

void* C_NSComboBox_ItemObjectValueAtIndex(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    id result_ = [nSComboBox itemObjectValueAtIndex:index];
    return result_;
}

void C_NSComboBox_NoteNumberOfItemsChanged(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox noteNumberOfItemsChanged];
}

void C_NSComboBox_ReloadData(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox reloadData];
}

void C_NSComboBox_ScrollItemAtIndexToTop(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox scrollItemAtIndexToTop:index];
}

void C_NSComboBox_ScrollItemAtIndexToVisible(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox scrollItemAtIndexToVisible:index];
}

void C_NSComboBox_DeselectItemAtIndex(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox deselectItemAtIndex:index];
}

void C_NSComboBox_SelectItemAtIndex(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox selectItemAtIndex:index];
}

void C_NSComboBox_SelectItemWithObjectValue(void* ptr, void* object) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox selectItemWithObjectValue:(id)object];
}

bool C_NSComboBox_HasVerticalScroller(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result_ = [nSComboBox hasVerticalScroller];
    return result_;
}

void C_NSComboBox_SetHasVerticalScroller(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setHasVerticalScroller:value];
}

CGSize C_NSComboBox_IntercellSpacing(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSSize result_ = [nSComboBox intercellSpacing];
    return result_;
}

void C_NSComboBox_SetIntercellSpacing(void* ptr, CGSize value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setIntercellSpacing:value];
}

bool C_NSComboBox_IsButtonBordered(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result_ = [nSComboBox isButtonBordered];
    return result_;
}

void C_NSComboBox_SetButtonBordered(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setButtonBordered:value];
}

double C_NSComboBox_ItemHeight(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    CGFloat result_ = [nSComboBox itemHeight];
    return result_;
}

void C_NSComboBox_SetItemHeight(void* ptr, double value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setItemHeight:value];
}

int C_NSComboBox_NumberOfVisibleItems(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result_ = [nSComboBox numberOfVisibleItems];
    return result_;
}

void C_NSComboBox_SetNumberOfVisibleItems(void* ptr, int value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setNumberOfVisibleItems:value];
}

void* C_NSComboBox_DataSource(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    id result_ = [nSComboBox dataSource];
    return result_;
}

void C_NSComboBox_SetDataSource(void* ptr, void* value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setDataSource:(id)value];
}

bool C_NSComboBox_UsesDataSource(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result_ = [nSComboBox usesDataSource];
    return result_;
}

void C_NSComboBox_SetUsesDataSource(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setUsesDataSource:value];
}

Array C_NSComboBox_ObjectValues(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSArray* result_ = [nSComboBox objectValues];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

int C_NSComboBox_NumberOfItems(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result_ = [nSComboBox numberOfItems];
    return result_;
}

int C_NSComboBox_IndexOfSelectedItem(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result_ = [nSComboBox indexOfSelectedItem];
    return result_;
}

void* C_NSComboBox_ObjectValueOfSelectedItem(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    id result_ = [nSComboBox objectValueOfSelectedItem];
    return result_;
}

bool C_NSComboBox_Completes(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result_ = [nSComboBox completes];
    return result_;
}

void C_NSComboBox_SetCompletes(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setCompletes:value];
}
