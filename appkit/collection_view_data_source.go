package appkit

// #include "collection_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewDataSource struct {
	NumberOfSectionsInCollectionView                             func(collectionView CollectionView) int
	CollectionView_NumberOfItemsInSection                        func(collectionView CollectionView, section int) int
	CollectionView_ItemForRepresentedObjectAtIndexPath           func(collectionView CollectionView, indexPath foundation.IndexPath) CollectionViewItem
	CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath func(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) View
}

func WrapCollectionViewDataSource(delegate *CollectionViewDataSource) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapCollectionViewDataSource(C.long(id))
	return objc.MakeObject(ptr)
}

//export collectionViewDataSource_NumberOfSectionsInCollectionView
func collectionViewDataSource_NumberOfSectionsInCollectionView(id int64, collectionView unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*CollectionViewDataSource)
	result := delegate.NumberOfSectionsInCollectionView(MakeCollectionView(collectionView))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_NumberOfItemsInSection
func collectionViewDataSource_CollectionView_NumberOfItemsInSection(id int64, collectionView unsafe.Pointer, section C.int) C.int {
	delegate := resources.Get(id).(*CollectionViewDataSource)
	result := delegate.CollectionView_NumberOfItemsInSection(MakeCollectionView(collectionView), int(section))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath
func collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath(id int64, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDataSource)
	result := delegate.CollectionView_ItemForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath
func collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(id int64, collectionView unsafe.Pointer, kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDataSource)
	result := delegate.CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(MakeCollectionView(collectionView), CollectionViewSupplementaryElementKind(foundation.MakeString(kind).String()), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export CollectionViewDataSource_RespondsTo
func CollectionViewDataSource_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*CollectionViewDataSource)
	switch selName {
	case "numberOfSectionsInCollectionView:":
		return delegate.NumberOfSectionsInCollectionView != nil
	case "collectionView:numberOfItemsInSection:":
		return delegate.CollectionView_NumberOfItemsInSection != nil
	case "collectionView:itemForRepresentedObjectAtIndexPath:":
		return delegate.CollectionView_ItemForRepresentedObjectAtIndexPath != nil
	case "collectionView:viewForSupplementaryElementOfKind:atIndexPath:":
		return delegate.CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath != nil
	default:
		return false
	}
}

//export deleteCollectionViewDataSource
func deleteCollectionViewDataSource(id int64) {
	resources.Delete(id)
}
