package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "status_bar_button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusBarButton interface {
	Button
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

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

func AllocStatusBarButton() *NSStatusBarButton {
	return MakeStatusBarButton(C.C_StatusBarButton_Alloc())
}

func (n *NSStatusBarButton) InitWithFrame(frameRect foundation.Rect) StatusBarButton {
	result := C.C_NSStatusBarButton_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeStatusBarButton(result)
}

func (n *NSStatusBarButton) InitWithCoder(coder foundation.Coder) StatusBarButton {
	result := C.C_NSStatusBarButton_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeStatusBarButton(result)
}

func (n *NSStatusBarButton) Init() StatusBarButton {
	result := C.C_NSStatusBarButton_Init(n.Ptr())
	return MakeStatusBarButton(result)
}

func (n *NSStatusBarButton) AppearsDisabled() bool {
	result := C.C_NSStatusBarButton_AppearsDisabled(n.Ptr())
	return bool(result)
}

func (n *NSStatusBarButton) SetAppearsDisabled(value bool) {
	C.C_NSStatusBarButton_SetAppearsDisabled(n.Ptr(), C.bool(value))
}
