package objc

// #import "object.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// PointerHolder is an interface for holding an objc pointer
type PointerHolder interface {
	// Ptr return the delegate objc pointer
	Ptr() unsafe.Pointer
}

// Object is interface for all objc NSObject type
type Object interface {
	PointerHolder
	fmt.Stringer
	// Retain0 call retain, but return no value, to avoiding subtypes return type conflicting
	Retain0()
	Release()
	// Autorelease0 call autorelease, but return no value, to avoiding subtypes return type conflicting
	Autorelease0()

	Copy() Object
	MutableCopy() Object

	Dealloc()
	Description() string
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

func MakeObject(ptr unsafe.Pointer) NSObject {
	return NSObject{ptr}
}

func (o NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func AllocObject() NSObject {
	result_ := C.C_NSObject_AllocObject()
	return MakeObject(result_)
}

func (o NSObject) Init() NSObject {
	result_ := C.C_NSObject_Init(o.Ptr())
	return MakeObject(result_)
}

func NewObject() NSObject {
	result_ := C.C_NSObject_NewObject()
	return MakeObject(result_)
}

func (o NSObject) Copy() Object {
	result_ := C.C_NSObject_Copy(o.Ptr())
	return MakeObject(result_)
}

func (o NSObject) MutableCopy() Object {
	result_ := C.C_NSObject_MutableCopy(o.Ptr())
	return MakeObject(result_)
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

func (o NSObject) Retain0() {
	C.Object_Retain(o.Ptr())
}

func (o NSObject) Retain() NSObject {
	p := C.Object_Retain(o.Ptr())
	return MakeObject(p)
}

func (o NSObject) Release() {
	C.Object_Release(o.Ptr())
}

func (o NSObject) Autorelease0() {
	C.Object_Autorelease(o.Ptr())
}

func (o NSObject) Autorelease() NSObject {
	p := C.Object_Autorelease(o.Ptr())
	return MakeObject(p)
}

func (o NSObject) Dealloc() {
	C.Object_Dealloc(o.Ptr())
}

func (o NSObject) Description() string {
	return C.GoString(C.Object_Description(o.Ptr()))
}

func (o NSObject) String() string {
	return o.Description()
}
