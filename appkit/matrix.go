package appkit

// #include "matrix.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Matrix interface {
	Control
	AddColumn()
	AddColumnWithCells(newCells []Cell)
	AddRow()
	AddRowWithCells(newCells []Cell)
	CellFrameAtRow_Column(row int, col int) foundation.Rect
	InsertColumn(column int)
	InsertColumn_WithCells(column int, newCells []Cell)
	InsertRow(row int)
	InsertRow_WithCells(row int, newCells []Cell)
	MakeCellAtRow_Column(row int, col int) Cell
	PutCell_AtRow_Column(newCell Cell, row int, col int)
	RemoveColumn(col int)
	RemoveRow(row int)
	RenewRows_Columns(newRows int, newCols int)
	SortUsingSelector(comparator objc.Selector)
	SetState_AtRow_Column(value int, row int, col int)
	SetToolTip_ForCell(toolTipString string, cell Cell)
	ToolTipForCell(cell Cell) string
	SelectCellAtRow_Column(row int, col int)
	SelectCellWithTag(tag int) bool
	SelectAll(sender objc.Object)
	SetSelectionFrom_To_Anchor_Highlight(startPos int, endPos int, anchorPos int, lit bool)
	DeselectAllCells()
	DeselectSelectedCell()
	CellAtRow_Column(row int, col int) Cell
	CellWithTag(tag int) Cell
	SelectText(sender objc.Object)
	SelectTextAtRow_Column(row int, col int) Cell
	TextShouldBeginEditing(textObject Text) bool
	TextDidBeginEditing(notification foundation.Notification)
	TextDidChange(notification foundation.Notification)
	TextShouldEndEditing(textObject Text) bool
	TextDidEndEditing(notification foundation.Notification)
	SetValidateSize(flag bool)
	SizeToCells()
	SetScrollable(flag bool)
	ScrollCellToVisibleAtRow_Column(row int, col int)
	DrawCellAtRow_Column(row int, col int)
	HighlightCell_AtRow_Column(flag bool, row int, col int)
	SendAction() bool
	SendAction_To_ForAllCells(selector objc.Selector, object objc.Object, flag bool)
	SendDoubleAction()
	Mode() MatrixMode
	SetMode(value MatrixMode)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	IsSelectionByRect() bool
	SetSelectionByRect(value bool)
	Prototype() Cell
	SetPrototype(value Cell)
	CellSize() foundation.Size
	SetCellSize(value foundation.Size)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	NumberOfColumns() int
	NumberOfRows() int
	AutorecalculatesCellSize() bool
	SetAutorecalculatesCellSize(value bool)
	KeyCell() Cell
	SetKeyCell(value Cell)
	SelectedCell() Cell
	SelectedCells() []Cell
	SelectedColumn() int
	SelectedRow() int
	Cells() []Cell
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	CellBackgroundColor() Color
	SetCellBackgroundColor(value Color)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	DrawsCellBackground() bool
	SetDrawsCellBackground(value bool)
	TabKeyTraversesCells() bool
	SetTabKeyTraversesCells(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	AutosizesCells() bool
	SetAutosizesCells(value bool)
	IsAutoscroll() bool
	SetAutoscroll(value bool)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	MouseDownFlags() int
}

type NSMatrix struct {
	NSControl
}

func MakeMatrix(ptr unsafe.Pointer) NSMatrix {
	return NSMatrix{
		NSControl: MakeControl(ptr),
	}
}

func AllocMatrix() NSMatrix {
	return MakeMatrix(C.C_Matrix_Alloc())
}

func (n NSMatrix) InitWithFrame(frameRect foundation.Rect) Matrix {
	result_ := C.C_NSMatrix_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeMatrix(result_)
}

