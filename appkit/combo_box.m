#import "combo_box.h"
#import <AppKit/NSComboBox.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_ComboBox_Alloc() {
    return [NSComboBox alloc];
}

void* C_NSComboBox_ComboBox_LabelWithAttributedString(void* attributedStringValue) {
    NSComboBox* result_ = [NSComboBox labelWithAttributedString:(NSAttributedString*)attributedStringValue];
    return result_;
}

void* C_NSComboBox_ComboBox_LabelWithString(void* stringValue) {
    NSComboBox* result_ = [NSComboBox labelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSComboBox_ComboBox_TextFieldWithString(void* stringValue) {
    NSComboBox* result_ = [NSComboBox textFieldWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSComboBox_ComboBox_WrappingLabelWithString(void* stringValue) {
    NSComboBox* result_ = [NSComboBox wrappingLabelWithString:(NSString*)stringValue];
    return result_;
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

void* C_NSComboBox_AllocComboBox() {
    NSComboBox* result_ = [NSComboBox alloc];
    return result_;
}

void* C_NSComboBox_NewComboBox() {
    NSComboBox* result_ = [NSComboBox new];
    return result_;
}

void* C_NSComboBox_Autorelease(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result_ = [nSComboBox autorelease];
    return result_;
}

void* C_NSComboBox_Retain(void* ptr) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSComboBox* result_ = [nSComboBox retain];
    return result_;
}

void C_NSComboBox_AddItemsWithObjectValues(void* ptr, Array objects) {
    NSComboBox* nSComboBox = (NSComboBox*)ptr;
    NSMutableArray* objcObjects = [NSMutableArray arrayWithCapacity:objects.len];
    if (objects.len > 0) {
    	void** objectsData = (void**)objects.data;
    	for (int i = 0; i < objects.len; i++) {
    		void* p = objectsData[i];
    		[objcObjects addObject:(NSObject*)p];
    	}
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

@interface NSComboBoxDataSourceAdaptor : NSObject <NSComboBoxDataSource>
@property (assign) uintptr_t goID;
@end

@implementation NSComboBoxDataSourceAdaptor


- (NSString*)comboBox:(NSComboBox*)comboBox completedString:(NSString*)_string {
    void* result_ = comboBoxDataSource_ComboBox_CompletedString([self goID], comboBox, _string);
    return (NSString*)result_;
}

- (NSUInteger)comboBox:(NSComboBox*)comboBox indexOfItemWithStringValue:(NSString*)_string {
    unsigned int result_ = comboBoxDataSource_ComboBox_IndexOfItemWithStringValue([self goID], comboBox, _string);
    return result_;
}

- (id)comboBox:(NSComboBox*)comboBox objectValueForItemAtIndex:(NSInteger)index {
    void* result_ = comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex([self goID], comboBox, index);
    return (id)result_;
}

- (NSInteger)numberOfItemsInComboBox:(NSComboBox*)comboBox {
    int result_ = comboBoxDataSource_NumberOfItemsInComboBox([self goID], comboBox);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ComboBoxDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapComboBoxDataSource(uintptr_t goID) {
    NSComboBoxDataSourceAdaptor* adaptor = [[NSComboBoxDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

@interface NSComboBoxDelegateAdaptor : NSObject <NSComboBoxDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSComboBoxDelegateAdaptor


- (void)comboBoxSelectionDidChange:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxSelectionDidChange([self goID], notification);
}

- (void)comboBoxSelectionIsChanging:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxSelectionIsChanging([self goID], notification);
}

- (void)comboBoxWillDismiss:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxWillDismiss([self goID], notification);
}

- (void)comboBoxWillPopUp:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxWillPopUp([self goID], notification);
}

- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidates:(NSArray*)candidates forSelectedRange:(NSRange)selectedRange {
    Array candidatesArray;
    int candidatescount = [candidates count];
    if (candidatescount > 0) {
    	void** candidatesData = malloc(candidatescount * sizeof(void*));
    	for (int i = 0; i < candidatescount; i++) {
    		 void* p = [candidates objectAtIndex:i];
    		 candidatesData[i] = p;
    	}
    	candidatesArray.data = candidatesData;
    	candidatesArray.len = candidatescount;
    }
    Array result_ = comboBoxDelegate_TextField_TextView_Candidates_ForSelectedRange([self goID], textField, textView, candidatesArray, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSTextCheckingResult*)p];
    	}
    }
    return objcResult_;
}

- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidatesForSelectedRange:(NSRange)selectedRange {
    Array result_ = comboBoxDelegate_TextField_TextView_CandidatesForSelectedRange([self goID], textField, textView, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSObject*)p];
    	}
    }
    return objcResult_;
}

- (BOOL)textField:(NSTextField*)textField textView:(NSTextView*)textView shouldSelectCandidateAtIndex:(NSUInteger)index {
    bool result_ = comboBoxDelegate_TextField_TextView_ShouldSelectCandidateAtIndex([self goID], textField, textView, index);
    return result_;
}

- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = comboBoxDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    comboBoxDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = comboBoxDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = comboBoxDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = comboBoxDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = comboBoxDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ComboBoxDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapComboBoxDelegate(uintptr_t goID) {
    NSComboBoxDelegateAdaptor* adaptor = [[NSComboBoxDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
