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
	result_ := C.C_NSAttributedString_Init(n.Ptr())
	return MakeAttributedString(result_)
}

func (n *NSAttributedString) String() string {
	result_ := C.C_NSAttributedString_String(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSAttributedString) Length() uint {
	result_ := C.C_NSAttributedString_Length(n.Ptr())
	return uint(result_)
}

func AttributedString_TextTypes() []string {
	result_ := C.C_NSAttributedString_AttributedString_TextTypes()
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func AttributedString_TextUnfilteredTypes() []string {
	result_ := C.C_NSAttributedString_AttributedString_TextUnfilteredTypes()
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
