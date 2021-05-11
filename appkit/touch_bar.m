#import <Appkit/Appkit.h>
#import "touch_bar.h"

void* C_TouchBar_Alloc() {
	return [NSTouchBar alloc];
}

void* C_NSTouchBar_Init(void* ptr) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	NSTouchBar* result = [nSTouchBar init];
	return result;
}

void* C_NSTouchBar_InitWithCoder(void* ptr, void* coder) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	NSTouchBar* result = [nSTouchBar initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSTouchBar_PrincipalItemIdentifier(void* ptr) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	NSString* result = [nSTouchBar principalItemIdentifier];
	return result;
}

void C_NSTouchBar_SetPrincipalItemIdentifier(void* ptr, void* value) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	[nSTouchBar setPrincipalItemIdentifier:(NSString*)value];
}

void* C_NSTouchBar_EscapeKeyReplacementItemIdentifier(void* ptr) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	NSString* result = [nSTouchBar escapeKeyReplacementItemIdentifier];
	return result;
}

void C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(void* ptr, void* value) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	[nSTouchBar setEscapeKeyReplacementItemIdentifier:(NSString*)value];
}

bool C_NSTouchBar_IsVisible(void* ptr) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	bool result = [nSTouchBar isVisible];
	return result;
}

void* C_NSTouchBar_CustomizationIdentifier(void* ptr) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	NSString* result = [nSTouchBar customizationIdentifier];
	return result;
}

void C_NSTouchBar_SetCustomizationIdentifier(void* ptr, void* value) {
	NSTouchBar* nSTouchBar = (NSTouchBar*)ptr;
	[nSTouchBar setCustomizationIdentifier:(NSString*)value];
}
