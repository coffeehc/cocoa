#import "text_field.h"
#import <AppKit/NSTextField.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

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

@interface NSTextFieldDelegateAdaptor : NSObject <NSTextFieldDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSTextFieldDelegateAdaptor


- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidates:(NSArray*)candidates forSelectedRange:(NSRange)selectedRange {
    Array candidatesArray;
    int candidatescount = [candidates count];
    if (candidatescount > 0) {
    	void** candidatesData = malloc(candidatescount * sizeof(void*));
    	for (int i = 0; i < candidatescount; i++) {
    		 void* p = [candidates objectAtIndex:i];
    		 candidatesData[i] = p;
    	}
    	candidatesArray.data = candidatesData;
    	candidatesArray.len = candidatescount;
    }
    Array result_ = textFieldDelegate_TextField_TextView_Candidates_ForSelectedRange([self goID], textField, textView, candidatesArray, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSTextCheckingResult*)p];
    	}
    }
    return objcResult_;
}

- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidatesForSelectedRange:(NSRange)selectedRange {
    Array result_ = textFieldDelegate_TextField_TextView_CandidatesForSelectedRange([self goID], textField, textView, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSObject*)p];
    	}
    }
    return objcResult_;
}

- (BOOL)textField:(NSTextField*)textField textView:(NSTextView*)textView shouldSelectCandidateAtIndex:(NSUInteger)index {
    bool result_ = textFieldDelegate_TextField_TextView_ShouldSelectCandidateAtIndex([self goID], textField, textView, index);
    return result_;
}

- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = textFieldDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    textFieldDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = textFieldDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = textFieldDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = textFieldDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = textFieldDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    textFieldDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    textFieldDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    textFieldDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TextFieldDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapTextFieldDelegate(uintptr_t goID) {
    NSTextFieldDelegateAdaptor* adaptor = [[NSTextFieldDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
