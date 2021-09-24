package foundation

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

func MakeAttributedString(ptr unsafe.Pointer) NSAttributedString {
	return NSAttributedString{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAttributedString() NSAttributedString {
	result_ := C.C_NSAttributedString_AllocAttributedString()
	return MakeAttributedString(result_)
}

func (n NSAttributedString) Init() NSAttributedString {
	result_ := C.C_NSAttributedString_Init(n.Ptr())
	return MakeAttributedString(result_)
}

func NewAttributedString() NSAttributedString {
	result_ := C.C_NSAttributedString_NewAttributedString()
	return MakeAttributedString(result_)
}

func (n NSAttributedString) Autorelease() NSAttributedString {
	result_ := C.C_NSAttributedString_Autorelease(n.Ptr())
	return MakeAttributedString(result_)
}

func (n NSAttributedString) Retain() NSAttributedString {
	result_ := C.C_NSAttributedString_Retain(n.Ptr())
	return MakeAttributedString(result_)
}

func (n NSAttributedString) String() string {
	result_ := C.C_NSAttributedString_String(n.Ptr())
	return MakeString(result_).String()
}

func (n NSAttributedString) Length() uint {
	result_ := C.C_NSAttributedString_Length(n.Ptr())
	return uint(result_)
}

type MutableAttributedString interface {
	AttributedString
	MutableString() MutableString
}

type NSMutableAttributedString struct {
	NSAttributedString
}

func MakeMutableAttributedString(ptr unsafe.Pointer) NSMutableAttributedString {
	return NSMutableAttributedString{
		NSAttributedString: MakeAttributedString(ptr),
	}
}

func AllocMutableAttributedString() NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_AllocMutableAttributedString()
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) Init() NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_Init(n.Ptr())
	return MakeMutableAttributedString(result_)
}

func NewMutableAttributedString() NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_NewMutableAttributedString()
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) Autorelease() NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_Autorelease(n.Ptr())
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) Retain() NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_Retain(n.Ptr())
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) MutableString() MutableString {
	result_ := C.C_NSMutableAttributedString_MutableString(n.Ptr())
	return MakeMutableString(result_)
}
