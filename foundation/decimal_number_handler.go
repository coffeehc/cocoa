package foundation

// #include "decimal_number_handler.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DecimalNumberHandler interface {
	objc.Object
}

type NSDecimalNumberHandler struct {
	objc.NSObject
}

func MakeDecimalNumberHandler(ptr unsafe.Pointer) *NSDecimalNumberHandler {
	if ptr == nil {
		return nil
	}
	return &NSDecimalNumberHandler{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocDecimalNumberHandler() *NSDecimalNumberHandler {
	return MakeDecimalNumberHandler(C.C_DecimalNumberHandler_Alloc())
}

func (n *NSDecimalNumberHandler) Init() DecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_Init(n.Ptr())
	return MakeDecimalNumberHandler(result_)
}

func DefaultDecimalNumberHandler() DecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_DefaultDecimalNumberHandler()
	return MakeDecimalNumberHandler(result_)
}
