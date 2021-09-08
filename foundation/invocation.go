package foundation

// #include "invocation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Invocation interface {
	objc.Object
	RetainArguments()
	Invoke()
	InvokeWithTarget(target objc.Object)
	Selector() objc.Selector
	SetSelector(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.Object)
	ArgumentsRetained() bool
	MethodSignature() MethodSignature
}

type NSInvocation struct {
	objc.NSObject
}

func MakeInvocation(ptr unsafe.Pointer) NSInvocation {
	return NSInvocation{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocInvocation() NSInvocation {
	result_ := C.C_NSInvocation_AllocInvocation()
	return MakeInvocation(result_)
}

func (n NSInvocation) Init() NSInvocation {
	result_ := C.C_NSInvocation_Init(n.Ptr())
	return MakeInvocation(result_)
}

func NewInvocation() NSInvocation {
	result_ := C.C_NSInvocation_NewInvocation()
	return MakeInvocation(result_)
}

func (n NSInvocation) Autorelease() NSInvocation {
	result_ := C.C_NSInvocation_Autorelease(n.Ptr())
	return MakeInvocation(result_)
}

func (n NSInvocation) Retain() NSInvocation {
	result_ := C.C_NSInvocation_Retain(n.Ptr())
	return MakeInvocation(result_)
}

func InvocationWithMethodSignature(sig MethodSignature) Invocation {
	result_ := C.C_NSInvocation_InvocationWithMethodSignature(objc.ExtractPtr(sig))
	return MakeInvocation(result_)
}

func (n NSInvocation) RetainArguments() {
	C.C_NSInvocation_RetainArguments(n.Ptr())
}

func (n NSInvocation) Invoke() {
	C.C_NSInvocation_Invoke(n.Ptr())
}

func (n NSInvocation) InvokeWithTarget(target objc.Object) {
	C.C_NSInvocation_InvokeWithTarget(n.Ptr(), objc.ExtractPtr(target))
}

func (n NSInvocation) Selector() objc.Selector {
	result_ := C.C_NSInvocation_Selector(n.Ptr())
	return objc.Selector(result_)
}

func (n NSInvocation) SetSelector(value objc.Selector) {
	C.C_NSInvocation_SetSelector(n.Ptr(), unsafe.Pointer(value))
}

func (n NSInvocation) Target() objc.Object {
	result_ := C.C_NSInvocation_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSInvocation) SetTarget(value objc.Object) {
	C.C_NSInvocation_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSInvocation) ArgumentsRetained() bool {
	result_ := C.C_NSInvocation_ArgumentsRetained(n.Ptr())
	return bool(result_)
}

func (n NSInvocation) MethodSignature() MethodSignature {
	result_ := C.C_NSInvocation_MethodSignature(n.Ptr())
	return MakeMethodSignature(result_)
}
