#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_TextFieldCell_Alloc();

void* C_NSTextFieldCell_InitTextCell(void* ptr, void* _string);
void* C_NSTextFieldCell_InitWithCoder(void* ptr, void* coder);
void* C_NSTextFieldCell_Init(void* ptr);
void* C_NSTextFieldCell_AllocTextFieldCell();
void* C_NSTextFieldCell_NewTextFieldCell();
void* C_NSTextFieldCell_Autorelease(void* ptr);
void* C_NSTextFieldCell_Retain(void* ptr);
void C_NSTextFieldCell_SetWantsNotificationForMarkedText(void* ptr, bool flag);
void* C_NSTextFieldCell_TextColor(void* ptr);
void C_NSTextFieldCell_SetTextColor(void* ptr, void* value);
unsigned int C_NSTextFieldCell_BezelStyle(void* ptr);
void C_NSTextFieldCell_SetBezelStyle(void* ptr, unsigned int value);
void* C_NSTextFieldCell_BackgroundColor(void* ptr);
void C_NSTextFieldCell_SetBackgroundColor(void* ptr, void* value);
bool C_NSTextFieldCell_DrawsBackground(void* ptr);
void C_NSTextFieldCell_SetDrawsBackground(void* ptr, bool value);
void* C_NSTextFieldCell_PlaceholderString(void* ptr);
void C_NSTextFieldCell_SetPlaceholderString(void* ptr, void* value);
void* C_NSTextFieldCell_PlaceholderAttributedString(void* ptr);
void C_NSTextFieldCell_SetPlaceholderAttributedString(void* ptr, void* value);
Array C_NSTextFieldCell_AllowedInputSourceLocales(void* ptr);
void C_NSTextFieldCell_SetAllowedInputSourceLocales(void* ptr, Array value);
