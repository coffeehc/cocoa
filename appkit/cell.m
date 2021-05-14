#import <Appkit/Appkit.h>
#import "cell.h"

void* C_Cell_Alloc() {
	return [NSCell alloc];
}

void* C_NSCell_InitImageCell(void* ptr, void* image) {
	NSCell* nSCell = (NSCell*)ptr;
	NSCell* result = [nSCell initImageCell:(NSImage*)image];
	return result;
}

void* C_NSCell_InitTextCell(void* ptr, void* _string) {
	NSCell* nSCell = (NSCell*)ptr;
	NSCell* result = [nSCell initTextCell:(NSString*)_string];
	return result;
}

void* C_NSCell_Init(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSCell* result = [nSCell init];
	return result;
}

void* C_NSCell_InitWithCoder(void* ptr, void* coder) {
	NSCell* nSCell = (NSCell*)ptr;
	NSCell* result = [nSCell initWithCoder:(NSCoder*)coder];
	return result;
}

void C_NSCell_SetNextState(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setNextState];
}

void* C_NSCell_SetUpFieldEditorAttributes(void* ptr, void* textObj) {
	NSCell* nSCell = (NSCell*)ptr;
	NSText* result = [nSCell setUpFieldEditorAttributes:(NSText*)textObj];
	return result;
}

void* C_NSCell_MenuForEvent_InRect_OfView(void* ptr, void* event, CGRect cellFrame, void* view) {
	NSCell* nSCell = (NSCell*)ptr;
	NSMenu* result = [nSCell menuForEvent:(NSEvent*)event inRect:cellFrame ofView:(NSView*)view];
	return result;
}

void C_NSCell_PerformClick(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell performClick:(id)sender];
}

void C_NSCell_TakeObjectValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeObjectValueFrom:(id)sender];
}

void C_NSCell_TakeIntegerValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeIntegerValueFrom:(id)sender];
}

void C_NSCell_TakeIntValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeIntValueFrom:(id)sender];
}

void C_NSCell_TakeStringValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeStringValueFrom:(id)sender];
}

void C_NSCell_TakeDoubleValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeDoubleValueFrom:(id)sender];
}

void C_NSCell_TakeFloatValueFrom(void* ptr, void* sender) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell takeFloatValueFrom:(id)sender];
}

bool C_NSCell_TrackMouse_InRect_OfView_UntilMouseUp(void* ptr, void* event, CGRect cellFrame, void* controlView, bool flag) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell trackMouse:(NSEvent*)event inRect:cellFrame ofView:(NSView*)controlView untilMouseUp:flag];
	return result;
}

bool C_NSCell_StartTrackingAt_InView(void* ptr, CGPoint startPoint, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell startTrackingAt:startPoint inView:(NSView*)controlView];
	return result;
}

bool C_NSCell_ContinueTracking_At_InView(void* ptr, CGPoint lastPoint, CGPoint currentPoint, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell continueTracking:lastPoint at:currentPoint inView:(NSView*)controlView];
	return result;
}

void C_NSCell_StopTracking_At_InView_MouseIsUp(void* ptr, CGPoint lastPoint, CGPoint stopPoint, void* controlView, bool flag) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell stopTracking:lastPoint at:stopPoint inView:(NSView*)controlView mouseIsUp:flag];
}

void C_NSCell_ResetCursorRect_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell resetCursorRect:cellFrame inView:(NSView*)controlView];
}

void C_NSCell_DrawFocusRingMaskWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell drawFocusRingMaskWithFrame:cellFrame inView:(NSView*)controlView];
}

CGRect C_NSCell_FocusRingMaskBoundsForFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	CGRect result = [nSCell focusRingMaskBoundsForFrame:cellFrame inView:(NSView*)controlView];
	return result;
}

void C_NSCell_CalcDrawInfo(void* ptr, CGRect rect) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell calcDrawInfo:rect];
}

CGSize C_NSCell_CellSizeForBounds(void* ptr, CGRect rect) {
	NSCell* nSCell = (NSCell*)ptr;
	CGSize result = [nSCell cellSizeForBounds:rect];
	return result;
}

CGRect C_NSCell_DrawingRectForBounds(void* ptr, CGRect rect) {
	NSCell* nSCell = (NSCell*)ptr;
	CGRect result = [nSCell drawingRectForBounds:rect];
	return result;
}

CGRect C_NSCell_ImageRectForBounds(void* ptr, CGRect rect) {
	NSCell* nSCell = (NSCell*)ptr;
	CGRect result = [nSCell imageRectForBounds:rect];
	return result;
}

