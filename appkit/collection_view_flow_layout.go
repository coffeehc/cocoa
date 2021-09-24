package appkit

// #include "collection_view_flow_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
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

func MakeCollectionViewFlowLayout(ptr unsafe.Pointer) NSCollectionViewFlowLayout {
	return NSCollectionViewFlowLayout{
		NSCollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func AllocCollectionViewFlowLayout() NSCollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_AllocCollectionViewFlowLayout()
	return MakeCollectionViewFlowLayout(result_)
}

func (n NSCollectionViewFlowLayout) Init() NSCollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_Init(n.Ptr())
	return MakeCollectionViewFlowLayout(result_)
}

func NewCollectionViewFlowLayout() NSCollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_NewCollectionViewFlowLayout()
	return MakeCollectionViewFlowLayout(result_)
}

func (n NSCollectionViewFlowLayout) Autorelease() NSCollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_Autorelease(n.Ptr())
	return MakeCollectionViewFlowLayout(result_)
}

func (n NSCollectionViewFlowLayout) Retain() NSCollectionViewFlowLayout {
	result_ := C.C_NSCollectionViewFlowLayout_Retain(n.Ptr())
	return MakeCollectionViewFlowLayout(result_)
}

func (n NSCollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	C.C_NSCollectionViewFlowLayout_CollapseSectionAtIndex(n.Ptr(), C.uint(sectionIndex))
}

func (n NSCollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	C.C_NSCollectionViewFlowLayout_ExpandSectionAtIndex(n.Ptr(), C.uint(sectionIndex))
}

func (n NSCollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionAtIndexIsCollapsed(n.Ptr(), C.uint(sectionIndex))
	return bool(result_)
}

func (n NSCollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	result_ := C.C_NSCollectionViewFlowLayout_ScrollDirection(n.Ptr())
	return CollectionViewScrollDirection(int(result_))
}

