package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextDelegate struct {
	TextDidChange          func(notification foundation.Notification)
	TextShouldBeginEditing func(textObject Text) bool
	TextDidBeginEditing    func(notification foundation.Notification)
	TextShouldEndEditing   func(textObject Text) bool
	TextDidEndEditing      func(notification foundation.Notification)
}

func WrapTextDelegate(delegate *TextDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapTextDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export TextDelegate_TextDidChange
func TextDelegate_TextDidChange(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TextDelegate)
	delegate.TextDidChange(foundation.MakeNotification(notification))
}

//export TextDelegate_TextShouldBeginEditing
func TextDelegate_TextShouldBeginEditing(id int64, textObject unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TextDelegate)
	result := delegate.TextShouldBeginEditing(MakeText(textObject))
	return C.bool(result)
}

//export TextDelegate_TextDidBeginEditing
func TextDelegate_TextDidBeginEditing(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TextDelegate)
	delegate.TextDidBeginEditing(foundation.MakeNotification(notification))
}

//export TextDelegate_TextShouldEndEditing
func TextDelegate_TextShouldEndEditing(id int64, textObject unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TextDelegate)
	result := delegate.TextShouldEndEditing(MakeText(textObject))
	return C.bool(result)
}

//export TextDelegate_TextDidEndEditing
func TextDelegate_TextDidEndEditing(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TextDelegate)
	delegate.TextDidEndEditing(foundation.MakeNotification(notification))
}

//export TextDelegate_RespondsTo
func TextDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*TextDelegate)
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

//export deleteTextDelegate
func deleteTextDelegate(id int64) {
	resources.Delete(id)
}
