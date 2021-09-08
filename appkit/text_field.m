#import <Appkit/Appkit.h>
#import "text_field.h"

void* C_TextField_Alloc() {
    return [NSTextField alloc];
}

void* C_NSTextField_TextField_LabelWithAttributedString(void* attributedStringValue) {
    NSTextField* result_ = [NSTextField labelWithAttributedString:(NSAttributedString*)attributedStringValue];
    return result_;
}

void* C_NSTextField_TextField_LabelWithString(void* stringValue) {
    NSTextField* result_ = [NSTextField labelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSTextField_TextFieldWithString(void* stringValue) {
    NSTextField* result_ = [NSTextField textFieldWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSTextField_TextField_WrappingLabelWithString(void* stringValue) {
    NSTextField* result_ = [NSTextField wrappingLabelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSTextField_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextField* result_ = [nSTextField initWithFrame:frameRect];
    return result_;
}

void* C_NSTextField_InitWithCoder(void* ptr, void* coder) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextField* result_ = [nSTextField initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTextField_Init(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextField* result_ = [nSTextField init];
    return result_;
}

void* C_NSTextField_AllocTextField() {
    NSTextField* result_ = [NSTextField alloc];
    return result_;
}

void* C_NSTextField_NewTextField() {
    NSTextField* result_ = [NSTextField new];
    return result_;
}

void* C_NSTextField_Autorelease(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextField* result_ = [nSTextField autorelease];
    return result_;
}

void* C_NSTextField_Retain(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextField* result_ = [nSTextField retain];
    return result_;
}

void C_NSTextField_SelectText(void* ptr, void* sender) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField selectText:(id)sender];
}

bool C_NSTextField_TextShouldBeginEditing(void* ptr, void* textObject) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField textShouldBeginEditing:(NSText*)textObject];
    return result_;
}

void C_NSTextField_TextDidBeginEditing(void* ptr, void* notification) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField textDidBeginEditing:(NSNotification*)notification];
}

void C_NSTextField_TextDidChange(void* ptr, void* notification) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField textDidChange:(NSNotification*)notification];
}

bool C_NSTextField_TextShouldEndEditing(void* ptr, void* textObject) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField textShouldEndEditing:(NSText*)textObject];
    return result_;
}

void C_NSTextField_TextDidEndEditing(void* ptr, void* notification) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField textDidEndEditing:(NSNotification*)notification];
}

bool C_NSTextField_IsSelectable(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField isSelectable];
    return result_;
}

void C_NSTextField_SetSelectable(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setSelectable:value];
}

bool C_NSTextField_IsEditable(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField isEditable];
    return result_;
}

void C_NSTextField_SetEditable(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setEditable:value];
}

bool C_NSTextField_AllowsEditingTextAttributes(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField allowsEditingTextAttributes];
    return result_;
}

void C_NSTextField_SetAllowsEditingTextAttributes(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setAllowsEditingTextAttributes:value];
}

bool C_NSTextField_ImportsGraphics(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField importsGraphics];
    return result_;
}

void C_NSTextField_SetImportsGraphics(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setImportsGraphics:value];
}

void* C_NSTextField_PlaceholderString(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSString* result_ = [nSTextField placeholderString];
    return result_;
}

void C_NSTextField_SetPlaceholderString(void* ptr, void* value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setPlaceholderString:(NSString*)value];
}

void* C_NSTextField_PlaceholderAttributedString(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSAttributedString* result_ = [nSTextField placeholderAttributedString];
    return result_;
}

void C_NSTextField_SetPlaceholderAttributedString(void* ptr, void* value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setPlaceholderAttributedString:(NSAttributedString*)value];
}

bool C_NSTextField_AllowsDefaultTighteningForTruncation(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField allowsDefaultTighteningForTruncation];
    return result_;
}

void C_NSTextField_SetAllowsDefaultTighteningForTruncation(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setAllowsDefaultTighteningForTruncation:value];
}

int C_NSTextField_MaximumNumberOfLines(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSInteger result_ = [nSTextField maximumNumberOfLines];
    return result_;
}

void C_NSTextField_SetMaximumNumberOfLines(void* ptr, int value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setMaximumNumberOfLines:value];
}

double C_NSTextField_PreferredMaxLayoutWidth(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    CGFloat result_ = [nSTextField preferredMaxLayoutWidth];
    return result_;
}

void C_NSTextField_SetPreferredMaxLayoutWidth(void* ptr, double value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setPreferredMaxLayoutWidth:value];
}

void* C_NSTextField_TextColor(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSColor* result_ = [nSTextField textColor];
    return result_;
}

void C_NSTextField_SetTextColor(void* ptr, void* value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setTextColor:(NSColor*)value];
}

void* C_NSTextField_BackgroundColor(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSColor* result_ = [nSTextField backgroundColor];
    return result_;
}

void C_NSTextField_SetBackgroundColor(void* ptr, void* value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setBackgroundColor:(NSColor*)value];
}

bool C_NSTextField_DrawsBackground(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField drawsBackground];
    return result_;
}

void C_NSTextField_SetDrawsBackground(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setDrawsBackground:value];
}

bool C_NSTextField_IsBezeled(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField isBezeled];
    return result_;
}

void C_NSTextField_SetBezeled(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setBezeled:value];
}

unsigned int C_NSTextField_BezelStyle(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSTextFieldBezelStyle result_ = [nSTextField bezelStyle];
    return result_;
}

void C_NSTextField_SetBezelStyle(void* ptr, unsigned int value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setBezelStyle:value];
}

bool C_NSTextField_IsBordered(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField isBordered];
    return result_;
}

void C_NSTextField_SetBordered(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setBordered:value];
}

bool C_NSTextField_AllowsCharacterPickerTouchBarItem(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField allowsCharacterPickerTouchBarItem];
    return result_;
}

void C_NSTextField_SetAllowsCharacterPickerTouchBarItem(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setAllowsCharacterPickerTouchBarItem:value];
}

bool C_NSTextField_IsAutomaticTextCompletionEnabled(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    BOOL result_ = [nSTextField isAutomaticTextCompletionEnabled];
    return result_;
}

void C_NSTextField_SetAutomaticTextCompletionEnabled(void* ptr, bool value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setAutomaticTextCompletionEnabled:value];
}

void* C_NSTextField_Delegate(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    id result_ = [nSTextField delegate];
    return result_;
}

void C_NSTextField_SetDelegate(void* ptr, void* value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setDelegate:(id)value];
}

unsigned int C_NSTextField_LineBreakStrategy(void* ptr) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    NSLineBreakStrategy result_ = [nSTextField lineBreakStrategy];
    return result_;
}

void C_NSTextField_SetLineBreakStrategy(void* ptr, unsigned int value) {
    NSTextField* nSTextField = (NSTextField*)ptr;
    [nSTextField setLineBreakStrategy:value];
}
