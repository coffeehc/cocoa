package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "tab_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TabView interface {
	View
	NumberOfTabViewItems() int
	TabViewType() TabViewType
	SetTabViewType(tabViewType TabViewType)
	TabPosition() TabPosition
	SetTabPosition(tabPosition TabPosition)
	TabViewBorderType() TabViewBorderType
	SetTabViewBorderType(tabViewBorderType TabViewBorderType)
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(allowsTruncatedLabels bool)
	SelectedTabViewItem() TabViewItem
	Font() Font
	SetFont(font Font)
	MinimumSize() foundation.Size
	ControlSize() ControlSize
	SetControlSize(controlSize ControlSize)
	AddTabViewItem(tabViewItem TabViewItem)
	InsertTabViewItem(tabViewItem TabViewItem, index int)
	RemoveTabViewItem(tabViewItem TabViewItem)
	IndexOfTabViewItem(tabViewItem TabViewItem) int
	TabViewItemAtIndex(index int) TabViewItem
	SelectFirstTabViewItem(sender objc.Object)
	SelectLastTabViewItem(sender objc.Object)
	SelectNextTabViewItem(sender objc.Object)
	SelectPreviousTabViewItem(sender objc.Object)
	SelectTabViewItem(tabViewItem TabViewItem)
	SelectTabViewItemAtIndex(index int)
	TabViewDidChangeNumberOfTabViewItems(callback func(view TabView))
	_tabViewDidChangeNumberOfTabViewItems() func(view TabView)
	ShouldSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem) bool)
	_shouldSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) bool
	WillSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem))
	_willSelectTabViewItem() func(view TabView, tabViewItem TabViewItem)
	DidSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem))
	_didSelectTabViewItem() func(view TabView, tabViewItem TabViewItem)
}

var _ TabView = (*NSTabView)(nil)

type NSTabView struct {
	NSView
	tabViewDidChangeNumberOfTabViewItems func(view TabView)
	shouldSelectTabViewItem              func(view TabView, tabViewItem TabViewItem) bool
	willSelectTabViewItem                func(view TabView, tabViewItem TabViewItem)
	didSelectTabViewItem                 func(view TabView, tabViewItem TabViewItem)
}

func MakeTabView(ptr unsafe.Pointer) *NSTabView {
	if ptr == nil {
		return nil
	}
	return &NSTabView{
		NSView: *MakeView(ptr),
	}
}
func (t *NSTabView) setDelegate() {
	id := resources.Register(t)
	C.TabView_SetDelegate(t.Ptr(), C.long(id))
}

func (t *NSTabView) NumberOfTabViewItems() int {
	return int(C.TabView_NumberOfTabViewItems(t.Ptr()))
}

func (t *NSTabView) TabViewType() TabViewType {
	return TabViewType(C.TabView_TabViewType(t.Ptr()))
}

func (t *NSTabView) SetTabViewType(tabViewType TabViewType) {
	C.TabView_SetTabViewType(t.Ptr(), C.ulong(tabViewType))
}

func (t *NSTabView) TabPosition() TabPosition {
	return TabPosition(C.TabView_TabPosition(t.Ptr()))
}

func (t *NSTabView) SetTabPosition(tabPosition TabPosition) {
	C.TabView_SetTabPosition(t.Ptr(), C.ulong(tabPosition))
}

func (t *NSTabView) TabViewBorderType() TabViewBorderType {
	return TabViewBorderType(C.TabView_TabViewBorderType(t.Ptr()))
}

func (t *NSTabView) SetTabViewBorderType(tabViewBorderType TabViewBorderType) {
	C.TabView_SetTabViewBorderType(t.Ptr(), C.ulong(tabViewBorderType))
}

func (t *NSTabView) AllowsTruncatedLabels() bool {
	return bool(C.TabView_AllowsTruncatedLabels(t.Ptr()))
}

func (t *NSTabView) SetAllowsTruncatedLabels(allowsTruncatedLabels bool) {
	C.TabView_SetAllowsTruncatedLabels(t.Ptr(), C.bool(allowsTruncatedLabels))
}

func (t *NSTabView) SelectedTabViewItem() TabViewItem {
	return MakeTabViewItem(C.TabView_SelectedTabViewItem(t.Ptr()))
}

