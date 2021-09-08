package appkit

// #include "collection_view_layout_invalidation_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewLayoutInvalidationContext interface {
	objc.Object
	InvalidateItemsAtIndexPaths(indexPaths foundation.Set)
	InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.Set)
	InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.Set)
	InvalidateEverything() bool
	InvalidateDataSourceCounts() bool
	ContentOffsetAdjustment() foundation.Point
	SetContentOffsetAdjustment(value foundation.Point)
	ContentSizeAdjustment() foundation.Size
	SetContentSizeAdjustment(value foundation.Size)
	InvalidatedItemIndexPaths() foundation.Set
}

type NSCollectionViewLayoutInvalidationContext struct {
	objc.NSObject
}

func MakeCollectionViewLayoutInvalidationContext(ptr unsafe.Pointer) NSCollectionViewLayoutInvalidationContext {
	return NSCollectionViewLayoutInvalidationContext{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewLayoutInvalidationContext() NSCollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_AllocCollectionViewLayoutInvalidationContext()
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) Init() NSCollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_Init(n.Ptr())
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func NewCollectionViewLayoutInvalidationContext() NSCollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_NewCollectionViewLayoutInvalidationContext()
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) Autorelease() NSCollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_Autorelease(n.Ptr())
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) Retain() NSCollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_Retain(n.Ptr())
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.Set) {
	C.C_NSCollectionViewLayoutInvalidationContext_InvalidateItemsAtIndexPaths(n.Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.Set) {
	C.C_NSCollectionViewLayoutInvalidationContext_InvalidateSupplementaryElementsOfKind_AtIndexPaths(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.Set) {
	C.C_NSCollectionViewLayoutInvalidationContext_InvalidateDecorationElementsOfKind_AtIndexPaths(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_InvalidateEverything(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_InvalidateDataSourceCounts(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_ContentOffsetAdjustment(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	C.C_NSCollectionViewLayoutInvalidationContext_SetContentOffsetAdjustment(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n NSCollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_ContentSizeAdjustment(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	C.C_NSCollectionViewLayoutInvalidationContext_SetContentSizeAdjustment(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	result_ := C.C_NSCollectionViewLayoutInvalidationContext_InvalidatedItemIndexPaths(n.Ptr())
	return foundation.MakeSet(result_)
}
