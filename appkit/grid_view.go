package appkit

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

func MakeGridView(ptr unsafe.Pointer) NSGridView {
	return NSGridView{
		NSView: MakeView(ptr),
	}
}

func GridViewWithNumberOfColumns_Rows(columnCount int, rowCount int) NSGridView {
	result_ := C.C_NSGridView_GridViewWithNumberOfColumns_Rows(C.int(columnCount), C.int(rowCount))
	return MakeGridView(result_)
}

func (n NSGridView) InitWithFrame(frameRect foundation.Rect) NSGridView {
	result_ := C.C_NSGridView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeGridView(result_)
}

func (n NSGridView) InitWithCoder(coder foundation.Coder) NSGridView {
	result_ := C.C_NSGridView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeGridView(result_)
}

func (n NSGridView) Init() NSGridView {
	result_ := C.C_NSGridView_Init(n.Ptr())
	return MakeGridView(result_)
}

func AllocGridView() NSGridView {
	result_ := C.C_NSGridView_AllocGridView()
	return MakeGridView(result_)
}

func NewGridView() NSGridView {
	result_ := C.C_NSGridView_NewGridView()
	return MakeGridView(result_)
}

func (n NSGridView) Autorelease() NSGridView {
	result_ := C.C_NSGridView_Autorelease(n.Ptr())
	return MakeGridView(result_)
}

func (n NSGridView) Retain() NSGridView {
	result_ := C.C_NSGridView_Retain(n.Ptr())
	return MakeGridView(result_)
}

func (n NSGridView) IndexOfColumn(column GridColumn) int {
	result_ := C.C_NSGridView_IndexOfColumn(n.Ptr(), objc.ExtractPtr(column))
	return int(result_)
}

func (n NSGridView) RowAtIndex(index int) GridRow {
	result_ := C.C_NSGridView_RowAtIndex(n.Ptr(), C.int(index))
	return MakeGridRow(result_)
}

func (n NSGridView) ColumnAtIndex(index int) GridColumn {
	result_ := C.C_NSGridView_ColumnAtIndex(n.Ptr(), C.int(index))
	return MakeGridColumn(result_)
}

func (n NSGridView) IndexOfRow(row GridRow) int {
	result_ := C.C_NSGridView_IndexOfRow(n.Ptr(), objc.ExtractPtr(row))
	return int(result_)
}

func (n NSGridView) AddRowWithViews(views []View) GridRow {
	var cViews C.Array
	if len(views) > 0 {
		cViewsData := make([]unsafe.Pointer, len(views))
		for idx, v := range views {
			cViewsData[idx] = objc.ExtractPtr(v)
		}
		cViews.data = unsafe.Pointer(&cViewsData[0])
		cViews.len = C.int(len(views))
	}
	result_ := C.C_NSGridView_AddRowWithViews(n.Ptr(), cViews)
	return MakeGridRow(result_)
}

func (n NSGridView) InsertRowAtIndex_WithViews(index int, views []View) GridRow {
	var cViews C.Array
	if len(views) > 0 {
		cViewsData := make([]unsafe.Pointer, len(views))
		for idx, v := range views {
			cViewsData[idx] = objc.ExtractPtr(v)
		}
		cViews.data = unsafe.Pointer(&cViewsData[0])
		cViews.len = C.int(len(views))
	}
	result_ := C.C_NSGridView_InsertRowAtIndex_WithViews(n.Ptr(), C.int(index), cViews)
	return MakeGridRow(result_)
}

func (n NSGridView) RemoveRowAtIndex(index int) {
	C.C_NSGridView_RemoveRowAtIndex(n.Ptr(), C.int(index))
}

func (n NSGridView) MoveRowAtIndex_ToIndex(fromIndex int, toIndex int) {
	C.C_NSGridView_MoveRowAtIndex_ToIndex(n.Ptr(), C.int(fromIndex), C.int(toIndex))
}

func (n NSGridView) AddColumnWithViews(views []View) GridColumn {
	var cViews C.Array
	if len(views) > 0 {
		cViewsData := make([]unsafe.Pointer, len(views))
		for idx, v := range views {
			cViewsData[idx] = objc.ExtractPtr(v)
		}
		cViews.data = unsafe.Pointer(&cViewsData[0])
		cViews.len = C.int(len(views))
	}
	result_ := C.C_NSGridView_AddColumnWithViews(n.Ptr(), cViews)
	return MakeGridColumn(result_)
}

