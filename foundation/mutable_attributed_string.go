package foundation

// #include "mutable_attributed_string.h"
import "C"
import (
	"unsafe"
)

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