func (n NSMatrix) InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(frameRect foundation.Rect, mode MatrixMode, cell Cell, rowsHigh int, colsWide int) Matrix {
	result_ := C.C_NSMatrix_InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), C.uint(uint(mode)), objc.ExtractPtr(cell), C.int(rowsHigh), C.int(colsWide))
	return MakeMatrix(result_)
}

func (n NSMatrix) InitWithCoder(coder foundation.Coder) Matrix {
	result_ := C.C_NSMatrix_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeMatrix(result_)
}

func (n NSMatrix) Init() Matrix {
	result_ := C.C_NSMatrix_Init(n.Ptr())
	return MakeMatrix(result_)
}

func (n NSMatrix) AddColumn() {
	C.C_NSMatrix_AddColumn(n.Ptr())
}

func (n NSMatrix) AddColumnWithCells(newCells []Cell) {
	var cNewCells C.Array
	if len(newCells) > 0 {
		cNewCellsData := make([]unsafe.Pointer, len(newCells))
		for idx, v := range newCells {
			cNewCellsData[idx] = objc.ExtractPtr(v)
		}
		cNewCells.data = unsafe.Pointer(&cNewCellsData[0])
		cNewCells.len = C.int(len(newCells))
	}
	C.C_NSMatrix_AddColumnWithCells(n.Ptr(), cNewCells)
}

func (n NSMatrix) AddRow() {
	C.C_NSMatrix_AddRow(n.Ptr())
}

func (n NSMatrix) AddRowWithCells(newCells []Cell) {
	var cNewCells C.Array
	if len(newCells) > 0 {
		cNewCellsData := make([]unsafe.Pointer, len(newCells))
		for idx, v := range newCells {
			cNewCellsData[idx] = objc.ExtractPtr(v)
		}
		cNewCells.data = unsafe.Pointer(&cNewCellsData[0])
		cNewCells.len = C.int(len(newCells))
	}
	C.C_NSMatrix_AddRowWithCells(n.Ptr(), cNewCells)
}

func (n NSMatrix) CellFrameAtRow_Column(row int, col int) foundation.Rect {
	result_ := C.C_NSMatrix_CellFrameAtRow_Column(n.Ptr(), C.int(row), C.int(col))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSMatrix) InsertColumn(column int) {
	C.C_NSMatrix_InsertColumn(n.Ptr(), C.int(column))
}

func (n NSMatrix) InsertColumn_WithCells(column int, newCells []Cell) {
	var cNewCells C.Array
	if len(newCells) > 0 {
		cNewCellsData := make([]unsafe.Pointer, len(newCells))
		for idx, v := range newCells {
			cNewCellsData[idx] = objc.ExtractPtr(v)
		}
		cNewCells.data = unsafe.Pointer(&cNewCellsData[0])
		cNewCells.len = C.int(len(newCells))
	}
	C.C_NSMatrix_InsertColumn_WithCells(n.Ptr(), C.int(column), cNewCells)
}

func (n NSMatrix) InsertRow(row int) {
	C.C_NSMatrix_InsertRow(n.Ptr(), C.int(row))
}

func (n NSMatrix) InsertRow_WithCells(row int, newCells []Cell) {
	var cNewCells C.Array
	if len(newCells) > 0 {
		cNewCellsData := make([]unsafe.Pointer, len(newCells))
		for idx, v := range newCells {
			cNewCellsData[idx] = objc.ExtractPtr(v)
		}
		cNewCells.data = unsafe.Pointer(&cNewCellsData[0])
		cNewCells.len = C.int(len(newCells))
	}
	C.C_NSMatrix_InsertRow_WithCells(n.Ptr(), C.int(row), cNewCells)
}

func (n NSMatrix) MakeCellAtRow_Column(row int, col int) Cell {
	result_ := C.C_NSMatrix_MakeCellAtRow_Column(n.Ptr(), C.int(row), C.int(col))
	return MakeCell(result_)
}

func (n NSMatrix) PutCell_AtRow_Column(newCell Cell, row int, col int) {
	C.C_NSMatrix_PutCell_AtRow_Column(n.Ptr(), objc.ExtractPtr(newCell), C.int(row), C.int(col))
}