func (t *NSTabView) Font() Font {
	return MakeFont(C.TabView_Font(t.Ptr()))
}

func (t *NSTabView) SetFont(font Font) {
	C.TabView_SetFont(t.Ptr(), toPointer(font))
}

func (t *NSTabView) MinimumSize() foundation.Size {
	return toSize(C.TabView_MinimumSize(t.Ptr()))
}

func (t *NSTabView) ControlSize() ControlSize {
	return ControlSize(C.TabView_ControlSize(t.Ptr()))
}

func (t *NSTabView) SetControlSize(controlSize ControlSize) {
	C.TabView_SetControlSize(t.Ptr(), C.ulong(controlSize))
}

func NewTabView(frame foundation.Rect) TabView {
	instance := MakeTabView(C.TabView_NewTabView(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (t *NSTabView) AddTabViewItem(tabViewItem TabViewItem) {
	C.TabView_AddTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) InsertTabViewItem(tabViewItem TabViewItem, index int) {
	C.TabView_InsertTabViewItem(t.Ptr(), toPointer(tabViewItem), C.long(index))
}

func (t *NSTabView) RemoveTabViewItem(tabViewItem TabViewItem) {
	C.TabView_RemoveTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) IndexOfTabViewItem(tabViewItem TabViewItem) int {
	return int(C.TabView_IndexOfTabViewItem(t.Ptr(), toPointer(tabViewItem)))
}

func (t *NSTabView) TabViewItemAtIndex(index int) TabViewItem {
	return MakeTabViewItem(C.TabView_TabViewItemAtIndex(t.Ptr(), C.long(index)))
}

func (t *NSTabView) SelectFirstTabViewItem(sender objc.Object) {
	C.TabView_SelectFirstTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectLastTabViewItem(sender objc.Object) {
	C.TabView_SelectLastTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectNextTabViewItem(sender objc.Object) {
	C.TabView_SelectNextTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectPreviousTabViewItem(sender objc.Object) {
	C.TabView_SelectPreviousTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectTabViewItem(tabViewItem TabViewItem) {
	C.TabView_SelectTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) SelectTabViewItemAtIndex(index int) {
	C.TabView_SelectTabViewItemAtIndex(t.Ptr(), C.long(index))
}

func (t *NSTabView) TabViewDidChangeNumberOfTabViewItems(callback func(view TabView)) {
	t.tabViewDidChangeNumberOfTabViewItems = callback
}

func (t *NSTabView) _tabViewDidChangeNumberOfTabViewItems() func(view TabView) {
	return t.tabViewDidChangeNumberOfTabViewItems
}

func (t *NSTabView) ShouldSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem) bool) {
	t.shouldSelectTabViewItem = callback
}

func (t *NSTabView) _shouldSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) bool {
	return t.shouldSelectTabViewItem
}

func (t *NSTabView) WillSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem)) {
	t.willSelectTabViewItem = callback
}

func (t *NSTabView) _willSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) {
	return t.willSelectTabViewItem
}

func (t *NSTabView) DidSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem)) {
	t.didSelectTabViewItem = callback
}

func (t *NSTabView) _didSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) {
	return t.didSelectTabViewItem
}

//export TabView_Delegate_TabViewDidChangeNumberOfTabViewItems
func TabView_Delegate_TabViewDidChangeNumberOfTabViewItems(id int64, view unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._tabViewDidChangeNumberOfTabViewItems()
	if callback != nil {
		callback(MakeTabView(view))
	}
}

//export TabView_Delegate_ShouldSelectTabViewItem
func TabView_Delegate_ShouldSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) bool {
	tabView := resources.Get(id).(TabView)
	callback := tabView._shouldSelectTabViewItem()
	if callback != nil {
		return callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
	return true
}

//export TabView_Delegate_WillSelectTabViewItem
func TabView_Delegate_WillSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._willSelectTabViewItem()
	if callback != nil {
		callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
}

//export TabView_Delegate_DidSelectTabViewItem
func TabView_Delegate_DidSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._didSelectTabViewItem()
	if callback != nil {
		callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
}
