package appkit

// #include "table_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TableView interface {
	Control
	ReloadData()
	ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet)
	MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.Object) View
	RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView
	ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View
	BeginUpdates()
	EndUpdates()
	MoveRowAtIndex_ToIndex(oldIndex int, newIndex int)
	InsertRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	RemoveRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	RowForView(view View) int
	ColumnForView(view View) int
	RegisterNib_ForIdentifier(nib Nib, identifier UserInterfaceItemIdentifier)
	IndicatorImageInTableColumn(tableColumn TableColumn) Image
	SetIndicatorImage_InTableColumn(image Image, tableColumn TableColumn)
	AddTableColumn(tableColumn TableColumn)
	RemoveTableColumn(tableColumn TableColumn)
	MoveColumn_ToColumn(oldIndex int, newIndex int)
	ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int
	TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn
	SelectColumnIndexes_ByExtendingSelection(indexes foundation.IndexSet, extend bool)
	DeselectColumn(column int)
	IsColumnSelected(column int) bool
	SelectRowIndexes_ByExtendingSelection(indexes foundation.IndexSet, extend bool)
	DeselectRow(row int)
	IsRowSelected(row int) bool
	SelectAll(sender objc.Object)
	DeselectAll(sender objc.Object)
	EditColumn_Row_WithEvent_Select(column int, row int, event Event, _select bool)
	DidAddRowView_ForRow(rowView TableRowView, row int)
	DidRemoveRowView_ForRow(rowView TableRowView, row int)
	RectOfColumn(column int) foundation.Rect
	RectOfRow(row int) foundation.Rect
	RowsInRect(rect foundation.Rect) foundation.Range
	ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet
	ColumnAtPoint(point foundation.Point) int
	RowAtPoint(point foundation.Point) int
	FrameOfCellAtColumn_Row(column int, row int) foundation.Rect
	SizeLastColumnToFit()
	NoteNumberOfRowsChanged()
	Tile()
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet)
	DrawRow_ClipRect(row int, clipRect foundation.Rect)
	DrawGridInClipRect(clipRect foundation.Rect)
	HighlightSelectionInClipRect(clipRect foundation.Rect)
	DrawBackgroundInClipRect(clipRect foundation.Rect)
	ScrollRowToVisible(row int)
	ScrollColumnToVisible(column int)
	CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IndexSet, mouseDownPoint foundation.Point) bool
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation)
	HideRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	UnhideRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	DataSource() objc.Object
	SetDataSource(value objc.Object)
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	ClickedColumn() int
	ClickedRow() int
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	RowHeight() coregraphics.Float
	SetRowHeight(value coregraphics.Float)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	Style() TableViewStyle
	SetStyle(value TableViewStyle)
	EffectiveStyle() TableViewStyle
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	GridColor() Color
	SetGridColor(value Color)
	GridStyleMask() TableViewGridLineStyle
	SetGridStyleMask(value TableViewGridLineStyle)
	EffectiveRowSizeStyle() TableViewRowSizeStyle
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	TableColumns() []TableColumn
	SelectedColumn() int
	SelectedColumnIndexes() foundation.IndexSet
	NumberOfSelectedColumns() int
	SelectedRow() int
	SelectedRowIndexes() foundation.IndexSet
	NumberOfSelectedRows() int
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	NumberOfColumns() int
	NumberOfRows() int
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)
	EditedColumn() int
	EditedRow() int
	HeaderView() TableHeaderView
	SetHeaderView(value TableHeaderView)
	CornerView() View
	SetCornerView(value View)
	ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle)
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	AutosaveName() TableViewAutosaveName
	SetAutosaveName(value TableViewAutosaveName)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	HighlightedTableColumn() TableColumn
	SetHighlightedTableColumn(value TableColumn)
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.SortDescriptor)
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)
	HiddenRowIndexes() foundation.IndexSet
}

type NSTableView struct {
	NSControl
}

func MakeTableView(ptr unsafe.Pointer) NSTableView {
	return NSTableView{
		NSControl: MakeControl(ptr),
	}
}

func (n NSTableView) InitWithCoder(coder foundation.Coder) NSTableView {
	result_ := C.C_NSTableView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableView(result_)
}

func (n NSTableView) InitWithFrame(frameRect foundation.Rect) NSTableView {
	result_ := C.C_NSTableView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableView(result_)
}

func (n NSTableView) Init() NSTableView {
	result_ := C.C_NSTableView_Init(n.Ptr())
	return MakeTableView(result_)
}

func AllocTableView() NSTableView {
	result_ := C.C_NSTableView_AllocTableView()
	return MakeTableView(result_)
}

func NewTableView() NSTableView {
	result_ := C.C_NSTableView_NewTableView()
	return MakeTableView(result_)
}

func (n NSTableView) Autorelease() NSTableView {
	result_ := C.C_NSTableView_Autorelease(n.Ptr())
	return MakeTableView(result_)
}

func (n NSTableView) Retain() NSTableView {
	result_ := C.C_NSTableView_Retain(n.Ptr())
	return MakeTableView(result_)
}

func (n NSTableView) ReloadData() {
	C.C_NSTableView_ReloadData(n.Ptr())
}

