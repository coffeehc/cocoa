#import <Appkit/Appkit.h>
#import "cell.h"

void* C_Cell_Alloc() {
    return [NSCell alloc];
}

void* C_NSCell_InitImageCell(void* ptr, void* image) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCell* result_ = [nSCell initImageCell:(NSImage*)image];
    return result_;
}

void* C_NSCell_InitTextCell(void* ptr, void* _string) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCell* result_ = [nSCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSCell_Init(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCell* result_ = [nSCell init];
    return result_;
}

void* C_NSCell_InitWithCoder(void* ptr, void* coder) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCell* result_ = [nSCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void C_NSCell_SetCellAttribute_To(void* ptr, unsigned int parameter, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setCellAttribute:parameter to:value];
}

int C_NSCell_CellAttribute(void* ptr, unsigned int parameter) {
    NSCell* nSCell = (NSCell*)ptr;
    NSInteger result_ = [nSCell cellAttribute:parameter];
    return result_;
}

void C_NSCell_SetNextState(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setNextState];
}

void* C_NSCell_SetUpFieldEditorAttributes(void* ptr, void* textObj) {
    NSCell* nSCell = (NSCell*)ptr;
    NSText* result_ = [nSCell setUpFieldEditorAttributes:(NSText*)textObj];
    return result_;
}

void* C_NSCell_MenuForEvent_InRect_OfView(void* ptr, void* event, CGRect cellFrame, void* view) {
    NSCell* nSCell = (NSCell*)ptr;
    NSMenu* result_ = [nSCell menuForEvent:(NSEvent*)event inRect:cellFrame ofView:(NSView*)view];
    return result_;
}

int C_NSCell_Compare(void* ptr, void* otherCell) {
    NSCell* nSCell = (NSCell*)ptr;
    NSComparisonResult result_ = [nSCell compare:(id)otherCell];
    return result_;
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
    BOOL result_ = [nSCell trackMouse:(NSEvent*)event inRect:cellFrame ofView:(NSView*)controlView untilMouseUp:flag];
    return result_;
}

bool C_NSCell_StartTrackingAt_InView(void* ptr, CGPoint startPoint, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell startTrackingAt:startPoint inView:(NSView*)controlView];
    return result_;
}

bool C_NSCell_ContinueTracking_At_InView(void* ptr, CGPoint lastPoint, CGPoint currentPoint, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell continueTracking:lastPoint at:currentPoint inView:(NSView*)controlView];
    return result_;
}

void C_NSCell_StopTracking_At_InView_MouseIsUp(void* ptr, CGPoint lastPoint, CGPoint stopPoint, void* controlView, bool flag) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell stopTracking:lastPoint at:stopPoint inView:(NSView*)controlView mouseIsUp:flag];
}

unsigned int C_NSCell_HitTestForEvent_InRect_OfView(void* ptr, void* event, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCellHitResult result_ = [nSCell hitTestForEvent:(NSEvent*)event inRect:cellFrame ofView:(NSView*)controlView];
    return result_;
}

void C_NSCell_ResetCursorRect_InView(void* ptr, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell resetCursorRect:cellFrame inView:(NSView*)controlView];
}

Array C_NSCell_DraggingImageComponentsWithFrame_InView(void* ptr, CGRect frame, void* view) {
    NSCell* nSCell = (NSCell*)ptr;
    NSArray* result_ = [nSCell draggingImageComponentsWithFrame:frame inView:(NSView*)view];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSCell_DrawFocusRingMaskWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell drawFocusRingMaskWithFrame:cellFrame inView:(NSView*)controlView];
}

CGRect C_NSCell_FocusRingMaskBoundsForFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    NSRect result_ = [nSCell focusRingMaskBoundsForFrame:cellFrame inView:(NSView*)controlView];
    return result_;
}

void C_NSCell_CalcDrawInfo(void* ptr, CGRect rect) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell calcDrawInfo:rect];
}

