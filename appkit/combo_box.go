package appkit

// #include "combo_box.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ComboBox interface {
	TextField
	AddItemsWithObjectValues(objects []objc.Object)
	AddItemWithObjectValue(object objc.Object)
	InsertItemWithObjectValue_AtIndex(object objc.Object, index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithObjectValue(object objc.Object)
	IndexOfItemWithObjectValue(object objc.Object) int
	ItemObjectValueAtIndex(index int) objc.Object
	NoteNumberOfItemsChanged()
	ReloadData()
	ScrollItemAtIndexToTop(index int)
	ScrollItemAtIndexToVisible(index int)
	DeselectItemAtIndex(index int)
	SelectItemAtIndex(index int)
	SelectItemWithObjectValue(object objc.Object)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	IsButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() coregraphics.Float
	SetItemHeight(value coregraphics.Float)
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	DataSource() objc.Object
	SetDataSource(value objc.Object)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	ObjectValues() []objc.Object
	NumberOfItems() int
	IndexOfSelectedItem() int
	ObjectValueOfSelectedItem() objc.Object
	Completes() bool
	SetCompletes(value bool)
}

type NSComboBox struct {
	NSTextField
}

func MakeComboBox(ptr unsafe.Pointer) NSComboBox {
	return NSComboBox{
		NSTextField: MakeTextField(ptr),
	}
}

func ComboBox_LabelWithAttributedString(attributedStringValue foundation.AttributedString) NSComboBox {
	result_ := C.C_NSComboBox_ComboBox_LabelWithAttributedString(objc.ExtractPtr(attributedStringValue))
	return MakeComboBox(result_)
}

func ComboBox_LabelWithString(stringValue string) NSComboBox {
	result_ := C.C_NSComboBox_ComboBox_LabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeComboBox(result_)
}

func ComboBox_TextFieldWithString(stringValue string) NSComboBox {
	result_ := C.C_NSComboBox_ComboBox_TextFieldWithString(foundation.NewString(stringValue).Ptr())
	return MakeComboBox(result_)
}

func ComboBox_WrappingLabelWithString(stringValue string) NSComboBox {
	result_ := C.C_NSComboBox_ComboBox_WrappingLabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeComboBox(result_)
}

func (n NSComboBox) InitWithFrame(frameRect foundation.Rect) NSComboBox {
	result_ := C.C_NSComboBox_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeComboBox(result_)
}

func (n NSComboBox) InitWithCoder(coder foundation.Coder) NSComboBox {
	result_ := C.C_NSComboBox_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeComboBox(result_)
}

func (n NSComboBox) Init() NSComboBox {
	result_ := C.C_NSComboBox_Init(n.Ptr())
	return MakeComboBox(result_)
}

func AllocComboBox() NSComboBox {
	result_ := C.C_NSComboBox_AllocComboBox()
	return MakeComboBox(result_)
}

func NewComboBox() NSComboBox {
	result_ := C.C_NSComboBox_NewComboBox()
	return MakeComboBox(result_)
}

func (n NSComboBox) Autorelease() NSComboBox {
	result_ := C.C_NSComboBox_Autorelease(n.Ptr())
	return MakeComboBox(result_)
}

func (n NSComboBox) Retain() NSComboBox {
	result_ := C.C_NSComboBox_Retain(n.Ptr())
	return MakeComboBox(result_)
}

func (n NSComboBox) AddItemsWithObjectValues(objects []objc.Object) {
	var cObjects C.Array
	if len(objects) > 0 {
		cObjectsData := make([]unsafe.Pointer, len(objects))
		for idx, v := range objects {
			cObjectsData[idx] = objc.ExtractPtr(v)
		}
		cObjects.data = unsafe.Pointer(&cObjectsData[0])
		cObjects.len = C.int(len(objects))
	}
	C.C_NSComboBox_AddItemsWithObjectValues(n.Ptr(), cObjects)
}

func (n NSComboBox) AddItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_AddItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n NSComboBox) InsertItemWithObjectValue_AtIndex(object objc.Object, index int) {
	C.C_NSComboBox_InsertItemWithObjectValue_AtIndex(n.Ptr(), objc.ExtractPtr(object), C.int(index))
}

func (n NSComboBox) RemoveAllItems() {
	C.C_NSComboBox_RemoveAllItems(n.Ptr())
}

func (n NSComboBox) RemoveItemAtIndex(index int) {
	C.C_NSComboBox_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSComboBox) RemoveItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_RemoveItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n NSComboBox) IndexOfItemWithObjectValue(object objc.Object) int {
	result_ := C.C_NSComboBox_IndexOfItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
	return int(result_)
}

func (n NSComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	result_ := C.C_NSComboBox_ItemObjectValueAtIndex(n.Ptr(), C.int(index))
	return objc.MakeObject(result_)
}

func (n NSComboBox) NoteNumberOfItemsChanged() {
	C.C_NSComboBox_NoteNumberOfItemsChanged(n.Ptr())
}

func (n NSComboBox) ReloadData() {
	C.C_NSComboBox_ReloadData(n.Ptr())
}

func (n NSComboBox) ScrollItemAtIndexToTop(index int) {
	C.C_NSComboBox_ScrollItemAtIndexToTop(n.Ptr(), C.int(index))
}

func (n NSComboBox) ScrollItemAtIndexToVisible(index int) {
	C.C_NSComboBox_ScrollItemAtIndexToVisible(n.Ptr(), C.int(index))
}

func (n NSComboBox) DeselectItemAtIndex(index int) {
	C.C_NSComboBox_DeselectItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSComboBox) SelectItemAtIndex(index int) {
	C.C_NSComboBox_SelectItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSComboBox) SelectItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_SelectItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n NSComboBox) HasVerticalScroller() bool {
	result_ := C.C_NSComboBox_HasVerticalScroller(n.Ptr())
	return bool(result_)
}

