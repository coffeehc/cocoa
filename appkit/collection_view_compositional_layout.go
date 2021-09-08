package appkit

// #include "collection_view_compositional_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewCompositionalLayout interface {
	CollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value CollectionViewCompositionalLayoutConfiguration)
}

type NSCollectionViewCompositionalLayout struct {
	NSCollectionViewLayout
}

func MakeCollectionViewCompositionalLayout(ptr unsafe.Pointer) NSCollectionViewCompositionalLayout {
	return NSCollectionViewCompositionalLayout{
		NSCollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (n NSCollectionViewCompositionalLayout) InitWithSection(section CollectionLayoutSection) NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_InitWithSection(n.Ptr(), objc.ExtractPtr(section))
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) InitWithSection_Configuration(section CollectionLayoutSection, configuration CollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(n.Ptr(), objc.ExtractPtr(section), objc.ExtractPtr(configuration))
	return MakeCollectionViewCompositionalLayout(result_)
}

func AllocCollectionViewCompositionalLayout() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_AllocCollectionViewCompositionalLayout()
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Autorelease() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_Autorelease(n.Ptr())
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Retain() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_Retain(n.Ptr())
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayout_Configuration(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayout) SetConfiguration(value CollectionViewCompositionalLayoutConfiguration) {
	C.C_NSCollectionViewCompositionalLayout_SetConfiguration(n.Ptr(), objc.ExtractPtr(value))
}
