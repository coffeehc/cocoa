package appkit

// #include "collection_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
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

func (n NSCollectionView) InitWithFrame(frameRect foundation.Rect) NSCollectionView {
	result_ := C.C_NSCollectionView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeCollectionView(result_)
}

func (n NSCollectionView) InitWithCoder(coder foundation.Coder) NSCollectionView {
	result_ := C.C_NSCollectionView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCollectionView(result_)
}

func (n NSCollectionView) Init() NSCollectionView {
	result_ := C.C_NSCollectionView_Init(n.Ptr())
	return MakeCollectionView(result_)
}

func AllocCollectionView() NSCollectionView {
	result_ := C.C_NSCollectionView_AllocCollectionView()
	return MakeCollectionView(result_)
}

func NewCollectionView() NSCollectionView {
	result_ := C.C_NSCollectionView_NewCollectionView()
	return MakeCollectionView(result_)
}

func (n NSCollectionView) Autorelease() NSCollectionView {
	result_ := C.C_NSCollectionView_Autorelease(n.Ptr())
	return MakeCollectionView(result_)
}

func (n NSCollectionView) Retain() NSCollectionView {
	result_ := C.C_NSCollectionView_Retain(n.Ptr())
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

type CollectionViewItem interface {
	ViewController
	ImageView() ImageView
	SetImageView(value ImageView)
	TextField() TextField
	SetTextField(value TextField)
	IsSelected() bool
	SetSelected(value bool)
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
}

type NSCollectionViewItem struct {
	NSViewController
}

func MakeCollectionViewItem(ptr unsafe.Pointer) NSCollectionViewItem {
	return NSCollectionViewItem{
		NSViewController: MakeViewController(ptr),
	}
}

func (n NSCollectionViewItem) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) InitWithCoder(coder foundation.Coder) NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) Init() NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_Init(n.Ptr())
	return MakeCollectionViewItem(result_)
}

func AllocCollectionViewItem() NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_AllocCollectionViewItem()
	return MakeCollectionViewItem(result_)
}

func NewCollectionViewItem() NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_NewCollectionViewItem()
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) Autorelease() NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_Autorelease(n.Ptr())
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) Retain() NSCollectionViewItem {
	result_ := C.C_NSCollectionViewItem_Retain(n.Ptr())
	return MakeCollectionViewItem(result_)
}

func (n NSCollectionViewItem) ImageView() ImageView {
	result_ := C.C_NSCollectionViewItem_ImageView(n.Ptr())
	return MakeImageView(result_)
}

func (n NSCollectionViewItem) SetImageView(value ImageView) {
	C.C_NSCollectionViewItem_SetImageView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionViewItem) TextField() TextField {
	result_ := C.C_NSCollectionViewItem_TextField(n.Ptr())
	return MakeTextField(result_)
}

