#import <Appkit/Appkit.h>
#import "combo_box.h"

void* C_ComboBox_Alloc() {
    return [NSComboBox alloc];
}

void* C_NSComboBox_InitWithFrame(void* ptr, CGRect frameRect) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result = [nSComboBox initWithFrame:frameRect];
    return result;
}

void* C_NSComboBox_InitWithCoder(void* ptr, void* coder) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result = [nSComboBox initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSComboBox_Init(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result = [nSComboBox init];
    return result;
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
    NSInteger result = [nSComboBox indexOfItemWithObjectValue:(id)object];
    return result;
}

void* C_NSComboBox_ItemObjectValueAtIndex(void* ptr, int index) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    id result = [nSComboBox itemObjectValueAtIndex:index];
    return result;
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
    BOOL result = [nSComboBox hasVerticalScroller];
    return result;
}

void C_NSComboBox_SetHasVerticalScroller(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setHasVerticalScroller:value];
}

CGSize C_NSComboBox_IntercellSpacing(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSSize result = [nSComboBox intercellSpacing];
    return result;
}

void C_NSComboBox_SetIntercellSpacing(void* ptr, CGSize value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setIntercellSpacing:value];
}

bool C_NSComboBox_IsButtonBordered(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result = [nSComboBox isButtonBordered];
    return result;
}

void C_NSComboBox_SetButtonBordered(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setButtonBordered:value];
}

double C_NSComboBox_ItemHeight(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    CGFloat result = [nSComboBox itemHeight];
    return result;
}

void C_NSComboBox_SetItemHeight(void* ptr, double value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setItemHeight:value];
}

int C_NSComboBox_NumberOfVisibleItems(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result = [nSComboBox numberOfVisibleItems];
    return result;
}

void C_NSComboBox_SetNumberOfVisibleItems(void* ptr, int value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setNumberOfVisibleItems:value];
}

bool C_NSComboBox_UsesDataSource(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result = [nSComboBox usesDataSource];
    return result;
}

void C_NSComboBox_SetUsesDataSource(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setUsesDataSource:value];
}

Array C_NSComboBox_ObjectValues(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSArray* result = [nSComboBox objectValues];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

int C_NSComboBox_NumberOfItems(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result = [nSComboBox numberOfItems];
    return result;
}

int C_NSComboBox_IndexOfSelectedItem(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSInteger result = [nSComboBox indexOfSelectedItem];
    return result;
}

void* C_NSComboBox_ObjectValueOfSelectedItem(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    id result = [nSComboBox objectValueOfSelectedItem];
    return result;
}

bool C_NSComboBox_Completes(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    BOOL result = [nSComboBox completes];
    return result;
}

void C_NSComboBox_SetCompletes(void* ptr, bool value) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    [nSComboBox setCompletes:value];
}
