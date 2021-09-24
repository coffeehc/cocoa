package appkit

// #include "text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TextField interface {
	Control
	SelectText(sender objc.Object)
	TextShouldBeginEditing(textObject Text) bool
	TextDidBeginEditing(notification foundation.Notification)
	TextDidChange(notification foundation.Notification)
	TextShouldEndEditing(textObject Text) bool
	TextDidEndEditing(notification foundation.Notification)
	IsSelectable() bool
	SetSelectable(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	PreferredMaxLayoutWidth() coregraphics.Float
	SetPreferredMaxLayoutWidth(value coregraphics.Float)
	TextColor() Color
	SetTextColor(value Color)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	IsBordered() bool
	SetBordered(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
}

type NSTextField struct {
	NSControl
}

func MakeTextField(ptr unsafe.Pointer) NSTextField {
	return NSTextField{
		NSControl: MakeControl(ptr),
	}
}

func TextField_LabelWithAttributedString(attributedStringValue foundation.AttributedString) NSTextField {
	result_ := C.C_NSTextField_TextField_LabelWithAttributedString(objc.ExtractPtr(attributedStringValue))
	return MakeTextField(result_)
}

func TextField_LabelWithString(stringValue string) NSTextField {
	result_ := C.C_NSTextField_TextField_LabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeTextField(result_)
}

func TextFieldWithString(stringValue string) NSTextField {
	result_ := C.C_NSTextField_TextFieldWithString(foundation.NewString(stringValue).Ptr())
	return MakeTextField(result_)
}

func TextField_WrappingLabelWithString(stringValue string) NSTextField {
	result_ := C.C_NSTextField_TextField_WrappingLabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeTextField(result_)
}

func (n NSTextField) InitWithFrame(frameRect foundation.Rect) NSTextField {
	result_ := C.C_NSTextField_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTextField(result_)
}

func (n NSTextField) InitWithCoder(coder foundation.Coder) NSTextField {
	result_ := C.C_NSTextField_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextField(result_)
}

func (n NSTextField) Init() NSTextField {
	result_ := C.C_NSTextField_Init(n.Ptr())
	return MakeTextField(result_)
}

func AllocTextField() NSTextField {
	result_ := C.C_NSTextField_AllocTextField()
	return MakeTextField(result_)
}

func NewTextField() NSTextField {
	result_ := C.C_NSTextField_NewTextField()
	return MakeTextField(result_)
}

func (n NSTextField) Autorelease() NSTextField {
	result_ := C.C_NSTextField_Autorelease(n.Ptr())
	return MakeTextField(result_)
}

func (n NSTextField) Retain() NSTextField {
	result_ := C.C_NSTextField_Retain(n.Ptr())
	return MakeTextField(result_)
}

func (n NSTextField) SelectText(sender objc.Object) {
	C.C_NSTextField_SelectText(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextField) TextShouldBeginEditing(textObject Text) bool {
	result_ := C.C_NSTextField_TextShouldBeginEditing(n.Ptr(), objc.ExtractPtr(textObject))
	return bool(result_)
}

func (n NSTextField) TextDidBeginEditing(notification foundation.Notification) {
	C.C_NSTextField_TextDidBeginEditing(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSTextField) TextDidChange(notification foundation.Notification) {
	C.C_NSTextField_TextDidChange(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSTextField) TextShouldEndEditing(textObject Text) bool {
	result_ := C.C_NSTextField_TextShouldEndEditing(n.Ptr(), objc.ExtractPtr(textObject))
	return bool(result_)
}

func (n NSTextField) TextDidEndEditing(notification foundation.Notification) {
	C.C_NSTextField_TextDidEndEditing(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSTextField) IsSelectable() bool {
	result_ := C.C_NSTextField_IsSelectable(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetSelectable(value bool) {
	C.C_NSTextField_SetSelectable(n.Ptr(), C.bool(value))
}

func (n NSTextField) IsEditable() bool {
	result_ := C.C_NSTextField_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetEditable(value bool) {
	C.C_NSTextField_SetEditable(n.Ptr(), C.bool(value))
}

func (n NSTextField) AllowsEditingTextAttributes() bool {
	result_ := C.C_NSTextField_AllowsEditingTextAttributes(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetAllowsEditingTextAttributes(value bool) {
	C.C_NSTextField_SetAllowsEditingTextAttributes(n.Ptr(), C.bool(value))
}

func (n NSTextField) ImportsGraphics() bool {
	result_ := C.C_NSTextField_ImportsGraphics(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetImportsGraphics(value bool) {
	C.C_NSTextField_SetImportsGraphics(n.Ptr(), C.bool(value))
}

func (n NSTextField) PlaceholderString() string {
	result_ := C.C_NSTextField_PlaceholderString(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTextField) SetPlaceholderString(value string) {
	C.C_NSTextField_SetPlaceholderString(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTextField) PlaceholderAttributedString() foundation.AttributedString {
	result_ := C.C_NSTextField_PlaceholderAttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSTextField) SetPlaceholderAttributedString(value foundation.AttributedString) {
	C.C_NSTextField_SetPlaceholderAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextField) AllowsDefaultTighteningForTruncation() bool {
	result_ := C.C_NSTextField_AllowsDefaultTighteningForTruncation(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	C.C_NSTextField_SetAllowsDefaultTighteningForTruncation(n.Ptr(), C.bool(value))
}

func (n NSTextField) MaximumNumberOfLines() int {
	result_ := C.C_NSTextField_MaximumNumberOfLines(n.Ptr())
	return int(result_)
}

func (n NSTextField) SetMaximumNumberOfLines(value int) {
	C.C_NSTextField_SetMaximumNumberOfLines(n.Ptr(), C.int(value))
}

func (n NSTextField) PreferredMaxLayoutWidth() coregraphics.Float {
	result_ := C.C_NSTextField_PreferredMaxLayoutWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTextField) SetPreferredMaxLayoutWidth(value coregraphics.Float) {
	C.C_NSTextField_SetPreferredMaxLayoutWidth(n.Ptr(), C.double(float64(value)))
}

func (n NSTextField) TextColor() Color {
	result_ := C.C_NSTextField_TextColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextField) SetTextColor(value Color) {
	C.C_NSTextField_SetTextColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextField) BackgroundColor() Color {
	result_ := C.C_NSTextField_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextField) SetBackgroundColor(value Color) {
	C.C_NSTextField_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextField) DrawsBackground() bool {
	result_ := C.C_NSTextField_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetDrawsBackground(value bool) {
	C.C_NSTextField_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n NSTextField) IsBezeled() bool {
	result_ := C.C_NSTextField_IsBezeled(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetBezeled(value bool) {
	C.C_NSTextField_SetBezeled(n.Ptr(), C.bool(value))
}

func (n NSTextField) BezelStyle() TextFieldBezelStyle {
	result_ := C.C_NSTextField_BezelStyle(n.Ptr())
	return TextFieldBezelStyle(uint(result_))
}

func (n NSTextField) SetBezelStyle(value TextFieldBezelStyle) {
	C.C_NSTextField_SetBezelStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSTextField) IsBordered() bool {
	result_ := C.C_NSTextField_IsBordered(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetBordered(value bool) {
	C.C_NSTextField_SetBordered(n.Ptr(), C.bool(value))
}

func (n NSTextField) AllowsCharacterPickerTouchBarItem() bool {
	result_ := C.C_NSTextField_AllowsCharacterPickerTouchBarItem(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	C.C_NSTextField_SetAllowsCharacterPickerTouchBarItem(n.Ptr(), C.bool(value))
}

func (n NSTextField) IsAutomaticTextCompletionEnabled() bool {
	result_ := C.C_NSTextField_IsAutomaticTextCompletionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextField) SetAutomaticTextCompletionEnabled(value bool) {
	C.C_NSTextField_SetAutomaticTextCompletionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextField) Delegate() objc.Object {
	result_ := C.C_NSTextField_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTextField) SetDelegate(value objc.Object) {
	C.C_NSTextField_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextField) LineBreakStrategy() LineBreakStrategy {
	result_ := C.C_NSTextField_LineBreakStrategy(n.Ptr())
	return LineBreakStrategy(uint(result_))
}

func (n NSTextField) SetLineBreakStrategy(value LineBreakStrategy) {
	C.C_NSTextField_SetLineBreakStrategy(n.Ptr(), C.uint(uint(value)))
}

type TextFieldDelegate struct {
	TextField_TextView_Candidates_ForSelectedRange          func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	TextField_TextView_CandidatesForSelectedRange           func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object
	TextField_TextView_ShouldSelectCandidateAtIndex         func(textField TextField, textView TextView, index uint) bool
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

func (delegate *TextFieldDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTextFieldDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export textFieldDelegate_TextField_TextView_Candidates_ForSelectedRange
func textFieldDelegate_TextField_TextView_Candidates_ForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, candidates C.Array, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	if candidates.len > 0 {
		defer C.free(candidates.data)
	}
	candidatesSlice := unsafe.Slice((*unsafe.Pointer)(candidates.data), int(candidates.len))
	var goCandidates = make([]foundation.TextCheckingResult, len(candidatesSlice))
	for idx, r := range candidatesSlice {
		goCandidates[idx] = foundation.MakeTextCheckingResult(r)
	}
	result := delegate.TextField_TextView_Candidates_ForSelectedRange(MakeTextField(textField), MakeTextView(textView), goCandidates, foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textFieldDelegate_TextField_TextView_CandidatesForSelectedRange
func textFieldDelegate_TextField_TextView_CandidatesForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.TextField_TextView_CandidatesForSelectedRange(MakeTextField(textField), MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textFieldDelegate_TextField_TextView_ShouldSelectCandidateAtIndex
func textFieldDelegate_TextField_TextView_ShouldSelectCandidateAtIndex(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, index C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.TextField_TextView_ShouldSelectCandidateAtIndex(MakeTextField(textField), MakeTextView(textView), uint(index))
	return C.bool(result)
}

//export textFieldDelegate_Control_IsValidObject
func textFieldDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export textFieldDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func textFieldDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export textFieldDelegate_Control_DidFailToFormatString_ErrorDescription
func textFieldDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export textFieldDelegate_Control_TextShouldBeginEditing
func textFieldDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export textFieldDelegate_Control_TextShouldEndEditing
func textFieldDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export textFieldDelegate_Control_TextView_DoCommandBySelector
func textFieldDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export textFieldDelegate_ControlTextDidBeginEditing
func textFieldDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export textFieldDelegate_ControlTextDidChange
func textFieldDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export textFieldDelegate_ControlTextDidEndEditing
func textFieldDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export TextFieldDelegate_RespondsTo
func TextFieldDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	switch selName {
	case "textField:textView:candidates:forSelectedRange:":
		return delegate.TextField_TextView_Candidates_ForSelectedRange != nil
	case "textField:textView:candidatesForSelectedRange:":
		return delegate.TextField_TextView_CandidatesForSelectedRange != nil
	case "textField:textView:shouldSelectCandidateAtIndex:":
		return delegate.TextField_TextView_ShouldSelectCandidateAtIndex != nil
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
