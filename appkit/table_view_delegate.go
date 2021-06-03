package appkit

// #include "table_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableViewDelegate struct {
	TableView_ViewForTableColumn_Row                           func(tableView TableView, tableColumn TableColumn, row int) View
	TableView_RowViewForRow                                    func(tableView TableView, row int) TableRowView
	TableView_DidAddRowView_ForRow                             func(tableView TableView, rowView TableRowView, row int)
	TableView_DidRemoveRowView_ForRow                          func(tableView TableView, rowView TableRowView, row int)
	TableView_IsGroupRow                                       func(tableView TableView, row int) bool
	TableView_WillDisplayCell_ForTableColumn_Row               func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	TableView_DataCellForTableColumn_Row                       func(tableView TableView, tableColumn TableColumn, row int) Cell
	TableView_ShouldShowCellExpansionForTableColumn_Row        func(tableView TableView, tableColumn TableColumn, row int) bool
	TableView_ShouldEditTableColumn_Row                        func(tableView TableView, tableColumn TableColumn, row int) bool
	TableView_HeightOfRow                                      func(tableView TableView, row int) coregraphics.Float
	TableView_SizeToFitWidthOfColumn                           func(tableView TableView, column int) coregraphics.Float
	SelectionShouldChangeInTableView                           func(tableView TableView) bool
	TableView_ShouldSelectRow                                  func(tableView TableView, row int) bool
	TableView_SelectionIndexesForProposedSelection             func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	TableView_ShouldSelectTableColumn                          func(tableView TableView, tableColumn TableColumn) bool
	TableViewSelectionIsChanging                               func(notification foundation.Notification)
	TableViewSelectionDidChange                                func(notification foundation.Notification)
	TableView_ShouldTypeSelectForEvent_WithCurrentSearchString func(tableView TableView, event Event, searchString string) bool
	TableView_TypeSelectStringForTableColumn_Row               func(tableView TableView, tableColumn TableColumn, row int) string
	TableView_NextTypeSelectMatchFromRow_ToRow_ForString       func(tableView TableView, startRow int, endRow int, searchString string) int
	TableView_ShouldReorderColumn_ToColumn                     func(tableView TableView, columnIndex int, newColumnIndex int) bool
	TableView_DidDragTableColumn                               func(tableView TableView, tableColumn TableColumn)
	TableViewColumnDidMove                                     func(notification foundation.Notification)
	TableViewColumnDidResize                                   func(notification foundation.Notification)
	TableView_DidClickTableColumn                              func(tableView TableView, tableColumn TableColumn)
	TableView_MouseDownInHeaderOfTableColumn                   func(tableView TableView, tableColumn TableColumn)
	TableView_ShouldTrackCell_ForTableColumn_Row               func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	TableView_RowActionsForRow_Edge                            func(tableView TableView, row int, edge TableRowActionEdge) []TableViewRowAction
	Control_IsValidObject                                      func(control Control, obj objc.Object) bool
	Control_DidFailToValidatePartialString_ErrorDescription    func(control Control, _string string, error string)
	Control_DidFailToFormatString_ErrorDescription             func(control Control, _string string, error string) bool
	Control_TextShouldBeginEditing                             func(control Control, fieldEditor Text) bool
	Control_TextShouldEndEditing                               func(control Control, fieldEditor Text) bool
	Control_TextView_DoCommandBySelector                       func(control Control, textView TextView, commandSelector objc.Selector) bool
	ControlTextDidBeginEditing                                 func(obj foundation.Notification)
	ControlTextDidChange                                       func(obj foundation.Notification)
	ControlTextDidEndEditing                                   func(obj foundation.Notification)
}

