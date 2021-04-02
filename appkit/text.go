package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type Text interface {
	View
	String() string
	SetString(string string)
	IsEditable() bool
	SetEditable(editable bool)
	IsSelectable() bool
	SetSelectable(selectable bool)
	IsFieldEditor() bool
	SetFieldEditor(fieldEditor bool)
	IsRichText() bool
	SetRichText(richText bool)
	MinSize() foundation.Size
	SetMinSize(minSize foundation.Size)
	MaxSize() foundation.Size
	SetMaxSize(maxSize foundation.Size)
	IsVerticallyResizable() bool
	SetVerticallyResizable(verticallyResizable bool)
	IsHorizontallyResizable() bool
	SetHorizontallyResizable(horizontallyResizable bool)
}

var _ Text = (*NSText)(nil)

type NSText struct {
	NSView
}

func MakeText(ptr unsafe.Pointer) *NSText {
	if ptr == nil {
		return nil
	}
	return &NSText{
		NSView: *MakeView(ptr),
	}
}

func (t *NSText) String() string {
	return C.GoString(C.Text_String(t.Ptr()))
}

func (t *NSText) SetString(string string) {
	cString := C.CString(string)
	defer C.free(unsafe.Pointer(cString))
	C.Text_SetString(t.Ptr(), cString)
}

func (t *NSText) IsEditable() bool {
	return bool(C.Text_IsEditable(t.Ptr()))
}

func (t *NSText) SetEditable(editable bool) {
	C.Text_SetEditable(t.Ptr(), C.bool(editable))
}

func (t *NSText) IsSelectable() bool {
	return bool(C.Text_IsSelectable(t.Ptr()))
}

func (t *NSText) SetSelectable(selectable bool) {
	C.Text_SetSelectable(t.Ptr(), C.bool(selectable))
}

func (t *NSText) IsFieldEditor() bool {
	return bool(C.Text_IsFieldEditor(t.Ptr()))
}

func (t *NSText) SetFieldEditor(fieldEditor bool) {
	C.Text_SetFieldEditor(t.Ptr(), C.bool(fieldEditor))
}

func (t *NSText) IsRichText() bool {
	return bool(C.Text_IsRichText(t.Ptr()))
}

func (t *NSText) SetRichText(richText bool) {
	C.Text_SetRichText(t.Ptr(), C.bool(richText))
}

func (t *NSText) MinSize() foundation.Size {
	return toSize(C.Text_MinSize(t.Ptr()))
}

func (t *NSText) SetMinSize(minSize foundation.Size) {
	C.Text_SetMinSize(t.Ptr(), toNSSize(minSize))
}

func (t *NSText) MaxSize() foundation.Size {
	return toSize(C.Text_MaxSize(t.Ptr()))
}

func (t *NSText) SetMaxSize(maxSize foundation.Size) {
	C.Text_SetMaxSize(t.Ptr(), toNSSize(maxSize))
}

func (t *NSText) IsVerticallyResizable() bool {
	return bool(C.Text_IsVerticallyResizable(t.Ptr()))
}

func (t *NSText) SetVerticallyResizable(verticallyResizable bool) {
	C.Text_SetVerticallyResizable(t.Ptr(), C.bool(verticallyResizable))
}

func (t *NSText) IsHorizontallyResizable() bool {
	return bool(C.Text_IsHorizontallyResizable(t.Ptr()))
}

func (t *NSText) SetHorizontallyResizable(horizontallyResizable bool) {
	C.Text_SetHorizontallyResizable(t.Ptr(), C.bool(horizontallyResizable))
}
