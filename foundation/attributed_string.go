package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "attributed_string.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AttributedString interface {
	objc.Object
	String() string
	Length() uint
}

type NSAttributedString struct {
	objc.NSObject
}

func MakeAttributedString(ptr unsafe.Pointer) *NSAttributedString {
	if ptr == nil {
		return nil
	}
	return &NSAttributedString{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocAttributedString() *NSAttributedString {
	return MakeAttributedString(C.C_AttributedString_Alloc())
}

func (n *NSAttributedString) Init() AttributedString {
	result := C.C_NSAttributedString_Init(n.Ptr())
	return MakeAttributedString(result)
}

func (n *NSAttributedString) String() string {
	result := C.C_NSAttributedString_String(n.Ptr())
	return MakeString(result).String()
}

func (n *NSAttributedString) Length() uint {
	result := C.C_NSAttributedString_Length(n.Ptr())
	return uint(result)
}
