package appkit

// #include "table_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func AllocTableView() NSTableView {
	return MakeTableView(C.C_TableView_Alloc())
}

func (n NSTableView) InitWithCoder(coder foundation.Coder) TableView {
	result_ := C.C_NSTableView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableView(result_)
}

func (n NSTableView) InitWithFrame(frameRect foundation.Rect) TableView {
	result_ := C.C_NSTableView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableView(result_)
}

func (n NSTableView) Init() TableView {
	result_ := C.C_NSTableView_Init(n.Ptr())
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
	result_KeySlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.key_data))[:result_.len:result_.len]
	result_ValueSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.value_data))[:result_.len:result_.len]
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
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
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
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
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
