package appkit

// #include "dragging_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type DraggingSource struct {
	DraggingSession_SourceOperationMaskForDraggingContext func(session DraggingSession, context DraggingContext) DragOperation // required
	DraggingSession_EndedAtPoint_Operation                func(session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	DraggingSession_MovedToPoint                          func(session DraggingSession, screenPoint foundation.Point)
	DraggingSession_WillBeginAtPoint                      func(session DraggingSession, screenPoint foundation.Point)
	IgnoreModifierKeysForDraggingSession                  func(session DraggingSession) bool
}

func (delegate *DraggingSource) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapDraggingSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export draggingSource_DraggingSession_SourceOperationMaskForDraggingContext
func draggingSource_DraggingSession_SourceOperationMaskForDraggingContext(hp C.uintptr_t, session unsafe.Pointer, context C.int) C.uint {
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
	result := delegate.DraggingSession_SourceOperationMaskForDraggingContext(MakeDraggingSession(session), DraggingContext(int(context)))
	return C.uint(uint(result))
}

//export draggingSource_DraggingSession_EndedAtPoint_Operation
func draggingSource_DraggingSession_EndedAtPoint_Operation(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
	delegate.DraggingSession_EndedAtPoint_Operation(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export draggingSource_DraggingSession_MovedToPoint
func draggingSource_DraggingSession_MovedToPoint(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint) {
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
	delegate.DraggingSession_MovedToPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_DraggingSession_WillBeginAtPoint
func draggingSource_DraggingSession_WillBeginAtPoint(hp C.uintptr_t, session unsafe.Pointer, screenPoint C.CGPoint) {
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
	delegate.DraggingSession_WillBeginAtPoint(MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))))
}

//export draggingSource_IgnoreModifierKeysForDraggingSession
func draggingSource_IgnoreModifierKeysForDraggingSession(hp C.uintptr_t, session unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
	result := delegate.IgnoreModifierKeysForDraggingSession(MakeDraggingSession(session))
	return C.bool(result)
}

//export DraggingSource_RespondsTo
func DraggingSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*DraggingSource)
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
