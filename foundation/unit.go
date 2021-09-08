package foundation

// #include "unit.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Unit interface {
	objc.Object
	Symbol() string
}

type NSUnit struct {
	objc.NSObject
}

func MakeUnit(ptr unsafe.Pointer) NSUnit {
	return NSUnit{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSUnit) InitWithSymbol(symbol string) NSUnit {
	result_ := C.C_NSUnit_InitWithSymbol(n.Ptr(), NewString(symbol).Ptr())
	return MakeUnit(result_)
}

func AllocUnit() NSUnit {
	result_ := C.C_NSUnit_AllocUnit()
	return MakeUnit(result_)
}

func (n NSUnit) Autorelease() NSUnit {
	result_ := C.C_NSUnit_Autorelease(n.Ptr())
	return MakeUnit(result_)
}

func (n NSUnit) Retain() NSUnit {
	result_ := C.C_NSUnit_Retain(n.Ptr())
	return MakeUnit(result_)
}

func (n NSUnit) Symbol() string {
	result_ := C.C_NSUnit_Symbol(n.Ptr())
	return MakeString(result_).String()
}