func (n NSCollectionViewItem) SetTextField(value TextField) {
	C.C_NSCollectionViewItem_SetTextField(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionViewItem) IsSelected() bool {
	result_ := C.C_NSCollectionViewItem_IsSelected(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewItem) SetSelected(value bool) {
	C.C_NSCollectionViewItem_SetSelected(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	result_ := C.C_NSCollectionViewItem_HighlightState(n.Ptr())
	return CollectionViewItemHighlightState(int(result_))
}

func (n NSCollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	C.C_NSCollectionViewItem_SetHighlightState(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionViewItem) CollectionView() CollectionView {
	result_ := C.C_NSCollectionViewItem_CollectionView(n.Ptr())
	return MakeCollectionView(result_)
}

func (n NSCollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	result_ := C.C_NSCollectionViewItem_DraggingImageComponents(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}

type CollectionViewDataSource interface {
	HasNumberOfSectionsInCollectionView() bool
	NumberOfSectionsInCollectionView(collectionView CollectionView) int
	CollectionView_NumberOfItemsInSection(collectionView CollectionView, section int) int
	CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) CollectionViewItem
	HasCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool
	CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) View
}

func CollectionViewDataSourceToObjc(protocol CollectionViewDataSource) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapCollectionViewDataSource(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewDataSource_NumberOfSectionsInCollectionView
func collectionViewDataSource_NumberOfSectionsInCollectionView(hp C.uintptr_t, collectionView unsafe.Pointer) C.int {
	protocol := cgo.Handle(hp).Value().(CollectionViewDataSource)
	result := protocol.NumberOfSectionsInCollectionView(MakeCollectionView(collectionView))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_NumberOfItemsInSection
func collectionViewDataSource_CollectionView_NumberOfItemsInSection(hp C.uintptr_t, collectionView unsafe.Pointer, section C.int) C.int {
	protocol := cgo.Handle(hp).Value().(CollectionViewDataSource)
	result := protocol.CollectionView_NumberOfItemsInSection(MakeCollectionView(collectionView), int(section))
	return C.int(result)
}

//export collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath
func collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewDataSource)
	result := protocol.CollectionView_ItemForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath
func collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewDataSource)
	result := protocol.CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(MakeCollectionView(collectionView), CollectionViewSupplementaryElementKind(foundation.MakeString(kind).String()), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export CollectionViewDataSource_RespondsTo
func CollectionViewDataSource_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(CollectionViewDataSource)
	_ = protocol
	switch selName {
	case "numberOfSectionsInCollectionView:":
		return protocol.HasNumberOfSectionsInCollectionView()
	case "collectionView:numberOfItemsInSection:":
		return true
	case "collectionView:itemForRepresentedObjectAtIndexPath:":
		return true
	case "collectionView:viewForSupplementaryElementOfKind:atIndexPath:":
		return protocol.HasCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath()
	default:
		return false
	}
}

type CollectionViewDelegate struct {
	CollectionView_ShouldSelectItemsAtIndexPaths                                  func(collectionView CollectionView, indexPaths foundation.Set) foundation.Set
	CollectionView_DidSelectItemsAtIndexPaths                                     func(collectionView CollectionView, indexPaths foundation.Set)
	CollectionView_ShouldDeselectItemsAtIndexPaths                                func(collectionView CollectionView, indexPaths foundation.Set) foundation.Set
	CollectionView_DidDeselectItemsAtIndexPaths                                   func(collectionView CollectionView, indexPaths foundation.Set)
	CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState                 func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.Set
	CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState                    func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath                func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath           func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath        func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	CollectionView_TransitionLayoutForOldLayout_NewLayout                         func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) CollectionViewTransitionLayout
	CollectionView_CanDragItemsAtIndexPaths_WithEvent                             func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	CollectionView_PasteboardWriterForItemAtIndexPath                             func(collectionView CollectionView, indexPath foundation.IndexPath) objc.Object
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths          func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	CollectionView_DraggingSession_EndedAtPoint_DragOperation                     func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	CollectionView_UpdateDraggingItemsForDrag                                     func(collectionView CollectionView, draggingInfo objc.Object)
	CollectionView_AcceptDrop_IndexPath_DropOperation                             func(collectionView CollectionView, draggingInfo objc.Object, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	CollectionView_CanDragItemsAtIndexes_WithEvent                                func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	CollectionView_PasteboardWriterForItemAtIndex                                 func(collectionView CollectionView, index uint) objc.Object
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes             func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	CollectionView_AcceptDrop_Index_DropOperation                                 func(collectionView CollectionView, draggingInfo objc.Object, index int, dropOperation CollectionViewDropOperation) bool
}

func (delegate *CollectionViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapCollectionViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_ShouldSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DidSelectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidSelectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_ShouldDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DidDeselectItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidDeselectItemsAtIndexPaths(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegate_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState
func collectionViewDelegate_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, highlightState C.int) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), CollectionViewItemHighlightState(int(highlightState)))
}

//export collectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegate_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath
func collectionViewDelegate_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, item unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(MakeCollectionView(collectionView), MakeCollectionViewItem(item), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath
func collectionViewDelegate_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath
func collectionViewDelegate_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, view unsafe.Pointer, elementKind unsafe.Pointer, indexPath unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(MakeCollectionView(collectionView), MakeView(view), CollectionViewSupplementaryElementKind(foundation.MakeString(elementKind).String()), foundation.MakeIndexPath(indexPath))
}

//export collectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout
func collectionViewDelegate_CollectionView_TransitionLayoutForOldLayout_NewLayout(hp C.uintptr_t, collectionView unsafe.Pointer, fromLayout unsafe.Pointer, toLayout unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout(MakeCollectionView(collectionView), MakeCollectionViewLayout(fromLayout), MakeCollectionViewLayout(toLayout))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent
func collectionViewDelegate_CollectionView_CanDragItemsAtIndexPaths_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexPaths unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent(MakeCollectionView(collectionView), foundation.MakeSet(indexPaths), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath
func collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndexPath(hp C.uintptr_t, collectionView unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndexPath(MakeCollectionView(collectionView), foundation.MakeIndexPath(indexPath))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths
func collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexPaths unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeSet(indexPaths))
}

//export collectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation
func collectionViewDelegate_CollectionView_DraggingSession_EndedAtPoint_DragOperation(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, operation C.uint) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), DragOperation(uint(operation)))
}