func (n NSMatrix) RemoveColumn(col int) {
	C.C_NSMatrix_RemoveColumn(n.Ptr(), C.int(col))
}

func (n NSMatrix) RemoveRow(row int) {
	C.C_NSMatrix_RemoveRow(n.Ptr(), C.int(row))
}

func (n NSMatrix) RenewRows_Columns(newRows int, newCols int) {
	C.C_NSMatrix_RenewRows_Columns(n.Ptr(), C.int(newRows), C.int(newCols))
}

func (n NSMatrix) SortUsingSelector(comparator objc.Selector) {
	C.C_NSMatrix_SortUsingSelector(n.Ptr(), unsafe.Pointer(comparator))
}

func (n NSMatrix) SetState_AtRow_Column(value int, row int, col int) {
	C.C_NSMatrix_SetState_AtRow_Column(n.Ptr(), C.int(value), C.int(row), C.int(col))
}

func (n NSMatrix) SetToolTip_ForCell(toolTipString string, cell Cell) {
	C.C_NSMatrix_SetToolTip_ForCell(n.Ptr(), foundation.NewString(toolTipString).Ptr(), objc.ExtractPtr(cell))
}

func (n NSMatrix) ToolTipForCell(cell Cell) string {
	result_ := C.C_NSMatrix_ToolTipForCell(n.Ptr(), objc.ExtractPtr(cell))
	return foundation.MakeString(result_).String()
}

func (n NSMatrix) SelectCellAtRow_Column(row int, col int) {
	C.C_NSMatrix_SelectCellAtRow_Column(n.Ptr(), C.int(row), C.int(col))
}

func (n NSMatrix) SelectCellWithTag(tag int) bool {
	result_ := C.C_NSMatrix_SelectCellWithTag(n.Ptr(), C.int(tag))
	return bool(result_)
}

func (n NSMatrix) SelectAll(sender objc.Object) {
	C.C_NSMatrix_SelectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSMatrix) SetSelectionFrom_To_Anchor_Highlight(startPos int, endPos int, anchorPos int, lit bool) {
	C.C_NSMatrix_SetSelectionFrom_To_Anchor_Highlight(n.Ptr(), C.int(startPos), C.int(endPos), C.int(anchorPos), C.bool(lit))
}

func (n NSMatrix) DeselectAllCells() {
	C.C_NSMatrix_DeselectAllCells(n.Ptr())
}

func (n NSMatrix) DeselectSelectedCell() {
	C.C_NSMatrix_DeselectSelectedCell(n.Ptr())
}

func (n NSMatrix) CellAtRow_Column(row int, col int) Cell {
	result_ := C.C_NSMatrix_CellAtRow_Column(n.Ptr(), C.int(row), C.int(col))
	return MakeCell(result_)
}

func (n NSMatrix) CellWithTag(tag int) Cell {
	result_ := C.C_NSMatrix_CellWithTag(n.Ptr(), C.int(tag))
	return MakeCell(result_)
}

