package appkit

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

func MakeSecureTextField(ptr unsafe.Pointer) NSSecureTextField {
	return NSSecureTextField{
		NSTextField: MakeTextField(ptr),
	}
}

func SecureTextField_LabelWithAttributedString(attributedStringValue foundation.AttributedString) NSSecureTextField {
	result_ := C.C_NSSecureTextField_SecureTextField_LabelWithAttributedString(objc.ExtractPtr(attributedStringValue))
	return MakeSecureTextField(result_)
}

func SecureTextField_LabelWithString(stringValue string) NSSecureTextField {
	result_ := C.C_NSSecureTextField_SecureTextField_LabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeSecureTextField(result_)
}

func SecureTextField_TextFieldWithString(stringValue string) NSSecureTextField {
	result_ := C.C_NSSecureTextField_SecureTextField_TextFieldWithString(foundation.NewString(stringValue).Ptr())
	return MakeSecureTextField(result_)
}

func SecureTextField_WrappingLabelWithString(stringValue string) NSSecureTextField {
	result_ := C.C_NSSecureTextField_SecureTextField_WrappingLabelWithString(foundation.NewString(stringValue).Ptr())
	return MakeSecureTextField(result_)
}

func (n NSSecureTextField) InitWithFrame(frameRect foundation.Rect) NSSecureTextField {
	result_ := C.C_NSSecureTextField_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSecureTextField(result_)
}

func (n NSSecureTextField) InitWithCoder(coder foundation.Coder) NSSecureTextField {
	result_ := C.C_NSSecureTextField_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSecureTextField(result_)
}

func (n NSSecureTextField) Init() NSSecureTextField {
	result_ := C.C_NSSecureTextField_Init(n.Ptr())
	return MakeSecureTextField(result_)
}

func AllocSecureTextField() NSSecureTextField {
	result_ := C.C_NSSecureTextField_AllocSecureTextField()
	return MakeSecureTextField(result_)
}

func NewSecureTextField() NSSecureTextField {
	result_ := C.C_NSSecureTextField_NewSecureTextField()
	return MakeSecureTextField(result_)
}

func (n NSSecureTextField) Autorelease() NSSecureTextField {
	result_ := C.C_NSSecureTextField_Autorelease(n.Ptr())
	return MakeSecureTextField(result_)
}

func (n NSSecureTextField) Retain() NSSecureTextField {
	result_ := C.C_NSSecureTextField_Retain(n.Ptr())
	return MakeSecureTextField(result_)
}
