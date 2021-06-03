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

func AllocCollectionLayoutGroup() NSCollectionLayoutGroup {
	return MakeCollectionLayoutGroup(C.C_CollectionLayoutGroup_Alloc())
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) CollectionLayoutGroup {
	cSubitemsData := make([]unsafe.Pointer, len(subitems))
	for idx, v := range subitems {
		cSubitemsData[idx] = objc.ExtractPtr(v)
	}
	cSubitems := C.Array{data: unsafe.Pointer(&cSubitemsData[0]), len: C.int(len(subitems))}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) CollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) CollectionLayoutGroup {
	cSubitemsData := make([]unsafe.Pointer, len(subitems))
	for idx, v := range subitems {
		cSubitemsData[idx] = objc.ExtractPtr(v)
	}
	cSubitems := C.Array{data: unsafe.Pointer(&cSubitemsData[0]), len: C.int(len(subitems))}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) CollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) VisualDescription() string {
	result_ := C.C_NSCollectionLayoutGroup_VisualDescription(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutGroup_Subitems(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]CollectionLayoutItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutGroup) SetSupplementaryItems(value []CollectionLayoutSupplementaryItem) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSCollectionLayoutGroup_SetSupplementaryItems(n.Ptr(), cValue)
}

func (n NSCollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutGroup_InterItemSpacing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutGroup) SetInterItemSpacing(value CollectionLayoutSpacing) {
	C.C_NSCollectionLayoutGroup_SetInterItemSpacing(n.Ptr(), objc.ExtractPtr(value))
}
