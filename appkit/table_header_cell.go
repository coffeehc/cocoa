package appkit

// #include "table_header_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableHeaderCell interface {
	TextFieldCell
	DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView View, ascending bool, priority int)
	SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect
}

type NSTableHeaderCell struct {
	NSTextFieldCell
}

func MakeTableHeaderCell(ptr unsafe.Pointer) *NSTableHeaderCell {
	if ptr == nil {
		return nil
	}
	return &NSTableHeaderCell{
		NSTextFieldCell: *MakeTextFieldCell(ptr),
	}
}

func AllocTableHeaderCell() *NSTableHeaderCell {
	return MakeTableHeaderCell(C.C_TableHeaderCell_Alloc())
}

func (n *NSTableHeaderCell) InitTextCell(_string string) TableHeaderCell {
	result_ := C.C_NSTableHeaderCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeTableHeaderCell(result_)
}

func (n *NSTableHeaderCell) InitWithCoder(coder foundation.Coder) TableHeaderCell {
	result_ := C.C_NSTableHeaderCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableHeaderCell(result_)
}

func (n *NSTableHeaderCell) Init() TableHeaderCell {
	result_ := C.C_NSTableHeaderCell_Init(n.Ptr())
	return MakeTableHeaderCell(result_)
}

func (n *NSTableHeaderCell) DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView View, ascending bool, priority int) {
	C.C_NSTableHeaderCell_DrawSortIndicatorWithFrame_InView_Ascending_Priority(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cellFrame))), objc.ExtractPtr(controlView), C.bool(ascending), C.int(priority))
}

func (n *NSTableHeaderCell) SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSTableHeaderCell_SortIndicatorRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}
