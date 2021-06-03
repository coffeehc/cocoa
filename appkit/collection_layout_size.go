package appkit

// #include "collection_layout_size.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutSize interface {
	objc.Object
	WidthDimension() CollectionLayoutDimension
	HeightDimension() CollectionLayoutDimension
}

type NSCollectionLayoutSize struct {
	objc.NSObject
}

func MakeCollectionLayoutSize(ptr unsafe.Pointer) NSCollectionLayoutSize {
	return NSCollectionLayoutSize{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutSize() NSCollectionLayoutSize {
	return MakeCollectionLayoutSize(C.C_CollectionLayoutSize_Alloc())
}

func CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(width CollectionLayoutDimension, height CollectionLayoutDimension) CollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(objc.ExtractPtr(width), objc.ExtractPtr(height))
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutSize_WidthDimension(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutSize_HeightDimension(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}
