package appkit

// #include "path_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type PathCell interface {
	ActionCell
	MouseEntered_WithFrame_InView(event Event, frame foundation.Rect, view View)
	MouseExited_WithFrame_InView(event Event, frame foundation.Rect, view View)
	RectOfPathComponentCell_WithFrame_InView(cell PathComponentCell, frame foundation.Rect, view View) foundation.Rect
	PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view View) PathComponentCell
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	ClickedPathComponentCell() PathComponentCell
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []PathComponentCell)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.URL)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
}

type NSPathCell struct {
	NSActionCell
}

func MakePathCell(ptr unsafe.Pointer) NSPathCell {
	return NSPathCell{
		NSActionCell: MakeActionCell(ptr),
	}
}

func (n NSPathCell) InitImageCell(image Image) NSPathCell {
	result_ := C.C_NSPathCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakePathCell(result_)
}

func (n NSPathCell) InitTextCell(_string string) NSPathCell {
	result_ := C.C_NSPathCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakePathCell(result_)
}

func (n NSPathCell) Init() NSPathCell {
	result_ := C.C_NSPathCell_Init(n.Ptr())
	return MakePathCell(result_)
}

func (n NSPathCell) InitWithCoder(coder foundation.Coder) NSPathCell {
	result_ := C.C_NSPathCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePathCell(result_)
}

func AllocPathCell() NSPathCell {
	result_ := C.C_NSPathCell_AllocPathCell()
	return MakePathCell(result_)
}

func NewPathCell() NSPathCell {
	result_ := C.C_NSPathCell_NewPathCell()
	return MakePathCell(result_)
}

func (n NSPathCell) Autorelease() NSPathCell {
	result_ := C.C_NSPathCell_Autorelease(n.Ptr())
	return MakePathCell(result_)
}

func (n NSPathCell) Retain() NSPathCell {
	result_ := C.C_NSPathCell_Retain(n.Ptr())
	return MakePathCell(result_)
}

func (n NSPathCell) MouseEntered_WithFrame_InView(event Event, frame foundation.Rect, view View) {
	C.C_NSPathCell_MouseEntered_WithFrame_InView(n.Ptr(), objc.ExtractPtr(event), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(view))
}

func (n NSPathCell) MouseExited_WithFrame_InView(event Event, frame foundation.Rect, view View) {
	C.C_NSPathCell_MouseExited_WithFrame_InView(n.Ptr(), objc.ExtractPtr(event), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(view))
}

func (n NSPathCell) RectOfPathComponentCell_WithFrame_InView(cell PathComponentCell, frame foundation.Rect, view View) foundation.Rect {
	result_ := C.C_NSPathCell_RectOfPathComponentCell_WithFrame_InView(n.Ptr(), objc.ExtractPtr(cell), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSPathCell) PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view View) PathComponentCell {
	result_ := C.C_NSPathCell_PathComponentCellAtPoint_WithFrame_InView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(view))
	return MakePathComponentCell(result_)
}

func (n NSPathCell) AllowedTypes() []string {
	result_ := C.C_NSPathCell_AllowedTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSPathCell) SetAllowedTypes(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathCell_SetAllowedTypes(n.Ptr(), cValue)
}

func (n NSPathCell) PathStyle() PathStyle {
	result_ := C.C_NSPathCell_PathStyle(n.Ptr())
	return PathStyle(int(result_))
}

func (n NSPathCell) SetPathStyle(value PathStyle) {
	C.C_NSPathCell_SetPathStyle(n.Ptr(), C.int(int(value)))
}

func (n NSPathCell) PlaceholderAttributedString() foundation.AttributedString {
	result_ := C.C_NSPathCell_PlaceholderAttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSPathCell) SetPlaceholderAttributedString(value foundation.AttributedString) {
	C.C_NSPathCell_SetPlaceholderAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathCell) PlaceholderString() string {
	result_ := C.C_NSPathCell_PlaceholderString(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPathCell) SetPlaceholderString(value string) {
	C.C_NSPathCell_SetPlaceholderString(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSPathCell) BackgroundColor() Color {
	result_ := C.C_NSPathCell_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSPathCell) SetBackgroundColor(value Color) {
	C.C_NSPathCell_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathCell) ClickedPathComponentCell() PathComponentCell {
	result_ := C.C_NSPathCell_ClickedPathComponentCell(n.Ptr())
	return MakePathComponentCell(result_)
}

func (n NSPathCell) PathComponentCells() []PathComponentCell {
	result_ := C.C_NSPathCell_PathComponentCells(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PathComponentCell, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakePathComponentCell(r)
	}
	return goResult_
}

func (n NSPathCell) SetPathComponentCells(value []PathComponentCell) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathCell_SetPathComponentCells(n.Ptr(), cValue)
}

func (n NSPathCell) DoubleAction() objc.Selector {
	result_ := C.C_NSPathCell_DoubleAction(n.Ptr())
	return objc.Selector(result_)
}

func (n NSPathCell) SetDoubleAction(value objc.Selector) {
	C.C_NSPathCell_SetDoubleAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSPathCell) URL() foundation.URL {
	result_ := C.C_NSPathCell_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSPathCell) SetURL(value foundation.URL) {
	C.C_NSPathCell_SetURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathCell) Delegate() objc.Object {
	result_ := C.C_NSPathCell_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSPathCell) SetDelegate(value objc.Object) {
	C.C_NSPathCell_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

type PathCellDelegate struct {
	PathCell_WillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	PathCell_WillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func (delegate *PathCellDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapPathCellDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export pathCellDelegate_PathCell_WillDisplayOpenPanel
func pathCellDelegate_PathCell_WillDisplayOpenPanel(hp C.uintptr_t, pathCell unsafe.Pointer, openPanel unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*PathCellDelegate)
	delegate.PathCell_WillDisplayOpenPanel(MakePathCell(pathCell), MakeOpenPanel(openPanel))
}

//export pathCellDelegate_PathCell_WillPopUpMenu
func pathCellDelegate_PathCell_WillPopUpMenu(hp C.uintptr_t, pathCell unsafe.Pointer, menu unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*PathCellDelegate)
	delegate.PathCell_WillPopUpMenu(MakePathCell(pathCell), MakeMenu(menu))
}

//export PathCellDelegate_RespondsTo
func PathCellDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*PathCellDelegate)
	switch selName {
	case "pathCell:willDisplayOpenPanel:":
		return delegate.PathCell_WillDisplayOpenPanel != nil
	case "pathCell:willPopUpMenu:":
		return delegate.PathCell_WillPopUpMenu != nil
	default:
		return false
	}
}
