package appkit

// #include "dragging_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DraggingSource struct {
	DraggingSession_SourceOperationMaskForDraggingContext func(session DraggingSession, context DraggingContext) DragOperation // required
	DraggingSession_EndedAtPoint_Operation                func(session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	DraggingSession_MovedToPoint                          func(session DraggingSession, screenPoint foundation.Point)
	DraggingSession_WillBeginAtPoint                      func(session DraggingSession, screenPoint foundation.Point)
	IgnoreModifierKeysForDraggingSession                  func(session DraggingSession) bool
}

func WrapDraggingSource(delegate *DraggingSource) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapDraggingSource(C.long(id))
	return objc.MakeObject(ptr)
}

//export draggingSource_DraggingSession_SourceOperationMaskForDraggingContext
func draggingSource_DraggingSession_SourceOperationMaskForDraggingContext(id int64, session unsafe.Pointer, context C.int) C.uint {
	delegate := resources.Get(id).(*DraggingSource)
	result := delegate.DraggingSession_SourceOperationMaskForDraggingContext(MakeDraggingSession(session), DraggingContext(int(context)))
	return C.uint(uint(result))
}

//export draggingSource_DraggingSession_EndedAtPoint_Operation
func draggingSource_DraggingSession_EndedAtPoint_Operation(id int64, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := resources.Get(id).(*DraggingSource)
	delegate.DraggingSession_EndedAtPoint_Operation(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export draggingSource_DraggingSession_MovedToPoint
func draggingSource_DraggingSession_MovedToPoint(id int64, session unsafe.Pointer, screenPoint C.CGPoint) {
	delegate := resources.Get(id).(*DraggingSource)
	delegate.DraggingSession_MovedToPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_DraggingSession_WillBeginAtPoint
func draggingSource_DraggingSession_WillBeginAtPoint(id int64, session unsafe.Pointer, screenPoint C.CGPoint) {
	delegate := resources.Get(id).(*DraggingSource)
	delegate.DraggingSession_WillBeginAtPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_IgnoreModifierKeysForDraggingSession
func draggingSource_IgnoreModifierKeysForDraggingSession(id int64, session unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*DraggingSource)
	result := delegate.IgnoreModifierKeysForDraggingSession(MakeDraggingSession(session))
	return C.bool(result)
}

//export DraggingSource_RespondsTo
func DraggingSource_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*DraggingSource)
	switch selName {
	case "draggingSession:sourceOperationMaskForDraggingContext:":
		return delegate.DraggingSession_SourceOperationMaskForDraggingContext != nil
	case "draggingSession:endedAtPoint:operation:":
		return delegate.DraggingSession_EndedAtPoint_Operation != nil
	case "draggingSession:movedToPoint:":
		return delegate.DraggingSession_MovedToPoint != nil
	case "draggingSession:willBeginAtPoint:":
		return delegate.DraggingSession_WillBeginAtPoint != nil
	case "ignoreModifierKeysForDraggingSession:":
		return delegate.IgnoreModifierKeysForDraggingSession != nil
	default:
		return false
	}
}

//export deleteDraggingSource
func deleteDraggingSource(id int64) {
	resources.Delete(id)
}
