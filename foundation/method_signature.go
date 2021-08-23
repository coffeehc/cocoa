package foundation

// #include "method_signature.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MethodSignature interface {
	objc.Object
	IsOneway() bool
	NumberOfArguments() uint
	FrameLength() uint
	MethodReturnLength() uint
}

type NSMethodSignature struct {
	objc.NSObject
}

func MakeMethodSignature(ptr unsafe.Pointer) NSMethodSignature {
	return NSMethodSignature{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocMethodSignature() NSMethodSignature {
	return MakeMethodSignature(C.C_MethodSignature_Alloc())
}

func (n NSMethodSignature) Init() MethodSignature {
	result_ := C.C_NSMethodSignature_Init(n.Ptr())
	return MakeMethodSignature(result_)
}

func (n NSMethodSignature) IsOneway() bool {
	result_ := C.C_NSMethodSignature_IsOneway(n.Ptr())
	return bool(result_)
}

func (n NSMethodSignature) NumberOfArguments() uint {
	result_ := C.C_NSMethodSignature_NumberOfArguments(n.Ptr())
	return uint(result_)
}

func (n NSMethodSignature) FrameLength() uint {
	result_ := C.C_NSMethodSignature_FrameLength(n.Ptr())
	return uint(result_)
}

func (n NSMethodSignature) MethodReturnLength() uint {
	result_ := C.C_NSMethodSignature_MethodReturnLength(n.Ptr())
	return uint(result_)
}
