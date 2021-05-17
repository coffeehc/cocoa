#import <Appkit/Appkit.h>
#import "tab_view.h"

void* C_TabView_Alloc() {
    return [NSTabView alloc];
}

void* C_NSTabView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result_ = [nSTabView initWithFrame:frameRect];
    return result_;
}

void* C_NSTabView_InitWithCoder(void* ptr, void* coder) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result_ = [nSTabView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTabView_Init(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result_ = [nSTabView init];
    return result_;
}

void C_NSTabView_AddTabViewItem(void* ptr, void* tabViewItem) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView addTabViewItem:(NSTabViewItem*)tabViewItem];
}

void C_NSTabView_InsertTabViewItem_AtIndex(void* ptr, void* tabViewItem, int index) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView insertTabViewItem:(NSTabViewItem*)tabViewItem atIndex:index];
}

void C_NSTabView_RemoveTabViewItem(void* ptr, void* tabViewItem) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView removeTabViewItem:(NSTabViewItem*)tabViewItem];
}

int C_NSTabView_IndexOfTabViewItem(void* ptr, void* tabViewItem) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSInteger result_ = [nSTabView indexOfTabViewItem:(NSTabViewItem*)tabViewItem];
    return result_;
}

int C_NSTabView_IndexOfTabViewItemWithIdentifier(void* ptr, void* identifier) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSInteger result_ = [nSTabView indexOfTabViewItemWithIdentifier:(id)identifier];
    return result_;
}

void* C_NSTabView_TabViewItemAtIndex(void* ptr, int index) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewItem* result_ = [nSTabView tabViewItemAtIndex:index];
    return result_;
}

void C_NSTabView_SelectFirstTabViewItem(void* ptr, void* sender) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectFirstTabViewItem:(id)sender];
}

void C_NSTabView_SelectLastTabViewItem(void* ptr, void* sender) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectLastTabViewItem:(id)sender];
}

void C_NSTabView_SelectNextTabViewItem(void* ptr, void* sender) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectNextTabViewItem:(id)sender];
}

void C_NSTabView_SelectPreviousTabViewItem(void* ptr, void* sender) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectPreviousTabViewItem:(id)sender];
}

void C_NSTabView_SelectTabViewItem(void* ptr, void* tabViewItem) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectTabViewItem:(NSTabViewItem*)tabViewItem];
}

void C_NSTabView_SelectTabViewItemAtIndex(void* ptr, int index) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectTabViewItemAtIndex:index];
}

void C_NSTabView_SelectTabViewItemWithIdentifier(void* ptr, void* identifier) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView selectTabViewItemWithIdentifier:(id)identifier];
}

void C_NSTabView_TakeSelectedTabViewItemFromSender(void* ptr, void* sender) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView takeSelectedTabViewItemFromSender:(id)sender];
}

void* C_NSTabView_TabViewItemAtPoint(void* ptr, CGPoint point) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewItem* result_ = [nSTabView tabViewItemAtPoint:point];
    return result_;
}

void* C_NSTabView_Delegate(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    id result_ = [nSTabView delegate];
    return result_;
}

void C_NSTabView_SetDelegate(void* ptr, void* value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setDelegate:(id)value];
}

int C_NSTabView_NumberOfTabViewItems(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSInteger result_ = [nSTabView numberOfTabViewItems];
    return result_;
}

Array C_NSTabView_TabViewItems(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSArray* result_ = [nSTabView tabViewItems];
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

void C_NSTabView_SetTabViewItems(void* ptr, Array value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTabViewItem*)(NSTabViewItem*)p];
    }
    [nSTabView setTabViewItems:objcValue];
}

unsigned int C_NSTabView_TabViewType(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewType result_ = [nSTabView tabViewType];
    return result_;
}

void C_NSTabView_SetTabViewType(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabViewType:value];
}

unsigned int C_NSTabView_TabPosition(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabPosition result_ = [nSTabView tabPosition];
    return result_;
}

void C_NSTabView_SetTabPosition(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabPosition:value];
}

unsigned int C_NSTabView_TabViewBorderType(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewBorderType result_ = [nSTabView tabViewBorderType];
    return result_;
}

void C_NSTabView_SetTabViewBorderType(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabViewBorderType:value];
}

void* C_NSTabView_SelectedTabViewItem(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewItem* result_ = [nSTabView selectedTabViewItem];
    return result_;
}

void* C_NSTabView_Font(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSFont* result_ = [nSTabView font];
    return result_;
}

void C_NSTabView_SetFont(void* ptr, void* value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setFont:(NSFont*)value];
}

bool C_NSTabView_DrawsBackground(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    BOOL result_ = [nSTabView drawsBackground];
    return result_;
}

void C_NSTabView_SetDrawsBackground(void* ptr, bool value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setDrawsBackground:value];
}

CGSize C_NSTabView_MinimumSize(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSSize result_ = [nSTabView minimumSize];
    return result_;
}

CGRect C_NSTabView_ContentRect(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSRect result_ = [nSTabView contentRect];
    return result_;
}

unsigned int C_NSTabView_ControlSize(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSControlSize result_ = [nSTabView controlSize];
    return result_;
}

void C_NSTabView_SetControlSize(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setControlSize:value];
}

bool C_NSTabView_AllowsTruncatedLabels(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    BOOL result_ = [nSTabView allowsTruncatedLabels];
    return result_;
}

void C_NSTabView_SetAllowsTruncatedLabels(void* ptr, bool value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setAllowsTruncatedLabels:value];
}
