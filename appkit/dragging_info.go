package appkit

// #include "dragging_info.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type DraggingInfo struct {
	SlideDraggedImageTo func(screenPoint foundation.Point) // required
	ResetSpringLoading  func()                             // required
}

func (delegate *DraggingInfo) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapDraggingInfo(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export draggingInfo_SlideDraggedImageTo
func draggingInfo_SlideDraggedImageTo(hp C.uintptr_t, screenPoint C.CGPoint) {
	delegate := cgo.Handle(hp).Value().(*DraggingInfo)
	delegate.SlideDraggedImageTo(foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingInfo_ResetSpringLoading
func draggingInfo_ResetSpringLoading(hp C.uintptr_t) {
	delegate := cgo.Handle(hp).Value().(*DraggingInfo)
	delegate.ResetSpringLoading()
}

//export DraggingInfo_RespondsTo
func DraggingInfo_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*DraggingInfo)
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
func deleteDraggingInfo(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