CGSize C_NSCell_CellSizeForBounds(void* ptr, CGRect rect) {
    NSCell* nSCell = (NSCell*)ptr;
    NSSize result_ = [nSCell cellSizeForBounds:rect];
    return result_;
}

CGRect C_NSCell_DrawingRectForBounds(void* ptr, CGRect rect) {
    NSCell* nSCell = (NSCell*)ptr;
    NSRect result_ = [nSCell drawingRectForBounds:rect];
    return result_;
}

CGRect C_NSCell_ImageRectForBounds(void* ptr, CGRect rect) {
    NSCell* nSCell = (NSCell*)ptr;
    NSRect result_ = [nSCell imageRectForBounds:rect];
    return result_;
}

CGRect C_NSCell_TitleRectForBounds(void* ptr, CGRect rect) {
    NSCell* nSCell = (NSCell*)ptr;
    NSRect result_ = [nSCell titleRectForBounds:rect];
    return result_;
}

void C_NSCell_DrawWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell drawWithFrame:cellFrame inView:(NSView*)controlView];
}

void* C_NSCell_HighlightColorWithFrame_InView(void* ptr, CGRect cellFrame, void* controlView) {
    NSCell* nSCell = (NSCell*)ptr;
    NSColor* result_ = [nSCell highlightColorWithFrame:cellFrame inView:(NSView*)controlView];
    return result_;
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
    NSTextView* result_ = [nSCell fieldEditorForView:(NSView*)controlView];
    return result_;
}

CGRect C_NSCell_ExpansionFrameWithFrame_InView(void* ptr, CGRect cellFrame, void* view) {
    NSCell* nSCell = (NSCell*)ptr;
    NSRect result_ = [nSCell expansionFrameWithFrame:cellFrame inView:(NSView*)view];
    return result_;
}

void C_NSCell_DrawWithExpansionFrame_InView(void* ptr, CGRect cellFrame, void* view) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell drawWithExpansionFrame:cellFrame inView:(NSView*)view];
}

void* C_NSCell_ObjectValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    id result_ = [nSCell objectValue];
    return result_;
}

void C_NSCell_SetObjectValue(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setObjectValue:(id)value];
}

bool C_NSCell_HasValidObjectValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell hasValidObjectValue];
    return result_;
}

int C_NSCell_IntegerValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSInteger result_ = [nSCell integerValue];
    return result_;
}

void C_NSCell_SetIntegerValue(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setIntegerValue:value];
}

void* C_NSCell_StringValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSString* result_ = [nSCell stringValue];
    return result_;
}

void C_NSCell_SetStringValue(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setStringValue:(NSString*)value];
}

double C_NSCell_DoubleValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    double result_ = [nSCell doubleValue];
    return result_;
}

void C_NSCell_SetDoubleValue(void* ptr, double value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setDoubleValue:value];
}

float C_NSCell_FloatValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    float result_ = [nSCell floatValue];
    return result_;
}

void C_NSCell_SetFloatValue(void* ptr, float value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setFloatValue:value];
}

unsigned int C_NSCell_Type(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSCellType result_ = [nSCell type];
    return result_;
}

void C_NSCell_SetType(void* ptr, unsigned int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setType:value];
}

bool C_NSCell_IsEnabled(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isEnabled];
    return result_;
}

void C_NSCell_SetEnabled(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setEnabled:value];
}

bool C_NSCell_AllowsUndo(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell allowsUndo];
    return result_;
}

void C_NSCell_SetAllowsUndo(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAllowsUndo:value];
}

bool C_NSCell_IsBezeled(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isBezeled];
    return result_;
}

void C_NSCell_SetBezeled(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setBezeled:value];
}

bool C_NSCell_IsBordered(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isBordered];
    return result_;
}

void C_NSCell_SetBordered(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setBordered:value];
}

bool C_NSCell_IsOpaque(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isOpaque];
    return result_;
}

