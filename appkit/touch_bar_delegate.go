package appkit

// #include "touch_bar_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TouchBarDelegate struct {
	TouchBar_MakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem
}

func WrapTouchBarDelegate(delegate *TouchBarDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapTouchBarDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export touchBarDelegate_TouchBar_MakeItemForIdentifier
func touchBarDelegate_TouchBar_MakeItemForIdentifier(id int64, touchBar unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*TouchBarDelegate)
	result := delegate.TouchBar_MakeItemForIdentifier(MakeTouchBar(touchBar), TouchBarItemIdentifier(foundation.MakeString(identifier).String()))
	return objc.ExtractPtr(result)
}

//export TouchBarDelegate_RespondsTo
func TouchBarDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*TouchBarDelegate)
	switch selName {
	case "touchBar:makeItemForIdentifier:":
		return delegate.TouchBar_MakeItemForIdentifier != nil
	default:
		return false
	}
}

//export deleteTouchBarDelegate
func deleteTouchBarDelegate(id int64) {
	resources.Delete(id)
}
