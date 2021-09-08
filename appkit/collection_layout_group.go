package appkit

// #include "collection_layout_group.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutGroup interface {
	CollectionLayoutItem
	VisualDescription() string
	Subitems() []CollectionLayoutItem
	SetSupplementaryItems(value []CollectionLayoutSupplementaryItem)
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value CollectionLayoutSpacing)
}

type NSCollectionLayoutGroup struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutGroup(ptr unsafe.Pointer) NSCollectionLayoutGroup {
	return NSCollectionLayoutGroup{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) NSCollectionLayoutGroup {
	var cSubitems C.Array
	if len(subitems) > 0 {
		cSubitemsData := make([]unsafe.Pointer, len(subitems))
		for idx, v := range subitems {
			cSubitemsData[idx] = objc.ExtractPtr(v)
		}
		cSubitems.data = unsafe.Pointer(&cSubitemsData[0])
		cSubitems.len = C.int(len(subitems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) NSCollectionLayoutGroup {
	var cSubitems C.Array
	if len(subitems) > 0 {
		cSubitemsData := make([]unsafe.Pointer, len(subitems))
		for idx, v := range subitems {
			cSubitemsData[idx] = objc.ExtractPtr(v)
		}
		cSubitems.data = unsafe.Pointer(&cSubitemsData[0])
		cSubitems.len = C.int(len(subitems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutGroup {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutGroup(result_)
}

func AllocCollectionLayoutGroup() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_AllocCollectionLayoutGroup()
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) Autorelease() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_Autorelease(n.Ptr())
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) Retain() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_Retain(n.Ptr())
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) VisualDescription() string {
	result_ := C.C_NSCollectionLayoutGroup_VisualDescription(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutGroup_Subitems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutGroup) SetSupplementaryItems(value []CollectionLayoutSupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutGroup_SetSupplementaryItems(n.Ptr(), cValue)
}

func (n NSCollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutGroup_InterItemSpacing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutGroup) SetInterItemSpacing(value CollectionLayoutSpacing) {
	C.C_NSCollectionLayoutGroup_SetInterItemSpacing(n.Ptr(), objc.ExtractPtr(value))
}
