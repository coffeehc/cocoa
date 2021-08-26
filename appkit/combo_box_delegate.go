package appkit

// #include "combo_box_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ComboBoxDelegate struct {
	ComboBoxSelectionDidChange                              func(notification foundation.Notification)
	ComboBoxSelectionIsChanging                             func(notification foundation.Notification)
	ComboBoxWillDismiss                                     func(notification foundation.Notification)
	ComboBoxWillPopUp                                       func(notification foundation.Notification)
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

func (delegate *ComboBoxDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapComboBoxDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export comboBoxDelegate_ComboBoxSelectionDidChange
func comboBoxDelegate_ComboBoxSelectionDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ComboBoxSelectionDidChange(foundation.MakeNotification(notification))
}

//export comboBoxDelegate_ComboBoxSelectionIsChanging
func comboBoxDelegate_ComboBoxSelectionIsChanging(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ComboBoxSelectionIsChanging(foundation.MakeNotification(notification))
}

//export comboBoxDelegate_ComboBoxWillDismiss
func comboBoxDelegate_ComboBoxWillDismiss(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ComboBoxWillDismiss(foundation.MakeNotification(notification))
}

//export comboBoxDelegate_ComboBoxWillPopUp
func comboBoxDelegate_ComboBoxWillPopUp(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ComboBoxWillPopUp(foundation.MakeNotification(notification))
}

//export comboBoxDelegate_TextField_TextView_Candidates_ForSelectedRange
func comboBoxDelegate_TextField_TextView_Candidates_ForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, candidates C.Array, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	if candidates.len > 0 {
		defer C.free(candidates.data)
	}
	candidatesSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(candidates.data))[:candidates.len:candidates.len]
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

//export comboBoxDelegate_TextField_TextView_CandidatesForSelectedRange
func comboBoxDelegate_TextField_TextView_CandidatesForSelectedRange(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
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

//export comboBoxDelegate_TextField_TextView_ShouldSelectCandidateAtIndex
func comboBoxDelegate_TextField_TextView_ShouldSelectCandidateAtIndex(hp C.uintptr_t, textField unsafe.Pointer, textView unsafe.Pointer, index C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.TextField_TextView_ShouldSelectCandidateAtIndex(MakeTextField(textField), MakeTextView(textView), uint(index))
	return C.bool(result)
}

//export comboBoxDelegate_Control_IsValidObject
func comboBoxDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export comboBoxDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func comboBoxDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export comboBoxDelegate_Control_DidFailToFormatString_ErrorDescription
func comboBoxDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export comboBoxDelegate_Control_TextShouldBeginEditing
func comboBoxDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export comboBoxDelegate_Control_TextShouldEndEditing
func comboBoxDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export comboBoxDelegate_Control_TextView_DoCommandBySelector
func comboBoxDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export comboBoxDelegate_ControlTextDidBeginEditing
func comboBoxDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export comboBoxDelegate_ControlTextDidChange
func comboBoxDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export comboBoxDelegate_ControlTextDidEndEditing
func comboBoxDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export ComboBoxDelegate_RespondsTo
func ComboBoxDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ComboBoxDelegate)
	switch selName {
	case "comboBoxSelectionDidChange:":
		return delegate.ComboBoxSelectionDidChange != nil
	case "comboBoxSelectionIsChanging:":
		return delegate.ComboBoxSelectionIsChanging != nil
	case "comboBoxWillDismiss:":
		return delegate.ComboBoxWillDismiss != nil
	case "comboBoxWillPopUp:":
		return delegate.ComboBoxWillPopUp != nil
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
