package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "error.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Error interface {
	objc.Object
	Code() int
	Domain() ErrorDomain
	LocalizedDescription() string
	LocalizedRecoverySuggestion() string
	LocalizedFailureReason() string
	RecoveryAttempter() objc.Object
	HelpAnchor() string
}

type NSError struct {
	objc.NSObject
}

func MakeError(ptr unsafe.Pointer) *NSError {
	if ptr == nil {
		return nil
	}
	return &NSError{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocError() *NSError {
	return MakeError(C.C_Error_Alloc())
}

func (n *NSError) Init() Error {
	result := C.C_NSError_Init(n.Ptr())
	return MakeError(result)
}

func (n *NSError) Code() int {
	result := C.C_NSError_Code(n.Ptr())
	return int(result)
}

func (n *NSError) Domain() ErrorDomain {
	result := C.C_NSError_Domain(n.Ptr())
	return ErrorDomain(MakeString(result).String())
}

func (n *NSError) LocalizedDescription() string {
	result := C.C_NSError_LocalizedDescription(n.Ptr())
	return MakeString(result).String()
}

func (n *NSError) LocalizedRecoverySuggestion() string {
	result := C.C_NSError_LocalizedRecoverySuggestion(n.Ptr())
	return MakeString(result).String()
}

func (n *NSError) LocalizedFailureReason() string {
	result := C.C_NSError_LocalizedFailureReason(n.Ptr())
	return MakeString(result).String()
}

func (n *NSError) RecoveryAttempter() objc.Object {
	result := C.C_NSError_RecoveryAttempter(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSError) HelpAnchor() string {
	result := C.C_NSError_HelpAnchor(n.Ptr())
	return MakeString(result).String()
}
