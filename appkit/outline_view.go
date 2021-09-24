package appkit

// #include "outline_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type OutlineView interface {
	TableView
	IsExpandable(item objc.Object) bool
	IsItemExpanded(item objc.Object) bool
	ExpandItem(item objc.Object)
	ExpandItem_ExpandChildren(item objc.Object, expandChildren bool)
	CollapseItem(item objc.Object)
	CollapseItem_CollapseChildren(item objc.Object, collapseChildren bool)
	ReloadItem(item objc.Object)
	ReloadItem_ReloadChildren(item objc.Object, reloadChildren bool)
	ItemAtRow(row int) objc.Object
	RowForItem(item objc.Object) int
	LevelForItem(item objc.Object) int
	LevelForRow(row int) int
	SetDropItem_DropChildIndex(item objc.Object, index int)
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
	ParentForItem(item objc.Object) objc.Object
	ChildIndexForItem(item objc.Object) int
	Child_OfItem(index int, item objc.Object) objc.Object
	NumberOfChildrenOfItem(item objc.Object) int
	FrameOfOutlineCellAtRow(row int) foundation.Rect
	InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions)
	MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.Object, toIndex int, newParent objc.Object)
	RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	OutlineTableColumn() TableColumn
	SetOutlineTableColumn(value TableColumn)
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	IndentationPerLevel() coregraphics.Float
	SetIndentationPerLevel(value coregraphics.Float)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
}

type NSOutlineView struct {
	NSTableView
}

func MakeOutlineView(ptr unsafe.Pointer) NSOutlineView {
	return NSOutlineView{
		NSTableView: MakeTableView(ptr),
	}
}

func (n NSOutlineView) InitWithCoder(coder foundation.Coder) NSOutlineView {
	result_ := C.C_NSOutlineView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeOutlineView(result_)
}

func (n NSOutlineView) InitWithFrame(frameRect foundation.Rect) NSOutlineView {
	result_ := C.C_NSOutlineView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeOutlineView(result_)
}

func (n NSOutlineView) Init() NSOutlineView {
	result_ := C.C_NSOutlineView_Init(n.Ptr())
	return MakeOutlineView(result_)
}

func AllocOutlineView() NSOutlineView {
	result_ := C.C_NSOutlineView_AllocOutlineView()
	return MakeOutlineView(result_)
}

func NewOutlineView() NSOutlineView {
	result_ := C.C_NSOutlineView_NewOutlineView()
	return MakeOutlineView(result_)
}

func (n NSOutlineView) Autorelease() NSOutlineView {
	result_ := C.C_NSOutlineView_Autorelease(n.Ptr())
	return MakeOutlineView(result_)
}

func (n NSOutlineView) Retain() NSOutlineView {
	result_ := C.C_NSOutlineView_Retain(n.Ptr())
	return MakeOutlineView(result_)
}

func (n NSOutlineView) IsExpandable(item objc.Object) bool {
	result_ := C.C_NSOutlineView_IsExpandable(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n NSOutlineView) IsItemExpanded(item objc.Object) bool {
	result_ := C.C_NSOutlineView_IsItemExpanded(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n NSOutlineView) ExpandItem(item objc.Object) {
	C.C_NSOutlineView_ExpandItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSOutlineView) ExpandItem_ExpandChildren(item objc.Object, expandChildren bool) {
	C.C_NSOutlineView_ExpandItem_ExpandChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(expandChildren))
}

func (n NSOutlineView) CollapseItem(item objc.Object) {
	C.C_NSOutlineView_CollapseItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSOutlineView) CollapseItem_CollapseChildren(item objc.Object, collapseChildren bool) {
	C.C_NSOutlineView_CollapseItem_CollapseChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(collapseChildren))
}

func (n NSOutlineView) ReloadItem(item objc.Object) {
	C.C_NSOutlineView_ReloadItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSOutlineView) ReloadItem_ReloadChildren(item objc.Object, reloadChildren bool) {
	C.C_NSOutlineView_ReloadItem_ReloadChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(reloadChildren))
}

func (n NSOutlineView) ItemAtRow(row int) objc.Object {
	result_ := C.C_NSOutlineView_ItemAtRow(n.Ptr(), C.int(row))
	return objc.MakeObject(result_)
}

