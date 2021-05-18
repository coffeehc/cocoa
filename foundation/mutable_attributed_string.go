package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "mutable_attributed_string.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MutableAttributedString interface {
	AttributedString
	MutableString() MutableString
}

type NSMutableAttributedString struct {
	NSAttributedString
}

func MakeMutableAttributedString(ptr unsafe.Pointer) *NSMutableAttributedString {
	if ptr == nil {
		return nil
	}
	return &NSMutableAttributedString{
		NSAttributedString: *MakeAttributedString(ptr),
	}
}

func AllocMutableAttributedString() *NSMutableAttributedString {
	return MakeMutableAttributedString(C.C_MutableAttributedString_Alloc())
}

func (n *NSMutableAttributedString) InitWithString(str string) MutableAttributedString {
	result_ := C.C_NSMutableAttributedString_InitWithString(n.Ptr(), NewString(str).Ptr())
	return MakeMutableAttributedString(result_)
}

func (n *NSMutableAttributedString) InitWithAttributedString(attrStr AttributedString) MutableAttributedString {
	result_ := C.C_NSMutableAttributedString_InitWithAttributedString(n.Ptr(), objc.ExtractPtr(attrStr))
	return MakeMutableAttributedString(result_)
}

func (n *NSMutableAttributedString) Init() MutableAttributedString {
	result_ := C.C_NSMutableAttributedString_Init(n.Ptr())
	return MakeMutableAttributedString(result_)
}

func (n *NSMutableAttributedString) MutableString() MutableString {
	result_ := C.C_NSMutableAttributedString_MutableString(n.Ptr())
	return MakeMutableString(result_)
}
