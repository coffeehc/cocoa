package appkit

// #include "color_space.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorSpace interface {
	objc.Object
	CGColorSpace() coregraphics.ColorSpaceRef
	ColorSpaceModel() ColorSpaceModel
	ICCProfileData() []byte
	LocalizedName() string
	NumberOfColorComponents() int
}

type NSColorSpace struct {
	objc.NSObject
}

func MakeColorSpace(ptr unsafe.Pointer) NSColorSpace {
	return NSColorSpace{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) NSColorSpace {
	result_ := C.C_NSColorSpace_InitWithCGColorSpace(n.Ptr(), unsafe.Pointer(cgColorSpace))
	return MakeColorSpace(result_)
}

func (n NSColorSpace) InitWithICCProfileData(iccData []byte) NSColorSpace {
	result_ := C.C_NSColorSpace_InitWithICCProfileData(n.Ptr(), foundation.NewData(iccData).Ptr())
	return MakeColorSpace(result_)
}

func AllocColorSpace() NSColorSpace {
	result_ := C.C_NSColorSpace_AllocColorSpace()
	return MakeColorSpace(result_)
}

func (n NSColorSpace) Init() NSColorSpace {
	result_ := C.C_NSColorSpace_Init(n.Ptr())
	return MakeColorSpace(result_)
}

func NewColorSpace() NSColorSpace {
	result_ := C.C_NSColorSpace_NewColorSpace()
	return MakeColorSpace(result_)
}

func (n NSColorSpace) Autorelease() NSColorSpace {
	result_ := C.C_NSColorSpace_Autorelease(n.Ptr())
	return MakeColorSpace(result_)
}

func (n NSColorSpace) Retain() NSColorSpace {
	result_ := C.C_NSColorSpace_Retain(n.Ptr())
	return MakeColorSpace(result_)
}

func AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	result_ := C.C_NSColorSpace_AvailableColorSpacesWithModel(C.int(int(model)))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ColorSpace, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeColorSpace(r)
	}
	return goResult_
}

func DeviceRGBColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_DeviceRGBColorSpace()
	return MakeColorSpace(result_)
}

func GenericRGBColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_GenericRGBColorSpace()
	return MakeColorSpace(result_)
}

func DeviceCMYKColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_DeviceCMYKColorSpace()
	return MakeColorSpace(result_)
}

func GenericCMYKColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_GenericCMYKColorSpace()
	return MakeColorSpace(result_)
}

func DeviceGrayColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_DeviceGrayColorSpace()
	return MakeColorSpace(result_)
}

func GenericGrayColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_GenericGrayColorSpace()
	return MakeColorSpace(result_)
}

func SRGBColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_SRGBColorSpace()
	return MakeColorSpace(result_)
}

func ExtendedSRGBColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_ExtendedSRGBColorSpace()
	return MakeColorSpace(result_)
}

func DisplayP3ColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_DisplayP3ColorSpace()
	return MakeColorSpace(result_)
}

func GenericGamma22GrayColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_GenericGamma22GrayColorSpace()
	return MakeColorSpace(result_)
}

func ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_ExtendedGenericGamma22GrayColorSpace()
	return MakeColorSpace(result_)
}

func AdobeRGB1998ColorSpace() ColorSpace {
	result_ := C.C_NSColorSpace_AdobeRGB1998ColorSpace()
	return MakeColorSpace(result_)
}

func (n NSColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	result_ := C.C_NSColorSpace_CGColorSpace(n.Ptr())
	return coregraphics.ColorSpaceRef(result_)
}

func (n NSColorSpace) ColorSpaceModel() ColorSpaceModel {
	result_ := C.C_NSColorSpace_ColorSpaceModel(n.Ptr())
	return ColorSpaceModel(int(result_))
}

func (n NSColorSpace) ICCProfileData() []byte {
	result_ := C.C_NSColorSpace_ICCProfileData(n.Ptr())
	return foundation.MakeData(result_).ToBytes()
}

func (n NSColorSpace) LocalizedName() string {
	result_ := C.C_NSColorSpace_LocalizedName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSColorSpace) NumberOfColorComponents() int {
	result_ := C.C_NSColorSpace_NumberOfColorComponents(n.Ptr())
	return int(result_)
}
