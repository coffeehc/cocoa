#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Control_Alloc();

void* C_NSControl_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSControl_InitWithCoder(void* ptr, void* coder);
void* C_NSControl_Init(void* ptr);
void* C_NSControl_AllocControl();
void* C_NSControl_NewControl();
void* C_NSControl_Autorelease(void* ptr);
void* C_NSControl_Retain(void* ptr);
void C_NSControl_TakeDoubleValueFrom(void* ptr, void* sender);
void C_NSControl_TakeFloatValueFrom(void* ptr, void* sender);
void C_NSControl_TakeIntValueFrom(void* ptr, void* sender);
void C_NSControl_TakeIntegerValueFrom(void* ptr, void* sender);
void C_NSControl_TakeObjectValueFrom(void* ptr, void* sender);
void C_NSControl_TakeStringValueFrom(void* ptr, void* sender);
void C_NSControl_DrawWithExpansionFrame_InView(void* ptr, CGRect contentFrame, void* view);
CGRect C_NSControl_ExpansionFrameWithFrame(void* ptr, CGRect contentFrame);
bool C_NSControl_AbortEditing(void* ptr);
void* C_NSControl_CurrentEditor(void* ptr);
void C_NSControl_ValidateEditing(void* ptr);
void C_NSControl_EditWithFrame_Editor_Delegate_Event(void* ptr, CGRect rect, void* textObj, void* delegate, void* event);
void C_NSControl_EndEditing(void* ptr, void* textObj);
void C_NSControl_SelectWithFrame_Editor_Delegate_Start_Length(void* ptr, CGRect rect, void* textObj, void* delegate, int selStart, int selLength);
CGSize C_NSControl_SizeThatFits(void* ptr, CGSize size);
void C_NSControl_SizeToFit(void* ptr);
bool C_NSControl_SendAction_To(void* ptr, void* action, void* target);
void C_NSControl_PerformClick(void* ptr, void* sender);
void C_NSControl_InvalidateIntrinsicContentSizeForCell(void* ptr, void* cell);
bool C_NSControl_IsEnabled(void* ptr);
void C_NSControl_SetEnabled(void* ptr, bool value);
double C_NSControl_DoubleValue(void* ptr);
void C_NSControl_SetDoubleValue(void* ptr, double value);
float C_NSControl_FloatValue(void* ptr);
void C_NSControl_SetFloatValue(void* ptr, float value);
int C_NSControl_IntegerValue(void* ptr);
void C_NSControl_SetIntegerValue(void* ptr, int value);
void* C_NSControl_ObjectValue(void* ptr);
void C_NSControl_SetObjectValue(void* ptr, void* value);
void* C_NSControl_StringValue(void* ptr);
void C_NSControl_SetStringValue(void* ptr, void* value);
void* C_NSControl_AttributedStringValue(void* ptr);
void C_NSControl_SetAttributedStringValue(void* ptr, void* value);
int C_NSControl_Alignment(void* ptr);
void C_NSControl_SetAlignment(void* ptr, int value);
void* C_NSControl_Font(void* ptr);
void C_NSControl_SetFont(void* ptr, void* value);
unsigned int C_NSControl_LineBreakMode(void* ptr);
void C_NSControl_SetLineBreakMode(void* ptr, unsigned int value);
bool C_NSControl_UsesSingleLineMode(void* ptr);
void C_NSControl_SetUsesSingleLineMode(void* ptr, bool value);
void* C_NSControl_Formatter(void* ptr);
void C_NSControl_SetFormatter(void* ptr, void* value);
int C_NSControl_BaseWritingDirection(void* ptr);
void C_NSControl_SetBaseWritingDirection(void* ptr, int value);
bool C_NSControl_AllowsExpansionToolTips(void* ptr);
void C_NSControl_SetAllowsExpansionToolTips(void* ptr, bool value);
unsigned int C_NSControl_ControlSize(void* ptr);
void C_NSControl_SetControlSize(void* ptr, unsigned int value);
bool C_NSControl_IsHighlighted(void* ptr);
void C_NSControl_SetHighlighted(void* ptr, bool value);
void* C_NSControl_Action(void* ptr);
void C_NSControl_SetAction(void* ptr, void* value);
void* C_NSControl_Target(void* ptr);
void C_NSControl_SetTarget(void* ptr, void* value);
bool C_NSControl_IsContinuous(void* ptr);
void C_NSControl_SetContinuous(void* ptr, bool value);
void C_NSControl_SetTag(void* ptr, int value);
bool C_NSControl_RefusesFirstResponder(void* ptr);
void C_NSControl_SetRefusesFirstResponder(void* ptr, bool value);
bool C_NSControl_IgnoresMultiClick(void* ptr);
void C_NSControl_SetIgnoresMultiClick(void* ptr, bool value);
void* C_NSControl_Cell(void* ptr);
void C_NSControl_SetCell(void* ptr, void* value);
