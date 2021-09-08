package appkit

// #include "cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Cell interface {
	objc.Object
	SetCellAttribute_To(parameter CellAttribute, value int)
	CellAttribute(parameter CellAttribute) int
	SetNextState()
	SetUpFieldEditorAttributes(textObj Text) Text
	MenuForEvent_InRect_OfView(event Event, cellFrame foundation.Rect, view View) Menu
	Compare(otherCell objc.Object) foundation.ComparisonResult
	PerformClick(sender objc.Object)
	TakeObjectValueFrom(sender objc.Object)
	TakeIntegerValueFrom(sender objc.Object)
	TakeIntValueFrom(sender objc.Object)
	TakeStringValueFrom(sender objc.Object)
	TakeDoubleValueFrom(sender objc.Object)
	TakeFloatValueFrom(sender objc.Object)
	TrackMouse_InRect_OfView_UntilMouseUp(event Event, cellFrame foundation.Rect, controlView View, flag bool) bool
	StartTrackingAt_InView(startPoint foundation.Point, controlView View) bool
	ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView View) bool
	StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView View, flag bool)
	HitTestForEvent_InRect_OfView(event Event, cellFrame foundation.Rect, controlView View) CellHitResult
	ResetCursorRect_InView(cellFrame foundation.Rect, controlView View)
	DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view View) []DraggingImageComponent
	DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView View)
	FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView View) foundation.Rect
	CalcDrawInfo(rect foundation.Rect)
	CellSizeForBounds(rect foundation.Rect) foundation.Size
	DrawingRectForBounds(rect foundation.Rect) foundation.Rect
	ImageRectForBounds(rect foundation.Rect) foundation.Rect
	TitleRectForBounds(rect foundation.Rect) foundation.Rect
	DrawWithFrame_InView(cellFrame foundation.Rect, controlView View)
	HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView View) Color
	DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView View)
	Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView View)
	EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView View, textObj Text, delegate objc.Object, event Event)
	SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView View, textObj Text, delegate objc.Object, selStart int, selLength int)
	EndEditing(textObj Text)
	FieldEditorForView(controlView View) TextView
	ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view View) foundation.Rect
	DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view View)
	ObjectValue() objc.Object
	SetObjectValue(value objc.Object)
	HasValidObjectValue() bool
	IntegerValue() int
	SetIntegerValue(value int)
	StringValue() string
	SetStringValue(value string)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	Type() CellType
	SetType(value CellType)
	IsEnabled() bool
	SetEnabled(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	IsBordered() bool
	SetBordered(value bool)
	IsOpaque() bool
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	InteriorBackgroundStyle() BackgroundStyle
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	NextState() int
	State() ControlStateValue
	SetState(value ControlStateValue)
	IsEditable() bool
	SetEditable(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	IsScrollable() bool
	SetScrollable(value bool)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() Font
	SetFont(value Font)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	Wraps() bool
	SetWraps(value bool)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	Title() string
	SetTitle(value string)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.Object)
	IsContinuous() bool
	SetContinuous(value bool)
	Image() Image
	SetImage(value Image)
	Tag() int
	SetTag(value int)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.Formatter)
	Menu() Menu
	SetMenu(value Menu)
	AcceptsFirstResponder() bool
	ShowsFirstResponder() bool
	SetShowsFirstResponder(value bool)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.Object)
	MouseDownFlags() int
	KeyEquivalent() string
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	CellSize() foundation.Size
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlView() View
	SetControlView(value View)
	IsHighlighted() bool
	SetHighlighted(value bool)
	SendsActionOnEndEditing() bool
	SetSendsActionOnEndEditing(value bool)
	WantsNotificationForMarkedText() bool
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
}

type NSCell struct {
	objc.NSObject
}

