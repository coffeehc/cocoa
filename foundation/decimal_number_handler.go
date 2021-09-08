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

func MakeDecimalNumberHandler(ptr unsafe.Pointer) NSDecimalNumberHandler {
	return NSDecimalNumberHandler{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDecimalNumberHandler() NSDecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_AllocDecimalNumberHandler()
	return MakeDecimalNumberHandler(result_)
}

func (n NSDecimalNumberHandler) Init() NSDecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_Init(n.Ptr())
	return MakeDecimalNumberHandler(result_)
}

func NewDecimalNumberHandler() NSDecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_NewDecimalNumberHandler()
	return MakeDecimalNumberHandler(result_)
}

func (n NSDecimalNumberHandler) Autorelease() NSDecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_Autorelease(n.Ptr())
	return MakeDecimalNumberHandler(result_)
}

func (n NSDecimalNumberHandler) Retain() NSDecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_Retain(n.Ptr())
	return MakeDecimalNumberHandler(result_)
}

func DefaultDecimalNumberHandler() DecimalNumberHandler {
	result_ := C.C_NSDecimalNumberHandler_DefaultDecimalNumberHandler()
	return MakeDecimalNumberHandler(result_)
}
