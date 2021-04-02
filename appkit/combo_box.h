#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool ComboBox_HasVerticalScroller(void* ptr);
void ComboBox_SetHasVerticalScroller(void* ptr, bool hasVerticalScroller);
NSSize ComboBox_IntercellSpacing(void* ptr);
void ComboBox_SetIntercellSpacing(void* ptr, NSSize intercellSpacing);
bool ComboBox_IsButtonBordered(void* ptr);
void ComboBox_SetButtonBordered(void* ptr, bool buttonBordered);
double ComboBox_ItemHeight(void* ptr);
void ComboBox_SetItemHeight(void* ptr, double itemHeight);
long ComboBox_NumberOfVisibleItems(void* ptr);
void ComboBox_SetNumberOfVisibleItems(void* ptr, long numberOfVisibleItems);
bool ComboBox_UsesDataSource(void* ptr);
void ComboBox_SetUsesDataSource(void* ptr, bool usesDataSource);
bool ComboBox_Completes(void* ptr);
void ComboBox_SetCompletes(void* ptr, bool completes);
Array ComboBox_ObjectValues(void* ptr);
long ComboBox_IndexOfSelectedItem(void* ptr);

void* ComboBox_NewComboBox(NSRect frame);
void ComboBox_AddItemsWithObjectValues(void* ptr, Array objects);
void ComboBox_AddItemWithObjectValue(void* ptr, void* object);
void ComboBox_InsertItemWithObjectValue(void* ptr, void* object, long index);
void ComboBox_RemoveAllItems(void* ptr);
void ComboBox_RemoveItemAtIndex(void* ptr, long index);
void ComboBox_RemoveItemWithObjectValue(void* ptr, void* object);
long ComboBox_IndexOfItemWithObjectValue(void* ptr, void* object);
void* ComboBox_ItemObjectValueAtIndex(void* ptr, long index);
void ComboBox_NoteNumberOfItemsChanged(void* ptr);
void ComboBox_ReloadData(void* ptr);
void ComboBox_ScrollItemAtIndexToTop(void* ptr, long index);
void ComboBox_ScrollItemAtIndexToVisible(void* ptr, long index);
void ComboBox_DeselectItemAtIndex(void* ptr, long index);
void ComboBox_SelectItemAtIndex(void* ptr, long index);
void ComboBox_SelectItemWithObjectValue(void* ptr, void* object);
