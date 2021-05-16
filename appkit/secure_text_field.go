package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "secure_text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SecureTextField interface {
	TextField
}

type NSSecureTextField struct {
	NSTextField
}

func MakeSecureTextField(ptr unsafe.Pointer) *NSSecureTextField {
	if ptr == nil {
		return nil
	}
	return &NSSecureTextField{
		NSTextField: *MakeTextField(ptr),
	}
}

func AllocSecureTextField() *NSSecureTextField {
	return MakeSecureTextField(C.C_SecureTextField_Alloc())
}

func (n *NSSecureTextField) InitWithFrame(frameRect foundation.Rect) SecureTextField {
	result_ := C.C_NSSecureTextField_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSecureTextField(result_)
}

func (n *NSSecureTextField) InitWithCoder(coder foundation.Coder) SecureTextField {
	result_ := C.C_NSSecureTextField_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSecureTextField(result_)
}

func (n *NSSecureTextField) Init() SecureTextField {
	result_ := C.C_NSSecureTextField_Init(n.Ptr())
	return MakeSecureTextField(result_)
}