func MakeCell(ptr unsafe.Pointer) NSCell {
	return NSCell{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSCell) InitImageCell(image Image) NSCell {
	result_ := C.C_NSCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakeCell(result_)
}

func (n NSCell) InitTextCell(_string string) NSCell {
	result_ := C.C_NSCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeCell(result_)
}

func (n NSCell) Init() NSCell {
	result_ := C.C_NSCell_Init(n.Ptr())
	return MakeCell(result_)
}

func (n NSCell) InitWithCoder(coder foundation.Coder) NSCell {
	result_ := C.C_NSCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCell(result_)
}

func AllocCell() NSCell {
	result_ := C.C_NSCell_AllocCell()
	return MakeCell(result_)
}

func NewCell() NSCell {
	result_ := C.C_NSCell_NewCell()
	return MakeCell(result_)
}

func (n NSCell) Autorelease() NSCell {
	result_ := C.C_NSCell_Autorelease(n.Ptr())
	return MakeCell(result_)
}

func (n NSCell) Retain() NSCell {
	result_ := C.C_NSCell_Retain(n.Ptr())
	return MakeCell(result_)
}

func (n NSCell) SetCellAttribute_To(parameter CellAttribute, value int) {
	C.C_NSCell_SetCellAttribute_To(n.Ptr(), C.uint(uint(parameter)), C.int(value))
}

func (n NSCell) CellAttribute(parameter CellAttribute) int {
	result_ := C.C_NSCell_CellAttribute(n.Ptr(), C.uint(uint(parameter)))
	return int(result_)
}

func (n NSCell) SetNextState() {
	C.C_NSCell_SetNextState(n.Ptr())
}

func (n NSCell) SetUpFieldEditorAttributes(textObj Text) Text {
	result_ := C.C_NSCell_SetUpFieldEditorAttributes(n.Ptr(), objc.ExtractPtr(textObj))
	return MakeText(result_)
}

func (n NSCell) MenuForEvent_InRect_OfView(event Event, cellFrame foundation.Rect, view View) Menu {
	result_ := C.C_NSCell_MenuForEvent_InRect_OfView(n.Ptr(), objc.ExtractPtr(event), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(view))
	return MakeMenu(result_)
}

func (n NSCell) Compare(otherCell objc.Object) foundation.ComparisonResult {
	result_ := C.C_NSCell_Compare(n.Ptr(), objc.ExtractPtr(otherCell))
	return foundation.ComparisonResult(int(result_))
}

func (n NSCell) PerformClick(sender objc.Object) {
	C.C_NSCell_PerformClick(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeObjectValueFrom(sender objc.Object) {
	C.C_NSCell_TakeObjectValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeIntegerValueFrom(sender objc.Object) {
	C.C_NSCell_TakeIntegerValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeIntValueFrom(sender objc.Object) {
	C.C_NSCell_TakeIntValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeStringValueFrom(sender objc.Object) {
	C.C_NSCell_TakeStringValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeDoubleValueFrom(sender objc.Object) {
	C.C_NSCell_TakeDoubleValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TakeFloatValueFrom(sender objc.Object) {
	C.C_NSCell_TakeFloatValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCell) TrackMouse_InRect_OfView_UntilMouseUp(event Event, cellFrame foundation.Rect, controlView View, flag bool) bool {
	result_ := C.C_NSCell_TrackMouse_InRect_OfView_UntilMouseUp(n.Ptr(), objc.ExtractPtr(event), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView), C.bool(flag))
	return bool(result_)
}

func (n NSCell) StartTrackingAt_InView(startPoint foundation.Point, controlView View) bool {
	result_ := C.C_NSCell_StartTrackingAt_InView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(startPoint))), objc.ExtractPtr(controlView))
	return bool(result_)
}

func (n NSCell) ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView View) bool {
	result_ := C.C_NSCell_ContinueTracking_At_InView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(lastPoint))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(currentPoint))), objc.ExtractPtr(controlView))
	return bool(result_)
}

func (n NSCell) StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView View, flag bool) {
	C.C_NSCell_StopTracking_At_InView_MouseIsUp(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(lastPoint))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(stopPoint))), objc.ExtractPtr(controlView), C.bool(flag))
}

func (n NSCell) HitTestForEvent_InRect_OfView(event Event, cellFrame foundation.Rect, controlView View) CellHitResult {
	result_ := C.C_NSCell_HitTestForEvent_InRect_OfView(n.Ptr(), objc.ExtractPtr(event), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
	return CellHitResult(uint(result_))
}

func (n NSCell) ResetCursorRect_InView(cellFrame foundation.Rect, controlView View) {
	C.C_NSCell_ResetCursorRect_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
}

func (n NSCell) DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view View) []DraggingImageComponent {
	result_ := C.C_NSCell_DraggingImageComponentsWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(view))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}

func (n NSCell) DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView View) {
	C.C_NSCell_DrawFocusRingMaskWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
}

func (n NSCell) FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView View) foundation.Rect {
	result_ := C.C_NSCell_FocusRingMaskBoundsForFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCell) CalcDrawInfo(rect foundation.Rect) {
	C.C_NSCell_CalcDrawInfo(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSCell) CellSizeForBounds(rect foundation.Rect) foundation.Size {
	result_ := C.C_NSCell_CellSizeForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCell) DrawingRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSCell_DrawingRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCell) ImageRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSCell_ImageRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCell) TitleRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSCell_TitleRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCell) DrawWithFrame_InView(cellFrame foundation.Rect, controlView View) {
	C.C_NSCell_DrawWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
}

