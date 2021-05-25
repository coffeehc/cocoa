package appkit

// #include "spring_loading_destination.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SpringLoadingDestination struct {
	SpringLoadingActivated_DraggingInfo func(activated bool, draggingInfo objc.Object)
	SpringLoadingHighlightChanged       func(draggingInfo objc.Object)
	SpringLoadingEntered                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingUpdated                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingExited                 func(draggingInfo objc.Object)
	DraggingEnded                       func(draggingInfo objc.Object)
}

func WrapSpringLoadingDestination(delegate *SpringLoadingDestination) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapSpringLoadingDestination(C.long(id))
	return objc.MakeObject(ptr)
}

//export springLoadingDestination_SpringLoadingActivated_DraggingInfo
func springLoadingDestination_SpringLoadingActivated_DraggingInfo(id int64, activated C.bool, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	delegate.SpringLoadingActivated_DraggingInfo(bool(activated), objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingHighlightChanged
func springLoadingDestination_SpringLoadingHighlightChanged(id int64, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	delegate.SpringLoadingHighlightChanged(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingEntered
func springLoadingDestination_SpringLoadingEntered(id int64, draggingInfo unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	result := delegate.SpringLoadingEntered(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingUpdated
func springLoadingDestination_SpringLoadingUpdated(id int64, draggingInfo unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	result := delegate.SpringLoadingUpdated(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingExited
func springLoadingDestination_SpringLoadingExited(id int64, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	delegate.SpringLoadingExited(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_DraggingEnded
func springLoadingDestination_DraggingEnded(id int64, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*SpringLoadingDestination)
	delegate.DraggingEnded(objc.MakeObject(draggingInfo))
}

//export SpringLoadingDestination_RespondsTo
func SpringLoadingDestination_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*SpringLoadingDestination)
	switch selName {
	case "springLoadingActivated:draggingInfo:":
		return delegate.SpringLoadingActivated_DraggingInfo != nil
	case "springLoadingHighlightChanged:":
		return delegate.SpringLoadingHighlightChanged != nil
	case "springLoadingEntered:":
		return delegate.SpringLoadingEntered != nil
	case "springLoadingUpdated:":
		return delegate.SpringLoadingUpdated != nil
	case "springLoadingExited:":
		return delegate.SpringLoadingExited != nil
	case "draggingEnded:":
		return delegate.DraggingEnded != nil
	default:
		return false
	}
}

//export deleteSpringLoadingDestination
func deleteSpringLoadingDestination(id int64) {
	resources.Delete(id)
}
