package appkit

// #include "dragging_info.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DraggingInfo struct {
	SlideDraggedImageTo func(screenPoint foundation.Point) // required
	ResetSpringLoading  func()                             // required
}

func WrapDraggingInfo(delegate *DraggingInfo) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapDraggingInfo(C.long(id))
	return objc.MakeObject(ptr)
}

//export draggingInfo_SlideDraggedImageTo
func draggingInfo_SlideDraggedImageTo(id int64, screenPoint C.CGPoint) {
	delegate := resources.Get(id).(*DraggingInfo)
	delegate.SlideDraggedImageTo(foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingInfo_ResetSpringLoading
func draggingInfo_ResetSpringLoading(id int64) {
	delegate := resources.Get(id).(*DraggingInfo)
	delegate.ResetSpringLoading()
}

//export DraggingInfo_RespondsTo
func DraggingInfo_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*DraggingInfo)
	switch selName {
	case "slideDraggedImageTo:":
		return delegate.SlideDraggedImageTo != nil
	case "resetSpringLoading:":
		return delegate.ResetSpringLoading != nil
	default:
		return false
	}
}

//export deleteDraggingInfo
func deleteDraggingInfo(id int64) {
	resources.Delete(id)
}
