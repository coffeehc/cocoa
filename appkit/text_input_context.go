package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_input_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextInputContext interface {
	objc.Object
	Activate()
	Deactivate()
	HandleEvent(event Event) bool
	DiscardMarkedText()
	InvalidateCharacterCoordinates()
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
}

type NSTextInputContext struct {
	objc.NSObject
}

func MakeTextInputContext(ptr unsafe.Pointer) *NSTextInputContext {
	if ptr == nil {
		return nil
	}
	return &NSTextInputContext{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTextInputContext() *NSTextInputContext {
	return MakeTextInputContext(C.C_TextInputContext_Alloc())
}

func (n *NSTextInputContext) Activate() {
	C.C_NSTextInputContext_Activate(n.Ptr())
}

func (n *NSTextInputContext) Deactivate() {
	C.C_NSTextInputContext_Deactivate(n.Ptr())
}

func (n *NSTextInputContext) HandleEvent(event Event) bool {
	result := C.C_NSTextInputContext_HandleEvent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSTextInputContext) DiscardMarkedText() {
	C.C_NSTextInputContext_DiscardMarkedText(n.Ptr())
}

func (n *NSTextInputContext) InvalidateCharacterCoordinates() {
	C.C_NSTextInputContext_InvalidateCharacterCoordinates(n.Ptr())
}

func TextInputContextLocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	result := C.C_NSTextInputContext_TextInputContextLocalizedNameForInputSource(foundation.NewString(string(inputSourceIdentifier)).Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSTextInputContext) AcceptsGlyphInfo() bool {
	result := C.C_NSTextInputContext_AcceptsGlyphInfo(n.Ptr())
	return bool(result)
}

func (n *NSTextInputContext) SetAcceptsGlyphInfo(value bool) {
	C.C_NSTextInputContext_SetAcceptsGlyphInfo(n.Ptr(), C.bool(value))
}

func (n *NSTextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	result := C.C_NSTextInputContext_SelectedKeyboardInputSource(n.Ptr())
	return TextInputSourceIdentifier(foundation.MakeString(result).String())
}

func (n *NSTextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	C.C_NSTextInputContext_SetSelectedKeyboardInputSource(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