CGRect C_NSCell_TitleRectForBounds(void* ptr, CGRect rect) {
	NSCell* nSCell = (NSCell*)ptr;
	CGRect result = [nSCell titleRectForBounds:rect];
	return result;
}

void C_NSCell_DrawWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell drawWithFrame:cellFrame inView:(NSView*)controlView];
}

void* C_NSCell_HighlightColorWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	NSColor* result = [nSCell highlightColorWithFrame:cellFrame inView:(NSView*)controlView];
	return result;
}

void C_NSCell_DrawInteriorWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell drawInteriorWithFrame:cellFrame inView:(NSView*)controlView];
}

void C_NSCell_Highlight_WithFrame_InView(void* ptr, bool flag, CGRect cellFrame, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell highlight:flag withFrame:cellFrame inView:(NSView*)controlView];
}

void C_NSCell_EditWithFrame_InView_Editor_Delegate_Event(void* ptr, CGRect rect, void* controlView, void* textObj, void* delegate, void* event) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell editWithFrame:rect inView:(NSView*)controlView editor:(NSText*)textObj delegate:(id)delegate event:(NSEvent*)event];
}

void C_NSCell_SelectWithFrame_InView_Editor_Delegate_Start_Length(void* ptr, CGRect rect, void* controlView, void* textObj, void* delegate, int selStart, int selLength) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell selectWithFrame:rect inView:(NSView*)controlView editor:(NSText*)textObj delegate:(id)delegate start:selStart length:selLength];
}

void C_NSCell_EndEditing(void* ptr, void* textObj) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell endEditing:(NSText*)textObj];
}

void* C_NSCell_FieldEditorForView(void* ptr, void* controlView) {
	NSCell* nSCell = (NSCell*)ptr;
	NSTextView* result = [nSCell fieldEditorForView:(NSView*)controlView];
	return result;
}

CGRect C_NSCell_ExpansionFrameWithFrame_InView(void* ptr, CGRect cellFrame, void* view) {
	NSCell* nSCell = (NSCell*)ptr;
	CGRect result = [nSCell expansionFrameWithFrame:cellFrame inView:(NSView*)view];
	return result;
}

void C_NSCell_DrawWithExpansionFrame_InView(void* ptr, CGRect cellFrame, void* view) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell drawWithExpansionFrame:cellFrame inView:(NSView*)view];
}

void* C_NSCell_ObjectValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	id result = [nSCell objectValue];
	return result;
}

void C_NSCell_SetObjectValue(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setObjectValue:(id)value];
}

bool C_NSCell_HasValidObjectValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell hasValidObjectValue];
	return result;
}

int C_NSCell_IntegerValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	int result = [nSCell integerValue];
	return result;
}

void C_NSCell_SetIntegerValue(void* ptr, int value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setIntegerValue:value];
}

void* C_NSCell_StringValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSString* result = [nSCell stringValue];
	return result;
}

void C_NSCell_SetStringValue(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setStringValue:(NSString*)value];
}

double C_NSCell_DoubleValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	double result = [nSCell doubleValue];
	return result;
}

void C_NSCell_SetDoubleValue(void* ptr, double value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setDoubleValue:value];
}

float C_NSCell_FloatValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	float result = [nSCell floatValue];
	return result;
}

void C_NSCell_SetFloatValue(void* ptr, float value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setFloatValue:value];
}

bool C_NSCell_IsEnabled(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isEnabled];
	return result;
}

void C_NSCell_SetEnabled(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setEnabled:value];
}

bool C_NSCell_AllowsUndo(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell allowsUndo];
	return result;
}

void C_NSCell_SetAllowsUndo(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setAllowsUndo:value];
}

bool C_NSCell_IsBezeled(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isBezeled];
	return result;
}

void C_NSCell_SetBezeled(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setBezeled:value];
}

bool C_NSCell_IsBordered(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isBordered];
	return result;
}

void C_NSCell_SetBordered(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setBordered:value];
}

bool C_NSCell_IsOpaque(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isOpaque];
	return result;
}

bool C_NSCell_AllowsMixedState(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell allowsMixedState];
	return result;
}

void C_NSCell_SetAllowsMixedState(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setAllowsMixedState:value];
}

int C_NSCell_NextState(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	int result = [nSCell nextState];
	return result;
}

int C_NSCell_State(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	int result = [nSCell state];
	return result;
}

void C_NSCell_SetState(void* ptr, int value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setState:value];
}

bool C_NSCell_IsEditable(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isEditable];
	return result;
}

void C_NSCell_SetEditable(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setEditable:value];
}

bool C_NSCell_IsSelectable(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isSelectable];
	return result;
}

