package appkit

// #include "grid_column.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

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
