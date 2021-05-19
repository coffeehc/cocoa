package appkit

// #include "combo_box.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func MakeComboBox(ptr unsafe.Pointer) *NSComboBox {
	if ptr == nil {
		return nil
	}
	return &NSComboBox{
		NSTextField: *MakeTextField(ptr),
	}
}

func AllocComboBox() *NSComboBox {
	return MakeComboBox(C.C_ComboBox_Alloc())
}

func (n *NSComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	result_ := C.C_NSComboBox_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeComboBox(result_)
}

func (n *NSComboBox) InitWithCoder(coder foundation.Coder) ComboBox {
	result_ := C.C_NSComboBox_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeComboBox(result_)
}

func (n *NSComboBox) Init() ComboBox {
	result_ := C.C_NSComboBox_Init(n.Ptr())
	return MakeComboBox(result_)
}

func (n *NSComboBox) AddItemsWithObjectValues(objects []objc.Object) {
	cObjectsData := make([]unsafe.Pointer, len(objects))
	for idx, v := range objects {
		cObjectsData[idx] = objc.ExtractPtr(v)
	}
	cObjects := C.Array{data: unsafe.Pointer(&cObjectsData[0]), len: C.int(len(objects))}
	C.C_NSComboBox_AddItemsWithObjectValues(n.Ptr(), cObjects)
}

func (n *NSComboBox) AddItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_AddItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSComboBox) InsertItemWithObjectValue_AtIndex(object objc.Object, index int) {
	C.C_NSComboBox_InsertItemWithObjectValue_AtIndex(n.Ptr(), objc.ExtractPtr(object), C.int(index))
}

func (n *NSComboBox) RemoveAllItems() {
	C.C_NSComboBox_RemoveAllItems(n.Ptr())
}

func (n *NSComboBox) RemoveItemAtIndex(index int) {
	C.C_NSComboBox_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSComboBox) RemoveItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_RemoveItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSComboBox) IndexOfItemWithObjectValue(object objc.Object) int {
	result_ := C.C_NSComboBox_IndexOfItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
	return int(result_)
}

func (n *NSComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	result_ := C.C_NSComboBox_ItemObjectValueAtIndex(n.Ptr(), C.int(index))
	return objc.MakeObject(result_)
}

func (n *NSComboBox) NoteNumberOfItemsChanged() {
	C.C_NSComboBox_NoteNumberOfItemsChanged(n.Ptr())
}

func (n *NSComboBox) ReloadData() {
	C.C_NSComboBox_ReloadData(n.Ptr())
}

func (n *NSComboBox) ScrollItemAtIndexToTop(index int) {
	C.C_NSComboBox_ScrollItemAtIndexToTop(n.Ptr(), C.int(index))
}

func (n *NSComboBox) ScrollItemAtIndexToVisible(index int) {
	C.C_NSComboBox_ScrollItemAtIndexToVisible(n.Ptr(), C.int(index))
}

func (n *NSComboBox) DeselectItemAtIndex(index int) {
	C.C_NSComboBox_DeselectItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSComboBox) SelectItemAtIndex(index int) {
	C.C_NSComboBox_SelectItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSComboBox) SelectItemWithObjectValue(object objc.Object) {
	C.C_NSComboBox_SelectItemWithObjectValue(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSComboBox) HasVerticalScroller() bool {
	result_ := C.C_NSComboBox_HasVerticalScroller(n.Ptr())
	return bool(result_)
}

func (n *NSComboBox) SetHasVerticalScroller(value bool) {
	C.C_NSComboBox_SetHasVerticalScroller(n.Ptr(), C.bool(value))
}

func (n *NSComboBox) IntercellSpacing() foundation.Size {
	result_ := C.C_NSComboBox_IntercellSpacing(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSComboBox) SetIntercellSpacing(value foundation.Size) {
	C.C_NSComboBox_SetIntercellSpacing(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSComboBox) IsButtonBordered() bool {
	result_ := C.C_NSComboBox_IsButtonBordered(n.Ptr())
	return bool(result_)
}

func (n *NSComboBox) SetButtonBordered(value bool) {
	C.C_NSComboBox_SetButtonBordered(n.Ptr(), C.bool(value))
}

func (n *NSComboBox) ItemHeight() coregraphics.Float {
	result_ := C.C_NSComboBox_ItemHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSComboBox) SetItemHeight(value coregraphics.Float) {
	C.C_NSComboBox_SetItemHeight(n.Ptr(), C.double(float64(value)))
}

func (n *NSComboBox) NumberOfVisibleItems() int {
	result_ := C.C_NSComboBox_NumberOfVisibleItems(n.Ptr())
	return int(result_)
}

func (n *NSComboBox) SetNumberOfVisibleItems(value int) {
	C.C_NSComboBox_SetNumberOfVisibleItems(n.Ptr(), C.int(value))
}

func (n *NSComboBox) DataSource() objc.Object {
	result_ := C.C_NSComboBox_DataSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSComboBox) SetDataSource(value objc.Object) {
	C.C_NSComboBox_SetDataSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSComboBox) UsesDataSource() bool {
	result_ := C.C_NSComboBox_UsesDataSource(n.Ptr())
	return bool(result_)
}

func (n *NSComboBox) SetUsesDataSource(value bool) {
	C.C_NSComboBox_SetUsesDataSource(n.Ptr(), C.bool(value))
}

func (n *NSComboBox) ObjectValues() []objc.Object {
	result_ := C.C_NSComboBox_ObjectValues(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}

func (n *NSComboBox) NumberOfItems() int {
	result_ := C.C_NSComboBox_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n *NSComboBox) IndexOfSelectedItem() int {
	result_ := C.C_NSComboBox_IndexOfSelectedItem(n.Ptr())
	return int(result_)
}

func (n *NSComboBox) ObjectValueOfSelectedItem() objc.Object {
	result_ := C.C_NSComboBox_ObjectValueOfSelectedItem(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSComboBox) Completes() bool {
	result_ := C.C_NSComboBox_Completes(n.Ptr())
	return bool(result_)
}

func (n *NSComboBox) SetCompletes(value bool) {
	C.C_NSComboBox_SetCompletes(n.Ptr(), C.bool(value))
}
