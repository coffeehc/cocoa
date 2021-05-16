package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "grid_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GridView interface {
	View
	IndexOfColumn(column GridColumn) int
	RowAtIndex(index int) GridRow
	ColumnAtIndex(index int) GridColumn
	IndexOfRow(row GridRow) int
	AddRowWithViews(views []View) GridRow
	InsertRowAtIndex_WithViews(index int, views []View) GridRow
	RemoveRowAtIndex(index int)
	MoveRowAtIndex_ToIndex(fromIndex int, toIndex int)
	AddColumnWithViews(views []View) GridColumn
	InsertColumnAtIndex_WithViews(index int, views []View) GridColumn
	RemoveColumnAtIndex(index int)
	MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int)
	CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell
	CellForView(view View) GridCell
	MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range)
	NumberOfRows() int
	NumberOfColumns() int
	ColumnSpacing() coregraphics.Float
	SetColumnSpacing(value coregraphics.Float)
	RowSpacing() coregraphics.Float
	SetRowSpacing(value coregraphics.Float)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

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

func AllocGridView() *NSGridView {
	return MakeGridView(C.C_GridView_Alloc())
}

func (n *NSGridView) InitWithFrame(frameRect foundation.Rect) GridView {
	result := C.C_NSGridView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeGridView(result)
}

func (n *NSGridView) InitWithCoder(coder foundation.Coder) GridView {
	result := C.C_NSGridView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeGridView(result)
}

func (n *NSGridView) Init() GridView {
	result := C.C_NSGridView_Init(n.Ptr())
	return MakeGridView(result)
}

func GridViewWithNumberOfColumns_Rows(columnCount int, rowCount int) GridView {
	result := C.C_NSGridView_GridViewWithNumberOfColumns_Rows(C.int(columnCount), C.int(rowCount))
	return MakeGridView(result)
}

func (n *NSGridView) IndexOfColumn(column GridColumn) int {
	result := C.C_NSGridView_IndexOfColumn(n.Ptr(), objc.ExtractPtr(column))
	return int(result)
}

func (n *NSGridView) RowAtIndex(index int) GridRow {
	result := C.C_NSGridView_RowAtIndex(n.Ptr(), C.int(index))
	return MakeGridRow(result)
}

func (n *NSGridView) ColumnAtIndex(index int) GridColumn {
	result := C.C_NSGridView_ColumnAtIndex(n.Ptr(), C.int(index))
	return MakeGridColumn(result)
}

func (n *NSGridView) IndexOfRow(row GridRow) int {
	result := C.C_NSGridView_IndexOfRow(n.Ptr(), objc.ExtractPtr(row))
	return int(result)
}

func (n *NSGridView) AddRowWithViews(views []View) GridRow {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	result := C.C_NSGridView_AddRowWithViews(n.Ptr(), cViews)
	return MakeGridRow(result)
}

func (n *NSGridView) InsertRowAtIndex_WithViews(index int, views []View) GridRow {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	result := C.C_NSGridView_InsertRowAtIndex_WithViews(n.Ptr(), C.int(index), cViews)
	return MakeGridRow(result)
}

func (n *NSGridView) RemoveRowAtIndex(index int) {
	C.C_NSGridView_RemoveRowAtIndex(n.Ptr(), C.int(index))
}

func (n *NSGridView) MoveRowAtIndex_ToIndex(fromIndex int, toIndex int) {
	C.C_NSGridView_MoveRowAtIndex_ToIndex(n.Ptr(), C.int(fromIndex), C.int(toIndex))
}

func (n *NSGridView) AddColumnWithViews(views []View) GridColumn {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	result := C.C_NSGridView_AddColumnWithViews(n.Ptr(), cViews)
	return MakeGridColumn(result)
}

func (n *NSGridView) InsertColumnAtIndex_WithViews(index int, views []View) GridColumn {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = objc.ExtractPtr(v)
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	result := C.C_NSGridView_InsertColumnAtIndex_WithViews(n.Ptr(), C.int(index), cViews)
	return MakeGridColumn(result)
}

func (n *NSGridView) RemoveColumnAtIndex(index int) {
	C.C_NSGridView_RemoveColumnAtIndex(n.Ptr(), C.int(index))
}

func (n *NSGridView) MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int) {
	C.C_NSGridView_MoveColumnAtIndex_ToIndex(n.Ptr(), C.int(fromIndex), C.int(toIndex))
}

func (n *NSGridView) CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell {
	result := C.C_NSGridView_CellAtColumnIndex_RowIndex(n.Ptr(), C.int(columnIndex), C.int(rowIndex))
	return MakeGridCell(result)
}

func (n *NSGridView) CellForView(view View) GridCell {
	result := C.C_NSGridView_CellForView(n.Ptr(), objc.ExtractPtr(view))
	return MakeGridCell(result)
}

func (n *NSGridView) MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range) {
	C.C_NSGridView_MergeCellsInHorizontalRange_VerticalRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(hRange)), *(*C.NSRange)(foundation.ToNSRangePointer(vRange)))
}

func (n *NSGridView) NumberOfRows() int {
	result := C.C_NSGridView_NumberOfRows(n.Ptr())
	return int(result)
}

func (n *NSGridView) NumberOfColumns() int {
	result := C.C_NSGridView_NumberOfColumns(n.Ptr())
	return int(result)
}

func (n *NSGridView) ColumnSpacing() coregraphics.Float {
	result := C.C_NSGridView_ColumnSpacing(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSGridView) SetColumnSpacing(value coregraphics.Float) {
	C.C_NSGridView_SetColumnSpacing(n.Ptr(), C.double(float64(value)))
}

func (n *NSGridView) RowSpacing() coregraphics.Float {
	result := C.C_NSGridView_RowSpacing(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSGridView) SetRowSpacing(value coregraphics.Float) {
	C.C_NSGridView_SetRowSpacing(n.Ptr(), C.double(float64(value)))
}

func (n *NSGridView) RowAlignment() GridRowAlignment {
	result := C.C_NSGridView_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result))
}

func (n *NSGridView) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridView_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n *NSGridView) XPlacement() GridCellPlacement {
	result := C.C_NSGridView_XPlacement(n.Ptr())
	return GridCellPlacement(int(result))
}

func (n *NSGridView) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridView_SetXPlacement(n.Ptr(), C.int(int(value)))
}

func (n *NSGridView) YPlacement() GridCellPlacement {
	result := C.C_NSGridView_YPlacement(n.Ptr())
	return GridCellPlacement(int(result))
}

func (n *NSGridView) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridView_SetYPlacement(n.Ptr(), C.int(int(value)))
}
