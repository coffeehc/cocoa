package appkit

// #include "collection_view_layout_attributes.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewLayoutAttributes interface {
	objc.Object
	RepresentedElementCategory() CollectionElementCategory
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IndexPath)
	RepresentedElementKind() string
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	Size() foundation.Size
	SetSize(value foundation.Size)
	Alpha() coregraphics.Float
	SetAlpha(value coregraphics.Float)
	IsHidden() bool
	SetHidden(value bool)
	ZIndex() int
	SetZIndex(value int)
}

type NSCollectionViewLayoutAttributes struct {
	objc.NSObject
}

func MakeCollectionViewLayoutAttributes(ptr unsafe.Pointer) *NSCollectionViewLayoutAttributes {
	if ptr == nil {
		return nil
	}
	return &NSCollectionViewLayoutAttributes{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCollectionViewLayoutAttributes() *NSCollectionViewLayoutAttributes {
	return MakeCollectionViewLayoutAttributes(C.C_CollectionViewLayoutAttributes_Alloc())
}

func (n *NSCollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_Init(n.Ptr())
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(foundation.NewString(string(decorationViewKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n *NSCollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	result_ := C.C_NSCollectionViewLayoutAttributes_RepresentedElementCategory(n.Ptr())
	return CollectionElementCategory(int(result_))
}

func (n *NSCollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	result_ := C.C_NSCollectionViewLayoutAttributes_IndexPath(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n *NSCollectionViewLayoutAttributes) SetIndexPath(value foundation.IndexPath) {
	C.C_NSCollectionViewLayoutAttributes_SetIndexPath(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSCollectionViewLayoutAttributes) RepresentedElementKind() string {
	result_ := C.C_NSCollectionViewLayoutAttributes_RepresentedElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSCollectionViewLayoutAttributes) Frame() foundation.Rect {
	result_ := C.C_NSCollectionViewLayoutAttributes_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	C.C_NSCollectionViewLayoutAttributes_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSCollectionViewLayoutAttributes) Size() foundation.Size {
	result_ := C.C_NSCollectionViewLayoutAttributes_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	C.C_NSCollectionViewLayoutAttributes_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSCollectionViewLayoutAttributes) Alpha() coregraphics.Float {
	result_ := C.C_NSCollectionViewLayoutAttributes_Alpha(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSCollectionViewLayoutAttributes) SetAlpha(value coregraphics.Float) {
	C.C_NSCollectionViewLayoutAttributes_SetAlpha(n.Ptr(), C.double(float64(value)))
}

func (n *NSCollectionViewLayoutAttributes) IsHidden() bool {
	result_ := C.C_NSCollectionViewLayoutAttributes_IsHidden(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionViewLayoutAttributes) SetHidden(value bool) {
	C.C_NSCollectionViewLayoutAttributes_SetHidden(n.Ptr(), C.bool(value))
}

func (n *NSCollectionViewLayoutAttributes) ZIndex() int {
	result_ := C.C_NSCollectionViewLayoutAttributes_ZIndex(n.Ptr())
	return int(result_)
}

func (n *NSCollectionViewLayoutAttributes) SetZIndex(value int) {
	C.C_NSCollectionViewLayoutAttributes_SetZIndex(n.Ptr(), C.int(value))
}
