package appkit

// #include "text.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type Text interface {
	View
	ToggleRuler(sender objc.Object)
	ReplaceCharactersInRange_WithRTF(_range foundation.Range, rtfData []byte)
	ReplaceCharactersInRange_WithRTFD(_range foundation.Range, rtfdData []byte)
	ReplaceCharactersInRange_WithString(_range foundation.Range, _string string)
	SelectAll(sender objc.Object)
	Copy_(sender objc.Object)
	Cut(sender objc.Object)
	Paste(sender objc.Object)
	CopyFont(sender objc.Object)
	PasteFont(sender objc.Object)
	CopyRuler(sender objc.Object)
	PasteRuler(sender objc.Object)
	Delete(sender objc.Object)
	ChangeFont(sender objc.Object)
	SetFont_Range(font Font, _range foundation.Range)
	AlignCenter(sender objc.Object)
	AlignLeft(sender objc.Object)
	AlignRight(sender objc.Object)
	SetTextColor_Range(color Color, _range foundation.Range)
	Superscript(sender objc.Object)
	Subscript(sender objc.Object)
	Unscript(sender objc.Object)
	Underline(sender objc.Object)
	ReadRTFDFromFile(path string) bool
	WriteRTFDToFile_Atomically(path string, flag bool) bool
	RTFDFromRange(_range foundation.Range) []byte
	RTFFromRange(_range foundation.Range) []byte
	CheckSpelling(sender objc.Object)
	ShowGuessPanel(sender objc.Object)
	SizeToFit()
	ScrollRangeToVisible(_range foundation.Range)
	String() string
	SetString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsEditable() bool
	SetEditable(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	IsFieldEditor() bool
	SetFieldEditor(value bool)
	IsRichText() bool
	SetRichText(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	IsRulerVisible() bool
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	Font() Font
	SetFont(value Font)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	TextColor() Color
	SetTextColor(value Color)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	IsVerticallyResizable() bool
	SetVerticallyResizable(value bool)
	IsHorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
}

type NSText struct {
	NSView
}

func MakeText(ptr unsafe.Pointer) NSText {
	return NSText{
		NSView: MakeView(ptr),
	}
}

func (n NSText) InitWithCoder(coder foundation.Coder) NSText {
	result_ := C.C_NSText_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeText(result_)
}

func (n NSText) InitWithFrame(frameRect foundation.Rect) NSText {
	result_ := C.C_NSText_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeText(result_)
}

func (n NSText) Init() NSText {
	result_ := C.C_NSText_Init(n.Ptr())
	return MakeText(result_)
}

func AllocText() NSText {
	result_ := C.C_NSText_AllocText()
	return MakeText(result_)
}

func NewText() NSText {
	result_ := C.C_NSText_NewText()
	return MakeText(result_)
}

func (n NSText) Autorelease() NSText {
	result_ := C.C_NSText_Autorelease(n.Ptr())
	return MakeText(result_)
}

func (n NSText) Retain() NSText {
	result_ := C.C_NSText_Retain(n.Ptr())
	return MakeText(result_)
}

func (n NSText) ToggleRuler(sender objc.Object) {
	C.C_NSText_ToggleRuler(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) ReplaceCharactersInRange_WithRTF(_range foundation.Range, rtfData []byte) {
	C.C_NSText_ReplaceCharactersInRange_WithRTF(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)), foundation.NewData(rtfData).Ptr())
}

func (n NSText) ReplaceCharactersInRange_WithRTFD(_range foundation.Range, rtfdData []byte) {
	C.C_NSText_ReplaceCharactersInRange_WithRTFD(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)), foundation.NewData(rtfdData).Ptr())
}

func (n NSText) ReplaceCharactersInRange_WithString(_range foundation.Range, _string string) {
	C.C_NSText_ReplaceCharactersInRange_WithString(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)), foundation.NewString(_string).Ptr())
}