func (n NSCollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	C.C_NSCollectionViewFlowLayout_SetScrollDirection(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionViewFlowLayout) MinimumLineSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewFlowLayout_MinimumLineSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewFlowLayout) SetMinimumLineSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewFlowLayout_SetMinimumLineSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewFlowLayout) MinimumInteritemSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewFlowLayout_MinimumInteritemSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewFlowLayout) SetMinimumInteritemSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewFlowLayout_SetMinimumInteritemSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_EstimatedItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetEstimatedItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewFlowLayout) ItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_ItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	result_ := C.C_NSCollectionViewFlowLayout_SectionInset(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	C.C_NSCollectionViewFlowLayout_SetSectionInset(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n NSCollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_HeaderReferenceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetHeaderReferenceSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	result_ := C.C_NSCollectionViewFlowLayout_FooterReferenceSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	C.C_NSCollectionViewFlowLayout_SetFooterReferenceSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionFootersPinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	C.C_NSCollectionViewFlowLayout_SetSectionFootersPinToVisibleBounds(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	result_ := C.C_NSCollectionViewFlowLayout_SectionHeadersPinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	C.C_NSCollectionViewFlowLayout_SetSectionHeadersPinToVisibleBounds(n.Ptr(), C.bool(value))
}

type CollectionViewFlowLayoutInvalidationContext interface {
	CollectionViewLayoutInvalidationContext
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

type NSCollectionViewFlowLayoutInvalidationContext struct {
	NSCollectionViewLayoutInvalidationContext
}

func MakeCollectionViewFlowLayoutInvalidationContext(ptr unsafe.Pointer) NSCollectionViewFlowLayoutInvalidationContext {
	return NSCollectionViewFlowLayoutInvalidationContext{
		NSCollectionViewLayoutInvalidationContext: MakeCollectionViewLayoutInvalidationContext(ptr),
	}
}

func AllocCollectionViewFlowLayoutInvalidationContext() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_AllocCollectionViewFlowLayoutInvalidationContext()
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Init() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Init(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func NewCollectionViewFlowLayoutInvalidationContext() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_NewCollectionViewFlowLayoutInvalidationContext()
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Autorelease() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Autorelease(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Retain() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Retain(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(n.Ptr(), C.bool(value))
}

type CollectionViewDelegateFlowLayout struct {
	CollectionView_Layout_SizeForItemAtIndexPath                                  func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	CollectionView_Layout_InsetForSectionAtIndex                                  func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	CollectionView_Layout_MinimumLineSpacingForSectionAtIndex                     func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) coregraphics.Float
	CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex                func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) coregraphics.Float
	CollectionView_Layout_ReferenceSizeForHeaderInSection                         func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	CollectionView_Layout_ReferenceSizeForFooterInSection                         func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	CollectionView_ShouldSelectItemsAtIndexPaths                                  func(collectionView CollectionView, indexPaths foundation.Set) foundation.Set
	CollectionView_DidSelectItemsAtIndexPaths                                     func(collectionView CollectionView, indexPaths foundation.Set)
	CollectionView_ShouldDeselectItemsAtIndexPaths                                func(collectionView CollectionView, indexPaths foundation.Set) foundation.Set
	CollectionView_DidDeselectItemsAtIndexPaths                                   func(collectionView CollectionView, indexPaths foundation.Set)
	CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState                 func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.Set
	CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState                    func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath                func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath           func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath        func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	CollectionView_TransitionLayoutForOldLayout_NewLayout                         func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) CollectionViewTransitionLayout
	CollectionView_CanDragItemsAtIndexPaths_WithEvent                             func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	CollectionView_PasteboardWriterForItemAtIndexPath                             func(collectionView CollectionView, indexPath foundation.IndexPath) objc.Object
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths          func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	CollectionView_DraggingSession_EndedAtPoint_DragOperation                     func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	CollectionView_UpdateDraggingItemsForDrag                                     func(collectionView CollectionView, draggingInfo objc.Object)
	CollectionView_AcceptDrop_IndexPath_DropOperation                             func(collectionView CollectionView, draggingInfo objc.Object, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	CollectionView_CanDragItemsAtIndexes_WithEvent                                func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	CollectionView_PasteboardWriterForItemAtIndex                                 func(collectionView CollectionView, index uint) objc.Object
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes             func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	CollectionView_AcceptDrop_Index_DropOperation                                 func(collectionView CollectionView, draggingInfo objc.Object, index int, dropOperation CollectionViewDropOperation) bool
}

func (delegate *CollectionViewDelegateFlowLayout) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewDelegateFlowLayout(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_SizeForItemAtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_Layout_SizeForItemAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, indexPath unsafe.Pointer) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_SizeForItemAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), foundation.MakeIndexPath(indexPath))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_InsetForSectionAtIndex
func collectionViewDelegateFlowLayout_CollectionView_Layout_InsetForSectionAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, section C.int) C.NSEdgeInsets {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_InsetForSectionAtIndex(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), int(section))
	return *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(result))
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumLineSpacingForSectionAtIndex
func collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, section C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), int(section))
	return C.double(float64(result))
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex
func collectionViewDelegateFlowLayout_CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, section C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), int(section))
	return C.double(float64(result))
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForHeaderInSection
func collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForHeaderInSection(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, section C.int) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_ReferenceSizeForHeaderInSection(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), int(section))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForFooterInSection
func collectionViewDelegateFlowLayout_CollectionView_Layout_ReferenceSizeForFooterInSection(hp C.uintptr_t, collectionView unsafe.Pointer, collectionViewLayout unsafe.Pointer, section C.int) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_Layout_ReferenceSizeForFooterInSection(MakeCollectionView(collectionView), MakeCollectionViewLayout(collectionViewLayout), int(section))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export collectionViewDelegateFlowLayout_CollectionView_ShouldSelectItemsAtIndexPaths
func collectionViewDelegateFlowLayout_CollectionView_ShouldSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_ShouldSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_DidSelectItemsAtIndexPaths
func collectionViewDelegateFlowLayout_CollectionView_DidSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DidSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegateFlowLayout_CollectionView_ShouldDeselectItemsAtIndexPaths
func collectionViewDelegateFlowLayout_CollectionView_ShouldDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_ShouldDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_DidDeselectItemsAtIndexPaths
func collectionViewDelegateFlowLayout_CollectionView_DidDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DidDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegateFlowLayout_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegateFlowLayout_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegateFlowLayout_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
}

