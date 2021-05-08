package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "value.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Value interface {
	objc.Object
	IsEqualToValue(value Value) bool
	NonretainedObjectValue() objc.Object
}

type NSValue struct {
	objc.NSObject
}

func MakeValue(ptr unsafe.Pointer) *NSValue {
	if ptr == nil {
		return nil
	}
	return &NSValue{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocValue() *NSValue {
	return MakeValue(C.C_Value_Alloc())
}

func (n *NSValue) InitWithCoder(coder Coder) Value {
	result := C.C_NSValue_InitWithCoder(n.Ptr(), toPointer(coder))
	return MakeValue(result)
}

func (n *NSValue) Init() Value {
	result := C.C_NSValue_Init(n.Ptr())
	return MakeValue(result)
}

func ValueWithNonretainedObject(anObject objc.Object) Value {
	result := C.C_NSValue_ValueWithNonretainedObject(toPointer(anObject))
	return MakeValue(result)
}

func (n *NSValue) IsEqualToValue(value Value) bool {
	result := C.C_NSValue_IsEqualToValue(n.Ptr(), toPointer(value))
	return bool(result)
}

func (n *NSValue) NonretainedObjectValue() objc.Object {
	result := C.C_NSValue_NonretainedObjectValue(n.Ptr())
	return objc.MakeObject(result)
}
