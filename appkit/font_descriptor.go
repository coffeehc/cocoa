package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeFontDescriptor(ptr unsafe.Pointer) *NSFontDescriptor {
	if ptr == nil {
		return nil
	}
	return &NSFontDescriptor{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocFontDescriptor() *NSFontDescriptor {
	return MakeFontDescriptor(C.C_FontDescriptor_Alloc())
}

func (n *NSFontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithDesign(n.Ptr(), foundation.NewString(string(design)).Ptr())
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) Init() FontDescriptor {
	result := C.C_NSFontDescriptor_Init(n.Ptr())
	return MakeFontDescriptor(result)
}

func FontDescriptorWithName_Matrix(fontName string, matrix foundation.AffineTransform) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithName_Matrix(foundation.NewString(fontName).Ptr(), objc.ExtractPtr(matrix))
	return MakeFontDescriptor(result)
}

func FontDescriptorWithName_Size(fontName string, size coregraphics.Float) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithName_Size(foundation.NewString(fontName).Ptr(), C.double(float64(size)))
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithFace(n.Ptr(), foundation.NewString(newFace).Ptr())
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithFamily(n.Ptr(), foundation.NewString(newFamily).Ptr())
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) FontDescriptorWithMatrix(matrix foundation.AffineTransform) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithMatrix(n.Ptr(), objc.ExtractPtr(matrix))
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) FontDescriptorWithSize(newPointSize coregraphics.Float) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithSize(n.Ptr(), C.double(float64(newPointSize)))
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	result := C.C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(n.Ptr(), C.uint(uint32(symbolicTraits)))
	return MakeFontDescriptor(result)
}

func (n *NSFontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	result := C.C_NSFontDescriptor_ObjectForKey(n.Ptr(), foundation.NewString(string(attribute)).Ptr())
	return objc.MakeObject(result)
}

func (n *NSFontDescriptor) Matrix() foundation.AffineTransform {
	result := C.C_NSFontDescriptor_Matrix(n.Ptr())
	return foundation.MakeAffineTransform(result)
}

func (n *NSFontDescriptor) PointSize() coregraphics.Float {
	result := C.C_NSFontDescriptor_PointSize(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFontDescriptor) PostscriptName() string {
	result := C.C_NSFontDescriptor_PostscriptName(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSFontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	result := C.C_NSFontDescriptor_SymbolicTraits(n.Ptr())
	return FontDescriptorSymbolicTraits(uint32(result))
}

func (n *NSFontDescriptor) RequiresFontAssetRequest() bool {
	result := C.C_NSFontDescriptor_RequiresFontAssetRequest(n.Ptr())
	return bool(result)
}
