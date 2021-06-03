package appkit

// #include "collection_view_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewItem interface {
	ViewController
	IsSelected() bool
	SetSelected(value bool)
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
}

type NSCollectionViewItem struct {
	NSViewController
}

func MakeCollectionViewItem(ptr unsafe.Pointer) NSCollectionViewItem {
	return NSCollectionViewItem{
		NSViewController: MakeViewController(ptr),
	}
}

func AllocCollectionViewItem() NSCollectionViewItem {
	return MakeCollectionViewItem(C.C_CollectionViewItem_Alloc())
}

func (n NSCollectionViewItem) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) CollectionViewItem {
	result_ := C.C_NSCollectionViewItem_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) InitWithCoder(coder foundation.Coder) CollectionViewItem {
	result_ := C.C_NSCollectionViewItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) Init() CollectionViewItem {
	result_ := C.C_NSCollectionViewItem_Init(n.Ptr())
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) IsSelected() bool {
	result_ := C.C_NSCollectionViewItem_IsSelected(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewItem) SetSelected(value bool) {
	C.C_NSCollectionViewItem_SetSelected(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	result_ := C.C_NSCollectionViewItem_HighlightState(n.Ptr())
	return CollectionViewItemHighlightState(int(result_))
}

func (n NSCollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	C.C_NSCollectionViewItem_SetHighlightState(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionViewItem) CollectionView() CollectionView {
	result_ := C.C_NSCollectionViewItem_CollectionView(n.Ptr())
	return MakeCollectionView(result_)
}

func (n NSCollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	result_ := C.C_NSCollectionViewItem_DraggingImageComponents(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}