func (n NSGridView) InsertColumnAtIndex_WithViews(index int, views []View) GridColumn {
	var cViews C.Array
	if len(views) > 0 {
		cViewsData := make([]unsafe.Pointer, len(views))
		for idx, v := range views {
			cViewsData[idx] = objc.ExtractPtr(v)
		}
		cViews.data = unsafe.Pointer(&cViewsData[0])
		cViews.len = C.int(len(views))
	}
	result_ := C.C_NSGridView_InsertColumnAtIndex_WithViews(n.Ptr(), C.int(index), cViews)
	return MakeGridColumn(result_)
}

func (n NSGridView) RemoveColumnAtIndex(index int) {
	C.C_NSGridView_RemoveColumnAtIndex(n.Ptr(), C.int(index))
}

func (n NSGridView) MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int) {
	C.C_NSGridView_MoveColumnAtIndex_ToIndex(n.Ptr(), C.int(fromIndex), C.int(toIndex))
}

func (n NSGridView) CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell {
	result_ := C.C_NSGridView_CellAtColumnIndex_RowIndex(n.Ptr(), C.int(columnIndex), C.int(rowIndex))
	return MakeGridCell(result_)
}

func (n NSGridView) CellForView(view View) GridCell {
	result_ := C.C_NSGridView_CellForView(n.Ptr(), objc.ExtractPtr(view))
	return MakeGridCell(result_)
}

func (n NSGridView) MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range) {
	C.C_NSGridView_MergeCellsInHorizontalRange_VerticalRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(hRange)), *(*C.NSRange)(foundation.ToNSRangePointer(vRange)))
}

func (n NSGridView) NumberOfRows() int {
	result_ := C.C_NSGridView_NumberOfRows(n.Ptr())
	return int(result_)
}

func (n NSGridView) NumberOfColumns() int {
	result_ := C.C_NSGridView_NumberOfColumns(n.Ptr())
	return int(result_)
}

func (n NSGridView) ColumnSpacing() coregraphics.Float {
	result_ := C.C_NSGridView_ColumnSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridView) SetColumnSpacing(value coregraphics.Float) {
	C.C_NSGridView_SetColumnSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSGridView) RowSpacing() coregraphics.Float {
	result_ := C.C_NSGridView_RowSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridView) SetRowSpacing(value coregraphics.Float) {
	C.C_NSGridView_SetRowSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSGridView) RowAlignment() GridRowAlignment {
	result_ := C.C_NSGridView_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result_))
}

func (n NSGridView) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridView_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSGridView) XPlacement() GridCellPlacement {
	result_ := C.C_NSGridView_XPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridView) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridView_SetXPlacement(n.Ptr(), C.int(int(value)))
}

func (n NSGridView) YPlacement() GridCellPlacement {
	result_ := C.C_NSGridView_YPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridView) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridView_SetYPlacement(n.Ptr(), C.int(int(value)))
}

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

