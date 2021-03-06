package appkit

// #include "control.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type Control interface {
	View
	TakeDoubleValueFrom(sender objc.Object)
	TakeFloatValueFrom(sender objc.Object)
	TakeIntValueFrom(sender objc.Object)
	TakeIntegerValueFrom(sender objc.Object)
	TakeObjectValueFrom(sender objc.Object)
	TakeStringValueFrom(sender objc.Object)
	DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view View)
	ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect
	AbortEditing() bool
	CurrentEditor() Text
	ValidateEditing()
	EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj Text, delegate objc.Object, event Event)
	EndEditing(textObj Text)
	SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj Text, delegate objc.Object, selStart int, selLength int)
	SizeThatFits(size foundation.Size) foundation.Size
	SizeToFit()
	SendAction_To(action objc.Selector, target objc.Object) bool
	PerformClick(sender objc.Object)
	InvalidateIntrinsicContentSizeForCell(cell Cell)
	IsEnabled() bool
	SetEnabled(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	IntegerValue() int
	SetIntegerValue(value int)
	ObjectValue() objc.Object
	SetObjectValue(value objc.Object)
	StringValue() string
	SetStringValue(value string)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() Font
	SetFont(value Font)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.Formatter)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	IsHighlighted() bool
	SetHighlighted(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.Object)
	IsContinuous() bool
	SetContinuous(value bool)
	SetTag(value int)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
	Cell() Cell
	SetCell(value Cell)
}

type NSControl struct {
	NSView
}

func MakeControl(ptr unsafe.Pointer) NSControl {
	return NSControl{
		NSView: MakeView(ptr),
	}
}

func (n NSControl) InitWithFrame(frameRect foundation.Rect) NSControl {
	result_ := C.C_NSControl_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeControl(result_)
}

func (n NSControl) InitWithCoder(coder foundation.Coder) NSControl {
	result_ := C.C_NSControl_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeControl(result_)
}

func (n NSControl) Init() NSControl {
	result_ := C.C_NSControl_Init(n.Ptr())
	return MakeControl(result_)
}

func AllocControl() NSControl {
	result_ := C.C_NSControl_AllocControl()
	return MakeControl(result_)
}

func NewControl() NSControl {
	result_ := C.C_NSControl_NewControl()
	return MakeControl(result_)
}

func (n NSControl) Autorelease() NSControl {
	result_ := C.C_NSControl_Autorelease(n.Ptr())
	return MakeControl(result_)
}

func (n NSControl) Retain() NSControl {
	result_ := C.C_NSControl_Retain(n.Ptr())
	return MakeControl(result_)
}

func (n NSControl) TakeDoubleValueFrom(sender objc.Object) {
	C.C_NSControl_TakeDoubleValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) TakeFloatValueFrom(sender objc.Object) {
	C.C_NSControl_TakeFloatValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) TakeIntValueFrom(sender objc.Object) {
	C.C_NSControl_TakeIntValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) TakeIntegerValueFrom(sender objc.Object) {
	C.C_NSControl_TakeIntegerValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) TakeObjectValueFrom(sender objc.Object) {
	C.C_NSControl_TakeObjectValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) TakeStringValueFrom(sender objc.Object) {
	C.C_NSControl_TakeStringValueFrom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view View) {
	C.C_NSControl_DrawWithExpansionFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentFrame))), objc.ExtractPtr(view))
}

func (n NSControl) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	result_ := C.C_NSControl_ExpansionFrameWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentFrame))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSControl) AbortEditing() bool {
	result_ := C.C_NSControl_AbortEditing(n.Ptr())
	return bool(result_)
}

func (n NSControl) CurrentEditor() Text {
	result_ := C.C_NSControl_CurrentEditor(n.Ptr())
	return MakeText(result_)
}

func (n NSControl) ValidateEditing() {
	C.C_NSControl_ValidateEditing(n.Ptr())
}

func (n NSControl) EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj Text, delegate objc.Object, event Event) {
	C.C_NSControl_EditWithFrame_Editor_Delegate_Event(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), objc.ExtractPtr(event))
}

func (n NSControl) EndEditing(textObj Text) {
	C.C_NSControl_EndEditing(n.Ptr(), objc.ExtractPtr(textObj))
}

func (n NSControl) SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj Text, delegate objc.Object, selStart int, selLength int) {
	C.C_NSControl_SelectWithFrame_Editor_Delegate_Start_Length(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), C.int(selStart), C.int(selLength))
}

