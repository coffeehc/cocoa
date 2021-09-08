package appkit

// #include "collection_layout_supplementary_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutSupplementaryItem interface {
	CollectionLayoutItem
	ItemAnchor() CollectionLayoutAnchor
	ContainerAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type NSCollectionLayoutSupplementaryItem struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutSupplementaryItem(ptr unsafe.Pointer) NSCollectionLayoutSupplementaryItem {
	return NSCollectionLayoutSupplementaryItem{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor, itemAnchor CollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutSupplementaryItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func AllocCollectionLayoutSupplementaryItem() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_AllocCollectionLayoutSupplementaryItem()
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) Autorelease() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) Retain() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_Retain(n.Ptr())
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ItemAnchor(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ContainerAnchor(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ElementKind() string {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutSupplementaryItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ZIndex(n.Ptr())
	return int(result_)
}

func (n NSCollectionLayoutSupplementaryItem) SetZIndex(value int) {
	C.C_NSCollectionLayoutSupplementaryItem_SetZIndex(n.Ptr(), C.int(value))
}
