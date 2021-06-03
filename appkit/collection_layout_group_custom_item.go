package appkit

// #include "collection_layout_group_custom_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutGroupCustomItem interface {
	objc.Object
	Frame() foundation.Rect
	ZIndex() int
}

type NSCollectionLayoutGroupCustomItem struct {
	objc.NSObject
}

func MakeCollectionLayoutGroupCustomItem(ptr unsafe.Pointer) NSCollectionLayoutGroupCustomItem {
	return NSCollectionLayoutGroupCustomItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutGroupCustomItem() NSCollectionLayoutGroupCustomItem {
	return MakeCollectionLayoutGroupCustomItem(C.C_CollectionLayoutGroupCustomItem_Alloc())
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrame(frame foundation.Rect) CollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))))
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), C.int(zIndex))
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func (n NSCollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutGroupCustomItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_ZIndex(n.Ptr())
	return int(result_)
}
