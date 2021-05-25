package appkit

// #include "path_cell_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PathCellDelegate struct {
	PathCell_WillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	PathCell_WillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func WrapPathCellDelegate(delegate *PathCellDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapPathCellDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export pathCellDelegate_PathCell_WillDisplayOpenPanel
func pathCellDelegate_PathCell_WillDisplayOpenPanel(id int64, pathCell unsafe.Pointer, openPanel unsafe.Pointer) {
	delegate := resources.Get(id).(*PathCellDelegate)
	delegate.PathCell_WillDisplayOpenPanel(MakePathCell(pathCell), MakeOpenPanel(openPanel))
}

//export pathCellDelegate_PathCell_WillPopUpMenu
func pathCellDelegate_PathCell_WillPopUpMenu(id int64, pathCell unsafe.Pointer, menu unsafe.Pointer) {
	delegate := resources.Get(id).(*PathCellDelegate)
	delegate.PathCell_WillPopUpMenu(MakePathCell(pathCell), MakeMenu(menu))
}

//export PathCellDelegate_RespondsTo
func PathCellDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*PathCellDelegate)
	switch selName {
	case "pathCell:willDisplayOpenPanel:":
		return delegate.PathCell_WillDisplayOpenPanel != nil
	case "pathCell:willPopUpMenu:":
		return delegate.PathCell_WillPopUpMenu != nil
	default:
		return false
	}
}

//export deletePathCellDelegate
func deletePathCellDelegate(id int64) {
	resources.Delete(id)
}
