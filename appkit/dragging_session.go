package appkit

// #include "dragging_session.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DraggingSession interface {
	objc.Object
	DraggingPasteboard() Pasteboard
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingSequenceNumber() int
	DraggingLocation() foundation.Point
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
}

type NSDraggingSession struct {
	objc.NSObject
}

func MakeDraggingSession(ptr unsafe.Pointer) NSDraggingSession {
	return NSDraggingSession{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDraggingSession() NSDraggingSession {
	result_ := C.C_NSDraggingSession_AllocDraggingSession()
	return MakeDraggingSession(result_)
}

func (n NSDraggingSession) Init() NSDraggingSession {
	result_ := C.C_NSDraggingSession_Init(n.Ptr())
	return MakeDraggingSession(result_)
}

func NewDraggingSession() NSDraggingSession {
	result_ := C.C_NSDraggingSession_NewDraggingSession()
	return MakeDraggingSession(result_)
}

func (n NSDraggingSession) Autorelease() NSDraggingSession {
	result_ := C.C_NSDraggingSession_Autorelease(n.Ptr())
	return MakeDraggingSession(result_)
}

func (n NSDraggingSession) Retain() NSDraggingSession {
	result_ := C.C_NSDraggingSession_Retain(n.Ptr())
	return MakeDraggingSession(result_)
}

func (n NSDraggingSession) DraggingPasteboard() Pasteboard {
	result_ := C.C_NSDraggingSession_DraggingPasteboard(n.Ptr())
	return MakePasteboard(result_)
}

func (n NSDraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	result_ := C.C_NSDraggingSession_AnimatesToStartingPositionsOnCancelOrFail(n.Ptr())
	return bool(result_)
}

func (n NSDraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	C.C_NSDraggingSession_SetAnimatesToStartingPositionsOnCancelOrFail(n.Ptr(), C.bool(value))
}

func (n NSDraggingSession) DraggingFormation() DraggingFormation {
	result_ := C.C_NSDraggingSession_DraggingFormation(n.Ptr())
	return DraggingFormation(int(result_))
}

func (n NSDraggingSession) SetDraggingFormation(value DraggingFormation) {
	C.C_NSDraggingSession_SetDraggingFormation(n.Ptr(), C.int(int(value)))
}

func (n NSDraggingSession) DraggingSequenceNumber() int {
	result_ := C.C_NSDraggingSession_DraggingSequenceNumber(n.Ptr())
	return int(result_)
}

func (n NSDraggingSession) DraggingLocation() foundation.Point {
	result_ := C.C_NSDraggingSession_DraggingLocation(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSDraggingSession) DraggingLeaderIndex() int {
	result_ := C.C_NSDraggingSession_DraggingLeaderIndex(n.Ptr())
	return int(result_)
}

func (n NSDraggingSession) SetDraggingLeaderIndex(value int) {
	C.C_NSDraggingSession_SetDraggingLeaderIndex(n.Ptr(), C.int(value))
}
