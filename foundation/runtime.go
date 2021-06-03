package foundation

// #import "runtime.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

// SelectorFromString returns the selector with a given name.
func SelectorFromString(name string) objc.Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return objc.Selector(C.Selector_SelectorFromString(cname))
}

// StringFromSelector returns a string representation of a given selector
func StringFromSelector(selector objc.Selector) string {
	return C.GoString(C.Selector_StringFromSelector(unsafe.Pointer(selector)))
}
