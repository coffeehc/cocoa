package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
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
}

type NSException struct {
	objc.NSObject
}

func MakeException(ptr unsafe.Pointer) *NSException {
	if ptr == nil {
		return nil
	}
	return &NSException{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocException() *NSException {
	return MakeException(C.C_Exception_Alloc())
}

func (n *NSException) Init() Exception {
	result := C.C_NSException_Init(n.Ptr())
	return MakeException(result)
}

func (n *NSException) Raise() {
	C.C_NSException_Raise(n.Ptr())
}

func (n *NSException) Name() ExceptionName {
	result := C.C_NSException_Name(n.Ptr())
	return ExceptionName(MakeString(result).String())
}

func (n *NSException) Reason() string {
	result := C.C_NSException_Reason(n.Ptr())
	return MakeString(result).String()
}
