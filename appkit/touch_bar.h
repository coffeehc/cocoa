#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TouchBar_Alloc();

void* C_NSTouchBar_Init(void* ptr);
void* C_NSTouchBar_InitWithCoder(void* ptr, void* coder);
void* C_NSTouchBar_AllocTouchBar();
void* C_NSTouchBar_NewTouchBar();
void* C_NSTouchBar_Autorelease(void* ptr);
void* C_NSTouchBar_Retain(void* ptr);
void* C_NSTouchBar_ItemForIdentifier(void* ptr, void* identifier);
void* C_NSTouchBar_Delegate(void* ptr);
void C_NSTouchBar_SetDelegate(void* ptr, void* value);
void* C_NSTouchBar_TemplateItems(void* ptr);
void C_NSTouchBar_SetTemplateItems(void* ptr, void* value);
Array C_NSTouchBar_DefaultItemIdentifiers(void* ptr);
void C_NSTouchBar_SetDefaultItemIdentifiers(void* ptr, Array value);
void* C_NSTouchBar_PrincipalItemIdentifier(void* ptr);
void C_NSTouchBar_SetPrincipalItemIdentifier(void* ptr, void* value);
void* C_NSTouchBar_EscapeKeyReplacementItemIdentifier(void* ptr);
void C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(void* ptr, void* value);
bool C_NSTouchBar_IsVisible(void* ptr);
Array C_NSTouchBar_ItemIdentifiers(void* ptr);
void* C_NSTouchBar_CustomizationIdentifier(void* ptr);
void C_NSTouchBar_SetCustomizationIdentifier(void* ptr, void* value);
Array C_NSTouchBar_CustomizationAllowedItemIdentifiers(void* ptr);
void C_NSTouchBar_SetCustomizationAllowedItemIdentifiers(void* ptr, Array value);
Array C_NSTouchBar_CustomizationRequiredItemIdentifiers(void* ptr);
void C_NSTouchBar_SetCustomizationRequiredItemIdentifiers(void* ptr, Array value);
bool C_NSTouchBar_AutomaticCustomizeTouchBarMenuItemEnabled();
void C_NSTouchBar_SetAutomaticCustomizeTouchBarMenuItemEnabled(bool value);
