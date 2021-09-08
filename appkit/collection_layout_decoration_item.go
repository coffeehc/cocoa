package appkit

// #include "collection_layout_decoration_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutDecorationItem interface {
	CollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type NSCollectionLayoutDecorationItem struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) NSCollectionLayoutDecorationItem {
	return NSCollectionLayoutDecorationItem{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(elementKind string) NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(foundation.NewString(elementKind).Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func CollectionLayoutDecorationItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutDecorationItem(result_)
}

func CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutDecorationItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutDecorationItem(result_)
}

func AllocCollectionLayoutDecorationItem() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_AllocCollectionLayoutDecorationItem()
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) Autorelease() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) Retain() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_Retain(n.Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) ElementKind() string {
	result_ := C.C_NSCollectionLayoutDecorationItem_ElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutDecorationItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutDecorationItem_ZIndex(n.Ptr())
	return int(result_)
}

func (n NSCollectionLayoutDecorationItem) SetZIndex(value int) {
	C.C_NSCollectionLayoutDecorationItem_SetZIndex(n.Ptr(), C.int(value))
}
