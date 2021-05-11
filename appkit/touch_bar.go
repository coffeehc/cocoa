package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "touch_bar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TouchBar interface {
	objc.Object
	PrincipalItemIdentifier() TouchBarItemIdentifier
	SetPrincipalItemIdentifier(value TouchBarItemIdentifier)
	EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier)
	IsVisible() bool
	CustomizationIdentifier() TouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value TouchBarCustomizationIdentifier)
}

type NSTouchBar struct {
	objc.NSObject
}

func MakeTouchBar(ptr unsafe.Pointer) *NSTouchBar {
	if ptr == nil {
		return nil
	}
	return &NSTouchBar{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTouchBar() *NSTouchBar {
	return MakeTouchBar(C.C_TouchBar_Alloc())
}

func (n *NSTouchBar) Init() TouchBar {
	result := C.C_NSTouchBar_Init(n.Ptr())
	return MakeTouchBar(result)
}

func (n *NSTouchBar) InitWithCoder(coder foundation.Coder) TouchBar {
	result := C.C_NSTouchBar_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTouchBar(result)
}

func (n *NSTouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	result := C.C_NSTouchBar_PrincipalItemIdentifier(n.Ptr())
	return TouchBarItemIdentifier(foundation.MakeString(result).String())
}

func (n *NSTouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	C.C_NSTouchBar_SetPrincipalItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSTouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	result := C.C_NSTouchBar_EscapeKeyReplacementItemIdentifier(n.Ptr())
	return TouchBarItemIdentifier(foundation.MakeString(result).String())
}

func (n *NSTouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	C.C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSTouchBar) IsVisible() bool {
	result := C.C_NSTouchBar_IsVisible(n.Ptr())
	return bool(result)
}

func (n *NSTouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	result := C.C_NSTouchBar_CustomizationIdentifier(n.Ptr())
	return TouchBarCustomizationIdentifier(foundation.MakeString(result).String())
}

func (n *NSTouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	C.C_NSTouchBar_SetCustomizationIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
