package appkit

// #include "path_control.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type PathControl interface {
	Control
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.URL)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	ClickedPathItem() PathControlItem
	IsEditable() bool
	SetEditable(value bool)
	PathItems() []PathControlItem
	SetPathItems(value []PathControlItem)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
}

type NSPathControl struct {
	NSControl
}

func MakePathControl(ptr unsafe.Pointer) NSPathControl {
	return NSPathControl{
		NSControl: MakeControl(ptr),
	}
}

func (n NSPathControl) InitWithFrame(frameRect foundation.Rect) NSPathControl {
	result_ := C.C_NSPathControl_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakePathControl(result_)
}

func (n NSPathControl) InitWithCoder(coder foundation.Coder) NSPathControl {
	result_ := C.C_NSPathControl_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePathControl(result_)
}

func (n NSPathControl) Init() NSPathControl {
	result_ := C.C_NSPathControl_Init(n.Ptr())
	return MakePathControl(result_)
}

func AllocPathControl() NSPathControl {
	result_ := C.C_NSPathControl_AllocPathControl()
	return MakePathControl(result_)
}

func NewPathControl() NSPathControl {
	result_ := C.C_NSPathControl_NewPathControl()
	return MakePathControl(result_)
}

func (n NSPathControl) Autorelease() NSPathControl {
	result_ := C.C_NSPathControl_Autorelease(n.Ptr())
	return MakePathControl(result_)
}

func (n NSPathControl) Retain() NSPathControl {
	result_ := C.C_NSPathControl_Retain(n.Ptr())
	return MakePathControl(result_)
}

func (n NSPathControl) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	C.C_NSPathControl_SetDraggingSourceOperationMask_ForLocal(n.Ptr(), C.uint(uint(mask)), C.bool(isLocal))
}

func (n NSPathControl) PathStyle() PathStyle {
	result_ := C.C_NSPathControl_PathStyle(n.Ptr())
	return PathStyle(int(result_))
}

func (n NSPathControl) SetPathStyle(value PathStyle) {
	C.C_NSPathControl_SetPathStyle(n.Ptr(), C.int(int(value)))
}

func (n NSPathControl) BackgroundColor() Color {
	result_ := C.C_NSPathControl_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSPathControl) SetBackgroundColor(value Color) {
	C.C_NSPathControl_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) DoubleAction() objc.Selector {
	result_ := C.C_NSPathControl_DoubleAction(n.Ptr())
	return objc.Selector(result_)
}

func (n NSPathControl) SetDoubleAction(value objc.Selector) {
	C.C_NSPathControl_SetDoubleAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSPathControl) URL() foundation.URL {
	result_ := C.C_NSPathControl_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSPathControl) SetURL(value foundation.URL) {
	C.C_NSPathControl_SetURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) Delegate() objc.Object {
	result_ := C.C_NSPathControl_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSPathControl) SetDelegate(value objc.Object) {
	C.C_NSPathControl_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) AllowedTypes() []string {
	result_ := C.C_NSPathControl_AllowedTypes(n.Ptr())
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

func (n NSPathControl) SetAllowedTypes(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathControl_SetAllowedTypes(n.Ptr(), cValue)
}

func (n NSPathControl) ClickedPathItem() PathControlItem {
	result_ := C.C_NSPathControl_ClickedPathItem(n.Ptr())
	return MakePathControlItem(result_)
}

func (n NSPathControl) IsEditable() bool {
	result_ := C.C_NSPathControl_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSPathControl) SetEditable(value bool) {
	C.C_NSPathControl_SetEditable(n.Ptr(), C.bool(value))
}

func (n NSPathControl) PathItems() []PathControlItem {
	result_ := C.C_NSPathControl_PathItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PathControlItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakePathControlItem(r)
	}
	return goResult_
}

func (n NSPathControl) SetPathItems(value []PathControlItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathControl_SetPathItems(n.Ptr(), cValue)
}

