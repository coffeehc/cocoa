package appkit

// #include "browser.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Browser interface {
	Control
	Tile()
	SelectedRowIndexesInColumn(column int) foundation.IndexSet
	SelectRowIndexes_InColumn(indexes foundation.IndexSet, column int)
	SelectedCellInColumn(column int) objc.Object
	SelectAll(sender objc.Object)
	SelectedRowInColumn(column int) int
	SelectRow_InColumn(row int, column int)
	LoadedCellAtRow_Column(row int, col int) objc.Object
	EditItemAtIndexPath_WithEvent_Select(indexPath foundation.IndexPath, event Event, _select bool)
	ItemAtIndexPath(indexPath foundation.IndexPath) objc.Object
	ItemAtRow_InColumn(row int, column int) objc.Object
	IndexPathForColumn(column int) foundation.IndexPath
	IsLeafItem(item objc.Object) bool
	ParentForItemsInColumn(column int) objc.Object
	Path() string
	SetPath(path string) bool
	PathToColumn(column int) string
	AddColumn()
	ValidateVisibleColumns()
	LoadColumnZero()
	ReloadColumn(column int)
	TitleOfColumn(column int) string
	SetTitle_OfColumn(_string string, column int)
	DrawTitleOfColumn_InRect(column int, rect foundation.Rect)
	TitleFrameOfColumn(column int) foundation.Rect
	NoteHeightOfRowsWithIndexesChanged_InColumn(indexSet foundation.IndexSet, columnIndex int)
	ReloadDataForRowIndexes_InColumn(rowIndexes foundation.IndexSet, column int)
	ScrollColumnToVisible(column int)
	ScrollColumnsLeftBy(shiftAmount int)
	ScrollColumnsRightBy(shiftAmount int)
	ScrollRowToVisible_InColumn(row int, column int)
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	CanDragRowsWithIndexes_InColumn_WithEvent(rowIndexes foundation.IndexSet, column int, event Event) bool
	FrameOfColumn(column int) foundation.Rect
	FrameOfInsideOfColumn(column int) foundation.Rect
	FrameOfRow_InColumn(row int, column int) foundation.Rect
	SendAction() bool
	DoClick(sender objc.Object)
	DoDoubleClick(sender objc.Object)
	ColumnContentWidthForColumnWidth(columnWidth coregraphics.Float) coregraphics.Float
	ColumnWidthForColumnContentWidth(columnContentWidth coregraphics.Float) coregraphics.Float
	WidthOfColumn(column int) coregraphics.Float
	SetWidth_OfColumn(columnWidth coregraphics.Float, columnIndex int)
	DefaultColumnWidth() coregraphics.Float
	SetDefaultColumnWidth(columnWidth coregraphics.Float)
	ReusesColumns() bool
	SetReusesColumns(value bool)
	MaxVisibleColumns() int
	SetMaxVisibleColumns(value int)
	AutohidesScroller() bool
	SetAutohidesScroller(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	MinColumnWidth() coregraphics.Float
	SetMinColumnWidth(value coregraphics.Float)
	SeparatesColumns() bool
	SetSeparatesColumns(value bool)
	TakesTitleFromPreviousColumn() bool
	SetTakesTitleFromPreviousColumn(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	CellPrototype() objc.Object
	SetCellPrototype(value objc.Object)
	AllowsBranchSelection() bool
	SetAllowsBranchSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	SelectedCell() objc.Object
	SelectedCells() []Cell
	SelectionIndexPath() foundation.IndexPath
	SetSelectionIndexPath(value foundation.IndexPath)
	SelectionIndexPaths() []foundation.IndexPath
	SetSelectionIndexPaths(value []foundation.IndexPath)
	PathSeparator() string
	SetPathSeparator(value string)
	SelectedColumn() int
	LastColumn() int
	SetLastColumn(value int)
	FirstVisibleColumn() int
	NumberOfVisibleColumns() int
	LastVisibleColumn() int
	IsLoaded() bool
	IsTitled() bool
	SetTitled(value bool)
	TitleHeight() coregraphics.Float
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	DoubleAction() *objc.Selector
	SetDoubleAction(value *objc.Selector)
	SendsActionOnArrowKeys() bool
	SetSendsActionOnArrowKeys(value bool)
	ClickedColumn() int
	ClickedRow() int
	ColumnsAutosaveName() BrowserColumnsAutosaveName
	SetColumnsAutosaveName(value BrowserColumnsAutosaveName)
	ColumnResizingType() BrowserColumnResizingType
	SetColumnResizingType(value BrowserColumnResizingType)
	PrefersAllColumnUserResizing() bool
	SetPrefersAllColumnUserResizing(value bool)
	RowHeight() coregraphics.Float
	SetRowHeight(value coregraphics.Float)
}

type NSBrowser struct {
	NSControl
}

func MakeBrowser(ptr unsafe.Pointer) *NSBrowser {
	if ptr == nil {
		return nil
	}
	return &NSBrowser{
		NSControl: *MakeControl(ptr),
	}
}

func AllocBrowser() *NSBrowser {
	return MakeBrowser(C.C_Browser_Alloc())
}

func (n *NSBrowser) InitWithFrame(frameRect foundation.Rect) Browser {
	result_ := C.C_NSBrowser_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeBrowser(result_)
}

func (n *NSBrowser) InitWithCoder(coder foundation.Coder) Browser {
	result_ := C.C_NSBrowser_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeBrowser(result_)
}

func (n *NSBrowser) Init() Browser {
	result_ := C.C_NSBrowser_Init(n.Ptr())
	return MakeBrowser(result_)
}

func (n *NSBrowser) Tile() {
	C.C_NSBrowser_Tile(n.Ptr())
}

func (n *NSBrowser) SelectedRowIndexesInColumn(column int) foundation.IndexSet {
	result_ := C.C_NSBrowser_SelectedRowIndexesInColumn(n.Ptr(), C.int(column))
	return foundation.MakeIndexSet(result_)
}

func (n *NSBrowser) SelectRowIndexes_InColumn(indexes foundation.IndexSet, column int) {
	C.C_NSBrowser_SelectRowIndexes_InColumn(n.Ptr(), objc.ExtractPtr(indexes), C.int(column))
}

func (n *NSBrowser) SelectedCellInColumn(column int) objc.Object {
	result_ := C.C_NSBrowser_SelectedCellInColumn(n.Ptr(), C.int(column))
	return objc.MakeObject(result_)
}

func (n *NSBrowser) SelectAll(sender objc.Object) {
	C.C_NSBrowser_SelectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSBrowser) SelectedRowInColumn(column int) int {
	result_ := C.C_NSBrowser_SelectedRowInColumn(n.Ptr(), C.int(column))
	return int(result_)
}

func (n *NSBrowser) SelectRow_InColumn(row int, column int) {
	C.C_NSBrowser_SelectRow_InColumn(n.Ptr(), C.int(row), C.int(column))
}

func (n *NSBrowser) LoadedCellAtRow_Column(row int, col int) objc.Object {
	result_ := C.C_NSBrowser_LoadedCellAtRow_Column(n.Ptr(), C.int(row), C.int(col))
	return objc.MakeObject(result_)
}

func (n *NSBrowser) EditItemAtIndexPath_WithEvent_Select(indexPath foundation.IndexPath, event Event, _select bool) {
	C.C_NSBrowser_EditItemAtIndexPath_WithEvent_Select(n.Ptr(), objc.ExtractPtr(indexPath), objc.ExtractPtr(event), C.bool(_select))
}

func (n *NSBrowser) ItemAtIndexPath(indexPath foundation.IndexPath) objc.Object {
	result_ := C.C_NSBrowser_ItemAtIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return objc.MakeObject(result_)
}

func (n *NSBrowser) ItemAtRow_InColumn(row int, column int) objc.Object {
	result_ := C.C_NSBrowser_ItemAtRow_InColumn(n.Ptr(), C.int(row), C.int(column))
	return objc.MakeObject(result_)
}

func (n *NSBrowser) IndexPathForColumn(column int) foundation.IndexPath {
	result_ := C.C_NSBrowser_IndexPathForColumn(n.Ptr(), C.int(column))
	return foundation.MakeIndexPath(result_)
}

func (n *NSBrowser) IsLeafItem(item objc.Object) bool {
	result_ := C.C_NSBrowser_IsLeafItem(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n *NSBrowser) ParentForItemsInColumn(column int) objc.Object {
	result_ := C.C_NSBrowser_ParentForItemsInColumn(n.Ptr(), C.int(column))
	return objc.MakeObject(result_)
}

func (n *NSBrowser) Path() string {
	result_ := C.C_NSBrowser_Path(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSBrowser) SetPath(path string) bool {
	result_ := C.C_NSBrowser_SetPath(n.Ptr(), foundation.NewString(path).Ptr())
	return bool(result_)
}

func (n *NSBrowser) PathToColumn(column int) string {
	result_ := C.C_NSBrowser_PathToColumn(n.Ptr(), C.int(column))
	return foundation.MakeString(result_).String()
}

func (n *NSBrowser) AddColumn() {
	C.C_NSBrowser_AddColumn(n.Ptr())
}

func (n *NSBrowser) ValidateVisibleColumns() {
	C.C_NSBrowser_ValidateVisibleColumns(n.Ptr())
}

func (n *NSBrowser) LoadColumnZero() {
	C.C_NSBrowser_LoadColumnZero(n.Ptr())
}

func (n *NSBrowser) ReloadColumn(column int) {
	C.C_NSBrowser_ReloadColumn(n.Ptr(), C.int(column))
}

func (n *NSBrowser) TitleOfColumn(column int) string {
	result_ := C.C_NSBrowser_TitleOfColumn(n.Ptr(), C.int(column))
	return foundation.MakeString(result_).String()
}

func (n *NSBrowser) SetTitle_OfColumn(_string string, column int) {
	C.C_NSBrowser_SetTitle_OfColumn(n.Ptr(), foundation.NewString(_string).Ptr(), C.int(column))
}

func (n *NSBrowser) DrawTitleOfColumn_InRect(column int, rect foundation.Rect) {
	C.C_NSBrowser_DrawTitleOfColumn_InRect(n.Ptr(), C.int(column), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSBrowser) TitleFrameOfColumn(column int) foundation.Rect {
	result_ := C.C_NSBrowser_TitleFrameOfColumn(n.Ptr(), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBrowser) NoteHeightOfRowsWithIndexesChanged_InColumn(indexSet foundation.IndexSet, columnIndex int) {
	C.C_NSBrowser_NoteHeightOfRowsWithIndexesChanged_InColumn(n.Ptr(), objc.ExtractPtr(indexSet), C.int(columnIndex))
}

func (n *NSBrowser) ReloadDataForRowIndexes_InColumn(rowIndexes foundation.IndexSet, column int) {
	C.C_NSBrowser_ReloadDataForRowIndexes_InColumn(n.Ptr(), objc.ExtractPtr(rowIndexes), C.int(column))
}

func (n *NSBrowser) ScrollColumnToVisible(column int) {
	C.C_NSBrowser_ScrollColumnToVisible(n.Ptr(), C.int(column))
}

func (n *NSBrowser) ScrollColumnsLeftBy(shiftAmount int) {
	C.C_NSBrowser_ScrollColumnsLeftBy(n.Ptr(), C.int(shiftAmount))
}

func (n *NSBrowser) ScrollColumnsRightBy(shiftAmount int) {
	C.C_NSBrowser_ScrollColumnsRightBy(n.Ptr(), C.int(shiftAmount))
}

func (n *NSBrowser) ScrollRowToVisible_InColumn(row int, column int) {
	C.C_NSBrowser_ScrollRowToVisible_InColumn(n.Ptr(), C.int(row), C.int(column))
}

func (n *NSBrowser) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	C.C_NSBrowser_SetDraggingSourceOperationMask_ForLocal(n.Ptr(), C.uint(uint(mask)), C.bool(isLocal))
}

func (n *NSBrowser) CanDragRowsWithIndexes_InColumn_WithEvent(rowIndexes foundation.IndexSet, column int, event Event) bool {
	result_ := C.C_NSBrowser_CanDragRowsWithIndexes_InColumn_WithEvent(n.Ptr(), objc.ExtractPtr(rowIndexes), C.int(column), objc.ExtractPtr(event))
	return bool(result_)
}

func (n *NSBrowser) FrameOfColumn(column int) foundation.Rect {
	result_ := C.C_NSBrowser_FrameOfColumn(n.Ptr(), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBrowser) FrameOfInsideOfColumn(column int) foundation.Rect {
	result_ := C.C_NSBrowser_FrameOfInsideOfColumn(n.Ptr(), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBrowser) FrameOfRow_InColumn(row int, column int) foundation.Rect {
	result_ := C.C_NSBrowser_FrameOfRow_InColumn(n.Ptr(), C.int(row), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBrowser) SendAction() bool {
	result_ := C.C_NSBrowser_SendAction(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) DoClick(sender objc.Object) {
	C.C_NSBrowser_DoClick(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSBrowser) DoDoubleClick(sender objc.Object) {
	C.C_NSBrowser_DoDoubleClick(n.Ptr(), objc.ExtractPtr(sender))
}

func Browser_RemoveSavedColumnsWithAutosaveName(name BrowserColumnsAutosaveName) {
	C.C_NSBrowser_Browser_RemoveSavedColumnsWithAutosaveName(foundation.NewString(string(name)).Ptr())
}

func (n *NSBrowser) ColumnContentWidthForColumnWidth(columnWidth coregraphics.Float) coregraphics.Float {
	result_ := C.C_NSBrowser_ColumnContentWidthForColumnWidth(n.Ptr(), C.double(float64(columnWidth)))
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) ColumnWidthForColumnContentWidth(columnContentWidth coregraphics.Float) coregraphics.Float {
	result_ := C.C_NSBrowser_ColumnWidthForColumnContentWidth(n.Ptr(), C.double(float64(columnContentWidth)))
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) WidthOfColumn(column int) coregraphics.Float {
	result_ := C.C_NSBrowser_WidthOfColumn(n.Ptr(), C.int(column))
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) SetWidth_OfColumn(columnWidth coregraphics.Float, columnIndex int) {
	C.C_NSBrowser_SetWidth_OfColumn(n.Ptr(), C.double(float64(columnWidth)), C.int(columnIndex))
}

func (n *NSBrowser) DefaultColumnWidth() coregraphics.Float {
	result_ := C.C_NSBrowser_DefaultColumnWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) SetDefaultColumnWidth(columnWidth coregraphics.Float) {
	C.C_NSBrowser_SetDefaultColumnWidth(n.Ptr(), C.double(float64(columnWidth)))
}

func (n *NSBrowser) ReusesColumns() bool {
	result_ := C.C_NSBrowser_ReusesColumns(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetReusesColumns(value bool) {
	C.C_NSBrowser_SetReusesColumns(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) MaxVisibleColumns() int {
	result_ := C.C_NSBrowser_MaxVisibleColumns(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) SetMaxVisibleColumns(value int) {
	C.C_NSBrowser_SetMaxVisibleColumns(n.Ptr(), C.int(value))
}

func (n *NSBrowser) AutohidesScroller() bool {
	result_ := C.C_NSBrowser_AutohidesScroller(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetAutohidesScroller(value bool) {
	C.C_NSBrowser_SetAutohidesScroller(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) BackgroundColor() Color {
	result_ := C.C_NSBrowser_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSBrowser) SetBackgroundColor(value Color) {
	C.C_NSBrowser_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowser) MinColumnWidth() coregraphics.Float {
	result_ := C.C_NSBrowser_MinColumnWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) SetMinColumnWidth(value coregraphics.Float) {
	C.C_NSBrowser_SetMinColumnWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSBrowser) SeparatesColumns() bool {
	result_ := C.C_NSBrowser_SeparatesColumns(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetSeparatesColumns(value bool) {
	C.C_NSBrowser_SetSeparatesColumns(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) TakesTitleFromPreviousColumn() bool {
	result_ := C.C_NSBrowser_TakesTitleFromPreviousColumn(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetTakesTitleFromPreviousColumn(value bool) {
	C.C_NSBrowser_SetTakesTitleFromPreviousColumn(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) Delegate() objc.Object {
	result_ := C.C_NSBrowser_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSBrowser) SetDelegate(value objc.Object) {
	C.C_NSBrowser_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowser) CellPrototype() objc.Object {
	result_ := C.C_NSBrowser_CellPrototype(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSBrowser) SetCellPrototype(value objc.Object) {
	C.C_NSBrowser_SetCellPrototype(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowser) AllowsBranchSelection() bool {
	result_ := C.C_NSBrowser_AllowsBranchSelection(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetAllowsBranchSelection(value bool) {
	C.C_NSBrowser_SetAllowsBranchSelection(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) AllowsEmptySelection() bool {
	result_ := C.C_NSBrowser_AllowsEmptySelection(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetAllowsEmptySelection(value bool) {
	C.C_NSBrowser_SetAllowsEmptySelection(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) AllowsMultipleSelection() bool {
	result_ := C.C_NSBrowser_AllowsMultipleSelection(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetAllowsMultipleSelection(value bool) {
	C.C_NSBrowser_SetAllowsMultipleSelection(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) AllowsTypeSelect() bool {
	result_ := C.C_NSBrowser_AllowsTypeSelect(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetAllowsTypeSelect(value bool) {
	C.C_NSBrowser_SetAllowsTypeSelect(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) SelectedCell() objc.Object {
	result_ := C.C_NSBrowser_SelectedCell(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSBrowser) SelectedCells() []Cell {
	result_ := C.C_NSBrowser_SelectedCells(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Cell, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCell(r)
	}
	return goResult_
}

func (n *NSBrowser) SelectionIndexPath() foundation.IndexPath {
	result_ := C.C_NSBrowser_SelectionIndexPath(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n *NSBrowser) SetSelectionIndexPath(value foundation.IndexPath) {
	C.C_NSBrowser_SetSelectionIndexPath(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowser) SelectionIndexPaths() []foundation.IndexPath {
	result_ := C.C_NSBrowser_SelectionIndexPaths(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]foundation.IndexPath, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeIndexPath(r)
	}
	return goResult_
}

func (n *NSBrowser) SetSelectionIndexPaths(value []foundation.IndexPath) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSBrowser_SetSelectionIndexPaths(n.Ptr(), cValue)
}

func (n *NSBrowser) PathSeparator() string {
	result_ := C.C_NSBrowser_PathSeparator(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSBrowser) SetPathSeparator(value string) {
	C.C_NSBrowser_SetPathSeparator(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSBrowser) SelectedColumn() int {
	result_ := C.C_NSBrowser_SelectedColumn(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) LastColumn() int {
	result_ := C.C_NSBrowser_LastColumn(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) SetLastColumn(value int) {
	C.C_NSBrowser_SetLastColumn(n.Ptr(), C.int(value))
}

func (n *NSBrowser) FirstVisibleColumn() int {
	result_ := C.C_NSBrowser_FirstVisibleColumn(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) NumberOfVisibleColumns() int {
	result_ := C.C_NSBrowser_NumberOfVisibleColumns(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) LastVisibleColumn() int {
	result_ := C.C_NSBrowser_LastVisibleColumn(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) IsLoaded() bool {
	result_ := C.C_NSBrowser_IsLoaded(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) IsTitled() bool {
	result_ := C.C_NSBrowser_IsTitled(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetTitled(value bool) {
	C.C_NSBrowser_SetTitled(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) TitleHeight() coregraphics.Float {
	result_ := C.C_NSBrowser_TitleHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) HasHorizontalScroller() bool {
	result_ := C.C_NSBrowser_HasHorizontalScroller(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetHasHorizontalScroller(value bool) {
	C.C_NSBrowser_SetHasHorizontalScroller(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) DoubleAction() *objc.Selector {
	result_ := C.C_NSBrowser_DoubleAction(n.Ptr())
	return objc.MakeSelector(result_)
}

func (n *NSBrowser) SetDoubleAction(value *objc.Selector) {
	C.C_NSBrowser_SetDoubleAction(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowser) SendsActionOnArrowKeys() bool {
	result_ := C.C_NSBrowser_SendsActionOnArrowKeys(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetSendsActionOnArrowKeys(value bool) {
	C.C_NSBrowser_SetSendsActionOnArrowKeys(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) ClickedColumn() int {
	result_ := C.C_NSBrowser_ClickedColumn(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) ClickedRow() int {
	result_ := C.C_NSBrowser_ClickedRow(n.Ptr())
	return int(result_)
}

func (n *NSBrowser) ColumnsAutosaveName() BrowserColumnsAutosaveName {
	result_ := C.C_NSBrowser_ColumnsAutosaveName(n.Ptr())
	return BrowserColumnsAutosaveName(foundation.MakeString(result_).String())
}

func (n *NSBrowser) SetColumnsAutosaveName(value BrowserColumnsAutosaveName) {
	C.C_NSBrowser_SetColumnsAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSBrowser) ColumnResizingType() BrowserColumnResizingType {
	result_ := C.C_NSBrowser_ColumnResizingType(n.Ptr())
	return BrowserColumnResizingType(uint(result_))
}

func (n *NSBrowser) SetColumnResizingType(value BrowserColumnResizingType) {
	C.C_NSBrowser_SetColumnResizingType(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBrowser) PrefersAllColumnUserResizing() bool {
	result_ := C.C_NSBrowser_PrefersAllColumnUserResizing(n.Ptr())
	return bool(result_)
}

func (n *NSBrowser) SetPrefersAllColumnUserResizing(value bool) {
	C.C_NSBrowser_SetPrefersAllColumnUserResizing(n.Ptr(), C.bool(value))
}

func (n *NSBrowser) RowHeight() coregraphics.Float {
	result_ := C.C_NSBrowser_RowHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBrowser) SetRowHeight(value coregraphics.Float) {
	C.C_NSBrowser_SetRowHeight(n.Ptr(), C.double(float64(value)))
}
