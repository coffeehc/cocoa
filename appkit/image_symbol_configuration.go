package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image_symbol_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ImageSymbolConfiguration interface {
	objc.Object
}

type NSImageSymbolConfiguration struct {
	objc.NSObject
}

func MakeImageSymbolConfiguration(ptr unsafe.Pointer) *NSImageSymbolConfiguration {
	if ptr == nil {
		return nil
	}
	return &NSImageSymbolConfiguration{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocImageSymbolConfiguration() *NSImageSymbolConfiguration {
	return MakeImageSymbolConfiguration(C.C_ImageSymbolConfiguration_Alloc())
}

func (n *NSImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_Init(n.Ptr())
	return MakeImageSymbolConfiguration(result_)
}

func ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(pointSize coregraphics.Float, weight FontWeight) ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(C.double(float64(pointSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeImageSymbolConfiguration(result_)
}

func ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(pointSize coregraphics.Float, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(C.double(float64(pointSize)), C.double(float64(coregraphics.Float(weight))), C.int(int(scale)))
	return MakeImageSymbolConfiguration(result_)
}

func ImageSymbolConfiguration_ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle(foundation.NewString(string(style)).Ptr())
	return MakeImageSymbolConfiguration(result_)
}

func ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(foundation.NewString(string(style)).Ptr(), C.int(int(scale)))
	return MakeImageSymbolConfiguration(result_)
}

func ImageSymbolConfiguration_ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	result_ := C.C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithScale(C.int(int(scale)))
	return MakeImageSymbolConfiguration(result_)
}
