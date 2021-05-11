package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "bitmap_image_rep.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type BitmapImageRep interface {
	ImageRep
	ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint coregraphics.Float, midPointColor Color, shadowColor Color, lightColor Color)
	SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.Object)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	IncrementalLoadFromData_Complete(data []byte, complete bool) int
	SetColor_AtX_Y(color Color, x int, y int)
	ColorAtX_Y(x int, y int) Color
	BitmapImageRepByRetaggingWithColorSpace(newSpace ColorSpace) BitmapImageRep
	BitsPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	IsPlanar() bool
	NumberOfPlanes() int
	SamplesPerPixel() int
	TIFFRepresentation() []byte
	CGImage() coregraphics.ImageRef
	ColorSpace() ColorSpace
}

type NSBitmapImageRep struct {
	NSImageRep
}

func MakeBitmapImageRep(ptr unsafe.Pointer) *NSBitmapImageRep {
	if ptr == nil {
		return nil
	}
	return &NSBitmapImageRep{
		NSImageRep: *MakeImageRep(ptr),
	}
}

func AllocBitmapImageRep() *NSBitmapImageRep {
	return MakeBitmapImageRep(C.C_BitmapImageRep_Alloc())
}

func (n *NSBitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	result := C.C_NSBitmapImageRep_InitWithCGImage(n.Ptr(), *(*C.CGImageRef)(coregraphics.ToCGImageRefPointer(cgImage)))
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	result := C.C_NSBitmapImageRep_InitWithData(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	result := C.C_NSBitmapImageRep_InitForIncrementalLoad(n.Ptr())
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) Init() BitmapImageRep {
	result := C.C_NSBitmapImageRep_Init(n.Ptr())
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) InitWithCoder(coder foundation.Coder) BitmapImageRep {
	result := C.C_NSBitmapImageRep_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeBitmapImageRep(result)
}

func BitmapImageRepImageRepWithData(data []byte) BitmapImageRep {
	result := C.C_NSBitmapImageRep_BitmapImageRepImageRepWithData(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint coregraphics.Float, midPointColor Color, shadowColor Color, lightColor Color) {
	C.C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(n.Ptr(), C.double(float64(midPoint)), objc.ExtractPtr(midPointColor), objc.ExtractPtr(shadowColor), objc.ExtractPtr(lightColor))
}

func (n *NSBitmapImageRep) SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.Object) {
	C.C_NSBitmapImageRep_SetProperty_WithValue(n.Ptr(), foundation.NewString(string(property)).Ptr(), objc.ExtractPtr(value))
}

func (n *NSBitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	result := C.C_NSBitmapImageRep_ValueForProperty(n.Ptr(), foundation.NewString(string(property)).Ptr())
	return objc.MakeObject(result)
}

func (n *NSBitmapImageRep) IncrementalLoadFromData_Complete(data []byte, complete bool) int {
	result := C.C_NSBitmapImageRep_IncrementalLoadFromData_Complete(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, C.bool(complete))
	return int(result)
}

func (n *NSBitmapImageRep) SetColor_AtX_Y(color Color, x int, y int) {
	C.C_NSBitmapImageRep_SetColor_AtX_Y(n.Ptr(), objc.ExtractPtr(color), C.int(x), C.int(y))
}

func (n *NSBitmapImageRep) ColorAtX_Y(x int, y int) Color {
	result := C.C_NSBitmapImageRep_ColorAtX_Y(n.Ptr(), C.int(x), C.int(y))
	return MakeColor(result)
}

func (n *NSBitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace ColorSpace) BitmapImageRep {
	result := C.C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(n.Ptr(), objc.ExtractPtr(newSpace))
	return MakeBitmapImageRep(result)
}

func (n *NSBitmapImageRep) BitsPerPixel() int {
	result := C.C_NSBitmapImageRep_BitsPerPixel(n.Ptr())
	return int(result)
}

func (n *NSBitmapImageRep) BytesPerPlane() int {
	result := C.C_NSBitmapImageRep_BytesPerPlane(n.Ptr())
	return int(result)
}

func (n *NSBitmapImageRep) BytesPerRow() int {
	result := C.C_NSBitmapImageRep_BytesPerRow(n.Ptr())
	return int(result)
}

func (n *NSBitmapImageRep) IsPlanar() bool {
	result := C.C_NSBitmapImageRep_IsPlanar(n.Ptr())
	return bool(result)
}

func (n *NSBitmapImageRep) NumberOfPlanes() int {
	result := C.C_NSBitmapImageRep_NumberOfPlanes(n.Ptr())
	return int(result)
}

func (n *NSBitmapImageRep) SamplesPerPixel() int {
	result := C.C_NSBitmapImageRep_SamplesPerPixel(n.Ptr())
	return int(result)
}

func (n *NSBitmapImageRep) TIFFRepresentation() []byte {
	result := C.C_NSBitmapImageRep_TIFFRepresentation(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSBitmapImageRep) CGImage() coregraphics.ImageRef {
	result := C.C_NSBitmapImageRep_CGImage(n.Ptr())
	return coregraphics.FromCGImageRefPointer(unsafe.Pointer(&result))
}

func (n *NSBitmapImageRep) ColorSpace() ColorSpace {
	result := C.C_NSBitmapImageRep_ColorSpace(n.Ptr())
	return MakeColorSpace(result)
}
