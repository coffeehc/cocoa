package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "responder.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Responder interface {
	objc.Object
	AcceptsFirstResponder() bool
	NextResponder() Responder
	SetNextResponder(nextResponder Responder)
	BecomeFirstResponder() bool
}

var _ Responder = (*NSResponder)(nil)

type NSResponder struct {
	objc.NSObject
}

func MakeResponder(ptr unsafe.Pointer) *NSResponder {
	if ptr == nil {
		return nil
	}
	return &NSResponder{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (r *NSResponder) AcceptsFirstResponder() bool {
	return bool(C.Responder_AcceptsFirstResponder(r.Ptr()))
}

func (r *NSResponder) NextResponder() Responder {
	return MakeResponder(C.Responder_NextResponder(r.Ptr()))
}

func (r *NSResponder) SetNextResponder(nextResponder Responder) {
	C.Responder_SetNextResponder(r.Ptr(), toPointer(nextResponder))
}

func (r *NSResponder) BecomeFirstResponder() bool {
	return bool(C.Responder_BecomeFirstResponder(r.Ptr()))
}
