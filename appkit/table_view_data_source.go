package appkit

// #include "table_view_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func WrapTableViewDataSource(delegate *TableViewDataSource) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapTableViewDataSource(C.long(id))
	return objc.MakeObject(ptr)
}

//export tableViewDataSource_NumberOfRowsInTableView
func tableViewDataSource_NumberOfRowsInTableView(id int64, tableView unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*TableViewDataSource)
	result := delegate.NumberOfRowsInTableView(MakeTableView(tableView))
	return C.int(result)
}

//export tableViewDataSource_TableView_ObjectValueForTableColumn_Row
func tableViewDataSource_TableView_ObjectValueForTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDataSource)
	result := delegate.TableView_ObjectValueForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row
func tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row(id int64, tableView unsafe.Pointer, object unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*TableViewDataSource)
	delegate.TableView_SetObjectValue_ForTableColumn_Row(MakeTableView(tableView), objc.MakeObject(object), MakeTableColumn(tableColumn), int(row))
}

//export tableViewDataSource_TableView_PasteboardWriterForRow
func tableViewDataSource_TableView_PasteboardWriterForRow(id int64, tableView unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDataSource)
	result := delegate.TableView_PasteboardWriterForRow(MakeTableView(tableView), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_AcceptDrop_Row_DropOperation
func tableViewDataSource_TableView_AcceptDrop_Row_DropOperation(id int64, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.bool {
	delegate := resources.Get(id).(*TableViewDataSource)
	result := delegate.TableView_AcceptDrop_Row_DropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.bool(result)
}

//export tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation
func tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation(id int64, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.uint {
	delegate := resources.Get(id).(*TableViewDataSource)
	result := delegate.TableView_ValidateDrop_ProposedRow_ProposedDropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.uint(uint(result))
}

//export tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes
func tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(id int64, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, rowIndexes unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDataSource)
	delegate.TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(rowIndexes))
}

//export tableViewDataSource_TableView_UpdateDraggingItemsForDrag
func tableViewDataSource_TableView_UpdateDraggingItemsForDrag(id int64, tableView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDataSource)
	delegate.TableView_UpdateDraggingItemsForDrag(MakeTableView(tableView), objc.MakeObject(draggingInfo))
}

//export tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation
func tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation(id int64, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := resources.Get(id).(*TableViewDataSource)
	delegate.TableView_DraggingSession_EndedAtPoint_Operation(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export tableViewDataSource_TableView_SortDescriptorsDidChange
func tableViewDataSource_TableView_SortDescriptorsDidChange(id int64, tableView unsafe.Pointer, oldDescriptors C.Array) {
	delegate := resources.Get(id).(*TableViewDataSource)
	defer C.free(oldDescriptors.data)
	oldDescriptorsSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(oldDescriptors.data))[:oldDescriptors.len:oldDescriptors.len]
	var goOldDescriptors = make([]foundation.SortDescriptor, len(oldDescriptorsSlice))
	for idx, r := range oldDescriptorsSlice {
		goOldDescriptors[idx] = foundation.MakeSortDescriptor(r)
	}
	delegate.TableView_SortDescriptorsDidChange(MakeTableView(tableView), goOldDescriptors)
}

//export TableViewDataSource_RespondsTo
func TableViewDataSource_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*TableViewDataSource)
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

//export deleteTableViewDataSource
func deleteTableViewDataSource(id int64) {
	resources.Delete(id)
}