func WrapTableViewDelegate(delegate *TableViewDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapTableViewDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export tableViewDelegate_TableView_ViewForTableColumn_Row
func tableViewDelegate_TableView_ViewForTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ViewForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_RowViewForRow
func tableViewDelegate_TableView_RowViewForRow(id int64, tableView unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_RowViewForRow(MakeTableView(tableView), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_DidAddRowView_ForRow
func tableViewDelegate_TableView_DidAddRowView_ForRow(id int64, tableView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_DidAddRowView_ForRow(MakeTableView(tableView), MakeTableRowView(rowView), int(row))
}

//export tableViewDelegate_TableView_DidRemoveRowView_ForRow
func tableViewDelegate_TableView_DidRemoveRowView_ForRow(id int64, tableView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_DidRemoveRowView_ForRow(MakeTableView(tableView), MakeTableRowView(rowView), int(row))
}

//export tableViewDelegate_TableView_IsGroupRow
func tableViewDelegate_TableView_IsGroupRow(id int64, tableView unsafe.Pointer, row C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_IsGroupRow(MakeTableView(tableView), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_WillDisplayCell_ForTableColumn_Row
func tableViewDelegate_TableView_WillDisplayCell_ForTableColumn_Row(id int64, tableView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_WillDisplayCell_ForTableColumn_Row(MakeTableView(tableView), objc.MakeObject(cell), MakeTableColumn(tableColumn), int(row))
}

//export tableViewDelegate_TableView_DataCellForTableColumn_Row
func tableViewDelegate_TableView_DataCellForTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_DataCellForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_ShouldShowCellExpansionForTableColumn_Row
func tableViewDelegate_TableView_ShouldShowCellExpansionForTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldShowCellExpansionForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_ShouldEditTableColumn_Row
func tableViewDelegate_TableView_ShouldEditTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldEditTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_HeightOfRow
func tableViewDelegate_TableView_HeightOfRow(id int64, tableView unsafe.Pointer, row C.int) C.double {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_HeightOfRow(MakeTableView(tableView), int(row))
	return C.double(float64(result))
}

//export tableViewDelegate_TableView_SizeToFitWidthOfColumn
func tableViewDelegate_TableView_SizeToFitWidthOfColumn(id int64, tableView unsafe.Pointer, column C.int) C.double {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_SizeToFitWidthOfColumn(MakeTableView(tableView), int(column))
	return C.double(float64(result))
}

//export tableViewDelegate_SelectionShouldChangeInTableView
func tableViewDelegate_SelectionShouldChangeInTableView(id int64, tableView unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.SelectionShouldChangeInTableView(MakeTableView(tableView))
	return C.bool(result)
}

//export tableViewDelegate_TableView_ShouldSelectRow
func tableViewDelegate_TableView_ShouldSelectRow(id int64, tableView unsafe.Pointer, row C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldSelectRow(MakeTableView(tableView), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_SelectionIndexesForProposedSelection
func tableViewDelegate_TableView_SelectionIndexesForProposedSelection(id int64, tableView unsafe.Pointer, proposedSelectionIndexes unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_SelectionIndexesForProposedSelection(MakeTableView(tableView), foundation.MakeIndexSet(proposedSelectionIndexes))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_ShouldSelectTableColumn
func tableViewDelegate_TableView_ShouldSelectTableColumn(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldSelectTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
	return C.bool(result)
}

//export tableViewDelegate_TableViewSelectionIsChanging
func tableViewDelegate_TableViewSelectionIsChanging(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableViewSelectionIsChanging(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableViewSelectionDidChange
func tableViewDelegate_TableViewSelectionDidChange(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableViewSelectionDidChange(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString
func tableViewDelegate_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(id int64, tableView unsafe.Pointer, event unsafe.Pointer, searchString unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(MakeTableView(tableView), MakeEvent(event), foundation.MakeString(searchString).String())
	return C.bool(result)
}

//export tableViewDelegate_TableView_TypeSelectStringForTableColumn_Row
func tableViewDelegate_TableView_TypeSelectStringForTableColumn_Row(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_TypeSelectStringForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return foundation.NewString(result).Ptr()
}

//export tableViewDelegate_TableView_NextTypeSelectMatchFromRow_ToRow_ForString
func tableViewDelegate_TableView_NextTypeSelectMatchFromRow_ToRow_ForString(id int64, tableView unsafe.Pointer, startRow C.int, endRow C.int, searchString unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_NextTypeSelectMatchFromRow_ToRow_ForString(MakeTableView(tableView), int(startRow), int(endRow), foundation.MakeString(searchString).String())
	return C.int(result)
}

//export tableViewDelegate_TableView_ShouldReorderColumn_ToColumn
func tableViewDelegate_TableView_ShouldReorderColumn_ToColumn(id int64, tableView unsafe.Pointer, columnIndex C.int, newColumnIndex C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldReorderColumn_ToColumn(MakeTableView(tableView), int(columnIndex), int(newColumnIndex))
	return C.bool(result)
}

//export tableViewDelegate_TableView_DidDragTableColumn
func tableViewDelegate_TableView_DidDragTableColumn(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_DidDragTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableViewColumnDidMove
func tableViewDelegate_TableViewColumnDidMove(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableViewColumnDidMove(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableViewColumnDidResize
func tableViewDelegate_TableViewColumnDidResize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableViewColumnDidResize(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableView_DidClickTableColumn
func tableViewDelegate_TableView_DidClickTableColumn(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_DidClickTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableView_MouseDownInHeaderOfTableColumn
func tableViewDelegate_TableView_MouseDownInHeaderOfTableColumn(id int64, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.TableView_MouseDownInHeaderOfTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableView_ShouldTrackCell_ForTableColumn_Row
func tableViewDelegate_TableView_ShouldTrackCell_ForTableColumn_Row(id int64, tableView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_ShouldTrackCell_ForTableColumn_Row(MakeTableView(tableView), MakeCell(cell), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_RowActionsForRow_Edge
func tableViewDelegate_TableView_RowActionsForRow_Edge(id int64, tableView unsafe.Pointer, row C.int, edge C.int) C.Array {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.TableView_RowActionsForRow_Edge(MakeTableView(tableView), int(row), TableRowActionEdge(int(edge)))
	cResultData := make([]unsafe.Pointer, len(result))
	for idx, v := range result {
		cResultData[idx] = objc.ExtractPtr(v)
	}
	cResult := C.Array{data: unsafe.Pointer(&cResultData[0]), len: C.int(len(result))}
	return cResult
}

//export tableViewDelegate_Control_IsValidObject
func tableViewDelegate_Control_IsValidObject(id int64, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export tableViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func tableViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export tableViewDelegate_Control_DidFailToFormatString_ErrorDescription
func tableViewDelegate_Control_DidFailToFormatString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export tableViewDelegate_Control_TextShouldBeginEditing
func tableViewDelegate_Control_TextShouldBeginEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export tableViewDelegate_Control_TextShouldEndEditing
func tableViewDelegate_Control_TextShouldEndEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export tableViewDelegate_Control_TextView_DoCommandBySelector
func tableViewDelegate_Control_TextView_DoCommandBySelector(id int64, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TableViewDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export tableViewDelegate_ControlTextDidBeginEditing
func tableViewDelegate_ControlTextDidBeginEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export tableViewDelegate_ControlTextDidChange
func tableViewDelegate_ControlTextDidChange(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export tableViewDelegate_ControlTextDidEndEditing
func tableViewDelegate_ControlTextDidEndEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*TableViewDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export TableViewDelegate_RespondsTo
func TableViewDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*TableViewDelegate)
	switch selName {
	case "tableView:viewForTableColumn:row:":
		return delegate.TableView_ViewForTableColumn_Row != nil
	case "tableView:rowViewForRow:":
		return delegate.TableView_RowViewForRow != nil
	case "tableView:didAddRowView:forRow:":
		return delegate.TableView_DidAddRowView_ForRow != nil
	case "tableView:didRemoveRowView:forRow:":
		return delegate.TableView_DidRemoveRowView_ForRow != nil
	case "tableView:isGroupRow:":
		return delegate.TableView_IsGroupRow != nil
	case "tableView:willDisplayCell:forTableColumn:row:":
		return delegate.TableView_WillDisplayCell_ForTableColumn_Row != nil
	case "tableView:dataCellForTableColumn:row:":
		return delegate.TableView_DataCellForTableColumn_Row != nil
	case "tableView:shouldShowCellExpansionForTableColumn:row:":
		return delegate.TableView_ShouldShowCellExpansionForTableColumn_Row != nil
	case "tableView:shouldEditTableColumn:row:":
		return delegate.TableView_ShouldEditTableColumn_Row != nil
	case "tableView:heightOfRow:":
		return delegate.TableView_HeightOfRow != nil
	case "tableView:sizeToFitWidthOfColumn:":
		return delegate.TableView_SizeToFitWidthOfColumn != nil
	case "selectionShouldChangeInTableView:":
		return delegate.SelectionShouldChangeInTableView != nil
	case "tableView:shouldSelectRow:":
		return delegate.TableView_ShouldSelectRow != nil
	case "tableView:selectionIndexesForProposedSelection:":
		return delegate.TableView_SelectionIndexesForProposedSelection != nil
	case "tableView:shouldSelectTableColumn:":
		return delegate.TableView_ShouldSelectTableColumn != nil
	case "tableViewSelectionIsChanging:":
		return delegate.TableViewSelectionIsChanging != nil
	case "tableViewSelectionDidChange:":
		return delegate.TableViewSelectionDidChange != nil
	case "tableView:shouldTypeSelectForEvent:withCurrentSearchString:":
		return delegate.TableView_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
	case "tableView:typeSelectStringForTableColumn:row:":
		return delegate.TableView_TypeSelectStringForTableColumn_Row != nil
	case "tableView:nextTypeSelectMatchFromRow:toRow:forString:":
		return delegate.TableView_NextTypeSelectMatchFromRow_ToRow_ForString != nil
	case "tableView:shouldReorderColumn:toColumn:":
		return delegate.TableView_ShouldReorderColumn_ToColumn != nil
	case "tableView:didDragTableColumn:":
		return delegate.TableView_DidDragTableColumn != nil
	case "tableViewColumnDidMove:":
		return delegate.TableViewColumnDidMove != nil
	case "tableViewColumnDidResize:":
		return delegate.TableViewColumnDidResize != nil
	case "tableView:didClickTableColumn:":
		return delegate.TableView_DidClickTableColumn != nil
	case "tableView:mouseDownInHeaderOfTableColumn:":
		return delegate.TableView_MouseDownInHeaderOfTableColumn != nil
	case "tableView:shouldTrackCell:forTableColumn:row:":
		return delegate.TableView_ShouldTrackCell_ForTableColumn_Row != nil
	case "tableView:rowActionsForRow:edge:":
		return delegate.TableView_RowActionsForRow_Edge != nil
	case "control:isValidObject:":
		return delegate.Control_IsValidObject != nil
	case "control:didFailToValidatePartialString:errorDescription:":
		return delegate.Control_DidFailToValidatePartialString_ErrorDescription != nil
	case "control:didFailToFormatString:errorDescription:":
		return delegate.Control_DidFailToFormatString_ErrorDescription != nil
	case "control:textShouldBeginEditing:":
		return delegate.Control_TextShouldBeginEditing != nil
	case "control:textShouldEndEditing:":
		return delegate.Control_TextShouldEndEditing != nil
	case "control:textView:doCommandBySelector:":
		return delegate.Control_TextView_DoCommandBySelector != nil
	case "controlTextDidBeginEditing:":
		return delegate.ControlTextDidBeginEditing != nil
	case "controlTextDidChange:":
		return delegate.ControlTextDidChange != nil
	case "controlTextDidEndEditing:":
		return delegate.ControlTextDidEndEditing != nil
	default:
		return false
	}
}

//export deleteTableViewDelegate
func deleteTableViewDelegate(id int64) {
	resources.Delete(id)
}
