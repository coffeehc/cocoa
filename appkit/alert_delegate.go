package appkit

// #include "alert_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AlertDelegate struct {
	AlertShowHelp func(alert Alert) bool
}

func WrapAlertDelegate(delegate *AlertDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapAlertDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export alertDelegate_AlertShowHelp
func alertDelegate_AlertShowHelp(id int64, alert unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*AlertDelegate)
	result := delegate.AlertShowHelp(MakeAlert(alert))
	return C.bool(result)
}

//export AlertDelegate_RespondsTo
func AlertDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*AlertDelegate)
	switch selName {
	case "alertShowHelp:":
		return delegate.AlertShowHelp != nil
	default:
		return false
	}
}

//export deleteAlertDelegate
func deleteAlertDelegate(id int64) {
	resources.Delete(id)
}
