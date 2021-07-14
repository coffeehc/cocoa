package uihelper

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "action.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

// ActionHandler is a callback function for an ActionTarget.
type ActionHandler func(sender objc.Object)

// CanSetAction is an interface for objc instance which can set a target and action
type CanSetAction interface {
	SetTarget(target objc.Object)
	SetAction(sel objc.Selector)
}

// ActionTarget hold an objc ActionTarget target instance
type ActionTarget struct {
	objc.NSObject
}

// NewAction create a new objc ActionTarget instance, from ActionHandler
func NewAction(handler ActionHandler) ActionTarget {
	if handler == nil {
		panic("handler is nil")
	}
	h := cgo.NewHandle(handler)
	return ActionTarget{
		NSObject: objc.MakeObject(C.C_NewAction(C.uintptr_t(h))),
	}
}

// SetAction set action for an ojbc instance, if it has hasTarget and setAction method.
func SetAction(instance CanSetAction, handler ActionHandler) {
	target := NewAction(handler)
	instance.SetTarget(target)
	instance.SetAction(foundation.SelectorFromString("onAction:"))
}

//export callAction
func callAction(hp C.uintptr_t, senderPtr unsafe.Pointer) {
	h := cgo.Handle(hp)
	handler := h.Value().(ActionHandler)
	handler(objc.MakeObject(senderPtr))
}

//export deleteAction
func deleteAction(hp C.uintptr_t) {
	h := cgo.Handle(hp)
	h.Delete()
}
