package appkit

// #include "text_field_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextFieldCell interface {
	ActionCell
	SetWantsNotificationForMarkedText(flag bool)
	TextColor() Color
	SetTextColor(value Color)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
}

type NSTextFieldCell struct {
	NSActionCell
}

func MakeTextFieldCell(ptr unsafe.Pointer) NSTextFieldCell {
	return NSTextFieldCell{
		NSActionCell: MakeActionCell(ptr),
	}
}

func AllocTextFieldCell() NSTextFieldCell {
	return MakeTextFieldCell(C.C_TextFieldCell_Alloc())
}

func (n NSTextFieldCell) InitTextCell(_string string) TextFieldCell {
	result_ := C.C_NSTextFieldCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeTextFieldCell(result_)
}

func (n NSTextFieldCell) InitWithCoder(coder foundation.Coder) TextFieldCell {
	result_ := C.C_NSTextFieldCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextFieldCell(result_)
}

func (n NSTextFieldCell) Init() TextFieldCell {
	result_ := C.C_NSTextFieldCell_Init(n.Ptr())
	return MakeTextFieldCell(result_)
}

func (n NSTextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	C.C_NSTextFieldCell_SetWantsNotificationForMarkedText(n.Ptr(), C.bool(flag))
}

func (n NSTextFieldCell) TextColor() Color {
	result_ := C.C_NSTextFieldCell_TextColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextFieldCell) SetTextColor(value Color) {
	C.C_NSTextFieldCell_SetTextColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextFieldCell) BezelStyle() TextFieldBezelStyle {
	result_ := C.C_NSTextFieldCell_BezelStyle(n.Ptr())
	return TextFieldBezelStyle(uint(result_))
}

func (n NSTextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	C.C_NSTextFieldCell_SetBezelStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSTextFieldCell) BackgroundColor() Color {
	result_ := C.C_NSTextFieldCell_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextFieldCell) SetBackgroundColor(value Color) {
	C.C_NSTextFieldCell_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextFieldCell) DrawsBackground() bool {
	result_ := C.C_NSTextFieldCell_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n NSTextFieldCell) SetDrawsBackground(value bool) {
	C.C_NSTextFieldCell_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n NSTextFieldCell) PlaceholderString() string {
	result_ := C.C_NSTextFieldCell_PlaceholderString(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTextFieldCell) SetPlaceholderString(value string) {
	C.C_NSTextFieldCell_SetPlaceholderString(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	result_ := C.C_NSTextFieldCell_PlaceholderAttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSTextFieldCell) SetPlaceholderAttributedString(value foundation.AttributedString) {
	C.C_NSTextFieldCell_SetPlaceholderAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextFieldCell) AllowedInputSourceLocales() []string {
	result_ := C.C_NSTextFieldCell_AllowedInputSourceLocales(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSTextFieldCell) SetAllowedInputSourceLocales(value []string) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = foundation.NewString(v).Ptr()
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextFieldCell_SetAllowedInputSourceLocales(n.Ptr(), cValue)
}