func (n NSControl) SizeThatFits(size foundation.Size) foundation.Size {
	result_ := C.C_NSControl_SizeThatFits(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSControl) SizeToFit() {
	C.C_NSControl_SizeToFit(n.Ptr())
}

func (n NSControl) SendAction_To(action objc.Selector, target objc.Object) bool {
	result_ := C.C_NSControl_SendAction_To(n.Ptr(), unsafe.Pointer(action), objc.ExtractPtr(target))
	return bool(result_)
}

func (n NSControl) PerformClick(sender objc.Object) {
	C.C_NSControl_PerformClick(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSControl) InvalidateIntrinsicContentSizeForCell(cell Cell) {
	C.C_NSControl_InvalidateIntrinsicContentSizeForCell(n.Ptr(), objc.ExtractPtr(cell))
}

func (n NSControl) IsEnabled() bool {
	result_ := C.C_NSControl_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetEnabled(value bool) {
	C.C_NSControl_SetEnabled(n.Ptr(), C.bool(value))
}

func (n NSControl) DoubleValue() float64 {
	result_ := C.C_NSControl_DoubleValue(n.Ptr())
	return float64(result_)
}

func (n NSControl) SetDoubleValue(value float64) {
	C.C_NSControl_SetDoubleValue(n.Ptr(), C.double(value))
}

func (n NSControl) FloatValue() float32 {
	result_ := C.C_NSControl_FloatValue(n.Ptr())
	return float32(result_)
}

func (n NSControl) SetFloatValue(value float32) {
	C.C_NSControl_SetFloatValue(n.Ptr(), C.float(value))
}

func (n NSControl) IntegerValue() int {
	result_ := C.C_NSControl_IntegerValue(n.Ptr())
	return int(result_)
}

func (n NSControl) SetIntegerValue(value int) {
	C.C_NSControl_SetIntegerValue(n.Ptr(), C.int(value))
}

func (n NSControl) ObjectValue() objc.Object {
	result_ := C.C_NSControl_ObjectValue(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSControl) SetObjectValue(value objc.Object) {
	C.C_NSControl_SetObjectValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSControl) StringValue() string {
	result_ := C.C_NSControl_StringValue(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSControl) SetStringValue(value string) {
	C.C_NSControl_SetStringValue(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSControl) AttributedStringValue() foundation.AttributedString {
	result_ := C.C_NSControl_AttributedStringValue(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSControl) SetAttributedStringValue(value foundation.AttributedString) {
	C.C_NSControl_SetAttributedStringValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSControl) Alignment() TextAlignment {
	result_ := C.C_NSControl_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSControl) SetAlignment(value TextAlignment) {
	C.C_NSControl_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSControl) Font() Font {
	result_ := C.C_NSControl_Font(n.Ptr())
	return MakeFont(result_)
}

func (n NSControl) SetFont(value Font) {
	C.C_NSControl_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSControl) LineBreakMode() LineBreakMode {
	result_ := C.C_NSControl_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result_))
}

func (n NSControl) SetLineBreakMode(value LineBreakMode) {
	C.C_NSControl_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSControl) UsesSingleLineMode() bool {
	result_ := C.C_NSControl_UsesSingleLineMode(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetUsesSingleLineMode(value bool) {
	C.C_NSControl_SetUsesSingleLineMode(n.Ptr(), C.bool(value))
}

func (n NSControl) Formatter() foundation.Formatter {
	result_ := C.C_NSControl_Formatter(n.Ptr())
	return foundation.MakeFormatter(result_)
}

func (n NSControl) SetFormatter(value foundation.Formatter) {
	C.C_NSControl_SetFormatter(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSControl) BaseWritingDirection() WritingDirection {
	result_ := C.C_NSControl_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result_))
}

func (n NSControl) SetBaseWritingDirection(value WritingDirection) {
	C.C_NSControl_SetBaseWritingDirection(n.Ptr(), C.int(int(value)))
}

func (n NSControl) AllowsExpansionToolTips() bool {
	result_ := C.C_NSControl_AllowsExpansionToolTips(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetAllowsExpansionToolTips(value bool) {
	C.C_NSControl_SetAllowsExpansionToolTips(n.Ptr(), C.bool(value))
}

func (n NSControl) ControlSize() ControlSize {
	result_ := C.C_NSControl_ControlSize(n.Ptr())
	return ControlSize(uint(result_))
}

func (n NSControl) SetControlSize(value ControlSize) {
	C.C_NSControl_SetControlSize(n.Ptr(), C.uint(uint(value)))
}

func (n NSControl) IsHighlighted() bool {
	result_ := C.C_NSControl_IsHighlighted(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetHighlighted(value bool) {
	C.C_NSControl_SetHighlighted(n.Ptr(), C.bool(value))
}

func (n NSControl) Action() objc.Selector {
	result_ := C.C_NSControl_Action(n.Ptr())
	return objc.Selector(result_)
}

func (n NSControl) SetAction(value objc.Selector) {
	C.C_NSControl_SetAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSControl) Target() objc.Object {
	result_ := C.C_NSControl_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSControl) SetTarget(value objc.Object) {
	C.C_NSControl_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSControl) IsContinuous() bool {
	result_ := C.C_NSControl_IsContinuous(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetContinuous(value bool) {
	C.C_NSControl_SetContinuous(n.Ptr(), C.bool(value))
}

func (n NSControl) SetTag(value int) {
	C.C_NSControl_SetTag(n.Ptr(), C.int(value))
}

func (n NSControl) RefusesFirstResponder() bool {
	result_ := C.C_NSControl_RefusesFirstResponder(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetRefusesFirstResponder(value bool) {
	C.C_NSControl_SetRefusesFirstResponder(n.Ptr(), C.bool(value))
}

func (n NSControl) IgnoresMultiClick() bool {
	result_ := C.C_NSControl_IgnoresMultiClick(n.Ptr())
	return bool(result_)
}

func (n NSControl) SetIgnoresMultiClick(value bool) {
	C.C_NSControl_SetIgnoresMultiClick(n.Ptr(), C.bool(value))
}

func (n NSControl) Cell() Cell {
	result_ := C.C_NSControl_Cell(n.Ptr())
	return MakeCell(result_)
}

func (n NSControl) SetCell(value Cell) {
	C.C_NSControl_SetCell(n.Ptr(), objc.ExtractPtr(value))
}

type ControlTextEditingDelegate struct {
	Control_IsValidObject                                   func(control Control, obj objc.Object) bool
	Control_DidFailToValidatePartialString_ErrorDescription func(control Control, _string string, error string)
	Control_DidFailToFormatString_ErrorDescription          func(control Control, _string string, error string) bool
	Control_TextShouldBeginEditing                          func(control Control, fieldEditor Text) bool
	Control_TextShouldEndEditing                            func(control Control, fieldEditor Text) bool
	Control_TextView_DoCommandBySelector                    func(control Control, textView TextView, commandSelector objc.Selector) bool
	ControlTextDidBeginEditing                              func(obj foundation.Notification)
	ControlTextDidChange                                    func(obj foundation.Notification)
	ControlTextDidEndEditing                                func(obj foundation.Notification)
}

func (delegate *ControlTextEditingDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapControlTextEditingDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export controlTextEditingDelegate_Control_IsValidObject
func controlTextEditingDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func controlTextEditingDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export controlTextEditingDelegate_Control_DidFailToFormatString_ErrorDescription
func controlTextEditingDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextShouldBeginEditing
func controlTextEditingDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextShouldEndEditing
func controlTextEditingDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextView_DoCommandBySelector
func controlTextEditingDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export controlTextEditingDelegate_ControlTextDidBeginEditing
func controlTextEditingDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export controlTextEditingDelegate_ControlTextDidChange
func controlTextEditingDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export controlTextEditingDelegate_ControlTextDidEndEditing
func controlTextEditingDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export ControlTextEditingDelegate_RespondsTo
func ControlTextEditingDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ControlTextEditingDelegate)
	switch selName {
	case "control:isValidObject:":
		return delegate.Control_IsValidObject != nil
	case "control:didFailToValidatePartialString:errorDescription:":
		return delegate.Control_DidFailToValidatePartialString_ErrorDescription != nil
	case "control:didFailToFormatString:errorDescription:":
		return delegate.Control_DidFailToFormatString_ErrorDescription != nil
	case "control:textShouldBeginEditing:":
		return delegate.Control_TextShouldBeginEditing != nil
	case "control:textShouldEndEditing:":
		return delegate.Control_TextShouldEndEditing != nil
	case "control:textView:doCommandBySelector:":
		return delegate.Control_TextView_DoCommandBySelector != nil
	case "controlTextDidBeginEditing:":
		return delegate.ControlTextDidBeginEditing != nil
	case "controlTextDidChange:":
		return delegate.ControlTextDidChange != nil
	case "controlTextDidEndEditing:":
		return delegate.ControlTextDidEndEditing != nil
	default:
		return false
	}
}
