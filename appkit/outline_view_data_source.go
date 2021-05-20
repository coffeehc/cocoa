package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "outline_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type OutlineViewDataSource struct {
	OutlineView_AcceptDrop_Item_ChildIndex                   func(outlineView OutlineView, info objc.Object, item objc.Object, index int) bool
	OutlineView_Child_OfItem                                 func(outlineView OutlineView, index int, item objc.Object) objc.Object
	OutlineView_DraggingSession_EndedAtPoint_Operation       func(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	OutlineView_DraggingSession_WillBeginAtPoint_ForItems    func(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	OutlineView_IsItemExpandable                             func(outlineView OutlineView, item objc.Object) bool
	OutlineView_ItemForPersistentObject                      func(outlineView OutlineView, object objc.Object) objc.Object
	OutlineView_NumberOfChildrenOfItem                       func(outlineView OutlineView, item objc.Object) int
	OutlineView_ObjectValueForTableColumn_ByItem             func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.Object
	OutlineView_PasteboardWriterForItem                      func(outlineView OutlineView, item objc.Object) objc.Object
	OutlineView_PersistentObjectForItem                      func(outlineView OutlineView, item objc.Object) objc.Object
	OutlineView_SetObjectValue_ForTableColumn_ByItem         func(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object)
	OutlineView_SortDescriptorsDidChange                     func(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	OutlineView_UpdateDraggingItemsForDrag                   func(outlineView OutlineView, draggingInfo objc.Object)
	OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex func(outlineView OutlineView, info objc.Object, item objc.Object, index int) DragOperation
}

func WrapOutlineViewDataSource(delegate *OutlineViewDataSource) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapOutlineViewDataSource(C.long(id))
	return objc.MakeObject(ptr)
}

//export OutlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex
func OutlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex(id int64, outlineView unsafe.Pointer, info unsafe.Pointer, item unsafe.Pointer, index C.int) C.bool {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_AcceptDrop_Item_ChildIndex(MakeOutlineView(outlineView), objc.MakeObject(info), objc.MakeObject(item), int(index))
	return C.bool(result)
}

//export OutlineViewDataSource_OutlineView_Child_OfItem
func OutlineViewDataSource_OutlineView_Child_OfItem(id int64, outlineView unsafe.Pointer, index C.int, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_Child_OfItem(MakeOutlineView(outlineView), int(index), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export OutlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation
func OutlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation(id int64, outlineView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	delegate.OutlineView_DraggingSession_EndedAtPoint_Operation(MakeOutlineView(outlineView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export OutlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems
func OutlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems(id int64, outlineView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, draggedItems C.Array) {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	defer C.free(draggedItems.data)
	draggedItemsSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(draggedItems.data))[:draggedItems.len:draggedItems.len]
	var goDraggedItems = make([]objc.Object, len(draggedItemsSlice))
	for idx, r := range draggedItemsSlice {
		goDraggedItems[idx] = objc.MakeObject(r)
	}
	delegate.OutlineView_DraggingSession_WillBeginAtPoint_ForItems(MakeOutlineView(outlineView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), goDraggedItems)
}

//export OutlineViewDataSource_OutlineView_IsItemExpandable
func OutlineViewDataSource_OutlineView_IsItemExpandable(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_IsItemExpandable(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export OutlineViewDataSource_OutlineView_ItemForPersistentObject
func OutlineViewDataSource_OutlineView_ItemForPersistentObject(id int64, outlineView unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_ItemForPersistentObject(MakeOutlineView(outlineView), objc.MakeObject(object))
	return objc.ExtractPtr(result)
}

//export OutlineViewDataSource_OutlineView_NumberOfChildrenOfItem
func OutlineViewDataSource_OutlineView_NumberOfChildrenOfItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_NumberOfChildrenOfItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.int(result)
}

//export OutlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem
func OutlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_ObjectValueForTableColumn_ByItem(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export OutlineViewDataSource_OutlineView_PasteboardWriterForItem
func OutlineViewDataSource_OutlineView_PasteboardWriterForItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_PasteboardWriterForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export OutlineViewDataSource_OutlineView_PersistentObjectForItem
func OutlineViewDataSource_OutlineView_PersistentObjectForItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_PersistentObjectForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export OutlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem
func OutlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem(id int64, outlineView unsafe.Pointer, object unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	delegate.OutlineView_SetObjectValue_ForTableColumn_ByItem(MakeOutlineView(outlineView), objc.MakeObject(object), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export OutlineViewDataSource_OutlineView_SortDescriptorsDidChange
func OutlineViewDataSource_OutlineView_SortDescriptorsDidChange(id int64, outlineView unsafe.Pointer, oldDescriptors C.Array) {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	defer C.free(oldDescriptors.data)
	oldDescriptorsSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(oldDescriptors.data))[:oldDescriptors.len:oldDescriptors.len]
	var goOldDescriptors = make([]foundation.SortDescriptor, len(oldDescriptorsSlice))
	for idx, r := range oldDescriptorsSlice {
		goOldDescriptors[idx] = foundation.MakeSortDescriptor(r)
	}
	delegate.OutlineView_SortDescriptorsDidChange(MakeOutlineView(outlineView), goOldDescriptors)
}

//export OutlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag
func OutlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag(id int64, outlineView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	delegate.OutlineView_UpdateDraggingItemsForDrag(MakeOutlineView(outlineView), objc.MakeObject(draggingInfo))
}

//export OutlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex
func OutlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(id int64, outlineView unsafe.Pointer, info unsafe.Pointer, item unsafe.Pointer, index C.int) C.uint {
	delegate := resources.Get(id).(*OutlineViewDataSource)
	result := delegate.OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(MakeOutlineView(outlineView), objc.MakeObject(info), objc.MakeObject(item), int(index))
	return C.uint(uint(result))
}

//export OutlineViewDataSource_RespondsTo
func OutlineViewDataSource_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*OutlineViewDataSource)
	switch selName {
	case "outlineView:acceptDrop:item:childIndex:":
		return delegate.OutlineView_AcceptDrop_Item_ChildIndex != nil
	case "outlineView:child:ofItem:":
		return delegate.OutlineView_Child_OfItem != nil
	case "outlineView:draggingSession:endedAtPoint:operation:":
		return delegate.OutlineView_DraggingSession_EndedAtPoint_Operation != nil
	case "outlineView:draggingSession:willBeginAtPoint:forItems:":
		return delegate.OutlineView_DraggingSession_WillBeginAtPoint_ForItems != nil
	case "outlineView:isItemExpandable:":
		return delegate.OutlineView_IsItemExpandable != nil
	case "outlineView:itemForPersistentObject:":
		return delegate.OutlineView_ItemForPersistentObject != nil
	case "outlineView:numberOfChildrenOfItem:":
		return delegate.OutlineView_NumberOfChildrenOfItem != nil
	case "outlineView:objectValueForTableColumn:byItem:":
		return delegate.OutlineView_ObjectValueForTableColumn_ByItem != nil
	case "outlineView:pasteboardWriterForItem:":
		return delegate.OutlineView_PasteboardWriterForItem != nil
	case "outlineView:persistentObjectForItem:":
		return delegate.OutlineView_PersistentObjectForItem != nil
	case "outlineView:setObjectValue:forTableColumn:byItem:":
		return delegate.OutlineView_SetObjectValue_ForTableColumn_ByItem != nil
	case "outlineView:sortDescriptorsDidChange:":
		return delegate.OutlineView_SortDescriptorsDidChange != nil
	case "outlineView:updateDraggingItemsForDrag:":
		return delegate.OutlineView_UpdateDraggingItemsForDrag != nil
	case "outlineView:validateDrop:proposedItem:proposedChildIndex:":
		return delegate.OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex != nil
	default:
		return false
	}
}

//export deleteOutlineViewDataSource
func deleteOutlineViewDataSource(id int64) {
	resources.Delete(id)
}
