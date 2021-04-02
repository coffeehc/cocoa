package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "secure_text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type SecureTextField interface {
	TextField
}

var _ SecureTextField = (*NSSecureTextField)(nil)

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
func (s *NSSecureTextField) setDelegate() {
	id := resources.Register(s)
	C.SecureTextField_SetDelegate(s.Ptr(), C.long(id))
}

func NewSecureTextField(frame foundation.Rect) SecureTextField {
	instance := MakeSecureTextField(C.SecureTextField_NewSecureTextField(toNSRect(frame)))
	instance.setDelegate()
	return instance
}