func (n NSOutlineView) RowForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_RowForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSOutlineView) LevelForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_LevelForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSOutlineView) LevelForRow(row int) int {
	result_ := C.C_NSOutlineView_LevelForRow(n.Ptr(), C.int(row))
	return int(result_)
}

func (n NSOutlineView) SetDropItem_DropChildIndex(item objc.Object, index int) {
	C.C_NSOutlineView_SetDropItem_DropChildIndex(n.Ptr(), objc.ExtractPtr(item), C.int(index))
}

func (n NSOutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	result_ := C.C_NSOutlineView_ShouldCollapseAutoExpandedItemsForDeposited(n.Ptr(), C.bool(deposited))
	return bool(result_)
}

func (n NSOutlineView) ParentForItem(item objc.Object) objc.Object {
	result_ := C.C_NSOutlineView_ParentForItem(n.Ptr(), objc.ExtractPtr(item))
	return objc.MakeObject(result_)
}

func (n NSOutlineView) ChildIndexForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_ChildIndexForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSOutlineView) Child_OfItem(index int, item objc.Object) objc.Object {
	result_ := C.C_NSOutlineView_Child_OfItem(n.Ptr(), C.int(index), objc.ExtractPtr(item))
	return objc.MakeObject(result_)
}

func (n NSOutlineView) NumberOfChildrenOfItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_NumberOfChildrenOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSOutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	result_ := C.C_NSOutlineView_FrameOfOutlineCellAtRow(n.Ptr(), C.int(row))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSOutlineView) InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions) {
	C.C_NSOutlineView_InsertItemsAtIndexes_InParent_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), C.uint(uint(animationOptions)))
}

func (n NSOutlineView) MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.Object, toIndex int, newParent objc.Object) {
	C.C_NSOutlineView_MoveItemAtIndex_InParent_ToIndex_InParent(n.Ptr(), C.int(fromIndex), objc.ExtractPtr(oldParent), C.int(toIndex), objc.ExtractPtr(newParent))
}

func (n NSOutlineView) RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions) {
	C.C_NSOutlineView_RemoveItemsAtIndexes_InParent_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), C.uint(uint(animationOptions)))
}

func (n NSOutlineView) StronglyReferencesItems() bool {
	result_ := C.C_NSOutlineView_StronglyReferencesItems(n.Ptr())
	return bool(result_)
}

func (n NSOutlineView) SetStronglyReferencesItems(value bool) {
	C.C_NSOutlineView_SetStronglyReferencesItems(n.Ptr(), C.bool(value))
}

func (n NSOutlineView) OutlineTableColumn() TableColumn {
	result_ := C.C_NSOutlineView_OutlineTableColumn(n.Ptr())
	return MakeTableColumn(result_)
}