func (n NSText) SelectAll(sender objc.Object) {
	C.C_NSText_SelectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Copy_(sender objc.Object) {
	C.C_NSText_Copy_(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Cut(sender objc.Object) {
	C.C_NSText_Cut(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Paste(sender objc.Object) {
	C.C_NSText_Paste(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) CopyFont(sender objc.Object) {
	C.C_NSText_CopyFont(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) PasteFont(sender objc.Object) {
	C.C_NSText_PasteFont(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) CopyRuler(sender objc.Object) {
	C.C_NSText_CopyRuler(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) PasteRuler(sender objc.Object) {
	C.C_NSText_PasteRuler(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Delete(sender objc.Object) {
	C.C_NSText_Delete(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) ChangeFont(sender objc.Object) {
	C.C_NSText_ChangeFont(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) SetFont_Range(font Font, _range foundation.Range) {
	C.C_NSText_SetFont_Range(n.Ptr(), objc.ExtractPtr(font), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSText) AlignCenter(sender objc.Object) {
	C.C_NSText_AlignCenter(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) AlignLeft(sender objc.Object) {
	C.C_NSText_AlignLeft(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) AlignRight(sender objc.Object) {
	C.C_NSText_AlignRight(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) SetTextColor_Range(color Color, _range foundation.Range) {
	C.C_NSText_SetTextColor_Range(n.Ptr(), objc.ExtractPtr(color), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSText) Superscript(sender objc.Object) {
	C.C_NSText_Superscript(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Subscript(sender objc.Object) {
	C.C_NSText_Subscript(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Unscript(sender objc.Object) {
	C.C_NSText_Unscript(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) Underline(sender objc.Object) {
	C.C_NSText_Underline(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) ReadRTFDFromFile(path string) bool {
	result_ := C.C_NSText_ReadRTFDFromFile(n.Ptr(), foundation.NewString(path).Ptr())
	return bool(result_)
}

func (n NSText) WriteRTFDToFile_Atomically(path string, flag bool) bool {
	result_ := C.C_NSText_WriteRTFDToFile_Atomically(n.Ptr(), foundation.NewString(path).Ptr(), C.bool(flag))
	return bool(result_)
}

func (n NSText) RTFDFromRange(_range foundation.Range) []byte {
	result_ := C.C_NSText_RTFDFromRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSText) RTFFromRange(_range foundation.Range) []byte {
	result_ := C.C_NSText_RTFFromRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSText) CheckSpelling(sender objc.Object) {
	C.C_NSText_CheckSpelling(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) ShowGuessPanel(sender objc.Object) {
	C.C_NSText_ShowGuessPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSText) SizeToFit() {
	C.C_NSText_SizeToFit(n.Ptr())
}

func (n NSText) ScrollRangeToVisible(_range foundation.Range) {
	C.C_NSText_ScrollRangeToVisible(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSText) String() string {
	result_ := C.C_NSText_String(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSText) SetString(value string) {
	C.C_NSText_SetString(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSText) BackgroundColor() Color {
	result_ := C.C_NSText_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSText) SetBackgroundColor(value Color) {
	C.C_NSText_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSText) DrawsBackground() bool {
	result_ := C.C_NSText_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n NSText) SetDrawsBackground(value bool) {
	C.C_NSText_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n NSText) IsEditable() bool {
	result_ := C.C_NSText_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSText) SetEditable(value bool) {
	C.C_NSText_SetEditable(n.Ptr(), C.bool(value))
}

func (n NSText) IsSelectable() bool {
	result_ := C.C_NSText_IsSelectable(n.Ptr())
	return bool(result_)
}

func (n NSText) SetSelectable(value bool) {
	C.C_NSText_SetSelectable(n.Ptr(), C.bool(value))
}

func (n NSText) IsFieldEditor() bool {
	result_ := C.C_NSText_IsFieldEditor(n.Ptr())
	return bool(result_)
}

func (n NSText) SetFieldEditor(value bool) {
	C.C_NSText_SetFieldEditor(n.Ptr(), C.bool(value))
}

func (n NSText) IsRichText() bool {
	result_ := C.C_NSText_IsRichText(n.Ptr())
	return bool(result_)
}

func (n NSText) SetRichText(value bool) {
	C.C_NSText_SetRichText(n.Ptr(), C.bool(value))
}

func (n NSText) ImportsGraphics() bool {
	result_ := C.C_NSText_ImportsGraphics(n.Ptr())
	return bool(result_)
}

func (n NSText) SetImportsGraphics(value bool) {
	C.C_NSText_SetImportsGraphics(n.Ptr(), C.bool(value))
}

func (n NSText) UsesFontPanel() bool {
	result_ := C.C_NSText_UsesFontPanel(n.Ptr())
	return bool(result_)
}

func (n NSText) SetUsesFontPanel(value bool) {
	C.C_NSText_SetUsesFontPanel(n.Ptr(), C.bool(value))
}

func (n NSText) IsRulerVisible() bool {
	result_ := C.C_NSText_IsRulerVisible(n.Ptr())
	return bool(result_)
}

func (n NSText) SelectedRange() foundation.Range {
	result_ := C.C_NSText_SelectedRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSText) SetSelectedRange(value foundation.Range) {
	C.C_NSText_SetSelectedRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(value)))
}

func (n NSText) Font() Font {
	result_ := C.C_NSText_Font(n.Ptr())
	return MakeFont(result_)
}

func (n NSText) SetFont(value Font) {
	C.C_NSText_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSText) Alignment() TextAlignment {
	result_ := C.C_NSText_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSText) SetAlignment(value TextAlignment) {
	C.C_NSText_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSText) TextColor() Color {
	result_ := C.C_NSText_TextColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSText) SetTextColor(value Color) {
	C.C_NSText_SetTextColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSText) BaseWritingDirection() WritingDirection {
	result_ := C.C_NSText_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result_))
}

func (n NSText) SetBaseWritingDirection(value WritingDirection) {
	C.C_NSText_SetBaseWritingDirection(n.Ptr(), C.int(int(value)))
}

func (n NSText) MaxSize() foundation.Size {
	result_ := C.C_NSText_MaxSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSText) SetMaxSize(value foundation.Size) {
	C.C_NSText_SetMaxSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSText) MinSize() foundation.Size {
	result_ := C.C_NSText_MinSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSText) SetMinSize(value foundation.Size) {
	C.C_NSText_SetMinSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSText) IsVerticallyResizable() bool {
	result_ := C.C_NSText_IsVerticallyResizable(n.Ptr())
	return bool(result_)
}

func (n NSText) SetVerticallyResizable(value bool) {
	C.C_NSText_SetVerticallyResizable(n.Ptr(), C.bool(value))
}

func (n NSText) IsHorizontallyResizable() bool {
	result_ := C.C_NSText_IsHorizontallyResizable(n.Ptr())
	return bool(result_)
}

func (n NSText) SetHorizontallyResizable(value bool) {
	C.C_NSText_SetHorizontallyResizable(n.Ptr(), C.bool(value))
}

func (n NSText) Delegate() objc.Object {
	result_ := C.C_NSText_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSText) SetDelegate(value objc.Object) {
	C.C_NSText_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

type TextDelegate struct {
	TextDidChange          func(notification foundation.Notification)
	TextShouldBeginEditing func(textObject Text) bool
	TextDidBeginEditing    func(notification foundation.Notification)
	TextShouldEndEditing   func(textObject Text) bool
	TextDidEndEditing      func(notification foundation.Notification)
}

func (delegate *TextDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTextDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export textDelegate_TextDidChange
func textDelegate_TextDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	delegate.TextDidChange(foundation.MakeNotification(notification))
}

//export textDelegate_TextShouldBeginEditing
func textDelegate_TextShouldBeginEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	result := delegate.TextShouldBeginEditing(MakeText(textObject))
	return C.bool(result)
}

//export textDelegate_TextDidBeginEditing
func textDelegate_TextDidBeginEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	delegate.TextDidBeginEditing(foundation.MakeNotification(notification))
}

//export textDelegate_TextShouldEndEditing
func textDelegate_TextShouldEndEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	result := delegate.TextShouldEndEditing(MakeText(textObject))
	return C.bool(result)
}

//export textDelegate_TextDidEndEditing
func textDelegate_TextDidEndEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	delegate.TextDidEndEditing(foundation.MakeNotification(notification))
}

//export TextDelegate_RespondsTo
func TextDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TextDelegate)
	switch selName {
	case "textDidChange:":
		return delegate.TextDidChange != nil
	case "textShouldBeginEditing:":
		return delegate.TextShouldBeginEditing != nil
	case "textDidBeginEditing:":
		return delegate.TextDidBeginEditing != nil
	case "textShouldEndEditing:":
		return delegate.TextShouldEndEditing != nil
	case "textDidEndEditing:":
		return delegate.TextDidEndEditing != nil
	default:
		return false
	}
}