int C_NSCell_BackgroundStyle(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSBackgroundStyle result_ = [nSCell backgroundStyle];
    return result_;
}

void C_NSCell_SetBackgroundStyle(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setBackgroundStyle:value];
}

int C_NSCell_InteriorBackgroundStyle(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSBackgroundStyle result_ = [nSCell interiorBackgroundStyle];
    return result_;
}

bool C_NSCell_AllowsMixedState(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell allowsMixedState];
    return result_;
}

void C_NSCell_SetAllowsMixedState(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAllowsMixedState:value];
}

int C_NSCell_NextState(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSInteger result_ = [nSCell nextState];
    return result_;
}

int C_NSCell_State(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSControlStateValue result_ = [nSCell state];
    return result_;
}

void C_NSCell_SetState(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setState:value];
}

bool C_NSCell_IsEditable(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isEditable];
    return result_;
}

void C_NSCell_SetEditable(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setEditable:value];
}

bool C_NSCell_IsSelectable(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isSelectable];
    return result_;
}

void C_NSCell_SetSelectable(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setSelectable:value];
}

bool C_NSCell_IsScrollable(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isScrollable];
    return result_;
}

void C_NSCell_SetScrollable(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setScrollable:value];
}

int C_NSCell_Alignment(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSTextAlignment result_ = [nSCell alignment];
    return result_;
}

void C_NSCell_SetAlignment(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAlignment:value];
}

void* C_NSCell_Font(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSFont* result_ = [nSCell font];
    return result_;
}

void C_NSCell_SetFont(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setFont:(NSFont*)value];
}

unsigned int C_NSCell_LineBreakMode(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSLineBreakMode result_ = [nSCell lineBreakMode];
    return result_;
}

void C_NSCell_SetLineBreakMode(void* ptr, unsigned int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setLineBreakMode:value];
}

bool C_NSCell_TruncatesLastVisibleLine(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell truncatesLastVisibleLine];
    return result_;
}

void C_NSCell_SetTruncatesLastVisibleLine(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setTruncatesLastVisibleLine:value];
}

bool C_NSCell_Wraps(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell wraps];
    return result_;
}

void C_NSCell_SetWraps(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setWraps:value];
}

int C_NSCell_BaseWritingDirection(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSWritingDirection result_ = [nSCell baseWritingDirection];
    return result_;
}

void C_NSCell_SetBaseWritingDirection(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setBaseWritingDirection:value];
}

void* C_NSCell_AttributedStringValue(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSAttributedString* result_ = [nSCell attributedStringValue];
    return result_;
}

void C_NSCell_SetAttributedStringValue(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAttributedStringValue:(NSAttributedString*)value];
}

bool C_NSCell_AllowsEditingTextAttributes(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell allowsEditingTextAttributes];
    return result_;
}

void C_NSCell_SetAllowsEditingTextAttributes(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAllowsEditingTextAttributes:value];
}

bool C_NSCell_ImportsGraphics(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell importsGraphics];
    return result_;
}

void C_NSCell_SetImportsGraphics(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setImportsGraphics:value];
}

void* C_NSCell_Title(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSString* result_ = [nSCell title];
    return result_;
}

void C_NSCell_SetTitle(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setTitle:(NSString*)value];
}

void* C_NSCell_Action(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    SEL result_ = [nSCell action];
    return result_;
}

void C_NSCell_SetAction(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setAction:(SEL)value];
}

void* C_NSCell_Target(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    id result_ = [nSCell target];
    return result_;
}

void C_NSCell_SetTarget(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setTarget:(id)value];
}

bool C_NSCell_IsContinuous(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isContinuous];
    return result_;
}

void C_NSCell_SetContinuous(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setContinuous:value];
}

void* C_NSCell_Image(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSImage* result_ = [nSCell image];
    return result_;
}

void C_NSCell_SetImage(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setImage:(NSImage*)value];
}

int C_NSCell_Tag(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSInteger result_ = [nSCell tag];
    return result_;
}

