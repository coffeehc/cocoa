#import <Appkit/Appkit.h>
#import "text.h"

void* C_Text_Alloc() {
    return [NSText alloc];
}

void* C_NSText_InitWithCoder(void* ptr, void* coder) {
    NSText* nSText = (NSText*)ptr;
    NSText* result_ = [nSText initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSText_InitWithFrame(void* ptr, CGRect frameRect) {
    NSText* nSText = (NSText*)ptr;
    NSText* result_ = [nSText initWithFrame:frameRect];
    return result_;
}

void* C_NSText_Init(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSText* result_ = [nSText init];
    return result_;
}

void C_NSText_ToggleRuler(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText toggleRuler:(id)sender];
}

void C_NSText_ReplaceCharactersInRange_WithRTF(void* ptr, NSRange _range, Array rtfData) {
    NSText* nSText = (NSText*)ptr;
    [nSText replaceCharactersInRange:_range withRTF:[[NSData alloc] initWithBytes:(Byte *)rtfData.data length:rtfData.len]];
}

void C_NSText_ReplaceCharactersInRange_WithRTFD(void* ptr, NSRange _range, Array rtfdData) {
    NSText* nSText = (NSText*)ptr;
    [nSText replaceCharactersInRange:_range withRTFD:[[NSData alloc] initWithBytes:(Byte *)rtfdData.data length:rtfdData.len]];
}

void C_NSText_ReplaceCharactersInRange_WithString(void* ptr, NSRange _range, void* _string) {
    NSText* nSText = (NSText*)ptr;
    [nSText replaceCharactersInRange:_range withString:(NSString*)_string];
}

void C_NSText_SelectAll(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText selectAll:(id)sender];
}

void C_NSText_Copy(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText copy:(id)sender];
}

void C_NSText_Cut(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText cut:(id)sender];
}

void C_NSText_Paste(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText paste:(id)sender];
}

void C_NSText_CopyFont(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText copyFont:(id)sender];
}

void C_NSText_PasteFont(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText pasteFont:(id)sender];
}

void C_NSText_CopyRuler(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText copyRuler:(id)sender];
}

void C_NSText_PasteRuler(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText pasteRuler:(id)sender];
}

void C_NSText_Delete(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText delete:(id)sender];
}

void C_NSText_ChangeFont(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText changeFont:(id)sender];
}

void C_NSText_SetFont_Range(void* ptr, void* font, NSRange _range) {
    NSText* nSText = (NSText*)ptr;
    [nSText setFont:(NSFont*)font range:_range];
}

void C_NSText_AlignCenter(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText alignCenter:(id)sender];
}

void C_NSText_AlignLeft(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText alignLeft:(id)sender];
}

void C_NSText_AlignRight(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText alignRight:(id)sender];
}

void C_NSText_SetTextColor_Range(void* ptr, void* color, NSRange _range) {
    NSText* nSText = (NSText*)ptr;
    [nSText setTextColor:(NSColor*)color range:_range];
}

void C_NSText_Superscript(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText superscript:(id)sender];
}

void C_NSText_Subscript(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText subscript:(id)sender];
}

void C_NSText_Unscript(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText unscript:(id)sender];
}

void C_NSText_Underline(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText underline:(id)sender];
}

bool C_NSText_ReadRTFDFromFile(void* ptr, void* path) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText readRTFDFromFile:(NSString*)path];
    return result_;
}

bool C_NSText_WriteRTFDToFile_Atomically(void* ptr, void* path, bool flag) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText writeRTFDToFile:(NSString*)path atomically:flag];
    return result_;
}

