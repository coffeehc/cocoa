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

func MakeUnit(ptr unsafe.Pointer) *NSUnit {
	if ptr == nil {
		return nil
	}
	return &NSUnit{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUnit() *NSUnit {
	return MakeUnit(C.C_Unit_Alloc())
}

func (n *NSUnit) InitWithSymbol(symbol string) Unit {
	result_ := C.C_NSUnit_InitWithSymbol(n.Ptr(), NewString(symbol).Ptr())
	return MakeUnit(result_)
}

func (n *NSUnit) Symbol() string {
	result_ := C.C_NSUnit_Symbol(n.Ptr())
	return MakeString(result_).String()
}
