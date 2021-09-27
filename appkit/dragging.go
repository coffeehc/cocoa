package appkit

// #include "dragging.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type DraggingSource interface {
	DraggingSession_SourceOperationMaskForDraggingContext(session DraggingSession, context DraggingContext) DragOperation
	HasDraggingSession_EndedAtPoint_Operation() bool
	DraggingSession_EndedAtPoint_Operation(session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	HasDraggingSession_MovedToPoint() bool
	DraggingSession_MovedToPoint(session DraggingSession, screenPoint foundation.Point)
	HasDraggingSession_WillBeginAtPoint() bool
	DraggingSession_WillBeginAtPoint(session DraggingSession, screenPoint foundation.Point)
	HasIgnoreModifierKeysForDraggingSession() bool
	IgnoreModifierKeysForDraggingSession(session DraggingSession) bool
}

func DraggingSourceToObjc(protocol DraggingSource) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapDraggingSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export draggingSource_DraggingSession_SourceOperationMaskForDraggingContext
func draggingSource_DraggingSession_SourceOperationMaskForDraggingContext(hp C.uintptr_t, session unsafe.Pointer, context C.int) C.uint {
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	result := protocol.DraggingSession_SourceOperationMaskForDraggingContext(MakeDraggingSession(session), DraggingContext(int(context)))
	return C.uint(uint(result))
}

