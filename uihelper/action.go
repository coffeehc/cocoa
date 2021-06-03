package uihelper

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "action.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"sync"
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

var actionsLock sync.RWMutex
var actionId int64 = 0
var actionsMap = map[int64]ActionHandler{}

// NewAction create a new objc ActionTarget instance, from ActionHandler
func NewAction(handler ActionHandler) ActionTarget {
	if handler == nil {
		panic("handler is nil")
	}
	actionsLock.Lock()
	actionId++
	id := actionId
	actionsMap[id] = handler
	actionsLock.Unlock()
	return ActionTarget{
		NSObject: objc.MakeObject(C.C_NewAction(C.long(id))),
	}
}

// SetAction set action for an ojbc instance, if it has hasTarget and setAction method.
func SetAction(instance CanSetAction, handler ActionHandler) {
	target := NewAction(handler)
	instance.SetTarget(target)
	instance.SetAction(foundation.SelectorFromString("onAction:"))
}

//export callAction
func callAction(id int64, senderPtr unsafe.Pointer) {
	actionsLock.RLock()
	handler := actionsMap[id]
	actionsLock.RUnlock()
	if handler == nil {
		panic("handler not found")
	}
	handler(objc.MakeObject(senderPtr))
}

//export deleteAction
func deleteAction(id int64) {
	actionsLock.Lock()
	delete(actionsMap, id)
	actionsLock.Unlock()
}
