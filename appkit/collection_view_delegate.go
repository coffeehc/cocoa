package appkit

// #include "collection_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type CollectionViewDelegate struct {
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

func (delegate *CollectionViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
}

//export collectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath
func collectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath
func collectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout
func collectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout(hp C.uintptr_t, collectionView unsafe.Pointer, fromLayout unsafe.Pointer, toLayout unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout(MakeCollectionView(collectionView), MakeCollectionViewLayout(fromLayout), MakeCollectionViewLayout(toLayout))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent
func collectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath
func collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation
func collectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export collectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag
func collectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_UpdateDraggingItemsForDrag(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo))
}

//export collectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation
func collectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, indexPath unsafe.Pointer, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_IndexPath_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), foundation.MakeIndexPath(indexPath), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent
func collectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexes unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexes_WithEvent(MakeCollectionView(collectionView), foundation.MakeIndexSet(indexes), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex
func collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, index C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndex(MakeCollectionView(collectionView), uint(index))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes
func collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(indexes))
}

//export collectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation
func collectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, index C.int, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_Index_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), int(index), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export CollectionViewDelegate_RespondsTo
func CollectionViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	switch selName {
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