void C_NSCell_SetTag(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setTag:value];
}

void* C_NSCell_Formatter(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSFormatter* result_ = [nSCell formatter];
    return result_;
}

void C_NSCell_SetFormatter(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setFormatter:(NSFormatter*)value];
}

void* C_NSCell_Cell_DefaultMenu() {
    NSMenu* result_ = [NSCell defaultMenu];
    return result_;
}

void* C_NSCell_Menu(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSMenu* result_ = [nSCell menu];
    return result_;
}

void C_NSCell_SetMenu(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setMenu:(NSMenu*)value];
}

bool C_NSCell_AcceptsFirstResponder(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell acceptsFirstResponder];
    return result_;
}

bool C_NSCell_ShowsFirstResponder(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell showsFirstResponder];
    return result_;
}

void C_NSCell_SetShowsFirstResponder(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setShowsFirstResponder:value];
}

bool C_NSCell_RefusesFirstResponder(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell refusesFirstResponder];
    return result_;
}

void C_NSCell_SetRefusesFirstResponder(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setRefusesFirstResponder:value];
}

void* C_NSCell_RepresentedObject(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    id result_ = [nSCell representedObject];
    return result_;
}

void C_NSCell_SetRepresentedObject(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setRepresentedObject:(id)value];
}

int C_NSCell_MouseDownFlags(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSInteger result_ = [nSCell mouseDownFlags];
    return result_;
}

bool C_NSCell_Cell_PrefersTrackingUntilMouseUp() {
    BOOL result_ = [NSCell prefersTrackingUntilMouseUp];
    return result_;
}

void* C_NSCell_KeyEquivalent(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSString* result_ = [nSCell keyEquivalent];
    return result_;
}

unsigned int C_NSCell_Cell_DefaultFocusRingType() {
    NSFocusRingType result_ = [NSCell defaultFocusRingType];
    return result_;
}

unsigned int C_NSCell_FocusRingType(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSFocusRingType result_ = [nSCell focusRingType];
    return result_;
}

void C_NSCell_SetFocusRingType(void* ptr, unsigned int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setFocusRingType:value];
}

CGSize C_NSCell_CellSize(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSSize result_ = [nSCell cellSize];
    return result_;
}

unsigned int C_NSCell_ControlSize(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSControlSize result_ = [nSCell controlSize];
    return result_;
}

void C_NSCell_SetControlSize(void* ptr, unsigned int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setControlSize:value];
}

void* C_NSCell_ControlView(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSView* result_ = [nSCell controlView];
    return result_;
}

void C_NSCell_SetControlView(void* ptr, void* value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setControlView:(NSView*)value];
}

bool C_NSCell_IsHighlighted(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell isHighlighted];
    return result_;
}

void C_NSCell_SetHighlighted(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setHighlighted:value];
}

bool C_NSCell_SendsActionOnEndEditing(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell sendsActionOnEndEditing];
    return result_;
}

void C_NSCell_SetSendsActionOnEndEditing(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setSendsActionOnEndEditing:value];
}

bool C_NSCell_WantsNotificationForMarkedText(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell wantsNotificationForMarkedText];
    return result_;
}

bool C_NSCell_UsesSingleLineMode(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    BOOL result_ = [nSCell usesSingleLineMode];
    return result_;
}

void C_NSCell_SetUsesSingleLineMode(void* ptr, bool value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setUsesSingleLineMode:value];
}

int C_NSCell_UserInterfaceLayoutDirection(void* ptr) {
    NSCell* nSCell = (NSCell*)ptr;
    NSUserInterfaceLayoutDirection result_ = [nSCell userInterfaceLayoutDirection];
    return result_;
}

void C_NSCell_SetUserInterfaceLayoutDirection(void* ptr, int value) {
    NSCell* nSCell = (NSCell*)ptr;
    [nSCell setUserInterfaceLayoutDirection:value];
}