//export draggingSource_DraggingSession_EndedAtPoint_Operation
func draggingSource_DraggingSession_EndedAtPoint_Operation(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	protocol.DraggingSession_EndedAtPoint_Operation(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export draggingSource_DraggingSession_MovedToPoint
func draggingSource_DraggingSession_MovedToPoint(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint) {
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	protocol.DraggingSession_MovedToPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_DraggingSession_WillBeginAtPoint
func draggingSource_DraggingSession_WillBeginAtPoint(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint) {
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	protocol.DraggingSession_WillBeginAtPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_IgnoreModifierKeysForDraggingSession
func draggingSource_IgnoreModifierKeysForDraggingSession(hp C.uintptr_t, session unsafe.Pointer) C.bool {
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	result := protocol.IgnoreModifierKeysForDraggingSession(MakeDraggingSession(session))
	return C.bool(result)
}

//export DraggingSource_RespondsTo
func DraggingSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(DraggingSource)
	_ = protocol
	switch selName {
	case "draggingSession:sourceOperationMaskForDraggingContext:":
		return true
	case "draggingSession:endedAtPoint:operation:":
		return protocol.HasDraggingSession_EndedAtPoint_Operation()
	case "draggingSession:movedToPoint:":
		return protocol.HasDraggingSession_MovedToPoint()
	case "draggingSession:willBeginAtPoint:":
		return protocol.HasDraggingSession_WillBeginAtPoint()
	case "ignoreModifierKeysForDraggingSession:":
		return protocol.HasIgnoreModifierKeysForDraggingSession()
	default:
		return false
	}
}

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
	case "wantsPeriodicDraggingUpdates":
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

type SpringLoadingDestination struct {
	SpringLoadingActivated_DraggingInfo func(activated bool, draggingInfo objc.Object) // required
	SpringLoadingHighlightChanged       func(draggingInfo objc.Object)                 // required
	SpringLoadingEntered                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingUpdated                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingExited                 func(draggingInfo objc.Object)
	DraggingEnded                       func(draggingInfo objc.Object)
}

func (delegate *SpringLoadingDestination) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapSpringLoadingDestination(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export springLoadingDestination_SpringLoadingActivated_DraggingInfo
func springLoadingDestination_SpringLoadingActivated_DraggingInfo(hp C.uintptr_t, activated C.bool, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingActivated_DraggingInfo(bool(activated), objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingHighlightChanged
func springLoadingDestination_SpringLoadingHighlightChanged(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingHighlightChanged(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingEntered
func springLoadingDestination_SpringLoadingEntered(hp C.uintptr_t, draggingInfo unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	result := delegate.SpringLoadingEntered(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingUpdated
func springLoadingDestination_SpringLoadingUpdated(hp C.uintptr_t, draggingInfo unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	result := delegate.SpringLoadingUpdated(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingExited
func springLoadingDestination_SpringLoadingExited(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingExited(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_DraggingEnded
func springLoadingDestination_DraggingEnded(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.DraggingEnded(objc.MakeObject(draggingInfo))
}

//export SpringLoadingDestination_RespondsTo
func SpringLoadingDestination_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
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

type DraggingInfo interface {
	SlideDraggedImageTo(screenPoint foundation.Point)
	ResetSpringLoading()
	DraggingPasteboard() Pasteboard
	DraggingSequenceNumber() int
	DraggingSource() objc.Object
	DraggingSourceOperationMask() DragOperation
	DraggingLocation() foundation.Point
	DraggingDestinationWindow() Window
	NumberOfValidItemsForDrop() int
	SetNumberOfValidItemsForDrop(value int)
	DraggedImageLocation() foundation.Point
	AnimatesToDestination() bool
	SetAnimatesToDestination(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	SpringLoadingHighlight() SpringLoadingHighlight
}

func DraggingInfoToObjc(protocol DraggingInfo) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapDraggingInfo(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export draggingInfo_SlideDraggedImageTo
func draggingInfo_SlideDraggedImageTo(hp C.uintptr_t, screenPoint C.CGPoint) {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	protocol.SlideDraggedImageTo(foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingInfo_ResetSpringLoading
func draggingInfo_ResetSpringLoading(hp C.uintptr_t) {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	protocol.ResetSpringLoading()
}

//export draggingInfo_DraggingPasteboard
func draggingInfo_DraggingPasteboard(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingPasteboard()
	return objc.ExtractPtr(result)
}

//export draggingInfo_DraggingSequenceNumber
func draggingInfo_DraggingSequenceNumber(hp C.uintptr_t) C.int {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingSequenceNumber()
	return C.int(result)
}

//export draggingInfo_DraggingSource
func draggingInfo_DraggingSource(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingSource()
	return objc.ExtractPtr(result)
}

//export draggingInfo_DraggingSourceOperationMask
func draggingInfo_DraggingSourceOperationMask(hp C.uintptr_t) C.uint {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingSourceOperationMask()
	return C.uint(uint(result))
}

//export draggingInfo_DraggingLocation
func draggingInfo_DraggingLocation(hp C.uintptr_t) C.CGPoint {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingLocation()
	return *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(result)))
}

//export draggingInfo_DraggingDestinationWindow
func draggingInfo_DraggingDestinationWindow(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingDestinationWindow()
	return objc.ExtractPtr(result)
}

//export draggingInfo_SetNumberOfValidItemsForDrop
func draggingInfo_SetNumberOfValidItemsForDrop(hp C.uintptr_t, value C.int) {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	protocol.SetNumberOfValidItemsForDrop(int(value))
}

//export draggingInfo_NumberOfValidItemsForDrop
func draggingInfo_NumberOfValidItemsForDrop(hp C.uintptr_t) C.int {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.NumberOfValidItemsForDrop()
	return C.int(result)
}

//export draggingInfo_DraggedImageLocation
func draggingInfo_DraggedImageLocation(hp C.uintptr_t) C.CGPoint {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggedImageLocation()
	return *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(result)))
}

//export draggingInfo_SetAnimatesToDestination
func draggingInfo_SetAnimatesToDestination(hp C.uintptr_t, value C.bool) {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	protocol.SetAnimatesToDestination(bool(value))
}

//export draggingInfo_AnimatesToDestination
func draggingInfo_AnimatesToDestination(hp C.uintptr_t) C.bool {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.AnimatesToDestination()
	return C.bool(result)
}

//export draggingInfo_SetDraggingFormation
func draggingInfo_SetDraggingFormation(hp C.uintptr_t, value C.int) {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	protocol.SetDraggingFormation(DraggingFormation(int(value)))
}

//export draggingInfo_DraggingFormation
func draggingInfo_DraggingFormation(hp C.uintptr_t) C.int {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.DraggingFormation()
	return C.int(int(result))
}

//export draggingInfo_SpringLoadingHighlight
func draggingInfo_SpringLoadingHighlight(hp C.uintptr_t) C.int {
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	result := protocol.SpringLoadingHighlight()
	return C.int(int(result))
}

//export DraggingInfo_RespondsTo
func DraggingInfo_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(DraggingInfo)
	_ = protocol
	switch selName {
	case "slideDraggedImageTo:":
		return true
	case "resetSpringLoading":
		return true
	case "draggingPasteboard":
		return true
	case "draggingSequenceNumber":
		return true
	case "draggingSource":
		return true
	case "draggingSourceOperationMask":
		return true
	case "draggingLocation":
		return true
	case "draggingDestinationWindow":
		return true
	case "setNumberOfValidItemsForDrop:":
		fallthrough
	case "numberOfValidItemsForDrop":
		return true
	case "draggedImageLocation":
		return true
	case "setAnimatesToDestination:":
		fallthrough
	case "animatesToDestination":
		return true
	case "setDraggingFormation:":
		fallthrough
	case "draggingFormation":
		return true
	case "springLoadingHighlight":
		return true
	default:
		return false
	}
}
