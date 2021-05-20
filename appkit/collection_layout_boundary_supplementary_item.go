package appkit

// #include "collection_layout_boundary_supplementary_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutBoundarySupplementaryItem interface {
	CollectionLayoutSupplementaryItem
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() foundation.Point
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

type NSCollectionLayoutBoundarySupplementaryItem struct {
	NSCollectionLayoutSupplementaryItem
}

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) *NSCollectionLayoutBoundarySupplementaryItem {
	if ptr == nil {
		return nil
	}
	return &NSCollectionLayoutBoundarySupplementaryItem{
		NSCollectionLayoutSupplementaryItem: *MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func AllocCollectionLayoutBoundarySupplementaryItem() *NSCollectionLayoutBoundarySupplementaryItem {
	return MakeCollectionLayoutBoundarySupplementaryItem(C.C_CollectionLayoutBoundarySupplementaryItem_Alloc())
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(absoluteOffset))))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(n.Ptr(), C.bool(value))
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Offset(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(n.Ptr())
	return RectAlignment(int(result_))
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(n.Ptr(), C.bool(value))
}
