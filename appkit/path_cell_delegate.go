package appkit

// #include "path_cell_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

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
