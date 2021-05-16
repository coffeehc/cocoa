#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_SavePanel_Alloc();

void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSSavePanel_Init(void* ptr);
void* C_NSSavePanel_SavePanel_();
int C_NSSavePanel_RunModal(void* ptr);
void C_NSSavePanel_ValidateVisibleColumns(void* ptr);
void C_NSSavePanel_Ok(void* ptr, void* sender);
void C_NSSavePanel_Cancel(void* ptr, void* sender);
void* C_NSSavePanel_URL(void* ptr);
void* C_NSSavePanel_Prompt(void* ptr);
void C_NSSavePanel_SetPrompt(void* ptr, void* value);
void* C_NSSavePanel_Message(void* ptr);
void C_NSSavePanel_SetMessage(void* ptr, void* value);
void* C_NSSavePanel_NameFieldLabel(void* ptr);
void C_NSSavePanel_SetNameFieldLabel(void* ptr, void* value);
void* C_NSSavePanel_NameFieldStringValue(void* ptr);
void C_NSSavePanel_SetNameFieldStringValue(void* ptr, void* value);
void* C_NSSavePanel_DirectoryURL(void* ptr);
void C_NSSavePanel_SetDirectoryURL(void* ptr, void* value);
void* C_NSSavePanel_AccessoryView(void* ptr);
void C_NSSavePanel_SetAccessoryView(void* ptr, void* value);
bool C_NSSavePanel_ShowsTagField(void* ptr);
void C_NSSavePanel_SetShowsTagField(void* ptr, bool value);
Array C_NSSavePanel_TagNames(void* ptr);
void C_NSSavePanel_SetTagNames(void* ptr, Array value);
bool C_NSSavePanel_CanCreateDirectories(void* ptr);
void C_NSSavePanel_SetCanCreateDirectories(void* ptr, bool value);
bool C_NSSavePanel_CanSelectHiddenExtension(void* ptr);
void C_NSSavePanel_SetCanSelectHiddenExtension(void* ptr, bool value);
bool C_NSSavePanel_ShowsHiddenFiles(void* ptr);
void C_NSSavePanel_SetShowsHiddenFiles(void* ptr, bool value);
bool C_NSSavePanel_IsExtensionHidden(void* ptr);
void C_NSSavePanel_SetExtensionHidden(void* ptr, bool value);
bool C_NSSavePanel_IsExpanded(void* ptr);
Array C_NSSavePanel_AllowedContentTypes(void* ptr);
void C_NSSavePanel_SetAllowedContentTypes(void* ptr, Array value);
bool C_NSSavePanel_AllowsOtherFileTypes(void* ptr);
void C_NSSavePanel_SetAllowsOtherFileTypes(void* ptr, bool value);
bool C_NSSavePanel_TreatsFilePackagesAsDirectories(void* ptr);
void C_NSSavePanel_SetTreatsFilePackagesAsDirectories(void* ptr, bool value);
