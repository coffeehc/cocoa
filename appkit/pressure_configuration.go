package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "pressure_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PressureConfiguration interface {
	objc.Object
	Set()
}

type NSPressureConfiguration struct {
	objc.NSObject
}

func MakePressureConfiguration(ptr unsafe.Pointer) *NSPressureConfiguration {
	if ptr == nil {
		return nil
	}
	return &NSPressureConfiguration{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPressureConfiguration() *NSPressureConfiguration {
	return MakePressureConfiguration(C.C_PressureConfiguration_Alloc())
}

func (n *NSPressureConfiguration) Init() PressureConfiguration {
	result := C.C_NSPressureConfiguration_Init(n.Ptr())
	return MakePressureConfiguration(result)
}

func (n *NSPressureConfiguration) Set() {
	C.C_NSPressureConfiguration_Set(n.Ptr())
}
