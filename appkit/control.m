#import <Appkit/Appkit.h>
#import "control.h"

void* C_Control_Alloc() {
    return [NSControl alloc];
}

void* C_NSControl_InitWithFrame(void* ptr, CGRect frameRect) {
    NSControl* nSControl = (NSControl*)ptr;
    NSControl* result = [nSControl initWithFrame:frameRect];
    return result;
}

void* C_NSControl_InitWithCoder(void* ptr, void* coder) {
    NSControl* nSControl = (NSControl*)ptr;
    NSControl* result = [nSControl initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSControl_Init(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSControl* result = [nSControl init];
    return result;
}

void C_NSControl_TakeDoubleValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeDoubleValueFrom:(id)sender];
}

void C_NSControl_TakeFloatValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeFloatValueFrom:(id)sender];
}

void C_NSControl_TakeIntValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeIntValueFrom:(id)sender];
}

void C_NSControl_TakeIntegerValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeIntegerValueFrom:(id)sender];
}

void C_NSControl_TakeObjectValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeObjectValueFrom:(id)sender];
}

void C_NSControl_TakeStringValueFrom(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl takeStringValueFrom:(id)sender];
}

void C_NSControl_DrawWithExpansionFrame_InView(void* ptr, CGRect contentFrame, void* view) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl drawWithExpansionFrame:contentFrame inView:(NSView*)view];
}

CGRect C_NSControl_ExpansionFrameWithFrame(void* ptr, CGRect contentFrame) {
    NSControl* nSControl = (NSControl*)ptr;
    NSRect result = [nSControl expansionFrameWithFrame:contentFrame];
    return result;
}

bool C_NSControl_AbortEditing(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl abortEditing];
    return result;
}

void* C_NSControl_CurrentEditor(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSText* result = [nSControl currentEditor];
    return result;
}

void C_NSControl_ValidateEditing(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl validateEditing];
}

void C_NSControl_EditWithFrame_Editor_Delegate_Event(void* ptr, CGRect rect, void* textObj, void* delegate, void* event) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl editWithFrame:rect editor:(NSText*)textObj delegate:(id)delegate event:(NSEvent*)event];
}

void C_NSControl_EndEditing(void* ptr, void* textObj) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl endEditing:(NSText*)textObj];
}

void C_NSControl_SelectWithFrame_Editor_Delegate_Start_Length(void* ptr, CGRect rect, void* textObj, void* delegate, int selStart, int selLength) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl selectWithFrame:rect editor:(NSText*)textObj delegate:(id)delegate start:selStart length:selLength];
}

CGSize C_NSControl_SizeThatFits(void* ptr, CGSize size) {
    NSControl* nSControl = (NSControl*)ptr;
    NSSize result = [nSControl sizeThatFits:size];
    return result;
}

void C_NSControl_SizeToFit(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl sizeToFit];
}

bool C_NSControl_SendAction_To(void* ptr, void* action, void* target) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl sendAction:(SEL)action to:(id)target];
    return result;
}

void C_NSControl_PerformClick(void* ptr, void* sender) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl performClick:(id)sender];
}

void C_NSControl_InvalidateIntrinsicContentSizeForCell(void* ptr, void* cell) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl invalidateIntrinsicContentSizeForCell:(NSCell*)cell];
}

bool C_NSControl_IsEnabled(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl isEnabled];
    return result;
}

void C_NSControl_SetEnabled(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setEnabled:value];
}

double C_NSControl_DoubleValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    double result = [nSControl doubleValue];
    return result;
}

void C_NSControl_SetDoubleValue(void* ptr, double value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setDoubleValue:value];
}

float C_NSControl_FloatValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    float result = [nSControl floatValue];
    return result;
}

void C_NSControl_SetFloatValue(void* ptr, float value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setFloatValue:value];
}

int C_NSControl_IntegerValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSInteger result = [nSControl integerValue];
    return result;
}

void C_NSControl_SetIntegerValue(void* ptr, int value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setIntegerValue:value];
}

void* C_NSControl_ObjectValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    id result = [nSControl objectValue];
    return result;
}

void C_NSControl_SetObjectValue(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setObjectValue:(id)value];
}

void* C_NSControl_StringValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSString* result = [nSControl stringValue];
    return result;
}

void C_NSControl_SetStringValue(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setStringValue:(NSString*)value];
}

void* C_NSControl_AttributedStringValue(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSAttributedString* result = [nSControl attributedStringValue];
    return result;
}

void C_NSControl_SetAttributedStringValue(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setAttributedStringValue:(NSAttributedString*)value];
}

int C_NSControl_Alignment(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSTextAlignment result = [nSControl alignment];
    return result;
}

void C_NSControl_SetAlignment(void* ptr, int value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setAlignment:value];
}

void* C_NSControl_Font(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSFont* result = [nSControl font];
    return result;
}

void C_NSControl_SetFont(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setFont:(NSFont*)value];
}

int C_NSControl_LineBreakMode(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSLineBreakMode result = [nSControl lineBreakMode];
    return result;
}

void C_NSControl_SetLineBreakMode(void* ptr, int value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setLineBreakMode:value];
}

bool C_NSControl_UsesSingleLineMode(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl usesSingleLineMode];
    return result;
}

void C_NSControl_SetUsesSingleLineMode(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setUsesSingleLineMode:value];
}

void* C_NSControl_Formatter(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSFormatter* result = [nSControl formatter];
    return result;
}

void C_NSControl_SetFormatter(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setFormatter:(NSFormatter*)value];
}

int C_NSControl_BaseWritingDirection(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSWritingDirection result = [nSControl baseWritingDirection];
    return result;
}

void C_NSControl_SetBaseWritingDirection(void* ptr, int value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setBaseWritingDirection:value];
}

bool C_NSControl_AllowsExpansionToolTips(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl allowsExpansionToolTips];
    return result;
}

void C_NSControl_SetAllowsExpansionToolTips(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setAllowsExpansionToolTips:value];
}

unsigned int C_NSControl_ControlSize(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSControlSize result = [nSControl controlSize];
    return result;
}

void C_NSControl_SetControlSize(void* ptr, unsigned int value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setControlSize:value];
}

bool C_NSControl_IsHighlighted(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl isHighlighted];
    return result;
}

void C_NSControl_SetHighlighted(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setHighlighted:value];
}

void* C_NSControl_Action(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    SEL result = [nSControl action];
    return result;
}

void C_NSControl_SetAction(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setAction:(SEL)value];
}

void* C_NSControl_Target(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    id result = [nSControl target];
    return result;
}

void C_NSControl_SetTarget(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setTarget:(id)value];
}

bool C_NSControl_IsContinuous(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl isContinuous];
    return result;
}

void C_NSControl_SetContinuous(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setContinuous:value];
}

bool C_NSControl_RefusesFirstResponder(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl refusesFirstResponder];
    return result;
}

void C_NSControl_SetRefusesFirstResponder(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setRefusesFirstResponder:value];
}

bool C_NSControl_IgnoresMultiClick(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    BOOL result = [nSControl ignoresMultiClick];
    return result;
}

void C_NSControl_SetIgnoresMultiClick(void* ptr, bool value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setIgnoresMultiClick:value];
}

void* C_NSControl_Cell(void* ptr) {
    NSControl* nSControl = (NSControl*)ptr;
    NSCell* result = [nSControl cell];
    return result;
}

void C_NSControl_SetCell(void* ptr, void* value) {
    NSControl* nSControl = (NSControl*)ptr;
    [nSControl setCell:(NSCell*)value];
}
