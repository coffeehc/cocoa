package appkit

// #include "collection_view_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewLayout interface {
	objc.Object
	PrepareLayout()
	LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point
	TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point
	PrepareForCollectionViewUpdates(updateItems []CollectionViewUpdateItem)
	FinalizeCollectionViewUpdates()
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IndexPath) CollectionViewLayoutAttributes
	InvalidateLayout()
	InvalidateLayoutWithContext(context CollectionViewLayoutInvalidationContext)
	ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool
	ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes CollectionViewLayoutAttributes, originalAttributes CollectionViewLayoutAttributes) bool
	InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext
	InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes CollectionViewLayoutAttributes, originalAttributes CollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext
	PrepareForAnimatedBoundsChange(oldBounds foundation.Rect)
	FinalizeAnimatedBoundsChange()
	RegisterNib_ForDecorationViewOfKind(nib Nib, elementKind CollectionViewDecorationElementKind)
	PrepareForTransitionFromLayout(oldLayout CollectionViewLayout)
	PrepareForTransitionToLayout(newLayout CollectionViewLayout)
	FinalizeLayoutTransition()
	CollectionView() CollectionView
	CollectionViewContentSize() foundation.Size
}

type NSCollectionViewLayout struct {
	objc.NSObject
}

func MakeCollectionViewLayout(ptr unsafe.Pointer) NSCollectionViewLayout {
	return NSCollectionViewLayout{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewLayout() NSCollectionViewLayout {
	return MakeCollectionViewLayout(C.C_CollectionViewLayout_Alloc())
}

func (n NSCollectionViewLayout) Init() CollectionViewLayout {
	result_ := C.C_NSCollectionViewLayout_Init(n.Ptr())
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionViewLayout) PrepareLayout() {
	C.C_NSCollectionViewLayout_PrepareLayout(n.Ptr())
}

func (n NSCollectionViewLayout) LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForElementsInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]CollectionViewLayoutAttributes, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionViewLayoutAttributes(r)
	}
	return goResult_
}

func (n NSCollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForItemAtIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForDecorationViewOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForDropTargetAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(pointInCollectionView))))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_LayoutAttributesForInterItemGapBeforeIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point {
	result_ := C.C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(proposedContentOffset))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayout) TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point {
	result_ := C.C_NSCollectionViewLayout_TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(proposedContentOffset))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(velocity))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []CollectionViewUpdateItem) {
	var cUpdateItems C.Array
	if len(updateItems) > 0 {
		cUpdateItemsData := make([]unsafe.Pointer, len(updateItems))
		for idx, v := range updateItems {
			cUpdateItemsData[idx] = objc.ExtractPtr(v)
		}
		cUpdateItems.data = unsafe.Pointer(&cUpdateItemsData[0])
		cUpdateItems.len = C.int(len(updateItems))
	}
	C.C_NSCollectionViewLayout_PrepareForCollectionViewUpdates(n.Ptr(), cUpdateItems)
}

func (n NSCollectionViewLayout) FinalizeCollectionViewUpdates() {
	C.C_NSCollectionViewLayout_FinalizeCollectionViewUpdates(n.Ptr())
}

func (n NSCollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	result_ := C.C_NSCollectionViewLayout_IndexPathsToInsertForSupplementaryViewOfKind(n.Ptr(), foundation.NewString(string(elementKind)).Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	result_ := C.C_NSCollectionViewLayout_IndexPathsToInsertForDecorationViewOfKind(n.Ptr(), foundation.NewString(string(elementKind)).Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionViewLayout) InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingItemAtIndexPath(n.Ptr(), objc.ExtractPtr(itemIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(elementIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(decorationIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	result_ := C.C_NSCollectionViewLayout_IndexPathsToDeleteForSupplementaryViewOfKind(n.Ptr(), foundation.NewString(string(elementKind)).Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	result_ := C.C_NSCollectionViewLayout_IndexPathsToDeleteForDecorationViewOfKind(n.Ptr(), foundation.NewString(string(elementKind)).Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingItemAtIndexPath(n.Ptr(), objc.ExtractPtr(itemIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(elementIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayout_FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(decorationIndexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayout) InvalidateLayout() {
	C.C_NSCollectionViewLayout_InvalidateLayout(n.Ptr())
}

func (n NSCollectionViewLayout) InvalidateLayoutWithContext(context CollectionViewLayoutInvalidationContext) {
	C.C_NSCollectionViewLayout_InvalidateLayoutWithContext(n.Ptr(), objc.ExtractPtr(context))
}

func (n NSCollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool {
	result_ := C.C_NSCollectionViewLayout_ShouldInvalidateLayoutForBoundsChange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(newBounds))))
	return bool(result_)
}

func (n NSCollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes CollectionViewLayoutAttributes, originalAttributes CollectionViewLayoutAttributes) bool {
	result_ := C.C_NSCollectionViewLayout_ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(n.Ptr(), objc.ExtractPtr(preferredAttributes), objc.ExtractPtr(originalAttributes))
	return bool(result_)
}

func (n NSCollectionViewLayout) InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayout_InvalidationContextForBoundsChange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(newBounds))))
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayout) InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes CollectionViewLayoutAttributes, originalAttributes CollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewLayout_InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(n.Ptr(), objc.ExtractPtr(preferredAttributes), objc.ExtractPtr(originalAttributes))
	return MakeCollectionViewLayoutInvalidationContext(result_)
}

func (n NSCollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds foundation.Rect) {
	C.C_NSCollectionViewLayout_PrepareForAnimatedBoundsChange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(oldBounds))))
}

func (n NSCollectionViewLayout) FinalizeAnimatedBoundsChange() {
	C.C_NSCollectionViewLayout_FinalizeAnimatedBoundsChange(n.Ptr())
}

func (n NSCollectionViewLayout) RegisterNib_ForDecorationViewOfKind(nib Nib, elementKind CollectionViewDecorationElementKind) {
	C.C_NSCollectionViewLayout_RegisterNib_ForDecorationViewOfKind(n.Ptr(), objc.ExtractPtr(nib), foundation.NewString(string(elementKind)).Ptr())
}

func (n NSCollectionViewLayout) PrepareForTransitionFromLayout(oldLayout CollectionViewLayout) {
	C.C_NSCollectionViewLayout_PrepareForTransitionFromLayout(n.Ptr(), objc.ExtractPtr(oldLayout))
}

func (n NSCollectionViewLayout) PrepareForTransitionToLayout(newLayout CollectionViewLayout) {
	C.C_NSCollectionViewLayout_PrepareForTransitionToLayout(n.Ptr(), objc.ExtractPtr(newLayout))
}

func (n NSCollectionViewLayout) FinalizeLayoutTransition() {
	C.C_NSCollectionViewLayout_FinalizeLayoutTransition(n.Ptr())
}

func (n NSCollectionViewLayout) CollectionView() CollectionView {
	result_ := C.C_NSCollectionViewLayout_CollectionView(n.Ptr())
	return MakeCollectionView(result_)
}

func (n NSCollectionViewLayout) CollectionViewContentSize() foundation.Size {
	result_ := C.C_NSCollectionViewLayout_CollectionViewContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}
