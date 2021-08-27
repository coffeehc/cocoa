package appkit

// #include "collection_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionView interface {
	View
	MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IndexPath) CollectionViewItem
	RegisterNib_ForItemWithIdentifier(nib Nib, identifier UserInterfaceItemIdentifier)
	MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IndexPath) View
	RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(nib Nib, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier)
	ReloadData()
	ReloadSections(sections foundation.IndexSet)
	ReloadItemsAtIndexPaths(indexPaths foundation.Set)
	NumberOfItemsInSection(section int) int
	InsertItemsAtIndexPaths(indexPaths foundation.Set)
	MoveItemAtIndexPath_ToIndexPath(indexPath foundation.IndexPath, newIndexPath foundation.IndexPath)
	DeleteItemsAtIndexPaths(indexPaths foundation.Set)
	InsertSections(sections foundation.IndexSet)
	MoveSection_ToSection(section int, newSection int)
	DeleteSections(sections foundation.IndexSet)
	ToggleSectionCollapse(sender objc.Object)
	SelectAll(sender objc.Object)
	DeselectAll(sender objc.Object)
	SelectItemsAtIndexPaths_ScrollPosition(indexPaths foundation.Set, scrollPosition CollectionViewScrollPosition)
	DeselectItemsAtIndexPaths(indexPaths foundation.Set)
	VisibleItems() []CollectionViewItem
	IndexPathsForVisibleItems() foundation.Set
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathForItem(item CollectionViewItem) foundation.IndexPath
	IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath
	ItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewItem
	SupplementaryViewForElementKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) View
	ScrollToItemsAtIndexPaths_ScrollPosition(indexPaths foundation.Set, scrollPosition CollectionViewScrollPosition)
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes
	ItemAtIndex(index uint) CollectionViewItem
	FrameForItemAtIndex(index uint) foundation.Rect
	FrameForItemAtIndex_WithNumberOfItems(index uint, numberOfItems uint) foundation.Rect
	SetDraggingSourceOperationMask_ForLocal(dragOperationMask DragOperation, localDestination bool)
	DataSource() objc.Object
	SetDataSource(value objc.Object)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	Content() []objc.Object
	SetContent(value []objc.Object)
	BackgroundView() View
	SetBackgroundView(value View)
	BackgroundColors() []Color
	SetBackgroundColors(value []Color)
	BackgroundViewScrollsWithContent() bool
	SetBackgroundViewScrollsWithContent(value bool)
	CollectionViewLayout() CollectionViewLayout
	SetCollectionViewLayout(value CollectionViewLayout)
	PrefetchDataSource() objc.Object
	SetPrefetchDataSource(value objc.Object)
	NumberOfSections() int
	IsSelectable() bool
	SetSelectable(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	SelectionIndexPaths() foundation.Set
	SetSelectionIndexPaths(value foundation.Set)
	IsFirstResponder() bool
	SelectionIndexes() foundation.IndexSet
	SetSelectionIndexes(value foundation.IndexSet)
}

type NSCollectionView struct {
	NSView
}

func MakeCollectionView(ptr unsafe.Pointer) NSCollectionView {
	return NSCollectionView{
		NSView: MakeView(ptr),
	}
}

func AllocCollectionView() NSCollectionView {
	return MakeCollectionView(C.C_CollectionView_Alloc())
}

func (n NSCollectionView) InitWithFrame(frameRect foundation.Rect) CollectionView {
	result_ := C.C_NSCollectionView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeCollectionView(result_)
}

func (n NSCollectionView) InitWithCoder(coder foundation.Coder) CollectionView {
	result_ := C.C_NSCollectionView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCollectionView(result_)
}

func (n NSCollectionView) Init() CollectionView {
	result_ := C.C_NSCollectionView_Init(n.Ptr())
	return MakeCollectionView(result_)
}