Array C_NSText_RTFDFromRange(void* ptr, NSRange _range) {
    NSText* nSText = (NSText*)ptr;
    NSData* result_ = [nSText RTFDFromRange:_range];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSText_RTFFromRange(void* ptr, NSRange _range) {
    NSText* nSText = (NSText*)ptr;
    NSData* result_ = [nSText RTFFromRange:_range];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void C_NSText_CheckSpelling(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText checkSpelling:(id)sender];
}

void C_NSText_ShowGuessPanel(void* ptr, void* sender) {
    NSText* nSText = (NSText*)ptr;
    [nSText showGuessPanel:(id)sender];
}

void C_NSText_SizeToFit(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    [nSText sizeToFit];
}

void C_NSText_ScrollRangeToVisible(void* ptr, NSRange _range) {
    NSText* nSText = (NSText*)ptr;
    [nSText scrollRangeToVisible:_range];
}

void* C_NSText_String(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSString* result_ = [nSText string];
    return result_;
}

void C_NSText_SetString(void* ptr, void* value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setString:(NSString*)value];
}

void* C_NSText_BackgroundColor(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSColor* result_ = [nSText backgroundColor];
    return result_;
}

void C_NSText_SetBackgroundColor(void* ptr, void* value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setBackgroundColor:(NSColor*)value];
}

bool C_NSText_DrawsBackground(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText drawsBackground];
    return result_;
}

void C_NSText_SetDrawsBackground(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setDrawsBackground:value];
}

bool C_NSText_IsEditable(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isEditable];
    return result_;
}

void C_NSText_SetEditable(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setEditable:value];
}

bool C_NSText_IsSelectable(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isSelectable];
    return result_;
}

void C_NSText_SetSelectable(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setSelectable:value];
}

bool C_NSText_IsFieldEditor(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isFieldEditor];
    return result_;
}

void C_NSText_SetFieldEditor(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setFieldEditor:value];
}

bool C_NSText_IsRichText(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isRichText];
    return result_;
}

void C_NSText_SetRichText(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setRichText:value];
}

bool C_NSText_ImportsGraphics(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText importsGraphics];
    return result_;
}

void C_NSText_SetImportsGraphics(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setImportsGraphics:value];
}

bool C_NSText_UsesFontPanel(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText usesFontPanel];
    return result_;
}

void C_NSText_SetUsesFontPanel(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setUsesFontPanel:value];
}

bool C_NSText_IsRulerVisible(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isRulerVisible];
    return result_;
}

NSRange C_NSText_SelectedRange(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSRange result_ = [nSText selectedRange];
    return result_;
}

void C_NSText_SetSelectedRange(void* ptr, NSRange value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setSelectedRange:value];
}

void* C_NSText_Font(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSFont* result_ = [nSText font];
    return result_;
}

void C_NSText_SetFont(void* ptr, void* value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setFont:(NSFont*)value];
}

int C_NSText_Alignment(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSTextAlignment result_ = [nSText alignment];
    return result_;
}

void C_NSText_SetAlignment(void* ptr, int value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setAlignment:value];
}

void* C_NSText_TextColor(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSColor* result_ = [nSText textColor];
    return result_;
}

void C_NSText_SetTextColor(void* ptr, void* value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setTextColor:(NSColor*)value];
}

int C_NSText_BaseWritingDirection(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSWritingDirection result_ = [nSText baseWritingDirection];
    return result_;
}

void C_NSText_SetBaseWritingDirection(void* ptr, int value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setBaseWritingDirection:value];
}

CGSize C_NSText_MaxSize(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSSize result_ = [nSText maxSize];
    return result_;
}

void C_NSText_SetMaxSize(void* ptr, CGSize value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setMaxSize:value];
}

CGSize C_NSText_MinSize(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    NSSize result_ = [nSText minSize];
    return result_;
}

void C_NSText_SetMinSize(void* ptr, CGSize value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setMinSize:value];
}

bool C_NSText_IsVerticallyResizable(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isVerticallyResizable];
    return result_;
}

void C_NSText_SetVerticallyResizable(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setVerticallyResizable:value];
}

bool C_NSText_IsHorizontallyResizable(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    BOOL result_ = [nSText isHorizontallyResizable];
    return result_;
}

void C_NSText_SetHorizontallyResizable(void* ptr, bool value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setHorizontallyResizable:value];
}

void* C_NSText_Delegate(void* ptr) {
    NSText* nSText = (NSText*)ptr;
    id result_ = [nSText delegate];
    return result_;
}

void C_NSText_SetDelegate(void* ptr, void* value) {
    NSText* nSText = (NSText*)ptr;
    [nSText setDelegate:(id)value];
}
