#import <Appkit/Appkit.h>
#import "toolbar.h"

void* C_Toolbar_Alloc() {
	return [NSToolbar alloc];
}

void* C_NSToolbar_InitWithIdentifier(void* ptr, void* identifier) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	NSToolbar* result = [nSToolbar initWithIdentifier:(NSString*)identifier];
	return result;
}

void* C_NSToolbar_Init(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	NSToolbar* result = [nSToolbar init];
	return result;
}

void C_NSToolbar_InsertItemWithItemIdentifier_AtIndex(void* ptr, void* itemIdentifier, int index) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar insertItemWithItemIdentifier:(NSString*)itemIdentifier atIndex:index];
}

void C_NSToolbar_RemoveItemAtIndex(void* ptr, int index) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar removeItemAtIndex:index];
}

void C_NSToolbar_RunCustomizationPalette(void* ptr, void* sender) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar runCustomizationPalette:(id)sender];
}

void C_NSToolbar_ValidateVisibleItems(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar validateVisibleItems];
}

void* C_NSToolbar_Identifier(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	NSString* result = [nSToolbar identifier];
	return result;
}

bool C_NSToolbar_ShowsBaselineSeparator(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar showsBaselineSeparator];
	return result;
}

void C_NSToolbar_SetShowsBaselineSeparator(void* ptr, bool value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setShowsBaselineSeparator:value];
}

bool C_NSToolbar_AllowsUserCustomization(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar allowsUserCustomization];
	return result;
}

void C_NSToolbar_SetAllowsUserCustomization(void* ptr, bool value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setAllowsUserCustomization:value];
}

bool C_NSToolbar_AllowsExtensionItems(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar allowsExtensionItems];
	return result;
}

void C_NSToolbar_SetAllowsExtensionItems(void* ptr, bool value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setAllowsExtensionItems:value];
}

void* C_NSToolbar_SelectedItemIdentifier(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	NSString* result = [nSToolbar selectedItemIdentifier];
	return result;
}

void C_NSToolbar_SetSelectedItemIdentifier(void* ptr, void* value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setSelectedItemIdentifier:(NSString*)value];
}

void* C_NSToolbar_CenteredItemIdentifier(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	NSString* result = [nSToolbar centeredItemIdentifier];
	return result;
}

void C_NSToolbar_SetCenteredItemIdentifier(void* ptr, void* value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setCenteredItemIdentifier:(NSString*)value];
}

bool C_NSToolbar_IsVisible(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar isVisible];
	return result;
}

void C_NSToolbar_SetVisible(void* ptr, bool value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setVisible:value];
}

bool C_NSToolbar_CustomizationPaletteIsRunning(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar customizationPaletteIsRunning];
	return result;
}

bool C_NSToolbar_AutosavesConfiguration(void* ptr) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	bool result = [nSToolbar autosavesConfiguration];
	return result;
}

void C_NSToolbar_SetAutosavesConfiguration(void* ptr, bool value) {
	NSToolbar* nSToolbar = (NSToolbar*)ptr;
	[nSToolbar setAutosavesConfiguration:value];
}
