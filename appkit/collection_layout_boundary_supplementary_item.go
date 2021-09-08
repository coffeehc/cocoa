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

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) NSCollectionLayoutBoundarySupplementaryItem {
	return NSCollectionLayoutBoundarySupplementaryItem{
		NSCollectionLayoutSupplementaryItem: MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(absoluteOffset))))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor, itemAnchor CollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutBoundarySupplementaryItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func AllocCollectionLayoutBoundarySupplementaryItem() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_AllocCollectionLayoutBoundarySupplementaryItem()
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Autorelease() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Retain() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Retain(n.Ptr())
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(n.Ptr(), C.bool(value))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Offset(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(n.Ptr())
	return RectAlignment(int(result_))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(n.Ptr(), C.bool(value))
}
