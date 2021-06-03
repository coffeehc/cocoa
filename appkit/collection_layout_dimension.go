package appkit

// #include "collection_layout_dimension.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutDimension interface {
	objc.Object
	Dimension() coregraphics.Float
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool
}

type NSCollectionLayoutDimension struct {
	objc.NSObject
}

func MakeCollectionLayoutDimension(ptr unsafe.Pointer) NSCollectionLayoutDimension {
	return NSCollectionLayoutDimension{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutDimension() NSCollectionLayoutDimension {
	return MakeCollectionLayoutDimension(C.C_CollectionLayoutDimension_Alloc())
}

func CollectionLayoutDimension_AbsoluteDimension(absoluteDimension coregraphics.Float) CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(C.double(float64(absoluteDimension)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_EstimatedDimension(estimatedDimension coregraphics.Float) CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(C.double(float64(estimatedDimension)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_FractionalHeightDimension(fractionalHeight coregraphics.Float) CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(C.double(float64(fractionalHeight)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_FractionalWidthDimension(fractionalWidth coregraphics.Float) CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(C.double(float64(fractionalWidth)))
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutDimension) Dimension() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutDimension_Dimension(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutDimension) IsAbsolute() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsAbsolute(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsEstimated() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsEstimated(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsFractionalHeight() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsFractionalHeight(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsFractionalWidth() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsFractionalWidth(n.Ptr())
	return bool(result_)
}
