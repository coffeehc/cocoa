package appkit

// #include "collection_layout_edge_spacing.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutEdgeSpacing interface {
	objc.Object
	Leading() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
}

type NSCollectionLayoutEdgeSpacing struct {
	objc.NSObject
}

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) NSCollectionLayoutEdgeSpacing {
	return NSCollectionLayoutEdgeSpacing{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(leading CollectionLayoutSpacing, top CollectionLayoutSpacing, trailing CollectionLayoutSpacing, bottom CollectionLayoutSpacing) NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(objc.ExtractPtr(leading), objc.ExtractPtr(top), objc.ExtractPtr(trailing), objc.ExtractPtr(bottom))
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func AllocCollectionLayoutEdgeSpacing() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_AllocCollectionLayoutEdgeSpacing()
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Autorelease() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Autorelease(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Retain() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Retain(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Leading(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Top(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Trailing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Bottom(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}
