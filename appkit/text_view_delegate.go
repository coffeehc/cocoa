package appkit

// #include "text_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TextViewDelegate struct {
	UndoManagerForTextView                                            func(view TextView) foundation.UndoManager
	TextView_WillDisplayToolTip_ForCharacterAtIndex                   func(textView TextView, tooltip string, characterIndex uint) string
	TextView_URLForContentsOfTextAttachment_AtIndex                   func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.URL
	TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange   func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	TextViewDidChangeSelection                                        func(notification foundation.Notification)
	TextView_WritablePasteboardTypesForCell_AtIndex                   func(view TextView, cell objc.Object, charIndex uint) []PasteboardType
	TextView_WriteCell_AtIndex_ToPasteboard_Type                      func(view TextView, cell objc.Object, charIndex uint, pboard Pasteboard, _type PasteboardType) bool
	TextView_ShouldChangeTextInRange_ReplacementString                func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	TextView_ShouldChangeTextInRanges_ReplacementStrings              func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	TextView_ShouldChangeTypingAttributes_ToAttributes                func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.Object
	TextViewDidChangeTypingAttributes                                 func(notification foundation.Notification)
	TextView_ClickedOnCell_InRect_AtIndex                             func(textView TextView, cell objc.Object, cellFrame foundation.Rect, charIndex uint)
	TextView_DoubleClickedOnCell_InRect_AtIndex                       func(textView TextView, cell objc.Object, cellFrame foundation.Rect, charIndex uint)
	TextView_ClickedOnLink_AtIndex                                    func(textView TextView, link objc.Object, charIndex uint) bool
	TextView_ShouldSetSpellingState_Range                             func(textView TextView, value int, affectedCharRange foundation.Range) int
	TextView_DraggedCell_InRect_Event_AtIndex                         func(view TextView, cell objc.Object, rect foundation.Rect, event Event, charIndex uint)
	TextView_WillShowSharingServicePicker_ForItems                    func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) SharingServicePicker
	TextView_DoCommandBySelector                                      func(textView TextView, commandSelector objc.Selector) bool
	TextView_Menu_ForEvent_AtIndex                                    func(view TextView, menu Menu, event Event, charIndex uint) Menu
	TextView_Candidates_ForSelectedRange                              func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	TextView_CandidatesForSelectedRange                               func(textView TextView, selectedRange foundation.Range) []objc.Object
	TextView_ShouldSelectCandidateAtIndex                             func(textView TextView, index uint) bool
	TextView_ShouldUpdateTouchBarItemIdentifiers                      func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	TextDidChange                                                     func(notification foundation.Notification)
	TextShouldBeginEditing                                            func(textObject Text) bool
	TextDidBeginEditing                                               func(notification foundation.Notification)
	TextShouldEndEditing                                              func(textObject Text) bool
	TextDidEndEditing                                                 func(notification foundation.Notification)
}

func (delegate *TextViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTextViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export textViewDelegate_UndoManagerForTextView
func textViewDelegate_UndoManagerForTextView(hp C.uintptr_t, view unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.UndoManagerForTextView(MakeTextView(view))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex
func textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex(hp C.uintptr_t, textView unsafe.Pointer, tooltip unsafe.Pointer, characterIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WillDisplayToolTip_ForCharacterAtIndex(MakeTextView(textView), foundation.MakeString(tooltip).String(), uint(characterIndex))
	return foundation.NewString(result).Ptr()
}

//export textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex
func textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, textAttachment unsafe.Pointer, charIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_URLForContentsOfTextAttachment_AtIndex(MakeTextView(textView), MakeTextAttachment(textAttachment), uint(charIndex))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange
func textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(hp C.uintptr_t, textView unsafe.Pointer, oldSelectedCharRange C.NSRange, newSelectedCharRange C.NSRange) C.NSRange {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&oldSelectedCharRange)), foundation.FromNSRangePointer(unsafe.Pointer(&newSelectedCharRange)))
	return *(*C.NSRange)(foundation.ToNSRangePointer(result))
}

