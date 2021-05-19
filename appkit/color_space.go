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

func MakeColorSpace(ptr unsafe.Pointer) *NSColorSpace {
	if ptr == nil {
		return nil
	}
	return &NSColorSpace{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocColorSpace() *NSColorSpace {
	return MakeColorSpace(C.C_ColorSpace_Alloc())
}

func (n *NSColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	result_ := C.C_NSColorSpace_InitWithCGColorSpace(n.Ptr(), *(*C.CGColorSpaceRef)(coregraphics.ToCGColorSpaceRefPointer(cgColorSpace)))
	return MakeColorSpace(result_)
}

func (n *NSColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	result_ := C.C_NSColorSpace_InitWithICCProfileData(n.Ptr(), C.Array{data: unsafe.Pointer(&iccData[0]), len: C.int(len(iccData))})
	return MakeColorSpace(result_)
}

func (n *NSColorSpace) Init() ColorSpace {
	result_ := C.C_NSColorSpace_Init(n.Ptr())
	return MakeColorSpace(result_)
}

func ColorSpace_AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	result_ := C.C_NSColorSpace_ColorSpace_AvailableColorSpacesWithModel(C.int(int(model)))
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
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

func (n *NSColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	result_ := C.C_NSColorSpace_CGColorSpace(n.Ptr())
	return coregraphics.FromCGColorSpaceRefPointer(unsafe.Pointer(&result_))
}

func (n *NSColorSpace) ColorSpaceModel() ColorSpaceModel {
	result_ := C.C_NSColorSpace_ColorSpaceModel(n.Ptr())
	return ColorSpaceModel(int(result_))
}

func (n *NSColorSpace) ICCProfileData() []byte {
	result_ := C.C_NSColorSpace_ICCProfileData(n.Ptr())
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	C.free(result_.data)
	return goResult_
}

func (n *NSColorSpace) LocalizedName() string {
	result_ := C.C_NSColorSpace_LocalizedName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSColorSpace) NumberOfColorComponents() int {
	result_ := C.C_NSColorSpace_NumberOfColorComponents(n.Ptr())
	return int(result_)
}
