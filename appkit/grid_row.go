package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "grid_row.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GridRow interface {
	objc.Object
	CellAtIndex(index int) GridCell
	MergeCellsInRange(_range foundation.Range)
	NumberOfCells() int
	IsHidden() bool
	SetHidden(value bool)
	TopPadding() coregraphics.Float
	SetTopPadding(value coregraphics.Float)
	BottomPadding() coregraphics.Float
	SetBottomPadding(value coregraphics.Float)
	Height() coregraphics.Float
	SetHeight(value coregraphics.Float)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	GridView() GridView
}

type NSGridRow struct {
	objc.NSObject
}

func MakeGridRow(ptr unsafe.Pointer) *NSGridRow {
	if ptr == nil {
		return nil
	}
	return &NSGridRow{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocGridRow() *NSGridRow {
	return MakeGridRow(C.C_GridRow_Alloc())
}

func (n *NSGridRow) Init() GridRow {
	result_ := C.C_NSGridRow_Init(n.Ptr())
	return MakeGridRow(result_)
}

func (n *NSGridRow) CellAtIndex(index int) GridCell {
	result_ := C.C_NSGridRow_CellAtIndex(n.Ptr(), C.int(index))
	return MakeGridCell(result_)
}

func (n *NSGridRow) MergeCellsInRange(_range foundation.Range) {
	C.C_NSGridRow_MergeCellsInRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n *NSGridRow) NumberOfCells() int {
	result_ := C.C_NSGridRow_NumberOfCells(n.Ptr())
	return int(result_)
}

func (n *NSGridRow) IsHidden() bool {
	result_ := C.C_NSGridRow_IsHidden(n.Ptr())
	return bool(result_)
}

func (n *NSGridRow) SetHidden(value bool) {
	C.C_NSGridRow_SetHidden(n.Ptr(), C.bool(value))
}

func (n *NSGridRow) TopPadding() coregraphics.Float {
	result_ := C.C_NSGridRow_TopPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSGridRow) SetTopPadding(value coregraphics.Float) {
	C.C_NSGridRow_SetTopPadding(n.Ptr(), C.double(float64(value)))
}

func (n *NSGridRow) BottomPadding() coregraphics.Float {
	result_ := C.C_NSGridRow_BottomPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSGridRow) SetBottomPadding(value coregraphics.Float) {
	C.C_NSGridRow_SetBottomPadding(n.Ptr(), C.double(float64(value)))
}

func (n *NSGridRow) Height() coregraphics.Float {
	result_ := C.C_NSGridRow_Height(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSGridRow) SetHeight(value coregraphics.Float) {
	C.C_NSGridRow_SetHeight(n.Ptr(), C.double(float64(value)))
}

func (n *NSGridRow) RowAlignment() GridRowAlignment {
	result_ := C.C_NSGridRow_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result_))
}

func (n *NSGridRow) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridRow_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n *NSGridRow) YPlacement() GridCellPlacement {
	result_ := C.C_NSGridRow_YPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n *NSGridRow) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridRow_SetYPlacement(n.Ptr(), C.int(int(value)))
}

func (n *NSGridRow) GridView() GridView {
	result_ := C.C_NSGridRow_GridView(n.Ptr())
	return MakeGridView(result_)
}
