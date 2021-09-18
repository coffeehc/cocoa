#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ComboBox_Alloc();

void* C_NSComboBox_ComboBox_LabelWithAttributedString(void* attributedStringValue);
void* C_NSComboBox_ComboBox_LabelWithString(void* stringValue);
void* C_NSComboBox_ComboBox_TextFieldWithString(void* stringValue);
void* C_NSComboBox_ComboBox_WrappingLabelWithString(void* stringValue);
void* C_NSComboBox_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSComboBox_InitWithCoder(void* ptr, void* coder);
void* C_NSComboBox_Init(void* ptr);
void* C_NSComboBox_AllocComboBox();
void* C_NSComboBox_NewComboBox();
void* C_NSComboBox_Autorelease(void* ptr);
void* C_NSComboBox_Retain(void* ptr);
void C_NSComboBox_AddItemsWithObjectValues(void* ptr, Array objects);
void C_NSComboBox_AddItemWithObjectValue(void* ptr, void* object);
void C_NSComboBox_InsertItemWithObjectValue_AtIndex(void* ptr, void* object, int index);
void C_NSComboBox_RemoveAllItems(void* ptr);
void C_NSComboBox_RemoveItemAtIndex(void* ptr, int index);
void C_NSComboBox_RemoveItemWithObjectValue(void* ptr, void* object);
int C_NSComboBox_IndexOfItemWithObjectValue(void* ptr, void* object);
void* C_NSComboBox_ItemObjectValueAtIndex(void* ptr, int index);
void C_NSComboBox_NoteNumberOfItemsChanged(void* ptr);
void C_NSComboBox_ReloadData(void* ptr);
void C_NSComboBox_ScrollItemAtIndexToTop(void* ptr, int index);
void C_NSComboBox_ScrollItemAtIndexToVisible(void* ptr, int index);
void C_NSComboBox_DeselectItemAtIndex(void* ptr, int index);
void C_NSComboBox_SelectItemAtIndex(void* ptr, int index);
void C_NSComboBox_SelectItemWithObjectValue(void* ptr, void* object);
bool C_NSComboBox_HasVerticalScroller(void* ptr);
void C_NSComboBox_SetHasVerticalScroller(void* ptr, bool value);
CGSize C_NSComboBox_IntercellSpacing(void* ptr);
void C_NSComboBox_SetIntercellSpacing(void* ptr, CGSize value);
bool C_NSComboBox_IsButtonBordered(void* ptr);
void C_NSComboBox_SetButtonBordered(void* ptr, bool value);
double C_NSComboBox_ItemHeight(void* ptr);
void C_NSComboBox_SetItemHeight(void* ptr, double value);
int C_NSComboBox_NumberOfVisibleItems(void* ptr);
void C_NSComboBox_SetNumberOfVisibleItems(void* ptr, int value);
void* C_NSComboBox_DataSource(void* ptr);
void C_NSComboBox_SetDataSource(void* ptr, void* value);
bool C_NSComboBox_UsesDataSource(void* ptr);
void C_NSComboBox_SetUsesDataSource(void* ptr, bool value);
Array C_NSComboBox_ObjectValues(void* ptr);
int C_NSComboBox_NumberOfItems(void* ptr);
int C_NSComboBox_IndexOfSelectedItem(void* ptr);
void* C_NSComboBox_ObjectValueOfSelectedItem(void* ptr);
bool C_NSComboBox_Completes(void* ptr);
void C_NSComboBox_SetCompletes(void* ptr, bool value);
