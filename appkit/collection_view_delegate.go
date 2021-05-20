package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "collection_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func WrapCollectionViewDelegate(delegate *CollectionViewDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapCollectionViewDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export CollectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths
func CollectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths
func CollectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DidSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export CollectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths
func CollectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths
func CollectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DidDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export CollectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState
func CollectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState
func CollectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
}

//export CollectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath
func CollectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(id int64, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export CollectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath
func CollectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(id int64, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export CollectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath
func CollectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(id int64, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export CollectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath
func CollectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(id int64, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export CollectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout
func CollectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout(id int64, collectionView unsafe.Pointer, fromLayout unsafe.Pointer, toLayout unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout(MakeCollectionView(collectionView), MakeCollectionViewLayout(fromLayout), MakeCollectionViewLayout(toLayout))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent
func CollectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent(id int64, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), MakeEvent(event))
	return C.bool(result)
}

//export CollectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath
func CollectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath(id int64, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths
func CollectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(id int64, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexPaths unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeSet(indexPaths))
}

//export CollectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation
func CollectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation(id int64, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export CollectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag
func CollectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag(id int64, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_UpdateDraggingItemsForDrag(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo))
}

//export CollectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation
func CollectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation(id int64, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, indexPath unsafe.Pointer, dropOperation C.int) C.bool {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_IndexPath_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), foundation.MakeIndexPath(indexPath), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export CollectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent
func CollectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent(id int64, collectionView unsafe.Pointer, indexes unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexes_WithEvent(MakeCollectionView(collectionView), foundation.MakeIndexSet(indexes), MakeEvent(event))
	return C.bool(result)
}

//export CollectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex
func CollectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex(id int64, collectionView unsafe.Pointer, index C.uint) unsafe.Pointer {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndex(MakeCollectionView(collectionView), uint(index))
	return objc.ExtractPtr(result)
}

//export CollectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes
func CollectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(id int64, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexes unsafe.Pointer) {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(indexes))
}

//export CollectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation
func CollectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation(id int64, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, index C.int, dropOperation C.int) C.bool {
	delegate := resources.Get(id).(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_Index_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), int(index), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export CollectionViewDelegate_RespondsTo
func CollectionViewDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*CollectionViewDelegate)
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

//export deleteCollectionViewDelegate
func deleteCollectionViewDelegate(id int64) {
	resources.Delete(id)
}
