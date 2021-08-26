package appkit

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
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.Object) []byte
	CanBeCompressedUsing(compression TIFFCompression) bool
	SetCompression_Factor(compression TIFFCompression, factor float32)
	SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.Object)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	IncrementalLoadFromData_Complete(data []byte, complete bool) int
	SetColor_AtX_Y(color Color, x int, y int)
	ColorAtX_Y(x int, y int) Color
	BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace ColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep
	BitmapImageRepByRetaggingWithColorSpace(newSpace ColorSpace) BitmapImageRep
	BitmapFormat() BitmapFormat
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

func MakeBitmapImageRep(ptr unsafe.Pointer) NSBitmapImageRep {
	return NSBitmapImageRep{
		NSImageRep: MakeImageRep(ptr),
	}
}

func AllocBitmapImageRep() NSBitmapImageRep {
	return MakeBitmapImageRep(C.C_BitmapImageRep_Alloc())
}

func (n NSBitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_InitWithCGImage(n.Ptr(), unsafe.Pointer(cgImage))
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_InitWithData(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_InitForIncrementalLoad(n.Ptr())
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) Init() BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_Init(n.Ptr())
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) InitWithCoder(coder foundation.Coder) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeBitmapImageRep(result_)
}

func BitmapImageRep_ImageRepWithData(data []byte) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_ImageRepWithData(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeBitmapImageRep(result_)
}

func BitmapImageRep_ImageRepsWithData(data []byte) []ImageRep {
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_ImageRepsWithData(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ImageRep, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeImageRep(r)
	}
	return goResult_
}

func (n NSBitmapImageRep) ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint coregraphics.Float, midPointColor Color, shadowColor Color, lightColor Color) {
	C.C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(n.Ptr(), C.double(float64(midPoint)), objc.ExtractPtr(midPointColor), objc.ExtractPtr(shadowColor), objc.ExtractPtr(lightColor))
}

