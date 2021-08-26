#import <Appkit/Appkit.h>
#import "pop_up_button.h"

void* C_PopUpButton_Alloc() {
    return [NSPopUpButton alloc];
}

void* C_NSPopUpButton_InitWithFrame_PullsDown(void* ptr, CGRect buttonFrame, bool flag) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSPopUpButton* result_ = [nSPopUpButton initWithFrame:buttonFrame pullsDown:flag];
    return result_;
}

void* C_NSPopUpButton_InitWithCoder(void* ptr, void* coder) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSPopUpButton* result_ = [nSPopUpButton initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPopUpButton_Init(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSPopUpButton* result_ = [nSPopUpButton init];
    return result_;
}

void C_NSPopUpButton_AddItemWithTitle(void* ptr, void* title) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton addItemWithTitle:(NSString*)title];
}

void C_NSPopUpButton_AddItemsWithTitles(void* ptr, Array itemTitles) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSMutableArray* objcItemTitles = [[NSMutableArray alloc] init];
    if (itemTitles.len > 0) {
    	void** itemTitlesData = (void**)itemTitles.data;
    	for (int i = 0; i < itemTitles.len; i++) {
    		void* p = itemTitlesData[i];
    		[objcItemTitles addObject:(NSString*)(NSString*)p];
    	}
    }
    [nSPopUpButton addItemsWithTitles:objcItemTitles];
}

void C_NSPopUpButton_InsertItemWithTitle_AtIndex(void* ptr, void* title, int index) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton insertItemWithTitle:(NSString*)title atIndex:index];
}

void C_NSPopUpButton_RemoveAllItems(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton removeAllItems];
}

void C_NSPopUpButton_RemoveItemWithTitle(void* ptr, void* title) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton removeItemWithTitle:(NSString*)title];
}

void C_NSPopUpButton_RemoveItemAtIndex(void* ptr, int index) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton removeItemAtIndex:index];
}

void C_NSPopUpButton_SelectItem(void* ptr, void* item) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton selectItem:(NSMenuItem*)item];
}

void C_NSPopUpButton_SelectItemAtIndex(void* ptr, int index) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton selectItemAtIndex:index];
}

bool C_NSPopUpButton_SelectItemWithTag(void* ptr, int tag) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    BOOL result_ = [nSPopUpButton selectItemWithTag:tag];
    return result_;
}

void C_NSPopUpButton_SelectItemWithTitle(void* ptr, void* title) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton selectItemWithTitle:(NSString*)title];
}

void* C_NSPopUpButton_ItemAtIndex(void* ptr, int index) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSMenuItem* result_ = [nSPopUpButton itemAtIndex:index];
    return result_;
}

void* C_NSPopUpButton_ItemTitleAtIndex(void* ptr, int index) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSString* result_ = [nSPopUpButton itemTitleAtIndex:index];
    return result_;
}

void* C_NSPopUpButton_ItemWithTitle(void* ptr, void* title) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSMenuItem* result_ = [nSPopUpButton itemWithTitle:(NSString*)title];
    return result_;
}

int C_NSPopUpButton_IndexOfItem(void* ptr, void* item) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfItem:(NSMenuItem*)item];
    return result_;
}

int C_NSPopUpButton_IndexOfItemWithTag(void* ptr, int tag) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfItemWithTag:tag];
    return result_;
}

int C_NSPopUpButton_IndexOfItemWithTitle(void* ptr, void* title) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfItemWithTitle:(NSString*)title];
    return result_;
}

int C_NSPopUpButton_IndexOfItemWithRepresentedObject(void* ptr, void* obj) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfItemWithRepresentedObject:(id)obj];
    return result_;
}

int C_NSPopUpButton_IndexOfItemWithTarget_AndAction(void* ptr, void* target, void* actionSelector) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfItemWithTarget:(id)target andAction:(SEL)actionSelector];
    return result_;
}

void C_NSPopUpButton_SynchronizeTitleAndSelectedItem(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton synchronizeTitleAndSelectedItem];
}

bool C_NSPopUpButton_PullsDown(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    BOOL result_ = [nSPopUpButton pullsDown];
    return result_;
}

void C_NSPopUpButton_SetPullsDown(void* ptr, bool value) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton setPullsDown:value];
}

bool C_NSPopUpButton_AutoenablesItems(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    BOOL result_ = [nSPopUpButton autoenablesItems];
    return result_;
}

void C_NSPopUpButton_SetAutoenablesItems(void* ptr, bool value) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton setAutoenablesItems:value];
}

void* C_NSPopUpButton_SelectedItem(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSMenuItem* result_ = [nSPopUpButton selectedItem];
    return result_;
}

void* C_NSPopUpButton_TitleOfSelectedItem(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSString* result_ = [nSPopUpButton titleOfSelectedItem];
    return result_;
}

int C_NSPopUpButton_IndexOfSelectedItem(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton indexOfSelectedItem];
    return result_;
}

int C_NSPopUpButton_SelectedTag(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton selectedTag];
    return result_;
}

int C_NSPopUpButton_NumberOfItems(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSInteger result_ = [nSPopUpButton numberOfItems];
    return result_;
}

Array C_NSPopUpButton_ItemArray(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSArray* result_ = [nSPopUpButton itemArray];
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

Array C_NSPopUpButton_ItemTitles(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSArray* result_ = [nSPopUpButton itemTitles];
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

void* C_NSPopUpButton_LastItem(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSMenuItem* result_ = [nSPopUpButton lastItem];
    return result_;
}

unsigned int C_NSPopUpButton_PreferredEdge(void* ptr) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    NSRectEdge result_ = [nSPopUpButton preferredEdge];
    return result_;
}

void C_NSPopUpButton_SetPreferredEdge(void* ptr, unsigned int value) {
    NSPopUpButton* nSPopUpButton = (NSPopUpButton*)ptr;
    [nSPopUpButton setPreferredEdge:value];
}
