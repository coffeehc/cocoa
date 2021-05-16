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
	result := C.C_NSGridCell_Init(n.Ptr())
	return MakeGridCell(result)
}

func (n *NSGridCell) Column() GridColumn {
	result := C.C_NSGridCell_Column(n.Ptr())
	return MakeGridColumn(result)
}

func (n *NSGridCell) Row() GridRow {
	result := C.C_NSGridCell_Row(n.Ptr())
	return MakeGridRow(result)
}

func (n *NSGridCell) ContentView() View {
	result := C.C_NSGridCell_ContentView(n.Ptr())
	return MakeView(result)
}

func (n *NSGridCell) SetContentView(value View) {
	C.C_NSGridCell_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func GridCell_EmptyContentView() View {
	result := C.C_NSGridCell_GridCell_EmptyContentView()
	return MakeView(result)
}

func (n *NSGridCell) CustomPlacementConstraints() []LayoutConstraint {
	result := C.C_NSGridCell_CustomPlacementConstraints(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]LayoutConstraint, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeLayoutConstraint(r)
	}
	return goResult
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
	result := C.C_NSGridCell_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result))
}

func (n *NSGridCell) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridCell_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n *NSGridCell) XPlacement() GridCellPlacement {
	result := C.C_NSGridCell_XPlacement(n.Ptr())
	return GridCellPlacement(int(result))
}

func (n *NSGridCell) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetXPlacement(n.Ptr(), C.int(int(value)))
}

func (n *NSGridCell) YPlacement() GridCellPlacement {
	result := C.C_NSGridCell_YPlacement(n.Ptr())
	return GridCellPlacement(int(result))
}

func (n *NSGridCell) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetYPlacement(n.Ptr(), C.int(int(value)))
}