//export textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges
func textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(hp C.uintptr_t, textView unsafe.Pointer, oldSelectedCharRanges C.Array, newSelectedCharRanges C.Array) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if oldSelectedCharRanges.len > 0 {
		defer C.free(oldSelectedCharRanges.data)
	}
	oldSelectedCharRangesSlice := unsafe.Slice((*unsafe.Pointer)(oldSelectedCharRanges.data), int(oldSelectedCharRanges.len))
	var goOldSelectedCharRanges = make([]foundation.Value, len(oldSelectedCharRangesSlice))
	for idx, r := range oldSelectedCharRangesSlice {
		goOldSelectedCharRanges[idx] = foundation.MakeValue(r)
	}
	if newSelectedCharRanges.len > 0 {
		defer C.free(newSelectedCharRanges.data)
	}
	newSelectedCharRangesSlice := unsafe.Slice((*unsafe.Pointer)(newSelectedCharRanges.data), int(newSelectedCharRanges.len))
	var goNewSelectedCharRanges = make([]foundation.Value, len(newSelectedCharRangesSlice))
	for idx, r := range newSelectedCharRangesSlice {
		goNewSelectedCharRanges[idx] = foundation.MakeValue(r)
	}
	result := delegate.TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(MakeTextView(textView), goOldSelectedCharRanges, goNewSelectedCharRanges)
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

//export textViewDelegate_TextViewDidChangeSelection
func textViewDelegate_TextViewDidChangeSelection(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextViewDidChangeSelection(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex
func textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, charIndex C.uint) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WritablePasteboardTypesForCell_AtIndex(MakeTextView(view), objc.MakeObject(cell), uint(charIndex))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type
func textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, charIndex C.uint, pboard unsafe.Pointer, _type unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WriteCell_AtIndex_ToPasteboard_Type(MakeTextView(view), objc.MakeObject(cell), uint(charIndex), MakePasteboard(pboard), PasteboardType(foundation.MakeString(_type).String()))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString
func textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString(hp C.uintptr_t, textView unsafe.Pointer, affectedCharRange C.NSRange, replacementString unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldChangeTextInRange_ReplacementString(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&affectedCharRange)), foundation.MakeString(replacementString).String())
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings
func textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings(hp C.uintptr_t, textView unsafe.Pointer, affectedRanges C.Array, replacementStrings C.Array) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if affectedRanges.len > 0 {
		defer C.free(affectedRanges.data)
	}
	affectedRangesSlice := unsafe.Slice((*unsafe.Pointer)(affectedRanges.data), int(affectedRanges.len))
	var goAffectedRanges = make([]foundation.Value, len(affectedRangesSlice))
	for idx, r := range affectedRangesSlice {
		goAffectedRanges[idx] = foundation.MakeValue(r)
	}
	if replacementStrings.len > 0 {
		defer C.free(replacementStrings.data)
	}
	replacementStringsSlice := unsafe.Slice((*unsafe.Pointer)(replacementStrings.data), int(replacementStrings.len))
	var goReplacementStrings = make([]string, len(replacementStringsSlice))
	for idx, r := range replacementStringsSlice {
		goReplacementStrings[idx] = foundation.MakeString(r).String()
	}
	result := delegate.TextView_ShouldChangeTextInRanges_ReplacementStrings(MakeTextView(textView), goAffectedRanges, goReplacementStrings)
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTypingAttributes_ToAttributes
func textViewDelegate_TextView_ShouldChangeTypingAttributes_ToAttributes(hp C.uintptr_t, textView unsafe.Pointer, oldTypingAttributes C.Dictionary, newTypingAttributes C.Dictionary) C.Dictionary {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if oldTypingAttributes.len > 0 {
		defer C.free(oldTypingAttributes.key_data)
		defer C.free(oldTypingAttributes.value_data)
	}
	oldTypingAttributesKeySlice := unsafe.Slice((*unsafe.Pointer)(oldTypingAttributes.key_data), int(oldTypingAttributes.len))
	oldTypingAttributesValueSlice := unsafe.Slice((*unsafe.Pointer)(oldTypingAttributes.value_data), int(oldTypingAttributes.len))
	var goOldTypingAttributes = make(map[string]objc.Object)
	for idx, k := range oldTypingAttributesKeySlice {
		v := oldTypingAttributesValueSlice[idx]
		goOldTypingAttributes[foundation.MakeString(k).String()] = objc.MakeObject(v)
	}
	if newTypingAttributes.len > 0 {
		defer C.free(newTypingAttributes.key_data)
		defer C.free(newTypingAttributes.value_data)
	}
	newTypingAttributesKeySlice := unsafe.Slice((*unsafe.Pointer)(newTypingAttributes.key_data), int(newTypingAttributes.len))
	newTypingAttributesValueSlice := unsafe.Slice((*unsafe.Pointer)(newTypingAttributes.value_data), int(newTypingAttributes.len))
	var goNewTypingAttributes = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range newTypingAttributesKeySlice {
		v := newTypingAttributesValueSlice[idx]
		goNewTypingAttributes[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	result := delegate.TextView_ShouldChangeTypingAttributes_ToAttributes(MakeTextView(textView), goOldTypingAttributes, goNewTypingAttributes)
	var cResult C.Dictionary
	if len(result) > 0 {
		cResultKeyData := make([]unsafe.Pointer, len(result))
		cResultValueData := make([]unsafe.Pointer, len(result))
		var idx = 0
		for k, v := range result {
			cResultKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cResultValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cResult.key_data = unsafe.Pointer(&cResultKeyData[0])
		cResult.value_data = unsafe.Pointer(&cResultValueData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextViewDidChangeTypingAttributes
func textViewDelegate_TextViewDidChangeTypingAttributes(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextViewDidChangeTypingAttributes(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex
func textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, cell unsafe.Pointer, cellFrame C.CGRect, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_ClickedOnCell_InRect_AtIndex(MakeTextView(textView), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&cellFrame))), uint(charIndex))
}

//export textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex
func textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, cell unsafe.Pointer, cellFrame C.CGRect, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_DoubleClickedOnCell_InRect_AtIndex(MakeTextView(textView), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&cellFrame))), uint(charIndex))
}

