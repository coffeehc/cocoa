package objc

// #import "object.h"
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// PointerHolder is a interface for holding a objc pointer
type PointerHolder interface {
	// Ptr return the delegate objc pointer
	Ptr() unsafe.Pointer
}

// Object is interface for all objc NSObject type
type Object interface {
	PointerHolder
	Dealloc()
}

// ExtractPtr return the objc ptr hold by Object. If is nil, or contains a nil pointer, return nil
func ExtractPtr(o PointerHolder) unsafe.Pointer {
	if o == nil {
		return nil
	}
	return o.Ptr()
}

// NSObject is wrapper for objc-NSObject
type NSObject struct {
	ptr unsafe.Pointer
}

func (o NSObject) PerformSelector(sel Selector) Object {
	rp := C.Object_PerformSelector(o.Ptr(), unsafe.Pointer(sel))
	return MakeObject(rp)
}

func (o NSObject) PerformSelector_WithObject(sel Selector, object Object) Object {
	var param unsafe.Pointer
	if object != nil {
		param = object.Ptr()
	} else {
		param = nil
	}
	rp := C.Object_PerformSelector_WithObject(o.Ptr(), unsafe.Pointer(sel), param)
	return MakeObject(rp)
}

func (o NSObject) PerformSelector_WithObject_WithObject(sel Selector, obj1, obj2 Object) Object {
	var param1, param2 unsafe.Pointer
	if obj1 != nil {
		param1 = obj1.Ptr()
	} else {
		param1 = nil
	}
	if obj2 != nil {
		param2 = obj2.Ptr()
	} else {
		param2 = nil
	}
	rp := C.Object_PerformSelector_WithObject_WithObject(o.Ptr(), unsafe.Pointer(sel), param1, param2)
	return MakeObject(rp)
}

func (o NSObject) Dealloc() {
	if o.Ptr() == nil {
		panic("objc object is null")
	}
	C.Object_Dealloc(o.Ptr())
}

func (o NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func MakeObject(ptr unsafe.Pointer) NSObject {
	return NSObject{ptr}
}

// AddDeallocHook add cocoa object dealloc hook
func AddDeallocHook(obj Object, hook func()) {
	if obj.Ptr() == nil {
		panic("cocoa pointer is nil")
	}
	h := cgo.NewHandle(hook)
	C.Dealloc_AddHook(obj.Ptr(), C.uintptr_t(h))
}

//export runDeallocTask
func runDeallocTask(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}
