package appkit

// #include "table_header_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableHeaderView interface {
	View
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
	TableView() TableView
	SetTableView(value TableView)
	DraggedColumn() int
	DraggedDistance() coregraphics.Float
	ResizedColumn() int
}

type NSTableHeaderView struct {
	NSView
}

func MakeTableHeaderView(ptr unsafe.Pointer) NSTableHeaderView {
	return NSTableHeaderView{
		NSView: MakeView(ptr),
	}
}

func AllocTableHeaderView() NSTableHeaderView {
	return MakeTableHeaderView(C.C_TableHeaderView_Alloc())
}

func (n NSTableHeaderView) InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	result_ := C.C_NSTableHeaderView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableHeaderView(result_)
}

func (n NSTableHeaderView) InitWithCoder(coder foundation.Coder) TableHeaderView {
	result_ := C.C_NSTableHeaderView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableHeaderView(result_)
}

func (n NSTableHeaderView) Init() TableHeaderView {
	result_ := C.C_NSTableHeaderView_Init(n.Ptr())
	return MakeTableHeaderView(result_)
}

func (n NSTableHeaderView) ColumnAtPoint(point foundation.Point) int {
	result_ := C.C_NSTableHeaderView_ColumnAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return int(result_)
}

func (n NSTableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	result_ := C.C_NSTableHeaderView_HeaderRectOfColumn(n.Ptr(), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTableHeaderView) TableView() TableView {
	result_ := C.C_NSTableHeaderView_TableView(n.Ptr())
	return MakeTableView(result_)
}

func (n NSTableHeaderView) SetTableView(value TableView) {
	C.C_NSTableHeaderView_SetTableView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableHeaderView) DraggedColumn() int {
	result_ := C.C_NSTableHeaderView_DraggedColumn(n.Ptr())
	return int(result_)
}

func (n NSTableHeaderView) DraggedDistance() coregraphics.Float {
	result_ := C.C_NSTableHeaderView_DraggedDistance(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTableHeaderView) ResizedColumn() int {
	result_ := C.C_NSTableHeaderView_ResizedColumn(n.Ptr())
	return int(result_)
}
