package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	result := C.C_NSColorSpace_InitWithCGColorSpace(n.Ptr(), *(*C.CGColorSpaceRef)(coregraphics.ToCGColorSpaceRefPointer(cgColorSpace)))
	return MakeColorSpace(result)
}

func (n *NSColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	result := C.C_NSColorSpace_InitWithICCProfileData(n.Ptr(), C.Array{data: unsafe.Pointer(&iccData[0]), len: C.int(len(iccData))})
	return MakeColorSpace(result)
}

func (n *NSColorSpace) Init() ColorSpace {
	result := C.C_NSColorSpace_Init(n.Ptr())
	return MakeColorSpace(result)
}

func (n *NSColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	result := C.C_NSColorSpace_CGColorSpace(n.Ptr())
	return coregraphics.FromCGColorSpaceRefPointer(unsafe.Pointer(&result))
}

func (n *NSColorSpace) ICCProfileData() []byte {
	result := C.C_NSColorSpace_ICCProfileData(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSColorSpace) LocalizedName() string {
	result := C.C_NSColorSpace_LocalizedName(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSColorSpace) NumberOfColorComponents() int {
	result := C.C_NSColorSpace_NumberOfColorComponents(n.Ptr())
	return int(result)
}
