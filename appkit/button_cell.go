package appkit

// #include "button_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ButtonCell interface {
	ActionCell
	SetPeriodicDelay_Interval(delay float32, interval float32)
	SetButtonType(_type ButtonType)
	MouseEntered(event Event)
	MouseExited(event Event)
	DrawBezelWithFrame_InView(frame foundation.Rect, controlView View)
	DrawImage_WithFrame_InView(image Image, frame foundation.Rect, controlView View)
	DrawTitle_WithFrame_InView(title foundation.AttributedString, frame foundation.Rect, controlView View) foundation.Rect
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.AttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	AlternateImage() Image
	SetAlternateImage(value Image)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	IsTransparent() bool
	SetTransparent(value bool)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	HighlightsBy() CellStyleMask
	SetHighlightsBy(value CellStyleMask)
	ShowsStateBy() CellStyleMask
	SetShowsStateBy(value CellStyleMask)
	Sound() Sound
	SetSound(value Sound)
}

type NSButtonCell struct {
	NSActionCell
}

func MakeButtonCell(ptr unsafe.Pointer) *NSButtonCell {
	if ptr == nil {
		return nil
	}
	return &NSButtonCell{
		NSActionCell: *MakeActionCell(ptr),
	}
}

func AllocButtonCell() *NSButtonCell {
	return MakeButtonCell(C.C_ButtonCell_Alloc())
}

func (n *NSButtonCell) InitWithCoder(coder foundation.Coder) ButtonCell {
	result_ := C.C_NSButtonCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeButtonCell(result_)
}

func (n *NSButtonCell) InitImageCell(image Image) ButtonCell {
	result_ := C.C_NSButtonCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakeButtonCell(result_)
}

func (n *NSButtonCell) InitTextCell(_string string) ButtonCell {
	result_ := C.C_NSButtonCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeButtonCell(result_)
}

func (n *NSButtonCell) Init() ButtonCell {
	result_ := C.C_NSButtonCell_Init(n.Ptr())
	return MakeButtonCell(result_)
}

func (n *NSButtonCell) SetPeriodicDelay_Interval(delay float32, interval float32) {
	C.C_NSButtonCell_SetPeriodicDelay_Interval(n.Ptr(), C.float(delay), C.float(interval))
}

func (n *NSButtonCell) SetButtonType(_type ButtonType) {
	C.C_NSButtonCell_SetButtonType(n.Ptr(), C.uint(uint(_type)))
}

func (n *NSButtonCell) MouseEntered(event Event) {
	C.C_NSButtonCell_MouseEntered(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSButtonCell) MouseExited(event Event) {
	C.C_NSButtonCell_MouseExited(n.Ptr(), objc.ExtractPtr(event))
}

func (n *NSButtonCell) DrawBezelWithFrame_InView(frame foundation.Rect, controlView View) {
	C.C_NSButtonCell_DrawBezelWithFrame_InView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(controlView))
}

func (n *NSButtonCell) DrawImage_WithFrame_InView(image Image, frame foundation.Rect, controlView View) {
	C.C_NSButtonCell_DrawImage_WithFrame_InView(n.Ptr(), objc.ExtractPtr(image), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(controlView))
}

func (n *NSButtonCell) DrawTitle_WithFrame_InView(title foundation.AttributedString, frame foundation.Rect, controlView View) foundation.Rect {
	result_ := C.C_NSButtonCell_DrawTitle_WithFrame_InView(n.Ptr(), objc.ExtractPtr(title), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(controlView))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSButtonCell) AlternateTitle() string {
	result_ := C.C_NSButtonCell_AlternateTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSButtonCell) SetAlternateTitle(value string) {
	C.C_NSButtonCell_SetAlternateTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	result_ := C.C_NSButtonCell_AttributedAlternateTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSButtonCell) SetAttributedAlternateTitle(value foundation.AttributedString) {
	C.C_NSButtonCell_SetAttributedAlternateTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSButtonCell) AttributedTitle() foundation.AttributedString {
	result_ := C.C_NSButtonCell_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSButtonCell) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSButtonCell_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSButtonCell) AlternateImage() Image {
	result_ := C.C_NSButtonCell_AlternateImage(n.Ptr())
	return MakeImage(result_)
}

