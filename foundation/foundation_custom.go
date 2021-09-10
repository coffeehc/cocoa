package foundation

// #include "foundation_custom.h"
import "C"

import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

// SelectorFromString returns the selector with a given name.
func SelectorFromString(name string) objc.Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return objc.Selector(C.Selector_SelectorFromString(cname))
}

// StringFromSelector returns a string representation of a given selector
func StringFromSelector(selector objc.Selector) string {
	return C.GoString(C.Selector_StringFromSelector(unsafe.Pointer(selector)))
}

// String wrap NSString
type String interface {
	objc.Object
	String() string
}

var _ String = (*NSString)(nil)

type NSString struct {
	objc.NSObject
}

func (N NSString) String() string {
	return C.GoString(C.String_Value(N.Ptr()))
}

func MakeString(ptr unsafe.Pointer) NSString {
	return NSString{
		NSObject: objc.MakeObject(ptr),
	}
}

func NewString(str string) NSString {
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

func MakeMutableString(ptr unsafe.Pointer) NSMutableString {
	return NSMutableString{
		NSString: MakeString(ptr),
	}
}

// Data wrap NSData
type Data interface {
	objc.Object
	ToBytes() []byte
}

type NSData struct {
	objc.NSObject
}

func MakeData(ptr unsafe.Pointer) NSData {
	return NSData{
		NSObject: objc.MakeObject(ptr),
	}
}

func NewData(bytes []byte) NSData {
	size := len(bytes)
	var p unsafe.Pointer
	if size == 0 {
		p = C.Data_New(unsafe.Pointer(nil), 0)
	} else {
		p = C.Data_New(unsafe.Pointer(&bytes[0]), C.int(size))
	}
	return MakeData(p)
}

func (d NSData) ToBytes() []byte {
	if d.Ptr() == nil {
		return nil
	}
	array := C.Data_ToBytes(d.Ptr())
	size := int(array.len)
	if size <= 0 {
		return nil
	}
	bytes := unsafe.Slice((*byte)(array.data), size)
	newBytes := make([]byte, size)
	copy(newBytes, bytes)
	return newBytes
}
