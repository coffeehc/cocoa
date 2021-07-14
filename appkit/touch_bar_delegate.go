package appkit

// #include "touch_bar_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TouchBarDelegate struct {
	TouchBar_MakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) TouchBarItem
}

func WrapTouchBarDelegate(delegate *TouchBarDelegate) objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTouchBarDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export touchBarDelegate_TouchBar_MakeItemForIdentifier
func touchBarDelegate_TouchBar_MakeItemForIdentifier(hp C.uintptr_t, touchBar unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TouchBarDelegate)
	result := delegate.TouchBar_MakeItemForIdentifier(MakeTouchBar(touchBar), TouchBarItemIdentifier(foundation.MakeString(identifier).String()))
	return objc.ExtractPtr(result)
}

//export TouchBarDelegate_RespondsTo
func TouchBarDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TouchBarDelegate)
	switch selName {
	case "touchBar:makeItemForIdentifier:":
		return delegate.TouchBar_MakeItemForIdentifier != nil
	default:
		return false
	}
}

//export deleteTouchBarDelegate
func deleteTouchBarDelegate(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
