#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* SavePanel_URL(void* ptr);
const char* SavePanel_Prompt(void* ptr);
void SavePanel_SetPrompt(void* ptr, const char* prompt);
const char* SavePanel_Message(void* ptr);
void SavePanel_SetMessage(void* ptr, const char* message);
const char* SavePanel_NameFieldLabel(void* ptr);
void SavePanel_SetNameFieldLabel(void* ptr, const char* nameFieldLabel);
void* SavePanel_DirectoryURL(void* ptr);
void SavePanel_SetDirectoryURL(void* ptr, void* directoryURL);
void* SavePanel_AccessoryView(void* ptr);
void SavePanel_SetAccessoryView(void* ptr, void* accessoryView);
bool SavePanel_ShowsTagField(void* ptr);
void SavePanel_SetShowsTagField(void* ptr, bool showsTagField);
Array SavePanel_TagNames(void* ptr);
void SavePanel_SetTagNames(void* ptr, Array tagNames);
bool SavePanel_CanCreateDirectories(void* ptr);
void SavePanel_SetCanCreateDirectories(void* ptr, bool canCreateDirectories);
bool SavePanel_CanSelectHiddenExtension(void* ptr);
void SavePanel_SetCanSelectHiddenExtension(void* ptr, bool canSelectHiddenExtension);
bool SavePanel_ShowsHiddenFiles(void* ptr);
void SavePanel_SetShowsHiddenFiles(void* ptr, bool showsHiddenFiles);
bool SavePanel_IsExtensionHidden(void* ptr);
void SavePanel_SetExtensionHidden(void* ptr, bool extensionHidden);
bool SavePanel_IsExpanded(void* ptr);
Array SavePanel_AllowedContentTypes(void* ptr);
void SavePanel_SetAllowedContentTypes(void* ptr, Array allowedContentTypes);
bool SavePanel_AllowsOtherFileTypes(void* ptr);
void SavePanel_SetAllowsOtherFileTypes(void* ptr, bool allowsOtherFileTypes);
bool SavePanel_TreatsFilePackagesAsDirectories(void* ptr);
void SavePanel_SetTreatsFilePackagesAsDirectories(void* ptr, bool treatsFilePackagesAsDirectories);

void* SavePanel_NewSavePanel();
unsigned long SavePanel_RunModal(void* ptr);
void SavePanel_ValidateVisibleColumns(void* ptr);

void SavePanel_BeginSheetModalForWindow(void* ptr, void* window, long handlerID);
void SavePanel_BeginWithCompletionHandler(void* ptr, long handlerID);

