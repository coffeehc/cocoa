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
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(customPlacementConstraints []LayoutConstraint)
	RowAlignment() GridRowAlignment
	SetRowAlignment(rowAlignment GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(xPlacement GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(yPlacement GridCellPlacement)
}

var _ GridCell = (*NSGridCell)(nil)

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

func (g *NSGridCell) Column() GridColumn {
	return MakeGridColumn(C.GridCell_Column(g.Ptr()))
}

func (g *NSGridCell) Row() GridRow {
	return MakeGridRow(C.GridCell_Row(g.Ptr()))
}

func (g *NSGridCell) ContentView() View {
	return MakeView(C.GridCell_ContentView(g.Ptr()))
}

func (g *NSGridCell) CustomPlacementConstraints() []LayoutConstraint {
	var cArray C.Array = C.GridCell_CustomPlacementConstraints(g.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]LayoutConstraint, len(result))
	for idx, r := range result {
		goArray[idx] = MakeLayoutConstraint(r)
	}
	return goArray
}

func (g *NSGridCell) SetCustomPlacementConstraints(customPlacementConstraints []LayoutConstraint) {
	c_customPlacementConstraintsData := make([]unsafe.Pointer, len(customPlacementConstraints))
	for idx, v := range customPlacementConstraints {
		c_customPlacementConstraintsData[idx] = v.Ptr()
	}
	c_customPlacementConstraints := C.Array{data: unsafe.Pointer(&c_customPlacementConstraintsData[0]), len: C.int(len(customPlacementConstraints))}
	C.GridCell_SetCustomPlacementConstraints(g.Ptr(), c_customPlacementConstraints)
}

func (g *NSGridCell) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridCell_RowAlignment(g.Ptr()))
}

func (g *NSGridCell) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridCell_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridCell) XPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridCell_XPlacement(g.Ptr()))
}

func (g *NSGridCell) SetXPlacement(xPlacement GridCellPlacement) {
	C.GridCell_SetXPlacement(g.Ptr(), C.long(xPlacement))
}

func (g *NSGridCell) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridCell_YPlacement(g.Ptr()))
}

func (g *NSGridCell) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridCell_SetYPlacement(g.Ptr(), C.long(yPlacement))
}
