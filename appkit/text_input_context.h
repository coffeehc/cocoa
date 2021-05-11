#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextInputContext_Alloc();

void C_NSTextInputContext_Activate(void* ptr);
void C_NSTextInputContext_Deactivate(void* ptr);
bool C_NSTextInputContext_HandleEvent(void* ptr, void* event);
void C_NSTextInputContext_DiscardMarkedText(void* ptr);
void C_NSTextInputContext_InvalidateCharacterCoordinates(void* ptr);
void* C_NSTextInputContext_TextInputContextLocalizedNameForInputSource(void* inputSourceIdentifier);
bool C_NSTextInputContext_AcceptsGlyphInfo(void* ptr);
void C_NSTextInputContext_SetAcceptsGlyphInfo(void* ptr, bool value);
void* C_NSTextInputContext_SelectedKeyboardInputSource(void* ptr);
void C_NSTextInputContext_SetSelectedKeyboardInputSource(void* ptr, void* value);
