package appkit

// #include "table_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TableViewDataSource struct {
	NumberOfRowsInTableView                                  func(tableView TableView) int
	TableView_ObjectValueForTableColumn_Row                  func(tableView TableView, tableColumn TableColumn, row int) objc.Object
	TableView_SetObjectValue_ForTableColumn_Row              func(tableView TableView, object objc.Object, tableColumn TableColumn, row int)
	TableView_PasteboardWriterForRow                         func(tableView TableView, row int) objc.Object
	TableView_AcceptDrop_Row_DropOperation                   func(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) bool
	TableView_ValidateDrop_ProposedRow_ProposedDropOperation func(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) DragOperation
	TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes func(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	TableView_UpdateDraggingItemsForDrag                     func(tableView TableView, draggingInfo objc.Object)
	TableView_DraggingSession_EndedAtPoint_Operation         func(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	TableView_SortDescriptorsDidChange                       func(tableView TableView, oldDescriptors []foundation.SortDescriptor)
}

func (delegate *TableViewDataSource) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTableViewDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export tableViewDataSource_NumberOfRowsInTableView
func tableViewDataSource_NumberOfRowsInTableView(hp C.uintptr_t, tableView unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	result := delegate.NumberOfRowsInTableView(MakeTableView(tableView))
	return C.int(result)
}

//export tableViewDataSource_TableView_ObjectValueForTableColumn_Row
func tableViewDataSource_TableView_ObjectValueForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	result := delegate.TableView_ObjectValueForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row
func tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, object unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	delegate.TableView_SetObjectValue_ForTableColumn_Row(MakeTableView(tableView), objc.MakeObject(object), MakeTableColumn(tableColumn), int(row))
}

//export tableViewDataSource_TableView_PasteboardWriterForRow
func tableViewDataSource_TableView_PasteboardWriterForRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	result := delegate.TableView_PasteboardWriterForRow(MakeTableView(tableView), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_AcceptDrop_Row_DropOperation
func tableViewDataSource_TableView_AcceptDrop_Row_DropOperation(hp C.uintptr_t, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	result := delegate.TableView_AcceptDrop_Row_DropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.bool(result)
}

//export tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation
func tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation(hp C.uintptr_t, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.uint {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	result := delegate.TableView_ValidateDrop_ProposedRow_ProposedDropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.uint(uint(result))
}

//export tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes
func tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(hp C.uintptr_t, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, rowIndexes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	delegate.TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(rowIndexes))
}

//export tableViewDataSource_TableView_UpdateDraggingItemsForDrag
func tableViewDataSource_TableView_UpdateDraggingItemsForDrag(hp C.uintptr_t, tableView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	delegate.TableView_UpdateDraggingItemsForDrag(MakeTableView(tableView), objc.MakeObject(draggingInfo))
}

//export tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation
func tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation(hp C.uintptr_t, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	delegate.TableView_DraggingSession_EndedAtPoint_Operation(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export tableViewDataSource_TableView_SortDescriptorsDidChange
func tableViewDataSource_TableView_SortDescriptorsDidChange(hp C.uintptr_t, tableView unsafe.Pointer, oldDescriptors C.Array) {
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	if oldDescriptors.len > 0 {
		defer C.free(oldDescriptors.data)
	}
	oldDescriptorsSlice := unsafe.Slice((*unsafe.Pointer)(oldDescriptors.data), int(oldDescriptors.len))
	var goOldDescriptors = make([]foundation.SortDescriptor, len(oldDescriptorsSlice))
	for idx, r := range oldDescriptorsSlice {
		goOldDescriptors[idx] = foundation.MakeSortDescriptor(r)
	}
	delegate.TableView_SortDescriptorsDidChange(MakeTableView(tableView), goOldDescriptors)
}

//export TableViewDataSource_RespondsTo
func TableViewDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TableViewDataSource)
	switch selName {
	case "numberOfRowsInTableView:":
		return delegate.NumberOfRowsInTableView != nil
	case "tableView:objectValueForTableColumn:row:":
		return delegate.TableView_ObjectValueForTableColumn_Row != nil
	case "tableView:setObjectValue:forTableColumn:row:":
		return delegate.TableView_SetObjectValue_ForTableColumn_Row != nil
	case "tableView:pasteboardWriterForRow:":
		return delegate.TableView_PasteboardWriterForRow != nil
	case "tableView:acceptDrop:row:dropOperation:":
		return delegate.TableView_AcceptDrop_Row_DropOperation != nil
	case "tableView:validateDrop:proposedRow:proposedDropOperation:":
		return delegate.TableView_ValidateDrop_ProposedRow_ProposedDropOperation != nil
	case "tableView:draggingSession:willBeginAtPoint:forRowIndexes:":
		return delegate.TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes != nil
	case "tableView:updateDraggingItemsForDrag:":
		return delegate.TableView_UpdateDraggingItemsForDrag != nil
	case "tableView:draggingSession:endedAtPoint:operation:":
		return delegate.TableView_DraggingSession_EndedAtPoint_Operation != nil
	case "tableView:sortDescriptorsDidChange:":
		return delegate.TableView_SortDescriptorsDidChange != nil
	default:
		return false
	}
}
