package appkit

// #include "dragging_destination.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type DraggingDestination struct {
	DraggingEntered              func(sender objc.Object) DragOperation
	WantsPeriodicDraggingUpdates func() bool
	DraggingUpdated              func(sender objc.Object) DragOperation
	DraggingEnded                func(sender objc.Object)
	DraggingExited               func(sender objc.Object)
	PrepareForDragOperation      func(sender objc.Object) bool
	PerformDragOperation         func(sender objc.Object) bool
	ConcludeDragOperation        func(sender objc.Object)
	UpdateDraggingItemsForDrag   func(sender objc.Object)
}

func (delegate *DraggingDestination) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapDraggingDestination(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export draggingDestination_DraggingEntered
func draggingDestination_DraggingEntered(hp C.uintptr_t, sender unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	result := delegate.DraggingEntered(objc.MakeObject(sender))
	return C.uint(uint(result))
}

//export draggingDestination_WantsPeriodicDraggingUpdates
func draggingDestination_WantsPeriodicDraggingUpdates(hp C.uintptr_t) C.bool {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	result := delegate.WantsPeriodicDraggingUpdates()
	return C.bool(result)
}

//export draggingDestination_DraggingUpdated
func draggingDestination_DraggingUpdated(hp C.uintptr_t, sender unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	result := delegate.DraggingUpdated(objc.MakeObject(sender))
	return C.uint(uint(result))
}

//export draggingDestination_DraggingEnded
func draggingDestination_DraggingEnded(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	delegate.DraggingEnded(objc.MakeObject(sender))
}

//export draggingDestination_DraggingExited
func draggingDestination_DraggingExited(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	delegate.DraggingExited(objc.MakeObject(sender))
}

//export draggingDestination_PrepareForDragOperation
func draggingDestination_PrepareForDragOperation(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	result := delegate.PrepareForDragOperation(objc.MakeObject(sender))
	return C.bool(result)
}

//export draggingDestination_PerformDragOperation
func draggingDestination_PerformDragOperation(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	result := delegate.PerformDragOperation(objc.MakeObject(sender))
	return C.bool(result)
}

//export draggingDestination_ConcludeDragOperation
func draggingDestination_ConcludeDragOperation(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	delegate.ConcludeDragOperation(objc.MakeObject(sender))
}

//export draggingDestination_UpdateDraggingItemsForDrag
func draggingDestination_UpdateDraggingItemsForDrag(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	delegate.UpdateDraggingItemsForDrag(objc.MakeObject(sender))
}

//export DraggingDestination_RespondsTo
func DraggingDestination_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*DraggingDestination)
	switch selName {
	case "draggingEntered:":
		return delegate.DraggingEntered != nil
	case "wantsPeriodicDraggingUpdates:":
		return delegate.WantsPeriodicDraggingUpdates != nil
	case "draggingUpdated:":
		return delegate.DraggingUpdated != nil
	case "draggingEnded:":
		return delegate.DraggingEnded != nil
	case "draggingExited:":
		return delegate.DraggingExited != nil
	case "prepareForDragOperation:":
		return delegate.PrepareForDragOperation != nil
	case "performDragOperation:":
		return delegate.PerformDragOperation != nil
	case "concludeDragOperation:":
		return delegate.ConcludeDragOperation != nil
	case "updateDraggingItemsForDrag:":
		return delegate.UpdateDraggingItemsForDrag != nil
	default:
		return false
	}
}
