#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_OpenPanel_Alloc();

void* C_NSOpenPanel_OpenPanel_WindowWithContentViewController(void* contentViewController);
void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSOpenPanel_Init(void* ptr);
void* C_NSOpenPanel_AllocOpenPanel();
void* C_NSOpenPanel_NewOpenPanel();
void* C_NSOpenPanel_Autorelease(void* ptr);
void* C_NSOpenPanel_Retain(void* ptr);
void* C_NSOpenPanel_OpenPanel_();
bool C_NSOpenPanel_CanChooseFiles(void* ptr);
void C_NSOpenPanel_SetCanChooseFiles(void* ptr, bool value);
bool C_NSOpenPanel_CanChooseDirectories(void* ptr);
void C_NSOpenPanel_SetCanChooseDirectories(void* ptr, bool value);
bool C_NSOpenPanel_ResolvesAliases(void* ptr);
void C_NSOpenPanel_SetResolvesAliases(void* ptr, bool value);
bool C_NSOpenPanel_AllowsMultipleSelection(void* ptr);
void C_NSOpenPanel_SetAllowsMultipleSelection(void* ptr, bool value);
bool C_NSOpenPanel_IsAccessoryViewDisclosed(void* ptr);
void C_NSOpenPanel_SetAccessoryViewDisclosed(void* ptr, bool value);
Array C_NSOpenPanel_URLs(void* ptr);
bool C_NSOpenPanel_CanDownloadUbiquitousContents(void* ptr);
void C_NSOpenPanel_SetCanDownloadUbiquitousContents(void* ptr, bool value);
bool C_NSOpenPanel_CanResolveUbiquitousConflicts(void* ptr);
void C_NSOpenPanel_SetCanResolveUbiquitousConflicts(void* ptr, bool value);
