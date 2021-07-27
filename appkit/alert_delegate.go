package appkit

// #include "alert_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type AlertDelegate struct {
	AlertShowHelp func(alert Alert) bool
}

func (delegate *AlertDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapAlertDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export alertDelegate_AlertShowHelp
func alertDelegate_AlertShowHelp(hp C.uintptr_t, alert unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*AlertDelegate)
	result := delegate.AlertShowHelp(MakeAlert(alert))
	return C.bool(result)
}

//export AlertDelegate_RespondsTo
func AlertDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*AlertDelegate)
	switch selName {
	case "alertShowHelp:":
		return delegate.AlertShowHelp != nil
	default:
		return false
	}
}

//export deleteAlertDelegate
func deleteAlertDelegate(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