func (n NSComboBox) SetHasVerticalScroller(value bool) {
	C.C_NSComboBox_SetHasVerticalScroller(n.Ptr(), C.bool(value))
}

func (n NSComboBox) IntercellSpacing() foundation.Size {
	result_ := C.C_NSComboBox_IntercellSpacing(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSComboBox) SetIntercellSpacing(value foundation.Size) {
	C.C_NSComboBox_SetIntercellSpacing(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSComboBox) IsButtonBordered() bool {
	result_ := C.C_NSComboBox_IsButtonBordered(n.Ptr())
	return bool(result_)
}

func (n NSComboBox) SetButtonBordered(value bool) {
	C.C_NSComboBox_SetButtonBordered(n.Ptr(), C.bool(value))
}

func (n NSComboBox) ItemHeight() coregraphics.Float {
	result_ := C.C_NSComboBox_ItemHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSComboBox) SetItemHeight(value coregraphics.Float) {
	C.C_NSComboBox_SetItemHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSComboBox) NumberOfVisibleItems() int {
	result_ := C.C_NSComboBox_NumberOfVisibleItems(n.Ptr())
	return int(result_)
}

func (n NSComboBox) SetNumberOfVisibleItems(value int) {
	C.C_NSComboBox_SetNumberOfVisibleItems(n.Ptr(), C.int(value))
}

func (n NSComboBox) DataSource() objc.Object {
	result_ := C.C_NSComboBox_DataSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSComboBox) SetDataSource(value objc.Object) {
	C.C_NSComboBox_SetDataSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSComboBox) UsesDataSource() bool {
	result_ := C.C_NSComboBox_UsesDataSource(n.Ptr())
	return bool(result_)
}

func (n NSComboBox) SetUsesDataSource(value bool) {
	C.C_NSComboBox_SetUsesDataSource(n.Ptr(), C.bool(value))
}

func (n NSComboBox) ObjectValues() []objc.Object {
	result_ := C.C_NSComboBox_ObjectValues(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}

func (n NSComboBox) NumberOfItems() int {
	result_ := C.C_NSComboBox_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n NSComboBox) IndexOfSelectedItem() int {
	result_ := C.C_NSComboBox_IndexOfSelectedItem(n.Ptr())
	return int(result_)
}

func (n NSComboBox) ObjectValueOfSelectedItem() objc.Object {
	result_ := C.C_NSComboBox_ObjectValueOfSelectedItem(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSComboBox) Completes() bool {
	result_ := C.C_NSComboBox_Completes(n.Ptr())
	return bool(result_)
}

func (n NSComboBox) SetCompletes(value bool) {
	C.C_NSComboBox_SetCompletes(n.Ptr(), C.bool(value))
}

type ComboBoxDataSource interface {
	HasComboBox_CompletedString() bool
	ComboBox_CompletedString(comboBox ComboBox, _string string) string
	HasComboBox_IndexOfItemWithStringValue() bool
	ComboBox_IndexOfItemWithStringValue(comboBox ComboBox, _string string) uint
	HasComboBox_ObjectValueForItemAtIndex() bool
	ComboBox_ObjectValueForItemAtIndex(comboBox ComboBox, index int) objc.Object
	HasNumberOfItemsInComboBox() bool
	NumberOfItemsInComboBox(comboBox ComboBox) int
}

func ComboBoxDataSourceToObjc(protocol ComboBoxDataSource) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapComboBoxDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export comboBoxDataSource_ComboBox_CompletedString
func comboBoxDataSource_ComboBox_CompletedString(hp C.uintptr_t, comboBox unsafe.Pointer, _string unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ComboBoxDataSource)
	result := protocol.ComboBox_CompletedString(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return foundation.NewString(result).Ptr()
}

//export comboBoxDataSource_ComboBox_IndexOfItemWithStringValue
func comboBoxDataSource_ComboBox_IndexOfItemWithStringValue(hp C.uintptr_t, comboBox unsafe.Pointer, _string unsafe.Pointer) C.uint {
	protocol := cgo.Handle(hp).Value().(ComboBoxDataSource)
	result := protocol.ComboBox_IndexOfItemWithStringValue(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return C.uint(result)
}

//export comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex
func comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex(hp C.uintptr_t, comboBox unsafe.Pointer, index C.int) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ComboBoxDataSource)
	result := protocol.ComboBox_ObjectValueForItemAtIndex(MakeComboBox(comboBox), int(index))
	return objc.ExtractPtr(result)
}

//export comboBoxDataSource_NumberOfItemsInComboBox
func comboBoxDataSource_NumberOfItemsInComboBox(hp C.uintptr_t, comboBox unsafe.Pointer) C.int {
	protocol := cgo.Handle(hp).Value().(ComboBoxDataSource)
	result := protocol.NumberOfItemsInComboBox(MakeComboBox(comboBox))
	return C.int(result)
}

//export ComboBoxDataSource_RespondsTo
func ComboBoxDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(ComboBoxDataSource)
	_ = protocol
	switch selName {
	case "comboBox:completedString:":
		return protocol.HasComboBox_CompletedString()
	case "comboBox:indexOfItemWithStringValue:":
		return protocol.HasComboBox_IndexOfItemWithStringValue()
	case "comboBox:objectValueForItemAtIndex:":
		return protocol.HasComboBox_ObjectValueForItemAtIndex()
	case "numberOfItemsInComboBox:":
		return protocol.HasNumberOfItemsInComboBox()
	default:
		return false
	}
}

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
