#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_WindowTabGroup_Alloc();

void* C_NSWindowTabGroup_AllocWindowTabGroup();
void* C_NSWindowTabGroup_Init(void* ptr);
void* C_NSWindowTabGroup_NewWindowTabGroup();
void* C_NSWindowTabGroup_Autorelease(void* ptr);
void* C_NSWindowTabGroup_Retain(void* ptr);
void C_NSWindowTabGroup_AddWindow(void* ptr, void* window);
void C_NSWindowTabGroup_InsertWindow_AtIndex(void* ptr, void* window, int index);
void C_NSWindowTabGroup_RemoveWindow(void* ptr, void* window);
void* C_NSWindowTabGroup_Identifier(void* ptr);
bool C_NSWindowTabGroup_IsOverviewVisible(void* ptr);
void C_NSWindowTabGroup_SetOverviewVisible(void* ptr, bool value);
bool C_NSWindowTabGroup_IsTabBarVisible(void* ptr);
Array C_NSWindowTabGroup_Windows(void* ptr);
void* C_NSWindowTabGroup_SelectedWindow(void* ptr);
void C_NSWindowTabGroup_SetSelectedWindow(void* ptr, void* value);
