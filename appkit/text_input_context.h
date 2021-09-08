#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextInputContext_Alloc();

void* C_NSTextInputContext_InitWithClient(void* ptr, void* client);
void* C_NSTextInputContext_AllocTextInputContext();
void* C_NSTextInputContext_Autorelease(void* ptr);
void* C_NSTextInputContext_Retain(void* ptr);
void C_NSTextInputContext_Activate(void* ptr);
void C_NSTextInputContext_Deactivate(void* ptr);
bool C_NSTextInputContext_HandleEvent(void* ptr, void* event);
void C_NSTextInputContext_DiscardMarkedText(void* ptr);
void C_NSTextInputContext_InvalidateCharacterCoordinates(void* ptr);
void* C_NSTextInputContext_TextInputContext_LocalizedNameForInputSource(void* inputSourceIdentifier);
void* C_NSTextInputContext_TextInputContext_CurrentInputContext();
void* C_NSTextInputContext_Client(void* ptr);
bool C_NSTextInputContext_AcceptsGlyphInfo(void* ptr);
void C_NSTextInputContext_SetAcceptsGlyphInfo(void* ptr, bool value);
Array C_NSTextInputContext_AllowedInputSourceLocales(void* ptr);
void C_NSTextInputContext_SetAllowedInputSourceLocales(void* ptr, Array value);
Array C_NSTextInputContext_KeyboardInputSources(void* ptr);
void* C_NSTextInputContext_SelectedKeyboardInputSource(void* ptr);
void C_NSTextInputContext_SetSelectedKeyboardInputSource(void* ptr, void* value);
