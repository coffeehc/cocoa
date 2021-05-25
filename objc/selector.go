package objc

// #import "selector.h"
import "C"
import "unsafe"

// Selector for select and hold method
type Selector struct {
	delegate unsafe.Pointer
}

func MakeSelector(ptr unsafe.Pointer) *Selector {
	return &Selector{
		delegate: ptr,
	}
}

// Ptr return delegate objc pointer
func (s *Selector) Ptr() unsafe.Pointer {
	return s.delegate
}

// SEL_RegisterName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value
func SEL_RegisterName(name string) *Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return MakeSelector(C.Selector_SEL_RegisterName(cname))
}

// Sel_GetName return selector name
func Sel_GetName(s *Selector) string {
	cstr := C.Selector_SEL_GetName(s.Ptr())
	return C.GoString(cstr)
}

// An opaque type that represents an Objective-C class.
type Class struct {
	delegate unsafe.Pointer
}

func MakeClass(ptr unsafe.Pointer) *Class {
	return &Class{
		delegate: ptr,
	}
}

type Method struct {
	delegate unsafe.Pointer
}

func MakeMethod(ptr unsafe.Pointer) *Method {
	return &Method{
		delegate: ptr,
	}
}