//export collectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag
func collectionViewDelegate_CollectionView_UpdateDraggingItemsForDrag(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_UpdateDraggingItemsForDrag(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo))
}

//export collectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation
func collectionViewDelegate_CollectionView_AcceptDrop_IndexPath_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, indexPath unsafe.Pointer, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_IndexPath_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), foundation.MakeIndexPath(indexPath), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent
func collectionViewDelegate_CollectionView_CanDragItemsAtIndexes_WithEvent(hp C.uintptr_t, collectionView unsafe.Pointer, indexes unsafe.Pointer, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_CanDragItemsAtIndexes_WithEvent(MakeCollectionView(collectionView), foundation.MakeIndexSet(indexes), MakeEvent(event))
	return C.bool(result)
}

//export collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex
func collectionViewDelegate_CollectionView_PasteboardWriterForItemAtIndex(hp C.uintptr_t, collectionView unsafe.Pointer, index C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_PasteboardWriterForItemAtIndex(MakeCollectionView(collectionView), uint(index))
	return objc.ExtractPtr(result)
}

//export collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes
func collectionViewDelegate_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(hp C.uintptr_t, collectionView unsafe.Pointer, session unsafe.Pointer, screenPoint C.CGPoint, indexes unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(MakeCollectionView(collectionView), MakeDraggingSession(session), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&screenPoint))), foundation.MakeIndexSet(indexes))
}

//export collectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation
func collectionViewDelegate_CollectionView_AcceptDrop_Index_DropOperation(hp C.uintptr_t, collectionView unsafe.Pointer, draggingInfo unsafe.Pointer, index C.int, dropOperation C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	result := delegate.CollectionView_AcceptDrop_Index_DropOperation(MakeCollectionView(collectionView), objc.MakeObject(draggingInfo), int(index), CollectionViewDropOperation(int(dropOperation)))
	return C.bool(result)
}

