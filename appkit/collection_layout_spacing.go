package appkit

// #include "collection_layout_spacing.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutSpacing interface {
	objc.Object
	Spacing() coregraphics.Float
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
}

type NSCollectionLayoutSpacing struct {
	objc.NSObject
}

func MakeCollectionLayoutSpacing(ptr unsafe.Pointer) NSCollectionLayoutSpacing {
	return NSCollectionLayoutSpacing{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutSpacing() NSCollectionLayoutSpacing {
	return MakeCollectionLayoutSpacing(C.C_CollectionLayoutSpacing_Alloc())
}

func CollectionLayoutSpacing_FixedSpacing(fixedSpacing coregraphics.Float) CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(C.double(float64(fixedSpacing)))
	return MakeCollectionLayoutSpacing(result_)
}

func CollectionLayoutSpacing_FlexibleSpacing(flexibleSpacing coregraphics.Float) CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(C.double(float64(flexibleSpacing)))
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutSpacing) Spacing() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutSpacing_Spacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutSpacing) IsFixedSpacing() bool {
	result_ := C.C_NSCollectionLayoutSpacing_IsFixedSpacing(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutSpacing) IsFlexibleSpacing() bool {
	result_ := C.C_NSCollectionLayoutSpacing_IsFlexibleSpacing(n.Ptr())
	return bool(result_)
}
