package appkit

// #include "text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func AllocTextField() NSTextField {
	return MakeTextField(C.C_TextField_Alloc())
}

func (n NSTextField) InitWithFrame(frameRect foundation.Rect) TextField {
	result_ := C.C_NSTextField_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTextField(result_)
}

func (n NSTextField) InitWithCoder(coder foundation.Coder) TextField {
	result_ := C.C_NSTextField_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextField(result_)
}

func (n NSTextField) Init() TextField {
	result_ := C.C_NSTextField_Init(n.Ptr())
	return MakeTextField(result_)
}

func TextField_LabelWithAttributedString(attributedStringValue foundation.AttributedString) TextField {
	result_ := C.C_NSTextField_TextField_LabelWithAttributedString(objc.ExtractPtr(attributedStringValue))
	return MakeTextField(result_)
}

func TextField_LabelWithString(stringValue string) TextField {
	result_ := C.C_NSTextField_TextField_LabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeTextField(result_)
}

func TextFieldWithString(stringValue string) TextField {
	result_ := C.C_NSTextField_TextFieldWithString(foundation.NewString(stringValue).Ptr())
	return MakeTextField(result_)
}

func TextField_WrappingLabelWithString(stringValue string) TextField {
	result_ := C.C_NSTextField_TextField_WrappingLabelWithString(foundation.NewString(stringValue).Ptr())
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
