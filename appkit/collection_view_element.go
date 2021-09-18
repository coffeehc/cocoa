package appkit

// #include "collection_view_element.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type CollectionViewElement struct {
	PrepareForReuse                            func()
	PreferredLayoutAttributesFittingAttributes func(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	ApplyLayoutAttributes                      func(layoutAttributes CollectionViewLayoutAttributes)
	WillTransitionFromLayout_ToLayout          func(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	DidTransitionFromLayout_ToLayout           func(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
}

func (delegate *CollectionViewElement) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewElement(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewElement_PrepareForReuse
func collectionViewElement_PrepareForReuse(hp C.uintptr_t) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
	delegate.PrepareForReuse()
}

//export collectionViewElement_PreferredLayoutAttributesFittingAttributes
func collectionViewElement_PreferredLayoutAttributesFittingAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
	result := delegate.PreferredLayoutAttributesFittingAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
	return objc.ExtractPtr(result)
}

//export collectionViewElement_ApplyLayoutAttributes
func collectionViewElement_ApplyLayoutAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
	delegate.ApplyLayoutAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
}

//export collectionViewElement_WillTransitionFromLayout_ToLayout
func collectionViewElement_WillTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
	delegate.WillTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewElement_DidTransitionFromLayout_ToLayout
func collectionViewElement_DidTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
	delegate.DidTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export CollectionViewElement_RespondsTo
func CollectionViewElement_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewElement)
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
