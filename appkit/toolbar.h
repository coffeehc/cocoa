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
void C_NSToolbar_SetConfigurationFromDictionary(void* ptr, Dictionary configDict);
void C_NSToolbar_ValidateVisibleItems(void* ptr);
void* C_NSToolbar_Delegate(void* ptr);
void C_NSToolbar_SetDelegate(void* ptr, void* value);
void* C_NSToolbar_Identifier(void* ptr);
unsigned int C_NSToolbar_DisplayMode(void* ptr);
void C_NSToolbar_SetDisplayMode(void* ptr, unsigned int value);
bool C_NSToolbar_ShowsBaselineSeparator(void* ptr);
void C_NSToolbar_SetShowsBaselineSeparator(void* ptr, bool value);
bool C_NSToolbar_AllowsUserCustomization(void* ptr);
void C_NSToolbar_SetAllowsUserCustomization(void* ptr, bool value);
bool C_NSToolbar_AllowsExtensionItems(void* ptr);
void C_NSToolbar_SetAllowsExtensionItems(void* ptr, bool value);
Array C_NSToolbar_Items(void* ptr);
Array C_NSToolbar_VisibleItems(void* ptr);
unsigned int C_NSToolbar_SizeMode(void* ptr);
void C_NSToolbar_SetSizeMode(void* ptr, unsigned int value);
void* C_NSToolbar_SelectedItemIdentifier(void* ptr);
void C_NSToolbar_SetSelectedItemIdentifier(void* ptr, void* value);
void* C_NSToolbar_CenteredItemIdentifier(void* ptr);
void C_NSToolbar_SetCenteredItemIdentifier(void* ptr, void* value);
bool C_NSToolbar_IsVisible(void* ptr);
void C_NSToolbar_SetVisible(void* ptr, bool value);
bool C_NSToolbar_CustomizationPaletteIsRunning(void* ptr);
bool C_NSToolbar_AutosavesConfiguration(void* ptr);
void C_NSToolbar_SetAutosavesConfiguration(void* ptr, bool value);
Dictionary C_NSToolbar_ConfigurationDictionary(void* ptr);
