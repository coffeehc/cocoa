package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "grid_row.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GridRow interface {
	objc.Object
	NumberOfCells() int
	IsHidden() bool
	SetHidden(hidden bool)
	TopPadding() float64
	SetTopPadding(topPadding float64)
	BottomPadding() float64
	SetBottomPadding(bottomPadding float64)
	Height() float64
	SetHeight(height float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(rowAlignment GridRowAlignment)
	YPlacement() GridCellPlacement
	SetYPlacement(yPlacement GridCellPlacement)
	GridView() GridView
	CellAtIndex(index int) GridCell
	MergeCellsInRange(r foundation.Range)
}

var _ GridRow = (*NSGridRow)(nil)

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

func (g *NSGridRow) NumberOfCells() int {
	return int(C.GridRow_NumberOfCells(g.Ptr()))
}

func (g *NSGridRow) IsHidden() bool {
	return bool(C.GridRow_IsHidden(g.Ptr()))
}

func (g *NSGridRow) SetHidden(hidden bool) {
	C.GridRow_SetHidden(g.Ptr(), C.bool(hidden))
}

func (g *NSGridRow) TopPadding() float64 {
	return float64(C.GridRow_TopPadding(g.Ptr()))
}

func (g *NSGridRow) SetTopPadding(topPadding float64) {
	C.GridRow_SetTopPadding(g.Ptr(), C.double(topPadding))
}

func (g *NSGridRow) BottomPadding() float64 {
	return float64(C.GridRow_BottomPadding(g.Ptr()))
}

func (g *NSGridRow) SetBottomPadding(bottomPadding float64) {
	C.GridRow_SetBottomPadding(g.Ptr(), C.double(bottomPadding))
}

func (g *NSGridRow) Height() float64 {
	return float64(C.GridRow_Height(g.Ptr()))
}

func (g *NSGridRow) SetHeight(height float64) {
	C.GridRow_SetHeight(g.Ptr(), C.double(height))
}

func (g *NSGridRow) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridRow_RowAlignment(g.Ptr()))
}

func (g *NSGridRow) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridRow_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridRow) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridRow_YPlacement(g.Ptr()))
}

func (g *NSGridRow) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridRow_SetYPlacement(g.Ptr(), C.long(yPlacement))
}

func (g *NSGridRow) GridView() GridView {
	return MakeGridView(C.GridRow_GridView(g.Ptr()))
}

func (g *NSGridRow) CellAtIndex(index int) GridCell {
	return MakeGridCell(C.GridRow_CellAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridRow) MergeCellsInRange(r foundation.Range) {
	C.GridRow_MergeCellsInRange(g.Ptr(), *(*C.NSRange)(r.ToNSRangePointer()))
}