//export textViewDelegate_TextView_ClickedOnLink_AtIndex
func textViewDelegate_TextView_ClickedOnLink_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, link unsafe.Pointer, charIndex C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ClickedOnLink_AtIndex(MakeTextView(textView), objc.MakeObject(link), uint(charIndex))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldSetSpellingState_Range
func textViewDelegate_TextView_ShouldSetSpellingState_Range(hp C.uintptr_t, textView unsafe.Pointer, value C.int, affectedCharRange C.NSRange) C.int {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldSetSpellingState_Range(MakeTextView(textView), int(value), foundation.FromNSRangePointer(unsafe.Pointer(&affectedCharRange)))
	return C.int(result)
}

//export textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex
func textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, rect C.CGRect, event unsafe.Pointer, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_DraggedCell_InRect_Event_AtIndex(MakeTextView(view), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&rect))), MakeEvent(event), uint(charIndex))
}

//export textViewDelegate_TextView_WillShowSharingServicePicker_ForItems
func textViewDelegate_TextView_WillShowSharingServicePicker_ForItems(hp C.uintptr_t, textView unsafe.Pointer, servicePicker unsafe.Pointer, items C.Array) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if items.len > 0 {
		defer C.free(items.data)
	}
	itemsSlice := unsafe.Slice((*unsafe.Pointer)(items.data), int(items.len))
	var goItems = make([]objc.Object, len(itemsSlice))
	for idx, r := range itemsSlice {
		goItems[idx] = objc.MakeObject(r)
	}
	result := delegate.TextView_WillShowSharingServicePicker_ForItems(MakeTextView(textView), MakeSharingServicePicker(servicePicker), goItems)
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_DoCommandBySelector
func textViewDelegate_TextView_DoCommandBySelector(hp C.uintptr_t, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_DoCommandBySelector(MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export textViewDelegate_TextView_Menu_ForEvent_AtIndex
func textViewDelegate_TextView_Menu_ForEvent_AtIndex(hp C.uintptr_t, view unsafe.Pointer, menu unsafe.Pointer, event unsafe.Pointer, charIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_Menu_ForEvent_AtIndex(MakeTextView(view), MakeMenu(menu), MakeEvent(event), uint(charIndex))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_Candidates_ForSelectedRange
func textViewDelegate_TextView_Candidates_ForSelectedRange(hp C.uintptr_t, textView unsafe.Pointer, candidates C.Array, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if candidates.len > 0 {
		defer C.free(candidates.data)
	}
	candidatesSlice := unsafe.Slice((*unsafe.Pointer)(candidates.data), int(candidates.len))
	var goCandidates = make([]foundation.TextCheckingResult, len(candidatesSlice))
	for idx, r := range candidatesSlice {
		goCandidates[idx] = foundation.MakeTextCheckingResult(r)
	}
	result := delegate.TextView_Candidates_ForSelectedRange(MakeTextView(textView), goCandidates, foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
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

//export textViewDelegate_TextView_CandidatesForSelectedRange
func textViewDelegate_TextView_CandidatesForSelectedRange(hp C.uintptr_t, textView unsafe.Pointer, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_CandidatesForSelectedRange(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
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

//export textViewDelegate_TextView_ShouldSelectCandidateAtIndex
func textViewDelegate_TextView_ShouldSelectCandidateAtIndex(hp C.uintptr_t, textView unsafe.Pointer, index C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldSelectCandidateAtIndex(MakeTextView(textView), uint(index))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers
func textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers(hp C.uintptr_t, textView unsafe.Pointer, identifiers C.Array) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if identifiers.len > 0 {
		defer C.free(identifiers.data)
	}
	identifiersSlice := unsafe.Slice((*unsafe.Pointer)(identifiers.data), int(identifiers.len))
	var goIdentifiers = make([]TouchBarItemIdentifier, len(identifiersSlice))
	for idx, r := range identifiersSlice {
		goIdentifiers[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	result := delegate.TextView_ShouldUpdateTouchBarItemIdentifiers(MakeTextView(textView), goIdentifiers)
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextDidChange
func textViewDelegate_TextDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidChange(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextShouldBeginEditing
func textViewDelegate_TextShouldBeginEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextShouldBeginEditing(MakeText(textObject))
	return C.bool(result)
}

//export textViewDelegate_TextDidBeginEditing
func textViewDelegate_TextDidBeginEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidBeginEditing(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextShouldEndEditing
func textViewDelegate_TextShouldEndEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextShouldEndEditing(MakeText(textObject))
	return C.bool(result)
}

//export textViewDelegate_TextDidEndEditing
func textViewDelegate_TextDidEndEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidEndEditing(foundation.MakeNotification(notification))
}

//export TextViewDelegate_RespondsTo
func TextViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	switch selName {
	case "undoManagerForTextView:":
		return delegate.UndoManagerForTextView != nil
	case "textView:willDisplayToolTip:forCharacterAtIndex:":
		return delegate.TextView_WillDisplayToolTip_ForCharacterAtIndex != nil
	case "textView:URLForContentsOfTextAttachment:atIndex:":
		return delegate.TextView_URLForContentsOfTextAttachment_AtIndex != nil
	case "textView:willChangeSelectionFromCharacterRange:toCharacterRange:":
		return delegate.TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange != nil
	case "textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:":
		return delegate.TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges != nil
	case "textViewDidChangeSelection:":
		return delegate.TextViewDidChangeSelection != nil
	case "textView:writablePasteboardTypesForCell:atIndex:":
		return delegate.TextView_WritablePasteboardTypesForCell_AtIndex != nil
	case "textView:writeCell:atIndex:toPasteboard:type:":
		return delegate.TextView_WriteCell_AtIndex_ToPasteboard_Type != nil
	case "textView:shouldChangeTextInRange:replacementString:":
		return delegate.TextView_ShouldChangeTextInRange_ReplacementString != nil
	case "textView:shouldChangeTextInRanges:replacementStrings:":
		return delegate.TextView_ShouldChangeTextInRanges_ReplacementStrings != nil
	case "textView:shouldChangeTypingAttributes:toAttributes:":
		return delegate.TextView_ShouldChangeTypingAttributes_ToAttributes != nil
	case "textViewDidChangeTypingAttributes:":
		return delegate.TextViewDidChangeTypingAttributes != nil
	case "textView:clickedOnCell:inRect:atIndex:":
		return delegate.TextView_ClickedOnCell_InRect_AtIndex != nil
	case "textView:doubleClickedOnCell:inRect:atIndex:":
		return delegate.TextView_DoubleClickedOnCell_InRect_AtIndex != nil
	case "textView:clickedOnLink:atIndex:":
		return delegate.TextView_ClickedOnLink_AtIndex != nil
	case "textView:shouldSetSpellingState:range:":
		return delegate.TextView_ShouldSetSpellingState_Range != nil
	case "textView:draggedCell:inRect:event:atIndex:":
		return delegate.TextView_DraggedCell_InRect_Event_AtIndex != nil
	case "textView:willShowSharingServicePicker:forItems:":
		return delegate.TextView_WillShowSharingServicePicker_ForItems != nil
	case "textView:doCommandBySelector:":
		return delegate.TextView_DoCommandBySelector != nil
	case "textView:menu:forEvent:atIndex:":
		return delegate.TextView_Menu_ForEvent_AtIndex != nil
	case "textView:candidates:forSelectedRange:":
		return delegate.TextView_Candidates_ForSelectedRange != nil
	case "textView:candidatesForSelectedRange:":
		return delegate.TextView_CandidatesForSelectedRange != nil
	case "textView:shouldSelectCandidateAtIndex:":
		return delegate.TextView_ShouldSelectCandidateAtIndex != nil
	case "textView:shouldUpdateTouchBarItemIdentifiers:":
		return delegate.TextView_ShouldUpdateTouchBarItemIdentifiers != nil
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