func (n NSOutlineView) SetOutlineTableColumn(value TableColumn) {
	C.C_NSOutlineView_SetOutlineTableColumn(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSOutlineView) AutoresizesOutlineColumn() bool {
	result_ := C.C_NSOutlineView_AutoresizesOutlineColumn(n.Ptr())
	return bool(result_)
}

func (n NSOutlineView) SetAutoresizesOutlineColumn(value bool) {
	C.C_NSOutlineView_SetAutoresizesOutlineColumn(n.Ptr(), C.bool(value))
}

func (n NSOutlineView) IndentationPerLevel() coregraphics.Float {
	result_ := C.C_NSOutlineView_IndentationPerLevel(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSOutlineView) SetIndentationPerLevel(value coregraphics.Float) {
	C.C_NSOutlineView_SetIndentationPerLevel(n.Ptr(), C.double(float64(value)))
}

func (n NSOutlineView) IndentationMarkerFollowsCell() bool {
	result_ := C.C_NSOutlineView_IndentationMarkerFollowsCell(n.Ptr())
	return bool(result_)
}

func (n NSOutlineView) SetIndentationMarkerFollowsCell(value bool) {
	C.C_NSOutlineView_SetIndentationMarkerFollowsCell(n.Ptr(), C.bool(value))
}

func (n NSOutlineView) AutosaveExpandedItems() bool {
	result_ := C.C_NSOutlineView_AutosaveExpandedItems(n.Ptr())
	return bool(result_)
}

func (n NSOutlineView) SetAutosaveExpandedItems(value bool) {
	C.C_NSOutlineView_SetAutosaveExpandedItems(n.Ptr(), C.bool(value))
}

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
	Control_TextView_DoCommandBySelector                         func(control Control, textView TextView, commandSelector objc.Selector) bool
	ControlTextDidBeginEditing                                   func(obj foundation.Notification)
	ControlTextDidChange                                         func(obj foundation.Notification)
	ControlTextDidEndEditing                                     func(obj foundation.Notification)
}

func (delegate *OutlineViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapOutlineViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export outlineViewDelegate_OutlineView_ShouldExpandItem
func outlineViewDelegate_OutlineView_ShouldExpandItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldExpandItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldCollapseItem
func outlineViewDelegate_OutlineView_ShouldCollapseItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldCollapseItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_TypeSelectStringForTableColumn_Item
func outlineViewDelegate_OutlineView_TypeSelectStringForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_TypeSelectStringForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return foundation.NewString(result).Ptr()
}

//export outlineViewDelegate_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString
func outlineViewDelegate_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(hp C.uintptr_t, outlineView unsafe.Pointer, startItem unsafe.Pointer, endItem unsafe.Pointer, searchString unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(MakeOutlineView(outlineView), objc.MakeObject(startItem), objc.MakeObject(endItem), foundation.MakeString(searchString).String())
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString
func outlineViewDelegate_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(hp C.uintptr_t, outlineView unsafe.Pointer, event unsafe.Pointer, searchString unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(MakeOutlineView(outlineView), MakeEvent(event), foundation.MakeString(searchString).String())
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldSelectTableColumn
func outlineViewDelegate_OutlineView_ShouldSelectTableColumn(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldSelectTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldSelectItem
func outlineViewDelegate_OutlineView_ShouldSelectItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldSelectItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_SelectionIndexesForProposedSelection
func outlineViewDelegate_OutlineView_SelectionIndexesForProposedSelection(hp C.uintptr_t, outlineView unsafe.Pointer, proposedSelectionIndexes unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_SelectionIndexesForProposedSelection(MakeOutlineView(outlineView), foundation.MakeIndexSet(proposedSelectionIndexes))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_SelectionShouldChangeInOutlineView
func outlineViewDelegate_SelectionShouldChangeInOutlineView(hp C.uintptr_t, outlineView unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.SelectionShouldChangeInOutlineView(MakeOutlineView(outlineView))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineViewSelectionIsChanging
func outlineViewDelegate_OutlineViewSelectionIsChanging(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewSelectionIsChanging(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewSelectionDidChange
func outlineViewDelegate_OutlineViewSelectionDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewSelectionDidChange(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineView_WillDisplayCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_WillDisplayCell_ForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_WillDisplayCell_ForTableColumn_Item(MakeOutlineView(outlineView), objc.MakeObject(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export outlineViewDelegate_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(MakeOutlineView(outlineView), objc.MakeObject(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
}

//export outlineViewDelegate_OutlineView_DataCellForTableColumn_Item
func outlineViewDelegate_OutlineView_DataCellForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_DataCellForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldShowOutlineCellForItem
func outlineViewDelegate_OutlineView_ShouldShowOutlineCellForItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldShowOutlineCellForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldShowCellExpansionForTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldShowCellExpansionForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldShowCellExpansionForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_ShouldReorderColumn_ToColumn
func outlineViewDelegate_OutlineView_ShouldReorderColumn_ToColumn(hp C.uintptr_t, outlineView unsafe.Pointer, columnIndex C.int, newColumnIndex C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldReorderColumn_ToColumn(MakeOutlineView(outlineView), int(columnIndex), int(newColumnIndex))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineViewColumnDidMove
func outlineViewDelegate_OutlineViewColumnDidMove(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewColumnDidMove(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewColumnDidResize
func outlineViewDelegate_OutlineViewColumnDidResize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewColumnDidResize(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemWillExpand
func outlineViewDelegate_OutlineViewItemWillExpand(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewItemWillExpand(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemDidExpand
func outlineViewDelegate_OutlineViewItemDidExpand(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewItemDidExpand(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemWillCollapse
func outlineViewDelegate_OutlineViewItemWillCollapse(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewItemWillCollapse(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineViewItemDidCollapse
func outlineViewDelegate_OutlineViewItemDidCollapse(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineViewItemDidCollapse(foundation.MakeNotification(notification))
}

//export outlineViewDelegate_OutlineView_ShouldEditTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldEditTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldEditTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_MouseDownInHeaderOfTableColumn
func outlineViewDelegate_OutlineView_MouseDownInHeaderOfTableColumn(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_MouseDownInHeaderOfTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_DidClickTableColumn
func outlineViewDelegate_OutlineView_DidClickTableColumn(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_DidClickTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_DidDragTableColumn
func outlineViewDelegate_OutlineView_DidDragTableColumn(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_DidDragTableColumn(MakeOutlineView(outlineView), MakeTableColumn(tableColumn))
}

//export outlineViewDelegate_OutlineView_HeightOfRowByItem
func outlineViewDelegate_OutlineView_HeightOfRowByItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.double {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_HeightOfRowByItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.double(float64(result))
}

//export outlineViewDelegate_OutlineView_SizeToFitWidthOfColumn
func outlineViewDelegate_OutlineView_SizeToFitWidthOfColumn(hp C.uintptr_t, outlineView unsafe.Pointer, column C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_SizeToFitWidthOfColumn(MakeOutlineView(outlineView), int(column))
	return C.double(float64(result))
}

//export outlineViewDelegate_OutlineView_TintConfigurationForItem
func outlineViewDelegate_OutlineView_TintConfigurationForItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_TintConfigurationForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ShouldTrackCell_ForTableColumn_Item
func outlineViewDelegate_OutlineView_ShouldTrackCell_ForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ShouldTrackCell_ForTableColumn_Item(MakeOutlineView(outlineView), MakeCell(cell), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_IsGroupItem
func outlineViewDelegate_OutlineView_IsGroupItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_IsGroupItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return C.bool(result)
}

//export outlineViewDelegate_OutlineView_DidAddRowView_ForRow
func outlineViewDelegate_OutlineView_DidAddRowView_ForRow(hp C.uintptr_t, outlineView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_DidAddRowView_ForRow(MakeOutlineView(outlineView), MakeTableRowView(rowView), int(row))
}

//export outlineViewDelegate_OutlineView_DidRemoveRowView_ForRow
func outlineViewDelegate_OutlineView_DidRemoveRowView_ForRow(hp C.uintptr_t, outlineView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.OutlineView_DidRemoveRowView_ForRow(MakeOutlineView(outlineView), MakeTableRowView(rowView), int(row))
}

//export outlineViewDelegate_OutlineView_RowViewForItem
func outlineViewDelegate_OutlineView_RowViewForItem(hp C.uintptr_t, outlineView unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_RowViewForItem(MakeOutlineView(outlineView), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_OutlineView_ViewForTableColumn_Item
func outlineViewDelegate_OutlineView_ViewForTableColumn_Item(hp C.uintptr_t, outlineView unsafe.Pointer, tableColumn unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.OutlineView_ViewForTableColumn_Item(MakeOutlineView(outlineView), MakeTableColumn(tableColumn), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export outlineViewDelegate_Control_IsValidObject
func outlineViewDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export outlineViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func outlineViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export outlineViewDelegate_Control_DidFailToFormatString_ErrorDescription
func outlineViewDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextShouldBeginEditing
func outlineViewDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextShouldEndEditing
func outlineViewDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export outlineViewDelegate_Control_TextView_DoCommandBySelector
func outlineViewDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export outlineViewDelegate_ControlTextDidBeginEditing
func outlineViewDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export outlineViewDelegate_ControlTextDidChange
func outlineViewDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export outlineViewDelegate_ControlTextDidEndEditing
func outlineViewDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export OutlineViewDelegate_RespondsTo
func OutlineViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*OutlineViewDelegate)
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