func MakeGridCell(ptr unsafe.Pointer) NSGridCell {
	return NSGridCell{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocGridCell() NSGridCell {
	result_ := C.C_NSGridCell_AllocGridCell()
	return MakeGridCell(result_)
}

func (n NSGridCell) Init() NSGridCell {
	result_ := C.C_NSGridCell_Init(n.Ptr())
	return MakeGridCell(result_)
}

func NewGridCell() NSGridCell {
	result_ := C.C_NSGridCell_NewGridCell()
	return MakeGridCell(result_)
}

func (n NSGridCell) Autorelease() NSGridCell {
	result_ := C.C_NSGridCell_Autorelease(n.Ptr())
	return MakeGridCell(result_)
}

func (n NSGridCell) Retain() NSGridCell {
	result_ := C.C_NSGridCell_Retain(n.Ptr())
	return MakeGridCell(result_)
}

func (n NSGridCell) Column() GridColumn {
	result_ := C.C_NSGridCell_Column(n.Ptr())
	return MakeGridColumn(result_)
}

func (n NSGridCell) Row() GridRow {
	result_ := C.C_NSGridCell_Row(n.Ptr())
	return MakeGridRow(result_)
}

func (n NSGridCell) ContentView() View {
	result_ := C.C_NSGridCell_ContentView(n.Ptr())
	return MakeView(result_)
}

func (n NSGridCell) SetContentView(value View) {
	C.C_NSGridCell_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func GridCell_EmptyContentView() View {
	result_ := C.C_NSGridCell_GridCell_EmptyContentView()
	return MakeView(result_)
}

func (n NSGridCell) CustomPlacementConstraints() []LayoutConstraint {
	result_ := C.C_NSGridCell_CustomPlacementConstraints(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func (n NSGridCell) SetCustomPlacementConstraints(value []LayoutConstraint) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSGridCell_SetCustomPlacementConstraints(n.Ptr(), cValue)
}

func (n NSGridCell) RowAlignment() GridRowAlignment {
	result_ := C.C_NSGridCell_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result_))
}

func (n NSGridCell) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridCell_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSGridCell) XPlacement() GridCellPlacement {
	result_ := C.C_NSGridCell_XPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridCell) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetXPlacement(n.Ptr(), C.int(int(value)))
}

func (n NSGridCell) YPlacement() GridCellPlacement {
	result_ := C.C_NSGridCell_YPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridCell) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridCell_SetYPlacement(n.Ptr(), C.int(int(value)))
}

type GridColumn interface {
	objc.Object
	CellAtIndex(index int) GridCell
	MergeCellsInRange(_range foundation.Range)
	GridView() GridView
	IsHidden() bool
	SetHidden(value bool)
	LeadingPadding() coregraphics.Float
	SetLeadingPadding(value coregraphics.Float)
	NumberOfCells() int
	TrailingPadding() coregraphics.Float
	SetTrailingPadding(value coregraphics.Float)
	Width() coregraphics.Float
	SetWidth(value coregraphics.Float)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
}

type NSGridColumn struct {
	objc.NSObject
}

