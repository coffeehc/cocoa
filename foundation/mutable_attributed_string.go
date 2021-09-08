package foundation

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

func MakeMutableAttributedString(ptr unsafe.Pointer) NSMutableAttributedString {
	return NSMutableAttributedString{
		NSAttributedString: MakeAttributedString(ptr),
	}
}

func (n NSMutableAttributedString) InitWithString(str string) NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_InitWithString(n.Ptr(), NewString(str).Ptr())
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) InitWithString_Attributes(str string, attrs map[AttributedStringKey]objc.Object) NSMutableAttributedString {
	var cAttrs C.Dictionary
	if len(attrs) > 0 {
		cAttrsKeyData := make([]unsafe.Pointer, len(attrs))
		cAttrsValueData := make([]unsafe.Pointer, len(attrs))
		var idx = 0
		for k, v := range attrs {
			cAttrsKeyData[idx] = NewString(string(k)).Ptr()
			cAttrsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cAttrs.key_data = unsafe.Pointer(&cAttrsKeyData[0])
		cAttrs.value_data = unsafe.Pointer(&cAttrsValueData[0])
		cAttrs.len = C.int(len(attrs))
	}
	result_ := C.C_NSMutableAttributedString_InitWithString_Attributes(n.Ptr(), NewString(str).Ptr(), cAttrs)
	return MakeMutableAttributedString(result_)
}

func (n NSMutableAttributedString) InitWithAttributedString(attrStr AttributedString) NSMutableAttributedString {
	result_ := C.C_NSMutableAttributedString_InitWithAttributedString(n.Ptr(), objc.ExtractPtr(attrStr))
	return MakeMutableAttributedString(result_)
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
