package appkit

// #include "table_row_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableRowView interface {
	View
	DrawBackgroundInRect(dirtyRect foundation.Rect)
	DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect)
	DrawSelectionInRect(dirtyRect foundation.Rect)
	DrawSeparatorInRect(dirtyRect foundation.Rect)
	ViewAtColumn(column int) objc.Object
	IsEmphasized() bool
	SetEmphasized(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	IsFloating() bool
	SetFloating(value bool)
	IsSelected() bool
	SetSelected(value bool)
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	IndentationForDropOperation() coregraphics.Float
	SetIndentationForDropOperation(value coregraphics.Float)
	IsTargetForDropOperation() bool
	SetTargetForDropOperation(value bool)
	IsGroupRowStyle() bool
	SetGroupRowStyle(value bool)
	NumberOfColumns() int
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	IsNextRowSelected() bool
	SetNextRowSelected(value bool)
	IsPreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
}

type NSTableRowView struct {
	NSView
}

func MakeTableRowView(ptr unsafe.Pointer) *NSTableRowView {
	if ptr == nil {
		return nil
	}
	return &NSTableRowView{
		NSView: *MakeView(ptr),
	}
}

func AllocTableRowView() *NSTableRowView {
	return MakeTableRowView(C.C_TableRowView_Alloc())
}

func (n *NSTableRowView) InitWithFrame(frameRect foundation.Rect) TableRowView {
	result_ := C.C_NSTableRowView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableRowView(result_)
}

func (n *NSTableRowView) InitWithCoder(coder foundation.Coder) TableRowView {
	result_ := C.C_NSTableRowView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableRowView(result_)
}

func (n *NSTableRowView) Init() TableRowView {
	result_ := C.C_NSTableRowView_Init(n.Ptr())
	return MakeTableRowView(result_)
}

func (n *NSTableRowView) DrawBackgroundInRect(dirtyRect foundation.Rect) {
	C.C_NSTableRowView_DrawBackgroundInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n *NSTableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect) {
	C.C_NSTableRowView_DrawDraggingDestinationFeedbackInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n *NSTableRowView) DrawSelectionInRect(dirtyRect foundation.Rect) {
	C.C_NSTableRowView_DrawSelectionInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n *NSTableRowView) DrawSeparatorInRect(dirtyRect foundation.Rect) {
	C.C_NSTableRowView_DrawSeparatorInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n *NSTableRowView) ViewAtColumn(column int) objc.Object {
	result_ := C.C_NSTableRowView_ViewAtColumn(n.Ptr(), C.int(column))
	return objc.MakeObject(result_)
}

func (n *NSTableRowView) IsEmphasized() bool {
	result_ := C.C_NSTableRowView_IsEmphasized(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetEmphasized(value bool) {
	C.C_NSTableRowView_SetEmphasized(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) InteriorBackgroundStyle() BackgroundStyle {
	result_ := C.C_NSTableRowView_InteriorBackgroundStyle(n.Ptr())
	return BackgroundStyle(int(result_))
}

func (n *NSTableRowView) IsFloating() bool {
	result_ := C.C_NSTableRowView_IsFloating(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetFloating(value bool) {
	C.C_NSTableRowView_SetFloating(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) IsSelected() bool {
	result_ := C.C_NSTableRowView_IsSelected(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetSelected(value bool) {
	C.C_NSTableRowView_SetSelected(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	result_ := C.C_NSTableRowView_SelectionHighlightStyle(n.Ptr())
	return TableViewSelectionHighlightStyle(int(result_))
}

func (n *NSTableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	C.C_NSTableRowView_SetSelectionHighlightStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSTableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	result_ := C.C_NSTableRowView_DraggingDestinationFeedbackStyle(n.Ptr())
	return TableViewDraggingDestinationFeedbackStyle(int(result_))
}

func (n *NSTableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	C.C_NSTableRowView_SetDraggingDestinationFeedbackStyle(n.Ptr(), C.int(int(value)))
}

func (n *NSTableRowView) IndentationForDropOperation() coregraphics.Float {
	result_ := C.C_NSTableRowView_IndentationForDropOperation(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTableRowView) SetIndentationForDropOperation(value coregraphics.Float) {
	C.C_NSTableRowView_SetIndentationForDropOperation(n.Ptr(), C.double(float64(value)))
}

func (n *NSTableRowView) IsTargetForDropOperation() bool {
	result_ := C.C_NSTableRowView_IsTargetForDropOperation(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetTargetForDropOperation(value bool) {
	C.C_NSTableRowView_SetTargetForDropOperation(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) IsGroupRowStyle() bool {
	result_ := C.C_NSTableRowView_IsGroupRowStyle(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetGroupRowStyle(value bool) {
	C.C_NSTableRowView_SetGroupRowStyle(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) NumberOfColumns() int {
	result_ := C.C_NSTableRowView_NumberOfColumns(n.Ptr())
	return int(result_)
}

func (n *NSTableRowView) BackgroundColor() Color {
	result_ := C.C_NSTableRowView_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSTableRowView) SetBackgroundColor(value Color) {
	C.C_NSTableRowView_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTableRowView) IsNextRowSelected() bool {
	result_ := C.C_NSTableRowView_IsNextRowSelected(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetNextRowSelected(value bool) {
	C.C_NSTableRowView_SetNextRowSelected(n.Ptr(), C.bool(value))
}

func (n *NSTableRowView) IsPreviousRowSelected() bool {
	result_ := C.C_NSTableRowView_IsPreviousRowSelected(n.Ptr())
	return bool(result_)
}

func (n *NSTableRowView) SetPreviousRowSelected(value bool) {
	C.C_NSTableRowView_SetPreviousRowSelected(n.Ptr(), C.bool(value))
}
