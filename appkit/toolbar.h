#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Toolbar_Alloc();

void* C_NSToolbar_InitWithIdentifier(void* ptr, void* identifier);
void* C_NSToolbar_Init(void* ptr);
void C_NSToolbar_InsertItemWithItemIdentifier_AtIndex(void* ptr, void* itemIdentifier, int index);
void C_NSToolbar_RemoveItemAtIndex(void* ptr, int index);
void C_NSToolbar_RunCustomizationPalette(void* ptr, void* sender);
void C_NSToolbar_ValidateVisibleItems(void* ptr);
void* C_NSToolbar_Identifier(void* ptr);
bool C_NSToolbar_ShowsBaselineSeparator(void* ptr);
void C_NSToolbar_SetShowsBaselineSeparator(void* ptr, bool value);
bool C_NSToolbar_AllowsUserCustomization(void* ptr);
void C_NSToolbar_SetAllowsUserCustomization(void* ptr, bool value);
bool C_NSToolbar_AllowsExtensionItems(void* ptr);
void C_NSToolbar_SetAllowsExtensionItems(void* ptr, bool value);
void* C_NSToolbar_SelectedItemIdentifier(void* ptr);
void C_NSToolbar_SetSelectedItemIdentifier(void* ptr, void* value);
void* C_NSToolbar_CenteredItemIdentifier(void* ptr);
void C_NSToolbar_SetCenteredItemIdentifier(void* ptr, void* value);
bool C_NSToolbar_IsVisible(void* ptr);
void C_NSToolbar_SetVisible(void* ptr, bool value);
bool C_NSToolbar_CustomizationPaletteIsRunning(void* ptr);
bool C_NSToolbar_AutosavesConfiguration(void* ptr);
void C_NSToolbar_SetAutosavesConfiguration(void* ptr, bool value);