//export CollectionViewDelegate_RespondsTo
func CollectionViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*CollectionViewDelegate)
	switch selName {
	case "collectionView:shouldSelectItemsAtIndexPaths:":
		return delegate.CollectionView_ShouldSelectItemsAtIndexPaths != nil
	case "collectionView:didSelectItemsAtIndexPaths:":
		return delegate.CollectionView_DidSelectItemsAtIndexPaths != nil
	case "collectionView:shouldDeselectItemsAtIndexPaths:":
		return delegate.CollectionView_ShouldDeselectItemsAtIndexPaths != nil
	case "collectionView:didDeselectItemsAtIndexPaths:":
		return delegate.CollectionView_DidDeselectItemsAtIndexPaths != nil
	case "collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:":
		return delegate.CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState != nil
	case "collectionView:didChangeItemsAtIndexPaths:toHighlightState:":
		return delegate.CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState != nil
	case "collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:":
		return delegate.CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath != nil
	case "collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:":
		return delegate.CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath != nil
	case "collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:":
		return delegate.CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath != nil
	case "collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:":
		return delegate.CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath != nil
	case "collectionView:transitionLayoutForOldLayout:newLayout:":
		return delegate.CollectionView_TransitionLayoutForOldLayout_NewLayout != nil
	case "collectionView:canDragItemsAtIndexPaths:withEvent:":
		return delegate.CollectionView_CanDragItemsAtIndexPaths_WithEvent != nil
	case "collectionView:pasteboardWriterForItemAtIndexPath:":
		return delegate.CollectionView_PasteboardWriterForItemAtIndexPath != nil
	case "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:":
		return delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths != nil
	case "collectionView:draggingSession:endedAtPoint:dragOperation:":
		return delegate.CollectionView_DraggingSession_EndedAtPoint_DragOperation != nil
	case "collectionView:updateDraggingItemsForDrag:":
		return delegate.CollectionView_UpdateDraggingItemsForDrag != nil
	case "collectionView:acceptDrop:indexPath:dropOperation:":
		return delegate.CollectionView_AcceptDrop_IndexPath_DropOperation != nil
	case "collectionView:canDragItemsAtIndexes:withEvent:":
		return delegate.CollectionView_CanDragItemsAtIndexes_WithEvent != nil
	case "collectionView:pasteboardWriterForItemAtIndex:":
		return delegate.CollectionView_PasteboardWriterForItemAtIndex != nil
	case "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:":
		return delegate.CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes != nil
	case "collectionView:acceptDrop:index:dropOperation:":
		return delegate.CollectionView_AcceptDrop_Index_DropOperation != nil
	default:
		return false
	}
}

type CollectionViewSectionHeaderView interface {
	HasPrepareForReuse() bool
	PrepareForReuse()
	HasPreferredLayoutAttributesFittingAttributes() bool
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	HasApplyLayoutAttributes() bool
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	HasWillTransitionFromLayout_ToLayout() bool
	WillTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasDidTransitionFromLayout_ToLayout() bool
	DidTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasSectionCollapseButton() bool
	SectionCollapseButton() Button
	SetSectionCollapseButton(value Button)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
}

func CollectionViewSectionHeaderViewToObjc(protocol CollectionViewSectionHeaderView) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapCollectionViewSectionHeaderView(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewSectionHeaderView_PrepareForReuse
func collectionViewSectionHeaderView_PrepareForReuse(hp C.uintptr_t) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.PrepareForReuse()
}

//export collectionViewSectionHeaderView_PreferredLayoutAttributesFittingAttributes
func collectionViewSectionHeaderView_PreferredLayoutAttributesFittingAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	result := protocol.PreferredLayoutAttributesFittingAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
	return objc.ExtractPtr(result)
}

//export collectionViewSectionHeaderView_ApplyLayoutAttributes
func collectionViewSectionHeaderView_ApplyLayoutAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.ApplyLayoutAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
}

//export collectionViewSectionHeaderView_WillTransitionFromLayout_ToLayout
func collectionViewSectionHeaderView_WillTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.WillTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewSectionHeaderView_DidTransitionFromLayout_ToLayout
func collectionViewSectionHeaderView_DidTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.DidTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewSectionHeaderView_SetSectionCollapseButton
func collectionViewSectionHeaderView_SetSectionCollapseButton(hp C.uintptr_t, value unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.SetSectionCollapseButton(MakeButton(value))
}

//export collectionViewSectionHeaderView_SectionCollapseButton
func collectionViewSectionHeaderView_SectionCollapseButton(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	result := protocol.SectionCollapseButton()
	return objc.ExtractPtr(result)
}

//export collectionViewSectionHeaderView_SetIdentifier
func collectionViewSectionHeaderView_SetIdentifier(hp C.uintptr_t, value unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	protocol.SetIdentifier(UserInterfaceItemIdentifier(foundation.MakeString(value).String()))
}

