package appkit

// #include "outline_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func MakeOutlineView(ptr unsafe.Pointer) *NSOutlineView {
	if ptr == nil {
		return nil
	}
	return &NSOutlineView{
		NSTableView: *MakeTableView(ptr),
	}
}

func AllocOutlineView() *NSOutlineView {
	return MakeOutlineView(C.C_OutlineView_Alloc())
}

func (n *NSOutlineView) InitWithCoder(coder foundation.Coder) OutlineView {
	result_ := C.C_NSOutlineView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeOutlineView(result_)
}

func (n *NSOutlineView) InitWithFrame(frameRect foundation.Rect) OutlineView {
	result_ := C.C_NSOutlineView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeOutlineView(result_)
}

func (n *NSOutlineView) Init() OutlineView {
	result_ := C.C_NSOutlineView_Init(n.Ptr())
	return MakeOutlineView(result_)
}

func (n *NSOutlineView) IsExpandable(item objc.Object) bool {
	result_ := C.C_NSOutlineView_IsExpandable(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n *NSOutlineView) IsItemExpanded(item objc.Object) bool {
	result_ := C.C_NSOutlineView_IsItemExpanded(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n *NSOutlineView) ExpandItem(item objc.Object) {
	C.C_NSOutlineView_ExpandItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSOutlineView) ExpandItem_ExpandChildren(item objc.Object, expandChildren bool) {
	C.C_NSOutlineView_ExpandItem_ExpandChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(expandChildren))
}

func (n *NSOutlineView) CollapseItem(item objc.Object) {
	C.C_NSOutlineView_CollapseItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSOutlineView) CollapseItem_CollapseChildren(item objc.Object, collapseChildren bool) {
	C.C_NSOutlineView_CollapseItem_CollapseChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(collapseChildren))
}

func (n *NSOutlineView) ReloadItem(item objc.Object) {
	C.C_NSOutlineView_ReloadItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSOutlineView) ReloadItem_ReloadChildren(item objc.Object, reloadChildren bool) {
	C.C_NSOutlineView_ReloadItem_ReloadChildren(n.Ptr(), objc.ExtractPtr(item), C.bool(reloadChildren))
}

func (n *NSOutlineView) ItemAtRow(row int) objc.Object {
	result_ := C.C_NSOutlineView_ItemAtRow(n.Ptr(), C.int(row))
	return objc.MakeObject(result_)
}

func (n *NSOutlineView) RowForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_RowForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n *NSOutlineView) LevelForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_LevelForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n *NSOutlineView) LevelForRow(row int) int {
	result_ := C.C_NSOutlineView_LevelForRow(n.Ptr(), C.int(row))
	return int(result_)
}

func (n *NSOutlineView) SetDropItem_DropChildIndex(item objc.Object, index int) {
	C.C_NSOutlineView_SetDropItem_DropChildIndex(n.Ptr(), objc.ExtractPtr(item), C.int(index))
}

func (n *NSOutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	result_ := C.C_NSOutlineView_ShouldCollapseAutoExpandedItemsForDeposited(n.Ptr(), C.bool(deposited))
	return bool(result_)
}

func (n *NSOutlineView) ParentForItem(item objc.Object) objc.Object {
	result_ := C.C_NSOutlineView_ParentForItem(n.Ptr(), objc.ExtractPtr(item))
	return objc.MakeObject(result_)
}

func (n *NSOutlineView) ChildIndexForItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_ChildIndexForItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n *NSOutlineView) Child_OfItem(index int, item objc.Object) objc.Object {
	result_ := C.C_NSOutlineView_Child_OfItem(n.Ptr(), C.int(index), objc.ExtractPtr(item))
	return objc.MakeObject(result_)
}

func (n *NSOutlineView) NumberOfChildrenOfItem(item objc.Object) int {
	result_ := C.C_NSOutlineView_NumberOfChildrenOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n *NSOutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	result_ := C.C_NSOutlineView_FrameOfOutlineCellAtRow(n.Ptr(), C.int(row))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSOutlineView) InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions) {
	C.C_NSOutlineView_InsertItemsAtIndexes_InParent_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), C.uint(uint(animationOptions)))
}

func (n *NSOutlineView) MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.Object, toIndex int, newParent objc.Object) {
	C.C_NSOutlineView_MoveItemAtIndex_InParent_ToIndex_InParent(n.Ptr(), C.int(fromIndex), objc.ExtractPtr(oldParent), C.int(toIndex), objc.ExtractPtr(newParent))
}

func (n *NSOutlineView) RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IndexSet, parent objc.Object, animationOptions TableViewAnimationOptions) {
	C.C_NSOutlineView_RemoveItemsAtIndexes_InParent_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), objc.ExtractPtr(parent), C.uint(uint(animationOptions)))
}

func (n *NSOutlineView) StronglyReferencesItems() bool {
	result_ := C.C_NSOutlineView_StronglyReferencesItems(n.Ptr())
	return bool(result_)
}

func (n *NSOutlineView) SetStronglyReferencesItems(value bool) {
	C.C_NSOutlineView_SetStronglyReferencesItems(n.Ptr(), C.bool(value))
}

func (n *NSOutlineView) OutlineTableColumn() TableColumn {
	result_ := C.C_NSOutlineView_OutlineTableColumn(n.Ptr())
	return MakeTableColumn(result_)
}

func (n *NSOutlineView) SetOutlineTableColumn(value TableColumn) {
	C.C_NSOutlineView_SetOutlineTableColumn(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSOutlineView) AutoresizesOutlineColumn() bool {
	result_ := C.C_NSOutlineView_AutoresizesOutlineColumn(n.Ptr())
	return bool(result_)
}

func (n *NSOutlineView) SetAutoresizesOutlineColumn(value bool) {
	C.C_NSOutlineView_SetAutoresizesOutlineColumn(n.Ptr(), C.bool(value))
}

func (n *NSOutlineView) IndentationPerLevel() coregraphics.Float {
	result_ := C.C_NSOutlineView_IndentationPerLevel(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSOutlineView) SetIndentationPerLevel(value coregraphics.Float) {
	C.C_NSOutlineView_SetIndentationPerLevel(n.Ptr(), C.double(float64(value)))
}

func (n *NSOutlineView) IndentationMarkerFollowsCell() bool {
	result_ := C.C_NSOutlineView_IndentationMarkerFollowsCell(n.Ptr())
	return bool(result_)
}

func (n *NSOutlineView) SetIndentationMarkerFollowsCell(value bool) {
	C.C_NSOutlineView_SetIndentationMarkerFollowsCell(n.Ptr(), C.bool(value))
}

func (n *NSOutlineView) AutosaveExpandedItems() bool {
	result_ := C.C_NSOutlineView_AutosaveExpandedItems(n.Ptr())
	return bool(result_)
}

func (n *NSOutlineView) SetAutosaveExpandedItems(value bool) {
	C.C_NSOutlineView_SetAutosaveExpandedItems(n.Ptr(), C.bool(value))
}
