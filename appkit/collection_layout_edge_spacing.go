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

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) *NSCollectionLayoutEdgeSpacing {
	if ptr == nil {
		return nil
	}
	return &NSCollectionLayoutEdgeSpacing{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutEdgeSpacing() *NSCollectionLayoutEdgeSpacing {
	return MakeCollectionLayoutEdgeSpacing(C.C_CollectionLayoutEdgeSpacing_Alloc())
}

func CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(leading CollectionLayoutSpacing, top CollectionLayoutSpacing, trailing CollectionLayoutSpacing, bottom CollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(objc.ExtractPtr(leading), objc.ExtractPtr(top), objc.ExtractPtr(trailing), objc.ExtractPtr(bottom))
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n *NSCollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Leading(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n *NSCollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Top(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n *NSCollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Trailing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n *NSCollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Bottom(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}