//export collectionViewSectionHeaderView_Identifier
func collectionViewSectionHeaderView_Identifier(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	result := protocol.Identifier()
	return foundation.NewString(string(result)).Ptr()
}

//export CollectionViewSectionHeaderView_RespondsTo
func CollectionViewSectionHeaderView_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(CollectionViewSectionHeaderView)
	_ = protocol
	switch selName {
	case "prepareForReuse":
		return protocol.HasPrepareForReuse()
	case "preferredLayoutAttributesFittingAttributes:":
		return protocol.HasPreferredLayoutAttributesFittingAttributes()
	case "applyLayoutAttributes:":
		return protocol.HasApplyLayoutAttributes()
	case "willTransitionFromLayout:toLayout:":
		return protocol.HasWillTransitionFromLayout_ToLayout()
	case "didTransitionFromLayout:toLayout:":
		return protocol.HasDidTransitionFromLayout_ToLayout()
	case "setSectionCollapseButton:":
		fallthrough
	case "sectionCollapseButton":
		return protocol.HasSectionCollapseButton()
	case "setIdentifier:":
		fallthrough
	case "identifier":
		return true
	default:
		return false
	}
}

type CollectionViewElement interface {
	HasPrepareForReuse() bool
	PrepareForReuse()
	HasPreferredLayoutAttributesFittingAttributes() bool
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	HasApplyLayoutAttributes() bool
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	HasWillTransitionFromLayout_ToLayout() bool
	WillTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasDidTransitionFromLayout_ToLayout() bool
	DidTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
}

func CollectionViewElementToObjc(protocol CollectionViewElement) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapCollectionViewElement(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export collectionViewElement_PrepareForReuse
func collectionViewElement_PrepareForReuse(hp C.uintptr_t) {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	protocol.PrepareForReuse()
}

//export collectionViewElement_PreferredLayoutAttributesFittingAttributes
func collectionViewElement_PreferredLayoutAttributesFittingAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	result := protocol.PreferredLayoutAttributesFittingAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
	return objc.ExtractPtr(result)
}

//export collectionViewElement_ApplyLayoutAttributes
func collectionViewElement_ApplyLayoutAttributes(hp C.uintptr_t, layoutAttributes unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	protocol.ApplyLayoutAttributes(MakeCollectionViewLayoutAttributes(layoutAttributes))
}

//export collectionViewElement_WillTransitionFromLayout_ToLayout
func collectionViewElement_WillTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	protocol.WillTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewElement_DidTransitionFromLayout_ToLayout
func collectionViewElement_DidTransitionFromLayout_ToLayout(hp C.uintptr_t, oldLayout unsafe.Pointer, newLayout unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	protocol.DidTransitionFromLayout_ToLayout(MakeCollectionViewLayout(oldLayout), MakeCollectionViewLayout(newLayout))
}

//export collectionViewElement_SetIdentifier
func collectionViewElement_SetIdentifier(hp C.uintptr_t, value unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	protocol.SetIdentifier(UserInterfaceItemIdentifier(foundation.MakeString(value).String()))
}

//export collectionViewElement_Identifier
func collectionViewElement_Identifier(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	result := protocol.Identifier()
	return foundation.NewString(string(result)).Ptr()
}

//export CollectionViewElement_RespondsTo
func CollectionViewElement_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(CollectionViewElement)
	_ = protocol
	switch selName {
	case "prepareForReuse":
		return protocol.HasPrepareForReuse()
	case "preferredLayoutAttributesFittingAttributes:":
		return protocol.HasPreferredLayoutAttributesFittingAttributes()
	case "applyLayoutAttributes:":
		return protocol.HasApplyLayoutAttributes()
	case "willTransitionFromLayout:toLayout:":
		return protocol.HasWillTransitionFromLayout_ToLayout()
	case "didTransitionFromLayout:toLayout:":
		return protocol.HasDidTransitionFromLayout_ToLayout()
	case "setIdentifier:":
		fallthrough
	case "identifier":
		return true
	default:
		return false
	}
}
