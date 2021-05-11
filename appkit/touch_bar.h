#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TouchBar_Alloc();

void* C_NSTouchBar_Init(void* ptr);
void* C_NSTouchBar_InitWithCoder(void* ptr, void* coder);
void* C_NSTouchBar_PrincipalItemIdentifier(void* ptr);
void C_NSTouchBar_SetPrincipalItemIdentifier(void* ptr, void* value);
void* C_NSTouchBar_EscapeKeyReplacementItemIdentifier(void* ptr);
void C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(void* ptr, void* value);
bool C_NSTouchBar_IsVisible(void* ptr);
void* C_NSTouchBar_CustomizationIdentifier(void* ptr);
void C_NSTouchBar_SetCustomizationIdentifier(void* ptr, void* value);
