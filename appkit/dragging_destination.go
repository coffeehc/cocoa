package appkit

// #include "dragging_destination.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
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

func WrapDraggingDestination(delegate *DraggingDestination) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapDraggingDestination(C.long(id))
	return objc.MakeObject(ptr)
}

//export draggingDestination_DraggingEntered
func draggingDestination_DraggingEntered(id int64, sender unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*DraggingDestination)
	result := delegate.DraggingEntered(objc.MakeObject(sender))
	return C.uint(uint(result))
}

//export draggingDestination_WantsPeriodicDraggingUpdates
func draggingDestination_WantsPeriodicDraggingUpdates(id int64) C.bool {
	delegate := resources.Get(id).(*DraggingDestination)
	result := delegate.WantsPeriodicDraggingUpdates()
	return C.bool(result)
}

//export draggingDestination_DraggingUpdated
func draggingDestination_DraggingUpdated(id int64, sender unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*DraggingDestination)
	result := delegate.DraggingUpdated(objc.MakeObject(sender))
	return C.uint(uint(result))
}

//export draggingDestination_DraggingEnded
func draggingDestination_DraggingEnded(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*DraggingDestination)
	delegate.DraggingEnded(objc.MakeObject(sender))
}

//export draggingDestination_DraggingExited
func draggingDestination_DraggingExited(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*DraggingDestination)
	delegate.DraggingExited(objc.MakeObject(sender))
}

//export draggingDestination_PrepareForDragOperation
func draggingDestination_PrepareForDragOperation(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*DraggingDestination)
	result := delegate.PrepareForDragOperation(objc.MakeObject(sender))
	return C.bool(result)
}

//export draggingDestination_PerformDragOperation
func draggingDestination_PerformDragOperation(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*DraggingDestination)
	result := delegate.PerformDragOperation(objc.MakeObject(sender))
	return C.bool(result)
}

//export draggingDestination_ConcludeDragOperation
func draggingDestination_ConcludeDragOperation(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*DraggingDestination)
	delegate.ConcludeDragOperation(objc.MakeObject(sender))
}

//export draggingDestination_UpdateDraggingItemsForDrag
func draggingDestination_UpdateDraggingItemsForDrag(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*DraggingDestination)
	delegate.UpdateDraggingItemsForDrag(objc.MakeObject(sender))
}

//export DraggingDestination_RespondsTo
func DraggingDestination_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*DraggingDestination)
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

//export deleteDraggingDestination
func deleteDraggingDestination(id int64) {
	resources.Delete(id)
}
