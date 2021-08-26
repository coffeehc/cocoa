package appkit

// #include "collection_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type CollectionViewDataSource struct {
	NumberOfSectionsInCollectionView                             func(collectionView CollectionView) int
	CollectionView_NumberOfItemsInSection                        func(collectionView CollectionView, section int) int                                   // required
	CollectionView_ItemForRepresentedObjectAtIndexPath           func(collectionView CollectionView, indexPath foundation.IndexPath) CollectionViewItem // required
	CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath func(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) View
}

func (delegate *CollectionViewDataSource) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewDataSource_NumberOfSectionsInCollectionView
func collectionViewDataSource_NumberOfSectionsInCollectionView(hp C.uintptr_t, collectionView unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDataSource)
	result := delegate.NumberOfSectionsInCollectionView(MakeCollectionView(collectionView))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_NumberOfItemsInSection
func collectionViewDataSource_CollectionView_NumberOfItemsInSection(hp C.uintptr_t, collectionView unsafe.Pointer, section C.int) C.int {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDataSource)
	result := delegate.CollectionView_NumberOfItemsInSection(MakeCollectionView(collectionView), int(section))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath
func collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDataSource)
	result := delegate.CollectionView_ItemForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath
func collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDataSource)
	result := delegate.CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(MakeCollectionView(collectionView), CollectionViewSupplementaryElementKind(foundation.MakeString(kind).String()), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export CollectionViewDataSource_RespondsTo
func CollectionViewDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewDataSource)
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
