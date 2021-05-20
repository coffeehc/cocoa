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

func MakeTintConfiguration(ptr unsafe.Pointer) *NSTintConfiguration {
	if ptr == nil {
		return nil
	}
	return &NSTintConfiguration{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTintConfiguration() *NSTintConfiguration {
	return MakeTintConfiguration(C.C_TintConfiguration_Alloc())
}

func (n *NSTintConfiguration) Init() TintConfiguration {
	result_ := C.C_NSTintConfiguration_Init(n.Ptr())
	return MakeTintConfiguration(result_)
}

func TintConfigurationWithFixedColor(color Color) TintConfiguration {
	result_ := C.C_NSTintConfiguration_TintConfigurationWithFixedColor(objc.ExtractPtr(color))
	return MakeTintConfiguration(result_)
}

func TintConfigurationWithPreferredColor(color Color) TintConfiguration {
	result_ := C.C_NSTintConfiguration_TintConfigurationWithPreferredColor(objc.ExtractPtr(color))
	return MakeTintConfiguration(result_)
}

func (n *NSTintConfiguration) AdaptsToUserAccentColor() bool {
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

func (n *NSTintConfiguration) BaseTintColor() Color {
	result_ := C.C_NSTintConfiguration_BaseTintColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSTintConfiguration) EquivalentContentTintColor() Color {
	result_ := C.C_NSTintConfiguration_EquivalentContentTintColor(n.Ptr())
	return MakeColor(result_)
}
