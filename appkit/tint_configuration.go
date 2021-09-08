package appkit

// #include "tint_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TintConfiguration interface {
	objc.Object
	AdaptsToUserAccentColor() bool
	BaseTintColor() Color
	EquivalentContentTintColor() Color
}

type NSTintConfiguration struct {
	objc.NSObject
}

func MakeTintConfiguration(ptr unsafe.Pointer) NSTintConfiguration {
	return NSTintConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func TintConfigurationWithFixedColor(color Color) NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_TintConfigurationWithFixedColor(objc.ExtractPtr(color))
	return MakeTintConfiguration(result_)
}

func TintConfigurationWithPreferredColor(color Color) NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_TintConfigurationWithPreferredColor(objc.ExtractPtr(color))
	return MakeTintConfiguration(result_)
}

func AllocTintConfiguration() NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_AllocTintConfiguration()
	return MakeTintConfiguration(result_)
}

func (n NSTintConfiguration) Init() NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_Init(n.Ptr())
	return MakeTintConfiguration(result_)
}

func NewTintConfiguration() NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_NewTintConfiguration()
	return MakeTintConfiguration(result_)
}

func (n NSTintConfiguration) Autorelease() NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_Autorelease(n.Ptr())
	return MakeTintConfiguration(result_)
}

func (n NSTintConfiguration) Retain() NSTintConfiguration {
	result_ := C.C_NSTintConfiguration_Retain(n.Ptr())
	return MakeTintConfiguration(result_)
}

func (n NSTintConfiguration) AdaptsToUserAccentColor() bool {
	result_ := C.C_NSTintConfiguration_AdaptsToUserAccentColor(n.Ptr())
	return bool(result_)
}

func DefaultTintConfiguration() TintConfiguration {
	result_ := C.C_NSTintConfiguration_DefaultTintConfiguration()
	return MakeTintConfiguration(result_)
}

func MonochromeTintConfiguration() TintConfiguration {
	result_ := C.C_NSTintConfiguration_MonochromeTintConfiguration()
	return MakeTintConfiguration(result_)
}

func (n NSTintConfiguration) BaseTintColor() Color {
	result_ := C.C_NSTintConfiguration_BaseTintColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTintConfiguration) EquivalentContentTintColor() Color {
	result_ := C.C_NSTintConfiguration_EquivalentContentTintColor(n.Ptr())
	return MakeColor(result_)
}