//export collectionViewDelegateFlowLayout_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegateFlowLayout_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegateFlowLayout_CollectionView_TransitionLayoutForOldLayout_NewLayout
func collectionViewDelegateFlowLayout_CollectionView_TransitionLayoutForOldLayout_NewLayout(hp C.uintptr_t, collectionView unsafe.Pointer, fromLayout unsafe.Pointer, toLayout unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout(MakeCollectionView(collectionView), MakeCollectionViewLayout(fromLayout), MakeCollectionViewLayout(toLayout))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexPaths_WithEvent
func collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexPaths_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndexPath
func collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths
func collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegateFlowLayout_CollectionView_DraggingSession_EndedAtPoint_DragOperation
func collectionViewDelegateFlowLayout_CollectionView_DraggingSession_EndedAtPoint_DragOperation(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export collectionViewDelegateFlowLayout_CollectionView_UpdateDraggingItemsForDrag
func collectionViewDelegateFlowLayout_CollectionView_UpdateDraggingItemsForDrag(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_UpdateDraggingItemsForDrag(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo))
}

//export collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_IndexPath_DropOperation
func collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_IndexPath_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, indexPath unsafe.Pointer, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_AcceptDrop_IndexPath_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), foundation.MakeIndexPath(indexPath), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexes_WithEvent
func collectionViewDelegateFlowLayout_CollectionView_CanDragItemsAtIndexes_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexes unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_CanDragItemsAtIndexes_WithEvent(MakeCollectionView(collectionView), foundation.MakeIndexSet(indexes), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndex
func collectionViewDelegateFlowLayout_CollectionView_PasteboardWriterForItemAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, index C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndex(MakeCollectionView(collectionView), uint(index))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes
func collectionViewDelegateFlowLayout_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(indexes))
}

//export collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_Index_DropOperation
func collectionViewDelegateFlowLayout_CollectionView_AcceptDrop_Index_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, index C.int, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	result := delegate.CollectionView_AcceptDrop_Index_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), int(index), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export CollectionViewDelegateFlowLayout_RespondsTo
func CollectionViewDelegateFlowLayout_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegateFlowLayout)
	switch selName {
	case "collectionView:layout:sizeForItemAtIndexPath:":
		return delegate.CollectionView_Layout_SizeForItemAtIndexPath != nil
	case "collectionView:layout:insetForSectionAtIndex:":
		return delegate.CollectionView_Layout_InsetForSectionAtIndex != nil
	case "collectionView:layout:minimumLineSpacingForSectionAtIndex:":
		return delegate.CollectionView_Layout_MinimumLineSpacingForSectionAtIndex != nil
	case "collectionView:layout:minimumInteritemSpacingForSectionAtIndex:":
		return delegate.CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex != nil
	case "collectionView:layout:referenceSizeForHeaderInSection:":
		return delegate.CollectionView_Layout_ReferenceSizeForHeaderInSection != nil
	case "collectionView:layout:referenceSizeForFooterInSection:":
		return delegate.CollectionView_Layout_ReferenceSizeForFooterInSection != nil
	case "collectionView:shouldSelectItemsAtIndexPaths:":
		return delegate.CollectionView_ShouldSelectItemsAtIndexPaths != nil
	case "collectionView:didSelectItemsAtIndexPaths:":
		return delegate.CollectionView_DidSelectItemsAtIndexPaths != nil
	case "collectionView:shouldDeselectItemsAtIndexPaths:":
		return delegate.CollectionView_ShouldDeselectItemsAtIndexPaths != nil
	case "collectionView:didDeselectItemsAtIndexPaths:":
		return delegate.CollectionView_DidDeselectItemsAtIndexPaths != nil
	case "collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:":
		return delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState != nil
	case "collectionView:didChangeItemsAtIndexPaths:toHighlightState:":
		return delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState != nil
	case "collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:":
		return delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath != nil
	case "collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:":
		return delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath != nil
	case "collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:":
		return delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath != nil
	case "collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:":
		return delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath != nil
	case "collectionView:transitionLayoutForOldLayout:newLayout:":
		return delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout != nil
	case "collectionView:canDragItemsAtIndexPaths:withEvent:":
		return delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent != nil
	case "collectionView:pasteboardWriterForItemAtIndexPath:":
		return delegate.CollectionView_PasteboardWriterForItemAtIndexPath != nil
	case "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:":
		return delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths != nil
	case "collectionView:draggingSession:endedAtPoint:dragOperation:":
		return delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation != nil
	case "collectionView:updateDraggingItemsForDrag:":
		return delegate.CollectionView_UpdateDraggingItemsForDrag != nil
	case "collectionView:acceptDrop:indexPath:dropOperation:":
		return delegate.CollectionView_AcceptDrop_IndexPath_DropOperation != nil
	case "collectionView:canDragItemsAtIndexes:withEvent:":
		return delegate.CollectionView_CanDragItemsAtIndexes_WithEvent != nil
	case "collectionView:pasteboardWriterForItemAtIndex:":
		return delegate.CollectionView_PasteboardWriterForItemAtIndex != nil
	case "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:":
		return delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes != nil
	case "collectionView:acceptDrop:index:dropOperation:":
		return delegate.CollectionView_AcceptDrop_Index_DropOperation != nil
	default:
		return false
	}
}
