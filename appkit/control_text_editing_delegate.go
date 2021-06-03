package appkit

// #include "control_text_editing_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ControlTextEditingDelegate struct {
	Control_IsValidObject                                   func(control Control, obj objc.Object) bool
	Control_DidFailToValidatePartialString_ErrorDescription func(control Control, _string string, error string)
	Control_DidFailToFormatString_ErrorDescription          func(control Control, _string string, error string) bool
	Control_TextShouldBeginEditing                          func(control Control, fieldEditor Text) bool
	Control_TextShouldEndEditing                            func(control Control, fieldEditor Text) bool
	Control_TextView_DoCommandBySelector                    func(control Control, textView TextView, commandSelector objc.Selector) bool
	ControlTextDidBeginEditing                              func(obj foundation.Notification)
	ControlTextDidChange                                    func(obj foundation.Notification)
	ControlTextDidEndEditing                                func(obj foundation.Notification)
}

func WrapControlTextEditingDelegate(delegate *ControlTextEditingDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapControlTextEditingDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export controlTextEditingDelegate_Control_IsValidObject
func controlTextEditingDelegate_Control_IsValidObject(id int64, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func controlTextEditingDelegate_Control_DidFailToValidatePartialString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export controlTextEditingDelegate_Control_DidFailToFormatString_ErrorDescription
func controlTextEditingDelegate_Control_DidFailToFormatString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextShouldBeginEditing
func controlTextEditingDelegate_Control_TextShouldBeginEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextShouldEndEditing
func controlTextEditingDelegate_Control_TextShouldEndEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export controlTextEditingDelegate_Control_TextView_DoCommandBySelector
func controlTextEditingDelegate_Control_TextView_DoCommandBySelector(id int64, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export controlTextEditingDelegate_ControlTextDidBeginEditing
func controlTextEditingDelegate_ControlTextDidBeginEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export controlTextEditingDelegate_ControlTextDidChange
func controlTextEditingDelegate_ControlTextDidChange(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export controlTextEditingDelegate_ControlTextDidEndEditing
func controlTextEditingDelegate_ControlTextDidEndEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export ControlTextEditingDelegate_RespondsTo
func ControlTextEditingDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ControlTextEditingDelegate)
	switch selName {
	case "control:isValidObject:":
		return delegate.Control_IsValidObject != nil
	case "control:didFailToValidatePartialString:errorDescription:":
		return delegate.Control_DidFailToValidatePartialString_ErrorDescription != nil
	case "control:didFailToFormatString:errorDescription:":
		return delegate.Control_DidFailToFormatString_ErrorDescription != nil
	case "control:textShouldBeginEditing:":
		return delegate.Control_TextShouldBeginEditing != nil
	case "control:textShouldEndEditing:":
		return delegate.Control_TextShouldEndEditing != nil
	case "control:textView:doCommandBySelector:":
		return delegate.Control_TextView_DoCommandBySelector != nil
	case "controlTextDidBeginEditing:":
		return delegate.ControlTextDidBeginEditing != nil
	case "controlTextDidChange:":
		return delegate.ControlTextDidChange != nil
	case "controlTextDidEndEditing:":
		return delegate.ControlTextDidEndEditing != nil
	default:
		return false
	}
}

//export deleteControlTextEditingDelegate
func deleteControlTextEditingDelegate(id int64) {
	resources.Delete(id)
}