func BitmapImageRep_TIFFRepresentationOfImageRepsInArray(array []ImageRep) []byte {
	var cArray C.Array
	if len(array) > 0 {
		cArrayData := make([]unsafe.Pointer, len(array))
		for idx, v := range array {
			cArrayData[idx] = objc.ExtractPtr(v)
		}
		cArray.data = unsafe.Pointer(&cArrayData[0])
		cArray.len = C.int(len(array))
	}
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray(cArray)
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func BitmapImageRep_TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(array []ImageRep, comp TIFFCompression, factor float32) []byte {
	var cArray C.Array
	if len(array) > 0 {
		cArrayData := make([]unsafe.Pointer, len(array))
		for idx, v := range array {
			cArrayData[idx] = objc.ExtractPtr(v)
		}
		cArray.data = unsafe.Pointer(&cArrayData[0])
		cArray.len = C.int(len(array))
	}
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(cArray, C.uint(uint(comp)), C.float(factor))
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func (n NSBitmapImageRep) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	result_ := C.C_NSBitmapImageRep_TIFFRepresentationUsingCompression_Factor(n.Ptr(), C.uint(uint(comp)), C.float(factor))
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func BitmapImageRep_RepresentationOfImageRepsInArray_UsingType_Properties(imageReps []ImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.Object) []byte {
	var cImageReps C.Array
	if len(imageReps) > 0 {
		cImageRepsData := make([]unsafe.Pointer, len(imageReps))
		for idx, v := range imageReps {
			cImageRepsData[idx] = objc.ExtractPtr(v)
		}
		cImageReps.data = unsafe.Pointer(&cImageRepsData[0])
		cImageReps.len = C.int(len(imageReps))
	}
	var cProperties C.Dictionary
	if len(properties) > 0 {
		cPropertiesKeyData := make([]unsafe.Pointer, len(properties))
		cPropertiesValueData := make([]unsafe.Pointer, len(properties))
		var idx = 0
		for k, v := range properties {
			cPropertiesKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cPropertiesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cProperties.key_data = unsafe.Pointer(&cPropertiesKeyData[0])
		cProperties.value_data = unsafe.Pointer(&cPropertiesValueData[0])
		cProperties.len = C.int(len(properties))
	}
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_RepresentationOfImageRepsInArray_UsingType_Properties(cImageReps, C.uint(uint(storageType)), cProperties)
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func (n NSBitmapImageRep) RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.Object) []byte {
	var cProperties C.Dictionary
	if len(properties) > 0 {
		cPropertiesKeyData := make([]unsafe.Pointer, len(properties))
		cPropertiesValueData := make([]unsafe.Pointer, len(properties))
		var idx = 0
		for k, v := range properties {
			cPropertiesKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cPropertiesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cProperties.key_data = unsafe.Pointer(&cPropertiesKeyData[0])
		cProperties.value_data = unsafe.Pointer(&cPropertiesValueData[0])
		cProperties.len = C.int(len(properties))
	}
	result_ := C.C_NSBitmapImageRep_RepresentationUsingType_Properties(n.Ptr(), C.uint(uint(storageType)), cProperties)
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func BitmapImageRep_LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	result_ := C.C_NSBitmapImageRep_BitmapImageRep_LocalizedNameForTIFFCompressionType(C.uint(uint(compression)))
	return foundation.MakeString(result_).String()
}

func (n NSBitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	result_ := C.C_NSBitmapImageRep_CanBeCompressedUsing(n.Ptr(), C.uint(uint(compression)))
	return bool(result_)
}

func (n NSBitmapImageRep) SetCompression_Factor(compression TIFFCompression, factor float32) {
	C.C_NSBitmapImageRep_SetCompression_Factor(n.Ptr(), C.uint(uint(compression)), C.float(factor))
}

func (n NSBitmapImageRep) SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.Object) {
	C.C_NSBitmapImageRep_SetProperty_WithValue(n.Ptr(), foundation.NewString(string(property)).Ptr(), objc.ExtractPtr(value))
}

func (n NSBitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	result_ := C.C_NSBitmapImageRep_ValueForProperty(n.Ptr(), foundation.NewString(string(property)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSBitmapImageRep) IncrementalLoadFromData_Complete(data []byte, complete bool) int {
	result_ := C.C_NSBitmapImageRep_IncrementalLoadFromData_Complete(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, C.bool(complete))
	return int(result_)
}

func (n NSBitmapImageRep) SetColor_AtX_Y(color Color, x int, y int) {
	C.C_NSBitmapImageRep_SetColor_AtX_Y(n.Ptr(), objc.ExtractPtr(color), C.int(x), C.int(y))
}

func (n NSBitmapImageRep) ColorAtX_Y(x int, y int) Color {
	result_ := C.C_NSBitmapImageRep_ColorAtX_Y(n.Ptr(), C.int(x), C.int(y))
	return MakeColor(result_)
}

func (n NSBitmapImageRep) BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace ColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_BitmapImageRepByConvertingToColorSpace_RenderingIntent(n.Ptr(), objc.ExtractPtr(targetSpace), C.int(int(renderingIntent)))
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace ColorSpace) BitmapImageRep {
	result_ := C.C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(n.Ptr(), objc.ExtractPtr(newSpace))
	return MakeBitmapImageRep(result_)
}

func (n NSBitmapImageRep) BitmapFormat() BitmapFormat {
	result_ := C.C_NSBitmapImageRep_BitmapFormat(n.Ptr())
	return BitmapFormat(uint(result_))
}

func (n NSBitmapImageRep) BitsPerPixel() int {
	result_ := C.C_NSBitmapImageRep_BitsPerPixel(n.Ptr())
	return int(result_)
}

func (n NSBitmapImageRep) BytesPerPlane() int {
	result_ := C.C_NSBitmapImageRep_BytesPerPlane(n.Ptr())
	return int(result_)
}

func (n NSBitmapImageRep) BytesPerRow() int {
	result_ := C.C_NSBitmapImageRep_BytesPerRow(n.Ptr())
	return int(result_)
}

func (n NSBitmapImageRep) IsPlanar() bool {
	result_ := C.C_NSBitmapImageRep_IsPlanar(n.Ptr())
	return bool(result_)
}

func (n NSBitmapImageRep) NumberOfPlanes() int {
	result_ := C.C_NSBitmapImageRep_NumberOfPlanes(n.Ptr())
	return int(result_)
}

func (n NSBitmapImageRep) SamplesPerPixel() int {
	result_ := C.C_NSBitmapImageRep_SamplesPerPixel(n.Ptr())
	return int(result_)
}

func (n NSBitmapImageRep) TIFFRepresentation() []byte {
	result_ := C.C_NSBitmapImageRep_TIFFRepresentation(n.Ptr())
	if result_.len > 0 {
		C.free(result_.data)
	}
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	return goResult_
}

func (n NSBitmapImageRep) CGImage() coregraphics.ImageRef {
	result_ := C.C_NSBitmapImageRep_CGImage(n.Ptr())
	return coregraphics.ImageRef(result_)
}

func (n NSBitmapImageRep) ColorSpace() ColorSpace {
	result_ := C.C_NSBitmapImageRep_ColorSpace(n.Ptr())
	return MakeColorSpace(result_)
}
