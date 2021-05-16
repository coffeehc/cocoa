package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "appearance.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Appearance interface {
	objc.Object
	Name() AppearanceName
	AllowsVibrancy() bool
}

type NSAppearance struct {
	objc.NSObject
}

func MakeAppearance(ptr unsafe.Pointer) *NSAppearance {
	if ptr == nil {
		return nil
	}
	return &NSAppearance{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocAppearance() *NSAppearance {
	return MakeAppearance(C.C_Appearance_Alloc())
}

func (n *NSAppearance) InitWithAppearanceNamed_Bundle(name AppearanceName, bundle foundation.Bundle) Appearance {
	result := C.C_NSAppearance_InitWithAppearanceNamed_Bundle(n.Ptr(), foundation.NewString(string(name)).Ptr(), objc.ExtractPtr(bundle))
	return MakeAppearance(result)
}

func (n *NSAppearance) InitWithCoder(coder foundation.Coder) Appearance {
	result := C.C_NSAppearance_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeAppearance(result)
}

func (n *NSAppearance) Init() Appearance {
	result := C.C_NSAppearance_Init(n.Ptr())
	return MakeAppearance(result)
}

func AppearanceNamed(name AppearanceName) Appearance {
	result := C.C_NSAppearance_AppearanceNamed(foundation.NewString(string(name)).Ptr())
	return MakeAppearance(result)
}

func (n *NSAppearance) Name() AppearanceName {
	result := C.C_NSAppearance_Name(n.Ptr())
	return AppearanceName(foundation.MakeString(result).String())
}

func CurrentDrawingAppearance() Appearance {
	result := C.C_NSAppearance_CurrentDrawingAppearance()
	return MakeAppearance(result)
}

func (n *NSAppearance) AllowsVibrancy() bool {
	result := C.C_NSAppearance_AllowsVibrancy(n.Ptr())
	return bool(result)
}