func (n NSCollectionView) MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IndexPath) CollectionViewItem {
	result_ := C.C_NSCollectionView_MakeItemWithIdentifier_ForIndexPath(n.Ptr(), foundation.NewString(string(identifier)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionView) RegisterNib_ForItemWithIdentifier(nib Nib, identifier UserInterfaceItemIdentifier) {
	C.C_NSCollectionView_RegisterNib_ForItemWithIdentifier(n.Ptr(), objc.ExtractPtr(nib), foundation.NewString(string(identifier)).Ptr())
}

func (n NSCollectionView) MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IndexPath) View {
	result_ := C.C_NSCollectionView_MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), foundation.NewString(string(identifier)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeView(result_)
}

func (n NSCollectionView) RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(nib Nib, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier) {
	C.C_NSCollectionView_RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(n.Ptr(), objc.ExtractPtr(nib), foundation.NewString(string(kind)).Ptr(), foundation.NewString(string(identifier)).Ptr())
}

func (n NSCollectionView) ReloadData() {
	C.C_NSCollectionView_ReloadData(n.Ptr())
}

func (n NSCollectionView) ReloadSections(sections foundation.IndexSet) {
	C.C_NSCollectionView_ReloadSections(n.Ptr(), objc.ExtractPtr(sections))
}

func (n NSCollectionView) ReloadItemsAtIndexPaths(indexPaths foundation.Set) {
	C.C_NSCollectionView_ReloadItemsAtIndexPaths(n.Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionView) NumberOfItemsInSection(section int) int {
	result_ := C.C_NSCollectionView_NumberOfItemsInSection(n.Ptr(), C.int(section))
	return int(result_)
}

func (n NSCollectionView) InsertItemsAtIndexPaths(indexPaths foundation.Set) {
	C.C_NSCollectionView_InsertItemsAtIndexPaths(n.Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionView) MoveItemAtIndexPath_ToIndexPath(indexPath foundation.IndexPath, newIndexPath foundation.IndexPath) {
	C.C_NSCollectionView_MoveItemAtIndexPath_ToIndexPath(n.Ptr(), objc.ExtractPtr(indexPath), objc.ExtractPtr(newIndexPath))
}

func (n NSCollectionView) DeleteItemsAtIndexPaths(indexPaths foundation.Set) {
	C.C_NSCollectionView_DeleteItemsAtIndexPaths(n.Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionView) InsertSections(sections foundation.IndexSet) {
	C.C_NSCollectionView_InsertSections(n.Ptr(), objc.ExtractPtr(sections))
}

func (n NSCollectionView) MoveSection_ToSection(section int, newSection int) {
	C.C_NSCollectionView_MoveSection_ToSection(n.Ptr(), C.int(section), C.int(newSection))
}

func (n NSCollectionView) DeleteSections(sections foundation.IndexSet) {
	C.C_NSCollectionView_DeleteSections(n.Ptr(), objc.ExtractPtr(sections))
}

func (n NSCollectionView) ToggleSectionCollapse(sender objc.Object) {
	C.C_NSCollectionView_ToggleSectionCollapse(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCollectionView) SelectAll(sender objc.Object) {
	C.C_NSCollectionView_SelectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCollectionView) DeselectAll(sender objc.Object) {
	C.C_NSCollectionView_DeselectAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSCollectionView) SelectItemsAtIndexPaths_ScrollPosition(indexPaths foundation.Set, scrollPosition CollectionViewScrollPosition) {
	C.C_NSCollectionView_SelectItemsAtIndexPaths_ScrollPosition(n.Ptr(), objc.ExtractPtr(indexPaths), C.uint(uint(scrollPosition)))
}

func (n NSCollectionView) DeselectItemsAtIndexPaths(indexPaths foundation.Set) {
	C.C_NSCollectionView_DeselectItemsAtIndexPaths(n.Ptr(), objc.ExtractPtr(indexPaths))
}

func (n NSCollectionView) VisibleItems() []CollectionViewItem {
	result_ := C.C_NSCollectionView_VisibleItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionViewItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionViewItem(r)
	}
	return goResult_
}

func (n NSCollectionView) IndexPathsForVisibleItems() foundation.Set {
	result_ := C.C_NSCollectionView_IndexPathsForVisibleItems(n.Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	result_ := C.C_NSCollectionView_IndexPathsForVisibleSupplementaryElementsOfKind(n.Ptr(), foundation.NewString(string(elementKind)).Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionView) IndexPathForItem(item CollectionViewItem) foundation.IndexPath {
	result_ := C.C_NSCollectionView_IndexPathForItem(n.Ptr(), objc.ExtractPtr(item))
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionView) IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath {
	result_ := C.C_NSCollectionView_IndexPathForItemAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.MakeIndexPath(result_)
}

func (n NSCollectionView) ItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewItem {
	result_ := C.C_NSCollectionView_ItemAtIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionView) SupplementaryViewForElementKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) View {
	result_ := C.C_NSCollectionView_SupplementaryViewForElementKind_AtIndexPath(n.Ptr(), foundation.NewString(string(elementKind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeView(result_)
}

func (n NSCollectionView) ScrollToItemsAtIndexPaths_ScrollPosition(indexPaths foundation.Set, scrollPosition CollectionViewScrollPosition) {
	C.C_NSCollectionView_ScrollToItemsAtIndexPaths_ScrollPosition(n.Ptr(), objc.ExtractPtr(indexPaths), C.uint(uint(scrollPosition)))
}

func (n NSCollectionView) LayoutAttributesForItemAtIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionView_LayoutAttributesForItemAtIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionView) LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	result_ := C.C_NSCollectionView_LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(n.Ptr(), foundation.NewString(string(kind)).Ptr(), objc.ExtractPtr(indexPath))
	return MakeCollectionViewLayoutAttributes(result_)
}

func (n NSCollectionView) ItemAtIndex(index uint) CollectionViewItem {
	result_ := C.C_NSCollectionView_ItemAtIndex(n.Ptr(), C.uint(index))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionView) FrameForItemAtIndex(index uint) foundation.Rect {
	result_ := C.C_NSCollectionView_FrameForItemAtIndex(n.Ptr(), C.uint(index))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionView) FrameForItemAtIndex_WithNumberOfItems(index uint, numberOfItems uint) foundation.Rect {
	result_ := C.C_NSCollectionView_FrameForItemAtIndex_WithNumberOfItems(n.Ptr(), C.uint(index), C.uint(numberOfItems))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionView) SetDraggingSourceOperationMask_ForLocal(dragOperationMask DragOperation, localDestination bool) {
	C.C_NSCollectionView_SetDraggingSourceOperationMask_ForLocal(n.Ptr(), C.uint(uint(dragOperationMask)), C.bool(localDestination))
}

func (n NSCollectionView) DataSource() objc.Object {
	result_ := C.C_NSCollectionView_DataSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCollectionView) SetDataSource(value objc.Object) {
	C.C_NSCollectionView_SetDataSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) Delegate() objc.Object {
	result_ := C.C_NSCollectionView_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCollectionView) SetDelegate(value objc.Object) {
	C.C_NSCollectionView_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) Content() []objc.Object {
	result_ := C.C_NSCollectionView_Content(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}

func (n NSCollectionView) SetContent(value []objc.Object) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionView_SetContent(n.Ptr(), cValue)
}

func (n NSCollectionView) BackgroundView() View {
	result_ := C.C_NSCollectionView_BackgroundView(n.Ptr())
	return MakeView(result_)
}

func (n NSCollectionView) SetBackgroundView(value View) {
	C.C_NSCollectionView_SetBackgroundView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) BackgroundColors() []Color {
	result_ := C.C_NSCollectionView_BackgroundColors(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Color, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeColor(r)
	}
	return goResult_
}

func (n NSCollectionView) SetBackgroundColors(value []Color) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionView_SetBackgroundColors(n.Ptr(), cValue)
}

func (n NSCollectionView) BackgroundViewScrollsWithContent() bool {
	result_ := C.C_NSCollectionView_BackgroundViewScrollsWithContent(n.Ptr())
	return bool(result_)
}

func (n NSCollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	C.C_NSCollectionView_SetBackgroundViewScrollsWithContent(n.Ptr(), C.bool(value))
}

func (n NSCollectionView) CollectionViewLayout() CollectionViewLayout {
	result_ := C.C_NSCollectionView_CollectionViewLayout(n.Ptr())
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionView) SetCollectionViewLayout(value CollectionViewLayout) {
	C.C_NSCollectionView_SetCollectionViewLayout(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) PrefetchDataSource() objc.Object {
	result_ := C.C_NSCollectionView_PrefetchDataSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSCollectionView) SetPrefetchDataSource(value objc.Object) {
	C.C_NSCollectionView_SetPrefetchDataSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) NumberOfSections() int {
	result_ := C.C_NSCollectionView_NumberOfSections(n.Ptr())
	return int(result_)
}

func (n NSCollectionView) IsSelectable() bool {
	result_ := C.C_NSCollectionView_IsSelectable(n.Ptr())
	return bool(result_)
}

func (n NSCollectionView) SetSelectable(value bool) {
	C.C_NSCollectionView_SetSelectable(n.Ptr(), C.bool(value))
}

func (n NSCollectionView) AllowsMultipleSelection() bool {
	result_ := C.C_NSCollectionView_AllowsMultipleSelection(n.Ptr())
	return bool(result_)
}

func (n NSCollectionView) SetAllowsMultipleSelection(value bool) {
	C.C_NSCollectionView_SetAllowsMultipleSelection(n.Ptr(), C.bool(value))
}

func (n NSCollectionView) AllowsEmptySelection() bool {
	result_ := C.C_NSCollectionView_AllowsEmptySelection(n.Ptr())
	return bool(result_)
}

func (n NSCollectionView) SetAllowsEmptySelection(value bool) {
	C.C_NSCollectionView_SetAllowsEmptySelection(n.Ptr(), C.bool(value))
}

func (n NSCollectionView) SelectionIndexPaths() foundation.Set {
	result_ := C.C_NSCollectionView_SelectionIndexPaths(n.Ptr())
	return foundation.MakeSet(result_)
}

func (n NSCollectionView) SetSelectionIndexPaths(value foundation.Set) {
	C.C_NSCollectionView_SetSelectionIndexPaths(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionView) IsFirstResponder() bool {
	result_ := C.C_NSCollectionView_IsFirstResponder(n.Ptr())
	return bool(result_)
}

func (n NSCollectionView) SelectionIndexes() foundation.IndexSet {
	result_ := C.C_NSCollectionView_SelectionIndexes(n.Ptr())
	return foundation.MakeIndexSet(result_)
}

func (n NSCollectionView) SetSelectionIndexes(value foundation.IndexSet) {
	C.C_NSCollectionView_SetSelectionIndexes(n.Ptr(), objc.ExtractPtr(value))
}
