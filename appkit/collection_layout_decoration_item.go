package appkit

// #include "collection_layout_decoration_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
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

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) *NSCollectionLayoutDecorationItem {
	if ptr == nil {
		return nil
	}
	return &NSCollectionLayoutDecorationItem{
		NSCollectionLayoutItem: *MakeCollectionLayoutItem(ptr),
	}
}

func AllocCollectionLayoutDecorationItem() *NSCollectionLayoutDecorationItem {
	return MakeCollectionLayoutDecorationItem(C.C_CollectionLayoutDecorationItem_Alloc())
}

func CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(foundation.NewString(elementKind).Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n *NSCollectionLayoutDecorationItem) ElementKind() string {
	result_ := C.C_NSCollectionLayoutDecorationItem_ElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSCollectionLayoutDecorationItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutDecorationItem_ZIndex(n.Ptr())
	return int(result_)
}

func (n *NSCollectionLayoutDecorationItem) SetZIndex(value int) {
	C.C_NSCollectionLayoutDecorationItem_SetZIndex(n.Ptr(), C.int(value))
}
