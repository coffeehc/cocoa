package appkit

// #include "text_field_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TextFieldDelegate struct {
	TextField_TextView_Candidates_ForSelectedRange          func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	TextField_TextView_CandidatesForSelectedRange           func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.Object
	TextField_TextView_ShouldSelectCandidateAtIndex         func(textField TextField, textView TextView, index uint) bool
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

func (delegate *TextFieldDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTextFieldDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export textFieldDelegate_TextField_TextView_Candidates_ForSelectedRange
func textFieldDelegate_TextField_TextView_Candidates_ForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, candidates C.Array, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	if candidates.len > 0 {
		defer C.free(candidates.data)
	}
	candidatesSlice := unsafe.Slice((*unsafe.Pointer)(candidates.data), int(candidates.len))
	var goCandidates = make([]foundation.TextCheckingResult, len(candidatesSlice))
	for idx, r := range candidatesSlice {
		goCandidates[idx] = foundation.MakeTextCheckingResult(r)
	}
	result := delegate.TextField_TextView_Candidates_ForSelectedRange(MakeTextField(textField), MakeTextView(textView), goCandidates, foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textFieldDelegate_TextField_TextView_CandidatesForSelectedRange
func textFieldDelegate_TextField_TextView_CandidatesForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.TextField_TextView_CandidatesForSelectedRange(MakeTextField(textField), MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textFieldDelegate_TextField_TextView_ShouldSelectCandidateAtIndex
func textFieldDelegate_TextField_TextView_ShouldSelectCandidateAtIndex(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, index C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.TextField_TextView_ShouldSelectCandidateAtIndex(MakeTextField(textField), MakeTextView(textView), uint(index))
	return C.bool(result)
}

//export textFieldDelegate_Control_IsValidObject
func textFieldDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export textFieldDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func textFieldDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export textFieldDelegate_Control_DidFailToFormatString_ErrorDescription
func textFieldDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export textFieldDelegate_Control_TextShouldBeginEditing
func textFieldDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export textFieldDelegate_Control_TextShouldEndEditing
func textFieldDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export textFieldDelegate_Control_TextView_DoCommandBySelector
func textFieldDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export textFieldDelegate_ControlTextDidBeginEditing
func textFieldDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export textFieldDelegate_ControlTextDidChange
func textFieldDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export textFieldDelegate_ControlTextDidEndEditing
func textFieldDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export TextFieldDelegate_RespondsTo
func TextFieldDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TextFieldDelegate)
	switch selName {
	case "textField:textView:candidates:forSelectedRange:":
		return delegate.TextField_TextView_Candidates_ForSelectedRange != nil
	case "textField:textView:candidatesForSelectedRange:":
		return delegate.TextField_TextView_CandidatesForSelectedRange != nil
	case "textField:textView:shouldSelectCandidateAtIndex:":
		return delegate.TextField_TextView_ShouldSelectCandidateAtIndex != nil
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