func (n NSTableView) ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet) {
	C.C_NSTableView_ReloadDataForRowIndexes_ColumnIndexes(n.Ptr(), objc.ExtractPtr(rowIndexes), objc.ExtractPtr(columnIndexes))
}

func (n NSTableView) MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.Object) View {
	result_ := C.C_NSTableView_MakeViewWithIdentifier_Owner(n.Ptr(), foundation.NewString(string(identifier)).Ptr(), objc.ExtractPtr(owner))
	return MakeView(result_)
}

func (n NSTableView) RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	result_ := C.C_NSTableView_RowViewAtRow_MakeIfNecessary(n.Ptr(), C.int(row), C.bool(makeIfNecessary))
	return MakeTableRowView(result_)
}

func (n NSTableView) ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	result_ := C.C_NSTableView_ViewAtColumn_Row_MakeIfNecessary(n.Ptr(), C.int(column), C.int(row), C.bool(makeIfNecessary))
	return MakeView(result_)
}

func (n NSTableView) BeginUpdates() {
	C.C_NSTableView_BeginUpdates(n.Ptr())
}

func (n NSTableView) EndUpdates() {
	C.C_NSTableView_EndUpdates(n.Ptr())
}

func (n NSTableView) MoveRowAtIndex_ToIndex(oldIndex int, newIndex int) {
	C.C_NSTableView_MoveRowAtIndex_ToIndex(n.Ptr(), C.int(oldIndex), C.int(newIndex))
}

func (n NSTableView) InsertRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	C.C_NSTableView_InsertRowsAtIndexes_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), C.uint(uint(animationOptions)))
}

func (n NSTableView) RemoveRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	C.C_NSTableView_RemoveRowsAtIndexes_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), C.uint(uint(animationOptions)))
}

func (n NSTableView) RowForView(view View) int {
	result_ := C.C_NSTableView_RowForView(n.Ptr(), objc.ExtractPtr(view))
	return int(result_)
}

func (n NSTableView) ColumnForView(view View) int {
	result_ := C.C_NSTableView_ColumnForView(n.Ptr(), objc.ExtractPtr(view))
	return int(result_)
}

func (n NSTableView) RegisterNib_ForIdentifier(nib Nib, identifier UserInterfaceItemIdentifier) {
	C.C_NSTableView_RegisterNib_ForIdentifier(n.Ptr(), objc.ExtractPtr(nib), foundation.NewString(string(identifier)).Ptr())
}

func (n NSTableView) IndicatorImageInTableColumn(tableColumn TableColumn) Image {
	result_ := C.C_NSTableView_IndicatorImageInTableColumn(n.Ptr(), objc.ExtractPtr(tableColumn))
	return MakeImage(result_)
}

func (n NSTableView) SetIndicatorImage_InTableColumn(image Image, tableColumn TableColumn) {
	C.C_NSTableView_SetIndicatorImage_InTableColumn(n.Ptr(), objc.ExtractPtr(image), objc.ExtractPtr(tableColumn))
}

func (n NSTableView) AddTableColumn(tableColumn TableColumn) {
	C.C_NSTableView_AddTableColumn(n.Ptr(), objc.ExtractPtr(tableColumn))
}

func (n NSTableView) RemoveTableColumn(tableColumn TableColumn) {
	C.C_NSTableView_RemoveTableColumn(n.Ptr(), objc.ExtractPtr(tableColumn))
}

func (n NSTableView) MoveColumn_ToColumn(oldIndex int, newIndex int) {
	C.C_NSTableView_MoveColumn_ToColumn(n.Ptr(), C.int(oldIndex), C.int(newIndex))
}

func (n NSTableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	result_ := C.C_NSTableView_ColumnWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return int(result_)
}

func (n NSTableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	result_ := C.C_NSTableView_TableColumnWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeTableColumn(result_)
}

func (n NSTableView) SelectColumnIndexes_ByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	C.C_NSTableView_SelectColumnIndexes_ByExtendingSelection(n.Ptr(), objc.ExtractPtr(indexes), C.bool(extend))
}

func (n NSTableView) DeselectColumn(column int) {
	C.C_NSTableView_DeselectColumn(n.Ptr(), C.int(column))
}

func (n NSTableView) IsColumnSelected(column int) bool {
	result_ := C.C_NSTableView_IsColumnSelected(n.Ptr(), C.int(column))
	return bool(result_)
}

func (n NSTableView) SelectRowIndexes_ByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	C.C_NSTableView_SelectRowIndexes_ByExtendingSelection(n.Ptr(), objc.ExtractPtr(indexes), C.bool(extend))
}

func (n NSTableView) DeselectRow(row int) {
	C.C_NSTableView_DeselectRow(n.Ptr(), C.int(row))
}

func (n NSTableView) IsRowSelected(row int) bool {
	result_ := C.C_NSTableView_IsRowSelected(n.Ptr(), C.int(row))
	return bool(result_)
}

