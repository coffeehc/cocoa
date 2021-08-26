package appkit

// #include "text_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

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
