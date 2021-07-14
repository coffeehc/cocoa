package appkit

// #include "path_control_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type PathControlDelegate struct {
	PathControl_ShouldDragPathComponentCell_WithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	PathControl_ValidateDrop                               func(pathControl PathControl, info objc.Object) DragOperation
	PathControl_AcceptDrop                                 func(pathControl PathControl, info objc.Object) bool
	PathControl_WillDisplayOpenPanel                       func(pathControl PathControl, openPanel OpenPanel)
	PathControl_WillPopUpMenu                              func(pathControl PathControl, menu Menu)
	PathControl_ShouldDragItem_WithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func WrapPathControlDelegate(delegate *PathControlDelegate) objc.Object {
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

//export deletePathControlDelegate
func deletePathControlDelegate(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
