package appkit

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
	Client() objc.Object
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	KeyboardInputSources() []TextInputSourceIdentifier
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
}

type NSTextInputContext struct {
	objc.NSObject
}

func MakeTextInputContext(ptr unsafe.Pointer) NSTextInputContext {
	return NSTextInputContext{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTextInputContext() NSTextInputContext {
	return MakeTextInputContext(C.C_TextInputContext_Alloc())
}

func (n NSTextInputContext) InitWithClient(client objc.Object) TextInputContext {
	result_ := C.C_NSTextInputContext_InitWithClient(n.Ptr(), objc.ExtractPtr(client))
	return MakeTextInputContext(result_)
}

func (n NSTextInputContext) Activate() {
	C.C_NSTextInputContext_Activate(n.Ptr())
}

func (n NSTextInputContext) Deactivate() {
	C.C_NSTextInputContext_Deactivate(n.Ptr())
}

func (n NSTextInputContext) HandleEvent(event Event) bool {
	result_ := C.C_NSTextInputContext_HandleEvent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSTextInputContext) DiscardMarkedText() {
	C.C_NSTextInputContext_DiscardMarkedText(n.Ptr())
}

func (n NSTextInputContext) InvalidateCharacterCoordinates() {
	C.C_NSTextInputContext_InvalidateCharacterCoordinates(n.Ptr())
}

func TextInputContext_LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	result_ := C.C_NSTextInputContext_TextInputContext_LocalizedNameForInputSource(foundation.NewString(string(inputSourceIdentifier)).Ptr())
	return foundation.MakeString(result_).String()
}

func TextInputContext_CurrentInputContext() TextInputContext {
	result_ := C.C_NSTextInputContext_TextInputContext_CurrentInputContext()
	return MakeTextInputContext(result_)
}

func (n NSTextInputContext) Client() objc.Object {
	result_ := C.C_NSTextInputContext_Client(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTextInputContext) AcceptsGlyphInfo() bool {
	result_ := C.C_NSTextInputContext_AcceptsGlyphInfo(n.Ptr())
	return bool(result_)
}

func (n NSTextInputContext) SetAcceptsGlyphInfo(value bool) {
	C.C_NSTextInputContext_SetAcceptsGlyphInfo(n.Ptr(), C.bool(value))
}

func (n NSTextInputContext) AllowedInputSourceLocales() []string {
	result_ := C.C_NSTextInputContext_AllowedInputSourceLocales(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSTextInputContext) SetAllowedInputSourceLocales(value []string) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = foundation.NewString(v).Ptr()
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextInputContext_SetAllowedInputSourceLocales(n.Ptr(), cValue)
}

func (n NSTextInputContext) KeyboardInputSources() []TextInputSourceIdentifier {
	result_ := C.C_NSTextInputContext_KeyboardInputSources(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextInputSourceIdentifier, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = TextInputSourceIdentifier(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	result_ := C.C_NSTextInputContext_SelectedKeyboardInputSource(n.Ptr())
	return TextInputSourceIdentifier(foundation.MakeString(result_).String())
}

func (n NSTextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	C.C_NSTextInputContext_SetSelectedKeyboardInputSource(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
