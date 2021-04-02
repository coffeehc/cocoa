#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool OpenPanel_CanChooseFiles(void* ptr);
void OpenPanel_SetCanChooseFiles(void* ptr, bool canChooseFiles);
bool OpenPanel_CanChooseDirectories(void* ptr);
void OpenPanel_SetCanChooseDirectories(void* ptr, bool canChooseDirectories);
bool OpenPanel_ResolvesAliases(void* ptr);
void OpenPanel_SetResolvesAliases(void* ptr, bool resolvesAliases);
bool OpenPanel_AllowsMultipleSelection(void* ptr);
void OpenPanel_SetAllowsMultipleSelection(void* ptr, bool allowsMultipleSelection);
bool OpenPanel_IsAccessoryViewDisclosed(void* ptr);
void OpenPanel_SetAccessoryViewDisclosed(void* ptr, bool accessoryViewDisclosed);
Array OpenPanel_URLs(void* ptr);
bool OpenPanel_CanDownloadUbiquitousContents(void* ptr);
void OpenPanel_SetCanDownloadUbiquitousContents(void* ptr, bool canDownloadUbiquitousContents);
bool OpenPanel_CanResolveUbiquitousConflicts(void* ptr);
void OpenPanel_SetCanResolveUbiquitousConflicts(void* ptr, bool canResolveUbiquitousConflicts);

void* OpenPanel_NewOpenPanel();
