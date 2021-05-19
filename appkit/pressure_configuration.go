package appkit

// #include "pressure_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PressureConfiguration interface {
	objc.Object
	Set()
	PressureBehavior() PressureBehavior
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

func (n *NSPressureConfiguration) InitWithPressureBehavior(pressureBehavior PressureBehavior) PressureConfiguration {
	result_ := C.C_NSPressureConfiguration_InitWithPressureBehavior(n.Ptr(), C.int(int(pressureBehavior)))
	return MakePressureConfiguration(result_)
}

func (n *NSPressureConfiguration) Init() PressureConfiguration {
	result_ := C.C_NSPressureConfiguration_Init(n.Ptr())
	return MakePressureConfiguration(result_)
}

func (n *NSPressureConfiguration) Set() {
	C.C_NSPressureConfiguration_Set(n.Ptr())
}

func (n *NSPressureConfiguration) PressureBehavior() PressureBehavior {
	result_ := C.C_NSPressureConfiguration_PressureBehavior(n.Ptr())
	return PressureBehavior(int(result_))
}
