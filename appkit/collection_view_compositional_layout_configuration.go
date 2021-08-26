package appkit

// #include "collection_view_compositional_layout_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewCompositionalLayoutConfiguration interface {
	objc.Object
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() coregraphics.Float
	SetInterSectionSpacing(value coregraphics.Float)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
}

type NSCollectionViewCompositionalLayoutConfiguration struct {
	objc.NSObject
}

func MakeCollectionViewCompositionalLayoutConfiguration(ptr unsafe.Pointer) NSCollectionViewCompositionalLayoutConfiguration {
	return NSCollectionViewCompositionalLayoutConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewCompositionalLayoutConfiguration() NSCollectionViewCompositionalLayoutConfiguration {
	return MakeCollectionViewCompositionalLayoutConfiguration(C.C_CollectionViewCompositionalLayoutConfiguration_Alloc())
}

func (n NSCollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_Init(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(n.Ptr())
	return CollectionViewScrollDirection(int(result_))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(n.Ptr())
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

func (n NSCollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(n.Ptr(), cValue)
}
