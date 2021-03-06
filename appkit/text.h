#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Text_Alloc();

void* C_NSText_InitWithCoder(void* ptr, void* coder);
void* C_NSText_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSText_Init(void* ptr);
void* C_NSText_AllocText();
void* C_NSText_NewText();
void* C_NSText_Autorelease(void* ptr);
void* C_NSText_Retain(void* ptr);
void C_NSText_ToggleRuler(void* ptr, void* sender);
void C_NSText_ReplaceCharactersInRange_WithRTF(void* ptr, NSRange _range, void* rtfData);
void C_NSText_ReplaceCharactersInRange_WithRTFD(void* ptr, NSRange _range, void* rtfdData);
void C_NSText_ReplaceCharactersInRange_WithString(void* ptr, NSRange _range, void* _string);
void C_NSText_SelectAll(void* ptr, void* sender);
void C_NSText_Copy_(void* ptr, void* sender);
void C_NSText_Cut(void* ptr, void* sender);
void C_NSText_Paste(void* ptr, void* sender);
void C_NSText_CopyFont(void* ptr, void* sender);
void C_NSText_PasteFont(void* ptr, void* sender);
void C_NSText_CopyRuler(void* ptr, void* sender);
void C_NSText_PasteRuler(void* ptr, void* sender);
void C_NSText_Delete(void* ptr, void* sender);
void C_NSText_ChangeFont(void* ptr, void* sender);
void C_NSText_SetFont_Range(void* ptr, void* font, NSRange _range);
void C_NSText_AlignCenter(void* ptr, void* sender);
void C_NSText_AlignLeft(void* ptr, void* sender);
void C_NSText_AlignRight(void* ptr, void* sender);
void C_NSText_SetTextColor_Range(void* ptr, void* color, NSRange _range);
void C_NSText_Superscript(void* ptr, void* sender);
void C_NSText_Subscript(void* ptr, void* sender);
void C_NSText_Unscript(void* ptr, void* sender);
void C_NSText_Underline(void* ptr, void* sender);
bool C_NSText_ReadRTFDFromFile(void* ptr, void* path);
bool C_NSText_WriteRTFDToFile_Atomically(void* ptr, void* path, bool flag);
void* C_NSText_RTFDFromRange(void* ptr, NSRange _range);
void* C_NSText_RTFFromRange(void* ptr, NSRange _range);
void C_NSText_CheckSpelling(void* ptr, void* sender);
void C_NSText_ShowGuessPanel(void* ptr, void* sender);
void C_NSText_SizeToFit(void* ptr);
void C_NSText_ScrollRangeToVisible(void* ptr, NSRange _range);
void* C_NSText_String(void* ptr);
void C_NSText_SetString(void* ptr, void* value);
void* C_NSText_BackgroundColor(void* ptr);
void C_NSText_SetBackgroundColor(void* ptr, void* value);
bool C_NSText_DrawsBackground(void* ptr);
void C_NSText_SetDrawsBackground(void* ptr, bool value);
bool C_NSText_IsEditable(void* ptr);
void C_NSText_SetEditable(void* ptr, bool value);
bool C_NSText_IsSelectable(void* ptr);
void C_NSText_SetSelectable(void* ptr, bool value);
bool C_NSText_IsFieldEditor(void* ptr);
void C_NSText_SetFieldEditor(void* ptr, bool value);
bool C_NSText_IsRichText(void* ptr);
void C_NSText_SetRichText(void* ptr, bool value);
bool C_NSText_ImportsGraphics(void* ptr);
void C_NSText_SetImportsGraphics(void* ptr, bool value);
bool C_NSText_UsesFontPanel(void* ptr);
void C_NSText_SetUsesFontPanel(void* ptr, bool value);
bool C_NSText_IsRulerVisible(void* ptr);
NSRange C_NSText_SelectedRange(void* ptr);
void C_NSText_SetSelectedRange(void* ptr, NSRange value);
void* C_NSText_Font(void* ptr);
void C_NSText_SetFont(void* ptr, void* value);
int C_NSText_Alignment(void* ptr);
void C_NSText_SetAlignment(void* ptr, int value);
void* C_NSText_TextColor(void* ptr);
void C_NSText_SetTextColor(void* ptr, void* value);
int C_NSText_BaseWritingDirection(void* ptr);
void C_NSText_SetBaseWritingDirection(void* ptr, int value);
CGSize C_NSText_MaxSize(void* ptr);
void C_NSText_SetMaxSize(void* ptr, CGSize value);
CGSize C_NSText_MinSize(void* ptr);
void C_NSText_SetMinSize(void* ptr, CGSize value);
bool C_NSText_IsVerticallyResizable(void* ptr);
void C_NSText_SetVerticallyResizable(void* ptr, bool value);
bool C_NSText_IsHorizontallyResizable(void* ptr);
void C_NSText_SetHorizontallyResizable(void* ptr, bool value);
void* C_NSText_Delegate(void* ptr);
void C_NSText_SetDelegate(void* ptr, void* value);

void* WrapTextDelegate(uintptr_t goID);
