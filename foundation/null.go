package foundation

// #include "null.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Null interface {
	objc.Object
}

type NSNull struct {
	objc.NSObject
}

func MakeNull(ptr unsafe.Pointer) NSNull {
	return NSNull{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocNull() NSNull {
	result_ := C.C_NSNull_AllocNull()
	return MakeNull(result_)
}

func (n NSNull) Init() NSNull {
	result_ := C.C_NSNull_Init(n.Ptr())
	return MakeNull(result_)
}

func NewNull() NSNull {
	result_ := C.C_NSNull_NewNull()
	return MakeNull(result_)
}

func (n NSNull) Autorelease() NSNull {
	result_ := C.C_NSNull_Autorelease(n.Ptr())
	return MakeNull(result_)
}

func (n NSNull) Retain() NSNull {
	result_ := C.C_NSNull_Retain(n.Ptr())
	return MakeNull(result_)
}

func Null_() Null {
	result_ := C.C_NSNull_Null_()
	return MakeNull(result_)
}
