package appkit

// #include "collection_view_flow_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type CollectionViewFlowLayout interface {
	CollectionViewLayout
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	MinimumLineSpacing() coregraphics.Float
	SetMinimumLineSpacing(value coregraphics.Float)
	MinimumInteritemSpacing() coregraphics.Float
	SetMinimumInteritemSpacing(value coregraphics.Float)
	EstimatedItemSize() foundation.Size
	SetEstimatedItemSize(value foundation.Size)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	HeaderReferenceSize() foundation.Size
	SetHeaderReferenceSize(value foundation.Size)
	FooterReferenceSize() foundation.Size
	SetFooterReferenceSize(value foundation.Size)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
}

type NSCollectionViewFlowLayout struct {
	NSCollectionViewLayout
}

func MakeCollectionViewFlowLayout(ptr unsafe.Pointer) *NSCollectionViewFlowLayout {
	if ptr == nil {
		return nil
	}
	return &NSCollectionViewFlowLayout{
		NSCollectionViewLayout: *MakeCollectionViewLayout(ptr),
	}
}

func AllocCollectionViewFlowLayout() *NSCollectionViewFlowLayout {
	return MakeCollectionViewFlowLayout(C.C_CollectionViewFlowLayout_Alloc())
}

func (n *NSCollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_Init(n.Ptr())
	return MakeCollectionViewFlowLayout(result_)
}

func (n *NSCollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	C.C_NSCollectionViewFlowLayout_CollapseSectionAtIndex(n.Ptr(), C.uint(sectionIndex))
}

func (n *NSCollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	C.C_NSCollectionViewFlowLayout_ExpandSectionAtIndex(n.Ptr(), C.uint(sectionIndex))
}

func (n *NSCollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionAtIndexIsCollapsed(n.Ptr(), C.uint(sectionIndex))
	return bool(result_)
}

func (n *NSCollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	result_ := C.C_NSCollectionViewFlowLayout_ScrollDirection(n.Ptr())
	return CollectionViewScrollDirection(int(result_))
}

func (n *NSCollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	C.C_NSCollectionViewFlowLayout_SetScrollDirection(n.Ptr(), C.int(int(value)))
}

func (n *NSCollectionViewFlowLayout) MinimumLineSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewFlowLayout_MinimumLineSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSCollectionViewFlowLayout) SetMinimumLineSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewFlowLayout_SetMinimumLineSpacing(n.Ptr(), C.double(float64(value)))
}

func (n *NSCollectionViewFlowLayout) MinimumInteritemSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewFlowLayout_MinimumInteritemSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSCollectionViewFlowLayout) SetMinimumInteritemSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewFlowLayout_SetMinimumInteritemSpacing(n.Ptr(), C.double(float64(value)))
}

func (n *NSCollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_EstimatedItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetEstimatedItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSCollectionViewFlowLayout) ItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_ItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSCollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	result_ := C.C_NSCollectionViewFlowLayout_SectionInset(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n *NSCollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	C.C_NSCollectionViewFlowLayout_SetSectionInset(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSCollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_HeaderReferenceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetHeaderReferenceSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSCollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_FooterReferenceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetFooterReferenceSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSCollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionFootersPinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	C.C_NSCollectionViewFlowLayout_SetSectionFootersPinToVisibleBounds(n.Ptr(), C.bool(value))
}

func (n *NSCollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionHeadersPinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	C.C_NSCollectionViewFlowLayout_SetSectionHeadersPinToVisibleBounds(n.Ptr(), C.bool(value))
}
