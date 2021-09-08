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

func MakePressureConfiguration(ptr unsafe.Pointer) NSPressureConfiguration {
	return NSPressureConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSPressureConfiguration) InitWithPressureBehavior(pressureBehavior PressureBehavior) NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_InitWithPressureBehavior(n.Ptr(), C.int(int(pressureBehavior)))
	return MakePressureConfiguration(result_)
}

func AllocPressureConfiguration() NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_AllocPressureConfiguration()
	return MakePressureConfiguration(result_)
}

func (n NSPressureConfiguration) Init() NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_Init(n.Ptr())
	return MakePressureConfiguration(result_)
}

func NewPressureConfiguration() NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_NewPressureConfiguration()
	return MakePressureConfiguration(result_)
}

func (n NSPressureConfiguration) Autorelease() NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_Autorelease(n.Ptr())
	return MakePressureConfiguration(result_)
}

func (n NSPressureConfiguration) Retain() NSPressureConfiguration {
	result_ := C.C_NSPressureConfiguration_Retain(n.Ptr())
	return MakePressureConfiguration(result_)
}

func (n NSPressureConfiguration) Set() {
	C.C_NSPressureConfiguration_Set(n.Ptr())
}

func (n NSPressureConfiguration) PressureBehavior() PressureBehavior {
	result_ := C.C_NSPressureConfiguration_PressureBehavior(n.Ptr())
	return PressureBehavior(int(result_))
}