func (n NSTableView) SelectAll(sender objc.Object) {
	C.C_NSTableView_SelectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTableView) DeselectAll(sender objc.Object) {
	C.C_NSTableView_DeselectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTableView) EditColumn_Row_WithEvent_Select(column int, row int, event Event, _select bool) {
	C.C_NSTableView_EditColumn_Row_WithEvent_Select(n.Ptr(), C.int(column), C.int(row), objc.ExtractPtr(event), C.bool(_select))
}

func (n NSTableView) DidAddRowView_ForRow(rowView TableRowView, row int) {
	C.C_NSTableView_DidAddRowView_ForRow(n.Ptr(), objc.ExtractPtr(rowView), C.int(row))
}

func (n NSTableView) DidRemoveRowView_ForRow(rowView TableRowView, row int) {
	C.C_NSTableView_DidRemoveRowView_ForRow(n.Ptr(), objc.ExtractPtr(rowView), C.int(row))
}

func (n NSTableView) RectOfColumn(column int) foundation.Rect {
	result_ := C.C_NSTableView_RectOfColumn(n.Ptr(), C.int(column))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTableView) RectOfRow(row int) foundation.Rect {
	result_ := C.C_NSTableView_RectOfRow(n.Ptr(), C.int(row))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTableView) RowsInRect(rect foundation.Rect) foundation.Range {
	result_ := C.C_NSTableView_RowsInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	result_ := C.C_NSTableView_ColumnIndexesInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.MakeIndexSet(result_)
}

func (n NSTableView) ColumnAtPoint(point foundation.Point) int {
	result_ := C.C_NSTableView_ColumnAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return int(result_)
}

func (n NSTableView) RowAtPoint(point foundation.Point) int {
	result_ := C.C_NSTableView_RowAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return int(result_)
}

func (n NSTableView) FrameOfCellAtColumn_Row(column int, row int) foundation.Rect {
	result_ := C.C_NSTableView_FrameOfCellAtColumn_Row(n.Ptr(), C.int(column), C.int(row))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTableView) SizeLastColumnToFit() {
	C.C_NSTableView_SizeLastColumnToFit(n.Ptr())
}

func (n NSTableView) NoteNumberOfRowsChanged() {
	C.C_NSTableView_NoteNumberOfRowsChanged(n.Ptr())
}

func (n NSTableView) Tile() {
	C.C_NSTableView_Tile(n.Ptr())
}

func (n NSTableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet) {
	C.C_NSTableView_NoteHeightOfRowsWithIndexesChanged(n.Ptr(), objc.ExtractPtr(indexSet))
}

func (n NSTableView) DrawRow_ClipRect(row int, clipRect foundation.Rect) {
	C.C_NSTableView_DrawRow_ClipRect(n.Ptr(), C.int(row), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))))
}

func (n NSTableView) DrawGridInClipRect(clipRect foundation.Rect) {
	C.C_NSTableView_DrawGridInClipRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))))
}

func (n NSTableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	C.C_NSTableView_HighlightSelectionInClipRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))))
}

func (n NSTableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	C.C_NSTableView_DrawBackgroundInClipRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))))
}

func (n NSTableView) ScrollRowToVisible(row int) {
	C.C_NSTableView_ScrollRowToVisible(n.Ptr(), C.int(row))
}

func (n NSTableView) ScrollColumnToVisible(column int) {
	C.C_NSTableView_ScrollColumnToVisible(n.Ptr(), C.int(column))
}

func (n NSTableView) CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IndexSet, mouseDownPoint foundation.Point) bool {
	result_ := C.C_NSTableView_CanDragRowsWithIndexes_AtPoint(n.Ptr(), objc.ExtractPtr(rowIndexes), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(mouseDownPoint))))
	return bool(result_)
}

func (n NSTableView) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	C.C_NSTableView_SetDraggingSourceOperationMask_ForLocal(n.Ptr(), C.uint(uint(mask)), C.bool(isLocal))
}

func (n NSTableView) SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation) {
	C.C_NSTableView_SetDropRow_DropOperation(n.Ptr(), C.int(row), C.uint(uint(dropOperation)))
}

func (n NSTableView) HideRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	C.C_NSTableView_HideRowsAtIndexes_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), C.uint(uint(rowAnimation)))
}

func (n NSTableView) UnhideRowsAtIndexes_WithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	C.C_NSTableView_UnhideRowsAtIndexes_WithAnimation(n.Ptr(), objc.ExtractPtr(indexes), C.uint(uint(rowAnimation)))
}

