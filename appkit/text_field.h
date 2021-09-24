#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TextField_Alloc();

void* C_NSTextField_TextField_LabelWithAttributedString(void* attributedStringValue);
void* C_NSTextField_TextField_LabelWithString(void* stringValue);
void* C_NSTextField_TextFieldWithString(void* stringValue);
void* C_NSTextField_TextField_WrappingLabelWithString(void* stringValue);
void* C_NSTextField_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTextField_InitWithCoder(void* ptr, void* coder);
void* C_NSTextField_Init(void* ptr);
void* C_NSTextField_AllocTextField();
void* C_NSTextField_NewTextField();
void* C_NSTextField_Autorelease(void* ptr);
void* C_NSTextField_Retain(void* ptr);
void C_NSTextField_SelectText(void* ptr, void* sender);
bool C_NSTextField_TextShouldBeginEditing(void* ptr, void* textObject);
void C_NSTextField_TextDidBeginEditing(void* ptr, void* notification);
void C_NSTextField_TextDidChange(void* ptr, void* notification);
bool C_NSTextField_TextShouldEndEditing(void* ptr, void* textObject);
void C_NSTextField_TextDidEndEditing(void* ptr, void* notification);
bool C_NSTextField_IsSelectable(void* ptr);
void C_NSTextField_SetSelectable(void* ptr, bool value);
bool C_NSTextField_IsEditable(void* ptr);
void C_NSTextField_SetEditable(void* ptr, bool value);
bool C_NSTextField_AllowsEditingTextAttributes(void* ptr);
void C_NSTextField_SetAllowsEditingTextAttributes(void* ptr, bool value);
bool C_NSTextField_ImportsGraphics(void* ptr);
void C_NSTextField_SetImportsGraphics(void* ptr, bool value);
void* C_NSTextField_PlaceholderString(void* ptr);
void C_NSTextField_SetPlaceholderString(void* ptr, void* value);
void* C_NSTextField_PlaceholderAttributedString(void* ptr);
void C_NSTextField_SetPlaceholderAttributedString(void* ptr, void* value);
bool C_NSTextField_AllowsDefaultTighteningForTruncation(void* ptr);
void C_NSTextField_SetAllowsDefaultTighteningForTruncation(void* ptr, bool value);
int C_NSTextField_MaximumNumberOfLines(void* ptr);
void C_NSTextField_SetMaximumNumberOfLines(void* ptr, int value);
double C_NSTextField_PreferredMaxLayoutWidth(void* ptr);
void C_NSTextField_SetPreferredMaxLayoutWidth(void* ptr, double value);
void* C_NSTextField_TextColor(void* ptr);
void C_NSTextField_SetTextColor(void* ptr, void* value);
void* C_NSTextField_BackgroundColor(void* ptr);
void C_NSTextField_SetBackgroundColor(void* ptr, void* value);
bool C_NSTextField_DrawsBackground(void* ptr);
void C_NSTextField_SetDrawsBackground(void* ptr, bool value);
bool C_NSTextField_IsBezeled(void* ptr);
void C_NSTextField_SetBezeled(void* ptr, bool value);
unsigned int C_NSTextField_BezelStyle(void* ptr);
void C_NSTextField_SetBezelStyle(void* ptr, unsigned int value);
bool C_NSTextField_IsBordered(void* ptr);
void C_NSTextField_SetBordered(void* ptr, bool value);
bool C_NSTextField_AllowsCharacterPickerTouchBarItem(void* ptr);
void C_NSTextField_SetAllowsCharacterPickerTouchBarItem(void* ptr, bool value);
bool C_NSTextField_IsAutomaticTextCompletionEnabled(void* ptr);
void C_NSTextField_SetAutomaticTextCompletionEnabled(void* ptr, bool value);
void* C_NSTextField_Delegate(void* ptr);
void C_NSTextField_SetDelegate(void* ptr, void* value);
unsigned int C_NSTextField_LineBreakStrategy(void* ptr);
void C_NSTextField_SetLineBreakStrategy(void* ptr, unsigned int value);

void* WrapTextFieldDelegate(uintptr_t goID);
