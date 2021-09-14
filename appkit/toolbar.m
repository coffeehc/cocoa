#import <Appkit/Appkit.h>
#import "toolbar.h"

void* C_Toolbar_Alloc() {
    return [NSToolbar alloc];
}

void* C_NSToolbar_InitWithIdentifier(void* ptr, void* identifier) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbar* result_ = [nSToolbar initWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSToolbar_Init(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbar* result_ = [nSToolbar init];
    return result_;
}

void* C_NSToolbar_AllocToolbar() {
    NSToolbar* result_ = [NSToolbar alloc];
    return result_;
}

void* C_NSToolbar_NewToolbar() {
    NSToolbar* result_ = [NSToolbar new];
    return result_;
}

void* C_NSToolbar_Autorelease(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbar* result_ = [nSToolbar autorelease];
    return result_;
}

void* C_NSToolbar_Retain(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbar* result_ = [nSToolbar retain];
    return result_;
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

void C_NSToolbar_SetConfigurationFromDictionary(void* ptr, Dictionary configDict) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSMutableDictionary* objcConfigDict = [NSMutableDictionary dictionaryWithCapacity:configDict.len];
    if (configDict.len > 0) {
    	void** configDictKeyData = (void**)configDict.key_data;
    	void** configDictValueData = (void**)configDict.value_data;
    	for (int i = 0; i < configDict.len; i++) {
    		void* kp = configDictKeyData[i];
    		void* vp = configDictValueData[i];
    		[objcConfigDict setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSToolbar setConfigurationFromDictionary:objcConfigDict];
}

void C_NSToolbar_ValidateVisibleItems(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar validateVisibleItems];
}

void* C_NSToolbar_Delegate(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    id result_ = [nSToolbar delegate];
    return result_;
}

void C_NSToolbar_SetDelegate(void* ptr, void* value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setDelegate:(id)value];
}

void* C_NSToolbar_Identifier(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbarIdentifier result_ = [nSToolbar identifier];
    return result_;
}

unsigned int C_NSToolbar_DisplayMode(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbarDisplayMode result_ = [nSToolbar displayMode];
    return result_;
}

void C_NSToolbar_SetDisplayMode(void* ptr, unsigned int value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setDisplayMode:value];
}

bool C_NSToolbar_ShowsBaselineSeparator(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar showsBaselineSeparator];
    return result_;
}

void C_NSToolbar_SetShowsBaselineSeparator(void* ptr, bool value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setShowsBaselineSeparator:value];
}

bool C_NSToolbar_AllowsUserCustomization(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar allowsUserCustomization];
    return result_;
}

void C_NSToolbar_SetAllowsUserCustomization(void* ptr, bool value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setAllowsUserCustomization:value];
}

bool C_NSToolbar_AllowsExtensionItems(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar allowsExtensionItems];
    return result_;
}

void C_NSToolbar_SetAllowsExtensionItems(void* ptr, bool value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setAllowsExtensionItems:value];
}

Array C_NSToolbar_Items(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSArray* result_ = [nSToolbar items];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSToolbar_VisibleItems(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSArray* result_ = [nSToolbar visibleItems];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

unsigned int C_NSToolbar_SizeMode(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbarSizeMode result_ = [nSToolbar sizeMode];
    return result_;
}

void C_NSToolbar_SetSizeMode(void* ptr, unsigned int value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setSizeMode:value];
}

void* C_NSToolbar_SelectedItemIdentifier(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbarItemIdentifier result_ = [nSToolbar selectedItemIdentifier];
    return result_;
}

void C_NSToolbar_SetSelectedItemIdentifier(void* ptr, void* value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setSelectedItemIdentifier:(NSString*)value];
}

void* C_NSToolbar_CenteredItemIdentifier(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSToolbarItemIdentifier result_ = [nSToolbar centeredItemIdentifier];
    return result_;
}

void C_NSToolbar_SetCenteredItemIdentifier(void* ptr, void* value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setCenteredItemIdentifier:(NSString*)value];
}

bool C_NSToolbar_IsVisible(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar isVisible];
    return result_;
}

void C_NSToolbar_SetVisible(void* ptr, bool value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setVisible:value];
}

bool C_NSToolbar_CustomizationPaletteIsRunning(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar customizationPaletteIsRunning];
    return result_;
}

bool C_NSToolbar_AutosavesConfiguration(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    BOOL result_ = [nSToolbar autosavesConfiguration];
    return result_;
}

void C_NSToolbar_SetAutosavesConfiguration(void* ptr, bool value) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    [nSToolbar setAutosavesConfiguration:value];
}

Dictionary C_NSToolbar_ConfigurationDictionary(void* ptr) {
    NSToolbar* nSToolbar = (NSToolbar*)ptr;
    NSDictionary* result_ = [nSToolbar configurationDictionary];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}
