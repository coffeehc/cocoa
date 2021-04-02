package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "status_bar_button.h"
import "C"
import (
	"unsafe"
)

type StatusBarButton interface {
	Button
	AppearsDisabled() bool
	SetAppearsDisabled(appearsDisabled bool)
}

var _ StatusBarButton = (*NSStatusBarButton)(nil)

type NSStatusBarButton struct {
	NSButton
}

func MakeStatusBarButton(ptr unsafe.Pointer) *NSStatusBarButton {
	if ptr == nil {
		return nil
	}
	return &NSStatusBarButton{
		NSButton: *MakeButton(ptr),
	}
}

func (s *NSStatusBarButton) AppearsDisabled() bool {
	return bool(C.StatusBarButton_AppearsDisabled(s.Ptr()))
}

func (s *NSStatusBarButton) SetAppearsDisabled(appearsDisabled bool) {
	C.StatusBarButton_SetAppearsDisabled(s.Ptr(), C.bool(appearsDisabled))
}
