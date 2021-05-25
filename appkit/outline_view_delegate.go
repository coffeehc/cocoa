package appkit

// #include "outline_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type OutlineViewDelegate struct {
	OutlineView_ShouldExpandItem                                 func(outlineView OutlineView, item objc.Object) bool
	OutlineView_ShouldCollapseItem                               func(outlineView OutlineView, item objc.Object) bool
	OutlineView_TypeSelectStringForTableColumn_Item              func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString     func(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.Object
	OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString func(outlineView OutlineView, event Event, searchString string) bool
	OutlineView_ShouldSelectTableColumn                          func(outlineView OutlineView, tableColumn TableColumn) bool
	OutlineView_ShouldSelectItem                                 func(outlineView OutlineView, item objc.Object) bool
	OutlineView_SelectionIndexesForProposedSelection             func(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	SelectionShouldChangeInOutlineView                           func(outlineView OutlineView) bool
	OutlineViewSelectionIsChanging                               func(notification foundation.Notification)
	OutlineViewSelectionDidChange                                func(notification foundation.Notification)
	OutlineView_WillDisplayCell_ForTableColumn_Item              func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	OutlineView_WillDisplayOutlineCell_ForTableColumn_Item       func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	OutlineView_DataCellForTableColumn_Item                      func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) Cell
	OutlineView_ShouldShowOutlineCellForItem                     func(outlineView OutlineView, item objc.Object) bool
	OutlineView_ShouldShowCellExpansionForTableColumn_Item       func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	OutlineView_ShouldReorderColumn_ToColumn                     func(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	OutlineViewColumnDidMove                                     func(notification foundation.Notification)
	OutlineViewColumnDidResize                                   func(notification foundation.Notification)
	OutlineViewItemWillExpand                                    func(notification foundation.Notification)
	OutlineViewItemDidExpand                                     func(notification foundation.Notification)
	OutlineViewItemWillCollapse                                  func(notification foundation.Notification)
	OutlineViewItemDidCollapse                                   func(notification foundation.Notification)
	OutlineView_ShouldEditTableColumn_Item                       func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	OutlineView_MouseDownInHeaderOfTableColumn                   func(outlineView OutlineView, tableColumn TableColumn)
	OutlineView_DidClickTableColumn                              func(outlineView OutlineView, tableColumn TableColumn)
	OutlineView_DidDragTableColumn                               func(outlineView OutlineView, tableColumn TableColumn)
	OutlineView_HeightOfRowByItem                                func(outlineView OutlineView, item objc.Object) coregraphics.Float
	OutlineView_SizeToFitWidthOfColumn                           func(outlineView OutlineView, column int) coregraphics.Float
	OutlineView_TintConfigurationForItem                         func(outlineView OutlineView, item objc.Object) TintConfiguration
	OutlineView_ShouldTrackCell_ForTableColumn_Item              func(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	OutlineView_IsGroupItem                                      func(outlineView OutlineView, item objc.Object) bool
	OutlineView_DidAddRowView_ForRow                             func(outlineView OutlineView, rowView TableRowView, row int)
	OutlineView_DidRemoveRowView_ForRow                          func(outlineView OutlineView, rowView TableRowView, row int)
	OutlineView_RowViewForItem                                   func(outlineView OutlineView, item objc.Object) TableRowView
	OutlineView_ViewForTableColumn_Item                          func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) View
	Control_IsValidObject                                        func(control Control, obj objc.Object) bool
	Control_DidFailToValidatePartialString_ErrorDescription      func(control Control, _string string, error string)
	Control_DidFailToFormatString_ErrorDescription               func(control Control, _string string, error string) bool
	Control_TextShouldBeginEditing                               func(control Control, fieldEditor Text) bool
	Control_TextShouldEndEditing                                 func(control Control, fieldEditor Text) bool
	Control_TextView_DoCommandBySelector                         func(control Control, textView TextView, commandSelector *objc.Selector) bool
	ControlTextDidBeginEditing                                   func(obj foundation.Notification)
	ControlTextDidChange                                         func(obj foundation.Notification)
	ControlTextDidEndEditing                                     func(obj foundation.Notification)
}

func WrapOutlineViewDelegate(delegate *OutlineViewDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapOutlineViewDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export outlineViewDelegate_OutlineView_ShouldExpandItem
func outlineViewDelegate_OutlineView_ShouldExpandItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldExpandItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldCollapseItem
func outlineViewDelegate_OutlineView_ShouldCollapseItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldCollapseItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_TypeSelectStringForTableColumn_Item
func outlineViewDelegate_OutlineView_TypeSelectStringForTableColumn_Item(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_TypeSelectStringForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return foundation.NewString(result).Ptr()
}

//export outlineViewDelegate_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString
func outlineViewDelegate_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(id int64, outlineView unsafe.Pointer, startItem unsafe.Pointer, endItem unsafe.Pointer, searchString unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(MakeOutlineView(outlineView), objc.MakeObject(startItem), objc.MakeObject(endItem), foundation.MakeString(searchString).String())
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString
func outlineViewDelegate_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(id int64, outlineView unsafe.Pointer, event unsafe.Pointer, searchString unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(MakeOutlineView(outlineView), MakeEvent(event), foundation.MakeString(searchString).String())
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldSelectTableColumn
func outlineViewDelegate_OutlineView_ShouldSelectTableColumn(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldSelectTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldSelectItem
func outlineViewDelegate_OutlineView_ShouldSelectItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldSelectItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_SelectionIndexesForProposedSelection
func outlineViewDelegate_OutlineView_SelectionIndexesForProposedSelection(id int64, outlineView unsafe.Pointer, proposedSelectionIndexes unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_SelectionIndexesForProposedSelection(MakeOutlineView(outlineView), foundation.MakeIndexSet(proposedSelectionIndexes))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_SelectionShouldChangeInOutlineView
func outlineViewDelegate_SelectionShouldChangeInOutlineView(id int64, outlineView unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.SelectionShouldChangeInOutlineView(MakeOutlineView(outlineView))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineViewSelectionIsChanging
func outlineViewDelegate_OutlineViewSelectionIsChanging(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewSelectionIsChanging(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewSelectionDidChange
func outlineViewDelegate_OutlineViewSelectionDidChange(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewSelectionDidChange(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineView_WillDisplayCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_WillDisplayCell_ForTableColumn_Item(id int64, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_WillDisplayCell_ForTableColumn_Item(MakeOutlineView(outlineView), objc.MakeObject(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export outlineViewDelegate_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(id int64, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(MakeOutlineView(outlineView), objc.MakeObject(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export outlineViewDelegate_OutlineView_DataCellForTableColumn_Item
func outlineViewDelegate_OutlineView_DataCellForTableColumn_Item(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_DataCellForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldShowOutlineCellForItem
func outlineViewDelegate_OutlineView_ShouldShowOutlineCellForItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldShowOutlineCellForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldShowCellExpansionForTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldShowCellExpansionForTableColumn_Item(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldShowCellExpansionForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldReorderColumn_ToColumn
func outlineViewDelegate_OutlineView_ShouldReorderColumn_ToColumn(id int64, outlineView unsafe.Pointer, columnIndex C.int, newColumnIndex C.int) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldReorderColumn_ToColumn(MakeOutlineView(outlineView), int(columnIndex), int(newColumnIndex))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineViewColumnDidMove
func outlineViewDelegate_OutlineViewColumnDidMove(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewColumnDidMove(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewColumnDidResize
func outlineViewDelegate_OutlineViewColumnDidResize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewColumnDidResize(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemWillExpand
func outlineViewDelegate_OutlineViewItemWillExpand(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewItemWillExpand(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemDidExpand
func outlineViewDelegate_OutlineViewItemDidExpand(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewItemDidExpand(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemWillCollapse
func outlineViewDelegate_OutlineViewItemWillCollapse(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewItemWillCollapse(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemDidCollapse
func outlineViewDelegate_OutlineViewItemDidCollapse(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineViewItemDidCollapse(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineView_ShouldEditTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldEditTableColumn_Item(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldEditTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_MouseDownInHeaderOfTableColumn
func outlineViewDelegate_OutlineView_MouseDownInHeaderOfTableColumn(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_MouseDownInHeaderOfTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_DidClickTableColumn
func outlineViewDelegate_OutlineView_DidClickTableColumn(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_DidClickTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_DidDragTableColumn
func outlineViewDelegate_OutlineView_DidDragTableColumn(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_DidDragTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_HeightOfRowByItem
func outlineViewDelegate_OutlineView_HeightOfRowByItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.double {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_HeightOfRowByItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.double(float64(result))
}

//export outlineViewDelegate_OutlineView_SizeToFitWidthOfColumn
func outlineViewDelegate_OutlineView_SizeToFitWidthOfColumn(id int64, outlineView unsafe.Pointer, column C.int) C.double {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_SizeToFitWidthOfColumn(MakeOutlineView(outlineView), int(column))
	return C.double(float64(result))
}

//export outlineViewDelegate_OutlineView_TintConfigurationForItem
func outlineViewDelegate_OutlineView_TintConfigurationForItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_TintConfigurationForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldTrackCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldTrackCell_ForTableColumn_Item(id int64, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldTrackCell_ForTableColumn_Item(MakeOutlineView(outlineView), MakeCell(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_IsGroupItem
func outlineViewDelegate_OutlineView_IsGroupItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_IsGroupItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_DidAddRowView_ForRow
func outlineViewDelegate_OutlineView_DidAddRowView_ForRow(id int64, outlineView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_DidAddRowView_ForRow(MakeOutlineView(outlineView), MakeTableRowView(rowView), int(row))
}

//export outlineViewDelegate_OutlineView_DidRemoveRowView_ForRow
func outlineViewDelegate_OutlineView_DidRemoveRowView_ForRow(id int64, outlineView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.OutlineView_DidRemoveRowView_ForRow(MakeOutlineView(outlineView), MakeTableRowView(rowView), int(row))
}

//export outlineViewDelegate_OutlineView_RowViewForItem
func outlineViewDelegate_OutlineView_RowViewForItem(id int64, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_RowViewForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ViewForTableColumn_Item
func outlineViewDelegate_OutlineView_ViewForTableColumn_Item(id int64, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.OutlineView_ViewForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_Control_IsValidObject
func outlineViewDelegate_Control_IsValidObject(id int64, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export outlineViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func outlineViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export outlineViewDelegate_Control_DidFailToFormatString_ErrorDescription
func outlineViewDelegate_Control_DidFailToFormatString_ErrorDescription(id int64, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextShouldBeginEditing
func outlineViewDelegate_Control_TextShouldBeginEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextShouldEndEditing
func outlineViewDelegate_Control_TextShouldEndEditing(id int64, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextView_DoCommandBySelector
func outlineViewDelegate_Control_TextView_DoCommandBySelector(id int64, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.MakeSelector(commandSelector))
	return C.bool(result)
}

//export outlineViewDelegate_ControlTextDidBeginEditing
func outlineViewDelegate_ControlTextDidBeginEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export outlineViewDelegate_ControlTextDidChange
func outlineViewDelegate_ControlTextDidChange(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export outlineViewDelegate_ControlTextDidEndEditing
func outlineViewDelegate_ControlTextDidEndEditing(id int64, obj unsafe.Pointer) {
	delegate := resources.Get(id).(*OutlineViewDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export OutlineViewDelegate_RespondsTo
func OutlineViewDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*OutlineViewDelegate)
	switch selName {
	case "outlineView:shouldExpandItem:":
		return delegate.OutlineView_ShouldExpandItem != nil
	case "outlineView:shouldCollapseItem:":
		return delegate.OutlineView_ShouldCollapseItem != nil
	case "outlineView:typeSelectStringForTableColumn:item:":
		return delegate.OutlineView_TypeSelectStringForTableColumn_Item != nil
	case "outlineView:nextTypeSelectMatchFromItem:toItem:forString:":
		return delegate.OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString != nil
	case "outlineView:shouldTypeSelectForEvent:withCurrentSearchString:":
		return delegate.OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
	case "outlineView:shouldSelectTableColumn:":
		return delegate.OutlineView_ShouldSelectTableColumn != nil
	case "outlineView:shouldSelectItem:":
		return delegate.OutlineView_ShouldSelectItem != nil
	case "outlineView:selectionIndexesForProposedSelection:":
		return delegate.OutlineView_SelectionIndexesForProposedSelection != nil
	case "selectionShouldChangeInOutlineView:":
		return delegate.SelectionShouldChangeInOutlineView != nil
	case "outlineViewSelectionIsChanging:":
		return delegate.OutlineViewSelectionIsChanging != nil
	case "outlineViewSelectionDidChange:":
		return delegate.OutlineViewSelectionDidChange != nil
	case "outlineView:willDisplayCell:forTableColumn:item:":
		return delegate.OutlineView_WillDisplayCell_ForTableColumn_Item != nil
	case "outlineView:willDisplayOutlineCell:forTableColumn:item:":
		return delegate.OutlineView_WillDisplayOutlineCell_ForTableColumn_Item != nil
	case "outlineView:dataCellForTableColumn:item:":
		return delegate.OutlineView_DataCellForTableColumn_Item != nil
	case "outlineView:shouldShowOutlineCellForItem:":
		return delegate.OutlineView_ShouldShowOutlineCellForItem != nil
	case "outlineView:shouldShowCellExpansionForTableColumn:item:":
		return delegate.OutlineView_ShouldShowCellExpansionForTableColumn_Item != nil
	case "outlineView:shouldReorderColumn:toColumn:":
		return delegate.OutlineView_ShouldReorderColumn_ToColumn != nil
	case "outlineViewColumnDidMove:":
		return delegate.OutlineViewColumnDidMove != nil
	case "outlineViewColumnDidResize:":
		return delegate.OutlineViewColumnDidResize != nil
	case "outlineViewItemWillExpand:":
		return delegate.OutlineViewItemWillExpand != nil
	case "outlineViewItemDidExpand:":
		return delegate.OutlineViewItemDidExpand != nil
	case "outlineViewItemWillCollapse:":
		return delegate.OutlineViewItemWillCollapse != nil
	case "outlineViewItemDidCollapse:":
		return delegate.OutlineViewItemDidCollapse != nil
	case "outlineView:shouldEditTableColumn:item:":
		return delegate.OutlineView_ShouldEditTableColumn_Item != nil
	case "outlineView:mouseDownInHeaderOfTableColumn:":
		return delegate.OutlineView_MouseDownInHeaderOfTableColumn != nil
	case "outlineView:didClickTableColumn:":
		return delegate.OutlineView_DidClickTableColumn != nil
	case "outlineView:didDragTableColumn:":
		return delegate.OutlineView_DidDragTableColumn != nil
	case "outlineView:heightOfRowByItem:":
		return delegate.OutlineView_HeightOfRowByItem != nil
	case "outlineView:sizeToFitWidthOfColumn:":
		return delegate.OutlineView_SizeToFitWidthOfColumn != nil
	case "outlineView:tintConfigurationForItem:":
		return delegate.OutlineView_TintConfigurationForItem != nil
	case "outlineView:shouldTrackCell:forTableColumn:item:":
		return delegate.OutlineView_ShouldTrackCell_ForTableColumn_Item != nil
	case "outlineView:isGroupItem:":
		return delegate.OutlineView_IsGroupItem != nil
	case "outlineView:didAddRowView:forRow:":
		return delegate.OutlineView_DidAddRowView_ForRow != nil
	case "outlineView:didRemoveRowView:forRow:":
		return delegate.OutlineView_DidRemoveRowView_ForRow != nil
	case "outlineView:rowViewForItem:":
		return delegate.OutlineView_RowViewForItem != nil
	case "outlineView:viewForTableColumn:item:":
		return delegate.OutlineView_ViewForTableColumn_Item != nil
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

//export deleteOutlineViewDelegate
func deleteOutlineViewDelegate(id int64) {
	resources.Delete(id)
}
