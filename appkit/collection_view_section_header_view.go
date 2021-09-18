package appkit

// #include "collection_view_section_header_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type CollectionViewSectionHeaderView struct {
	PrepareForReuse                            func()
	PreferredLayoutAttributesFittingAttributes func(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	ApplyLayoutAttributes                      func(layoutAttributes CollectionViewLayoutAttributes)
	WillTransitionFromLayout_ToLayout          func(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	DidTransitionFromLayout_ToLayout           func(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
}

func (delegate *CollectionViewSectionHeaderView) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewSectionHeaderView(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewSectionHeaderView_PrepareForReuse
func collectionViewSectionHeaderView_PrepareForReuse(hp C.uintptr_t) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	delegate.PrepareForReuse()
}

//export collectionViewSectionHeaderView_PreferredLayoutAttributesFittingAttributes
func collectionViewSectionHeaderView_PreferredLayoutAttributesFittingAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	result := delegate.PreferredLayoutAttributesFittingAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
	return objc.ExtractPtr(result)
}

//export collectionViewSectionHeaderView_ApplyLayoutAttributes
func collectionViewSectionHeaderView_ApplyLayoutAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	delegate.ApplyLayoutAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
}

//export collectionViewSectionHeaderView_WillTransitionFromLayout_ToLayout
func collectionViewSectionHeaderView_WillTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	delegate.WillTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewSectionHeaderView_DidTransitionFromLayout_ToLayout
func collectionViewSectionHeaderView_DidTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	delegate.DidTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export CollectionViewSectionHeaderView_RespondsTo
func CollectionViewSectionHeaderView_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewSectionHeaderView)
	switch selName {
	case "prepareForReuse":
		return delegate.PrepareForReuse != nil
	case "preferredLayoutAttributesFittingAttributes:":
		return delegate.PreferredLayoutAttributesFittingAttributes != nil
	case "applyLayoutAttributes:":
		return delegate.ApplyLayoutAttributes != nil
	case "willTransitionFromLayout:toLayout:":
		return delegate.WillTransitionFromLayout_ToLayout != nil
	case "didTransitionFromLayout:toLayout:":
		return delegate.DidTransitionFromLayout_ToLayout != nil
	default:
		return false
	}
}
