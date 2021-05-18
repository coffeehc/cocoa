package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "grid_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GridCell interface {
	objc.Object
	Column() GridColumn
	Row() GridRow
	ContentView() View
	SetContentView(value View)
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(value []LayoutConstraint)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

type NSGridCell struct {
	objc.NSObject
}

func MakeGridCell(ptr unsafe.Pointer) *NSGridCell {
	if ptr == nil {
		return nil
	}
	return &NSGridCell{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocGridCell() *NSGridCell {
	return MakeGridCell(C.C_GridCell_Alloc())
}

func (n *NSGridCell) Init() GridCell {
	result_ := C.C_NSGridCell_Init(n.Ptr())
	return MakeGridCell(result_)
}

func (n *NSGridCell) Column() GridColumn {
	result_ := C.C_NSGridCell_Column(n.Ptr())
	return MakeGridColumn(result_)
}

func (n *NSGridCell) Row() GridRow {
	result_ := C.C_NSGridCell_Row(n.Ptr())
	return MakeGridRow(result_)
}

func (n *NSGridCell) ContentView() View {
	result_ := C.C_NSGridCell_ContentView(n.Ptr())
	return MakeView(result_)
}

func (n *NSGridCell) SetContentView(value View) {
	C.C_NSGridCell_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func GridCell_EmptyContentView() View {
	result_ := C.C_NSGridCell_GridCell_EmptyContentView()
	return MakeView(result_)
}

func (n *NSGridCell) CustomPlacementConstraints() []LayoutConstraint {
	result_ := C.C_NSGridCell_CustomPlacementConstraints(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func (n *NSGridCell) SetCustomPlacementConstraints(value []LayoutConstraint) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSGridCell_SetCustomPlacementConstraints(n.Ptr(), cValue)
}

func (n *NSGridCell) RowAlignment() GridRowAlignment {
	result_ := C.C_NSGridCell_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result_))
}

func (n *NSGridCell) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridCell_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n *NSGridCell) XPlacement() GridCellPlacement {
	result_ := C.C_NSGridCell_XPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n *NSGridCell) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetXPlacement(n.Ptr(), C.int(int(value)))
}

func (n *NSGridCell) YPlacement() GridCellPlacement {
	result_ := C.C_NSGridCell_YPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n *NSGridCell) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetYPlacement(n.Ptr(), C.int(int(value)))
}
