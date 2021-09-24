#import "touch_bar.h"
#import <AppKit/NSTouchBar.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TouchBar_Alloc() {
    return [NSTouchBar alloc];
}

void* C_NSTouchBar_Init(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBar* result_ = [nSTouchBar init];
    return result_;
}

void* C_NSTouchBar_InitWithCoder(void* ptr, void* coder) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBar* result_ = [nSTouchBar initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTouchBar_AllocTouchBar() {
    NSTouchBar* result_ = [NSTouchBar alloc];
    return result_;
}

void* C_NSTouchBar_NewTouchBar() {
    NSTouchBar* result_ = [NSTouchBar new];
    return result_;
}

void* C_NSTouchBar_Autorelease(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBar* result_ = [nSTouchBar autorelease];
    return result_;
}

void* C_NSTouchBar_Retain(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBar* result_ = [nSTouchBar retain];
    return result_;
}

void* C_NSTouchBar_ItemForIdentifier(void* ptr, void* identifier) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBarItem* result_ = [nSTouchBar itemForIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSTouchBar_Delegate(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    id result_ = [nSTouchBar delegate];
    return result_;
}

void C_NSTouchBar_SetDelegate(void* ptr, void* value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    [nSTouchBar setDelegate:(id)value];
}

void* C_NSTouchBar_TemplateItems(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSSet* result_ = [nSTouchBar templateItems];
    return result_;
}

void C_NSTouchBar_SetTemplateItems(void* ptr, void* value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    [nSTouchBar setTemplateItems:(NSSet*)value];
}

Array C_NSTouchBar_DefaultItemIdentifiers(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSArray* result_ = [nSTouchBar defaultItemIdentifiers];
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

void C_NSTouchBar_SetDefaultItemIdentifiers(void* ptr, Array value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTouchBar setDefaultItemIdentifiers:objcValue];
}

void* C_NSTouchBar_PrincipalItemIdentifier(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBarItemIdentifier result_ = [nSTouchBar principalItemIdentifier];
    return result_;
}

void C_NSTouchBar_SetPrincipalItemIdentifier(void* ptr, void* value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    [nSTouchBar setPrincipalItemIdentifier:(NSString*)value];
}

void* C_NSTouchBar_EscapeKeyReplacementItemIdentifier(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBarItemIdentifier result_ = [nSTouchBar escapeKeyReplacementItemIdentifier];
    return result_;
}

void C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(void* ptr, void* value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    [nSTouchBar setEscapeKeyReplacementItemIdentifier:(NSString*)value];
}

bool C_NSTouchBar_IsVisible(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    BOOL result_ = [nSTouchBar isVisible];
    return result_;
}

Array C_NSTouchBar_ItemIdentifiers(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSArray* result_ = [nSTouchBar itemIdentifiers];
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

void* C_NSTouchBar_CustomizationIdentifier(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSTouchBarCustomizationIdentifier result_ = [nSTouchBar customizationIdentifier];
    return result_;
}

void C_NSTouchBar_SetCustomizationIdentifier(void* ptr, void* value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    [nSTouchBar setCustomizationIdentifier:(NSString*)value];
}

Array C_NSTouchBar_CustomizationAllowedItemIdentifiers(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSArray* result_ = [nSTouchBar customizationAllowedItemIdentifiers];
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

void C_NSTouchBar_SetCustomizationAllowedItemIdentifiers(void* ptr, Array value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTouchBar setCustomizationAllowedItemIdentifiers:objcValue];
}

Array C_NSTouchBar_CustomizationRequiredItemIdentifiers(void* ptr) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSArray* result_ = [nSTouchBar customizationRequiredItemIdentifiers];
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

void C_NSTouchBar_SetCustomizationRequiredItemIdentifiers(void* ptr, Array value) {
    NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTouchBar setCustomizationRequiredItemIdentifiers:objcValue];
}

bool C_NSTouchBar_AutomaticCustomizeTouchBarMenuItemEnabled() {
    BOOL result_ = [NSTouchBar automaticCustomizeTouchBarMenuItemEnabled];
    return result_;
}

void C_NSTouchBar_SetAutomaticCustomizeTouchBarMenuItemEnabled(bool value) {
    [NSTouchBar setAutomaticCustomizeTouchBarMenuItemEnabled:value];
}
