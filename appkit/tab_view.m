#import <Appkit/Appkit.h>
#import "tab_view.h"

void* C_TabView_Alloc() {
    return [NSTabView alloc];
}

void* C_NSTabView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result = [nSTabView initWithFrame:frameRect];
    return result;
}

void* C_NSTabView_InitWithCoder(void* ptr, void* coder) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result = [nSTabView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSTabView_Init(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabView* result = [nSTabView init];
    return result;
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
    NSInteger result = [nSTabView indexOfTabViewItem:(NSTabViewItem*)tabViewItem];
    return result;
}

int C_NSTabView_IndexOfTabViewItemWithIdentifier(void* ptr, void* identifier) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSInteger result = [nSTabView indexOfTabViewItemWithIdentifier:(id)identifier];
    return result;
}

void* C_NSTabView_TabViewItemAtIndex(void* ptr, int index) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewItem* result = [nSTabView tabViewItemAtIndex:index];
    return result;
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
    NSTabViewItem* result = [nSTabView tabViewItemAtPoint:point];
    return result;
}

int C_NSTabView_NumberOfTabViewItems(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSInteger result = [nSTabView numberOfTabViewItems];
    return result;
}

Array C_NSTabView_TabViewItems(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSArray* result = [nSTabView tabViewItems];
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
    NSTabViewType result = [nSTabView tabViewType];
    return result;
}

void C_NSTabView_SetTabViewType(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabViewType:value];
}

unsigned int C_NSTabView_TabPosition(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabPosition result = [nSTabView tabPosition];
    return result;
}

void C_NSTabView_SetTabPosition(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabPosition:value];
}

unsigned int C_NSTabView_TabViewBorderType(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewBorderType result = [nSTabView tabViewBorderType];
    return result;
}

void C_NSTabView_SetTabViewBorderType(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setTabViewBorderType:value];
}

void* C_NSTabView_SelectedTabViewItem(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSTabViewItem* result = [nSTabView selectedTabViewItem];
    return result;
}

void* C_NSTabView_Font(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSFont* result = [nSTabView font];
    return result;
}

void C_NSTabView_SetFont(void* ptr, void* value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setFont:(NSFont*)value];
}

bool C_NSTabView_DrawsBackground(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    BOOL result = [nSTabView drawsBackground];
    return result;
}

void C_NSTabView_SetDrawsBackground(void* ptr, bool value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setDrawsBackground:value];
}

CGSize C_NSTabView_MinimumSize(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSSize result = [nSTabView minimumSize];
    return result;
}

CGRect C_NSTabView_ContentRect(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSRect result = [nSTabView contentRect];
    return result;
}

unsigned int C_NSTabView_ControlSize(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    NSControlSize result = [nSTabView controlSize];
    return result;
}

void C_NSTabView_SetControlSize(void* ptr, unsigned int value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setControlSize:value];
}

bool C_NSTabView_AllowsTruncatedLabels(void* ptr) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    BOOL result = [nSTabView allowsTruncatedLabels];
    return result;
}

void C_NSTabView_SetAllowsTruncatedLabels(void* ptr, bool value) {
    NSTabView* nSTabView = (NSTabView*)ptr;
    [nSTabView setAllowsTruncatedLabels:value];
}
