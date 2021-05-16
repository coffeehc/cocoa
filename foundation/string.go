package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "string.h"
import "C"

import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type String interface {
	objc.Object
	String() string
}

var _ String = (*NSString)(nil)

type NSString struct {
	objc.NSObject
}

func (N *NSString) String() string {
	return C.GoString(C.String_Value(N.Ptr()))
}

func MakeString(ptr unsafe.Pointer) *NSString {
	return &NSString{
		NSObject: *objc.MakeObject(ptr),
	}
}

func NewString(str string) *NSString {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	ptr := C.String_New(cstr)
	return MakeString(ptr)
}

type MutableString interface {
	String
}

type NSMutableString struct {
	NSString
}

func MakeMutableString(ptr unsafe.Pointer) *NSMutableString {
	return &NSMutableString{
		NSString: *MakeString(ptr),
	}
}
