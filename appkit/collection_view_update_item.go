package appkit

// #include "collection_view_update_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewUpdateItem interface {
	objc.Object
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
}

type NSCollectionViewUpdateItem struct {
	objc.NSObject
}

func MakeCollectionViewUpdateItem(ptr unsafe.Pointer) NSCollectionViewUpdateItem {
	return NSCollectionViewUpdateItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewUpdateItem() NSCollectionViewUpdateItem {
	return MakeCollectionViewUpdateItem(C.C_CollectionViewUpdateItem_Alloc())
}

func (n NSCollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Init(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	result_ := C.C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	result_ := C.C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	result_ := C.C_NSCollectionViewUpdateItem_UpdateAction(n.Ptr())
	return CollectionUpdateAction(int(result_))
}
