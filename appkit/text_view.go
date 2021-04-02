package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type TextView interface {
	Text
	AllowsUndo() bool
	SetAllowsUndo(allowsUndo bool)
	TextContainer() TextContainer
	SetTextContainer(textContainer TextContainer)
}

var _ TextView = (*NSTextView)(nil)

type NSTextView struct {
	NSText
}

func MakeTextView(ptr unsafe.Pointer) *NSTextView {
	if ptr == nil {
		return nil
	}
	return &NSTextView{
		NSText: *MakeText(ptr),
	}
}

func (t *NSTextView) AllowsUndo() bool {
	return bool(C.TextView_AllowsUndo(t.Ptr()))
}

func (t *NSTextView) SetAllowsUndo(allowsUndo bool) {
	C.TextView_SetAllowsUndo(t.Ptr(), C.bool(allowsUndo))
}

func (t *NSTextView) TextContainer() TextContainer {
	return MakeTextContainer(C.TextView_TextContainer(t.Ptr()))
}

func (t *NSTextView) SetTextContainer(textContainer TextContainer) {
	C.TextView_SetTextContainer(t.Ptr(), toPointer(textContainer))
}

func NewTextView(frame foundation.Rect) TextView {
	return MakeTextView(C.TextView_NewTextView(toNSRect(frame)))
}

func ScrollableTextView() ScrollView {
	return MakeScrollView(C.TextView_ScrollableTextView())
}
