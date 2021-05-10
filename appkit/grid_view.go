package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "grid_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type GridView interface {
	View
	NumberOfRows() int
	NumberOfColumns() int
	ColumnSpacing() float64
	SetColumnSpacing(columnSpacing float64)
	RowSpacing() float64
	SetRowSpacing(rowSpacing float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(rowAlignment GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(xPlacement GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(yPlacement GridCellPlacement)
	IndexOfColumn(column GridColumn) int
	IndexOfRow(row GridRow) int
	ColumnAtIndex(index int) GridColumn
	RowAtIndex(index int) GridRow
	AddRowWithViews(views []View) GridRow
	InsertRowAtIndex(index int, views []View) GridRow
	RemoveRowAtIndex(index int)
	MoveRowAtIndex(fromIndex int, toIndex int)
	AddColumnWithViews(views []View) GridColumn
	InsertColumnAtIndex(index int, views []View) GridColumn
	RemoveColumnAtIndex(index int)
	MoveColumnAtIndex(fromIndex int, toIndex int)
	CellAt(columnIndex int, rowIndex int) GridCell
	CellForView(view View) GridCell
	MergeCellsInHorizontalRange(hRange foundation.Range, vRange foundation.Range)
}

var _ GridView = (*NSGridView)(nil)

type NSGridView struct {
	NSView
}

func MakeGridView(ptr unsafe.Pointer) *NSGridView {
	if ptr == nil {
		return nil
	}
	return &NSGridView{
		NSView: *MakeView(ptr),
	}
}

func (g *NSGridView) NumberOfRows() int {
	return int(C.GridView_NumberOfRows(g.Ptr()))
}

func (g *NSGridView) NumberOfColumns() int {
	return int(C.GridView_NumberOfColumns(g.Ptr()))
}

func (g *NSGridView) ColumnSpacing() float64 {
	return float64(C.GridView_ColumnSpacing(g.Ptr()))
}

func (g *NSGridView) SetColumnSpacing(columnSpacing float64) {
	C.GridView_SetColumnSpacing(g.Ptr(), C.double(columnSpacing))
}

func (g *NSGridView) RowSpacing() float64 {
	return float64(C.GridView_RowSpacing(g.Ptr()))
}

func (g *NSGridView) SetRowSpacing(rowSpacing float64) {
	C.GridView_SetRowSpacing(g.Ptr(), C.double(rowSpacing))
}

func (g *NSGridView) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridView_RowAlignment(g.Ptr()))
}

func (g *NSGridView) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridView_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridView) XPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridView_XPlacement(g.Ptr()))
}

func (g *NSGridView) SetXPlacement(xPlacement GridCellPlacement) {
	C.GridView_SetXPlacement(g.Ptr(), C.long(xPlacement))
}

func (g *NSGridView) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridView_YPlacement(g.Ptr()))
}

func (g *NSGridView) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridView_SetYPlacement(g.Ptr(), C.long(yPlacement))
}

func NewGridView(frameRect foundation.Rect) GridView {
	return MakeGridView(C.GridView_NewGridView(toNSRect(frameRect)))
}

func GridViewWithNumberOfColumns(columnCount int, rowCount int) GridView {
	return MakeGridView(C.GridView_GridViewWithNumberOfColumns(C.long(columnCount), C.long(rowCount)))
}

func (g *NSGridView) IndexOfColumn(column GridColumn) int {
	return int(C.GridView_IndexOfColumn(g.Ptr(), toPointer(column)))
}

func (g *NSGridView) IndexOfRow(row GridRow) int {
	return int(C.GridView_IndexOfRow(g.Ptr(), toPointer(row)))
}

func (g *NSGridView) ColumnAtIndex(index int) GridColumn {
	return MakeGridColumn(C.GridView_ColumnAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridView) RowAtIndex(index int) GridRow {
	return MakeGridRow(C.GridView_RowAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridView) AddRowWithViews(views []View) GridRow {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	return MakeGridRow(C.GridView_AddRowWithViews(g.Ptr(), c_views))
}

func (g *NSGridView) InsertRowAtIndex(index int, views []View) GridRow {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	return MakeGridRow(C.GridView_InsertRowAtIndex(g.Ptr(), C.long(index), c_views))
}

func (g *NSGridView) RemoveRowAtIndex(index int) {
	C.GridView_RemoveRowAtIndex(g.Ptr(), C.long(index))
}

func (g *NSGridView) MoveRowAtIndex(fromIndex int, toIndex int) {
	C.GridView_MoveRowAtIndex(g.Ptr(), C.long(fromIndex), C.long(toIndex))
}

func (g *NSGridView) AddColumnWithViews(views []View) GridColumn {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	return MakeGridColumn(C.GridView_AddColumnWithViews(g.Ptr(), c_views))
}

func (g *NSGridView) InsertColumnAtIndex(index int, views []View) GridColumn {
	c_viewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		c_viewsData[idx] = v.Ptr()
	}
	c_views := C.Array{data: unsafe.Pointer(&c_viewsData[0]), len: C.int(len(views))}
	return MakeGridColumn(C.GridView_InsertColumnAtIndex(g.Ptr(), C.long(index), c_views))
}

func (g *NSGridView) RemoveColumnAtIndex(index int) {
	C.GridView_RemoveColumnAtIndex(g.Ptr(), C.long(index))
}

func (g *NSGridView) MoveColumnAtIndex(fromIndex int, toIndex int) {
	C.GridView_MoveColumnAtIndex(g.Ptr(), C.long(fromIndex), C.long(toIndex))
}

func (g *NSGridView) CellAt(columnIndex int, rowIndex int) GridCell {
	return MakeGridCell(C.GridView_CellAt(g.Ptr(), C.long(columnIndex), C.long(rowIndex)))
}

func (g *NSGridView) CellForView(view View) GridCell {
	return MakeGridCell(C.GridView_CellForView(g.Ptr(), toPointer(view)))
}

func (g *NSGridView) MergeCellsInHorizontalRange(hRange foundation.Range, vRange foundation.Range) {
	C.GridView_MergeCellsInHorizontalRange(g.Ptr(), *(*C.NSRange)(hRange.ToNSRangePointer()), *(*C.NSRange)(vRange.ToNSRangePointer()))
}