func (n NSTableView) DataSource() objc.Object {
	result_ := C.C_NSTableView_DataSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTableView) SetDataSource(value objc.Object) {
	C.C_NSTableView_SetDataSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) UsesStaticContents() bool {
	result_ := C.C_NSTableView_UsesStaticContents(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetUsesStaticContents(value bool) {
	C.C_NSTableView_SetUsesStaticContents(n.Ptr(), C.bool(value))
}

func (n NSTableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	result_ := C.C_NSTableView_RegisteredNibsByIdentifier(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[UserInterfaceItemIdentifier]Nib)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[UserInterfaceItemIdentifier(foundation.MakeString(k).String())] = MakeNib(v)
	}
	return goResult_
}

func (n NSTableView) DoubleAction() objc.Selector {
	result_ := C.C_NSTableView_DoubleAction(n.Ptr())
	return objc.Selector(result_)
}

func (n NSTableView) SetDoubleAction(value objc.Selector) {
	C.C_NSTableView_SetDoubleAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSTableView) ClickedColumn() int {
	result_ := C.C_NSTableView_ClickedColumn(n.Ptr())
	return int(result_)
}

func (n NSTableView) ClickedRow() int {
	result_ := C.C_NSTableView_ClickedRow(n.Ptr())
	return int(result_)
}

func (n NSTableView) AllowsColumnReordering() bool {
	result_ := C.C_NSTableView_AllowsColumnReordering(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsColumnReordering(value bool) {
	C.C_NSTableView_SetAllowsColumnReordering(n.Ptr(), C.bool(value))
}

func (n NSTableView) AllowsColumnResizing() bool {
	result_ := C.C_NSTableView_AllowsColumnResizing(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsColumnResizing(value bool) {
	C.C_NSTableView_SetAllowsColumnResizing(n.Ptr(), C.bool(value))
}

func (n NSTableView) AllowsMultipleSelection() bool {
	result_ := C.C_NSTableView_AllowsMultipleSelection(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsMultipleSelection(value bool) {
	C.C_NSTableView_SetAllowsMultipleSelection(n.Ptr(), C.bool(value))
}

func (n NSTableView) AllowsEmptySelection() bool {
	result_ := C.C_NSTableView_AllowsEmptySelection(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsEmptySelection(value bool) {
	C.C_NSTableView_SetAllowsEmptySelection(n.Ptr(), C.bool(value))
}

func (n NSTableView) AllowsColumnSelection() bool {
	result_ := C.C_NSTableView_AllowsColumnSelection(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsColumnSelection(value bool) {
	C.C_NSTableView_SetAllowsColumnSelection(n.Ptr(), C.bool(value))
}

func (n NSTableView) UsesAutomaticRowHeights() bool {
	result_ := C.C_NSTableView_UsesAutomaticRowHeights(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetUsesAutomaticRowHeights(value bool) {
	C.C_NSTableView_SetUsesAutomaticRowHeights(n.Ptr(), C.bool(value))
}

func (n NSTableView) IntercellSpacing() foundation.Size {
	result_ := C.C_NSTableView_IntercellSpacing(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSTableView) SetIntercellSpacing(value foundation.Size) {
	C.C_NSTableView_SetIntercellSpacing(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSTableView) RowHeight() coregraphics.Float {
	result_ := C.C_NSTableView_RowHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTableView) SetRowHeight(value coregraphics.Float) {
	C.C_NSTableView_SetRowHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSTableView) BackgroundColor() Color {
	result_ := C.C_NSTableView_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTableView) SetBackgroundColor(value Color) {
	C.C_NSTableView_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) UsesAlternatingRowBackgroundColors() bool {
	result_ := C.C_NSTableView_UsesAlternatingRowBackgroundColors(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	C.C_NSTableView_SetUsesAlternatingRowBackgroundColors(n.Ptr(), C.bool(value))
}

func (n NSTableView) Style() TableViewStyle {
	result_ := C.C_NSTableView_Style(n.Ptr())
	return TableViewStyle(int(result_))
}

func (n NSTableView) SetStyle(value TableViewStyle) {
	C.C_NSTableView_SetStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableView) EffectiveStyle() TableViewStyle {
	result_ := C.C_NSTableView_EffectiveStyle(n.Ptr())
	return TableViewStyle(int(result_))
}

func (n NSTableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	result_ := C.C_NSTableView_SelectionHighlightStyle(n.Ptr())
	return TableViewSelectionHighlightStyle(int(result_))
}

func (n NSTableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	C.C_NSTableView_SetSelectionHighlightStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableView) GridColor() Color {
	result_ := C.C_NSTableView_GridColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTableView) SetGridColor(value Color) {
	C.C_NSTableView_SetGridColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) GridStyleMask() TableViewGridLineStyle {
	result_ := C.C_NSTableView_GridStyleMask(n.Ptr())
	return TableViewGridLineStyle(uint(result_))
}

func (n NSTableView) SetGridStyleMask(value TableViewGridLineStyle) {
	C.C_NSTableView_SetGridStyleMask(n.Ptr(), C.uint(uint(value)))
}

func (n NSTableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	result_ := C.C_NSTableView_EffectiveRowSizeStyle(n.Ptr())
	return TableViewRowSizeStyle(int(result_))
}

func (n NSTableView) RowSizeStyle() TableViewRowSizeStyle {
	result_ := C.C_NSTableView_RowSizeStyle(n.Ptr())
	return TableViewRowSizeStyle(int(result_))
}

func (n NSTableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	C.C_NSTableView_SetRowSizeStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableView) TableColumns() []TableColumn {
	result_ := C.C_NSTableView_TableColumns(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TableColumn, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTableColumn(r)
	}
	return goResult_
}

func (n NSTableView) SelectedColumn() int {
	result_ := C.C_NSTableView_SelectedColumn(n.Ptr())
	return int(result_)
}

func (n NSTableView) SelectedColumnIndexes() foundation.IndexSet {
	result_ := C.C_NSTableView_SelectedColumnIndexes(n.Ptr())
	return foundation.MakeIndexSet(result_)
}

func (n NSTableView) NumberOfSelectedColumns() int {
	result_ := C.C_NSTableView_NumberOfSelectedColumns(n.Ptr())
	return int(result_)
}

func (n NSTableView) SelectedRow() int {
	result_ := C.C_NSTableView_SelectedRow(n.Ptr())
	return int(result_)
}

func (n NSTableView) SelectedRowIndexes() foundation.IndexSet {
	result_ := C.C_NSTableView_SelectedRowIndexes(n.Ptr())
	return foundation.MakeIndexSet(result_)
}

func (n NSTableView) NumberOfSelectedRows() int {
	result_ := C.C_NSTableView_NumberOfSelectedRows(n.Ptr())
	return int(result_)
}

func (n NSTableView) AllowsTypeSelect() bool {
	result_ := C.C_NSTableView_AllowsTypeSelect(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAllowsTypeSelect(value bool) {
	C.C_NSTableView_SetAllowsTypeSelect(n.Ptr(), C.bool(value))
}

func (n NSTableView) NumberOfColumns() int {
	result_ := C.C_NSTableView_NumberOfColumns(n.Ptr())
	return int(result_)
}

func (n NSTableView) NumberOfRows() int {
	result_ := C.C_NSTableView_NumberOfRows(n.Ptr())
	return int(result_)
}

func (n NSTableView) FloatsGroupRows() bool {
	result_ := C.C_NSTableView_FloatsGroupRows(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetFloatsGroupRows(value bool) {
	C.C_NSTableView_SetFloatsGroupRows(n.Ptr(), C.bool(value))
}

func (n NSTableView) EditedColumn() int {
	result_ := C.C_NSTableView_EditedColumn(n.Ptr())
	return int(result_)
}

func (n NSTableView) EditedRow() int {
	result_ := C.C_NSTableView_EditedRow(n.Ptr())
	return int(result_)
}

func (n NSTableView) HeaderView() TableHeaderView {
	result_ := C.C_NSTableView_HeaderView(n.Ptr())
	return MakeTableHeaderView(result_)
}

func (n NSTableView) SetHeaderView(value TableHeaderView) {
	C.C_NSTableView_SetHeaderView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) CornerView() View {
	result_ := C.C_NSTableView_CornerView(n.Ptr())
	return MakeView(result_)
}

func (n NSTableView) SetCornerView(value View) {
	C.C_NSTableView_SetCornerView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	result_ := C.C_NSTableView_ColumnAutoresizingStyle(n.Ptr())
	return TableViewColumnAutoresizingStyle(uint(result_))
}

func (n NSTableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	C.C_NSTableView_SetColumnAutoresizingStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSTableView) AutosaveTableColumns() bool {
	result_ := C.C_NSTableView_AutosaveTableColumns(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetAutosaveTableColumns(value bool) {
	C.C_NSTableView_SetAutosaveTableColumns(n.Ptr(), C.bool(value))
}

func (n NSTableView) AutosaveName() TableViewAutosaveName {
	result_ := C.C_NSTableView_AutosaveName(n.Ptr())
	return TableViewAutosaveName(foundation.MakeString(result_).String())
}

func (n NSTableView) SetAutosaveName(value TableViewAutosaveName) {
	C.C_NSTableView_SetAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSTableView) Delegate() objc.Object {
	result_ := C.C_NSTableView_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTableView) SetDelegate(value objc.Object) {
	C.C_NSTableView_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) HighlightedTableColumn() TableColumn {
	result_ := C.C_NSTableView_HighlightedTableColumn(n.Ptr())
	return MakeTableColumn(result_)
}

func (n NSTableView) SetHighlightedTableColumn(value TableColumn) {
	C.C_NSTableView_SetHighlightedTableColumn(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableView) VerticalMotionCanBeginDrag() bool {
	result_ := C.C_NSTableView_VerticalMotionCanBeginDrag(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetVerticalMotionCanBeginDrag(value bool) {
	C.C_NSTableView_SetVerticalMotionCanBeginDrag(n.Ptr(), C.bool(value))
}

func (n NSTableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	result_ := C.C_NSTableView_DraggingDestinationFeedbackStyle(n.Ptr())
	return TableViewDraggingDestinationFeedbackStyle(int(result_))
}

func (n NSTableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	C.C_NSTableView_SetDraggingDestinationFeedbackStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableView) SortDescriptors() []foundation.SortDescriptor {
	result_ := C.C_NSTableView_SortDescriptors(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.SortDescriptor, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeSortDescriptor(r)
	}
	return goResult_
}

func (n NSTableView) SetSortDescriptors(value []foundation.SortDescriptor) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTableView_SetSortDescriptors(n.Ptr(), cValue)
}

func (n NSTableView) RowActionsVisible() bool {
	result_ := C.C_NSTableView_RowActionsVisible(n.Ptr())
	return bool(result_)
}

func (n NSTableView) SetRowActionsVisible(value bool) {
	C.C_NSTableView_SetRowActionsVisible(n.Ptr(), C.bool(value))
}

func (n NSTableView) HiddenRowIndexes() foundation.IndexSet {
	result_ := C.C_NSTableView_HiddenRowIndexes(n.Ptr())
	return foundation.MakeIndexSet(result_)
}

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

func (delegate *TableViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTableViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export tableViewDelegate_TableView_ViewForTableColumn_Row
func tableViewDelegate_TableView_ViewForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ViewForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_RowViewForRow
func tableViewDelegate_TableView_RowViewForRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_RowViewForRow(MakeTableView(tableView), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_DidAddRowView_ForRow
func tableViewDelegate_TableView_DidAddRowView_ForRow(hp C.uintptr_t, tableView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_DidAddRowView_ForRow(MakeTableView(tableView), MakeTableRowView(rowView), int(row))
}

//export tableViewDelegate_TableView_DidRemoveRowView_ForRow
func tableViewDelegate_TableView_DidRemoveRowView_ForRow(hp C.uintptr_t, tableView unsafe.Pointer, rowView unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_DidRemoveRowView_ForRow(MakeTableView(tableView), MakeTableRowView(rowView), int(row))
}

//export tableViewDelegate_TableView_IsGroupRow
func tableViewDelegate_TableView_IsGroupRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_IsGroupRow(MakeTableView(tableView), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_WillDisplayCell_ForTableColumn_Row
func tableViewDelegate_TableView_WillDisplayCell_ForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_WillDisplayCell_ForTableColumn_Row(MakeTableView(tableView), objc.MakeObject(cell), MakeTableColumn(tableColumn), int(row))
}

//export tableViewDelegate_TableView_DataCellForTableColumn_Row
func tableViewDelegate_TableView_DataCellForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_DataCellForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_ShouldShowCellExpansionForTableColumn_Row
func tableViewDelegate_TableView_ShouldShowCellExpansionForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldShowCellExpansionForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_ShouldEditTableColumn_Row
func tableViewDelegate_TableView_ShouldEditTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldEditTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_HeightOfRow
func tableViewDelegate_TableView_HeightOfRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_HeightOfRow(MakeTableView(tableView), int(row))
	return C.double(float64(result))
}

//export tableViewDelegate_TableView_SizeToFitWidthOfColumn
func tableViewDelegate_TableView_SizeToFitWidthOfColumn(hp C.uintptr_t, tableView unsafe.Pointer, column C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_SizeToFitWidthOfColumn(MakeTableView(tableView), int(column))
	return C.double(float64(result))
}

//export tableViewDelegate_SelectionShouldChangeInTableView
func tableViewDelegate_SelectionShouldChangeInTableView(hp C.uintptr_t, tableView unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.SelectionShouldChangeInTableView(MakeTableView(tableView))
	return C.bool(result)
}

//export tableViewDelegate_TableView_ShouldSelectRow
func tableViewDelegate_TableView_ShouldSelectRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldSelectRow(MakeTableView(tableView), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_SelectionIndexesForProposedSelection
func tableViewDelegate_TableView_SelectionIndexesForProposedSelection(hp C.uintptr_t, tableView unsafe.Pointer, proposedSelectionIndexes unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_SelectionIndexesForProposedSelection(MakeTableView(tableView), foundation.MakeIndexSet(proposedSelectionIndexes))
	return objc.ExtractPtr(result)
}

//export tableViewDelegate_TableView_ShouldSelectTableColumn
func tableViewDelegate_TableView_ShouldSelectTableColumn(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldSelectTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
	return C.bool(result)
}

//export tableViewDelegate_TableViewSelectionIsChanging
func tableViewDelegate_TableViewSelectionIsChanging(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableViewSelectionIsChanging(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableViewSelectionDidChange
func tableViewDelegate_TableViewSelectionDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableViewSelectionDidChange(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString
func tableViewDelegate_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(hp C.uintptr_t, tableView unsafe.Pointer, event unsafe.Pointer, searchString unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(MakeTableView(tableView), MakeEvent(event), foundation.MakeString(searchString).String())
	return C.bool(result)
}

//export tableViewDelegate_TableView_TypeSelectStringForTableColumn_Row
func tableViewDelegate_TableView_TypeSelectStringForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_TypeSelectStringForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return foundation.NewString(result).Ptr()
}

//export tableViewDelegate_TableView_NextTypeSelectMatchFromRow_ToRow_ForString
func tableViewDelegate_TableView_NextTypeSelectMatchFromRow_ToRow_ForString(hp C.uintptr_t, tableView unsafe.Pointer, startRow C.int, endRow C.int, searchString unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_NextTypeSelectMatchFromRow_ToRow_ForString(MakeTableView(tableView), int(startRow), int(endRow), foundation.MakeString(searchString).String())
	return C.int(result)
}

//export tableViewDelegate_TableView_ShouldReorderColumn_ToColumn
func tableViewDelegate_TableView_ShouldReorderColumn_ToColumn(hp C.uintptr_t, tableView unsafe.Pointer, columnIndex C.int, newColumnIndex C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldReorderColumn_ToColumn(MakeTableView(tableView), int(columnIndex), int(newColumnIndex))
	return C.bool(result)
}

//export tableViewDelegate_TableView_DidDragTableColumn
func tableViewDelegate_TableView_DidDragTableColumn(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_DidDragTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableViewColumnDidMove
func tableViewDelegate_TableViewColumnDidMove(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableViewColumnDidMove(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableViewColumnDidResize
func tableViewDelegate_TableViewColumnDidResize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableViewColumnDidResize(foundation.MakeNotification(notification))
}

//export tableViewDelegate_TableView_DidClickTableColumn
func tableViewDelegate_TableView_DidClickTableColumn(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_DidClickTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableView_MouseDownInHeaderOfTableColumn
func tableViewDelegate_TableView_MouseDownInHeaderOfTableColumn(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.TableView_MouseDownInHeaderOfTableColumn(MakeTableView(tableView), MakeTableColumn(tableColumn))
}

//export tableViewDelegate_TableView_ShouldTrackCell_ForTableColumn_Row
func tableViewDelegate_TableView_ShouldTrackCell_ForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, cell unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_ShouldTrackCell_ForTableColumn_Row(MakeTableView(tableView), MakeCell(cell), MakeTableColumn(tableColumn), int(row))
	return C.bool(result)
}

//export tableViewDelegate_TableView_RowActionsForRow_Edge
func tableViewDelegate_TableView_RowActionsForRow_Edge(hp C.uintptr_t, tableView unsafe.Pointer, row C.int, edge C.int) C.Array {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.TableView_RowActionsForRow_Edge(MakeTableView(tableView), int(row), TableRowActionEdge(int(edge)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export tableViewDelegate_Control_IsValidObject
func tableViewDelegate_Control_IsValidObject(hp C.uintptr_t, control unsafe.Pointer, obj unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.Control_IsValidObject(MakeControl(control), objc.MakeObject(obj))
	return C.bool(result)
}

//export tableViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription
func tableViewDelegate_Control_DidFailToValidatePartialString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.Control_DidFailToValidatePartialString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
}

//export tableViewDelegate_Control_DidFailToFormatString_ErrorDescription
func tableViewDelegate_Control_DidFailToFormatString_ErrorDescription(hp C.uintptr_t, control unsafe.Pointer, _string unsafe.Pointer, error unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.Control_DidFailToFormatString_ErrorDescription(MakeControl(control), foundation.MakeString(_string).String(), foundation.MakeString(error).String())
	return C.bool(result)
}

//export tableViewDelegate_Control_TextShouldBeginEditing
func tableViewDelegate_Control_TextShouldBeginEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.Control_TextShouldBeginEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export tableViewDelegate_Control_TextShouldEndEditing
func tableViewDelegate_Control_TextShouldEndEditing(hp C.uintptr_t, control unsafe.Pointer, fieldEditor unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.Control_TextShouldEndEditing(MakeControl(control), MakeText(fieldEditor))
	return C.bool(result)
}

//export tableViewDelegate_Control_TextView_DoCommandBySelector
func tableViewDelegate_Control_TextView_DoCommandBySelector(hp C.uintptr_t, control unsafe.Pointer, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	result := delegate.Control_TextView_DoCommandBySelector(MakeControl(control), MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export tableViewDelegate_ControlTextDidBeginEditing
func tableViewDelegate_ControlTextDidBeginEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.ControlTextDidBeginEditing(foundation.MakeNotification(obj))
}

//export tableViewDelegate_ControlTextDidChange
func tableViewDelegate_ControlTextDidChange(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.ControlTextDidChange(foundation.MakeNotification(obj))
}

//export tableViewDelegate_ControlTextDidEndEditing
func tableViewDelegate_ControlTextDidEndEditing(hp C.uintptr_t, obj unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
	delegate.ControlTextDidEndEditing(foundation.MakeNotification(obj))
}

//export TableViewDelegate_RespondsTo
func TableViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TableViewDelegate)
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

type TableViewDataSource interface {
	HasNumberOfRowsInTableView() bool
	NumberOfRowsInTableView(tableView TableView) int
	HasTableView_ObjectValueForTableColumn_Row() bool
	TableView_ObjectValueForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) objc.Object
	HasTableView_SetObjectValue_ForTableColumn_Row() bool
	TableView_SetObjectValue_ForTableColumn_Row(tableView TableView, object objc.Object, tableColumn TableColumn, row int)
	HasTableView_PasteboardWriterForRow() bool
	TableView_PasteboardWriterForRow(tableView TableView, row int) objc.Object
	HasTableView_AcceptDrop_Row_DropOperation() bool
	TableView_AcceptDrop_Row_DropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) bool
	HasTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool
	TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) DragOperation
	HasTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool
	TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	HasTableView_UpdateDraggingItemsForDrag() bool
	TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo objc.Object)
	HasTableView_DraggingSession_EndedAtPoint_Operation() bool
	TableView_DraggingSession_EndedAtPoint_Operation(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	HasTableView_SortDescriptorsDidChange() bool
	TableView_SortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor)
}

func TableViewDataSourceToObjc(protocol TableViewDataSource) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapTableViewDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export tableViewDataSource_NumberOfRowsInTableView
func tableViewDataSource_NumberOfRowsInTableView(hp C.uintptr_t, tableView unsafe.Pointer) C.int {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	result := protocol.NumberOfRowsInTableView(MakeTableView(tableView))
	return C.int(result)
}

//export tableViewDataSource_TableView_ObjectValueForTableColumn_Row
func tableViewDataSource_TableView_ObjectValueForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	result := protocol.TableView_ObjectValueForTableColumn_Row(MakeTableView(tableView), MakeTableColumn(tableColumn), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row
func tableViewDataSource_TableView_SetObjectValue_ForTableColumn_Row(hp C.uintptr_t, tableView unsafe.Pointer, object unsafe.Pointer, tableColumn unsafe.Pointer, row C.int) {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	protocol.TableView_SetObjectValue_ForTableColumn_Row(MakeTableView(tableView), objc.MakeObject(object), MakeTableColumn(tableColumn), int(row))
}

//export tableViewDataSource_TableView_PasteboardWriterForRow
func tableViewDataSource_TableView_PasteboardWriterForRow(hp C.uintptr_t, tableView unsafe.Pointer, row C.int) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	result := protocol.TableView_PasteboardWriterForRow(MakeTableView(tableView), int(row))
	return objc.ExtractPtr(result)
}

//export tableViewDataSource_TableView_AcceptDrop_Row_DropOperation
func tableViewDataSource_TableView_AcceptDrop_Row_DropOperation(hp C.uintptr_t, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.bool {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	result := protocol.TableView_AcceptDrop_Row_DropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.bool(result)
}

//export tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation
func tableViewDataSource_TableView_ValidateDrop_ProposedRow_ProposedDropOperation(hp C.uintptr_t, tableView unsafe.Pointer, info unsafe.Pointer, row C.int, dropOperation C.uint) C.uint {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	result := protocol.TableView_ValidateDrop_ProposedRow_ProposedDropOperation(MakeTableView(tableView), objc.MakeObject(info), int(row), TableViewDropOperation(uint(dropOperation)))
	return C.uint(uint(result))
}

//export tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes
func tableViewDataSource_TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(hp C.uintptr_t, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, rowIndexes unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	protocol.TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(rowIndexes))
}

//export tableViewDataSource_TableView_UpdateDraggingItemsForDrag
func tableViewDataSource_TableView_UpdateDraggingItemsForDrag(hp C.uintptr_t, tableView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	protocol.TableView_UpdateDraggingItemsForDrag(MakeTableView(tableView), objc.MakeObject(draggingInfo))
}

//export tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation
func tableViewDataSource_TableView_DraggingSession_EndedAtPoint_Operation(hp C.uintptr_t, tableView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	protocol.TableView_DraggingSession_EndedAtPoint_Operation(MakeTableView(tableView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export tableViewDataSource_TableView_SortDescriptorsDidChange
func tableViewDataSource_TableView_SortDescriptorsDidChange(hp C.uintptr_t, tableView unsafe.Pointer, oldDescriptors C.Array) {
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	if oldDescriptors.len > 0 {
		defer C.free(oldDescriptors.data)
	}
	oldDescriptorsSlice := unsafe.Slice((*unsafe.Pointer)(oldDescriptors.data), int(oldDescriptors.len))
	var goOldDescriptors = make([]foundation.SortDescriptor, len(oldDescriptorsSlice))
	for idx, r := range oldDescriptorsSlice {
		goOldDescriptors[idx] = foundation.MakeSortDescriptor(r)
	}
	protocol.TableView_SortDescriptorsDidChange(MakeTableView(tableView), goOldDescriptors)
}

//export TableViewDataSource_RespondsTo
func TableViewDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(TableViewDataSource)
	_ = protocol
	switch selName {
	case "numberOfRowsInTableView:":
		return protocol.HasNumberOfRowsInTableView()
	case "tableView:objectValueForTableColumn:row:":
		return protocol.HasTableView_ObjectValueForTableColumn_Row()
	case "tableView:setObjectValue:forTableColumn:row:":
		return protocol.HasTableView_SetObjectValue_ForTableColumn_Row()
	case "tableView:pasteboardWriterForRow:":
		return protocol.HasTableView_PasteboardWriterForRow()
	case "tableView:acceptDrop:row:dropOperation:":
		return protocol.HasTableView_AcceptDrop_Row_DropOperation()
	case "tableView:validateDrop:proposedRow:proposedDropOperation:":
		return protocol.HasTableView_ValidateDrop_ProposedRow_ProposedDropOperation()
	case "tableView:draggingSession:willBeginAtPoint:forRowIndexes:":
		return protocol.HasTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes()
	case "tableView:updateDraggingItemsForDrag:":
		return protocol.HasTableView_UpdateDraggingItemsForDrag()
	case "tableView:draggingSession:endedAtPoint:operation:":
		return protocol.HasTableView_DraggingSession_EndedAtPoint_Operation()
	case "tableView:sortDescriptorsDidChange:":
		return protocol.HasTableView_SortDescriptorsDidChange()
	default:
		return false
	}
}
