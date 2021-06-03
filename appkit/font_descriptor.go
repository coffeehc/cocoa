package appkit

// #include "font_descriptor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type FontDescriptor interface {
	objc.Object
	FontDescriptorWithFace(newFace string) FontDescriptor
	FontDescriptorWithFamily(newFamily string) FontDescriptor
	FontDescriptorWithMatrix(matrix foundation.AffineTransform) FontDescriptor
	FontDescriptorWithSize(newPointSize coregraphics.Float) FontDescriptor
	FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.Set) []FontDescriptor
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.Set) FontDescriptor
	ObjectForKey(attribute FontDescriptorAttributeName) objc.Object
	Matrix() foundation.AffineTransform
	PointSize() coregraphics.Float
	PostscriptName() string
	SymbolicTraits() FontDescriptorSymbolicTraits
	RequiresFontAssetRequest() bool
}

type NSFontDescriptor struct {
	objc.NSObject
}

func MakeFontDescriptor(ptr unsafe.Pointer) NSFontDescriptor {
	return NSFontDescriptor{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocFontDescriptor() NSFontDescriptor {
	return MakeFontDescriptor(C.C_FontDescriptor_Alloc())
}

func (n NSFontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithDesign(n.Ptr(), foundation.NewString(string(design)).Ptr())
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) Init() FontDescriptor {
	result_ := C.C_NSFontDescriptor_Init(n.Ptr())
	return MakeFontDescriptor(result_)
}

func FontDescriptorWithName_Matrix(fontName string, matrix foundation.AffineTransform) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithName_Matrix(foundation.NewString(fontName).Ptr(), objc.ExtractPtr(matrix))
	return MakeFontDescriptor(result_)
}

func FontDescriptorWithName_Size(fontName string, size coregraphics.Float) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithName_Size(foundation.NewString(fontName).Ptr(), C.double(float64(size)))
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithFace(n.Ptr(), foundation.NewString(newFace).Ptr())
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithFamily(n.Ptr(), foundation.NewString(newFamily).Ptr())
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) FontDescriptorWithMatrix(matrix foundation.AffineTransform) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithMatrix(n.Ptr(), objc.ExtractPtr(matrix))
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) FontDescriptorWithSize(newPointSize coregraphics.Float) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithSize(n.Ptr(), C.double(float64(newPointSize)))
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	result_ := C.C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(n.Ptr(), C.uint(uint32(symbolicTraits)))
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.Set) []FontDescriptor {
	result_ := C.C_NSFontDescriptor_MatchingFontDescriptorsWithMandatoryKeys(n.Ptr(), objc.ExtractPtr(mandatoryKeys))
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]FontDescriptor, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeFontDescriptor(r)
	}
	return goResult_
}

func (n NSFontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.Set) FontDescriptor {
	result_ := C.C_NSFontDescriptor_MatchingFontDescriptorWithMandatoryKeys(n.Ptr(), objc.ExtractPtr(mandatoryKeys))
	return MakeFontDescriptor(result_)
}

func (n NSFontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	result_ := C.C_NSFontDescriptor_ObjectForKey(n.Ptr(), foundation.NewString(string(attribute)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSFontDescriptor) Matrix() foundation.AffineTransform {
	result_ := C.C_NSFontDescriptor_Matrix(n.Ptr())
	return foundation.MakeAffineTransform(result_)
}

func (n NSFontDescriptor) PointSize() coregraphics.Float {
	result_ := C.C_NSFontDescriptor_PointSize(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFontDescriptor) PostscriptName() string {
	result_ := C.C_NSFontDescriptor_PostscriptName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSFontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	result_ := C.C_NSFontDescriptor_SymbolicTraits(n.Ptr())
	return FontDescriptorSymbolicTraits(uint32(result_))
}

func (n NSFontDescriptor) RequiresFontAssetRequest() bool {
	result_ := C.C_NSFontDescriptor_RequiresFontAssetRequest(n.Ptr())
	return bool(result_)
}
