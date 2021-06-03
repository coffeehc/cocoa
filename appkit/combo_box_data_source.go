package appkit

// #include "combo_box_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ComboBoxDataSource struct {
	ComboBox_CompletedString            func(comboBox ComboBox, _string string) string
	ComboBox_IndexOfItemWithStringValue func(comboBox ComboBox, _string string) uint
	ComboBox_ObjectValueForItemAtIndex  func(comboBox ComboBox, index int) objc.Object
	NumberOfItemsInComboBox             func(comboBox ComboBox) int
}

func WrapComboBoxDataSource(delegate *ComboBoxDataSource) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapComboBoxDataSource(C.long(id))
	return objc.MakeObject(ptr)
}

//export comboBoxDataSource_ComboBox_CompletedString
func comboBoxDataSource_ComboBox_CompletedString(id int64, comboBox unsafe.Pointer, _string unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*ComboBoxDataSource)
	result := delegate.ComboBox_CompletedString(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return foundation.NewString(result).Ptr()
}

//export comboBoxDataSource_ComboBox_IndexOfItemWithStringValue
func comboBoxDataSource_ComboBox_IndexOfItemWithStringValue(id int64, comboBox unsafe.Pointer, _string unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*ComboBoxDataSource)
	result := delegate.ComboBox_IndexOfItemWithStringValue(MakeComboBox(comboBox), foundation.MakeString(_string).String())
	return C.uint(result)
}

//export comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex
func comboBoxDataSource_ComboBox_ObjectValueForItemAtIndex(id int64, comboBox unsafe.Pointer, index C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*ComboBoxDataSource)
	result := delegate.ComboBox_ObjectValueForItemAtIndex(MakeComboBox(comboBox), int(index))
	return objc.ExtractPtr(result)
}

//export comboBoxDataSource_NumberOfItemsInComboBox
func comboBoxDataSource_NumberOfItemsInComboBox(id int64, comboBox unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*ComboBoxDataSource)
	result := delegate.NumberOfItemsInComboBox(MakeComboBox(comboBox))
	return C.int(result)
}

//export ComboBoxDataSource_RespondsTo
func ComboBoxDataSource_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ComboBoxDataSource)
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
func deleteComboBoxDataSource(id int64) {
	resources.Delete(id)
}
