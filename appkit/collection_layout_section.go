package appkit

// #include "collection_layout_section.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionLayoutSection interface {
	objc.Object
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	InterGroupSpacing() coregraphics.Float
	SetInterGroupSpacing(value coregraphics.Float)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []CollectionLayoutDecorationItem)
}

type NSCollectionLayoutSection struct {
	objc.NSObject
}

func MakeCollectionLayoutSection(ptr unsafe.Pointer) NSCollectionLayoutSection {
	return NSCollectionLayoutSection{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionLayoutSection() NSCollectionLayoutSection {
	return MakeCollectionLayoutSection(C.C_CollectionLayoutSection_Alloc())
}

func CollectionLayoutSection_SectionWithGroup(group CollectionLayoutGroup) CollectionLayoutSection {
	result_ := C.C_NSCollectionLayoutSection_CollectionLayoutSection_SectionWithGroup(objc.ExtractPtr(group))
	return MakeCollectionLayoutSection(result_)
}

func (n NSCollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	result_ := C.C_NSCollectionLayoutSection_OrthogonalScrollingBehavior(n.Ptr())
	return CollectionLayoutSectionOrthogonalScrollingBehavior(int(result_))
}

func (n NSCollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	C.C_NSCollectionLayoutSection_SetOrthogonalScrollingBehavior(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionLayoutSection) InterGroupSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutSection_InterGroupSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutSection) SetInterGroupSpacing(value coregraphics.Float) {
	C.C_NSCollectionLayoutSection_SetInterGroupSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	result_ := C.C_NSCollectionLayoutSection_ContentInsets(n.Ptr())
	return FromNSDirectionalEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	C.C_NSCollectionLayoutSection_SetContentInsets(n.Ptr(), *(*C.NSDirectionalEdgeInsets)(ToNSDirectionalEdgeInsetsPointer(value)))
}

func (n NSCollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	result_ := C.C_NSCollectionLayoutSection_SupplementariesFollowContentInsets(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	C.C_NSCollectionLayoutSection_SetSupplementariesFollowContentInsets(n.Ptr(), C.bool(value))
}

func (n NSCollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutSection_BoundarySupplementaryItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]CollectionLayoutBoundarySupplementaryItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutBoundarySupplementaryItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutSection) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutSection_SetBoundarySupplementaryItems(n.Ptr(), cValue)
}

func (n NSCollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutSection_DecorationItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]CollectionLayoutDecorationItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutDecorationItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutSection) SetDecorationItems(value []CollectionLayoutDecorationItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutSection_SetDecorationItems(n.Ptr(), cValue)
}
