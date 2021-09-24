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
	result_ := C.C_NSCollectionViewLayout_AllocCollectionViewLayout()
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionViewLayout) Init() NSCollectionViewLayout {
	result_ := C.C_NSCollectionViewLayout_Init(n.Ptr())
	return MakeCollectionViewLayout(result_)
}

func NewCollectionViewLayout() NSCollectionViewLayout {
	result_ := C.C_NSCollectionViewLayout_NewCollectionViewLayout()
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionViewLayout) Autorelease() NSCollectionViewLayout {
	result_ := C.C_NSCollectionViewLayout_Autorelease(n.Ptr())
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionViewLayout) Retain() NSCollectionViewLayout {
	result_ := C.C_NSCollectionViewLayout_Retain(n.Ptr())
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
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
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

func MakeCollectionViewLayoutAttributes(ptr unsafe.Pointer) NSCollectionViewLayoutAttributes {
	return NSCollectionViewLayoutAttributes{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IndexPath) NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(foundation.NewString(string(decorationViewKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func AllocCollectionViewLayoutAttributes() NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_AllocCollectionViewLayoutAttributes()
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayoutAttributes) Init() NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_Init(n.Ptr())
	return MakeCollectionViewLayoutAttributes(result_)
}

func NewCollectionViewLayoutAttributes() NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_NewCollectionViewLayoutAttributes()
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayoutAttributes) Autorelease() NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_Autorelease(n.Ptr())
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayoutAttributes) Retain() NSCollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionViewLayoutAttributes_Retain(n.Ptr())
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	result_ := C.C_NSCollectionViewLayoutAttributes_RepresentedElementCategory(n.Ptr())
	return CollectionElementCategory(int(result_))
}

func (n NSCollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	result_ := C.C_NSCollectionViewLayoutAttributes_IndexPath(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionViewLayoutAttributes) SetIndexPath(value foundation.IndexPath) {
	C.C_NSCollectionViewLayoutAttributes_SetIndexPath(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionViewLayoutAttributes) RepresentedElementKind() string {
	result_ := C.C_NSCollectionViewLayoutAttributes_RepresentedElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionViewLayoutAttributes) Frame() foundation.Rect {
	result_ := C.C_NSCollectionViewLayoutAttributes_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	C.C_NSCollectionViewLayoutAttributes_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSCollectionViewLayoutAttributes) Size() foundation.Size {
	result_ := C.C_NSCollectionViewLayoutAttributes_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	C.C_NSCollectionViewLayoutAttributes_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewLayoutAttributes) Alpha() coregraphics.Float {
	result_ := C.C_NSCollectionViewLayoutAttributes_Alpha(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewLayoutAttributes) SetAlpha(value coregraphics.Float) {
	C.C_NSCollectionViewLayoutAttributes_SetAlpha(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewLayoutAttributes) IsHidden() bool {
	result_ := C.C_NSCollectionViewLayoutAttributes_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewLayoutAttributes) SetHidden(value bool) {
	C.C_NSCollectionViewLayoutAttributes_SetHidden(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewLayoutAttributes) ZIndex() int {
	result_ := C.C_NSCollectionViewLayoutAttributes_ZIndex(n.Ptr())
	return int(result_)
}

func (n NSCollectionViewLayoutAttributes) SetZIndex(value int) {
	C.C_NSCollectionViewLayoutAttributes_SetZIndex(n.Ptr(), C.int(value))
}

type CollectionViewUpdateItem interface {
	objc.Object
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
}

type NSCollectionViewUpdateItem struct {
	objc.NSObject
}

func MakeCollectionViewUpdateItem(ptr unsafe.Pointer) NSCollectionViewUpdateItem {
	return NSCollectionViewUpdateItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewUpdateItem() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_AllocCollectionViewUpdateItem()
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Init() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Init(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func NewCollectionViewUpdateItem() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_NewCollectionViewUpdateItem()
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Autorelease() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Autorelease(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) Retain() NSCollectionViewUpdateItem {
	result_ := C.C_NSCollectionViewUpdateItem_Retain(n.Ptr())
	return MakeCollectionViewUpdateItem(result_)
}

func (n NSCollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	result_ := C.C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	result_ := C.C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	result_ := C.C_NSCollectionViewUpdateItem_UpdateAction(n.Ptr())
	return CollectionUpdateAction(int(result_))
}

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
