package appkit

// #include "table_column.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableColumn interface {
	objc.Object
	SizeToFit()
	TableView() TableView
	SetTableView(value TableView)
	Width() coregraphics.Float
	SetWidth(value coregraphics.Float)
	MinWidth() coregraphics.Float
	SetMinWidth(value coregraphics.Float)
	MaxWidth() coregraphics.Float
	SetMaxWidth(value coregraphics.Float)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	Title() string
	SetTitle(value string)
	HeaderCell() TableHeaderCell
	SetHeaderCell(value TableHeaderCell)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	IsEditable() bool
	SetEditable(value bool)
	SortDescriptorPrototype() foundation.SortDescriptor
	SetSortDescriptorPrototype(value foundation.SortDescriptor)
	IsHidden() bool
	SetHidden(value bool)
	HeaderToolTip() string
	SetHeaderToolTip(value string)
}

type NSTableColumn struct {
	objc.NSObject
}

func MakeTableColumn(ptr unsafe.Pointer) *NSTableColumn {
	if ptr == nil {
		return nil
	}
	return &NSTableColumn{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTableColumn() *NSTableColumn {
	return MakeTableColumn(C.C_TableColumn_Alloc())
}

func (n *NSTableColumn) InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	result_ := C.C_NSTableColumn_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeTableColumn(result_)
}

func (n *NSTableColumn) InitWithCoder(coder foundation.Coder) TableColumn {
	result_ := C.C_NSTableColumn_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableColumn(result_)
}

func (n *NSTableColumn) Init() TableColumn {
	result_ := C.C_NSTableColumn_Init(n.Ptr())
	return MakeTableColumn(result_)
}

func (n *NSTableColumn) SizeToFit() {
	C.C_NSTableColumn_SizeToFit(n.Ptr())
}

func (n *NSTableColumn) TableView() TableView {
	result_ := C.C_NSTableColumn_TableView(n.Ptr())
	return MakeTableView(result_)
}

func (n *NSTableColumn) SetTableView(value TableView) {
	C.C_NSTableColumn_SetTableView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTableColumn) Width() coregraphics.Float {
	result_ := C.C_NSTableColumn_Width(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTableColumn) SetWidth(value coregraphics.Float) {
	C.C_NSTableColumn_SetWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSTableColumn) MinWidth() coregraphics.Float {
	result_ := C.C_NSTableColumn_MinWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTableColumn) SetMinWidth(value coregraphics.Float) {
	C.C_NSTableColumn_SetMinWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSTableColumn) MaxWidth() coregraphics.Float {
	result_ := C.C_NSTableColumn_MaxWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTableColumn) SetMaxWidth(value coregraphics.Float) {
	C.C_NSTableColumn_SetMaxWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSTableColumn) ResizingMask() TableColumnResizingOptions {
	result_ := C.C_NSTableColumn_ResizingMask(n.Ptr())
	return TableColumnResizingOptions(uint(result_))
}

func (n *NSTableColumn) SetResizingMask(value TableColumnResizingOptions) {
	C.C_NSTableColumn_SetResizingMask(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTableColumn) Title() string {
	result_ := C.C_NSTableColumn_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSTableColumn) SetTitle(value string) {
	C.C_NSTableColumn_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSTableColumn) HeaderCell() TableHeaderCell {
	result_ := C.C_NSTableColumn_HeaderCell(n.Ptr())
	return MakeTableHeaderCell(result_)
}

func (n *NSTableColumn) SetHeaderCell(value TableHeaderCell) {
	C.C_NSTableColumn_SetHeaderCell(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTableColumn) Identifier() UserInterfaceItemIdentifier {
	result_ := C.C_NSTableColumn_Identifier(n.Ptr())
	return UserInterfaceItemIdentifier(foundation.MakeString(result_).String())
}

func (n *NSTableColumn) SetIdentifier(value UserInterfaceItemIdentifier) {
	C.C_NSTableColumn_SetIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSTableColumn) IsEditable() bool {
	result_ := C.C_NSTableColumn_IsEditable(n.Ptr())
	return bool(result_)
}

func (n *NSTableColumn) SetEditable(value bool) {
	C.C_NSTableColumn_SetEditable(n.Ptr(), C.bool(value))
}

func (n *NSTableColumn) SortDescriptorPrototype() foundation.SortDescriptor {
	result_ := C.C_NSTableColumn_SortDescriptorPrototype(n.Ptr())
	return foundation.MakeSortDescriptor(result_)
}

func (n *NSTableColumn) SetSortDescriptorPrototype(value foundation.SortDescriptor) {
	C.C_NSTableColumn_SetSortDescriptorPrototype(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTableColumn) IsHidden() bool {
	result_ := C.C_NSTableColumn_IsHidden(n.Ptr())
	return bool(result_)
}

func (n *NSTableColumn) SetHidden(value bool) {
	C.C_NSTableColumn_SetHidden(n.Ptr(), C.bool(value))
}

func (n *NSTableColumn) HeaderToolTip() string {
	result_ := C.C_NSTableColumn_HeaderToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSTableColumn) SetHeaderToolTip(value string) {
	C.C_NSTableColumn_SetHeaderToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}
