package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "combo_box.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

// ComboBox is a view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value
type ComboBox interface {
	TextField

	HasVerticalScroller() bool
	SetHasVerticalScroller(hasVerticalScroller bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(intercellSpacing foundation.Size)
	IsButtonBordered() bool
	SetButtonBordered(buttonBordered bool)
	ItemHeight() float64
	SetItemHeight(itemHeight float64)
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(numberOfVisibleItems int)
	UsesDataSource() bool
	SetUsesDataSource(usesDataSource bool)
	Completes() bool
	SetCompletes(completes bool)
	ObjectValues() []objc.Object
	IndexOfSelectedItem() int
	AddItemsWithObjectValues(objects []objc.Object)
	AddItemWithObjectValue(object objc.Object)
	InsertItemWithObjectValue(object objc.Object, index int)
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
}

var _ ComboBox = (*NSComboBox)(nil)

type NSComboBox struct {
	NSTextField
}

func MakeComboBox(ptr unsafe.Pointer) *NSComboBox {
	if ptr == nil {
		return nil
	}
	return &NSComboBox{
		NSTextField: *MakeTextField(ptr),
	}
}

func (c *NSComboBox) HasVerticalScroller() bool {
	return bool(C.ComboBox_HasVerticalScroller(c.Ptr()))
}

func (c *NSComboBox) SetHasVerticalScroller(hasVerticalScroller bool) {
	C.ComboBox_SetHasVerticalScroller(c.Ptr(), C.bool(hasVerticalScroller))
}

func (c *NSComboBox) IntercellSpacing() foundation.Size {
	return toSize(C.ComboBox_IntercellSpacing(c.Ptr()))
}

func (c *NSComboBox) SetIntercellSpacing(intercellSpacing foundation.Size) {
	C.ComboBox_SetIntercellSpacing(c.Ptr(), toNSSize(intercellSpacing))
}

func (c *NSComboBox) IsButtonBordered() bool {
	return bool(C.ComboBox_IsButtonBordered(c.Ptr()))
}

func (c *NSComboBox) SetButtonBordered(buttonBordered bool) {
	C.ComboBox_SetButtonBordered(c.Ptr(), C.bool(buttonBordered))
}

func (c *NSComboBox) ItemHeight() float64 {
	return float64(C.ComboBox_ItemHeight(c.Ptr()))
}

func (c *NSComboBox) SetItemHeight(itemHeight float64) {
	C.ComboBox_SetItemHeight(c.Ptr(), C.double(itemHeight))
}

func (c *NSComboBox) NumberOfVisibleItems() int {
	return int(C.ComboBox_NumberOfVisibleItems(c.Ptr()))
}

func (c *NSComboBox) SetNumberOfVisibleItems(numberOfVisibleItems int) {
	C.ComboBox_SetNumberOfVisibleItems(c.Ptr(), C.long(numberOfVisibleItems))
}

func (c *NSComboBox) UsesDataSource() bool {
	return bool(C.ComboBox_UsesDataSource(c.Ptr()))
}

func (c *NSComboBox) SetUsesDataSource(usesDataSource bool) {
	C.ComboBox_SetUsesDataSource(c.Ptr(), C.bool(usesDataSource))
}

func (c *NSComboBox) Completes() bool {
	return bool(C.ComboBox_Completes(c.Ptr()))
}

func (c *NSComboBox) SetCompletes(completes bool) {
	C.ComboBox_SetCompletes(c.Ptr(), C.bool(completes))
}

func (c *NSComboBox) ObjectValues() []objc.Object {
	var cArray C.Array = C.ComboBox_ObjectValues(c.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]objc.Object, len(result))
	for idx, r := range result {
		goArray[idx] = objc.MakeObject(r)
	}
	return goArray
}

func (c *NSComboBox) IndexOfSelectedItem() int {
	return int(C.ComboBox_IndexOfSelectedItem(c.Ptr()))
}

func NewComboBox(frame foundation.Rect) ComboBox {
	return MakeComboBox(C.ComboBox_NewComboBox(toNSRect(frame)))
}

func (c *NSComboBox) AddItemsWithObjectValues(objects []objc.Object) {
	c_objectsData := make([]unsafe.Pointer, len(objects))
	for idx, v := range objects {
		c_objectsData[idx] = v.Ptr()
	}
	c_objects := C.Array{data: unsafe.Pointer(&c_objectsData[0]), len: C.int(len(objects))}
	C.ComboBox_AddItemsWithObjectValues(c.Ptr(), c_objects)
}

func (c *NSComboBox) AddItemWithObjectValue(object objc.Object) {
	C.ComboBox_AddItemWithObjectValue(c.Ptr(), toPointer(object))
}

func (c *NSComboBox) InsertItemWithObjectValue(object objc.Object, index int) {
	C.ComboBox_InsertItemWithObjectValue(c.Ptr(), toPointer(object), C.long(index))
}

func (c *NSComboBox) RemoveAllItems() {
	C.ComboBox_RemoveAllItems(c.Ptr())
}

func (c *NSComboBox) RemoveItemAtIndex(index int) {
	C.ComboBox_RemoveItemAtIndex(c.Ptr(), C.long(index))
}

func (c *NSComboBox) RemoveItemWithObjectValue(object objc.Object) {
	C.ComboBox_RemoveItemWithObjectValue(c.Ptr(), toPointer(object))
}

func (c *NSComboBox) IndexOfItemWithObjectValue(object objc.Object) int {
	return int(C.ComboBox_IndexOfItemWithObjectValue(c.Ptr(), toPointer(object)))
}

func (c *NSComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	return objc.MakeObject(C.ComboBox_ItemObjectValueAtIndex(c.Ptr(), C.long(index)))
}

func (c *NSComboBox) NoteNumberOfItemsChanged() {
	C.ComboBox_NoteNumberOfItemsChanged(c.Ptr())
}

func (c *NSComboBox) ReloadData() {
	C.ComboBox_ReloadData(c.Ptr())
}

func (c *NSComboBox) ScrollItemAtIndexToTop(index int) {
	C.ComboBox_ScrollItemAtIndexToTop(c.Ptr(), C.long(index))
}

func (c *NSComboBox) ScrollItemAtIndexToVisible(index int) {
	C.ComboBox_ScrollItemAtIndexToVisible(c.Ptr(), C.long(index))
}

func (c *NSComboBox) DeselectItemAtIndex(index int) {
	C.ComboBox_DeselectItemAtIndex(c.Ptr(), C.long(index))
}

func (c *NSComboBox) SelectItemAtIndex(index int) {
	C.ComboBox_SelectItemAtIndex(c.Ptr(), C.long(index))
}

func (c *NSComboBox) SelectItemWithObjectValue(object objc.Object) {
	C.ComboBox_SelectItemWithObjectValue(c.Ptr(), toPointer(object))
}
