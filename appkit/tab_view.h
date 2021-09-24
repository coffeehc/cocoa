#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TabView_Alloc();

void* C_NSTabView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTabView_InitWithCoder(void* ptr, void* coder);
void* C_NSTabView_Init(void* ptr);
void* C_NSTabView_AllocTabView();
void* C_NSTabView_NewTabView();
void* C_NSTabView_Autorelease(void* ptr);
void* C_NSTabView_Retain(void* ptr);
void C_NSTabView_AddTabViewItem(void* ptr, void* tabViewItem);
void C_NSTabView_InsertTabViewItem_AtIndex(void* ptr, void* tabViewItem, int index);
void C_NSTabView_RemoveTabViewItem(void* ptr, void* tabViewItem);
int C_NSTabView_IndexOfTabViewItem(void* ptr, void* tabViewItem);
int C_NSTabView_IndexOfTabViewItemWithIdentifier(void* ptr, void* identifier);
void* C_NSTabView_TabViewItemAtIndex(void* ptr, int index);
void C_NSTabView_SelectFirstTabViewItem(void* ptr, void* sender);
void C_NSTabView_SelectLastTabViewItem(void* ptr, void* sender);
void C_NSTabView_SelectNextTabViewItem(void* ptr, void* sender);
void C_NSTabView_SelectPreviousTabViewItem(void* ptr, void* sender);
void C_NSTabView_SelectTabViewItem(void* ptr, void* tabViewItem);
void C_NSTabView_SelectTabViewItemAtIndex(void* ptr, int index);
void C_NSTabView_SelectTabViewItemWithIdentifier(void* ptr, void* identifier);
void C_NSTabView_TakeSelectedTabViewItemFromSender(void* ptr, void* sender);
void* C_NSTabView_TabViewItemAtPoint(void* ptr, CGPoint point);
void* C_NSTabView_Delegate(void* ptr);
void C_NSTabView_SetDelegate(void* ptr, void* value);
int C_NSTabView_NumberOfTabViewItems(void* ptr);
Array C_NSTabView_TabViewItems(void* ptr);
void C_NSTabView_SetTabViewItems(void* ptr, Array value);
unsigned int C_NSTabView_TabViewType(void* ptr);
void C_NSTabView_SetTabViewType(void* ptr, unsigned int value);
unsigned int C_NSTabView_TabPosition(void* ptr);
void C_NSTabView_SetTabPosition(void* ptr, unsigned int value);
unsigned int C_NSTabView_TabViewBorderType(void* ptr);
void C_NSTabView_SetTabViewBorderType(void* ptr, unsigned int value);
void* C_NSTabView_SelectedTabViewItem(void* ptr);
void* C_NSTabView_Font(void* ptr);
void C_NSTabView_SetFont(void* ptr, void* value);
bool C_NSTabView_DrawsBackground(void* ptr);
void C_NSTabView_SetDrawsBackground(void* ptr, bool value);
CGSize C_NSTabView_MinimumSize(void* ptr);
CGRect C_NSTabView_ContentRect(void* ptr);
unsigned int C_NSTabView_ControlSize(void* ptr);
void C_NSTabView_SetControlSize(void* ptr, unsigned int value);
bool C_NSTabView_AllowsTruncatedLabels(void* ptr);
void C_NSTabView_SetAllowsTruncatedLabels(void* ptr, bool value);

void* WrapTabViewDelegate(uintptr_t goID);
