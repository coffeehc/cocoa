package objc

// #import "selector.h"
import "C"
import "unsafe"

// Selector for select and hold method
type Selector unsafe.Pointer

// SEL_RegisterName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value
func SEL_RegisterName(name string) Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Selector(C.Selector_SEL_RegisterName(cname))
}

// Sel_GetName return selector name
func Sel_GetName(s Selector) string {
	cstr := C.Selector_SEL_GetName(unsafe.Pointer(s))
	return C.GoString(cstr)
}
