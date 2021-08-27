package appkit

// #include "outline_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
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

func (delegate *OutlineViewDataSource) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapOutlineViewDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export outlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex
func outlineViewDataSource_OutlineView_AcceptDrop_Item_ChildIndex(hp C.uintptr_t, outlineView unsafe.Pointer, info unsafe.Pointer, item unsafe.Pointer, index C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_AcceptDrop_Item_ChildIndex(MakeOutlineView(outlineView), objc.MakeObject(info), objc.MakeObject(item), int(index))
	return C.bool(result)
}

//export outlineViewDataSource_OutlineView_Child_OfItem
func outlineViewDataSource_OutlineView_Child_OfItem(hp C.uintptr_t, outlineView unsafe.Pointer, index C.int, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_Child_OfItem(MakeOutlineView(outlineView), int(index), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation
func outlineViewDataSource_OutlineView_DraggingSession_EndedAtPoint_Operation(hp C.uintptr_t, outlineView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	delegate.OutlineView_DraggingSession_EndedAtPoint_Operation(MakeOutlineView(outlineView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export outlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems
func outlineViewDataSource_OutlineView_DraggingSession_WillBeginAtPoint_ForItems(hp C.uintptr_t, outlineView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, draggedItems C.Array) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	if draggedItems.len > 0 {
		defer C.free(draggedItems.data)
	}
	draggedItemsSlice := unsafe.Slice((*unsafe.Pointer)(draggedItems.data), int(draggedItems.len))
	var goDraggedItems = make([]objc.Object, len(draggedItemsSlice))
	for idx, r := range draggedItemsSlice {
		goDraggedItems[idx] = objc.MakeObject(r)
	}
	delegate.OutlineView_DraggingSession_WillBeginAtPoint_ForItems(MakeOutlineView(outlineView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), goDraggedItems)
}

//export outlineViewDataSource_OutlineView_IsItemExpandable
func outlineViewDataSource_OutlineView_IsItemExpandable(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_IsItemExpandable(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDataSource_OutlineView_ItemForPersistentObject
func outlineViewDataSource_OutlineView_ItemForPersistentObject(hp C.uintptr_t, outlineView unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_ItemForPersistentObject(MakeOutlineView(outlineView), objc.MakeObject(object))
	return objc.ExtractPtr(result)
}

//export outlineViewDataSource_OutlineView_NumberOfChildrenOfItem
func outlineViewDataSource_OutlineView_NumberOfChildrenOfItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_NumberOfChildrenOfItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.int(result)
}

//export outlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem
func outlineViewDataSource_OutlineView_ObjectValueForTableColumn_ByItem(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_ObjectValueForTableColumn_ByItem(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDataSource_OutlineView_PasteboardWriterForItem
func outlineViewDataSource_OutlineView_PasteboardWriterForItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_PasteboardWriterForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDataSource_OutlineView_PersistentObjectForItem
func outlineViewDataSource_OutlineView_PersistentObjectForItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_PersistentObjectForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem
func outlineViewDataSource_OutlineView_SetObjectValue_ForTableColumn_ByItem(hp C.uintptr_t, outlineView unsafe.Pointer, object unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	delegate.OutlineView_SetObjectValue_ForTableColumn_ByItem(MakeOutlineView(outlineView), objc.MakeObject(object), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export outlineViewDataSource_OutlineView_SortDescriptorsDidChange
func outlineViewDataSource_OutlineView_SortDescriptorsDidChange(hp C.uintptr_t, outlineView unsafe.Pointer, oldDescriptors C.Array) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	if oldDescriptors.len > 0 {
		defer C.free(oldDescriptors.data)
	}
	oldDescriptorsSlice := unsafe.Slice((*unsafe.Pointer)(oldDescriptors.data), int(oldDescriptors.len))
	var goOldDescriptors = make([]foundation.SortDescriptor, len(oldDescriptorsSlice))
	for idx, r := range oldDescriptorsSlice {
		goOldDescriptors[idx] = foundation.MakeSortDescriptor(r)
	}
	delegate.OutlineView_SortDescriptorsDidChange(MakeOutlineView(outlineView), goOldDescriptors)
}

//export outlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag
func outlineViewDataSource_OutlineView_UpdateDraggingItemsForDrag(hp C.uintptr_t, outlineView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	delegate.OutlineView_UpdateDraggingItemsForDrag(MakeOutlineView(outlineView), objc.MakeObject(draggingInfo))
}

//export outlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex
func outlineViewDataSource_OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(hp C.uintptr_t, outlineView unsafe.Pointer, info unsafe.Pointer, item unsafe.Pointer, index C.int) C.uint {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
	result := delegate.OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(MakeOutlineView(outlineView), objc.MakeObject(info), objc.MakeObject(item), int(index))
	return C.uint(uint(result))
}

//export OutlineViewDataSource_RespondsTo
func OutlineViewDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*OutlineViewDataSource)
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
