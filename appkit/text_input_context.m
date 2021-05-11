#import <Appkit/Appkit.h>
#import "text_input_context.h"

void* C_TextInputContext_Alloc() {
	return [NSTextInputContext alloc];
}

void C_NSTextInputContext_Activate(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext activate];
}

void C_NSTextInputContext_Deactivate(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext deactivate];
}

bool C_NSTextInputContext_HandleEvent(void* ptr, void* event) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	bool result = [nSTextInputContext handleEvent:(NSEvent*)event];
	return result;
}

void C_NSTextInputContext_DiscardMarkedText(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext discardMarkedText];
}

void C_NSTextInputContext_InvalidateCharacterCoordinates(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext invalidateCharacterCoordinates];
}

void* C_NSTextInputContext_TextInputContextLocalizedNameForInputSource(void* inputSourceIdentifier) {
	NSString* result = [NSTextInputContext localizedNameForInputSource:(NSString*)inputSourceIdentifier];
	return result;
}

bool C_NSTextInputContext_AcceptsGlyphInfo(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	bool result = [nSTextInputContext acceptsGlyphInfo];
	return result;
}

void C_NSTextInputContext_SetAcceptsGlyphInfo(void* ptr, bool value) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext setAcceptsGlyphInfo:value];
}

void* C_NSTextInputContext_SelectedKeyboardInputSource(void* ptr) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	NSString* result = [nSTextInputContext selectedKeyboardInputSource];
	return result;
}

void C_NSTextInputContext_SetSelectedKeyboardInputSource(void* ptr, void* value) {
	NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
	[nSTextInputContext setSelectedKeyboardInputSource:(NSString*)value];
}