func (n NSMatrix) SelectText(sender objc.Object) {
	C.C_NSMatrix_SelectText(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSMatrix) SelectTextAtRow_Column(row int, col int) Cell {
	result_ := C.C_NSMatrix_SelectTextAtRow_Column(n.Ptr(), C.int(row), C.int(col))
	return MakeCell(result_)
}

func (n NSMatrix) TextShouldBeginEditing(textObject Text) bool {
	result_ := C.C_NSMatrix_TextShouldBeginEditing(n.Ptr(), objc.ExtractPtr(textObject))
	return bool(result_)
}

func (n NSMatrix) TextDidBeginEditing(notification foundation.Notification) {
	C.C_NSMatrix_TextDidBeginEditing(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSMatrix) TextDidChange(notification foundation.Notification) {
	C.C_NSMatrix_TextDidChange(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSMatrix) TextShouldEndEditing(textObject Text) bool {
	result_ := C.C_NSMatrix_TextShouldEndEditing(n.Ptr(), objc.ExtractPtr(textObject))
	return bool(result_)
}

func (n NSMatrix) TextDidEndEditing(notification foundation.Notification) {
	C.C_NSMatrix_TextDidEndEditing(n.Ptr(), objc.ExtractPtr(notification))
}

func (n NSMatrix) SetValidateSize(flag bool) {
	C.C_NSMatrix_SetValidateSize(n.Ptr(), C.bool(flag))
}

func (n NSMatrix) SizeToCells() {
	C.C_NSMatrix_SizeToCells(n.Ptr())
}

func (n NSMatrix) SetScrollable(flag bool) {
	C.C_NSMatrix_SetScrollable(n.Ptr(), C.bool(flag))
}

func (n NSMatrix) ScrollCellToVisibleAtRow_Column(row int, col int) {
	C.C_NSMatrix_ScrollCellToVisibleAtRow_Column(n.Ptr(), C.int(row), C.int(col))
}

func (n NSMatrix) DrawCellAtRow_Column(row int, col int) {
	C.C_NSMatrix_DrawCellAtRow_Column(n.Ptr(), C.int(row), C.int(col))
}

func (n NSMatrix) HighlightCell_AtRow_Column(flag bool, row int, col int) {
	C.C_NSMatrix_HighlightCell_AtRow_Column(n.Ptr(), C.bool(flag), C.int(row), C.int(col))
}

func (n NSMatrix) SendAction() bool {
	result_ := C.C_NSMatrix_SendAction(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SendAction_To_ForAllCells(selector objc.Selector, object objc.Object, flag bool) {
	C.C_NSMatrix_SendAction_To_ForAllCells(n.Ptr(), unsafe.Pointer(selector), objc.ExtractPtr(object), C.bool(flag))
}

func (n NSMatrix) SendDoubleAction() {
	C.C_NSMatrix_SendDoubleAction(n.Ptr())
}

func (n NSMatrix) Mode() MatrixMode {
	result_ := C.C_NSMatrix_Mode(n.Ptr())
	return MatrixMode(uint(result_))
}

func (n NSMatrix) SetMode(value MatrixMode) {
	C.C_NSMatrix_SetMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSMatrix) AllowsEmptySelection() bool {
	result_ := C.C_NSMatrix_AllowsEmptySelection(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetAllowsEmptySelection(value bool) {
	C.C_NSMatrix_SetAllowsEmptySelection(n.Ptr(), C.bool(value))
}

func (n NSMatrix) IsSelectionByRect() bool {
	result_ := C.C_NSMatrix_IsSelectionByRect(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetSelectionByRect(value bool) {
	C.C_NSMatrix_SetSelectionByRect(n.Ptr(), C.bool(value))
}

func (n NSMatrix) Prototype() Cell {
	result_ := C.C_NSMatrix_Prototype(n.Ptr())
	return MakeCell(result_)
}

func (n NSMatrix) SetPrototype(value Cell) {
	C.C_NSMatrix_SetPrototype(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMatrix) CellSize() foundation.Size {
	result_ := C.C_NSMatrix_CellSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSMatrix) SetCellSize(value foundation.Size) {
	C.C_NSMatrix_SetCellSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSMatrix) IntercellSpacing() foundation.Size {
	result_ := C.C_NSMatrix_IntercellSpacing(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSMatrix) SetIntercellSpacing(value foundation.Size) {
	C.C_NSMatrix_SetIntercellSpacing(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSMatrix) NumberOfColumns() int {
	result_ := C.C_NSMatrix_NumberOfColumns(n.Ptr())
	return int(result_)
}

func (n NSMatrix) NumberOfRows() int {
	result_ := C.C_NSMatrix_NumberOfRows(n.Ptr())
	return int(result_)
}

func (n NSMatrix) AutorecalculatesCellSize() bool {
	result_ := C.C_NSMatrix_AutorecalculatesCellSize(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetAutorecalculatesCellSize(value bool) {
	C.C_NSMatrix_SetAutorecalculatesCellSize(n.Ptr(), C.bool(value))
}

func (n NSMatrix) KeyCell() Cell {
	result_ := C.C_NSMatrix_KeyCell(n.Ptr())
	return MakeCell(result_)
}

func (n NSMatrix) SetKeyCell(value Cell) {
	C.C_NSMatrix_SetKeyCell(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMatrix) SelectedCell() Cell {
	result_ := C.C_NSMatrix_SelectedCell(n.Ptr())
	return MakeCell(result_)
}

func (n NSMatrix) SelectedCells() []Cell {
	result_ := C.C_NSMatrix_SelectedCells(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Cell, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCell(r)
	}
	return goResult_
}

func (n NSMatrix) SelectedColumn() int {
	result_ := C.C_NSMatrix_SelectedColumn(n.Ptr())
	return int(result_)
}

func (n NSMatrix) SelectedRow() int {
	result_ := C.C_NSMatrix_SelectedRow(n.Ptr())
	return int(result_)
}

func (n NSMatrix) Cells() []Cell {
	result_ := C.C_NSMatrix_Cells(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Cell, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCell(r)
	}
	return goResult_
}

func (n NSMatrix) BackgroundColor() Color {
	result_ := C.C_NSMatrix_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSMatrix) SetBackgroundColor(value Color) {
	C.C_NSMatrix_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMatrix) CellBackgroundColor() Color {
	result_ := C.C_NSMatrix_CellBackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSMatrix) SetCellBackgroundColor(value Color) {
	C.C_NSMatrix_SetCellBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMatrix) DrawsBackground() bool {
	result_ := C.C_NSMatrix_DrawsBackground(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetDrawsBackground(value bool) {
	C.C_NSMatrix_SetDrawsBackground(n.Ptr(), C.bool(value))
}

func (n NSMatrix) DrawsCellBackground() bool {
	result_ := C.C_NSMatrix_DrawsCellBackground(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetDrawsCellBackground(value bool) {
	C.C_NSMatrix_SetDrawsCellBackground(n.Ptr(), C.bool(value))
}

func (n NSMatrix) TabKeyTraversesCells() bool {
	result_ := C.C_NSMatrix_TabKeyTraversesCells(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetTabKeyTraversesCells(value bool) {
	C.C_NSMatrix_SetTabKeyTraversesCells(n.Ptr(), C.bool(value))
}

func (n NSMatrix) Delegate() objc.Object {
	result_ := C.C_NSMatrix_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSMatrix) SetDelegate(value objc.Object) {
	C.C_NSMatrix_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMatrix) AutosizesCells() bool {
	result_ := C.C_NSMatrix_AutosizesCells(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetAutosizesCells(value bool) {
	C.C_NSMatrix_SetAutosizesCells(n.Ptr(), C.bool(value))
}

func (n NSMatrix) IsAutoscroll() bool {
	result_ := C.C_NSMatrix_IsAutoscroll(n.Ptr())
	return bool(result_)
}

func (n NSMatrix) SetAutoscroll(value bool) {
	C.C_NSMatrix_SetAutoscroll(n.Ptr(), C.bool(value))
}

func (n NSMatrix) DoubleAction() objc.Selector {
	result_ := C.C_NSMatrix_DoubleAction(n.Ptr())
	return objc.Selector(result_)
}

func (n NSMatrix) SetDoubleAction(value objc.Selector) {
	C.C_NSMatrix_SetDoubleAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSMatrix) MouseDownFlags() int {
	result_ := C.C_NSMatrix_MouseDownFlags(n.Ptr())
	return int(result_)
}
