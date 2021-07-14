package appkit

// #include "combo_box_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ComboBoxDataSource struct {
	ComboBox_CompletedString            func(comboBox ComboBox, _string string) string
	ComboBox_IndexOfItemWithStringValue func(comboBox ComboBox, _string string) uint
	ComboBox_ObjectValueForItemAtIndex  func(comboBox ComboBox, index int) objc.Object
	NumberOfItemsInComboBox             func(comboBox ComboBox) int
}

func WrapComboBoxDataSource(delegate *ComboBoxDataSource) objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapComboBoxDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export comboBoxDataSource_ComboBox_CompletedString
func comboBoxDataSource_ComboBox_CompletedString(hp C.uintptr_t, comboBox unsafe.Pointer, _string unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDataSource)
	result := delegate.ComboBox_CompletedString(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return foundation.NewString(result).Ptr()
}

//export comboBoxDataSource_ComboBox_IndexOfItemWithStringValue
func comboBoxDataSource_ComboBox_IndexOfItemWithStringValue(hp C.uintptr_t, comboBox unsafe.Pointer, _string unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDataSource)
	result := delegate.ComboBox_IndexOfItemWithStringValue(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return C.uint(result)
}

//export comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex
func comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex(hp C.uintptr_t, comboBox unsafe.Pointer, index C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDataSource)
	result := delegate.ComboBox_ObjectValueForItemAtIndex(MakeComboBox(comboBox), int(index))
	return objc.ExtractPtr(result)
}

//export comboBoxDataSource_NumberOfItemsInComboBox
func comboBoxDataSource_NumberOfItemsInComboBox(hp C.uintptr_t, comboBox unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*ComboBoxDataSource)
	result := delegate.NumberOfItemsInComboBox(MakeComboBox(comboBox))
	return C.int(result)
}

//export ComboBoxDataSource_RespondsTo
func ComboBoxDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ComboBoxDataSource)
	switch selName {
	case "comboBox:completedString:":
		return delegate.ComboBox_CompletedString != nil
	case "comboBox:indexOfItemWithStringValue:":
		return delegate.ComboBox_IndexOfItemWithStringValue != nil
	case "comboBox:objectValueForItemAtIndex:":
		return delegate.ComboBox_ObjectValueForItemAtIndex != nil
	case "numberOfItemsInComboBox:":
		return delegate.NumberOfItemsInComboBox != nil
	default:
		return false
	}
}

//export deleteComboBoxDataSource
func deleteComboBoxDataSource(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
