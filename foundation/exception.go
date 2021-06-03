package foundation

// #include "exception.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Exception interface {
	objc.Object
	Raise()
	Name() ExceptionName
	Reason() string
	CallStackReturnAddresses() []Number
	CallStackSymbols() []string
}

type NSException struct {
	objc.NSObject
}

func MakeException(ptr unsafe.Pointer) NSException {
	return NSException{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocException() NSException {
	return MakeException(C.C_Exception_Alloc())
}

func (n NSException) Init() Exception {
	result_ := C.C_NSException_Init(n.Ptr())
	return MakeException(result_)
}

func (n NSException) Raise() {
	C.C_NSException_Raise(n.Ptr())
}

func (n NSException) Name() ExceptionName {
	result_ := C.C_NSException_Name(n.Ptr())
	return ExceptionName(MakeString(result_).String())
}

func (n NSException) Reason() string {
	result_ := C.C_NSException_Reason(n.Ptr())
	return MakeString(result_).String()
}

func (n NSException) CallStackReturnAddresses() []Number {
	result_ := C.C_NSException_CallStackReturnAddresses(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Number, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeNumber(r)
	}
	return goResult_
}

func (n NSException) CallStackSymbols() []string {
	result_ := C.C_NSException_CallStackSymbols(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
