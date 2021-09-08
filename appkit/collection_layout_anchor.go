package appkit

// #include "collection_layout_anchor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutAnchor interface {
	objc.Object
	Edges() DirectionalRectEdge
	Offset() foundation.Point
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
}

type NSCollectionLayoutAnchor struct {
	objc.NSObject
}

func MakeCollectionLayoutAnchor(ptr unsafe.Pointer) NSCollectionLayoutAnchor {
	return NSCollectionLayoutAnchor{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges(edges DirectionalRectEdge) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(C.uint(uint(edges)))
	return MakeCollectionLayoutAnchor(result_)
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(C.uint(uint(edges)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(absoluteOffset))))
	return MakeCollectionLayoutAnchor(result_)
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(C.uint(uint(edges)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(fractionalOffset))))
	return MakeCollectionLayoutAnchor(result_)
}

func AllocCollectionLayoutAnchor() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_AllocCollectionLayoutAnchor()
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Autorelease() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_Autorelease(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Retain() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_Retain(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Edges() DirectionalRectEdge {
	result_ := C.C_NSCollectionLayoutAnchor_Edges(n.Ptr())
	return DirectionalRectEdge(uint(result_))
}

func (n NSCollectionLayoutAnchor) Offset() foundation.Point {
	result_ := C.C_NSCollectionLayoutAnchor_Offset(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutAnchor) IsAbsoluteOffset() bool {
	result_ := C.C_NSCollectionLayoutAnchor_IsAbsoluteOffset(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutAnchor) IsFractionalOffset() bool {
	result_ := C.C_NSCollectionLayoutAnchor_IsFractionalOffset(n.Ptr())
	return bool(result_)
}