func (n NSPathControl) PlaceholderAttributedString() foundation.AttributedString {
	result_ := C.C_NSPathControl_PlaceholderAttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSPathControl) SetPlaceholderAttributedString(value foundation.AttributedString) {
	C.C_NSPathControl_SetPlaceholderAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) PlaceholderString() string {
	result_ := C.C_NSPathControl_PlaceholderString(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPathControl) SetPlaceholderString(value string) {
	C.C_NSPathControl_SetPlaceholderString(n.Ptr(), foundation.NewString(value).Ptr())
}

type PathControlDelegate struct {
	PathControl_ShouldDragPathComponentCell_WithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	PathControl_ValidateDrop                               func(pathControl PathControl, info objc.Object) DragOperation
	PathControl_AcceptDrop                                 func(pathControl PathControl, info objc.Object) bool
	PathControl_WillDisplayOpenPanel                       func(pathControl PathControl, openPanel OpenPanel)
	PathControl_WillPopUpMenu                              func(pathControl PathControl, menu Menu)
	PathControl_ShouldDragItem_WithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func (delegate *PathControlDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapPathControlDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export pathControlDelegate_PathControl_ShouldDragPathComponentCell_WithPasteboard
func pathControlDelegate_PathControl_ShouldDragPathComponentCell_WithPasteboard(hp C.uintptr_t, pathControl unsafe.Pointer, pathComponentCell unsafe.Pointer, pasteboard unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	result := delegate.PathControl_ShouldDragPathComponentCell_WithPasteboard(MakePathControl(pathControl), MakePathComponentCell(pathComponentCell), MakePasteboard(pasteboard))
	return C.bool(result)
}

//export pathControlDelegate_PathControl_ValidateDrop
func pathControlDelegate_PathControl_ValidateDrop(hp C.uintptr_t, pathControl unsafe.Pointer, info unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	result := delegate.PathControl_ValidateDrop(MakePathControl(pathControl), objc.MakeObject(info))
	return C.uint(uint(result))
}

//export pathControlDelegate_PathControl_AcceptDrop
func pathControlDelegate_PathControl_AcceptDrop(hp C.uintptr_t, pathControl unsafe.Pointer, info unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	result := delegate.PathControl_AcceptDrop(MakePathControl(pathControl), objc.MakeObject(info))
	return C.bool(result)
}

//export pathControlDelegate_PathControl_WillDisplayOpenPanel
func pathControlDelegate_PathControl_WillDisplayOpenPanel(hp C.uintptr_t, pathControl unsafe.Pointer, openPanel unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	delegate.PathControl_WillDisplayOpenPanel(MakePathControl(pathControl), MakeOpenPanel(openPanel))
}

//export pathControlDelegate_PathControl_WillPopUpMenu
func pathControlDelegate_PathControl_WillPopUpMenu(hp C.uintptr_t, pathControl unsafe.Pointer, menu unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	delegate.PathControl_WillPopUpMenu(MakePathControl(pathControl), MakeMenu(menu))
}

//export pathControlDelegate_PathControl_ShouldDragItem_WithPasteboard
func pathControlDelegate_PathControl_ShouldDragItem_WithPasteboard(hp C.uintptr_t, pathControl unsafe.Pointer, pathItem unsafe.Pointer, pasteboard unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	result := delegate.PathControl_ShouldDragItem_WithPasteboard(MakePathControl(pathControl), MakePathControlItem(pathItem), MakePasteboard(pasteboard))
	return C.bool(result)
}

//export PathControlDelegate_RespondsTo
func PathControlDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*PathControlDelegate)
	switch selName {
	case "pathControl:shouldDragPathComponentCell:withPasteboard:":
		return delegate.PathControl_ShouldDragPathComponentCell_WithPasteboard != nil
	case "pathControl:validateDrop:":
		return delegate.PathControl_ValidateDrop != nil
	case "pathControl:acceptDrop:":
		return delegate.PathControl_AcceptDrop != nil
	case "pathControl:willDisplayOpenPanel:":
		return delegate.PathControl_WillDisplayOpenPanel != nil
	case "pathControl:willPopUpMenu:":
		return delegate.PathControl_WillPopUpMenu != nil
	case "pathControl:shouldDragItem:withPasteboard:":
		return delegate.PathControl_ShouldDragItem_WithPasteboard != nil
	default:
		return false
	}
}