func (n NSCell) HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView View) Color {
	result_ := C.C_NSCell_HighlightColorWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
	return MakeColor(result_)
}

func (n NSCell) DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView View) {
	C.C_NSCell_DrawInteriorWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
}

func (n NSCell) Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView View) {
	C.C_NSCell_Highlight_WithFrame_InView(n.Ptr(), C.bool(flag), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView))
}

func (n NSCell) EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView View, textObj Text, delegate objc.Object, event Event) {
	C.C_NSCell_EditWithFrame_InView_Editor_Delegate_Event(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(controlView), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), objc.ExtractPtr(event))
}

func (n NSCell) SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView View, textObj Text, delegate objc.Object, selStart int, selLength int) {
	C.C_NSCell_SelectWithFrame_InView_Editor_Delegate_Start_Length(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(controlView), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), C.int(selStart), C.int(selLength))
}

func (n NSCell) EndEditing(textObj Text) {
	C.C_NSCell_EndEditing(n.Ptr(), objc.ExtractPtr(textObj))
}

func (n NSCell) FieldEditorForView(controlView View) TextView {
	result_ := C.C_NSCell_FieldEditorForView(n.Ptr(), objc.ExtractPtr(controlView))
	return MakeTextView(result_)
}

func (n NSCell) ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view View) foundation.Rect {
	result_ := C.C_NSCell_ExpansionFrameWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCell) DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view View) {
	C.C_NSCell_DrawWithExpansionFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(view))
}

