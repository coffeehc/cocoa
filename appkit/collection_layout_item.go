package appkit

// #include "collection_layout_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutItem interface {
	objc.Object
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value CollectionLayoutEdgeSpacing)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
}

type NSCollectionLayoutItem struct {
	objc.NSObject
}

func MakeCollectionLayoutItem(ptr unsafe.Pointer) NSCollectionLayoutItem {
	return NSCollectionLayoutItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutItem() NSCollectionLayoutItem {
	return MakeCollectionLayoutItem(C.C_CollectionLayoutItem_Alloc())
}

func CollectionLayoutItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) CollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutItem(result_)
}

func CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) CollectionLayoutItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutItem(result_)
}

func (n NSCollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutItem_LayoutSize(n.Ptr())
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutItem_SupplementaryItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutSupplementaryItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutSupplementaryItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutItem_EdgeSpacing(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutItem) SetEdgeSpacing(value CollectionLayoutEdgeSpacing) {
	C.C_NSCollectionLayoutItem_SetEdgeSpacing(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	result_ := C.C_NSCollectionLayoutItem_ContentInsets(n.Ptr())
	return FromNSDirectionalEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	C.C_NSCollectionLayoutItem_SetContentInsets(n.Ptr(), *(*C.NSDirectionalEdgeInsets)(ToNSDirectionalEdgeInsetsPointer(value)))
}