func (n *NSButtonCell) SetAlternateImage(value Image) {
	C.C_NSButtonCell_SetAlternateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSButtonCell) ImagePosition() CellImagePosition {
	result_ := C.C_NSButtonCell_ImagePosition(n.Ptr())
	return CellImagePosition(uint(result_))
}

func (n *NSButtonCell) SetImagePosition(value CellImagePosition) {
	C.C_NSButtonCell_SetImagePosition(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) ImageScaling() ImageScaling {
	result_ := C.C_NSButtonCell_ImageScaling(n.Ptr())
	return ImageScaling(uint(result_))
}

func (n *NSButtonCell) SetImageScaling(value ImageScaling) {
	C.C_NSButtonCell_SetImageScaling(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) SetKeyEquivalent(value string) {
	C.C_NSButtonCell_SetKeyEquivalent(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	result_ := C.C_NSButtonCell_KeyEquivalentModifierMask(n.Ptr())
	return EventModifierFlags(uint(result_))
}

func (n *NSButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	C.C_NSButtonCell_SetKeyEquivalentModifierMask(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) BackgroundColor() Color {
	result_ := C.C_NSButtonCell_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSButtonCell) SetBackgroundColor(value Color) {
	C.C_NSButtonCell_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSButtonCell) BezelStyle() BezelStyle {
	result_ := C.C_NSButtonCell_BezelStyle(n.Ptr())
	return BezelStyle(uint(result_))
}

func (n *NSButtonCell) SetBezelStyle(value BezelStyle) {
	C.C_NSButtonCell_SetBezelStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) ImageDimsWhenDisabled() bool {
	result_ := C.C_NSButtonCell_ImageDimsWhenDisabled(n.Ptr())
	return bool(result_)
}

func (n *NSButtonCell) SetImageDimsWhenDisabled(value bool) {
	C.C_NSButtonCell_SetImageDimsWhenDisabled(n.Ptr(), C.bool(value))
}

func (n *NSButtonCell) IsTransparent() bool {
	result_ := C.C_NSButtonCell_IsTransparent(n.Ptr())
	return bool(result_)
}

func (n *NSButtonCell) SetTransparent(value bool) {
	C.C_NSButtonCell_SetTransparent(n.Ptr(), C.bool(value))
}

func (n *NSButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	result_ := C.C_NSButtonCell_ShowsBorderOnlyWhileMouseInside(n.Ptr())
	return bool(result_)
}

func (n *NSButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	C.C_NSButtonCell_SetShowsBorderOnlyWhileMouseInside(n.Ptr(), C.bool(value))
}

func (n *NSButtonCell) HighlightsBy() CellStyleMask {
	result_ := C.C_NSButtonCell_HighlightsBy(n.Ptr())
	return CellStyleMask(uint(result_))
}

func (n *NSButtonCell) SetHighlightsBy(value CellStyleMask) {
	C.C_NSButtonCell_SetHighlightsBy(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) ShowsStateBy() CellStyleMask {
	result_ := C.C_NSButtonCell_ShowsStateBy(n.Ptr())
	return CellStyleMask(uint(result_))
}

func (n *NSButtonCell) SetShowsStateBy(value CellStyleMask) {
	C.C_NSButtonCell_SetShowsStateBy(n.Ptr(), C.uint(uint(value)))
}

func (n *NSButtonCell) Sound() Sound {
	result_ := C.C_NSButtonCell_Sound(n.Ptr())
	return MakeSound(result_)
}

func (n *NSButtonCell) SetSound(value Sound) {
	C.C_NSButtonCell_SetSound(n.Ptr(), objc.ExtractPtr(value))
}
