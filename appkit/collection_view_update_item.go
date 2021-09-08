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
	result_ := C.C_NSCollectionViewUpdateItem_AllocCollectionViewUpdateItem()
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Init() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Init(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func NewCollectionViewUpdateItem() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_NewCollectionViewUpdateItem()
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Autorelease() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Autorelease(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Retain() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Retain(n.Ptr())
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