void C_NSCell_SetSelectable(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setSelectable:value];
}

bool C_NSCell_IsScrollable(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isScrollable];
	return result;
}

void C_NSCell_SetScrollable(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setScrollable:value];
}

void* C_NSCell_Font(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSFont* result = [nSCell font];
	return result;
}

void C_NSCell_SetFont(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setFont:(NSFont*)value];
}

bool C_NSCell_TruncatesLastVisibleLine(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell truncatesLastVisibleLine];
	return result;
}

void C_NSCell_SetTruncatesLastVisibleLine(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setTruncatesLastVisibleLine:value];
}

bool C_NSCell_Wraps(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell wraps];
	return result;
}

void C_NSCell_SetWraps(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setWraps:value];
}

void* C_NSCell_AttributedStringValue(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSAttributedString* result = [nSCell attributedStringValue];
	return result;
}

void C_NSCell_SetAttributedStringValue(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setAttributedStringValue:(NSAttributedString*)value];
}

bool C_NSCell_AllowsEditingTextAttributes(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell allowsEditingTextAttributes];
	return result;
}

void C_NSCell_SetAllowsEditingTextAttributes(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setAllowsEditingTextAttributes:value];
}

bool C_NSCell_ImportsGraphics(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell importsGraphics];
	return result;
}

void C_NSCell_SetImportsGraphics(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setImportsGraphics:value];
}

void* C_NSCell_Title(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSString* result = [nSCell title];
	return result;
}

void C_NSCell_SetTitle(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setTitle:(NSString*)value];
}

void* C_NSCell_Action(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	SEL result = [nSCell action];
	return result;
}

void C_NSCell_SetAction(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setAction:(SEL)value];
}

void* C_NSCell_Target(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	id result = [nSCell target];
	return result;
}

void C_NSCell_SetTarget(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setTarget:(id)value];
}

bool C_NSCell_IsContinuous(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isContinuous];
	return result;
}

void C_NSCell_SetContinuous(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setContinuous:value];
}

void* C_NSCell_Image(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSImage* result = [nSCell image];
	return result;
}

void C_NSCell_SetImage(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setImage:(NSImage*)value];
}

int C_NSCell_Tag(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	int result = [nSCell tag];
	return result;
}

void C_NSCell_SetTag(void* ptr, int value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setTag:value];
}

void* C_NSCell_Menu(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSMenu* result = [nSCell menu];
	return result;
}

void C_NSCell_SetMenu(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setMenu:(NSMenu*)value];
}

bool C_NSCell_AcceptsFirstResponder(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell acceptsFirstResponder];
	return result;
}

bool C_NSCell_ShowsFirstResponder(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell showsFirstResponder];
	return result;
}

void C_NSCell_SetShowsFirstResponder(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setShowsFirstResponder:value];
}

bool C_NSCell_RefusesFirstResponder(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell refusesFirstResponder];
	return result;
}

void C_NSCell_SetRefusesFirstResponder(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setRefusesFirstResponder:value];
}

void* C_NSCell_RepresentedObject(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	id result = [nSCell representedObject];
	return result;
}

void C_NSCell_SetRepresentedObject(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setRepresentedObject:(id)value];
}

int C_NSCell_MouseDownFlags(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	int result = [nSCell mouseDownFlags];
	return result;
}

void* C_NSCell_KeyEquivalent(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSString* result = [nSCell keyEquivalent];
	return result;
}

CGSize C_NSCell_CellSize(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	CGSize result = [nSCell cellSize];
	return result;
}

void* C_NSCell_ControlView(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	NSView* result = [nSCell controlView];
	return result;
}

void C_NSCell_SetControlView(void* ptr, void* value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setControlView:(NSView*)value];
}

bool C_NSCell_IsHighlighted(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell isHighlighted];
	return result;
}

void C_NSCell_SetHighlighted(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setHighlighted:value];
}

bool C_NSCell_SendsActionOnEndEditing(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell sendsActionOnEndEditing];
	return result;
}

void C_NSCell_SetSendsActionOnEndEditing(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setSendsActionOnEndEditing:value];
}

bool C_NSCell_WantsNotificationForMarkedText(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell wantsNotificationForMarkedText];
	return result;
}

bool C_NSCell_UsesSingleLineMode(void* ptr) {
	NSCell* nSCell = (NSCell*)ptr;
	bool result = [nSCell usesSingleLineMode];
	return result;
}

void C_NSCell_SetUsesSingleLineMode(void* ptr, bool value) {
	NSCell* nSCell = (NSCell*)ptr;
	[nSCell setUsesSingleLineMode:value];
}