func MakeGridColumn(ptr unsafe.Pointer) NSGridColumn {
	return NSGridColumn{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocGridColumn() NSGridColumn {
	result_ := C.C_NSGridColumn_AllocGridColumn()
	return MakeGridColumn(result_)
}

func (n NSGridColumn) Init() NSGridColumn {
	result_ := C.C_NSGridColumn_Init(n.Ptr())
	return MakeGridColumn(result_)
}

func NewGridColumn() NSGridColumn {
	result_ := C.C_NSGridColumn_NewGridColumn()
	return MakeGridColumn(result_)
}

func (n NSGridColumn) Autorelease() NSGridColumn {
	result_ := C.C_NSGridColumn_Autorelease(n.Ptr())
	return MakeGridColumn(result_)
}

func (n NSGridColumn) Retain() NSGridColumn {
	result_ := C.C_NSGridColumn_Retain(n.Ptr())
	return MakeGridColumn(result_)
}

func (n NSGridColumn) CellAtIndex(index int) GridCell {
	result_ := C.C_NSGridColumn_CellAtIndex(n.Ptr(), C.int(index))
	return MakeGridCell(result_)
}

func (n NSGridColumn) MergeCellsInRange(_range foundation.Range) {
	C.C_NSGridColumn_MergeCellsInRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSGridColumn) GridView() GridView {
	result_ := C.C_NSGridColumn_GridView(n.Ptr())
	return MakeGridView(result_)
}

func (n NSGridColumn) IsHidden() bool {
	result_ := C.C_NSGridColumn_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSGridColumn) SetHidden(value bool) {
	C.C_NSGridColumn_SetHidden(n.Ptr(), C.bool(value))
}

func (n NSGridColumn) LeadingPadding() coregraphics.Float {
	result_ := C.C_NSGridColumn_LeadingPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridColumn) SetLeadingPadding(value coregraphics.Float) {
	C.C_NSGridColumn_SetLeadingPadding(n.Ptr(), C.double(float64(value)))
}

func (n NSGridColumn) NumberOfCells() int {
	result_ := C.C_NSGridColumn_NumberOfCells(n.Ptr())
	return int(result_)
}

func (n NSGridColumn) TrailingPadding() coregraphics.Float {
	result_ := C.C_NSGridColumn_TrailingPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridColumn) SetTrailingPadding(value coregraphics.Float) {
	C.C_NSGridColumn_SetTrailingPadding(n.Ptr(), C.double(float64(value)))
}

func (n NSGridColumn) Width() coregraphics.Float {
	result_ := C.C_NSGridColumn_Width(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridColumn) SetWidth(value coregraphics.Float) {
	C.C_NSGridColumn_SetWidth(n.Ptr(), C.double(float64(value)))
}

func (n NSGridColumn) XPlacement() GridCellPlacement {
	result_ := C.C_NSGridColumn_XPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridColumn) SetXPlacement(value GridCellPlacement) {
	C.C_NSGridColumn_SetXPlacement(n.Ptr(), C.int(int(value)))
}

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

func MakeGridRow(ptr unsafe.Pointer) NSGridRow {
	return NSGridRow{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocGridRow() NSGridRow {
	result_ := C.C_NSGridRow_AllocGridRow()
	return MakeGridRow(result_)
}

func (n NSGridRow) Init() NSGridRow {
	result_ := C.C_NSGridRow_Init(n.Ptr())
	return MakeGridRow(result_)
}

func NewGridRow() NSGridRow {
	result_ := C.C_NSGridRow_NewGridRow()
	return MakeGridRow(result_)
}

func (n NSGridRow) Autorelease() NSGridRow {
	result_ := C.C_NSGridRow_Autorelease(n.Ptr())
	return MakeGridRow(result_)
}

func (n NSGridRow) Retain() NSGridRow {
	result_ := C.C_NSGridRow_Retain(n.Ptr())
	return MakeGridRow(result_)
}

func (n NSGridRow) CellAtIndex(index int) GridCell {
	result_ := C.C_NSGridRow_CellAtIndex(n.Ptr(), C.int(index))
	return MakeGridCell(result_)
}

func (n NSGridRow) MergeCellsInRange(_range foundation.Range) {
	C.C_NSGridRow_MergeCellsInRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSGridRow) NumberOfCells() int {
	result_ := C.C_NSGridRow_NumberOfCells(n.Ptr())
	return int(result_)
}

func (n NSGridRow) IsHidden() bool {
	result_ := C.C_NSGridRow_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSGridRow) SetHidden(value bool) {
	C.C_NSGridRow_SetHidden(n.Ptr(), C.bool(value))
}

func (n NSGridRow) TopPadding() coregraphics.Float {
	result_ := C.C_NSGridRow_TopPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridRow) SetTopPadding(value coregraphics.Float) {
	C.C_NSGridRow_SetTopPadding(n.Ptr(), C.double(float64(value)))
}

func (n NSGridRow) BottomPadding() coregraphics.Float {
	result_ := C.C_NSGridRow_BottomPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridRow) SetBottomPadding(value coregraphics.Float) {
	C.C_NSGridRow_SetBottomPadding(n.Ptr(), C.double(float64(value)))
}

func (n NSGridRow) Height() coregraphics.Float {
	result_ := C.C_NSGridRow_Height(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSGridRow) SetHeight(value coregraphics.Float) {
	C.C_NSGridRow_SetHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSGridRow) RowAlignment() GridRowAlignment {
	result_ := C.C_NSGridRow_RowAlignment(n.Ptr())
	return GridRowAlignment(int(result_))
}

func (n NSGridRow) SetRowAlignment(value GridRowAlignment) {
	C.C_NSGridRow_SetRowAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSGridRow) YPlacement() GridCellPlacement {
	result_ := C.C_NSGridRow_YPlacement(n.Ptr())
	return GridCellPlacement(int(result_))
}

func (n NSGridRow) SetYPlacement(value GridCellPlacement) {
	C.C_NSGridRow_SetYPlacement(n.Ptr(), C.int(int(value)))
}

func (n NSGridRow) GridView() GridView {
	result_ := C.C_NSGridRow_GridView(n.Ptr())
	return MakeGridView(result_)
}
