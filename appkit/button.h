#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Button_Alloc();

void* C_NSButton_Button_CheckboxWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSButton_ButtonWithImage_Target_Action(void* image, void* target, void* action);
void* C_NSButton_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action);
void* C_NSButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSButton_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSButton_InitWithCoder(void* ptr, void* coder);
void* C_NSButton_Init(void* ptr);
void* C_NSButton_AllocButton();
void* C_NSButton_NewButton();
void* C_NSButton_Autorelease(void* ptr);
void* C_NSButton_Retain(void* ptr);
void C_NSButton_SetButtonType(void* ptr, unsigned int _type);
void C_NSButton_SetPeriodicDelay_Interval(void* ptr, float delay, float interval);
void C_NSButton_CompressWithPrioritizedCompressionOptions(void* ptr, Array prioritizedOptions);
CGSize C_NSButton_MinimumSizeWithPrioritizedCompressionOptions(void* ptr, Array prioritizedOptions);
void C_NSButton_SetNextState(void* ptr);
void C_NSButton_Highlight(void* ptr, bool flag);
void* C_NSButton_ContentTintColor(void* ptr);
void C_NSButton_SetContentTintColor(void* ptr, void* value);
bool C_NSButton_HasDestructiveAction(void* ptr);
void C_NSButton_SetHasDestructiveAction(void* ptr, bool value);
void* C_NSButton_AlternateTitle(void* ptr);
void C_NSButton_SetAlternateTitle(void* ptr, void* value);
void* C_NSButton_AttributedTitle(void* ptr);
void C_NSButton_SetAttributedTitle(void* ptr, void* value);
void* C_NSButton_AttributedAlternateTitle(void* ptr);
void C_NSButton_SetAttributedAlternateTitle(void* ptr, void* value);
void* C_NSButton_Title(void* ptr);
void C_NSButton_SetTitle(void* ptr, void* value);
void* C_NSButton_Sound(void* ptr);
void C_NSButton_SetSound(void* ptr, void* value);
bool C_NSButton_IsSpringLoaded(void* ptr);
void C_NSButton_SetSpringLoaded(void* ptr, bool value);
int C_NSButton_MaxAcceleratorLevel(void* ptr);
void C_NSButton_SetMaxAcceleratorLevel(void* ptr, int value);
void* C_NSButton_Image(void* ptr);
void C_NSButton_SetImage(void* ptr, void* value);
void* C_NSButton_AlternateImage(void* ptr);
void C_NSButton_SetAlternateImage(void* ptr, void* value);
unsigned int C_NSButton_ImagePosition(void* ptr);
void C_NSButton_SetImagePosition(void* ptr, unsigned int value);
bool C_NSButton_IsBordered(void* ptr);
void C_NSButton_SetBordered(void* ptr, bool value);
bool C_NSButton_IsTransparent(void* ptr);
void C_NSButton_SetTransparent(void* ptr, bool value);
unsigned int C_NSButton_BezelStyle(void* ptr);
void C_NSButton_SetBezelStyle(void* ptr, unsigned int value);
void* C_NSButton_BezelColor(void* ptr);
void C_NSButton_SetBezelColor(void* ptr, void* value);
bool C_NSButton_ShowsBorderOnlyWhileMouseInside(void* ptr);
void C_NSButton_SetShowsBorderOnlyWhileMouseInside(void* ptr, bool value);
bool C_NSButton_ImageHugsTitle(void* ptr);
void C_NSButton_SetImageHugsTitle(void* ptr, bool value);
unsigned int C_NSButton_ImageScaling(void* ptr);
void C_NSButton_SetImageScaling(void* ptr, unsigned int value);
void* C_NSButton_ActiveCompressionOptions(void* ptr);
bool C_NSButton_AllowsMixedState(void* ptr);
void C_NSButton_SetAllowsMixedState(void* ptr, bool value);
int C_NSButton_State(void* ptr);
void C_NSButton_SetState(void* ptr, int value);
void* C_NSButton_KeyEquivalent(void* ptr);
void C_NSButton_SetKeyEquivalent(void* ptr, void* value);
unsigned int C_NSButton_KeyEquivalentModifierMask(void* ptr);
void C_NSButton_SetKeyEquivalentModifierMask(void* ptr, unsigned int value);
void* C_NSButton_SymbolConfiguration(void* ptr);
void C_NSButton_SetSymbolConfiguration(void* ptr, void* value);
