package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image_symbol_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ImageSymbolConfiguration interface {
	objc.Object
}

var _ ImageSymbolConfiguration = (*NSImageSymbolConfiguration)(nil)

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

func ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	return MakeImageSymbolConfiguration(C.ImageSymbolConfiguration_ConfigurationWithScale(C.long(scale)))
}

func ConfigurationWithPointSize(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	return MakeImageSymbolConfiguration(C.ImageSymbolConfiguration_ConfigurationWithPointSize(C.double(pointSize), C.double(weight)))
}

func ConfigurationWithPointSizeAndScale(pointSize float64, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	return MakeImageSymbolConfiguration(C.ImageSymbolConfiguration_ConfigurationWithPointSizeAndScale(C.double(pointSize), C.double(weight), C.long(scale)))
}

func ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	cStyle := C.CString(string(style))
	defer C.free(unsafe.Pointer(cStyle))
	return MakeImageSymbolConfiguration(C.ImageSymbolConfiguration_ConfigurationWithTextStyle(cStyle))
}

func configurationWithTextStyleAndScale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	cStyle := C.CString(string(style))
	defer C.free(unsafe.Pointer(cStyle))
	return MakeImageSymbolConfiguration(C.ImageSymbolConfiguration_configurationWithTextStyleAndScale(cStyle, C.long(scale)))
}