func (n NSCell) ObjectValue() objc.Object {
	result_ := C.C_NSCell_ObjectValue(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCell) SetObjectValue(value objc.Object) {
	C.C_NSCell_SetObjectValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) HasValidObjectValue() bool {
	result_ := C.C_NSCell_HasValidObjectValue(n.Ptr())
	return bool(result_)
}

func (n NSCell) IntegerValue() int {
	result_ := C.C_NSCell_IntegerValue(n.Ptr())
	return int(result_)
}

func (n NSCell) SetIntegerValue(value int) {
	C.C_NSCell_SetIntegerValue(n.Ptr(), C.int(value))
}

func (n NSCell) StringValue() string {
	result_ := C.C_NSCell_StringValue(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCell) SetStringValue(value string) {
	C.C_NSCell_SetStringValue(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSCell) DoubleValue() float64 {
	result_ := C.C_NSCell_DoubleValue(n.Ptr())
	return float64(result_)
}

func (n NSCell) SetDoubleValue(value float64) {
	C.C_NSCell_SetDoubleValue(n.Ptr(), C.double(value))
}

func (n NSCell) FloatValue() float32 {
	result_ := C.C_NSCell_FloatValue(n.Ptr())
	return float32(result_)
}

func (n NSCell) SetFloatValue(value float32) {
	C.C_NSCell_SetFloatValue(n.Ptr(), C.float(value))
}

func (n NSCell) Type() CellType {
	result_ := C.C_NSCell_Type(n.Ptr())
	return CellType(uint(result_))
}

func (n NSCell) SetType(value CellType) {
	C.C_NSCell_SetType(n.Ptr(), C.uint(uint(value)))
}

func (n NSCell) IsEnabled() bool {
	result_ := C.C_NSCell_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetEnabled(value bool) {
	C.C_NSCell_SetEnabled(n.Ptr(), C.bool(value))
}

func (n NSCell) AllowsUndo() bool {
	result_ := C.C_NSCell_AllowsUndo(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetAllowsUndo(value bool) {
	C.C_NSCell_SetAllowsUndo(n.Ptr(), C.bool(value))
}

func (n NSCell) IsBezeled() bool {
	result_ := C.C_NSCell_IsBezeled(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetBezeled(value bool) {
	C.C_NSCell_SetBezeled(n.Ptr(), C.bool(value))
}

func (n NSCell) IsBordered() bool {
	result_ := C.C_NSCell_IsBordered(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetBordered(value bool) {
	C.C_NSCell_SetBordered(n.Ptr(), C.bool(value))
}

func (n NSCell) IsOpaque() bool {
	result_ := C.C_NSCell_IsOpaque(n.Ptr())
	return bool(result_)
}

func (n NSCell) BackgroundStyle() BackgroundStyle {
	result_ := C.C_NSCell_BackgroundStyle(n.Ptr())
	return BackgroundStyle(int(result_))
}

func (n NSCell) SetBackgroundStyle(value BackgroundStyle) {
	C.C_NSCell_SetBackgroundStyle(n.Ptr(), C.int(int(value)))
}

func (n NSCell) InteriorBackgroundStyle() BackgroundStyle {
	result_ := C.C_NSCell_InteriorBackgroundStyle(n.Ptr())
	return BackgroundStyle(int(result_))
}

func (n NSCell) AllowsMixedState() bool {
	result_ := C.C_NSCell_AllowsMixedState(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetAllowsMixedState(value bool) {
	C.C_NSCell_SetAllowsMixedState(n.Ptr(), C.bool(value))
}

func (n NSCell) NextState() int {
	result_ := C.C_NSCell_NextState(n.Ptr())
	return int(result_)
}

func (n NSCell) State() ControlStateValue {
	result_ := C.C_NSCell_State(n.Ptr())
	return ControlStateValue(int(result_))
}

func (n NSCell) SetState(value ControlStateValue) {
	C.C_NSCell_SetState(n.Ptr(), C.int(int(value)))
}

func (n NSCell) IsEditable() bool {
	result_ := C.C_NSCell_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetEditable(value bool) {
	C.C_NSCell_SetEditable(n.Ptr(), C.bool(value))
}

func (n NSCell) IsSelectable() bool {
	result_ := C.C_NSCell_IsSelectable(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetSelectable(value bool) {
	C.C_NSCell_SetSelectable(n.Ptr(), C.bool(value))
}

func (n NSCell) IsScrollable() bool {
	result_ := C.C_NSCell_IsScrollable(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetScrollable(value bool) {
	C.C_NSCell_SetScrollable(n.Ptr(), C.bool(value))
}

func (n NSCell) Alignment() TextAlignment {
	result_ := C.C_NSCell_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSCell) SetAlignment(value TextAlignment) {
	C.C_NSCell_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSCell) Font() Font {
	result_ := C.C_NSCell_Font(n.Ptr())
	return MakeFont(result_)
}

func (n NSCell) SetFont(value Font) {
	C.C_NSCell_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) LineBreakMode() LineBreakMode {
	result_ := C.C_NSCell_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result_))
}

func (n NSCell) SetLineBreakMode(value LineBreakMode) {
	C.C_NSCell_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSCell) TruncatesLastVisibleLine() bool {
	result_ := C.C_NSCell_TruncatesLastVisibleLine(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetTruncatesLastVisibleLine(value bool) {
	C.C_NSCell_SetTruncatesLastVisibleLine(n.Ptr(), C.bool(value))
}

func (n NSCell) Wraps() bool {
	result_ := C.C_NSCell_Wraps(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetWraps(value bool) {
	C.C_NSCell_SetWraps(n.Ptr(), C.bool(value))
}

func (n NSCell) BaseWritingDirection() WritingDirection {
	result_ := C.C_NSCell_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result_))
}

func (n NSCell) SetBaseWritingDirection(value WritingDirection) {
	C.C_NSCell_SetBaseWritingDirection(n.Ptr(), C.int(int(value)))
}

func (n NSCell) AttributedStringValue() foundation.AttributedString {
	result_ := C.C_NSCell_AttributedStringValue(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSCell) SetAttributedStringValue(value foundation.AttributedString) {
	C.C_NSCell_SetAttributedStringValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) AllowsEditingTextAttributes() bool {
	result_ := C.C_NSCell_AllowsEditingTextAttributes(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetAllowsEditingTextAttributes(value bool) {
	C.C_NSCell_SetAllowsEditingTextAttributes(n.Ptr(), C.bool(value))
}

func (n NSCell) ImportsGraphics() bool {
	result_ := C.C_NSCell_ImportsGraphics(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetImportsGraphics(value bool) {
	C.C_NSCell_SetImportsGraphics(n.Ptr(), C.bool(value))
}

func (n NSCell) Title() string {
	result_ := C.C_NSCell_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCell) SetTitle(value string) {
	C.C_NSCell_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSCell) Action() objc.Selector {
	result_ := C.C_NSCell_Action(n.Ptr())
	return objc.Selector(result_)
}

func (n NSCell) SetAction(value objc.Selector) {
	C.C_NSCell_SetAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSCell) Target() objc.Object {
	result_ := C.C_NSCell_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCell) SetTarget(value objc.Object) {
	C.C_NSCell_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) IsContinuous() bool {
	result_ := C.C_NSCell_IsContinuous(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetContinuous(value bool) {
	C.C_NSCell_SetContinuous(n.Ptr(), C.bool(value))
}

func (n NSCell) Image() Image {
	result_ := C.C_NSCell_Image(n.Ptr())
	return MakeImage(result_)
}

func (n NSCell) SetImage(value Image) {
	C.C_NSCell_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) Tag() int {
	result_ := C.C_NSCell_Tag(n.Ptr())
	return int(result_)
}

func (n NSCell) SetTag(value int) {
	C.C_NSCell_SetTag(n.Ptr(), C.int(value))
}

func (n NSCell) Formatter() foundation.Formatter {
	result_ := C.C_NSCell_Formatter(n.Ptr())
	return foundation.MakeFormatter(result_)
}

func (n NSCell) SetFormatter(value foundation.Formatter) {
	C.C_NSCell_SetFormatter(n.Ptr(), objc.ExtractPtr(value))
}

func Cell_DefaultMenu() Menu {
	result_ := C.C_NSCell_Cell_DefaultMenu()
	return MakeMenu(result_)
}

func (n NSCell) Menu() Menu {
	result_ := C.C_NSCell_Menu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSCell) SetMenu(value Menu) {
	C.C_NSCell_SetMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) AcceptsFirstResponder() bool {
	result_ := C.C_NSCell_AcceptsFirstResponder(n.Ptr())
	return bool(result_)
}

func (n NSCell) ShowsFirstResponder() bool {
	result_ := C.C_NSCell_ShowsFirstResponder(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetShowsFirstResponder(value bool) {
	C.C_NSCell_SetShowsFirstResponder(n.Ptr(), C.bool(value))
}

func (n NSCell) RefusesFirstResponder() bool {
	result_ := C.C_NSCell_RefusesFirstResponder(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetRefusesFirstResponder(value bool) {
	C.C_NSCell_SetRefusesFirstResponder(n.Ptr(), C.bool(value))
}

func (n NSCell) RepresentedObject() objc.Object {
	result_ := C.C_NSCell_RepresentedObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCell) SetRepresentedObject(value objc.Object) {
	C.C_NSCell_SetRepresentedObject(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) MouseDownFlags() int {
	result_ := C.C_NSCell_MouseDownFlags(n.Ptr())
	return int(result_)
}

func Cell_PrefersTrackingUntilMouseUp() bool {
	result_ := C.C_NSCell_Cell_PrefersTrackingUntilMouseUp()
	return bool(result_)
}

func (n NSCell) KeyEquivalent() string {
	result_ := C.C_NSCell_KeyEquivalent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func Cell_DefaultFocusRingType() FocusRingType {
	result_ := C.C_NSCell_Cell_DefaultFocusRingType()
	return FocusRingType(uint(result_))
}

func (n NSCell) FocusRingType() FocusRingType {
	result_ := C.C_NSCell_FocusRingType(n.Ptr())
	return FocusRingType(uint(result_))
}

func (n NSCell) SetFocusRingType(value FocusRingType) {
	C.C_NSCell_SetFocusRingType(n.Ptr(), C.uint(uint(value)))
}

func (n NSCell) CellSize() foundation.Size {
	result_ := C.C_NSCell_CellSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCell) ControlSize() ControlSize {
	result_ := C.C_NSCell_ControlSize(n.Ptr())
	return ControlSize(uint(result_))
}

func (n NSCell) SetControlSize(value ControlSize) {
	C.C_NSCell_SetControlSize(n.Ptr(), C.uint(uint(value)))
}

func (n NSCell) ControlView() View {
	result_ := C.C_NSCell_ControlView(n.Ptr())
	return MakeView(result_)
}

func (n NSCell) SetControlView(value View) {
	C.C_NSCell_SetControlView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCell) IsHighlighted() bool {
	result_ := C.C_NSCell_IsHighlighted(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetHighlighted(value bool) {
	C.C_NSCell_SetHighlighted(n.Ptr(), C.bool(value))
}

func (n NSCell) SendsActionOnEndEditing() bool {
	result_ := C.C_NSCell_SendsActionOnEndEditing(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetSendsActionOnEndEditing(value bool) {
	C.C_NSCell_SetSendsActionOnEndEditing(n.Ptr(), C.bool(value))
}

func (n NSCell) WantsNotificationForMarkedText() bool {
	result_ := C.C_NSCell_WantsNotificationForMarkedText(n.Ptr())
	return bool(result_)
}

func (n NSCell) UsesSingleLineMode() bool {
	result_ := C.C_NSCell_UsesSingleLineMode(n.Ptr())
	return bool(result_)
}

func (n NSCell) SetUsesSingleLineMode(value bool) {
	C.C_NSCell_SetUsesSingleLineMode(n.Ptr(), C.bool(value))
}

func (n NSCell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	result_ := C.C_NSCell_UserInterfaceLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result_))
}

func (n NSCell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	C.C_NSCell_SetUserInterfaceLayoutDirection(n.Ptr(), C.int(int(value)))
}
